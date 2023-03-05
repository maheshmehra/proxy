package utils

import (
	"encoding/json"
	"github.com/maheshmehra/proxy/objects"
	"io/ioutil"
	"log"
	"os"
)

// ConfigSetter : function for read data from config and set to struct
func ConfigSetter(fileName string) bool {
	// reading file
	file, err := os.Open(fileName)

	if err != nil {
		log.Fatalln("Unable to read config data ", err.Error())
		return false
	}

	// reading byte from file
	byteValue, err := ioutil.ReadAll(file)

	if err != nil {
		log.Fatalln("Error while reading from file ", err.Error())
		return false
	}

	err = json.Unmarshal(byteValue, &objects.ProxyConfigObj)

	if err != nil {
		log.Fatalln("Unable to decode file data in struct ", err.Error())
		return false
	}

	// closing file
	err = file.Close()

	if err != nil {
		log.Fatalln("Error while closing file ", err.Error())
		return false
	}

	return true
}
