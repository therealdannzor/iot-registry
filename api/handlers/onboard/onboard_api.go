package api

import (
	"github.com/therealdannzor/iot-registry/backend"
)

// OnboardAPI is API implementation of /onboard root endpoint
type OnboardAPI struct {
	Backend backend.SensorBackend
}
