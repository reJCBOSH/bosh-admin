package openapi

import (
	"bosh-admin/global"
	"bosh-admin/model"
	"crypto/hmac"
	"crypto/sha256"
	"encoding/hex"
	"errors"

	"gorm.io/gorm"
)

type OpenapiSvc struct {
	db *gorm.DB
}

func NewOpenAPISvc() *OpenapiSvc {
	return &OpenapiSvc{
		db: global.GormDB,
	}
}

// AppKeyInfo 应用密钥信息
type AppKeyInfo struct {
	AppId     string
	AppKey    string
	SecretKey string
}

// GetAppKeyInfo 获取应用密钥信息
func (svc *OpenapiSvc) GetAppKeyInfo(appKey string) (*AppKeyInfo, error) {
	var app model.SysApp
	err := svc.db.Where("app_key = ? AND status = 1", appKey).First(&app).Error
	if err != nil {
		if errors.Is(err, gorm.ErrRecordNotFound) {
			return nil, errors.New("无效的appKey")
		}
		return nil, err
	}
	return &AppKeyInfo{
		AppId:     app.AppId,
		AppKey:    app.AppKey,
		SecretKey: app.SecretKey,
	}, nil
}

// GenerateSignature 生成签名
func (svc *OpenapiSvc) GenerateSignature(secretKey string, timestamp, nonce string, body string) string {
	// 按照规则拼接签名字符串
	signStr := timestamp + "\n" + nonce + "\n" + body
	// 使用HMAC-SHA256算法生成签名
	h := hmac.New(sha256.New, []byte(secretKey))
	h.Write([]byte(signStr))
	return hex.EncodeToString(h.Sum(nil))
}

// VerifySignature 验证签名
func (svc *OpenapiSvc) VerifySignature(appKey, timestamp, nonce, body, signature string) error {
	// 获取应用密钥信息
	appKeyInfo, err := svc.GetAppKeyInfo(appKey)
	if err != nil {
		return err
	}

	// 生成签名
	expectedSignature := svc.GenerateSignature(appKeyInfo.SecretKey, timestamp, nonce, body)

	// 比较签名
	if signature != expectedSignature {
		return errors.New("签名验证失败")
	}

	return nil
}

// CheckAppPermission 检查应用权限
func (svc *OpenapiSvc) CheckAppPermission(appId, module, apiPath, method string) error {
	// 检查是否存在对应的权限记录
	var count int64
	err := svc.db.Model(&model.SysAppPerm{}).
		Where("app_id = ? AND module = ? AND api_path = ? AND method = ? AND status = 1",
			appId, module, apiPath, method).
		Count(&count).Error

	if err != nil {
		return err
	}

	if count == 0 {
		return errors.New("应用无权限访问该接口")
	}

	return nil
}

// GetAppPermissions 获取应用所有权限
func (svc *OpenapiSvc) GetAppPermissions(appId string) ([]model.SysAppPerm, error) {
	var permissions []model.SysAppPerm
	err := svc.db.Where("app_id = ? AND status = 1", appId).Find(&permissions).Error
	return permissions, err
}
