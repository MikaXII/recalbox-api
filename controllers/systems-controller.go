package controllers

import (
	"fmt"
	"io"
	"io/ioutil"
	"net/http"
	"os"

	"github.com/gin-gonic/gin"
	"gitlab.com/MikaXII/recalbox-api/config"
	"gitlab.com/MikaXII/recalbox-api/models"
)

var systemsEndpoint string
var systemsPath string
var gamelistPath string

func RomGroupV1(r *gin.RouterGroup, config *configuration.Configuration) {
	systemsEndpoint = config.ListEndpoint.SystemsEndpoint
	systemsPath = config.Fs.SystemsPath

	r.GET(systemsEndpoint, getSystemList)
	r.GET(systemsEndpoint+"/:systemId/", getRomsBySytem)
	r.GET(systemsEndpoint+"/:systemId/hash", getRomsHashBySystem)
	r.GET(systemsEndpoint+"/:systemId/gamelist", getGamelist)

	r.POST(systemsEndpoint+"/:systemId", uploadRoms)
}

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

func getRomsBySytem(c *gin.Context) {
	listFiles := []models.System{}
	systemID := c.Param("systemId")
	files, _ := ioutil.ReadDir(systemsPath + systemID)
	for _, f := range files {
		listFiles = append(listFiles, models.System{Name: f.Name()})
	}
	c.JSON(http.StatusOK, listFiles)
}

func getRomsHashBySystem(c *gin.Context) {
	supportedHash := models.SupportedHash()
	systemID := c.Param("systemId")
	files, _ := ioutil.ReadDir(systemsPath + systemID)
	romInfo := []models.RomInfo{}

	for _, f := range files {
		filePath := systemsPath + systemID + "/" + f.Name()
		rom := models.NewRom(filePath, f, supportedHash)
		romInfo = append(romInfo, *rom.Info())
	}
	c.JSON(http.StatusOK, romInfo)
}

func getGamelist(c *gin.Context) {
	systemID := c.Param("systemId")
	gamelist, _ := ioutil.ReadFile(systemsPath + "/" + systemID + gamelistPath)
	c.JSON(http.StatusOK, gamelist)
}

func uploadRoms(c *gin.Context) {

	form, _ := c.MultipartForm()
	files := form.File["upload[]"]
	systemID := c.Param("systemId")

	for _, file := range files {
		src, _ := file.Open()
		defer src.Close()
		println(systemsPath + systemID + "/" + file.Filename)
		dst, _ := os.Create(systemsPath + systemID + "/" + file.Filename)
		defer dst.Close()

		io.Copy(dst, src)
	}
	c.String(http.StatusOK, fmt.Sprintf("%d files uploaded!", len(files)))
}
