package model

type SysRoleMenu struct {
	RoleId uint `gorm:"comment:角色id"`
	MenuId uint `gorm:"comment:菜单id"`
}

func (SysRoleMenu) TableName() string {
	return "sys_role_menu"
}

func (SysRoleMenu) TableComment() string {
	return "系统角色菜单关系表"
}
