/*
 * @Author: Matheus Leal
 * @Date: 2022-07-01 22:53:53
 * @Last Modified by: Matheus Leal
 * @Last Modified time: 2022-07-03 14:12:29
 */
package handler

import (
	"encoding/json"
	"net/http"

	"github.com/julienschmidt/httprouter"
)

type Link struct {
	Description string `json:"description"`
	Method      []Link `json:"method"`
}

type Index struct {
	Description string `json:"description"`
	Href        string `json:"href"`
}

/**
 * Index HTTP route handler.
 *
 * @param   w				http.ResponseWriter			The header map that will be sent by WriteHeader
 *					req			http.Request						Specifies the HTTP method
 * @return
 */
func IndexHandler(w http.ResponseWriter, req *http.Request, _ httprouter.Params) {
	data := Link{}

	w.Header().Set("Content-Type", "application/json")
	w.Header().Set("Vary", "Accept-Encoding, Origin")

	w.WriteHeader(http.StatusOK)

	json.NewEncoder(w).Encode(data)
}
