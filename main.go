package main

import (
	"github-missing-api/trending"
	"net/http"
	"os"
	"time"

	"github.com/gin-gonic/gin"
	"github.com/patrickmn/go-cache"
)

func main() {
	var port = "8080"
	if _port := os.Getenv("PORT"); _port != "" {
		port = _port
	}
	var goCache = cache.New(1*time.Hour, 1*time.Hour)
	r := gin.Default()
	trend := r.Group("/trending")
	{
		trend.GET("/languages", func(c *gin.Context) {
			c.JSON(http.StatusOK, trending.Langs)
		})
		trend.GET("/repositories", func(c *gin.Context) {
			var language = c.Query("language")
			var since = c.DefaultQuery("since", "daily")
			var cacheKey = "repositories::nolang" + since
			if language != "" {
				cacheKey = "repositories::" + language + since
			}
			if json, found := goCache.Get(cacheKey); found {
				c.JSON(http.StatusOK, json)
			} else {
				var json = trending.FetchRepositories(language, since)
				goCache.Set(cacheKey, json, cache.DefaultExpiration)
				c.JSON(http.StatusOK, json)
			}
		})
	}
	r.Run(":" + port) // listen and serve on 0.0.0.0:8080
}
