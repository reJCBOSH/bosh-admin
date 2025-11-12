package auth

import (
	"slices"
	"strings"
	"time"

	"bosh-admin/core/ctx"
	"bosh-admin/core/db"
	"bosh-admin/core/exception"
	"bosh-admin/global"
	"bosh-admin/model"

	"github.com/duke-git/lancet/v2/random"
	"github.com/duke-git/lancet/v2/slice"
	"github.com/golang-jwt/jwt/v5"
)

type JWTSvc struct {
	accessSecret    []byte // access token密钥
	refreshSecret   []byte // refresh token密钥
	accessDuration  int64  // access token有效时长
	refreshDuration int64  // refresh token有效时长
	bufferDuration  int64  // 缓冲时长
}

func NewJWTSvc() *JWTSvc {
	return &JWTSvc{
		accessSecret:    []byte(global.Config.JWT.AccessSecret),
		refreshSecret:   []byte(global.Config.JWT.RefreshSecret),
		accessDuration:  global.Config.JWT.AccessDuration,
		refreshDuration: global.Config.JWT.RefreshDuration,
		bufferDuration:  global.Config.JWT.BufferDuration,
	}
}

func (svc *JWTSvc) GetAccessToken(c *ctx.Context) (string, error) {
	// 从请求头中获取Authorization
	headerAuth := c.Request.Header.Get("Authorization")
	if headerAuth == "" {
		return "", exception.NewException("请求头中未携带Authorization")
	}
	// 分割Authorization
	authParts := strings.Split(headerAuth, " ")
	if len(authParts) != 2 || authParts[0] != "Bearer" {
		return "", exception.NewException("请求头中Authorization格式有误")
	}
	return authParts[1], nil
}

// enhanceClaims 完善补充RegisteredClaims
func enhanceClaims(claims *jwt.RegisteredClaims, issuer string, duration int64) int64 {
	nowTime := time.Now().Local()
	expiresAt := nowTime.Add(time.Duration(duration) * time.Second)
	claims.Issuer = issuer
	claims.IssuedAt = jwt.NewNumericDate(nowTime)
	claims.NotBefore = jwt.NewNumericDate(nowTime)
	claims.ExpiresAt = jwt.NewNumericDate(expiresAt)
	return expiresAt.UnixMilli()
}

// GenerateUserAccessToken 生成用户access token
func (svc *JWTSvc) GenerateUserAccessToken(claims *UserAccessClaims) (string, int64, error) {
	expiresAt := enhanceClaims(&claims.RegisteredClaims, global.Config.Server.Name, svc.accessDuration)
	// 使用指定的签名方法签名token
	token := jwt.NewWithClaims(jwt.SigningMethodHS256, claims)
	// 加密生成token字符串
	tokenStr, err := token.SignedString(svc.accessSecret)
	return tokenStr, expiresAt, err
}

// GenerateMemberAccessToken 生成成员access token
func (svc *JWTSvc) GenerateMemberAccessToken(claims *MemberAccessClaims) (string, int64, error) {
	expiresAt := enhanceClaims(&claims.RegisteredClaims, global.Config.Server.Name, svc.accessDuration)
	// 使用指定的签名方法签名token
	token := jwt.NewWithClaims(jwt.SigningMethodHS256, claims)
	// 加密生成token字符串
	tokenStr, err := token.SignedString(svc.accessSecret)
	return tokenStr, expiresAt, err
}

// GenerateRefreshToken 生成refresh token
func (svc *JWTSvc) GenerateRefreshToken(claims *RefreshTokenClaims) (string, int64, error) {
	expiresAt := enhanceClaims(&claims.RegisteredClaims, global.Config.Server.Name, svc.refreshDuration)
	// 使用指定的签名方法签名token
	token := jwt.NewWithClaims(jwt.SigningMethodHS256, claims)
	// 加密生成token字符串
	tokenStr, err := token.SignedString(svc.refreshSecret)
	return tokenStr, expiresAt, err
}

// ParseUserAccessToken 解析用户access token
func (svc *JWTSvc) ParseUserAccessToken(tokenStr string) (*UserAccessClaims, error) {
	token, err := jwt.ParseWithClaims(tokenStr, &UserAccessClaims{}, func(token *jwt.Token) (any, error) {
		return svc.accessSecret, nil
	})
	if err != nil {
		return nil, err
	}
	if token != nil && token.Valid {
		if claims, ok := token.Claims.(*UserAccessClaims); ok {
			return claims, nil
		}
	}
	return nil, exception.NewException("无效的用户access token")
}

// ParseMemberAccessToken 解析成员access token
func (svc *JWTSvc) ParseMemberAccessToken(tokenStr string) (*MemberAccessClaims, error) {
	token, err := jwt.ParseWithClaims(tokenStr, &MemberAccessClaims{}, func(token *jwt.Token) (any, error) {
		return svc.accessSecret, nil
	})
	if err != nil {
		return nil, err
	}
	if token != nil && token.Valid {
		if claims, ok := token.Claims.(*MemberAccessClaims); ok {
			return claims, nil
		}
	}
	return nil, exception.NewException("无效的成员access token")
}

// ParseRefreshToken 解析refresh token
func (svc *JWTSvc) ParseRefreshToken(tokenStr string) (*RefreshTokenClaims, error) {
	token, err := jwt.ParseWithClaims(tokenStr, &RefreshTokenClaims{}, func(token *jwt.Token) (any, error) {
		return svc.refreshSecret, nil
	})
	if err != nil {
		return nil, err
	}
	if token != nil && token.Valid {
		if claims, ok := token.Claims.(*RefreshTokenClaims); ok {
			return claims, nil
		}
	}
	return nil, exception.NewException("无效refresh token")
}

// TokenValidate token校验
func (svc *JWTSvc) TokenValidate(claims jwt.RegisteredClaims, subject, audience string) error {
	if claims.ID == "" {
		return exception.NewException("无效token")
	}
	if claims.Issuer != global.Config.Server.Name {
		return exception.NewException("无效token")
	}
	if claims.Subject != subject {
		return exception.NewException("无效token")
	}
	if !(slices.Contains(claims.Audience, "all") || slices.Contains(claims.Audience, audience)) {
		return exception.NewException("无效token")
	}
	return nil
}

// GetUserAccessClaims 获取UserAccessClaims
func (svc *JWTSvc) GetUserAccessClaims(c *ctx.Context) (*UserAccessClaims, error) {
	tokenStr, err := svc.GetAccessToken(c)
	if err != nil {
		return nil, err
	}
	return svc.ParseUserAccessToken(tokenStr)
}

// GetUserClaims 获取UserClaims
func (svc *JWTSvc) GetUserClaims(c *ctx.Context) *UserClaims {
	if claims, exists := c.Get("userAccessClaims"); !exists {
		if cl, err := svc.GetUserAccessClaims(c); err != nil {
			return nil
		} else {
			return cl.User
		}
	} else {
		waitUse := claims.(*UserAccessClaims)
		return waitUse.User
	}
}

// GetMemberAccessClaims 获取MemberAccessClaims
func (svc *JWTSvc) GetMemberAccessClaims(c *ctx.Context) (*MemberAccessClaims, error) {
	tokenStr, err := svc.GetAccessToken(c)
	if err != nil {
		return nil, err
	}
	return svc.ParseMemberAccessToken(tokenStr)
}

// GetMemberClaims 获取MemberClaims
func (svc *JWTSvc) GetMemberClaims(c *ctx.Context) *MemberClaims {
	if claims, exists := c.Get("memberAccessClaims"); !exists {
		if cl, err := svc.GetMemberAccessClaims(c); err != nil {
			return nil
		} else {
			return cl.Member
		}
	} else {
		waitUse := claims.(*MemberAccessClaims)
		return waitUse.Member
	}
}

func (svc *JWTSvc) UserLogin(user *model.SysUser) (string, string, int64, error) {
	var claims = &UserAccessClaims{
		User: &UserClaims{
			UserId:   user.Id,
			Username: user.Username,
			Nickname: user.Nickname,
			RoleId:   user.RoleId,
			RoleCode: user.Role.RoleCode,
			DataAuth: user.Role.DataAuth,
			DeptId:   user.DeptId,
			DeptCode: user.Dept.DeptCode,
			DeptPath: user.Dept.DeptPath,
		},
	}
	uuid, err := random.UUIdV4()
	if err != nil {
		return "", "", 0, exception.NewException("生成UUID失败")
	}
	claims.ID = uuid
	var audience []string
	if user.Role.RoleCode == global.SuperAdmin && user.Dept.DeptCode == global.SystemAdmin {
		audience = []string{JwtAudienceAll}
	} else {
		audience = []string{JwtAudienceApi, JwtAudienceStatic}
	}
	claims.Audience = audience
	claims.Subject = JwtSubjectAccess
	accessToken, expiresAt, err := svc.GenerateUserAccessToken(claims)
	if err != nil {
		return "", "", 0, err
	}
	// 为refresh token创建简化的claims，使用相同的UUID
	refreshClaims := &RefreshTokenClaims{}
	refreshClaims.ID = uuid // 使用相同的UUID，便于拉黑时同时失效
	refreshClaims.Subject = JwtSubjectRefresh
	refreshClaims.Audience = audience
	refreshClaims.BindId = user.Id
	refreshToken, _, err := svc.GenerateRefreshToken(refreshClaims)
	if err != nil {
		return "", "", 0, err
	}
	return accessToken, refreshToken, expiresAt, nil
}

// RefreshToken 刷新token
func (svc *JWTSvc) RefreshToken(refreshToken string) (string, string, int64, error) {
	// 验证refresh token
	refreshClaims, err := svc.ParseRefreshToken(refreshToken)
	if err != nil {
		return "", "", 0, err
	}
	// refresh token 临近过期则刷新
	if refreshClaims.ExpiresAt.Unix()-time.Now().Local().Unix() <= svc.bufferDuration {
		refreshToken, _, err = svc.GenerateRefreshToken(refreshClaims)
		if err != nil {
			return "", "", 0, err
		}
	}
	var accessToken string
	var expiresAt int64
	if slice.Contain(refreshClaims.Audience, JwtAudienceAll) || slice.Contain(refreshClaims.Audience, JwtAudienceApi) {
		// 获取用户信息
		var user *model.SysUser
		err = db.GormDB().Where("id = ?", refreshClaims.BindId).Preload("Role").Preload("Dept").First(&user).Error
		if err != nil {
			return "", "", 0, err
		}
		var userClaims = &UserClaims{
			UserId:   user.Id,
			Username: user.Username,
			Nickname: user.Nickname,
			RoleId:   user.RoleId,
			RoleCode: user.Role.RoleCode,
			DataAuth: user.Role.DataAuth,
			DeptId:   user.DeptId,
			DeptCode: user.Dept.DeptCode,
			DeptPath: user.Dept.DeptPath,
		}
		userAccessClaims := new(UserAccessClaims)
		userAccessClaims.ID = refreshClaims.ID
		userAccessClaims.Audience = refreshClaims.Audience
		userAccessClaims.Subject = JwtSubjectAccess
		userAccessClaims.User = userClaims
		accessToken, expiresAt, err = svc.GenerateUserAccessToken(userAccessClaims)
	}
	if slice.Contain(refreshClaims.Audience, JwtAudienceApp) {
		// 获取成员信息
	}
	if err != nil {
		return "", "", 0, err
	}
	return accessToken, refreshToken, expiresAt, nil
}
