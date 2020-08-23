package backend

import (
	"errors"
	"sync"
	"time"
)

const NUMBER_OF_DEVICES = 100
const MAX_SERVED_REQUESTS = 10

// APIcall interface contains the only API call the registry in which it initially uses
type APIcall interface {
	// Register is the API function in which the system initiates a registeration process of the generated DevEUIs (in shortcode format).
	// The backend enqueues up to 10 in-flight requests to process at a time until it has finished them all.
	// Available HTTP status codes are:
	// 	- 200 (OK): all the (100) devices were successfully registered and returned as a JSON body
	// 	- 500 (error): the registration process was interrupted/cancelled or have failed in some way
	Register(wholeID string) error

	// GetList returns the SensorList
	GetList() [NUMBER_OF_DEVICES]string
}

// SensorBackend is the backend of the registry
type SensorBackend struct {
	// SensorList contains an array of unique DevEUIs that already is or are to be registered.
	// This is filled with 100x randomly generated 16-char hex strings at system initialisation.
	SensorList [NUMBER_OF_DEVICES]string

	// registry contains a lookup from short code formatted DevUIs (of length 5) to full the DevEUI (of length 16).
	// It is assumed that entries of this mapping has been registered by the API
	registry map[string]string

	// scheduled is a queue which contains the up to 10 unique DevEUIs that are considered in-flight to register
	scheduled [MAX_SERVED_REQUESTS]string

	// scLen keeps track of the amount of scheduled items and functions like a tail index for the queue
	scLen int

	// prevent race conditions
	mu *sync.Mutex
}

func (sb SensorBackend) Register(wholeID string) error {
	shortCode, err := shortFmt(wholeID)
	if err != nil {
		return err
	}

	if sb.registry[shortCode] != "" {
		return errors.New("device already registered")
	}

	time.Sleep(75 * time.Millisecond) // simulate system processing time to register a device
	sb.registry[shortCode] = wholeID  // device added to registry

	return nil
}

func (sb SensorBackend) GetList() [NUMBER_OF_DEVICES]string {
	return sb.SensorList
}
