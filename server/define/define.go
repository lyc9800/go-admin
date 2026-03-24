package define

import (
	"time"

	"github.com/golang-jwt/jwt/v5"
)

var (
	// 密钥
	Jwtkey = "sys-admin"
	// token过期时间,7天过期
	TokenExpire = time.Now().Add(time.Second * 3600 * 24 * 7).Unix()
	// 刷新token过期时间,14天过期
	RefreshTokenExpire = time.Now().Add(time.Second * 3600 * 24 * 14).Unix()
	// 默认分页显示的条数
	DefaultSize = 10
	// 头像保存目录
	StaticResource = "static/resource"
)

type UserClaim struct {
	Id      uint
	Name    string
	IsAdmin bool // 是否为管理员
	RoleId  uint
	jwt.RegisteredClaims
}
