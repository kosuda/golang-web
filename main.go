package main

import (
	"github.com/kosuda/golang-web/config"
	"github.com/kosuda/golang-web/router"
	"log"
	"net/http"
	"strconv"
)

func main() {
	c := config.Configuration.Common.API
	log.Printf("Listen Address %s:%d", c.Host, c.Port)
	log.Fatal(http.ListenAndServe(c.Host+":"+strconv.Itoa(c.Port), router.New()))
}
