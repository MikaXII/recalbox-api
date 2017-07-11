package main

import (
	"net/http"
	"gitlab.com/MikaXII/recalbox-api/router"
)

func main() {
	mux := http.NewServeMux();
	//mux.Handle("/api/", apiHandler{})

	router.LoadAllEndpoint(mux)

	http.ListenAndServe(":8080", mux);
}







