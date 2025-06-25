package adslot

import (
	"encoding/json"
	"errors"
	"fmt"
	"github.com/jinzhu/copier"
	aGloabl "github.com/nzmxd/app-insight/global"
	"github.com/nzmxd/app-insight/model/adslot"
	"github.com/nzmxd/app-insight/model/adslot/request"
	"github.com/nzmxd/bserver/global"
	"go.uber.org/zap"
)

type AdSlotPushService struct{}
type CoreParamService[T adslot.AdSlotCoreParams] struct{}

func (s *CoreParamService[T]) Create(params T) (err error) {
	defer func() {
		if errors.Is(err, DuplicatePushAdUnitErr) {
			return
		}
		adUnitPushLog := params.GetAdUnitPushLog() // 提前拿到，避免 defer 中重复调用
		adUnitPushLog.PushStatus = 0
		if err == nil {
			adUnitPushLog.PushStatus = 1
		}
		if e := adUnitPushLogService.CreatAdUnitPushLog(adUnitPushLog); e != nil {
			global.LOG.Error("插入广告位推送日志失败", zap.Error(e))
		}
	}()

	// Step 1: 计算 MD5 并设置
	var md5 string
	if md5, err = params.GetParamsDictMd5(); err != nil {
		return err
	}
	params.SetParamsDictMd5(md5)

	// Step 2: 校验合法性
	if !params.ValidateCoreParams() {
		return InvalidAdUnitErr
	}
	// Stop 3: 根据推送日志本地去重
	exists, e := adUnitPushLogService.ExistsByParamsDictMd5(md5)
	if e != nil {
		global.LOG.Error("广告位根据推送日志去除失败", zap.Error(e))
	}
	if exists {
		return DuplicatePushAdUnitErr
	}

	// Step 4: 最终去重
	var count int64
	if err = aGloabl.SpRawDB.Model(&params).Where("params_dict_md5 = ?", md5).Count(&count).Error; err != nil {
		return err
	}
	if count > 0 {
		return DuplicateAdUnitErr
	}

	// Step 4: 执行插入
	return aGloabl.SpRawDB.Create(&params).Error
}

func (s *CoreParamService[T]) ConvertApplovin(req request.ScrapyApplovinCoreParamsRequest) (*adslot.ScrapyApplovinCoreParams, error) {
	_, sdkKeyOk := req.ParamsDict["sdk_key"]
	_, pkgNameOk := req.ParamsDict["package_name"]
	if !sdkKeyOk || !pkgNameOk {
		return nil, errors.New("params参数错误，缺少sdk_key或者package_name")
	}
	scrapyApplovinCoreParams := new(adslot.ScrapyApplovinCoreParams)
	paramsDict, _ := json.Marshal(&req.ParamsDict)
	scrapyApplovinCoreParams.ParamsId = req.ParamsId
	scrapyApplovinCoreParams.ParamsDict = string(paramsDict)
	scrapyApplovinCoreParams.ParamsDictMd5 = req.ParamsDictMd5
	scrapyApplovinCoreParams.Os = req.Os
	scrapyApplovinCoreParams.SourceApp = req.SourceApp
	scrapyApplovinCoreParams.BkInt = req.BkInt
	scrapyApplovinCoreParams.BkString = req.BkString
	scrapyApplovinCoreParams.Geo = req.Geo
	scrapyApplovinCoreParams.Lang = req.Lang
	scrapyApplovinCoreParams.Version = req.Version
	scrapyApplovinCoreParams.IsAvailable = req.IsAvailable
	return scrapyApplovinCoreParams, nil
}

func (s *CoreParamService[T]) ConvertAdmob(req request.ScrapyAdmobCoreParamsRequest) (*adslot.ScrapyAdmobCoreParams, error) {
	if err := s.checkRequiredParams(
		req.ParamsDict, "client", "slotname", "source_app", "admob_account_id",
		"os_type", "adunit_type", "format"); err != nil {
		return nil, err
	}
	scrapyParams := new(adslot.ScrapyAdmobCoreParams)
	// 自动复制字段
	if err := copier.Copy(scrapyParams, &req); err != nil {
		return nil, err
	}
	// 特殊字段处理
	paramsDictBytes, _ := json.Marshal(req.ParamsDict)
	scrapyParams.ParamsDict = string(paramsDictBytes)
	return scrapyParams, nil
}

func (s *CoreParamService[T]) checkRequiredParams(params map[string]interface{}, keys ...string) error {
	for _, key := range keys {
		if _, ok := params[key]; !ok {
			return fmt.Errorf("params参数错误，缺少%s", key)
		}
	}
	return nil
}
