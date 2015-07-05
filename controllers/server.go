package controllers

import (
	"github.com/codegangsta/negroni"
	"github.com/gorilla/mux"
	"net/http"
)

func StartREST(hostname string, staticdir string) {
	router := mux.NewRouter().StrictSlash(true)
	router.HandleFunc("/index", IndexHandler)
	router.PathPrefix("/").Handler(http.FileServer(http.Dir(staticdir)))

	n := negroni.New()
	n.UseHandler(router)

	n.Run(hostname)
}
