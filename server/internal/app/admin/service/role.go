package service

import (
	"fmt"
	"kiwi/internal/app/admin/model"
	"kiwi/internal/app/admin/requests"
	commonModel "kiwi/internal/common/model"
	"kiwi/internal/common/service"

	"gorm.io/gorm"
)

type Role struct {
	service.BaseService[model.Role]
}

func NewRole(args ...any) *Role {
	srv := &Role{}
	srv.BaseService = service.NewSrv(srv.BaseService, args...)
	return srv
}

// ListRole
// 根据条件查询列表
func (srv *Role) ListRole(condition *commonModel.PageQuery[*requests.QueryRole]) (*commonModel.ResPage[model.Role], error) {
	return service.GetList[model.Role](condition, func(qu *requests.QueryRole, d *gorm.DB) *gorm.DB {
		if qu.KeyWords != "" {
			d = d.Where("name like ? or remark like ?", "%"+qu.KeyWords+"%", "%"+qu.KeyWords+"%")
		}

		return d
	})
}

func (srv *Role) GetRoleMenus(Role_code string, args ...any) ([]*model.Menu, error) {
	menu := make([]*model.Menu, 0)

	rm := srv.GetDB().Model(&model.RoleMenu{})
	if Role_code != "" {
		rm = rm.Where("role_code =?", Role_code)
	}

	sqlStr := rm.ToSQL(func(tx *gorm.DB) *gorm.DB { return tx.Find(nil) })
	qry := srv.GetDB().Model(&model.Menu{}).Joins(fmt.Sprintf(`left join (%s) rm on rm.menu_code = sys_menu.code`, sqlStr)).
		Select("sys_menu.*,case when rm.role_code is not null then 1 else 0 end as is_have")

	if len(args) > 0 {
		isHave := args[0].(int)
		qry = qry.Where("is_have = ?", isHave)
	}

	err := qry.Find(&menu).Error
	return menu, err
}

func (srv *Role) SaveRoleMenus(Role_code string, menu_codes []string) error {
	rm := srv.GetDB().Model(&model.RoleMenu{}).Where("role_code =?", Role_code)
	err := rm.Delete(&model.RoleMenu{}).Error
	if err != nil {
		return err
	}

	rm = srv.GetDB().Model(&model.RoleMenu{})
	for _, code := range menu_codes {
		rm = rm.Create(&model.RoleMenu{
			RoleCode: Role_code,
			MenuCode: code,
		})
	}

	return rm.Error
}
