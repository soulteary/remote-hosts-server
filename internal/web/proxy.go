package web

import (
	"crypto/md5"
	"fmt"
	"net/http"
	"strings"
	"time"

	"github.com/gin-gonic/gin"
)

func anyResMatched(url string, matches []string) bool {
	for _, match := range matches {
		exist := strings.HasPrefix(url, match)
		if exist {
			return true
		}
	}
	return false
}

func optimizeResourceCacheTime(cacheRes []string) gin.HandlerFunc {
	// ViewHandler support dist handler from UI
	// https://github.com/gin-gonic/gin/issues/1222
	data := []byte(time.Now().String())
	/* #nosec */
	etag := fmt.Sprintf("W/%x", md5.Sum(data))
	return func(c *gin.Context) {
		if anyResMatched(c.Request.RequestURI, cacheRes) {
			c.Header("Cache-Control", "public, max-age=31536000")
			c.Header("ETag", etag)
			if match := c.GetHeader("If-None-Match"); match != "" {
				if strings.Contains(match, etag) {
					c.Status(http.StatusNotModified)
					return
				}
			}
		}
	}
}
