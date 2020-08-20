// DO NOT EDIT THIS FILE. This file will be overwritten when re-running go-raml.
package api

import (
	"github.com/gorilla/mux"

	v1 "github.com/therealdannzor/iot-registry/api/handlers/onboard"
)

func initRoutes(r *mux.Router) {

	OnboardInterfaceRoutes(r, v1.OnboardAPI{})
}
