package api

import (
	"fmt"
	"net/http"

	"github.com/gorilla/mux"
	"github.com/rs/cors"
	ob "github.com/therealdannzor/iot-registry/api/handlers/onboard"
	be "github.com/therealdannzor/iot-registry/backend"
)

// Expose exposes the APIs to the backend `back` at ip:port `address`
func Expose(back be.APIcall, address string) {
	// add cors restrictions
	c := cors.New(cors.Options{
		AllowedOrigins: []string{"*"},
		AllowedMethods: []string{"POST"},
	})

	// create a router and register the URL path
	r := mux.NewRouter()
	initRoutes(r)

	// connect backend implementation to the API
	ob.SetWorker(back)

	fmt.Println("API exposed. Listening on " + address + " ...")
	// start server at given address
	err := http.ListenAndServe(address, c.Handler(r))
	if err != nil {
		fmt.Println("Could not start server: ", err)
		return
	}

}
