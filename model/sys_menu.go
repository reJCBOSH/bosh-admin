package model

import (
	"bosh-admin/core/db"
)

type SysMenu struct {
	db.BasicModel
	Path          string    `gorm:"not null;comment:路由路径" json:"path"`
	Name          string    `gorm:"type:varchar(30);not null;comment:路由名称" json:"name"`
	Redirect      string    `gorm:"comment:路由重定向" json:"redirect"`
	Component     string    `gorm:"comment:按需加载需要展示的页面" json:"component"`
	ParentId      uint      `gorm:"default:0;comment:父级菜单id" json:"parentId"`
	MenuType      int       `gorm:"default:0;comment:菜单类型 0菜单 1iframe 2外链 3按钮" json:"menuType"`
	Title         string    `gorm:"type:varchar(30);not null;comment:菜单名称" json:"title"`
	Icon          string    `gorm:"type:varchar(30);not null;comment:菜单图标" json:"icon"`
	DisplayOrder  int       `gorm:"type:int;default:0;comment:菜单排序" json:"displayOrder"`
	ShowBadge     bool      `gorm:"default:0;comment:是否显示徽章" json:"showBadge"`
	ShowTextBadge string    `gorm:"default:0;comment:文本徽章" json:"showTextBadge"`
	IsHide        bool      `gorm:"default:0;comment:是否在菜单中隐藏" json:"isHide"`
	IsHideTab     bool      `gorm:"default:0;comment:是否在标签页中隐藏" json:"isHideTab"`
	Link          string    `gorm:"not null;comment:外链" json:"link"`
	IsIframe      bool      `gorm:"default:0;comment:是否内嵌iframe" json:"isIframe"`
	KeepAlive     bool      `gorm:"default:0;comment:是否缓存改路由页面" json:"keepAlive"`
	FixedTab      bool      `gorm:"default:0;comment:固定标签页" json:"fixedTab"`
	AuthMark      string    `gorm:"not null;comment:权限标识" json:"authMark"`
	Children      []SysMenu `gorm:"-" json:"children"`
}

func (SysMenu) TableName() string {
	return "sys_menu"
}

func (SysMenu) TableComment() string {
	return "系统菜单表"
}
