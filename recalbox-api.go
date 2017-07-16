package main

import (
	"gitlab.com/MikaXII/recalbox-api/httprouter"
	"net/http"
	"log"
	"gitlab.com/MikaXII/recalbox-api/controllers"
)

func main() {

	router := httprouter.New()
	loadEndpoints(router)
	log.Fatal(http.ListenAndServe(":8080", router))
}


func loadEndpoints(r *httprouter.Router){
	apiV1 := r.NewGroup("/v1")
	recalroutes.SystemGroupV1(apiV1)
}







