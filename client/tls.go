package client

import "crypto/tls"

// TODO: supply endpoint and cert details
//conn, err := grpc.Dial(address, grpc.WithTransportCredentials(credentials.NewClientTLSFromCert(nil, "")))
var fidentTSLConfig = &tls.Config{
	InsecureSkipVerify: true, // FOR TESTING ONLY, TODO: Remove this verify/distribute fident.io cert
}
