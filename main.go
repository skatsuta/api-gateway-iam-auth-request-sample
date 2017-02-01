package main

import (
	"flag"
	"fmt"
	"io"
	"log"
	"net/http"
	"os"
	"time"

	"github.com/aws/aws-sdk-go/aws/credentials"
	"github.com/aws/aws-sdk-go/aws/signer/v4"
)

func main() {
	// Flags
	url := flag.String("u", "", "endpoint URL of API Gateway")
	credPath := flag.String("c", "", "shared credentials file path")
	profile := flag.String("p", "default", "profile name of credentials")
	region := flag.String("r", "us-east-1", "API region")
	flag.Parse()

	// Create a new HTTP request
	req, err := http.NewRequest("GET", *url, nil)
	if err != nil {
		log.Fatalf("failed to create a new HTTP request: %v", err)
	}

	// Sign the request with the given credentials
	cred := credentials.NewSharedCredentials(*credPath, *profile)
	signer := v4.NewSigner(cred)
	if _, err := signer.Sign(req, nil, "execute-api", *region, time.Now()); err != nil {
		log.Fatalf("failed to sign a HTTP request: %v", err)
	}

	// Print Authorization header for debugging
	log.Println("Authorization:", req.Header.Get("Authorization"))

	// Send the HTTP request
	client := new(http.Client)
	resp, err := client.Do(req)
	if err != nil {
		log.Fatalf("failed to fetch response: %v", err)
	}

	defer resp.Body.Close()

	// Print a response to stdout
	fmt.Println("\nResponse:")
	io.Copy(os.Stdout, resp.Body)
}
