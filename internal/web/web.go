package web

import (
	"embed"
	"encoding/json"
	"fmt"
	"gateway/internal/file"
	"io"
	"io/fs"
	"net/http"

	"github.com/gin-contrib/gzip"
	"github.com/gin-gonic/gin"
)

//go:embed assets/favicon.png
var Favicon embed.FS

//go:embed pages/index.html
var HomePage []byte

//go:embed pages/confirm.html
var ConfirmPage []byte

//go:embed assets
var Assets embed.FS

const (
	API_DATA    = "/api/data"
	API_COMPARE = "/api/compare"
	API_PREPARE = "/api/prepare"
)

func getConfig() string {
	config := map[string]string{
		"Data":    API_DATA,
		"Compare": API_COMPARE,
		"Prepare": API_PREPARE,
	}
	b, err := json.Marshal(config)
	if err != nil {
		return "{}"
	}
	return string(b)
}

func API(port string) {
	gin.SetMode(gin.ReleaseMode)
	r := gin.Default()
	r.Use(gzip.Gzip(gzip.DefaultCompression))
	r.Use(optimizeResourceCacheTime([]string{"/favicon.png", "/assets/"}))

	r.Any("/", func(c *gin.Context) {
		c.Data(http.StatusOK, "text/html; charset=utf-8", HomePage)
	})

	r.Any("/confirm", func(c *gin.Context) {
		c.Data(http.StatusOK, "text/html; charset=utf-8", ConfirmPage)
	})

	favicon, _ := fs.Sub(Favicon, "assets")
	r.Any("/favicon.png", func(c *gin.Context) {
		c.FileFromFS("favicon.png", http.FS(favicon))
	})

	css, _ := fs.Sub(Assets, "assets/css")
	r.StaticFS("/assets/css", http.FS(css))

	js, _ := fs.Sub(Assets, "assets/js")
	r.StaticFS("/assets/js", http.FS(js))

	r.GET(API_DATA, func(c *gin.Context) {
		c.Data(http.StatusOK, "plain/text; charset=utf-8", []byte(file.GetHostsFileContent()))
		c.Abort()
	})

	r.GET(API_COMPARE, func(c *gin.Context) {
		c.JSON(200, gin.H{
			"status":  "posted",
			"data":    file.GetHostsFileContent(),
			"prepare": file.ReadPrepareFile(),
		})
		c.Abort()
	})

	r.POST(API_PREPARE, func(c *gin.Context) {
		body, err := io.ReadAll(c.Request.Body)
		if err != nil {
			// TODO: 提示处理出错
			fmt.Println(err)
			c.Data(http.StatusOK, "plain/text; charset=utf-8", []byte("请求有问题"))
			c.Abort()
			return
		}

		if body == nil {
			// TODO: 提示处理出错
			fmt.Println(err)
			c.Data(http.StatusOK, "plain/text; charset=utf-8", []byte("请求有问题"))
			c.Abort()
			return
		}

		success := file.SaveHostsFileContent(body)
		if success {
			c.Data(http.StatusOK, "plain/text; charset=utf-8", []byte("保存成功"))
		} else {
			c.Data(http.StatusOK, "plain/text; charset=utf-8", []byte("保存失败"))
		}
		c.Abort()
	})

	r.GET("/api/config.js", func(c *gin.Context) {
		c.Data(http.StatusOK, "application/javascript; charset=utf-8", []byte(`window.$API$ = `+getConfig()))
		c.Abort()
	})

	r.Run(":" + port)
}
