package main

import (
	"fmt"

	client "github.com/fident/go-manage"
)

const (
	keyfilePath = "./cubex_cloud_testkey.json"
)

func main() {
	// Create new fident client with path to keyfile.json and fident instance address
	// (Note you can use 'client.FidentInstanceAddressSharedLocal' variable to connect using deckard env vars)
	testClient, err := client.New(keyfilePath, client.FidentInstanceAddressSharedLocal)
	if err != nil {
		panic(err)
	}

	// Query all identity ids for project
	ids, err := testClient.GetAllIdentityIDsForProject()
	if err != nil {
		panic(err)
	}

	fmt.Printf("All IDs Result: %v\n", ids)

	// Query account details using client
	/*
		details, err := testClient.GetAccountDetailsForIdentityID("EFIDFIID-ZGVT5I6L4-MISCR-V5UX35S")
		if err != nil {
			panic(err)
		}

		fmt.Printf("Details Result: %v\n", details)
	*/

	// Query last login timestamp using client
	/*
		lastlogin, err := testClient.GetLastLoginTimestampForIdentityID("EFIDFIID-ZGVT5I6L4-MISCR-V5UX35S")
		if err != nil {
			panic(err)
		}

		fmt.Printf("Timestamp Result: %v\n", lastlogin)
	*/

	// Adding management permission example
	/*
		err = testClient.AddManagementPermission("EFIDFIID-ZGTMFQ1EO-MISCR-IAPZNMF", permissions.PermissionAll)
		if err != nil {
			panic(err)
		}
	*/

	// Creating user account example
	/*
		id, _, err := testClient.CreateUserAccount("tetcom", false)
		if err != nil {
			panic(err)
		}

		fmt.Printf("Created account with ID: %v\n", id)
	*/
}
