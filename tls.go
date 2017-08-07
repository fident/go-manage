package manage

import (
	"crypto/tls"
	"crypto/x509"
	"errors"
)

// Fident GRPC certificate
var fidentGRPCCertificate = `-----BEGIN CERTIFICATE-----
MIIC+zCCAeOgAwIBAgIJAKqY7FjWykkAMA0GCSqGSIb3DQEBBQUAMBQxEjAQBgNV
BAMMCWxvY2FsaG9zdDAeFw0xNjA1MjMxMDU1MzRaFw0yNjA1MjExMDU1MzRaMBQx
EjAQBgNVBAMMCWxvY2FsaG9zdDCCASIwDQYJKoZIhvcNAQEBBQADggEPADCCAQoC
ggEBAL2ceyYn8A+1+rf8ysvmOqDrNDo+fOSp40sbg706dHpkm1z5CGs3mwqkb638
vHeiDFINQZR+FLoaQjxWmMSmyI05YS9zcwMd4r1lWpjRrAb3awjDm7nD63puFvyS
0e1W6rGMFzU6bAPqNEHP8pt35cqfJt7NtLqt/et9QHiH9Q0sOnTcY1ITutDsnA0B
owtRFRExsVyRGtCkczFL6cT2/wXrpcdkb/803Hz3KKRUStkfACl6OrKMNz2xfLiI
T5gd8CP89GnyqDRkp2brAitabMzz637pw1bHFJgxoKSu5lg/71OJ2T5sQ2rRrovh
nC3KQ7DzD9W4iY5mj3Pc+LfNt2sCAwEAAaNQME4wHQYDVR0OBBYEFNwKCQFalNDE
MSC8wrqpUKMgjK12MB8GA1UdIwQYMBaAFNwKCQFalNDEMSC8wrqpUKMgjK12MAwG
A1UdEwQFMAMBAf8wDQYJKoZIhvcNAQEFBQADggEBAASkrccLQF3FEp7wI/otvy/N
ZmY2jnQ7sAFegp9mLWtGCmnEKiOMxJj7LdgaRzFTiYo/nI+g16Vq/Qm0LnIVLNT2
yPhK3Wc+PsSHd/HK4vieBml2jAonMysk/rV/95lMF0ll2ZFm4AadkD6+R+dL9Lt+
MgDltl0SvQ0xSVmfIyKqQubd2e0caI5yOL7xZsP38xLi3Ou+KcjyP6PP4m9BD7i9
gqyKWdvnD3FeTn2z4pC2iQcCPQnxFNTzaSYeQvE9c3/9JMXMW5LEChb6zEBMCrsC
XQP3Oyk7j0T4X7M8+vWLRhqiJ1j7TV40qNVWSgLxHRp6H0c2sCFrlcywqBl2EAk=
-----END CERTIFICATE-----`

var fidentTSLConfig *tls.Config

// internal TLS init
func initTLS() error {
	cp := x509.NewCertPool()
	if !cp.AppendCertsFromPEM([]byte(fidentGRPCCertificate)) {
		return errors.New("credentials: failed to append certificates")
	}
	fidentTSLConfig = &tls.Config{RootCAs: cp}
	return nil
}

func overrideFidentCerficate(certPEM string) {
	if certPEM != "" {
		fidentGRPCCertificate = certPEM
	}
}
