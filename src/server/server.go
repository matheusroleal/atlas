package server

import (
	"fmt"
	"log"
	"net/http"
	"os"

	"github.com/julienschmidt/httprouter"

	"github.com/matheusroleal/atlas/src/handler"
)

func RunServer() {
	router := httprouter.New()

	// API setup Routes
	router.GET("/", handler.IndexHandler)
	router.GET("/healthcheck", handler.HealthcheckHandler)

	// Track Routes
	router.POST("/track", handler.SegmentCreate)

	addr := fmt.Sprintf("%s:%s", os.Getenv("HOST"), os.Getenv("PORT"))
	log.Println()
	log.Printf("[ATLAS API] listening on %s", addr)
	log.Println()

	log.Fatal(http.ListenAndServe(addr, router))
}
