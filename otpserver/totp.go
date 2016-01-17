package otpserver

import (
	"github.com/ant0ine/go-json-rest/rest"
	otp "github.com/hgfischer/go-otp"
	"net/http"
)

// VerifyTOTP verifies a token with current secret
func VerifyTOTP(token, secret string) bool {
	totp := &otp.TOTP{
		Secret:         secret,
		IsBase32Secret: false,
		Length:         GoogleAuthenticatorDigits,
		Period:         GoogleAuthenticatorPeriod,
		WindowBack:     TOTPDefaultWindowBack,
		WindowForward:  TOTPDefaultWindowForward,
	}

	return totp.Verify(token)
}

// GenerateTOTPAuthURL generates a TOTP auth URL for Google Authenticator
func GenerateTOTPAuthURL(issuer, accountName, secret string) string {
	return generateAuthURL(TOTP, issuer, accountName, secret)
}

func verifyTOTPHandler(w rest.ResponseWriter, req *rest.Request) {
	type verifyTOTPRequest struct {
		Token  string `json:"token"`
		Secret string `json:"secret"`
	}

	body := verifyTOTPRequest{}
	if err := req.DecodeJsonPayload(&body); err != nil {
		rest.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}

	isValid := VerifyTOTP(body.Token, body.Secret)
	w.WriteJson(map[string]interface{}{"isValid": isValid})
}

func generateTOTPAuthURLHandler(w rest.ResponseWriter, req *rest.Request) {
	type generateTOTPAuthURLRequest struct {
		Issuer      string `json:"issuer"`
		AccountName string `json:"accountName"`
		Secret      string `json:"secret"`
	}

	body := generateTOTPAuthURLRequest{}
	if err := req.DecodeJsonPayload(&body); err != nil {
		rest.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}

	authURL := GenerateTOTPAuthURL(body.Issuer, body.AccountName, body.Secret)
	w.WriteJson(map[string]interface{}{"authRUL": authURL})
}
