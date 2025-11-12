package model

type SysRoleDept struct {
	RoleId uint `gorm:"comment:角色id"`
	DeptId uint `gorm:"comment:部门id"`
}

func (SysRoleDept) TableName() string {
	return "sys_role_dept"
}

func (SysRoleDept) TableComment() string {
	return "系统角色部门关系表"
}
