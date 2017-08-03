package client

import (
	"context"
	"crypto/tls"
	"errors"
	"time"

	"github.com/fident/go-proto/fident"

	"google.golang.org/grpc"
	"google.golang.org/grpc/credentials"
	"google.golang.org/grpc/metadata"
)

// GetLastLoginTimestampForIdentityID retrieves last login time for given identity
func (i *Instance) GetLastLoginTimestampForIdentityID(identityID string) (time.Time, error) {
	r, err := i.GetLastLoginTimestampsForIdentityIDs([]string{identityID})
	if err != nil {
		return time.Time{}, err
	}

	if t, ok := r[identityID]; ok {
		return t, nil
	}

	return time.Time{}, errors.New("no successful login found for ID")
}

// GetLastLoginTimestampsForIdentityIDs retrieves last login times for given identitys
func (i *Instance) GetLastLoginTimestampsForIdentityIDs(identityIDs []string) (map[string]time.Time, error) {
	meta, err := i.preflightAuth()
	if err != nil {
		return map[string]time.Time{}, err
	}
	ctx := metadata.NewContext(context.Background(), meta)

	// TODO: supply endpoint and cert details
	//conn, err := grpc.Dial(address, grpc.WithTransportCredentials(credentials.NewClientTLSFromCert(nil, "")))
	tlc := &tls.Config{
		InsecureSkipVerify: true, // FOR TESTING ONLY, TODO: Remove this verify/distribute fident.io cert
	}

	conn, err := grpc.Dial(i.fidentEndpoint, grpc.WithTransportCredentials(credentials.NewTLS(tlc)))
	if err != nil {
		return map[string]time.Time{}, err
	}
	defer conn.Close()

	c := fident.NewAuthClient(conn)
	res, err := c.GetLastLoginTimestamps(ctx, &fident.LoginTimestampRequest{
		IdentityId: identityIDs,
	})
	if err != nil {
		return map[string]time.Time{}, err
	}

	fin := map[string]time.Time{}
	for id, unix := range res.Results {
		fin[id] = time.Unix(unix, 0)
	}

	return fin, nil
}
