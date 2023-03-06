package connections

import (
	"fmt"
	"github.com/maheshmehra/proxy/lib"
	"github.com/maheshmehra/proxy/objects"
	"io"
	"net/http"
	"strings"
)

// HandleConnections : function is used for handle requests
func HandleConnections(writer http.ResponseWriter, req *http.Request) {
	// handling exceptions
	defer lib.HandlePanic()

	// Base URL
	baseURL := ""

	// checking for exchange
	if strings.Contains(req.RequestURI, "?exchange=binance") {
		baseURL = objects.ProxyConfigObj.BinanceBaseURL
		req.RequestURI = strings.Replace(req.RequestURI, "?exchange=binance", "", 1)
	} else if strings.Contains(req.RequestURI, "?exchange=kucoin") {
		baseURL = objects.ProxyConfigObj.KuCoinBaseURL
		req.RequestURI = strings.Replace(req.RequestURI, "?exchange=kucoin", "", 1)
	} else if strings.Contains(req.RequestURI, "?exchange=gateio") {
		baseURL = objects.ProxyConfigObj.GateIOBaseURL
		req.RequestURI = strings.Replace(req.RequestURI, "?exchange=gateio", "", 1)
	}

	// hitting exchange api
	res, err := http.Get(baseURL + req.RequestURI)

	if err != nil {
		fmt.Println(err)
	}

	if err == nil {
		// reading response
		byteVal, _ := io.ReadAll(res.Body)

		// writing response in writer object
		_, err = fmt.Fprintf(writer, string(byteVal))

		// checking if any error occurred
		if err != nil {
			fmt.Println(err)
		}
	}
}
