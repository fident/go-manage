package main

import (
	"fmt"

	client "github.com/fident/go-manage"
)

const (
	keyfilePath = "./testkey.json"
)

func main() {
	testClient, err := client.New(keyfilePath, client.FidentInstanceAddressLocal)
	if err != nil {
		panic(err)
	}

	lastlogin, err := testClient.GetAccountDetailsForIdentityID("EFIDFIID-ZGVT5I6L4-MISCR-V5UX35S")
	if err != nil {
		panic(err)
	}

	fmt.Printf("Result: %v\n", lastlogin)
}
