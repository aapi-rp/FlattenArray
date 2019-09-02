package base

import (
	"os"
	"time"
)

//ReadTimeout server connection read timeout
var ReadTimeout = 650 * time.Second

//WriteTimeout server connection write timeout
var WriteTimeout = 650 * time.Second

//IdleTimeout server connection idle timeout
var IdleTimeout = 670 * time.Second

var defaultPort = "8234"
var defaultApiHost = "localhost:8234"
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
		returnval = defaultUrlScheme
	}

	return returnval
}
