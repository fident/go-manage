package manage

import (
	"fmt"
	"os"
)

const (
	//FidentInstanceAddressLocalHost is the address of local fident instance
	FidentInstanceAddressLocalHost = "localhost:50052"

	// FidentInstanceAddressSharedLocal is the address of shared internal fident instance
	FidentInstanceAddressSharedLocal = "HOST$PORT"

	// FidentInstancePublic is the address of production/public fident instance
	FidentInstancePublic = "fident.io:50052"

	// Deckard env var keys
	deckardFidentGRPCHostKey = "FIDENT_SERVICE_HOST_GRPC"
	deckardFidentGRPCPortKey = "FIDENT_SERVICE_PORT_GRPC"
)

func getAddressFromLocal() string {
	host := os.Getenv(deckardFidentGRPCHostKey)
	port := os.Getenv(deckardFidentGRPCPortKey)
	r := fmt.Sprintf("%s:%s", host, port)
	return r
}
