package main

import (
	"net/http"
	"time"

	"github.com/jetsly/github-missing-api/trending"

	"github.com/gin-gonic/gin"
	"github.com/patrickmn/go-cache"
)

func getCacheKey(language string, since string, category string) string {
	var cacheKey = "nolang" + since
	if language != "" {
		cacheKey = "" + language + since
	}
	return category + "::" + cacheKey
}

func main() {
	const port = "8000"
	goCache := cache.New(1*time.Hour, 1*time.Hour)
	r := gin.Default()
	trend := r.Group("/trending")
	{
		trend.GET("/languages", func(c *gin.Context) {
			c.JSON(http.StatusOK, trending.Langs)
		})
		trend.GET("/repositories", func(c *gin.Context) {
			var language = c.Query("language")
			var since = c.DefaultQuery("since", "daily")
			var cacheKey = getCacheKey(language, since, "repositories")
			if json, found := goCache.Get(cacheKey); found {
				c.JSON(http.StatusOK, json)
			} else {
				var json = trending.FetchRepositories(language, since)
				goCache.Set(cacheKey, json, cache.DefaultExpiration)
				c.JSON(http.StatusOK, json)
			}
		})
		trend.GET("/developers", func(c *gin.Context) {
			var language = c.Query("language")
			var since = c.DefaultQuery("since", "daily")
			var cacheKey = getCacheKey(language, since, "developers")
			if json, found := goCache.Get(cacheKey); found {
				c.JSON(http.StatusOK, json)
			} else {
				var json = trending.FetchDevelopers(language, since)
				goCache.Set(cacheKey, json, cache.DefaultExpiration)
				c.JSON(http.StatusOK, json)
			}
		})
	}
	r.Run(":" + port) // listen and serve on 0.0.0.0:8080
}
