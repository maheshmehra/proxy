package main

import (
	"github.com/maheshmehra/proxy/connections"
	"github.com/maheshmehra/proxy/objects"
	"github.com/maheshmehra/proxy/utils"
	"github.com/rs/cors"
	"log"
	"net/http"
)

func main() {

	// loading config data
	status := utils.ConfigSetter("config.json")

	// checking status
	if !status {
		return
	}

	// handling cors
	c := cors.New(cors.Options{
		AllowedOrigins: []string{objects.ProxyConfigObj.ClientServer},
		AllowedMethods: []string{"GET", "POST", "PUT", "DELETE"},
	})

	// handling requests
	http.HandleFunc("/", connections.HandleConnections)

	err := http.ListenAndServe(":80", c.Handler(http.DefaultServeMux))

	if err != nil {
		log.Fatalln(err)
	}
}
