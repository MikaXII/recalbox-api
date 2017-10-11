package main

import (
	"log"

	"github.com/gin-gonic/gin"
	"gitlab.com/MikaXII/recalbox-api/config"
	"gitlab.com/MikaXII/recalbox-api/controllers"
)

func main() {

	router := gin.Default()
	loadEndpoints(router)

	err := router.Run()
	if err != nil {
		log.Fatal(err)
	}
}

func loadEndpoints(r *gin.Engine) {

	config := configuration.LoadConfig(gin.Mode())
	authGuard := gin.Accounts{"foo": "bar"}
	basicAuth := gin.BasicAuth(authGuard)
	// version, _ := semver.Make(config.Version)
	// 	rangeV1, _ := semver.ParseRange(">0.0.0")
	//	rangeV2, _ := semver.ParseRange(">=2.0.0 <3.0.0")

	controllers.APIInfo(r, config.Version)
	v1 := r.Group("/v1", basicAuth)
	{
		controllers.RomGroupV1(v1, config)
		controllers.BiosGroupV1(v1, config)
	}
}
