package backend

import (
	"crypto/rand"
	"encoding/hex"
	"errors"
	"fmt"
	"sync"
)

// New creates a backend for the sensor registry service
func New() (SensorBackend, error) {
	reg := make(map[string]string)
	b := SensorBackend{registry: reg}
	list, err := newBatch()
	if err != nil {
		return b, err
	}

	b.SensorList = list
	b.mu = new(sync.Mutex)

	return b, nil
}

// newBatch generates a full sequence of DevEUIs (UUIDs). Each DevEUI is a 16-character hex string.
func newBatch() ([NUMBER_OF_DEVICES]string, error) {
	tmp := make(map[string]string)
	var res [NUMBER_OF_DEVICES]string

	for i := 0; i < NUMBER_OF_DEVICES; {
		ID, err := randHex()
		if err != nil {
			return res, err
		}
		short, err := shortFmt(ID)
		if err != nil {
			return res, err
		}

		// make sure we add unique short code IDs
		if tmp[short] == "" {
			tmp[short] = ID
			i++
		}
	}

	var index int
	// the key is short code and value is full code -> add the full code only
	for _, value := range tmp {
		res[index] = value
		index++
	}

	return res, nil
}

// randHex creates a 16-character hex string
func randHex() (string, error) {
	b := make([]byte, 8) // one byte -> two hex characters
	if _, err := rand.Read(b); err != nil {
		return "", err
	}

	return hex.EncodeToString(b), nil
}

// shortFmt accepts a 16 character hex string and returns the last 5 characters
func shortFmt(hexstr string) (string, error) {
	if len(hexstr) != 16 {
		fmt.Println(hexstr)
		return "", errors.New("incorrect hex string length, must be 16")
	}

	return hexstr[len(hexstr)-5:], nil
}
