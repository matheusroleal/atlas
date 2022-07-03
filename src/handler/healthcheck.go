/*
 * @Author: Matheus Leal
 * @Date: 2022-07-01 22:53:39
 * @Last Modified by: Matheus Leal
 * @Last Modified time: 2022-07-03 14:12:04
 */
package handler

import (
	"net/http"

	"github.com/julienschmidt/httprouter"
)

/**
 * Healthcheck HTTP route handler.
 *
 * @param   w				http.ResponseWriter			The header map that will be sent by WriteHeader
 *					req			http.Request						Specifies the HTTP method
 * @return
 */
func HealthcheckHandler(w http.ResponseWriter, req *http.Request, _ httprouter.Params) {
	w.Header().Set("Content-Type", "text/html;charset=UTF-8")
	w.Header().Set("Vary", "Accept-Encoding, Origin")

	w.WriteHeader(http.StatusOK)

	w.Write([]byte("WORKING\n"))
}
