package client

import (
	"context"
	"errors"
	"time"

	"github.com/fident/go-proto/fident"

	"google.golang.org/grpc"
	"google.golang.org/grpc/credentials"
	"google.golang.org/grpc/metadata"
)

// AccountDetail item structure
type AccountDetail struct {
	IdentityID string
	Username   string
	Attributes map[string]string
	Created    time.Time
}

// GetFirstNameAttribute returns the first name from account detail attributes
func (a *AccountDetail) GetFirstNameAttribute() string {
	return a.Attributes[attributeKeyFirstName]
}

// GetLastNameAttribute returns the last name from account detail attributes
func (a *AccountDetail) GetLastNameAttribute() string {
	return a.Attributes[attributeKeyLastName]
}

// GetEmailAddress returns email address for account
func (a *AccountDetail) GetEmailAddress() string {
	return a.Username
}

// GetAccountDetailsForIdentityID retrieves account details for given identity id
func (i *Instance) GetAccountDetailsForIdentityID(identityID string) (AccountDetail, error) {
	r, err := i.GetAccountDetailsForIdentityIDs([]string{identityID})
	if err != nil {
		return AccountDetail{}, err
	}

	if len(r) > 0 {
		return r[0], nil
	}

	return AccountDetail{}, errors.New("no details found for ID")
}

// GetAccountDetailsForIdentityIDs retrieves account details for given identity ids
func (i *Instance) GetAccountDetailsForIdentityIDs(identityIDs []string) ([]AccountDetail, error) {
	meta, err := i.preflightAuth()
	if err != nil {
		return []AccountDetail{}, err
	}
	ctx := metadata.NewContext(context.Background(), meta)

	conn, err := grpc.Dial(i.fidentEndpoint, grpc.WithTransportCredentials(credentials.NewTLS(fidentTSLConfig)))
	if err != nil {
		return []AccountDetail{}, err
	}
	defer conn.Close()

	c := fident.NewAuthClient(conn)
	res, err := c.GetAccountDetails(ctx, &fident.AccountDetailRequest{
		IdentityId: identityIDs,
	})
	if err != nil {
		return []AccountDetail{}, err
	}

	fin := []AccountDetail{}
	for _, det := range res.Results {
		A := AccountDetail{
			IdentityID: det.GetIdentityId(),
			Username:   det.GetUsername(),
			Created:    time.Unix(det.GetCreated(), 0),
		}

		at := map[string]string{}
		for _, ar := range det.Attributes {
			at[ar.GetKey()] = ar.GetValue()
		}
		A.Attributes = at
		fin = append(fin, A)
	}

	return fin, nil
}
