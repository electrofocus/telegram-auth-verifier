package tgverifier

import (
	"crypto/hmac"
	"crypto/sha256"
	"encoding/hex"
	"errors"
	"fmt"
)

// ErrInvalidCreds represents error in case of having invalid Telegram auth credentials
var ErrInvalidCreds = errors.New("invalid telegram creds")

// Credentials are Telegram Login credentials available for parsing from JSON.
type Credentials struct {
	ID        int64  `json:"id"`
	FirstName string `json:"first_name"`
	LastName  string `json:"last_name"`
	Username  string `json:"username"`
	PhotoURL  string `json:"photo_url"`
	AuthDate  int64  `json:"auth_date"`
	Hash      string `json:"hash"`
}

// Verify checks if the credentials are from Telegram.
// Returns nil error if credentials are from Telegram.
func (c *Credentials) Verify(token []byte) error {
	secret := sha256.Sum256(token)

	checkString := c.String()

	authCode := computeHmac256([]byte(checkString), secret[:])
	hexAuthCode := hex.EncodeToString(authCode)

	if hexAuthCode != c.Hash {
		return ErrInvalidCreds
	}

	return nil
}

// String builds credentials string, excluding hash field.
func (c *Credentials) String() string {
	s := fmt.Sprintf("auth_date=%d", c.AuthDate)

	if c.FirstName != "" {
		s += fmt.Sprintf("\nfirst_name=%s", c.FirstName)
	}

	s += fmt.Sprintf("\nid=%d", c.ID)

	if c.LastName != "" {
		s += fmt.Sprintf("\nlast_name=%s", c.LastName)
	}

	if c.PhotoURL != "" {
		s += fmt.Sprintf("\nphoto_url=%s", c.PhotoURL)
	}

	if c.Username != "" {
		s += fmt.Sprintf("\nusername=%s", c.Username)
	}

	return s
}

func computeHmac256(msg []byte, key []byte) []byte {
	h := hmac.New(sha256.New, key)
	h.Write(msg)
	return h.Sum(nil)
}
