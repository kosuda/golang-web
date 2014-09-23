package router

import (
	"github.com/julienschmidt/httprouter"
	"github.com/kosuda/golang-web/controller"
	"log"
)

// New Instanse function
func New() *httprouter.Router {
	router := httprouter.New()
	setup(router)
	log.Print("router setup done")
	return router
}

func setup(router *httprouter.Router) {
	router.GET("/api/user", controller.UserGet)
	router.GET("/api/user/:id", controller.UserGet)
	router.PUT("/api/user", controller.UserUpsert)
	router.PUT("/api/user/:id", controller.UserUpsert)
	router.DELETE("/api/user", controller.UserDelete)
	router.DELETE("/api/user/:id", controller.UserDelete)
	router.PATCH("/api/user/:id", controller.UserUpdate)
}
