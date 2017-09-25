package recalroutes

import (
	"github.com/gin-gonic/gin"
)

func ApiInfo(r *gin.Engine, version string) {
	r.GET("/", func(c *gin.Context) {
		c.JSON(200, gin.H{"version": version})
	})

}
