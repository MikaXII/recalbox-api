package controllers

import (
	"github.com/gin-gonic/gin"
)

func ApiInfo(r *gin.Engine, version string) {
	r.GET("/", func(c *gin.Context) {
		var listRoutes []string
		for _, route := range r.Routes() {
			listRoutes = append(listRoutes, route.Path)
		}
		c.JSON(200, gin.H{"version": version, "listAvailableEndpoint": listRoutes})
	})

}
