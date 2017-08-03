package main

import (
	"fmt"

	"github.com/fident/go-manage/client"
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

	lastlogin, err := testClient.GetLastLoginTimestampForIdentityID("EFIDFIID-ZGVT5I6L4-MISCR-V5UX35S")
	if err != nil {
		panic(err)
	}

	fmt.Printf("Last login was %s\n", lastlogin.String())
}
