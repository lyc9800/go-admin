package middleware

import (
	"bytes"
	"encoding/json"
	"fmt"
	"io"
	"net/http"
	"strings"
	"time"

	"server/models"

	"github.com/gin-gonic/gin"
)

// ResponseWriter 自定义的 ResponseWriter
type ResponseWriter struct {
	gin.ResponseWriter
	body *bytes.Buffer
}

func (w ResponseWriter) Write(b []byte) (int, error) {
	if w.body != nil {
		w.body.Write(b)
	}
	return w.ResponseWriter.Write(b)
}

// SimpleLogger 安全的日志中间件（不包含 IP 查询）
func SimpleLogger() gin.HandlerFunc {
	return func(c *gin.Context) {
		// 开始时间
		startTime := time.Now()

		// 排除不需要记录的路由
		requestURL := c.Request.URL.Path
		if shouldSkipLog(requestURL) {
			c.Next()
			return
		}

		// 创建自定义的 ResponseWriter
		writer := ResponseWriter{
			ResponseWriter: c.Writer,
			body:           bytes.NewBuffer([]byte{}),
		}
		c.Writer = &writer

		// 处理请求
		c.Next()

		// 结束时间
		useTime := time.Since(startTime).Milliseconds()

		// 获取客户端IP
		clientIP := c.ClientIP()

		// 获取操作系统和浏览器信息
		os, browser := getSafeOSAndBrowser(c)
		fmt.Printf("=== 浏览器信息调试 ===\n")
		fmt.Printf("User-Agent: %s\n", c.Request.Header.Get("User-Agent"))
		fmt.Printf("识别结果: OS=%s, Browser=%s\n", os, browser)
		fmt.Printf("========================\n")
		// 构建日志记录
		sysLog := models.SysLog{
			Browser:      fmt.Sprintf("%s_%s", os, browser),
			RequestUri:   requestURL,
			ClassMethod:  c.HandlerName(),
			HttpMethod:   c.Request.Method,
			UseTime:      useTime,
			Ip:           clientIP,
			StatusCode:   c.Writer.Status(),
			Result:       truncateString(writer.body.String(), 1000),
			RequestParam: getSafeRequestParam(c),
			Country:      getCountryByIP(clientIP),
			Province:     getProvinceByIP(clientIP),
			City:         getCityByIP(clientIP),
			Region:       "本地",
			Isp:          "本地",
		}

		// 安全保存到数据库
		go func(log models.SysLog) {
			defer func() {
				if r := recover(); r != nil {
					fmt.Printf("保存日志时发生panic: %v\n", r)
				}
			}()

			err := models.DB.Create(&log).Error
			if err != nil {
				fmt.Printf("保存日志失败: %v\n", err)
			} else {
				fmt.Printf("日志保存成功: %s %s (ID: %d)\n",
					log.HttpMethod, log.RequestUri, log.ID)
			}
		}(sysLog)
	}
}

// 安全的操作系统和浏览器信息获取
func getSafeOSAndBrowser(c *gin.Context) (string, string) {
	defer func() {
		if r := recover(); r != nil {
			fmt.Printf("获取OS/Browser时panic: %v\n", r)
		}
	}()

	userAgent := c.Request.Header.Get("User-Agent")
	if userAgent == "" {
		return "未知", "未知"
	}

	ua := strings.ToLower(userAgent)
	var os, browser string

	// 检测操作系统
	switch {
	case strings.Contains(ua, "windows"):
		os = "Windows"
	case strings.Contains(ua, "mac"):
		os = "macOS"
	case strings.Contains(ua, "linux"):
		os = "Linux"
	case strings.Contains(ua, "android"):
		os = "Android"
	case strings.Contains(ua, "iphone") || strings.Contains(ua, "ipad"):
		os = "iOS"
	default:
		os = "未知"
	}

	// 检测浏览器
	switch {
	case strings.Contains(ua, "chrome") && !strings.Contains(ua, "edg"):
		browser = "Chrome"
	case strings.Contains(ua, "firefox"):
		browser = "Firefox"
	case strings.Contains(ua, "safari") && !strings.Contains(ua, "chrome"):
		browser = "Safari"
	case strings.Contains(ua, "edge") || strings.Contains(ua, "edg"):
		browser = "Edge"
	default:
		browser = "未知"
	}

	return os, browser
}

// 安全的请求参数获取
func getSafeRequestParam(c *gin.Context) string {
	defer func() {
		if r := recover(); r != nil {
			fmt.Printf("获取请求参数时panic: %v\n", r)
		}
	}()

	if c.Request.Method == http.MethodGet {
		query := c.Request.URL.RawQuery
		if query != "" {
			return truncateString(query, 500)
		}
		return ""
	}

	// 对于POST等请求，尝试读取body
	body, err := io.ReadAll(c.Request.Body)
	if err != nil {
		return ""
	}

	// 重新设置请求体
	c.Request.Body = io.NopCloser(bytes.NewBuffer(body))

	// 如果是JSON，尝试格式化
	if strings.Contains(c.Request.Header.Get("Content-Type"), "application/json") {
		var data interface{}
		if err := json.Unmarshal(body, &data); err == nil {
			if formatted, err := json.Marshal(data); err == nil {
				return truncateString(string(formatted), 500)
			}
		}
	}

	return truncateString(string(body), 500)
}

// 简单的IP地址解析
func getCountryByIP(ip string) string {
	if ip == "::1" || ip == "127.0.0.1" {
		return "本地"
	}
	return "中国"
}

func getProvinceByIP(ip string) string {
	if ip == "::1" || ip == "127.0.0.1" {
		return "本地"
	}
	return "未知"
}

func getCityByIP(ip string) string {
	if ip == "::1" || ip == "127.0.0.1" {
		return "本地"
	}
	return "未知"
}

// 判断是否需要跳过日志记录
func shouldSkipLog(url string) bool {
	skipPaths := []string{
		"/favicon.ico",
		"/static/",
		"/uploads/",
		"/images/",
		"/css/",
		"/js/",
		"/log",
		"/uploadFile",
		"/api/log", // 添加日志接口
	}

	for _, path := range skipPaths {
		if strings.Contains(url, path) {
			return true
		}
	}
	return false
}

// 截断字符串
func truncateString(str string, maxLength int) string {
	if len(str) > maxLength {
		return str[:maxLength] + "..."
	}
	return str
}
