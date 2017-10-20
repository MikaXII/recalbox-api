package controllers

import (
	"fmt"
	"io"
	"io/ioutil"
	"log"
	"net/http"
	"os"

	"github.com/gin-gonic/gin"
	"gitlab.com/MikaXII/recalbox-api/config"
	"gitlab.com/MikaXII/recalbox-api/models"
)

var systemsEndpoint string
var systemsPath string
var gamelistPath string

// RomGroupV1 Regroup path for v1 endpoint
func RomGroupV1(r *gin.RouterGroup, config *configuration.Configuration) {
	systemsEndpoint = config.ListEndpoint.SystemsEndpoint
	systemsPath = config.Fs.SystemsPath
	r.GET(systemsEndpoint, getSystemList)
	r.GET(systemsEndpoint+"/:systemId/", getRomsBySytem)
	r.GET(systemsEndpoint+"/:systemId/hash", getRomsHashBySystem)
	r.GET(systemsEndpoint+"/:systemId/gamelist", getGamelist)

	r.POST(systemsEndpoint+"/:systemId", uploadRoms)
}

// getSystemList Get system list
func getSystemList(c *gin.Context) {
	listFiles := []models.System{}
	files, _ := ioutil.ReadDir(systemsPath)
	for _, f := range files {
		if f.IsDir() {
			listFiles = append(listFiles, models.System{Name: f.Name()})
		}
	}
	c.JSON(http.StatusOK, listFiles)
}

// getRomsBySytem Get all rom in a system folder
func getRomsBySytem(c *gin.Context) {
	listFiles := []models.System{}
	systemID := c.Param("systemId")
	files, _ := ioutil.ReadDir(systemsPath + systemID)
	for _, f := range files {
		listFiles = append(listFiles, models.System{Name: f.Name()})
	}
	c.JSON(http.StatusOK, listFiles)
}

// getRomsHashBySystem Get All rom's hash in a system folder
func getRomsHashBySystem(c *gin.Context) {
	systemID := c.Param("systemId")
	files, _ := ioutil.ReadDir(systemsPath + systemID)
	romInfo := []models.Media{}

	for _, f := range files {
		filePath := systemsPath + systemID + "/" + f.Name()
		rom := models.NewMedia(filePath, f)
		romInfo = append(romInfo, *rom)
	}
	c.JSON(http.StatusOK, romInfo)
}

// getGamelist Get the gamelist of a system
func getGamelist(c *gin.Context) {
	systemID := c.Param("systemId")
	gamelist, _ := ioutil.ReadFile(systemsPath + "/" + systemID + gamelistPath)
	c.JSON(http.StatusOK, gamelist)
}

// uploadRoms Uploading a rom with system name
func uploadRoms(c *gin.Context) {

	form, _ := c.MultipartForm()
	files := form.File["upload[]"]
	systemID := c.Param("systemId")

	for _, file := range files {
		src, _ := file.Open()
		defer func() {
			err := src.Close()
			if err != nil {
				log.Fatal(err)
			}
		}()
		println(systemsPath + systemID + "/" + file.Filename)
		dst, _ := os.Create(systemsPath + systemID + "/" + file.Filename)
		defer func() {
			err := dst.Close()
			if err != nil {
				log.Fatal(err)
			}
		}()

		_, err := io.Copy(dst, src)
		if err != nil {
			log.Fatal(err)
		}

	}
	c.String(http.StatusOK, fmt.Sprintf("%d files uploaded!", len(files)))
}
