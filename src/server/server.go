/*
 * @Author: Matheus Leal
 * @Date: 2022-07-01 22:54:13
 * @Last Modified by:   Matheus Leal
 * @Last Modified time: 2022-07-01 22:54:13
 */
package server

import (
	"fmt"
	"net/http"
	"os"

	"github.com/julienschmidt/httprouter"
	"github.com/matheusroleal/atlas/src/handler"
	log "github.com/sirupsen/logrus"
)

func RunServer() {
	router := httprouter.New()

	// API setup Routes
	router.GET("/", handler.IndexHandler)
	router.GET("/healthcheck", handler.HealthcheckHandler)

	// Track Routes
	router.POST("/track", handler.TrackCreate)

	// Segments Routes
	router.POST("/segment", handler.SegmentCreate)

	addr := fmt.Sprintf("%s:%s", os.Getenv("HOST"), os.Getenv("PORT"))
	log.Info()
	log.Info("[ATLAS API] listening on %s", addr)
	log.Info()

	log.Error(http.ListenAndServe(addr, router))
}
