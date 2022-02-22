package main

import (
	"log"
	"net/http"

	"github.com/julienschmidt/httprouter"
	"github.com/yufy/go-example/swagger/handlers"
)

func main() {
	router := httprouter.New()

	router.POST("/v1/users", handlers.CreateUser)
	router.GET("/v1/users/:id", handlers.User)
	router.PUT("/v1/users/:id", handlers.UpdateUser)
	router.DELETE("/v1/users/:id", handlers.DeleteUser)
	router.GET("/v1/users", handlers.Users)

	log.Fatal(http.ListenAndServe(":13225", router))
}
