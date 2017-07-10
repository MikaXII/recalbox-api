package main

import (
	"net/http"
	"time"
)

func main() {

	mux := http.NewServeMux();

	mux.Handle("/api/", apiHandler{})

	mux.Handle("/", Home(time.RFC1123))


	http.ListenAndServe(":8080", mux);
}

func Home(format string) http.Handler {
	fn := func(w http.ResponseWriter, r *http.Request) {
		println(r.Method)

		tm := time.Now().Format(format)
		w.Write([]byte("The time is: " + tm ))
	}
	return http.HandlerFunc(fn)
}

