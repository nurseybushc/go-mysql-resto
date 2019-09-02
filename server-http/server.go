package main

import (
	"net/http"

	"github.com/boantp/go-mysql-resto/api"
	"github.com/julienschmidt/httprouter"
)

func main() {
	router := httprouter.New()

	//Get API Controller instance
	api := api.NewApiController()

	//POST nearby restaurant
	router.POST("/restaurant/create", api.GetNearbyRestaurant)
	router.GET("/restaurants", api.GetAllRestaurants)
	//GET detail restaurant
	router.GET("/restaurant/:restaurant_id", api.GetDetailRestaurant)
	//POST reservation restaurant
	router.POST("/restaurant/reservation/create", api.CreateReservationRestaurant)
	http.ListenAndServe(":3000", router)
}
