package controllers

import (
	"io/ioutil"
	"net/http"

	"github.com/gin-gonic/gin"
	"gitlab.com/MikaXII/recalbox-api/config"
	"gitlab.com/MikaXII/recalbox-api/utils"
)

// const ROM_DIR = "/recalbox/share/roms/"
var biosEndpoint string
var biosPath string

// Bios represents a bios file
type Bios struct {
	Name string
	Hash string
}

// BiosGroupV1 regroup path for v1 endpoint
func BiosGroupV1(r *gin.RouterGroup, config *configuration.Configuration) {
	biosEndpoint = config.ListEndpoint.BiosEndpoint
	biosPath = config.Fs.BiosPath
	r.GET(biosEndpoint, getListBios)

}

// getListBios -> get Bios list...
func getListBios(c *gin.Context) {
	listFiles := []Bios{}
	files, _ := ioutil.ReadDir(biosPath)

	for _, f := range files {
		filePath := biosPath + "/" + f.Name()
		if f.IsDir() {
		} else {
			listFiles = append(listFiles, Bios{Name: f.Name(), Hash: utils.MD5ToString(filePath)})
		}
	}
	c.JSON(http.StatusOK, listFiles)
}
