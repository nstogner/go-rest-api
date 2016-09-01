package main

import (
	"net/http"
	"os"

	"github.com/Sirupsen/logrus"
	"github.com/nstogner/go-rest-api/part2/rest"
)

func main() {
	port, ok := os.LookupEnv("PORT")
	if !ok {
		// Default to port 8080
		port = "8080"
	}

	logrus.WithField("port", port).Info("listening for http traffic")
	// If ListenAndServe returns an error we will log and exit
	logrus.Fatal(http.ListenAndServe(":"+port, rest.API()))
}
