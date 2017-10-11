package main

import (
	"flag"
	"io"
	"log"
	"net/http"
	"os"
	"time"

	"github.com/aws/aws-sdk-go/aws/credentials"
	"github.com/aws/aws-sdk-go/aws/signer/v4"
	"github.com/k0kubun/pp"
)

func main() {
	// Flags
	url := flag.String("u", "", "endpoint URL of API Gateway")
	credPath := flag.String("c", "", "shared credentials file path (default is ~/.aws/credentials)")
	profile := flag.String("p", "default", "profile name of credentials")
	region := flag.String("r", "us-east-1", "API region")
	method := flag.String("m", "GET", "HTTP method")
	verbose := flag.Bool("v", false, "verbose output")
	host := flag.String("host", "", "Host header (default is extracted from URL)")
	flag.Parse()

	// Create a new HTTP request
	req, err := http.NewRequest(*method, *url, nil)
	if err != nil {
		log.Fatalf("failed to create a new HTTP request: %v", err)
	}

	if *host != "" {
		req.URL.Host = *host
	}

	// Sign the request with the given credentials
	cred := credentials.NewSharedCredentials(*credPath, *profile)
	signer := v4.NewSigner(cred)
	if _, err := signer.Sign(req, nil, "execute-api", *region, time.Now()); err != nil {
		log.Fatalf("failed to sign a HTTP request: %v", err)
	}

	// Add no cache header
	req.Header.Add("Cache-Control", "no-cache")

	// Print Authorization header for debugging
	if *verbose {
		pp.Println("Request header:", req.Header)
	}

	// Send the HTTP request
	client := new(http.Client)
	resp, err := client.Do(req)
	if err != nil {
		log.Fatalf("failed to fetch response: %v", err)
	}

	defer resp.Body.Close()

	if *verbose {
		pp.Println("Response header:", resp.Header)
	}

	// Print a response to stdout
	io.Copy(os.Stdout, resp.Body)
}
