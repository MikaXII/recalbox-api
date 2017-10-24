package main

import (
	"log"

	"github.com/gin-contrib/cors"
	"github.com/gin-gonic/gin"
	"gitlab.com/MikaXII/recalbox-api/recalapi"
)

func main() {

	err := GetMainEngine().Run()
	if err != nil {
		log.Fatal(err)
	}
}

// GetMainEngine get the main engine ?
func GetMainEngine() *gin.Engine {
	router := gin.Default()
	loadEndpoints(router)

	return router
}

// loadEndpoints load all endpoints
func loadEndpoints(r *gin.Engine) {
	r.Use(cors.Default())
	config := recalapi.LoadConfig(gin.Mode())
	authGuard := gin.Accounts{"foo": "bar"}
	basicAuth := gin.BasicAuth(authGuard)
	// version, _ := semver.Make(config.Version)
	// 	rangeV1, _ := semver.ParseRange(">0.0.0")
	//	rangeV2, _ := semver.ParseRange(">=2.0.0 <3.0.0")

	recalapi.APIInfo(r, config.Version)
	v1 := r.Group("/v1", basicAuth)
	{
		recalapi.RomGroupV1(v1, config)
		recalapi.BiosGroupV1(v1, config)
	}
}
