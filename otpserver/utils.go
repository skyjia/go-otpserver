package otpserver

import (
	"encoding/base32"
	"fmt"
	"net/url"
	"strings"
)

const (
	// URIScheme is the URI scheme
	URIScheme = "otpauth"
)

// Generate the URL according to Google Authentcator's Key Uri Format
//   https://github.com/google/google-authenticator/wiki/Key-Uri-Format
func generateAuthURL(otpType, issuer, accountName, secret string) string {
	encodedSecret := encodeSecret(secret)

	u := &url.URL{
		Scheme: URIScheme,
		Host:   otpType,
		Path:   fmt.Sprintf("%s:%s", issuer, accountName),
	}

	q := u.Query()
	q.Set("issuer", issuer)
	q.Set("secret", encodedSecret)
	u.RawQuery = q.Encode()

	return u.String()
}

//
func encodeSecret(secret string) string {
	secretBytes := []byte(secret)
	encodedString := base32.StdEncoding.EncodeToString(secretBytes)
	return strings.TrimRight(encodedString, "=")
}
