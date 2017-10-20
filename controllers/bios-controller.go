package controllers

import (
	"io/ioutil"
	"net/http"

	"github.com/gin-gonic/gin"
	"gitlab.com/MikaXII/recalbox-api/config"
	"gitlab.com/MikaXII/recalbox-api/models"
)

// const ROM_DIR = "/recalbox/share/roms/"
var biosEndpoint string
var biosPath string

// BiosGroupV1 regroup path for v1 endpoint
func BiosGroupV1(r *gin.RouterGroup, config *configuration.Configuration) {
	biosEndpoint = config.ListEndpoint.BiosEndpoint
	biosPath = config.Fs.BiosPath
	r.GET(biosEndpoint, getListBios)
}

// getListBios -> get Bios list...
func getListBios(c *gin.Context) {
	listFiles := []models.Media{}
	files, _ := ioutil.ReadDir(biosPath)
	for _, f := range files {
		filePath := biosPath + "/" + f.Name()
		if f.IsDir() {
		} else {
			listFiles = append(listFiles, *models.NewMedia(filePath, f))
		}
	}
	c.JSON(http.StatusOK, listFiles)
}
