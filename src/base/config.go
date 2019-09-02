package base

import (
	"os"
)

var defaultPort = "8111"
var defaultApiHost = "localhost:8111"
var defaultUrlScheme = "http"

func GetPort() string {
	returnval := os.Getenv("port")

	if returnval == "" {
		returnval = defaultPort
	}

	return returnval
}

func GetApiHost() string {
	returnval := os.Getenv("api_host")

	if returnval == "" {
		returnval = defaultApiHost
	}

	return returnval
}

func GetUrlScheme() string {
	returnval := os.Getenv("url_scheme")

	if returnval == "" {
		returnval = defaultApiHost
	}

	return returnval
}
