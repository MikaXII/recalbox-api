package main

import (
	"github.com/blang/semver"
	"github.com/gin-gonic/gin"
	"gitlab.com/MikaXII/recalbox-api/config"
	"gitlab.com/MikaXII/recalbox-api/controllers"
)

func main() {

	router := gin.Default()
	loadEndpoints(router)
	router.Run()
}

func loadEndpoints(r *gin.Engine) {

	config := configuration.LoadConfig(gin.Mode())

	version, _ := semver.Make(config.Version)
	rangeV1, _ := semver.ParseRange(">0.0.0")
	rangeV2, _ := semver.ParseRange(">=2.0.0 <3.0.0")

	controllers.ApiInfo(r, config.Version)
	if rangeV1(version) {
		v1 := r.Group("/v1")
		{
			controllers.RomGroupV1(v1, config)
		}
	}

	if rangeV2(version) {
		println("loaded ?")
	}
}
