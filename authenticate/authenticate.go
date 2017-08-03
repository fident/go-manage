package authenticate

import (
	"context"
	"crypto/tls"

	"github.com/fident/go-manage/key"
	"github.com/fident/go-proto/fident"

	jwt "github.com/dgrijalva/jwt-go"
	"google.golang.org/grpc"
	"google.golang.org/grpc/credentials"
)

const challengeParam = "challenge_token"

// GetToken retrieves a new authentication token for Fident requests
func GetToken(key key.Key, fidentEndpoint string) (string, error) {

	// TODO: supply endpoint and cert details
	//conn, err := grpc.Dial(address, grpc.WithTransportCredentials(credentials.NewClientTLSFromCert(nil, "")))
	tlc := &tls.Config{
		InsecureSkipVerify: true, // FOR TESTING ONLY, TODO: Remove this verify/distribute fident.io cert
	}

	conn, err := grpc.Dial(fidentEndpoint, grpc.WithTransportCredentials(credentials.NewTLS(tlc)))
	if err != nil {
		return "", err
	}

	defer conn.Close()

	c := fident.NewAuthClient(conn)

	// Contact fident and print out challenge
	r, err := c.GetAuthenticationChallenge(context.Background(), &fident.AuthChallengeRequest{
		IdentityId: key.IdentityID,
		ProjectId:  key.ProjectID,
	})
	if err != nil {
		return "", err
	}

	// 'Optional' verification of challenge source here
	//(endpoint is already verified by GRPC connection so there should be no point)

	challengeResponseToken := jwt.NewWithClaims(jwt.SigningMethodRS256, jwt.MapClaims{
		challengeParam: r.Challenge,
	})

	response, err := challengeResponseToken.SignedString(key.PrivateKey)
	if err != nil {
		return "", err
	}

	tr, err := c.PerformAuthentication(context.Background(), &fident.PerformAuthRequest{
		IdentityId:        key.IdentityID,
		KeyHandle:         key.KeyHandle,
		ProjectId:         key.ProjectID,
		ChallengeResponse: response,
	})
	if err != nil {
		return "", err
	}

	return tr.GetToken(), nil
}
