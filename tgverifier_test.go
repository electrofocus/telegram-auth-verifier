package tgverifier

import (
	"encoding/json"
	"log"
	"testing"
)

func TestVerifyAuthentic(t *testing.T) {
	b := []byte(`{
		"id": 151272257,
		"first_name": "Амир",
		"last_name": "Муллагалиев",
		"username": "amirmullagaliev",
		"photo_url": "https://t.me/i/userpic/320/vnCWM1N8YDg12Q_m87JFNG9Ey2cy3YoOCCP6AkFkMfY.jpg",
		"auth_date": 1615827772,
		"hash": "71df938daeec09e0843bc37d766bf2d8e5efac14512d1ffb51263562c0e10a8e"
	}`)

	var cred Credentials
	if err := json.Unmarshal(b, &cred); err != nil {
		log.Println(err)
	}

	if err := cred.Verify([]byte("1622902058:AAEbul3sgosfAvYvd8S-B6zDqMfwGZYL7wk")); err != nil {
		t.Error("credentials are not authentic")
	}
}

func TestVerifyNotAuthentic(t *testing.T) {
	b := []byte(`{
		"id": 151272257,
		"first_name": "Амир",
		"last_name": "Муллагалиев",
		"username": "amirmullagaliev",
		"photo_url": "https://t.me/i/userpic/320/vnCWM1N8YDg12Q_m87JFNG9Ey2cy3YoOCCP6AkFkMfY.jpg",
		"auth_date": 1615827772,
		"hash": "71df938daeec09e0843bc37d766f2d8e5efac14512d1ffb51263562c0e10a8e"
	}`)

	var cred Credentials
	if err := json.Unmarshal(b, &cred); err != nil {
		log.Println(err)
	}

	if err := cred.Verify([]byte("1622902058:AAEbul3sgosfAvYvd8S-B6zDqMfwGZYL7wk")); err == nil {
		t.Error("not authentic credentials verified as authentic")
	}
}
