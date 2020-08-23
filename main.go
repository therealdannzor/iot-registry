package main

import (
	"fmt"

	"github.com/therealdannzor/iot-registry/api"
	"github.com/therealdannzor/iot-registry/backend"
)

func main() {
	back, err := backend.New()
	if err != nil {
		fmt.Println(err)
		return
	}

	api.Expose(back, "127.0.0.1:8080")
}
