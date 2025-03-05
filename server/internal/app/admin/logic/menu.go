package logic

import (
	"errors"
	"kiwi/internal/app/admin/model"
	"kiwi/internal/app/admin/requests"
	"kiwi/internal/app/admin/service"
	commonModel "kiwi/internal/common/model"
)

func SaveRoleMenus(code string, menu_codes []string) error {
	return service.NewRole().SaveRoleMenus(code, menu_codes)
}

func GetRoleMenus(Code string) ([]*model.Menu, error) {
	data, err := service.NewRole().GetRoleMenus(Code)
	if err != nil {
		return nil, err
	}

	list := MakeTree(data)
	return list, nil
}

// 创建数据
func CreateMenu(r *requests.CreateMenu) error {
	srv := service.NewMenu()

	data, err := srv.GetByField("code", r.Code)
	if err != nil {
		return err
	}

	if data.ID != 0 {
		return errors.New("角色编码已存在")
	}

	info := &model.Menu{
		MenuBase: model.MenuBase{
			Code:        r.Code,
			Name:        r.Name,
			Description: r.Description,
			Frontpath:   r.Frontpath,
			Order:       r.Order,
			Icon:        r.Icon,
			ParentCode:  r.ParentCode,
			Visible:     true,
		},
	}

	return service.NewMenu().Create(info)
}

// 更新数据
func UpdateMenu(r *requests.UpdateMenu) error {
	data := make(map[string]any)
	data["code"] = r.Code
	data["name"] = r.Name
	data["description"] = r.Description
	data["frontpath"] = r.Frontpath
	// data["condition"] = r.Condition
	data["order"] = r.Order
	data["icon"] = r.Icon
	// data["method"] = r.Method
	data["parent_code"] = r.ParentCode
	data["visible"] = true

	return service.NewMenu().UpdateByMap(uint(r.Id), data)
}

// 根据ID查询数据
func GetMenuById(id uint) (*model.Menu, error) {
	return service.NewMenu().GetById(id)
}

func GetCatalog() ([]*model.Menu, error) {
	return service.NewMenu().GetCatalog()
}

// 根据条件查询数据
func ListMenu(condition *commonModel.PageQuery[*requests.QueryMenu]) (*commonModel.ResPage[model.Menu], error) {
	res, err := service.NewMenu().ListMenu(condition)
	if err != nil {
		return nil, err
	}

	res.List = MakeTree(res.List)
	return res, nil
}

// 删除数据
func DeleteMenu(id uint) error {
	return service.NewMenu().Delete(id)
}

func GetMenuTree(Role_code string) ([]*model.Menu, error) {
	list, err := service.NewRole().GetRoleMenus(Role_code)
	if err != nil {
		return nil, err
	}

	return MakeTree(list), nil
}

func MakeTree(list []*model.Menu) []*model.Menu {
	findFirst := func(list []*model.Menu) []*model.Menu {
		// 如果 parentCode 为空，则返回第一个元素
		if len(list) == 0 {
			return nil
		}

		rtnList := make([]*model.Menu, 0)

		for _, menu := range list {
			if menu.ParentCode == "" {
				rtnList = append(rtnList, menu)
			}
		}
		return rtnList
	}

	findChild := func(list []*model.Menu, parentCode string) []*model.Menu {
		// 查找所有子菜单
		rtnList := make([]*model.Menu, 0)

		for _, menu := range list {
			if menu.ParentCode == parentCode {
				rtnList = append(rtnList, menu)
			}
		}
		return rtnList
	}

	fList := findFirst(list)

	for _, menu := range fList {
		data := findChild(list, menu.Code)
		menu.Children = data
	}

	return fList
}
