package main

import (
	"bytes"
	"encoding/json"
	"fmt"
	"log"
	"net/http"
	"os"
	"time"
)

// Variables
var err error
var ver int
var newline string
var domain string
var port string
var certPath string
var TC string
var colorReset string
var colorRed string
var email string

func main() {
	// Variables
	ver := "0.1.0"
	newline := "\n"
	certPath := "/var/goacme/certs"
	TC := "T"
	log.Print("Starting...")
	time.Sleep(1500)
	// Debug Details
	fmt.Print("goacme v" + ver + newline)
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
	log.Print("You must read and agree to the Let's Encrypt Terms and Conditions, listed here (https://letsencrypt.org/documents/LE-SA-v1.2-November-15-2017.pdf)" + newline)
	fmt.Print("Do you agree? [Y/N] ")
	fmt.Scan(&TC)
	// Ironically, putting the if Terms and Conditions are "True"/"Y" will close if the TC != "N" rather then "Y"
	// In other words,
	if TC != "Y" { // Don't change "Y" to "N" or else the user will only be able to continue if they say no.
		os.Exit(3)
	}
	log.Println("Creating Certificate Path...")
	if _, err := os.Stat(certPath); os.IsNotExist(err) {
		os.Mkdir(certPath, os.ModeDir)
	}
	log.Print("Successful." + newline)
	fmt.Println("----------------------------------------------" + newline)
	log.Print("Please add an e-mail address (used for certificate warnings): ")
	fmt.Scan(&email)
	log.Print("Adding account..." + newline)
	acctData, err := json.Marshal(map[string]string{
		"email": email,
	})
	newAcct, err := http.Post("https://acme-v02.api.letsencrypt.org/acme/new-acct", "application/json", bytes.NewBuffer(acctData))
	if err != nil {
		log.Fatal(err)
	}
	defer newAcct.Body.Close()
	fmt.Println(newAcct)
}
