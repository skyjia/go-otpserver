package otpserver

import (
	"encoding/base32"
	"fmt"
	"strings"
)

func generateAuthURL(otpType, issuer, accountName, secret string) string {
	encodedSecret := getTrimedBase32Secret(secret)
	url := fmt.Sprintf("otpauth://%s/%s:%s?secret=%s&issuer=%s", otpType, issuer, accountName, encodedSecret, issuer)
	return url
}

func getTrimedBase32Secret(secret string) string {
	secretBytes := []byte(secret)
	encodedString := base32.StdEncoding.EncodeToString(secretBytes)
	return strings.TrimRight(encodedString, "=")
}
