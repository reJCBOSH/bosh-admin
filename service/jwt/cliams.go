package jwt

import "github.com/golang-jwt/jwt/v5"

const (
	SubjectAccess  = "access"
	SubjectRefresh = "refresh"
	AudienceAll    = "all"
	AudienceApi    = "api"
	AudienceApp    = "app"
	AudienceStatic = "static"
)

type UserClaims struct {
	UserId   uint   `json:"userId"`   // 用户id
	Username string `json:"username"` // 用户名
	Nickname string `json:"nickname"` // 昵称
}

type UserAccessClaims struct {
	jwt.RegisteredClaims
	User *UserClaims
}

type MemberClaims struct {
	MemberId uint `json:"memberId"` // 成员id
}

type MemberAccessClaims struct {
	jwt.RegisteredClaims
	Member *MemberClaims
}

// RefreshTokenClaims 用于refresh token的claims，只包含关联access token的信息
type RefreshTokenClaims struct {
	jwt.RegisteredClaims
	BindId uint `json:"bindId"`
}
