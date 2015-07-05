package controllers

import (
	"encoding/json"
	"github.com/shymega/shelves/models"
	"net/http"
)

func IndexHandler(rw http.ResponseWriter, r *http.Request) {
	iha := models.IndexHandlerResponse{}
	iha.Doc = "http://shelves.readthedocs.org/api/v1"
	iha.Output = ""
	iha.StatusCode = 112

	js, err := json.Marshal(iha)
	if err != nil {
		panic(err)
	}

	rw.Header().Set("Content-Type", "application/json")
	rw.Write(js)
}
