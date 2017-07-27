package main

import (
	"fmt"
	"log"

	"github.com/fident/go-manage/authenticate"
	"github.com/fident/go-manage/key"
	"github.com/fident/go-manage/token"
)

const (
	fidentInstanceAddress = "localhost:50052"
	keyfilePath           = "./testkey.json"
)

func main() {
	// Read key file
	key, err := key.FromFile(keyfilePath)
	if err != nil {
		log.Fatalf("failed to read keyfile: %v", err)
	}

	// Use key to get token
	tok, err := authenticate.GetToken(key, fidentInstanceAddress)
	if err != nil {
		log.Fatalf("failed to get token: %v", err)
	}

	// Parse token for local management
	res, err := token.Parse(tok)
	if err != nil {
		log.Fatalf("failed to parse token: %v", err)
	}

	fmt.Printf("Token: %v\n", res)
	// Make calls using token (Auto-handling token expiry)

	fmt.Printf("Needs refresh: %v\n", res.WithinExpiryRange())
}
