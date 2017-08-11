package manage

import (
	"log"

	"github.com/fident/go-manage/authenticate"
	"github.com/fident/go-manage/key"
	"github.com/fident/go-manage/tls"
	"github.com/fident/go-manage/token"

	"google.golang.org/grpc/metadata"
)

// Instance is go-managem client instance
type Instance struct {
	authenticated  bool
	key            key.Key
	activeToken    token.Token
	fidentEndpoint string
}

// New client instance
func New(keyfilePath, fidentInstanceAddress string) (Instance, error) {
	in := Instance{}
	err := in.Init(keyfilePath, fidentInstanceAddress)
	if err != nil {
		return in, err
	}
	return in, nil
}

// IsAuthenticated checks clients authentication status
func (i *Instance) IsAuthenticated() bool {
	return i.authenticated
}

// Init client with a new keyfile and endpoint
func (i *Instance) Init(keyfilePath, fidentInstanceAddress string) error {
	// Get fident instance
	if fidentInstanceAddress == FidentInstanceAddressSharedLocal {
		fidentInstanceAddress = getAddressFromLocal()
	}
	// Read key file
	key, err := key.FromFile(keyfilePath)
	if err != nil {
		log.Fatalf("failed to read keyfile: %v", err)
		return err
	}
	i.key = key
	i.fidentEndpoint = fidentInstanceAddress
	tls.InitTLS()
	return nil
}

func (i *Instance) updateToken() error {
	// Use key to get token
	tok, err := authenticate.GetToken(i.key, i.fidentEndpoint)
	if err != nil {
		return err
	}

	// Parse token for local management
	res, err := token.Parse(tok)
	if err != nil {
		return err
	}
	i.activeToken = res
	i.authenticated = true
	return nil
}

// preflightAuth checks current active token status before making request handling token expiry & refresh
func (i *Instance) preflightAuth() (metadata.MD, error) {
	if i.authenticated == false || i.activeToken.GetValue() == "" || i.activeToken.WithinExpiryRange() {
		err := i.updateToken()
		if err != nil {
			return metadata.MD{}, err
		}
	}

	meta := metadata.New(map[string]string{projectIDMetaKey: i.key.ProjectID, tokenMetaKey: i.activeToken.GetValue()})
	return meta, nil
}
