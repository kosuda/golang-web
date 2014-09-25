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
	router.GET("/api/redis/user/:id", controller.RedisUserGet)
	router.GET("/api/redis/user", controller.RedisUserGet)
	router.PUT("/api/user", controller.UserUpsert)
	router.PUT("/api/user/:id", controller.UserUpsert)
	router.PUT("/api/redis/user/:id", controller.RedisUserWrite)
	router.DELETE("/api/user", controller.UserDelete)
	router.DELETE("/api/user/:id", controller.UserDelete)
	router.PATCH("/api/user/:id", controller.UserUpdate)

	// TOTEC
	router.GET("/api/musics", controller.MusicGet)
	router.GET("/api/musics/:id", controller.MusicGet)
	router.POST("/api/musics", controller.MusicUpsert)
	router.PUT("/api/musics/:id", controller.MusicUpsert)
	router.DELETE("/api/musics/:id", controller.MusicDelete)

	router.POST("/api/musics/:id/play", controller.HistoryUpsert)
}
