/*
 * @Author: Matheus Leal
 * @Date: 2022-07-01 22:53:53
 * @Last Modified by:   Matheus Leal
 * @Last Modified time: 2022-07-01 22:53:53
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

func IndexHandler(w http.ResponseWriter, req *http.Request, _ httprouter.Params) {
	data := Link{}

	w.Header().Set("Content-Type", "application/json")
	w.Header().Set("Vary", "Accept-Encoding, Origin")

	w.WriteHeader(http.StatusOK)

	json.NewEncoder(w).Encode(data)
}
