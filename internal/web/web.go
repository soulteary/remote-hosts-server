package web

import (
	"embed"
	"encoding/json"
	"fmt"
	"gateway/internal/file"
	"io"
	"io/fs"
	"net/http"
	"strings"

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
	API_DATA    = "/api/hosts"
	API_DIFF = "/api/diff"
	API_PREPARE = "/api/prepare"
	PAGE_DIFF   = "/confirm"
)

func getConfig() string {
	config := map[string]string{
		"Data":    API_DATA,
		"Diff": API_DIFF,
		"Prepare": API_PREPARE,
	}
	b, err := json.Marshal(config)
	if err != nil {
		return "{}"
	}
	return string(b)
}

func API(port string, mode string) {
	gin.SetMode(gin.ReleaseMode)
	r := gin.Default()
	r.Use(gzip.Gzip(gzip.DefaultCompression))
	r.Use(optimizeResourceCacheTime([]string{"/favicon.png", "/assets/"}))

	r.Any("/", func(c *gin.Context) {
		c.Data(http.StatusOK, "text/html; charset=utf-8", HomePage)
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
		c.Data(http.StatusOK, "plain/text; charset=utf-8", []byte(file.GetHostsFileContent("stable")))
		c.Abort()
	})

	if strings.ToUpper(mode) != "SIMPLE" {
		r.Any(PAGE_DIFF, func(c *gin.Context) {
			c.Data(http.StatusOK, "text/html; charset=utf-8", ConfirmPage)
		})

		r.GET(API_DIFF, func(c *gin.Context) {
			c.JSON(200, gin.H{
				"code":    0,
				"data":    file.GetHostsFileContent("stable"),
				"prepare": file.GetHostsFileContent("prepare"),
			})
			c.Abort()
		})
	}

	r.POST(API_PREPARE, func(c *gin.Context) {
		body, err := io.ReadAll(c.Request.Body)
		if err != nil {
			c.JSON(200, gin.H{
				"code":    1,
				"message": "Error: There is a problem with submitting the data.",
			})
			c.Abort()
			fmt.Println(err)
			return
		}

		if body == nil {
			c.JSON(200, gin.H{
				"code":    2,
				"message": "Error: No commit data detected.",
			})
			c.Abort()
			return
		}

		success := false
		if mode == "SIMPLE" {
			success = file.SaveHostsFileContent("stable", body)
		} else {
			success = file.SaveHostsFileContent("prepare", body)
		}

		if success {
			if mode == "SIMPLE" {
				c.JSON(200, gin.H{
					"code":    0,
					"message": "Hosts data saved successfully.",
				})
			} else {
				c.JSON(200, gin.H{
					"code":    0,
					"message": "The data submission is successful.",
					"next":    PAGE_DIFF,
				})
			}
		} else {
			c.JSON(200, gin.H{
				"code":    0,
				"message": "Failed to save data",
			})
		}
		c.Abort()
	})

	r.GET("/api/config.js", func(c *gin.Context) {
		c.Data(http.StatusOK, "application/javascript; charset=utf-8", []byte(`window.$API$ = `+getConfig()))
		c.Abort()
	})

	r.Run(":" + port)
}
