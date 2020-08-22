package api

import be "github.com/therealdannzor/iot-registry/backend"

// Global worker (exported)
var APIworker be.APIcall

func SetWorker(worker be.APIcall) {
	APIworker = worker
}
