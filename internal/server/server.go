package server

import "fmt"

// Start listening and serving requests on given port
func Start(port string) {
	router := setRouter()

	router.Run(fmt.Sprintf(":%s", port))
}
