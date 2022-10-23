package web

import (
	"embed"
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

	r.GET("/api/data.txt", func(c *gin.Context) {
		c.Data(http.StatusOK, "plain/text; charset=utf-8", []byte(file.ReadFile()))
		c.Abort()
	})

	r.GET("/api/diff", func(c *gin.Context) {
		c.JSON(200, gin.H{
			"status":  "posted",
			"data":    file.ReadFile(),
			"prepare": file.ReadPrepareFile(),
		})
		c.Abort()
	})

	r.POST("/api/prepare", func(c *gin.Context) {
		body, _ := io.ReadAll(c.Request.Body)
		if body != nil {
			c.Data(http.StatusOK, "plain/text; charset=utf-8", []byte(body))
		} else {
			c.Data(http.StatusOK, "plain/text; charset=utf-8", []byte("请求有问题"))
		}
		c.Abort()

	})

	r.Run(":" + port)
}
