package service

import (
	"kiwi/internal/app/admin/model"
	"kiwi/internal/app/admin/requests"
	commonModel "kiwi/internal/common/model"
	"kiwi/internal/common/service"

	"gorm.io/gorm"
)

type Menu struct {
	service.BaseService[model.Menu]
}

func NewMenu(args ...any) *Menu {
	srv := &Menu{}
	srv.BaseService = service.NewSrv(srv.BaseService, args...)
	return srv
}

// GetCatalog 获取目录
func (srv *Menu) GetCatalog() ([]*model.Menu, error) {
	var menus []*model.Menu
	err := srv.GetDB().Model(&model.Menu{}).Where("parent_code = ?", "").Find(&menus).Error
	return menus, err
}

// AddMenus 添加菜单
func (srv *Menu) AddMenu(menu *model.Menu) error {
	return srv.GetDB().Create(menu).Error
}

// 检查菜单是否存在
func (srv *Menu) CheckMenuExist(menu *model.Menu) (bool, error) {
	var count int64
	err := srv.GetDB().Model(&model.Menu{}).Where("code = ?", menu.Code).Count(&count).Error
	if err != nil {
		return false, err
	}
	return count > 0, nil
}

// ListMenu
// 根据条件查询列表
func (srv *Menu) ListMenu(condition *commonModel.PageQuery[*requests.QueryMenu]) (*commonModel.ResPage[model.Menu], error) {
	return service.GetList[model.Menu](condition, func(qu *requests.QueryMenu, d *gorm.DB) *gorm.DB {
		if qu.KeyWords != "" {
			d = d.Where("name like ? or remark like ?", "%"+qu.KeyWords+"%", "%"+qu.KeyWords+"%")
		}

		return d
	})
}
