package migrations

import (
	"time"

	"bosh-admin/core/db"
	"bosh-admin/core/log"
	"bosh-admin/global"
	"bosh-admin/model"
	"bosh-admin/util"

	"github.com/go-gormigrate/gormigrate/v2"
	"gorm.io/gorm"
)

var InitSchema = &gormigrate.Migration{
	ID: "202507090943_init_schema",
	Migrate: func(tx *gorm.DB) error {
		err := tx.Migrator().AutoMigrate(
			&model.Resource{},
			&model.SysApi{},
			&model.SysBlackJwt{},
			&model.SysDept{},
			&model.SysLoginRecord{},
			&model.SysMenu{},
			&model.SysOperationRecord{},
			&model.SysRole{},
			&model.SysRoleDept{},
			&model.SysRoleMenu{},
			&model.SysUser{},
		)
		if err != nil {
			log.Error("初始化数据表失败", err)
			return err
		}
		dept := model.SysDept{
			DeptName:     "系统管理",
			DeptCode:     global.SystemAdmin,
			Remark:       "系统管理",
			Status:       1,
			DisplayOrder: 9999,
		}
		err = tx.Create(&dept).Error
		if err != nil {
			log.Error("初始化部门数据失败", err)
			return err
		}
		role := model.SysRole{
			RoleName: "超级管理员",
			RoleCode: global.SuperAdmin,
			Status:   1,
			Remark:   "超级管理员",
			DataAuth: 1,
		}
		err = tx.Create(&role).Error
		if err != nil {
			log.Error("初始化角色数据表失败", err)
			return err
		}
		defaultPwd, _ := util.BcryptHash(global.DefaultPassword)
		user := model.SysUser{
			Username:     global.SuperAdmin,
			Password:     defaultPwd,
			PwdUpdatedAt: db.CustomTime(time.Now()),
			Nickname:     "超级管理员",
			Gender:       0,
			Introduce:    "行天之道，总司一切",
			Status:       1,
			RoleId:       role.Id,
			DeptId:       dept.Id,
		}
		err = tx.Create(&user).Error
		if err != nil {
			log.Error("初始化用户数据表失败", err)
			return err
		}
		menus := []model.SysMenu{
			{Path: "/system", Name: "System", ParentId: 0, Title: "系统管理", Icon: "ri:settings-3-line"},
			{Path: "/system/user", Name: "SystemUser", Component: "system/user/index", ParentId: 1, Title: "用户管理", Icon: "ri:admin-line", KeepAlive: true},
			{Path: "/system/role", Name: "SystemRole", Component: "system/role/index", ParentId: 1, Title: "角色管理", Icon: "ri:admin-fill", KeepAlive: true},
			{Path: "/system/menu", Name: "SystemMenu", Component: "system/menu/index", ParentId: 1, Title: "菜单管理", Icon: "ep:menu", KeepAlive: true},
			{Path: "/system/dept", Name: "SystemDept", Component: "system/dept/index", ParentId: 1, Title: "部门管理", Icon: "ri:git-branch-line", KeepAlive: true},
			{Path: "/monitor", Name: "Monitor", ParentId: 0, Title: "系统监控", Icon: "ep:monitor"},
			{Path: "/monitor/loginRecord", Name: "LoginRecord", Component: "monitor/loginRecord/index", ParentId: 6, Title: "登录日志", Icon: "ri:window-line", KeepAlive: true},
			{Path: "/monitor/operationRecord", Name: "OperationRecord", Component: "monitor/operationRecord/index", ParentId: 6, Title: "操作日志", Icon: "ri:history-fill", KeepAlive: true},
			// 用户按钮
			{ParentId: 2, MenuType: 3, Title: "新增用户", AuthMark: "sysUser:add"},
			{ParentId: 2, MenuType: 3, Title: "修改用户", AuthMark: "sysUser:edit"},
			{ParentId: 2, MenuType: 3, Title: "删除用户", AuthMark: "sysUser:del"},
			{ParentId: 2, MenuType: 3, Title: "设置用户状态", AuthMark: "sysUser:status"},
			{ParentId: 2, MenuType: 3, Title: "重置密码", AuthMark: "sysUser:resetPassword"},
			// 角色按钮
			{ParentId: 3, MenuType: 3, Title: "新增角色", AuthMark: "sysRole:add"},
			{ParentId: 3, MenuType: 3, Title: "修改角色", AuthMark: "sysRole:edit"},
			{ParentId: 3, MenuType: 3, Title: "删除角色", AuthMark: "sysRole:del"},
			{ParentId: 3, MenuType: 3, Title: "设置角色状态", AuthMark: "sysRole:status"},
			{ParentId: 3, MenuType: 3, Title: "设置菜单权限", AuthMark: "sysRole:menuAuth"},
			{ParentId: 3, MenuType: 3, Title: "设置数据权限", AuthMark: "sysRole:dataAuth"},
			// 菜单按钮
			{ParentId: 4, MenuType: 3, Title: "新增菜单", AuthMark: "sysMenu:add"},
			{ParentId: 4, MenuType: 3, Title: "修改菜单", AuthMark: "sysMenu:edit"},
			{ParentId: 4, MenuType: 3, Title: "删除菜单", AuthMark: "sysMenu:del"},
			// 部门按钮
			{ParentId: 5, MenuType: 3, Title: "新增部门", AuthMark: "sysDept:add"},
			{ParentId: 5, MenuType: 3, Title: "修改部门", AuthMark: "sysDept:edit"},
			{ParentId: 5, MenuType: 3, Title: "删除部门", AuthMark: "sysDept:del"},
			// 登录日志按钮
			{ParentId: 7, MenuType: 3, Title: "删除日志", AuthMark: "sysLoginRecord:del"},
			{ParentId: 7, MenuType: 3, Title: "批量删除日志", AuthMark: "sysLoginRecord:batchDel"},
			// 操作日志按钮
			{ParentId: 8, MenuType: 3, Title: "查看详情", AuthMark: "sysOperationRecord:view"},
			{ParentId: 8, MenuType: 3, Title: "删除日志", AuthMark: "sysOperationRecord:del"},
			{ParentId: 8, MenuType: 3, Title: "批量删除日志", AuthMark: "sysOperationRecord:batchDel"},
		}
		err = tx.Create(&menus).Error
		if err != nil {
			log.Error("初始化菜单数据表失败", err)
			return err
		}
		return nil
	},
	Rollback: func(tx *gorm.DB) error {
		return tx.Migrator().DropTable(
			&model.Resource{},
			&model.SysApi{},
			&model.SysBlackJwt{},
			&model.SysDept{},
			&model.SysLoginRecord{},
			&model.SysMenu{},
			&model.SysOperationRecord{},
			&model.SysRole{},
			&model.SysRoleDept{},
			&model.SysRoleMenu{},
			&model.SysUser{},
		)
	},
}
