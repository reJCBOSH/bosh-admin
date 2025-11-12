package menu

import (
	"bosh-admin/core/db"
	"bosh-admin/core/exception"
	"bosh-admin/global"
	"bosh-admin/model"
)

type SysMenuSvc struct{}

func NewSysMenuSvc() *SysMenuSvc {
	return &SysMenuSvc{}
}

func (svc *SysMenuSvc) GetMenuTree() ([]model.SysMenu, error) {
	treeMap, err := getMenuTreeMap()
	menuTree := treeMap[0]
	for i := 0; i < len(menuTree); i++ {
		err = getMenuChildrenList(&menuTree[i], treeMap)
	}
	return menuTree, err
}

// getMenuTreeMap 获取菜单Map
func getMenuTreeMap() (map[uint][]model.SysMenu, error) {
	var allMenus []model.SysMenu
	treeMap := make(map[uint][]model.SysMenu)
	err := db.GormDB().Order("display_order DESC").Find(&allMenus).Error
	for _, v := range allMenus {
		treeMap[v.ParentId] = append(treeMap[v.ParentId], v)
	}
	return treeMap, err
}

// getMenuChildrenList 获取子菜单列表
func getMenuChildrenList(menu *model.SysMenu, treeMap map[uint][]model.SysMenu) (err error) {
	menu.Children = treeMap[menu.Id]
	for i := 0; i < len(menu.Children); i++ {
		err = getMenuChildrenList(&menu.Children[i], treeMap)
	}
	return err
}

func (svc *SysMenuSvc) GetMenuList(title string, pageNo, pageSize int) ([]model.SysMenu, int64, error) {
	var list []model.SysMenu
	var total int64
	var err error
	query := db.GormDB().Model(&model.SysMenu{})
	if title != "" {
		query = query.Where("title LIKE ?", "%"+title+"%")
	}
	if pageNo > 0 && pageSize > 0 {
		if err = query.Count(&total).Error; err != nil {
			return nil, 0, exception.NewException("查询菜单数量失败", err)
		}
		query = query.Scopes(db.PageScope(pageSize, pageNo))
	}
	err = query.Find(&list).Error
	if err != nil {
		return nil, 0, exception.NewException("查询菜单列表失败", err)
	}
	return list, total, nil
}

func (svc *SysMenuSvc) GetMenuById(id uint) (*model.SysMenu, error) {
	menu, err := db.QueryById[model.SysMenu](id)
	if err != nil {
		return nil, exception.NewException("查询菜单失败", err)
	}
	return menu, nil
}

func (svc *SysMenuSvc) AddMenu(menu AddMenuReq) error {
	var err error
	if menu.MenuType < 3 {
		var count int64
		err = db.GormDB().Where("menu_type < ?", 3).Where("name = ?", menu.Name).Count(&count).Error
		if err != nil {
			return exception.NewException("查询重名菜单失败", err)
		}
		if count > 0 {
			return exception.NewException("路由名称已存在，必须保持唯一")
		}
	} else {
		var count int64
		err = db.GormDB().Where("parent_id = ?", menu.ParentId).Where("auth_mark = ?", menu.AuthMark).Count(&count).Error
		if err != nil {
			return exception.NewException("查询重名按钮失败", err)
		}
		if count > 0 {
			return exception.NewException("权限标识已存在，必须保持唯一")
		}
	}
	err = db.Create(&menu, new(model.SysMenu).TableName())
	if err != nil {
		return exception.NewException("新增菜单失败", err)
	}
	return nil
}

func (svc *SysMenuSvc) EditMenu(menu EditMenuReq) error {
	originMenu, err := db.QueryById[model.SysMenu](menu.Id)
	if err != nil {
		return exception.NewException("查询菜单失败", err)
	}
	if menu.MenuType < 3 && menu.Name != originMenu.Name {
		var count int64
		err = db.GormDB().Where("menu_type < ?", 3).Where("name = ?", menu.Name).Where("id != ?", menu.Id).Count(&count).Error
		if err != nil {
			return exception.NewException("查询重名菜单失败", err)
		}
		if count > 0 {
			return exception.NewException("路由名称已存在，必须保持唯一")
		}
	}
	err = db.Updates(&menu, new(model.SysMenu).TableName())
	if err != nil {
		return exception.NewException("修改菜单失败", err)
	}
	return nil
}

func (svc *SysMenuSvc) DelMenu(id any) error {
	originMenu, err := db.QueryById[model.SysMenu](id)
	if err != nil {
		return exception.NewException("查询菜单失败", err)
	}
	if originMenu.MenuType < 3 {
		var count int64
		err = db.GormDB().Where("parent_id = ?", originMenu.Id).Where("menu_type < ?", 3).Count(&count).Error
		if err != nil {
			return exception.NewException("查询子菜单失败", err)
		}
		if count > 0 {
			return exception.NewException("存在子菜单，无法删除")
		}
	}
	tx := db.Begin()
	// 删除菜单
	err = tx.Delete(&model.SysMenu{}, id).Error
	if err != nil {
		tx.Rollback()
		return exception.NewException("删除菜单失败", err)
	}
	// 删除角色-菜单关联
	err = tx.Where("menu_id = ?", id).Delete(&model.SysRoleMenu{}).Error
	if err != nil {
		tx.Rollback()
		return exception.NewException("删除角色-菜单关联失败", err)
	}
	// 删除按钮子菜单及角色-按钮关联
	if originMenu.MenuType < 3 {
		var btnIds []uint
		err = tx.Model(&model.SysMenu{}).Where("parent_id = ?", originMenu.Id).Where("menu_type = ?", 3).Pluck("id", &btnIds).Error
		if err != nil {
			tx.Rollback()
			return exception.NewException("查询按钮子菜单失败", err)
		}
		if len(btnIds) > 0 {
			// 删除按钮子菜单
			err = tx.Delete(&model.SysMenu{}, btnIds).Error
			if err != nil {
				tx.Rollback()
				return exception.NewException("删除按钮子菜单失败", err)
			}
			// 删除角色-按钮关联
			err = tx.Where("menu_id IN ?", btnIds).Delete(&model.SysRoleMenu{}).Error
			if err != nil {
				tx.Rollback()
				return exception.NewException("删除角色-按钮关联失败", err)
			}
		}
	}
	tx.Commit()
	return nil
}

// getAsyncRoutesChildrenList 获取art admin子菜单列表
func getAsyncRoutesChildrenList(menu *ArtMenu, treeMap map[uint][]ArtMenu) error {
	var err error
	menu.Children = treeMap[menu.Id]
	for i := 0; i < len(menu.Children); i++ {
		err = getAsyncRoutesChildrenList(&menu.Children[i], treeMap)
	}
	return err
}

// GetAsyncRoutes 获取art admin菜单
func (svc *SysMenuSvc) GetAsyncRoutes(roleId uint, roleCode string) ([]ArtMenu, error) {
	var roleMenuIds []uint
	if roleCode != global.SuperAdmin {
		err := db.GormDB().Model(&model.SysRoleMenu{}).Where("role_id = ?", roleId).Pluck("menu_id", &roleMenuIds).Error
		if err != nil {
			return nil, exception.NewException("查询角色菜单失败", err)
		}
	}
	var buttons []model.SysMenu
	buttonQuery := db.GormDB()
	if roleCode != global.SuperAdmin {
		buttonQuery = buttonQuery.Where("id IN ?", roleMenuIds)
	}
	err := buttonQuery.Where("menu_type = ?", 3).Order("display_order DESC").Find(&buttons).Error
	if err != nil {
		return nil, exception.NewException("查询按钮失败", err)
	}
	btnMap := make(map[uint][]model.SysMenu)
	for _, button := range buttons {
		btnMap[button.ParentId] = append(btnMap[button.ParentId], button)
	}
	var menus []model.SysMenu
	menuQuery := db.GormDB()
	if roleCode != global.SuperAdmin {
		menuQuery = menuQuery.Where("id IN ?", roleMenuIds)
	}
	err = menuQuery.Where("menu_type != ?", 3).Order("display_order DESC").Find(&menus).Error
	if err != nil {
		return nil, exception.NewException("查询菜单失败", err)
	}
	menuMap := make(map[uint][]ArtMenu)
	for _, menu := range menus {
		artMenu := ArtMenu{
			Id:        menu.Id,
			ParentId:  menu.ParentId,
			Path:      menu.Path,
			Name:      menu.Name,
			Redirect:  menu.Redirect,
			Component: menu.Component,
			Meta: ArtMenuMeta{
				Title:         menu.Title,
				Icon:          menu.Icon,
				ShowBadge:     menu.ShowBadge,
				ShowTextBadge: menu.ShowTextBadge,
				IsHide:        menu.IsHide,
				IsHideTab:     menu.IsHideTab,
				Link:          menu.Link,
				IsIframe:      menu.IsIframe,
				KeepAlive:     menu.KeepAlive,
				FixedTab:      menu.FixedTab,
			},
		}
		if btnArr, ok := btnMap[menu.Id]; ok {
			var auths []ArtAuthItem
			for _, btn := range btnArr {
				auths = append(auths, ArtAuthItem{
					Title:    btn.Title,
					AuthMark: btn.AuthMark,
				})
			}
			artMenu.Meta.AuthList = auths
		}
		menuMap[menu.ParentId] = append(menuMap[menu.ParentId], artMenu)
	}
	routers := menuMap[0]
	for i := 0; i < len(routers); i++ {
		err = getAsyncRoutesChildrenList(&routers[i], menuMap)
		if err != nil {
			return nil, exception.NewException("获取子菜单列表失败", err)
		}
	}
	return routers, nil
}
