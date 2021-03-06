package manage

import (
	"context"

	"github.com/fident/go-manage/fidentapi"
	"github.com/fident/go-manage/tls"

	"google.golang.org/grpc"
	"google.golang.org/grpc/credentials"
	"google.golang.org/grpc/metadata"
)

// GetAllIdentityIDsForProject retrieves all identity ids for your fident project
func (i *Instance) GetAllIdentityIDsForProject() ([]string, error) {
	meta, err := i.preflightAuth()
	if err != nil {
		return []string{}, err
	}
	ctx := metadata.NewOutgoingContext(context.Background(), meta)

	conn, err := grpc.Dial(i.fidentEndpoint, grpc.WithTransportCredentials(credentials.NewTLS(tls.FidentTSLConfig)))
	if err != nil {
		return []string{}, err
	}
	defer conn.Close()

	c := fidentapi.NewAuthClient(conn)
	res, err := c.GetAllIdentityIDs(ctx, &fidentapi.IdentityIDsRequest{})
	if err != nil {
		return []string{}, err
	}

	return res.IdentityId, nil
}
