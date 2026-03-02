package openapi

// CreateAppRequest 创建应用请求
type CreateAppRequest struct {
	AppName string `json:"appName" validate:"required"`
	AppDesc string `json:"appDesc"`
	Remark  string `json:"remark"`
}

// CreateAppResponse 创建应用响应
type CreateAppResponse struct {
	AppId     string `json:"appId"`
	AppSecret string `json:"appSecret"`
}

// CreateAppKeyRequest 创建API密钥请求
type CreateAppKeyRequest struct {
	AppId     string `json:"appId" validate:"required"`
	ExpiredAt string `json:"expiredAt"` // 过期时间，格式: "2006-01-02 15:04:05"
	Remark    string `json:"remark"`
}

// CreateAppKeyResponse 创建API密钥响应
type CreateAppKeyResponse struct {
	AppKey    string `json:"appKey"`
	SecretKey string `json:"secretKey"`
}

// OpenAPIErrorResponse 错误响应
type OpenAPIErrorResponse struct {
	Code    int    `json:"code"`
	Message string `json:"message"`
}