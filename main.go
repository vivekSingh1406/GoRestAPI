package main

import (
	"fmt"
	"go-project/routers"
)

func main() {
	// Setup the router
	r := routers.SetupRouter()

	// port 5000
	if err := r.Run(":5000"); err != nil {
		fmt.Printf("Error starting server: %v\n", err)
	}
}
