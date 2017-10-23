package controllers

import (
	"github.com/gin-contrib/static"
	"github.com/gin-gonic/gin"
)

// APIInfo set root endpoint with all info
func APIInfo(r *gin.Engine, version string) {
	r.Use(static.Serve("/", static.LocalFile("client/dist", true)))

	r.GET("/status", func(c *gin.Context) {
		var listRoutes []string
		for _, route := range r.Routes() {
			listRoutes = append(listRoutes, route.Path)
		}
		c.JSON(200, gin.H{"version": version, "listAvailableEndpoint": listRoutes})
	})

}
