package recalroutes

import (
	"net/http"
	"io/ioutil"
	"encoding/json"
	"gitlab.com/MikaXII/recalbox-api/httprouter"
	"gitlab.com/MikaXII/recalbox-api/utils"
)

const ROM_DIR = "/recalbox/share/roms/"
const BASE_PATH = "/systems"

type systemName struct {
	Name string
}

type RomHash struct {
	Name string
	Crc string
	Md5 string
	Sha1 string
}

type RomHashes []RomHash

func SystemGroupV1(r *httprouter.RouteGroup) {
	r.GET(BASE_PATH, getSystemList)
	r.GET(BASE_PATH +"/:id/roms", getRomsBySytem)
	r.GET(BASE_PATH +"/:id/hash", getRomsHash)
}

func response(w http.ResponseWriter, data []byte) {
	w.Header().Set("Content-Type", "application/json")
	w.Write(data)
}

func getSystemList(w http.ResponseWriter, r *http.Request, ps httprouter.Params) {

			listFiles := []systemName{}
			files, _ := ioutil.ReadDir(ROM_DIR);
			for _, f := range files {
				listFiles = append(listFiles, systemName{Name:f.Name()})
			}

			jsonFiles, _ := json.Marshal(listFiles)

			response(w,[]byte(jsonFiles))
}

func getRomsBySytem(w http.ResponseWriter, r *http.Request, ps httprouter.Params) {

	listFiles := []systemName{}
	systemId := ps.ByName("id")
	files, _ := ioutil.ReadDir(ROM_DIR + systemId );
	for _, f := range files {
		listFiles = append(listFiles, systemName{Name: f.Name()})
	}
	jsonFiles, _ := json.Marshal(listFiles)
	response(w,[]byte(jsonFiles))
}

func getRomsHash(w http.ResponseWriter, r *http.Request, ps httprouter.Params) {

	systemId := ps.ByName("id")
	files, _ := ioutil.ReadDir(ROM_DIR + systemId );
	romH := RomHashes{}

	for _, f := range files {
		filePath := ROM_DIR + systemId +"/" + f.Name()
		rom := RomHash{Name: f.Name(),
			Crc: utils.CRC32ToString(filePath),
			Md5: utils.MD5ToString(filePath),
			Sha1: utils.SHA1ToString(filePath)}
		romH = append(romH, rom)
	}
	jsonFiles, _ := json.Marshal(romH)
	response(w,[]byte(jsonFiles))
}

