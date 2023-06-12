package controllers

import (
	"log"
	"net/http"
)

func (server *Service) Run(addr string) {
	log.Fatal(http.ListenAndServe(addr, server.Router))
}
