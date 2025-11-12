package middleware

import (
	"bytes"
	"io"
	"net/http"
	"net/url"
	"strings"
	"sync"
	"time"

	"bosh-admin/core/ctx"
	"bosh-admin/core/db"
	"bosh-admin/core/log"
	"bosh-admin/model"
	"bosh-admin/module/auth"
	"bosh-admin/util"

	"github.com/gin-gonic/gin"
	jsoniter "github.com/json-iterator/go"
	ua "github.com/mssola/user_agent"
)

var respPool sync.Pool
var bufferSize = 1024

func init() {
	respPool.New = func() interface{} {
		return make([]byte, bufferSize)
	}
}

func OperationRecord() gin.HandlerFunc {
	return ctx.Handler(func(c *ctx.Context) {
		var body []byte
		if c.Request.Method != http.MethodGet {
			var err error
			body, err = io.ReadAll(c.Request.Body)
			if err != nil {
				log.Error("读取请求体失败:", err.Error())
			} else {
				c.Request.Body = io.NopCloser(bytes.NewBuffer(body))
			}
		} else {
			query := c.Request.URL.RawQuery
			query, _ = url.QueryUnescape(query)
			split := strings.Split(query, "&")
			m := make(map[string]string)
			for _, v := range split {
				kv := strings.Split(v, "=")
				if len(kv) == 2 {
					m[kv[0]] = kv[1]
				}
			}
			body, _ = jsoniter.Marshal(&m)
		}
		var userId uint
		var username string
		claims := auth.NewJWTSvc().GetUserClaims(c)
		if claims != nil && claims.UserId != 0 {
			userId = claims.UserId
			username = claims.Username
		}
		requestIP := c.ClientIP()
		userAgent := c.Request.UserAgent()
		record := model.SysOperationRecord{
			Uid:       userId,
			Username:  username,
			Method:    c.Request.Method,
			Path:      c.Request.URL.Path,
			UserAgent: userAgent,
			RequestIP: requestIP,
		}
		record.RequestRegion = util.IP2Region(requestIP)
		UA := ua.New(userAgent)
		record.RequestOS = UA.OS()
		record.RequestBrowser, _ = UA.Browser()
		requestHeader, err := jsoniter.Marshal(c.Request.Header)
		if err != nil {
			log.Error("解析请求头失败:", err.Error())
		}
		record.RequestHeader = string(requestHeader)

		// 上传文件时候 中间件日志进行裁断操作
		if strings.Contains(c.GetHeader("Content-Type"), "multipart/form-data") {
			record.RequestBody = "[文件]"
		} else {
			if len(body) > bufferSize {
				record.RequestBody = "[超出记录长度]"
			} else {
				record.RequestBody = string(body)
			}
		}

		writer := responseBodyWriter{
			ResponseWriter: c.Writer,
			body:           &bytes.Buffer{},
		}
		c.Writer = writer
		now := time.Now()

		c.Next()

		latency := time.Since(now).Milliseconds()
		record.Latency = latency
		record.Status = c.Writer.Status()
		record.ResponseBody = writer.body.String()
		responseHeader, err := jsoniter.Marshal(writer.Header())
		if err != nil {
			log.Error("解析响应头失败:", err.Error())
		}
		record.ResponseHeader = string(responseHeader)

		if strings.Contains(c.Writer.Header().Get("Pragma"), "public") ||
			strings.Contains(c.Writer.Header().Get("Expires"), "0") ||
			strings.Contains(c.Writer.Header().Get("Cache-Control"), "must-revalidate, post-check=0, pre-check=0") ||
			strings.Contains(c.Writer.Header().Get("Content-Type"), "application/force-download") ||
			strings.Contains(c.Writer.Header().Get("Content-Type"), "application/octet-stream") ||
			strings.Contains(c.Writer.Header().Get("Content-Type"), "application/vnd.ms-excel") ||
			strings.Contains(c.Writer.Header().Get("Content-Type"), "application/download") ||
			strings.Contains(c.Writer.Header().Get("Content-Disposition"), "attachment") ||
			strings.Contains(c.Writer.Header().Get("Content-Transfer-Encoding"), "binary") {
			if len(record.ResponseBody) > bufferSize {
				// 截断
				record.ResponseBody = "[超出记录长度]"
			}
		}

		if err = db.Create(&record); err != nil {
			log.Error("保存操作记录失败:", err.Error())
		}
	})
}

type responseBodyWriter struct {
	gin.ResponseWriter
	body *bytes.Buffer
}

func (r responseBodyWriter) Write(b []byte) (int, error) {
	r.body.Write(b)
	return r.ResponseWriter.Write(b)
}
