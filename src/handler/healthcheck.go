package handler

import (
	"net/http"

	"github.com/julienschmidt/httprouter"
)

func HealthcheckHandler(w http.ResponseWriter, req *http.Request, _ httprouter.Params) {
	w.Header().Set("Content-Type", "text/html;charset=UTF-8")
	w.Header().Set("Vary", "Accept-Encoding, Origin")

	w.WriteHeader(http.StatusOK)

	w.Write([]byte("WORKING\n"))
}
