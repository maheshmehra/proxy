package connectivity

import (
	"fmt"
	"github.com/maheshmehra/proxy/connections"
	"github.com/maheshmehra/proxy/lib"
	"github.com/maheshmehra/proxy/objects"
	"github.com/rs/cors"
	"net/http"
)

// HostHttp : function for open port
func HostHttp() bool {
	// handling error
	defer lib.HandlePanic()

	// variable for error status
	isConnected := true

	// handling cors
	c := cors.New(cors.Options{
		AllowedOrigins: []string{objects.ProxyConfigObj.ClientServer},
		AllowedMethods: []string{"GET", "POST", "PUT", "DELETE"},
	})

	// handling requests
	http.HandleFunc("/", connections.HandleConnections)

	go func() {
		// opening connection on port 8080
		err := http.ListenAndServe(":8080", c.Handler(http.DefaultServeMux))

		if err != nil {
			fmt.Println(err)
			isConnected = false
		}

		fmt.Println("Profiler started in port :8080")
	}()

	// returning final response
	return isConnected
}
