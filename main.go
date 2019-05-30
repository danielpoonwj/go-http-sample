package main

import (
	"fmt"
	"log"
	"net/http"
	"os"
	"strconv"

	"github.com/gorilla/mux"
)

const defaultPort int = 8080

func getPort() int {
	port, ok := os.LookupEnv("PORT")
	if ok {
		portInt, err := strconv.Atoi(port)
		if err != nil {
			log.Fatal(err)
		}

		return portInt
	}

	return defaultPort
}

func main() {
	router := mux.NewRouter()

	router.HandleFunc("/heartbeat", heartbeatHandler).Methods("GET")
	router.HandleFunc("/", echoHandler).Methods("GET")

	router.Use(loggingMiddleware)

	port := getPort()
	log.Printf("Starting server on port %d\n", port)
	http.ListenAndServe(fmt.Sprintf(":%d", port), router)
}
