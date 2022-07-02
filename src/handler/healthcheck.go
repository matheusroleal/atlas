/*
 * @Author: Matheus Leal
 * @Date: 2022-07-01 22:53:39
 * @Last Modified by:   Matheus Leal
 * @Last Modified time: 2022-07-01 22:53:39
 */
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
