package main

import (
	"fmt"
	"log"
	"net/http"
	"os"
	"runtime"
)

// Variables
var err error
var ver int
var newline string
var domain string
var port string
var certPath string

func main() {
	// Variables
	ver := "0.1.0"
	newline := "\n"
	port := ":80"
	certPath := "/var/goacme/certs"
	// Main function
	log.Print("Starting...")
	// Debug Details
	fmt.Print("goacme v" + ver + " is running on " + runtime.GOOS + " machine" + newline)
	// Get the domain to get a certificate for
	fmt.Print("Enter the domain you would like to add: ")
	fmt.Scanln(&domain)
	// Ping the Let's Encrypt directory
	log.Print("Sending initial ping...")
	dirResp, err := http.Get("https://acme-v02.api.letsencrypt.org/directory")
	if err != nil {
		log.Fatal(err)
	}
	defer dirResp.Body.Close()
	log.Print("Successful." + newline)
	// Create HTTP server on port 80
	log.Print("Allowing incoming requests...")
	if err != nil {
		fmt.Print("There seems to be an error with making a request on port " + port + newline)
		fmt.Print("This is most likely because your firewall is blocking port " + port)
	}
	log.Print("Successful.")
	if _, err := os.Stat(certPath); os.IsNotExist(err) {
		os.Mkdir(certPath, os.ModeDir)
	}
}
