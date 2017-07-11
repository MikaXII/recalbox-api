package models

import (
	"net/http"
	"io/ioutil"
	"encoding/json"
	"strings"
)

type systemName struct {
	Name string
}

func ApiSystem(path string, mux *http.ServeMux) {

	mux.Handle(path, getSystemList())


	mux.Handle(path + "/", getRomsBySytem())
}


func getSystemList() http.Handler{
	fn := func(w http.ResponseWriter, r *http.Request) {


		switch r.Method {
		case "GET":
			listFiles := []systemName{}
			files, _ := ioutil.ReadDir("/");
			for _, f := range files {
				listFiles = append(listFiles, systemName{Name:f.Name()})
			}

			jsonFiles, _ := json.Marshal(listFiles)

			w.Write([]byte(jsonFiles))
			break;

		/*case "POST":
			w.Write([]byte("{toto: prout}"))
			break;*/
		}

	}
	return http.HandlerFunc(fn)
}

func getRomsBySytem() http.Handler {
	fn := func(w http.ResponseWriter, r *http.Request) {
		switch r.Method {
		case "GET":

			listFiles := []systemName{}
			files, _ := ioutil.ReadDir("/");
			for _, f := range files {
				listFiles = append(listFiles, systemName{Name:f.Name()})
			}

			jsonFiles, _ := json.Marshal(listFiles)

			w.Write([]byte(jsonFiles))
			break;

			/*case "POST":
				w.Write([]byte("{toto: prout}"))
				break;*/
		}

	}
	return http.HandlerFunc(fn)
}