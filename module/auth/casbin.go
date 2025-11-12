package auth

import (
	"strconv"
	"sync"

	"bosh-admin/core/db"
	"bosh-admin/core/exception"
	"bosh-admin/core/log"

	"github.com/casbin/casbin/v2"
	"github.com/casbin/casbin/v2/model"
	gormAdapter "github.com/casbin/gorm-adapter/v3"
)

type CasbinSvc struct{}

func NewCasbinSvc() *CasbinSvc {
	return &CasbinSvc{}
}

var (
	syncedEnforcer *casbin.SyncedEnforcer
	once           sync.Once
)

// CasbinInfo 访问权限
type CasbinInfo struct {
	Path   string `json:"path"`   // 路径
	Method string `json:"method"` // 方法
}

func (svc *CasbinSvc) Casbin() *casbin.SyncedEnforcer {
	once.Do(func() {
		a, _ := gormAdapter.NewAdapterByDB(db.GormDB())
		text := `
		[request_definition]
		r = sub, obj, act
		
		[policy_definition]
		p = sub, obj, act
		
		[role_definition]
		g = _, _
		
		[policy_effect]
		e = some(where (p.eft == allow))
		
		[matchers]
		m = r.sub == p.sub && keyMatch2(r.obj,p.obj) && r.act == p.act
		`
		m, err := model.NewModelFromString(text)
		if err != nil {
			log.Error("字符串加载模型失败!", err.Error())
			return
		}
		syncedEnforcer, _ = casbin.NewSyncedEnforcer(m, a)
	})
	_ = syncedEnforcer.LoadPolicy()
	return syncedEnforcer
}

// ClearCasbin 清除匹配的访问权限
func (svc *CasbinSvc) ClearCasbin(v int, p ...string) bool {
	e := svc.Casbin()
	success, _ := e.RemoveFilteredPolicy(v, p...)
	return success
}

// UpdateCasbin 更新访问权限
func (svc *CasbinSvc) UpdateCasbin(roleId uint, casbinInfos []CasbinInfo) error {
	roleIdStr := strconv.Itoa(int(roleId))
	svc.ClearCasbin(0, roleIdStr)
	var rules [][]string
	for _, v := range casbinInfos {
		rules = append(rules, []string{roleIdStr, v.Path, v.Method})
	}
	e := svc.Casbin()
	success, err := e.AddPolicies(rules)
	if !success {
		return exception.NewException("存在相同访问权限，添加失败，请联系超级管理员", err)
	}
	return nil
}

// UpdateCasbinApi 更新访问api
func (svc *CasbinSvc) UpdateCasbinApi(oldPath, newPath, oldMethod, newMethod string) error {
	e := svc.Casbin()
	success, err := e.UpdatePolicy([]string{"", oldPath, oldMethod}, []string{"", newPath, newMethod})
	if !success {
		return exception.NewException("更新访问权限失败", err)
	}
	return db.GormDB().Model(&gormAdapter.CasbinRule{}).Where("v1 = ? AND v2 = ?", oldPath, oldMethod).
		Updates(map[string]interface{}{
			"v1": newPath,
			"v2": newMethod,
		}).Error
}

// GetCasbinByRoleId 通过roleId获取访问权限
func (svc *CasbinSvc) GetCasbinByRoleId(roleId uint) ([]CasbinInfo, error) {
	e := svc.Casbin()
	roleIdStr := strconv.Itoa(int(roleId))
	list, err := e.GetFilteredPolicy(0, roleIdStr)
	if err != nil {
		return nil, exception.NewException("获取访问权限失败", err)
	}
	var rules []CasbinInfo
	for _, v := range list {
		rules = append(rules, CasbinInfo{
			Path:   v[1],
			Method: v[2],
		})
	}
	return rules, nil
}

func (svc *CasbinSvc) RemoveCasbin(path, method string) error {
	e := svc.Casbin()
	success, err := e.RemovePolicy("", path, method)
	if !success {
		return exception.NewException("移除访问权限失败", err)
	}
	return db.GormDB().Model(&gormAdapter.CasbinRule{}).Where("api_path = ?", path).Where("api_method = ?", method).
		Delete(&gormAdapter.CasbinRule{}).Error
}
