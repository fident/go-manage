package manage

import (
	"context"
	"net/http"
	"time"

	"github.com/fident/go-manage/fidentapi"
	"github.com/fident/go-manage/tls"

	"google.golang.org/grpc"
	"google.golang.org/grpc/credentials"
	"google.golang.org/grpc/metadata"
)

// CreateUserAccount creates a new user account returning identity ID and optionally a temporary session cookie
func (i *Instance) CreateUserAccount(email string, tempSessionCookie bool) (string, *http.Cookie, error) {
	meta, err := i.preflightAuth()
	if err != nil {
		return "", nil, err
	}
	ctx := metadata.NewOutgoingContext(context.Background(), meta)

	conn, err := grpc.Dial(i.fidentEndpoint, grpc.WithTransportCredentials(credentials.NewTLS(tls.FidentTSLConfig)))
	if err != nil {
		return "", nil, err
	}
	defer conn.Close()

	c := fidentapi.NewAuthClient(conn)
	res, err := c.CreateUserAccount(ctx, &fidentapi.CreateUserAccountRequest{
		EmailAddress:        email,
		IssueTemporaryToken: tempSessionCookie,
	})
	if err != nil {
		return "", nil, err
	}

	var cookie *http.Cookie

	if tempSessionCookie {
		cookie = &http.Cookie{
			Name:     res.TemporaryToken.Name,
			Value:    res.TemporaryToken.Value,
			Domain:   res.TemporaryToken.Domain,
			Expires:  time.Unix(res.TemporaryToken.ExpiresEpoch, 0),
			Secure:   res.TemporaryToken.Secure,
			HttpOnly: res.TemporaryToken.HttpOnly,
		}
	}

	return res.IdentityId, cookie, err
}
