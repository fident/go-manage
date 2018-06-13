package manage

import (
	"context"

	"github.com/fident/go-manage/fidentapi"
	"github.com/fident/go-manage/permissions"
	"github.com/fident/go-manage/tls"

	"google.golang.org/grpc"
	"google.golang.org/grpc/credentials"
	"google.golang.org/grpc/metadata"
)

// GetManagementPermissions retrieves all management permissions for given identity id
func (i *Instance) GetManagementPermissions(identityID string) ([]permissions.Entry, error) {
	meta, err := i.preflightAuth()
	if err != nil {
		return []permissions.Entry{}, err
	}
	ctx := metadata.NewOutgoingContext(context.Background(), meta)

	conn, err := grpc.Dial(i.fidentEndpoint, grpc.WithTransportCredentials(credentials.NewTLS(tls.FidentTSLConfig)))
	if err != nil {
		return []permissions.Entry{}, err
	}
	defer conn.Close()

	c := fidentapi.NewAuthClient(conn)
	res, err := c.GetManagementPermissionsForIdentityIDs(ctx, &fidentapi.GetManagementPermissionsRequest{
		IdentityId: []string{identityID},
	})

	if err != nil {
		return []permissions.Entry{}, err
	}

	fin := []permissions.Entry{}
	for _, re := range res.Results {
		for _, pem := range re.Permissions {
			fin = append(fin, permissions.Entry(pem))
		}
	}

	return fin, nil
}

// AddManagementPermission adds given management permission to given identity id
func (i *Instance) AddManagementPermission(identityID string, permission permissions.Entry) error {
	meta, err := i.preflightAuth()
	if err != nil {
		return err
	}

	ctx := metadata.NewOutgoingContext(context.Background(), meta)
	conn, err := grpc.Dial(i.fidentEndpoint, grpc.WithTransportCredentials(credentials.NewTLS(tls.FidentTSLConfig)))
	if err != nil {
		return err
	}
	defer conn.Close()

	c := fidentapi.NewAuthClient(conn)
	_, err = c.AddManagementPermissionToIdentityIDs(ctx, &fidentapi.AddManagementPermissionRequest{
		IdentityId: []string{identityID},
		Permission: string(permission),
	})
	if err != nil {
		return err
	}

	return nil
}

// RemoveManagementPermission removes given management permission from given identity id
func (i *Instance) RemoveManagementPermission(identityID string, permission permissions.Entry) error {
	meta, err := i.preflightAuth()
	if err != nil {
		return err
	}

	ctx := metadata.NewOutgoingContext(context.Background(), meta)
	conn, err := grpc.Dial(i.fidentEndpoint, grpc.WithTransportCredentials(credentials.NewTLS(tls.FidentTSLConfig)))
	if err != nil {
		return err
	}
	defer conn.Close()

	c := fidentapi.NewAuthClient(conn)
	_, err = c.RemoveManagementPermissionFromIdentityIDs(ctx, &fidentapi.RemoveManagementPermissionRequest{
		IdentityId: []string{identityID},
		Permission: string(permission),
	})
	if err != nil {
		return err
	}

	return nil
}
