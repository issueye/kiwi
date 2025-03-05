package service

import (
	"kiwi/internal/app/admin/model"
	"kiwi/internal/app/admin/requests"
	commonModel "kiwi/internal/common/model"
	"kiwi/internal/common/service"

	"gorm.io/gorm"
)

type Dicts struct {
	service.BaseService[model.DictsInfo]
}

func NewDicts(args ...any) *Dicts {
	srv := &Dicts{}
	srv.BaseService = service.NewSrv(srv.BaseService, args...)
	return srv
}

// ListDicts
// 根据条件查询列表
func (r *Dicts) ListDicts(condition *commonModel.PageQuery[*requests.QueryDicts]) (*commonModel.ResPage[model.DictsInfo], error) {
	return service.GetList[model.DictsInfo](condition, func(qu *requests.QueryDicts, d *gorm.DB) *gorm.DB {
		if qu.KeyWords != "" {
			d = d.Where("name like ? or remark like ?", "%"+qu.KeyWords+"%", "%"+qu.KeyWords+"%")
		}

		return d
	})
}

type DictDetail struct {
	service.BaseService[model.DictDetail]
}

func NewDictDetail(args ...any) *DictDetail {
	srv := &DictDetail{}
	srv.BaseService = service.NewSrv(srv.BaseService, args...)
	return srv
}

func (r *DictDetail) Save(data *model.DictDetail) error {
	return r.GetDB().Model(&model.DictDetail{}).Where("id = ?", data.ID).Save(data).Error
}

func (r *DictDetail) List(condition *commonModel.PageQuery[*requests.QueryDictsDetail]) (*commonModel.ResPage[model.DictDetail], error) {
	return service.GetList[model.DictDetail](condition, func(qu *requests.QueryDictsDetail, d *gorm.DB) *gorm.DB {
		if qu.KeyWords != "" {
			d = d.Where("name like ? or remark like ?", "%"+qu.KeyWords+"%", "%"+qu.KeyWords+"%")
		}

		if qu.DictCode != "" {
			d = d.Where("dict_code = ?", qu.DictCode)
		}

		return d
	})
}

func (r *Dicts) NotExistCreate(data *model.DictsInfo) error {
	// 检查字典是否存在
	var count int64
	err := r.GetDB().Model(&model.DictsInfo{}).Where("code = ?", data.Code).Count(&count).Error
	if err != nil {
		return err
	}

	if count == 0 {
		return r.GetDB().Create(data).Error
	}

	if count > 0 {
		for _, detail := range data.Details {
			// 检查字典明细是否存在
			err = r.GetDB().Model(&model.DictDetail{}).Where("dict_code =?", data.Code).Where("key = ?", detail.Key).Count(&count).Error
			if err != nil {
				return err
			}

			if count == 0 {
				err = r.GetDB().Create(&detail).Error
				if err != nil {
					return err
				}
			}
		}
	}

	return nil
}
