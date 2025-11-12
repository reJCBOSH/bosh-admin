package menu

import (
	"bosh-admin/core/db"
	"bosh-admin/module/basic"
)

type GetMenuListReq struct {
	basic.Pagination
	Title string `json:"title" form:"title"`
}

type MenuBasicItem struct {
	Path          string `json:"path" form:"path" validate:"required_if=MenuType 0|required_if=MenuType 1|required_if=MenuType 2"` // 路由路径
	Name          string `json:"name" form:"name" validate:"required_if=MenuType 0|required_if=MenuType 1|required_if=MenuType 2"` // 路由名称(必须保持唯一)
	Redirect      string `json:"redirect" form:"redirect"`                                                                         // 路由重定向
	Component     string `json:"component" form:"component"`                                                                       // 按需加载需要展示的页面
	ParentId      uint   `json:"parentId" form:"parentId"`                                                                         // 父级菜单id
	MenuType      int    `json:"menuType" form:"menuType" validate:"oneof=0 1 2 3"`                                                // 菜单类型 0菜单 1iframe 2外链 3按钮
	Title         string `json:"title" form:"title" validate:"required"`                                                           // 菜单名称
	Icon          string `json:"icon" form:"icon"`                                                                                 // 菜单图标
	DisplayOrder  int    `json:"displayOrder" form:"displayOrder" validate:"required,gte=0,lte=9999"`                              // 菜单排序
	ShowBadge     bool   `json:"showBadge" form:"showBadge"`                                                                       // 是否显示徽章
	ShowTextBadge string `json:"showTextBadge" form:"showTextBadge"`                                                               // 文本徽章
	IsHide        bool   `json:"isHide" form:"isHide"`                                                                             // 是否在菜单中隐藏
	IsHideTab     bool   `json:"isHideTab" form:"isHideTab"`                                                                       // 是否在标签页中隐藏
	Link          string `json:"link" form:"link"`                                                                                 // 外部链接
	IsIframe      bool   `json:"isIframe" form:"isIframe"`                                                                         // 是否为iframe
	KeepAlive     bool   `json:"keepAlive" form:"keepAlive"`                                                                       // 是否缓存
	FixedTab      bool   `json:"fixedTab" form:"fixedTab"`                                                                         // 是否固定标签页
	AuthMark      string `json:"authMark" form:"authMark" validate:"required_if=MenuType 3"`                                       // 权限标识
}

type MenuItem struct {
	db.BasicModel
	MenuBasicItem
}

type GetMenuDetailResp struct {
	MenuItem
}

type AddMenuReq struct {
	db.AddBasicModel
	MenuBasicItem
}

type EditMenuReq struct {
	db.EditBasicModel
	MenuBasicItem
}

// ArtMenu 路由菜单
type ArtMenu struct {
	Id        uint        `json:"-"`                   // 菜单Id
	ParentId  uint        `json:"-"`                   // 父级菜单Id
	Path      string      `json:"path"`                // 路由地址
	Name      string      `json:"name"`                // 路由名称(必须保持唯一)
	Redirect  string      `json:"redirect,omitempty"`  // 路由重定向
	Component string      `json:"component,omitempty"` // 按需加载需要展示的页面
	Meta      ArtMenuMeta `json:"meta"`                // 路由元信息
	Children  []ArtMenu   `json:"children,omitempty"`  // 子路由配置项
}

// ArtMenuMeta 路由元信息
type ArtMenuMeta struct {
	Title         string        `json:"title"`                   // 菜单名称
	Icon          string        `json:"icon,omitempty"`          // 菜单图标
	ShowBadge     bool          `json:"showBadge"`               // 是否显示徽章
	ShowTextBadge string        `json:"showTextBadge,omitempty"` // 文本徽章
	IsHide        bool          `json:"isHide"`                  // 是否在菜单中隐藏
	IsHideTab     bool          `json:"isHideTab"`               // 是否在标签页中隐藏
	Link          string        `json:"link,omitempty"`          // 外部链接
	IsIframe      bool          `json:"isIframe"`                // 是否为iframe
	KeepAlive     bool          `json:"keepAlive"`               // 是否缓存
	FixedTab      bool          `json:"fixedTab"`                // 是否固定标签页
	AuthList      []ArtAuthItem `json:"authList"`                // 操作权限
}

type ArtAuthItem struct {
	Title    string `json:"title"`    // 权限名称
	AuthMark string `json:"authMark"` // 权限标识
}
