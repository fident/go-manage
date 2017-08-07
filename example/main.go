package main

import (
	"fmt"

	client "github.com/fident/go-manage"
)

const (
	keyfilePath = "./testkey.json"
)

func main() {
	// Create new fident client with path to keyfile.json and fident instance address
	testClient, err := client.New(keyfilePath, client.FidentInstanceAddressLocal)
	if err != nil {
		panic(err)
	}

	// Query account details using client
	details, err := testClient.GetAccountDetailsForIdentityID("EFIDFIID-ZGVT5I6L4-MISCR-V5UX35S")
	if err != nil {
		panic(err)
	}

	fmt.Printf("Details Result: %v\n", details)

	// Query last login timestamp using client
	lastlogin, err := testClient.GetLastLoginTimestampForIdentityID("EFIDFIID-ZGVT5I6L4-MISCR-V5UX35S")
	if err != nil {
		panic(err)
	}

	fmt.Printf("Timestamp Result: %v\n", lastlogin)

}
