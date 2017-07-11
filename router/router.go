package router

import (
	"net/http"
	"gitlab.com/MikaXII/recalbox-api/models"
)


func LoadAllEndpoint(mux *http.ServeMux) {

	models.ApiSystem("/systems", mux)
}

