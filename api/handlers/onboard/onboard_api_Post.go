package api

import (
	"encoding/json"
	"errors"
	"fmt"
	"net/http"
	"os"
	"os/signal"

	"github.com/therealdannzor/iot-registry/backend"
)

// Result contains the response to the client after registration is fully, or partially, complete
type Result struct {
	DeviceIDs []string `json:"deveuis"`
}

// done is used to signal that the whole registration process is completed.
// Since each HTTP request in Go is its own goroutine and this is a globally declared
// variable, the response will propagate to all active consumers of this service.
var done = make(chan bool)

// sig is used to signal that the system is interrupted before it can finish.
// By using the same principle as the `done` channel above, all clients will know if
// the service backend is stops before it manages to finish the regiration of all IDs.
var sig = make(chan os.Signal, 1)

// Post is the handler for POST /onboard
func (api OnboardAPI) Post(w http.ResponseWriter, r *http.Request) {
	// reqCh is a channel tied to each API request goroutine
	var reqCh = make(chan string, 20)

	if APIworker == nil {
		errorResponse(w, http.StatusInternalServerError, errors.New("missing backend"))
		return
	}

	var IDs []string
	res := Result{DeviceIDs: IDs}

	signal.Notify(sig, os.Interrupt)
	go func() {
		defer os.Exit(1)
		<-sig
		fmt.Println("\nAbort process detected: finish in-flight requests before shutdown")
		fmt.Println("DevEUIs: ", res.DeviceIDs)
		// TODO: send HTTP response with registered DevEUIs
	}()

	go func() {
		for {
			req, stream := <-reqCh
			if stream {
				fmt.Println("Register: ", req)
				res.DeviceIDs = append(res.DeviceIDs, req)
				APIworker.Register(req)
			} else {
				fmt.Println("End of stream")
				done <- true
				break
			}
		}
	}()

	allDevices := APIworker.GetList()
	for i := 0; i < backend.NUMBER_OF_DEVICES; i++ {
		oneID := allDevices[i]
		reqCh <- oneID
	}

	close(reqCh)
	<-done

	sendResponse(w, res)
}

func sendResponse(w http.ResponseWriter, res Result) {
	jsonResp, err := json.Marshal(res)
	if err != nil {
		errorResponse(w, http.StatusInternalServerError, err)
		return
	}
	response(w, http.StatusOK, jsonResp)
}
