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
	RoleId   uint   `json:"roleId"`   // 角色id
	RoleCode string `json:"roleCode"` // 角色标识
	DataAuth int    `json:"dataAuth"` // 数据权限
	DeptId   uint   `json:"deptId"`   // 部门id
	DeptCode string `json:"deptCode"` // 部门标识
	DeptPath string `json:"deptPath"` // 部门路径
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
