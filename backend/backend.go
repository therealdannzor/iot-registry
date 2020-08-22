package backend

import (
	"errors"
	"sync"
	"time"
)

const NUMBER_OF_DEVICES = 20
const MAX_SERVED_REQUESTS = 10

// APIcall interface contains the only API call the registry in which it initially uses
type APIcall interface {
	// Register is the API function in which the system initiates a registeration process of the generated DevEUIs (in shortcode format).
	// The backend enqueues up to 10 in-flight requests to process at a time until it has finished them all.
	// Available HTTP status codes are:
	// 	- 200 (OK): all the (100) devices were successfully registered and returned as a JSON body
	// 	- 500 (error): the registration process was interrupted/cancelled or have failed in some way;
	//                 the service proceeds with handling the in-flight requests (up to 10) and then exits
	// 	- 503 (error): the service is currently operating at full capacity, the user should try again later
	Register(wholeID string) error
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

	time.Sleep(50 * time.Millisecond) // simulate system processing time to register a device
	sb.registry[shortCode] = wholeID  // device added to registry

	return nil
}

// Enqueue adds a short-format DevEUI to the in-flight queue
func (sb SensorBackend) Enqueue(shortcodeID string) error {
	sb.mu.Lock()
	defer sb.mu.Unlock()

	// throttle the request if operating at max capacity
	if sb.scLen == 10 {
		return errors.New("already processing maximum in-flight requests")
	}

	if shortcodeID == "" || len(shortcodeID) != 5 {
		return errors.New("incorrect format of device ID, must be of length 5")
	}

	sb.scheduled[sb.scLen] = shortcodeID
	sb.scLen++
	return nil
}

// Pop removes an ID from the queue and returns it
func (sb SensorBackend) Pop() (string, error) {
	sb.mu.Lock()
	defer sb.mu.Unlock()
	var emptystring string
	var ID string

	switch scLen := len(sb.scheduled); scLen {
	case 0:
		return emptystring, errors.New("nothing to register")
	case 1:
		var empty [10]string
		ID, sb.scheduled = sb.scheduled[0], empty
	default:
		ID, sb.scheduled = sb.scheduled[0], shiftLeft(sb.scheduled)
	}

	sb.scLen--
	return ID, nil
}

func shiftLeft(arr [MAX_SERVED_REQUESTS]string) [MAX_SERVED_REQUESTS]string {
	lastIndex := MAX_SERVED_REQUESTS - 1

	// move all elements one index left
	for i := 0; i < lastIndex; i++ {
		arr[i] = arr[i+1]
	}
	// reset last entry to avoid a duplicate of the last element
	arr[lastIndex] = ""

	return arr
}
