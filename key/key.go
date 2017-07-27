package key

import (
	"crypto"
	"crypto/rand"
	"crypto/rsa"
	"crypto/sha256"
	"crypto/x509"
	"encoding/base64"
	"encoding/json"
	"encoding/pem"
	"errors"
	"io/ioutil"
)

// Key is the structure for app authentication JSON files
type Key struct {
	UserType      string          `json:"type"`
	IdentityID    string          `json:"identity_id"`
	ProjectID     string          `json:"project_id"`
	ServiceName   string          `json:"service_name"`
	PrivateKeyPEM string          `json:"private_key"`
	PrivateKey    *rsa.PrivateKey `json:"-"`
	KeyHandle     string          `json:"key_handle"`
}

func FromFile(path string) (Key, error) {
	b, err := ioutil.ReadFile(path)
	if err != nil {
		return Key{}, err
	}
	var res Key
	json.Unmarshal(b, &res)
	var errt error
	res.PrivateKey, errt = parseRSAPrivateKeyFromPEM([]byte(res.PrivateKeyPEM))
	if errt != nil {
		return Key{}, errt
	}
	return res, nil
}

func parseRSAPrivateKeyFromPEM(key []byte) (*rsa.PrivateKey, error) {
	block, _ := pem.Decode(key)
	if block == nil {
		return nil, errors.New("Invalid PEM key")
	}

	var parsedKey interface{}
	parsedKey, err := x509.ParsePKCS1PrivateKey(block.Bytes)
	if err != nil {
		parsedKey, err = x509.ParsePKCS8PrivateKey(block.Bytes)
		if err != nil {
			return nil, err
		}
	}

	var pkey *rsa.PrivateKey
	var ok bool
	if pkey, ok = parsedKey.(*rsa.PrivateKey); !ok {
		return nil, errors.New("Invalid private key")
	}

	return pkey, nil
}

// SignAuthChallenge signs given auth challenge
func (k *Key) SignAuthChallenge(challenge string) (string, error) {
	hash := sha256.New()
	hash.Write([]byte(challenge))
	hashresult := hash.Sum(nil)
	result, err := rsa.SignPKCS1v15(rand.Reader, k.PrivateKey, crypto.SHA256, hashresult)
	if err != nil {
		return "", err
	}
	uEnc := base64.URLEncoding.EncodeToString([]byte(result))
	return uEnc, nil
}
