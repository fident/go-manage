package main

import (
	"fmt"

	client "github.com/fident/go-manage"
)

const (
	fidentInstanceAddress = "localhost:50052"
	keyfilePath           = "./testkey.json"
)

func main() {
	testClient, err := client.New(keyfilePath, fidentInstanceAddress)
	if err != nil {
		panic(err)
	}

	lastlogin, err := testClient.GetAccountDetailsForIdentityID("EFIDFIID-ZGVT5I6L4-MISCR-V5UX35S")
	if err != nil {
		panic(err)
	}

	fmt.Printf("Got %v\n", lastlogin)
}
