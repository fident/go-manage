package main

import (
	"fmt"
	"log"

	"github.com/fident/go-manage/authenticate"
)

const (
	fidentInstanceAddress = "localhost:50052"
	keyfilePath           = "./testkey.json"
)

func main() {
	token, err := authenticate.GetToken(keyfilePath, fidentInstanceAddress)
	if err != nil {
		log.Fatalf("failed to get token: %v", err)
	}

	fmt.Printf("Token: %s\n", token)
}
