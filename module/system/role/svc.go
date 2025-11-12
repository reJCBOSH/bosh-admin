package role

import (
	"bosh-admin/core/db"
	"bosh-admin/core/exception"
	"bosh-admin/global"
	"bosh-admin/model"
)

type SysRoleSvc struct{}

func NewSysRoleSvc() *SysRoleSvc {
	return &SysRoleSvc{}
}

func (svc *SysRoleSvc) GetRoleList(roleName, roleCode string, status *int, pageNo, pageSize int) ([]model.SysRole, int64, error) {
	var list []model.SysRole
	var total int64
	var err error
	query := db.GormDB().Model(&model.SysRole{})
	if roleName != "" {
		query = query.Where("role_name LIKE ?", "%"+roleName+"%")
	}
	if roleCode != "" {
		query = query.Where("role_code LIKE ?", "%"+roleCode+"%")
	}
	if status != nil {
		query = query.Where("status = ?", *status)
	}
	if pageNo > 0 && pageSize > 0 {
		err = query.Count(&total).Error
		if err != nil {
			return nil, 0, exception.NewException("查询角色数量失败", err)
		}
		query = query.Scopes(db.PageScope(pageNo, pageSize))
	}
	err = query.Find(&list).Error
	if err != nil {
		return nil, 0, exception.NewException("查询角色列表失败", err)
	}
	return list, total, nil
}

func (svc *SysRoleSvc) GetRoleById(id any) (*model.SysRole, error) {
	role, err := db.QueryById[model.SysRole](id)
	if err != nil {
		return nil, exception.NewException("查询角色失败", err)
	}
	return role, nil
}

func (svc *SysRoleSvc) AddRole(role AddRoleReq) error {
	var count int64
	err := db.GormDB().Where("role_code = ?", role.RoleCode).Count(&count).Error
	if err != nil {
		return exception.NewException("查询角色失败", err)
	}
	if count > 0 {
		return exception.NewException("角色编码已存在")
	}
	err = db.Create(&role, new(model.SysRole).TableName())
	if err != nil {
		return exception.NewException("新增角色失败", err)
	}
	return nil
}

func (svc *SysRoleSvc) EditRole(role EditRoleReq) error {
	_, err := db.QueryById[model.SysRole](role.Id)
	if err != nil {
		return exception.NewException("查询角色失败", err)
	}
	err = db.Updates(&role, new(model.SysRole).TableName())
	if err != nil {
		return exception.NewException("修改角色失败", err)
	}
	return nil
}

func (svc *SysRoleSvc) DelRole(id any) error {
	originRole, err := db.QueryById[model.SysRole](id)
	if err != nil {
		return exception.NewException("查询角色失败", err)
	}
	if originRole.RoleCode == global.SuperAdmin {
		return exception.NewException("超级管理员角色不能删除")
	}
	var count int64
	err = db.GormDB().Model(&model.SysUser{}).Where("role_id = ?", id).Count(&count).Error
	if err != nil {
		return exception.NewException("查询角色用户失败", err)
	}
	if count > 0 {
		return exception.NewException("角色下存在用户，请先删除用户")
	}
	tx := db.Begin()
	// 删除角色
	err = tx.Delete(&model.SysRole{}, id).Error
	if err != nil {
		tx.Rollback()
		return exception.NewException("删除角色失败", err)
	}
	// 删除角色-菜单关联
	err = tx.Where("role_id = ?", id).Delete(&model.SysRoleMenu{}).Error
	if err != nil {
		tx.Rollback()
		return exception.NewException("删除角色菜单关联失败", err)
	}
	// 删除角色-部门关联
	err = tx.Where("role_id = ?", id).Delete(&model.SysRoleDept{}).Error
	if err != nil {
		tx.Rollback()
		return exception.NewException("删除角色部门关联失败", err)
	}
	tx.Commit()
	return nil
}

func (svc *SysRoleSvc) GetRoleMenu(roleId any) ([]model.SysMenu, error) {
	role, err := db.QueryById[model.SysRole](roleId)
	if err != nil {
		return nil, exception.NewException("查询角色失败", err)
	}
	var menuIds []uint
	err = db.GormDB().Model(&model.SysRoleMenu{}).Where("role_id = ?", role.Id).Pluck("menu_id", &menuIds).Error
	if err != nil {
		return nil, exception.NewException("查询角色菜单失败", err)
	}
	var menus []model.SysMenu
	err = db.GormDB().Model(&model.SysMenu{}).Where("id IN ?", menuIds).Order("display_order DESC,id ASC").Find(&menus).Error
	if err != nil {
		return nil, exception.NewException("查询角色菜单失败", err)
	}
	return menus, nil
}

func (svc *SysRoleSvc) GetRoleMenuIds(roleId any) ([]uint, error) {
	role, err := db.QueryById[model.SysRole](roleId)
	if err != nil {
		return nil, exception.NewException("查询角色失败", err)
	}
	var menuIds []uint
	err = db.GormDB().Model(&model.SysRoleMenu{}).Where("role_id = ?", role.Id).Pluck("menu_id", &menuIds).Error
	if err != nil {
		return nil, exception.NewException("查询角色菜单失败", err)
	}
	return menuIds, nil
}

func (svc *SysRoleSvc) SetRoleMenuAuth(roleId uint, menuIds []uint) error {
	role, err := db.QueryById[model.SysRole](roleId)
	if err != nil {
		return exception.NewException("查询角色失败", err)
	}
	if role.RoleCode == global.SuperAdmin {
		return exception.NewException("超级管理员角色不能修改权限")
	}
	var menuNum int64
	err = db.GormDB().Model(&model.SysMenu{}).Where("id IN ?", menuIds).Count(&menuNum).Error
	if err != nil {
		return exception.NewException("查询菜单失败", err)
	}
	if menuNum != int64(len(menuIds)) {
		return exception.NewException("菜单权限数据错误")
	}
	var roleMenus []model.SysRoleMenu
	for _, menuId := range menuIds {
		roleMenus = append(roleMenus, model.SysRoleMenu{
			RoleId: role.Id,
			MenuId: menuId,
		})
	}
	tx := db.Begin()
	err = tx.Where("role_id = ?", role.Id).Delete(&model.SysRoleMenu{}).Error
	if err != nil {
		tx.Rollback()
		return exception.NewException("删除角色菜单失败", err)
	}
	err = tx.Create(&roleMenus).Error
	if err != nil {
		tx.Rollback()
		return exception.NewException("设置角色菜单失败", err)
	}
	tx.Commit()
	return nil
}

func (svc *SysRoleSvc) GetRoleDeptIds(roleId any) ([]uint, error) {
	role, err := db.QueryById[model.SysRole](roleId)
	if err != nil {
		return nil, exception.NewException("查询角色失败", err)
	}
	var deptIds []uint
	err = db.GormDB().Model(&model.SysRoleDept{}).Where("role_id = ?", role.Id).Pluck("dept_id", &deptIds).Error
	if err != nil {
		return nil, exception.NewException("查询角色部门失败", err)
	}
	return deptIds, nil
}

func (svc *SysRoleSvc) SetRoleDataAuth(roleId uint, dataAuth int, deptIds []uint) error {
	if dataAuth == 5 && len(deptIds) == 0 {
		return exception.NewException("数据权限为自定义时，部门不能为空")
	}
	role, err := db.QueryById[model.SysRole](roleId)
	if err != nil {
		return exception.NewException("查询角色失败", err)
	}
	var roleDepts []model.SysRoleDept
	if dataAuth == 5 {
		var deptNum int64
		err = db.GormDB().Model(&model.SysDept{}).Where("id IN ?", deptIds).Count(&deptNum).Error
		if err != nil {
			return exception.NewException("查询部门失败", err)
		}
		if deptNum != int64(len(deptIds)) {
			return exception.NewException("部门权限数据错误")
		}
		for _, deptId := range deptIds {
			roleDepts = append(roleDepts, model.SysRoleDept{
				RoleId: role.Id,
				DeptId: deptId,
			})
		}
	}
	tx := db.Begin()
	if role.DataAuth != dataAuth {
		err = tx.Model(&model.SysRole{}).Where("id = ?", role.Id).Update("data_auth", dataAuth).Error
		if err != nil {
			tx.Rollback()
			return exception.NewException("设置角色数据权限失败", err)
		}
	}
	if dataAuth == 5 {
		err = tx.Where("role_id = ?", role.Id).Delete(&model.SysRoleDept{}).Error
		if err != nil {
			tx.Rollback()
			return exception.NewException("删除角色部门失败", err)
		}
		err = tx.Create(&roleDepts).Error
		if err != nil {
			tx.Rollback()
			return exception.NewException("设置角色部门失败", err)
		}
	}
	tx.Commit()
	return nil
}

func (svc *SysRoleSvc) SetRoleStatus(currentRoleId, roleId uint, status int) error {
	role, err := db.QueryById[model.SysRole](roleId)
	if err != nil {
		return exception.NewException("查询角色失败", err)
	}
	if role.RoleCode == global.SuperAdmin {
		return exception.NewException("超级管理员角色不能修改状态")
	}
	if role.Status == status {
		return nil
	}
	if status == 1 {
		var menuNum int64
		err = db.GormDB().Model(&model.SysRoleMenu{}).Where("role_id = ?", role.Id).Count(&menuNum).Error
		if err != nil {
			return exception.NewException("查询角色菜单失败", err)
		}
		if menuNum == 0 {
			return exception.NewException("请先分配菜单权限")
		}
		if role.DataAuth == 0 {
			return exception.NewException("请先分配数据权限")
		}
	} else {
		if currentRoleId == roleId {
			return exception.NewException("无法禁用当前操作员角色")
		}
	}
	err = db.GormDB().Model(&model.SysRole{}).Where("id = ?", role.Id).Update("status", status).Error
	if err != nil {
		return exception.NewException("设置角色状态失败", err)
	}
	return nil
}
