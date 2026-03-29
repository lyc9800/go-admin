package define

import (
	"os"
	"path/filepath"
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
	StaticResource = "static"
	// 邮箱地址,授权码，地址
	EmailFrom     = "823654832@qq.com"
	EmailPassword = "tilalchlkcjbbfif"
	EmailHost     = "smtp.qq.com"
	EmailPort     = "587"
	// ip2region 存放路径
	Ip2RegionPath = getIP2RegionPath()
)

// getIP2RegionPath 获取 IP2Region 数据库文件的绝对路径
func getIP2RegionPath() string {
	// 获取当前工作目录
	wd, err := os.Getwd()
	if err != nil {
		return "static/ip2region_v4.xdb"
	}

	// 尝试多个可能的路径
	possiblePaths := []string{
		filepath.Join(wd, "static", "ip2region_v4.xdb"),
		filepath.Join(wd, "static/ip2region_v4.xdb"),
		filepath.Join(wd, "server", "static", "ip2region_v4.xdb"),
		"static/ip2region_v4.xdb",
		"./static/ip2region_v4.xdb",
		"server/static/ip2region_v4.xdb",
		"./server/static/ip2region_v4.xdb",
	}

	for _, path := range possiblePaths {
		if _, err := os.Stat(path); err == nil {
			absPath, _ := filepath.Abs(path)
			return absPath
		}
	}

	// 如果都没找到，返回相对路径
	return "static/ip2region_v4.xdb"
}

type UserClaim struct {
	Id      uint
	Name    string
	IsAdmin bool // 是否为管理员
	RoleId  uint
	jwt.RegisteredClaims
}
