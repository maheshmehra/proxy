package lib

import (
	"fmt"
	"github.com/ztrue/tracerr"
)

// HandlePanic : handle panic function, it recovers application from exception
func HandlePanic() {
	// if any error occurred handle
	if err := recover(); err != nil {
		text := tracerr.Sprint(err.(error))
		fmt.Println(text)
	}
}
