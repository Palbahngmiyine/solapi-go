package authenticator

import (
	"crypto/hmac"
	cr "crypto/rand"
	"crypto/sha256"
	"encoding/hex"
	"fmt"
	"time"
)

// Authenticator holds authentication information
type Authenticator struct {
	APIKey    string
	APISecret string
}

// NewAuthenticator creates a new Authenticator instance
func NewAuthenticator(apiKey, apiSecret string) *Authenticator {
	return &Authenticator{
		APIKey:    apiKey,
		APISecret: apiSecret,
	}
}

// RandomString returns a random string
func RandomString(n int) string {
	b := make([]byte, n)
	if _, err := cr.Read(b); err != nil {
		panic(err) // Critical error if we can't get random bytes
	}

	return hex.EncodeToString(b)
}

// GetAuthorization gets the authorization
func (a *Authenticator) GetAuthorization() string {
	salt := RandomString(20)
	date := time.Now().Format(time.RFC3339)
	h := hmac.New(sha256.New, []byte(a.APISecret))
	h.Write([]byte(date + salt))
	signature := hex.EncodeToString(h.Sum(nil))
	authorization := fmt.Sprintf("HMAC-SHA256 apiKey=%s, date=%s, salt=%s, signature=%s", a.APIKey, date, salt, signature)
	return authorization
}
