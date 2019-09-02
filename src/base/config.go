package base

import (
	"os"
)

var defaultPort = "8111"

func GetPort() string {
	port := os.Getenv("port")

	if port == "" {
		port = defaultPort
	}

	return port
}
