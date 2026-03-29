package utils

import (
	"strings"

	"github.com/gin-gonic/gin"
)

func GetClientIp(c *gin.Context) string {
	clientIP := c.ClientIP()
	remoteIP := c.RemoteIP()
	ip := c.Request.Header.Get("X-Forwarded-For")
	if strings.Contains(ip, "127.0.0.1") || ip == "" {
		ip = c.Request.Header.Get("X-Real-IP")
	}
	if ip == "" {
		ip = "127.0.0.1"
	}
	if remoteIP != "127.0.0.1" {
		ip = remoteIP
	}
	if clientIP != "127.0.0.1" {
		ip = clientIP
	}
	return ip

}
