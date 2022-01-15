package tgverifier

import (
	"crypto/hmac"
	"crypto/sha256"
	"encoding/hex"
	"errors"
	"fmt"
)

// ErrInvalidCredentials represents error in case of having invalid Telegram auth credentials.
var ErrInvalidCredentials = errors.New("invalid telegram credentials")

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
	var (
		secret      = sha256.Sum256(token)
		checkString = c.String()
		authCode    = computeHmac256([]byte(checkString), secret[:])
		hexAuthCode = hex.EncodeToString(authCode)
	)

	if hexAuthCode != c.Hash {
		return ErrInvalidCredentials
	}

	return nil
}

// String builds credentials string, excluding hash field.
func (c *Credentials) String() string {
	s := fmt.Sprint("auth_date=", c.AuthDate, "\nfirst_name=", c.FirstName, "\nid=", c.ID)

	if c.LastName != "" {
		s = fmt.Sprint(s, "\nlast_name=", c.LastName)
	}

	if c.PhotoURL != "" {
		s = fmt.Sprint(s, "\nphoto_url=", c.PhotoURL)
	}

	if c.Username != "" {
		s = fmt.Sprint(s, "\nusername=", c.Username)
	}

	return s
}

func computeHmac256(msg []byte, key []byte) []byte {
	h := hmac.New(sha256.New, key)
	h.Write(msg)
	return h.Sum(nil)
}
