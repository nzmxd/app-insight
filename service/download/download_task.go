package download

import (
	"app-insight/core/downloader"
	aGlobal "app-insight/global"
	"app-insight/model/apprank"
	"app-insight/model/download"
	"app-insight/model/download/request"
	"context"
	"database/sql"
	"encoding/json"
	"errors"
	"fmt"
	"github.com/hibiken/asynq"
	"github.com/minio/minio-go/v7"
	"github.com/nzmxd/bserver/global"
	"github.com/nzmxd/bserver/utils"
	"github.com/nzmxd/bserver/utils/upload"
	"go.uber.org/zap"
	"gorm.io/gorm"
	"os"
	"strconv"
	"strings"
	"time"
)

type DownloadTaskService struct {
}

func (d *DownloadTaskService) CreateDownloadTask(appDownloadTask *download.AppDownloadTask) (err error) {
	if appDownloadTask.AppID == nil || *appDownloadTask.AppID == "" {
		return errors.New("AppID 不能为空")
	}

	// 判断 AppID 是否已存在
	var count int64
	err = global.DB.Model(&download.AppDownloadTask{}).
		Where("app_id = ?", *appDownloadTask.AppID).
		Count(&count).Error
	if err != nil {
		return err
	}
	if count > 0 {
		// 已存在，不插入
		global.LOG.Info("该 AppID 已存在，跳过插入", zap.String("appID", *appDownloadTask.AppID))
		return nil
	}

	// 设置默认值
	if appDownloadTask.Status == nil {
		appDownloadTask.Status = utils.Ptr(download.StatusPending)
	}
	if appDownloadTask.RetryCount == nil {
		appDownloadTask.RetryCount = utils.Ptr(0)
	}

	// 插入记录
	err = global.DB.Create(appDownloadTask).Error
	return err
}

// DeleteDownloadTask 删除SDK匹配结果记录
// Author [nzmxd](https://github.com/nzmxd)
func (d *DownloadTaskService) DeleteDownloadTask(id string) (err error) {
	err = global.DB.Delete(&download.AppDownloadTask{}, "id = ?", id).Error
	return err
}

// DeleteDownloadTaskByIds 批量删除SDK匹配结果记录
// Author [nzmxd](https://github.com/nzmxd)
func (d *DownloadTaskService) DeleteDownloadTaskByIds(ids []string) (err error) {
	err = global.DB.Delete(&[]download.AppDownloadTask{}, "id in ?", ids).Error
	return err
}

// UpdateDownloadTask 更新SDK匹配结果记录
// Author [nzmxd](https://github.com/nzmxd)
func (d *DownloadTaskService) UpdateDownloadTask(appDownloadTask *download.AppDownloadTask) (err error) {
	err = global.DB.Model(&download.AppDownloadTask{}).Where("id = ?", appDownloadTask.ID).Updates(&appDownloadTask).Error
	return err
}

// GetDownloadTask 根据id获取SDK匹配结果记录
// Author [nzmxd](https://github.com/nzmxd)
func (d *DownloadTaskService) GetDownloadTask(id string) (downloadTask download.AppDownloadTask, err error) {
	err = global.DB.Where("id = ?", id).First(&downloadTask).Error
	return
}

// GetDownloadTaskByAppId 根据app_id获取SDK匹配结果记录
// Author [nzmxd](https://github.com/nzmxd)
func (d *DownloadTaskService) GetDownloadTaskByAppId(appId string) (downloadTask download.AppDownloadTask, err error) {
	err = global.DB.Where("app_id = ?", appId).First(&downloadTask).Error
	return
}

// GetDownloadStats 获取下载任务状态统计
// Author [nzmxd](https://github.com/nzmxd)
func (d *DownloadTaskService) GetDownloadStats(search request.DownloadStatsSearch) (*download.DownloadStatsResult, error) {
	var result download.DownloadStatsResult
	db := global.DB.Model(&download.AppDownloadTask{})

	// 所有匹配创建时间、开始时间、完成时间的范围（可为空）
	start := search.StartTime
	end := search.EndTime

	// 1. 等待中的任务（pending）按 CreatedAt
	if err := db.Session(&gorm.Session{}).Where("status = ?", download.StatusPending).
		Where("created_at BETWEEN ? AND ?", start, end).
		Count(&result.PendingCount).Error; err != nil {
		return nil, err
	}

	// 2. 任务队列中任务（downloading / retrying），按 StartedAt
	if err := db.Session(&gorm.Session{}).Where("status IN ?", []int{download.StatusDownloading, download.StatusRetrying}).
		Where("started_at BETWEEN ? AND ?", start, end).
		Count(&result.QueueCount).Error; err != nil {
		return nil, err
	}

	// 3. 重试的任务（retrying）
	if err := db.Session(&gorm.Session{}).Where("status = ?", download.StatusRetrying).
		Where("started_at BETWEEN ? AND ?", start, end).
		Count(&result.RetryCount).Error; err != nil {
		return nil, err
	}

	// 4. 失败的任务（failed）
	if err := db.Session(&gorm.Session{}).Where("status = ?", download.StatusFailed).
		Where("started_at BETWEEN ? AND ?", start, end).
		Count(&result.FailedCount).Error; err != nil {
		return nil, err
	}

	// 5. 下载完成任务（success），按 FinishedAt 分 GP 与非 GP 上架
	if err := db.Session(&gorm.Session{}).Where("status = ?", download.StatusSuccess).
		Where("finished_at BETWEEN ? AND ?", start, end).
		Where("is_gp_listing = ?", true).
		Count(&result.FinishedGpCount).Error; err != nil {
		return nil, err
	}

	if err := db.Session(&gorm.Session{}).Where("status = ?", download.StatusSuccess).
		Where("finished_at BETWEEN ? AND ?", start, end).
		Where("is_gp_listing = ?", false).
		Count(&result.FinishedNonGpCount).Error; err != nil {
		return nil, err
	}

	// 6. 汇总总任务数（只看创建时间）
	if err := db.Session(&gorm.Session{}).Where("created_at BETWEEN ? AND ?", start, end).
		Count(&result.TotalCount).Error; err != nil {
		return nil, err
	}

	// 7. 计算平均下载延时（单位：秒）
	var avg sql.NullFloat64
	if err := db.Session(&gorm.Session{}).
		Select("AVG(TIMESTAMPDIFF(SECOND, started_at, finished_at))").
		Where("status = ?", download.StatusSuccess).
		Where("finished_at BETWEEN ? AND ?", search.StartTime, search.EndTime).
		Scan(&avg).Error; err != nil {
		return nil, err
	}
	// 转换成 float64，如果为 NULL 则设为 0
	if avg.Valid {
		result.AvgDownloadSeconds = avg.Float64
	} else {
		result.AvgDownloadSeconds = 0
	}

	return &result, nil
}

// BatchInsertByAppId 批量插入下载任务，如果 app_id 已存在则跳过插入
// Author [nzmxd](https://github.com/nzmxd)
func (d *DownloadTaskService) BatchInsertByAppId(appIDs []string) (inserted int64, skipped int64, err error) {
	if len(appIDs) == 0 {
		return 0, 0, nil
	}

	// 去重 appID
	uniqueMap := make(map[string]struct{}, len(appIDs))
	for _, id := range appIDs {
		uniqueMap[id] = struct{}{}
	}
	var deduped []string
	for id := range uniqueMap {
		deduped = append(deduped, id)
	}

	// 查已有 app_id
	var existingIDs []string
	if err = global.DB.Model(&download.AppDownloadTask{}).
		Where("app_id IN ?", deduped).
		Pluck("app_id", &existingIDs).Error; err != nil {
		return
	}

	existingMap := make(map[string]struct{}, len(existingIDs))
	for _, id := range existingIDs {
		existingMap[id] = struct{}{}
	}

	var toInsert []download.AppDownloadTask
	for _, id := range deduped {
		if _, exists := existingMap[id]; exists {
			continue
		}
		idCopy := id
		toInsert = append(toInsert, download.AppDownloadTask{
			AppID:  &idCopy,
			Status: utils.Ptr(download.StatusPending),
		})
	}

	// 插入
	if len(toInsert) > 0 {
		err = global.DB.Create(&toInsert).Error
		if err != nil {
			return
		}
		inserted = int64(len(toInsert))
	}

	skipped = int64(len(deduped)) - inserted
	return
}

// EnqueueDownloadTask 将下载任务加入到下载队列
// Author [nzmxd](https://github.com/nzmxd)
func (d *DownloadTaskService) EnqueueDownloadTask(appDownloadTask download.AppDownloadTask) error {
	payload, err := json.Marshal(appDownloadTask)
	if err != nil {
		return err
	}
	downloadTask := asynq.NewTask(aGlobal.DownloadTask, payload, asynq.Retention(24*time.Hour))
	_, err = aGlobal.AsynqClient.Enqueue(downloadTask,
		asynq.Queue("download"),
		asynq.TaskID(*appDownloadTask.AppID),
		asynq.Timeout(15*time.Minute),
		asynq.MaxRetry(aGlobal.Config.Downloader.MaxRetry))
	return err
}

// UpdateTimeoutTasks 更新超时任务状态
// Author [nzmxd](https://github.com/nzmxd)
func (d *DownloadTaskService) UpdateTimeoutTasks() error {
	oneHourAgo := time.Now().Add(-1 * time.Hour)
	result := global.DB.Model(&download.AppDownloadTask{}).Where("status IN (?) AND updated_at < ?", []int{download.StatusQueued, download.StatusDownloading}, oneHourAgo).
		Updates(map[string]interface{}{
			"status":      0,
			"retry_count": 0,
		})
	if result.Error != nil {
		return result.Error
	}
	global.LOG.Info(fmt.Sprintf("成功更新了 %d 条超时记录", result.RowsAffected))
	return nil
}

// ProcessTask 实现Asynq任务队列处理接口，处理下载任务
func (d *DownloadTaskService) ProcessTask(ctx context.Context, t *asynq.Task) error {
	var downloadTask download.AppDownloadTask
	maxRetry, _ := asynq.GetMaxRetry(ctx)
	retryCount, _ := asynq.GetRetryCount(ctx)
	isFinalRetry := retryCount >= maxRetry

	defer func() {
		downloadTaskJson, _ := json.Marshal(downloadTask)
		_, _ = t.ResultWriter().Write(downloadTaskJson)
		if err := d.UpdateDownloadTask(&downloadTask); err != nil {
			global.LOG.Error("Failed to update download task", zap.Error(err))
		}
	}()

	// Step 1: 解码任务
	if err := json.Unmarshal(t.Payload(), &downloadTask); err != nil {
		return d.setTaskError(&downloadTask, download.StatusFailed, fmt.Errorf("unmarshal payload failed: %v", err), true)
	}

	// Step 2: 设置为下载中
	downloadTask.Status = utils.Ptr(download.StatusDownloading)
	downloadTask.Source = utils.Ptr(aGlobal.Config.Downloader.Source)
	_ = d.UpdateDownloadTask(&downloadTask) // 可异步

	// Step 3: 获取 App 详情
	appDetail := new(apprank.AppDetail)
	genericAppDetail, err := aGlobal.Downloader.GetAppDetail(*downloadTask.AppID)
	bts, _ := json.Marshal(genericAppDetail)
	_ = json.Unmarshal(bts, &appDetail)
	if err != nil {
		status := download.StatusRetrying
		if isFinalRetry || errors.Is(err, downloader.AppNotFoundErr) {
			status = download.StatusFailed
		}
		return d.setTaskError(&downloadTask, status, fmt.Errorf("get app detail failed: %w", err), errors.Is(err, downloader.AppNotFoundErr))
	}

	// Step 4: 验证 GooglePlay 上架状态
	if err = aGlobal.DownloaderList["googleplay"].Validate(appDetail.RealPackageName, ""); err != nil {
		appDetail.IsGoogleListing = false
	} else {
		appDetail.IsGoogleListing = true
	}
	_ = global.CH.Create(appDetail).Error
	// Step 5: 校验大小限制
	if sizeMB, err := strconv.Atoi(appDetail.Size); err == nil {
		limitBytes := aGlobal.Config.Downloader.LimitSize * 1024 * 1024
		if sizeMB > limitBytes {
			errMsg := fmt.Sprintf("download limit exceeded: %d > %d", sizeMB, limitBytes)
			return d.setTaskError(&downloadTask, download.StatusFailed, fmt.Errorf(errMsg), true)
		}
	} else {
		global.LOG.Error("Convert app size failed", zap.String("size", appDetail.Size), zap.Error(err))
	}

	// Step 6: 下载文件
	var filePath string
	if appDetail.DownloadURL != "" {
		filePath, err = downloader.DownloadFile(appDetail.DownloadURL, aGlobal.Config.Downloader.SavePath, aGlobal.Config.Downloader.Proxy, aGlobal.Config.Downloader.Timeout)
	} else {
		filePath, err = aGlobal.Downloader.Download(*downloadTask.AppID, "", aGlobal.Config.Downloader.SavePath)
	}

	if err != nil {
		status := download.StatusRetrying
		if isFinalRetry {
			status = download.StatusFailed
		}
		return d.setTaskError(&downloadTask, status, fmt.Errorf("download failed: %w", err), false)
	}

	if !strings.HasSuffix(filePath, "APK") {
		status := download.StatusRetrying
		return d.setTaskError(&downloadTask, status, fmt.Errorf("download failed: %w", errors.New("文件后缀检查失败")), false)
	}

	// Step 7: 成功，填充数据
	now := time.Now()
	downloadTask.Status = utils.Ptr(download.StatusSuccess)
	downloadTask.FilePath = &filePath
	downloadTask.VersionName = &appDetail.VersionName
	downloadTask.VersionCode = &appDetail.VersionCode
	downloadTask.Developer = &appDetail.Developer
	downloadTask.IsGpListing = &appDetail.IsGoogleListing
	downloadTask.FinishedAt = &now

	// Step 8: 判断是否需要上传到minio
	if aGlobal.Config.Downloader.UploadEnable {
		remoteFilePath, err := upload.MinioClient.UploadLocalFile(filePath, "android")
		if err != nil {
			global.LOG.Error("Upload local file failed", zap.String("appID", appDetail.AppID), zap.Error(err))
			return err
		}
		_ = os.Remove(filePath)
		downloadTask.FilePath = &remoteFilePath
	}
	global.LOG.Info("应用下载成功", zap.String("packageName", *downloadTask.AppID), zap.String("filePath", *downloadTask.FilePath))
	return nil
}

// setTaskError 封装错误处理
func (d *DownloadTaskService) setTaskError(task *download.AppDownloadTask, status int, err error, skipRetry bool) error {
	task.Status = utils.Ptr(status)
	msg := err.Error()
	task.ErrorMessage = &msg
	utils.IntPtrAddOne(&task.RetryCount)

	global.LOG.Error("ProcessTask error",
		zap.String("app_id", *task.AppID),
		zap.String("error", msg),
		zap.Int("status", status),
	)

	if skipRetry {
		return asynq.SkipRetry
	}
	return err
}

// GetDownloadTaskByAppIdAndVersion 根据appId和versionCode查询下载任务
func (d *DownloadTaskService) GetDownloadTaskByAppIdAndVersion(appId string, versionCode string) (downloadTask download.AppDownloadTask, err error) {
	err = global.DB.
		Where("app_id = ?", appId).
		Where("version_code = ?", versionCode).
		First(&downloadTask).Error
	return
}

// GetDownloadUrl 获取下载链接
func (d *DownloadTaskService) GetDownloadUrl(req request.DownloadUrlSearch) (string, error) {
	var task download.AppDownloadTask

	// 查找符合条件的记录
	db := global.DB.Where("app_id = ? AND version_code = ?", *req.AppID)
	if req.VersionCode != nil {
		db = db.Where("version_code = ?", *req.VersionCode)
	}
	err := db.First(&task).Error
	if err != nil {
		return "", fmt.Errorf("未找到下载任务记录: %w", err)
	}

	// 判断 file_path 是否存在并有效
	if task.FilePath != nil && *task.FilePath != "" {
		// 去掉前缀获取 MinIO 的 object key
		baseUrl := global.CONFIG.Minio.BucketUrl
		if strings.HasPrefix(*task.FilePath, baseUrl) {
			objectKey := strings.TrimPrefix(*task.FilePath, baseUrl)

			// 校验文件是否存在
			exists, err := upload.MinioClient.Client.StatObject(
				context.Background(),
				global.CONFIG.Minio.BucketName,
				strings.TrimPrefix(objectKey, "/"),
				minio.StatObjectOptions{},
			)
			if err == nil && exists.Size > 0 {
				return *task.FilePath, nil
			}
		}
	}
	var versionCode string
	// fallback 方式获取新的下载链接
	if req.VersionCode != nil {
		versionCode = *req.VersionCode
	}
	downloadUrl := aGlobal.Downloader.GetAppDownloadUrl(*req.AppID, versionCode)
	if downloadUrl == "" {
		return "", fmt.Errorf("获取下载链接失败")
	}
	return downloadUrl, nil
}
