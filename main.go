package main

import (
	"fmt"
	"log"

	"github.com/fident/go-manage/authenticate"
	"github.com/fident/go-manage/token"
)

const (
	fidentInstanceAddress = "localhost:50052"
	keyfilePath           = "./testkey.json"
)

func main() {
	tok, err := authenticate.GetToken(keyfilePath, fidentInstanceAddress)
	if err != nil {
		log.Fatalf("failed to get token: %v", err)
	}

	res, err := token.Parse(tok)
	if err != nil {
		log.Fatalf("failed to parse token: %v", err)
	}

	fmt.Printf("Token: %v\n", res)

	fmt.Printf("Needs refresh: %v\n", res.WithinExpiryRange())
}
