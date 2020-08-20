// THIS FILE IS SAFE TO EDIT. It will not be overwritten when rerunning go-raml.
package api

import (
	"encoding/json"
	"net/http"

	"github.com/therealdannzor/iot-registry/api/types"
)

// Post is the handler for POST /onboard
func (api OnboardAPI) Post(w http.ResponseWriter, r *http.Request) {
	var reqBody types.Request

	// decode request
	if err := json.NewDecoder(r.Body).Decode(&reqBody); err != nil {
		w.WriteHeader(400)
		return
	}

	w.Header().Set("Content-Type", "application/json")
	var respBody string
	json.NewEncoder(w).Encode(&respBody)
}
