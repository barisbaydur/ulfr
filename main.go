package main

import (
	"log"
	"net/http"
	"ulfr/config"
	"ulfr/controllers"

	"github.com/julienschmidt/httprouter"
	"github.com/rs/cors"
)

func main() {
	controllers.Init()

	c := cors.New(cors.Options{
		AllowedOrigins: []string{"*"},
		AllowedMethods: []string{"POST", "GET", "OPTIONS"},
		AllowedHeaders: []string{"Content-Type"},
	})

	log.Println("Server started on localhost:8080")

	if err := http.ListenAndServe(":"+config.Port, c.Handler(Routes())); err != nil {
		log.Fatal(err)
	}
}

func Routes() *httprouter.Router {
	router := httprouter.New()

	// Fire Routes
	router.NotFound = http.HandlerFunc(controllers.FireULFR)
	router.POST("/fire", controllers.Fire{}.Trigger)
	router.GET("/fire", controllers.Fire{}.Index)
	router.GET("/fire/show/:id", controllers.Fire{}.Get)
	router.GET("/fire/delete/:id", controllers.Fire{}.Delete)

	// Dashboard Routes
	router.GET("/dashboard", controllers.Dashboard{}.Index)

	// Domain Routes
	router.GET("/domain", controllers.Domain{}.Index)
	router.GET("/domain/create", controllers.Domain{}.Add)
	router.POST("/domain/create", controllers.Domain{}.Create)
	router.GET("/domain/delete/:id", controllers.Domain{}.Delete)
	router.GET("/domain/update/:id", controllers.Domain{}.Get)
	router.POST("/domain/update/:id", controllers.Domain{}.Update)

	// Path Routes
	router.GET("/path", controllers.Path{}.Index)
	router.GET("/path/create", controllers.Path{}.Add)
	router.POST("/path/create", controllers.Path{}.Create)
	router.GET("/path/delete/:id", controllers.Path{}.Delete)
	router.GET("/path/update/:id", controllers.Path{}.Get)
	router.POST("/path/update/:id", controllers.Path{}.Update)

	// Static Routes
	router.ServeFiles("/assets/*filepath", http.Dir("assets/"))
	router.ServeFiles("/data/*filepath", http.Dir("data/"))

	return router
}
