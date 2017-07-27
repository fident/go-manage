package token

import (
	"encoding/base64"
	"encoding/json"
	"errors"
	"strings"
	"time"
)

// Token type
type Token struct {
	value   string
	expires time.Time
}

type rawTok struct {
	Expires int64 `json:"exp"`
}

// WithinExpiryRange checks if token will expire within next 12 hours or has already expired (due refresh)
func (t *Token) WithinExpiryRange() bool {
	if (t.expires.Unix()-time.Now().Unix()) < (3600*12) || t.value == "" {
		return true
	}
	return false
}

// GetValue returns current tokens value
func (t *Token) GetValue() string {
	return t.value
}

// Parse token
func Parse(token string) (Token, error) {
	parts := strings.Split(token, ".")
	if len(parts) != 3 {
		return Token{}, errors.New("Invalid token")
	}

	claimBytes, err := decodeSegment(parts[1])
	if err != nil {
		return Token{}, err
	}

	var re rawTok
	err = json.Unmarshal(claimBytes, &re)
	if err != nil {
		return Token{}, err
	}

	res := Token{
		value:   token,
		expires: time.Unix(re.Expires, 0),
	}

	return res, nil
}

// decode JWT specific base64url encoding with padding stripped
func decodeSegment(seg string) ([]byte, error) {
	if l := len(seg) % 4; l > 0 {
		seg += strings.Repeat("=", 4-l)
	}
	return base64.URLEncoding.DecodeString(seg)
}
