package analysis

import (
	aGloabl "github.com/nzmxd/app-insight/global"
	"github.com/nzmxd/app-insight/model/analysis"
)

type SdkMetadataService struct{}

// CreateSdkMetadata 创建SdkMetadata表记录
// Author [yourname](https://github.com/yourname)
func (s *SdkMetadataService) CreateSdkMetadata(SdkMetadata *analysis.SdkMetadata) (err error) {
	err = aGloabl.AppRankDB.Create(SdkMetadata).Error
	return err
}

// DeleteSdkMetadata 删除SdkMetadata表记录
// Author [yourname](https://github.com/yourname)
func (s *SdkMetadataService) DeleteSdkMetadata(ID string) (err error) {
	err = aGloabl.AppRankDB.Delete(&analysis.SdkMetadata{}, "id = ?", ID).Error
	return err
}

// DeleteSdkMetadataByIds 批量删除SdkMetadata表记录
// Author [yourname](https://github.com/yourname)
func (s *SdkMetadataService) DeleteSdkMetadataByIds(IDs []string) (err error) {
	err = aGloabl.AppRankDB.Delete(&[]analysis.SdkMetadata{}, "id in ?", IDs).Error
	return err
}

// UpdateSdkMetadata 更新SdkMetadata表记录
// Author [yourname](https://github.com/yourname)
func (s *SdkMetadataService) UpdateSdkMetadata(SdkMetadata analysis.SdkMetadata) (err error) {
	err = aGloabl.AppRankDB.Model(&analysis.SdkMetadata{}).Where("id = ?", SdkMetadata.ID).Updates(&SdkMetadata).Error
	return err
}

// GetSdkMetadata 根据ID获取SdkMetadata表记录
// Author [yourname](https://github.com/yourname)
func (s *SdkMetadataService) GetSdkMetadata(ID string) (SdkMetadata analysis.SdkMetadata, err error) {
	err = aGloabl.AppRankDB.Where("id = ?", ID).First(&SdkMetadata).Error
	return
}
