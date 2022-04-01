package authentication

import (
	"crypto/rand"
	"encoding/hex"
)

func GenerateApiKey() string {
	b := make([]byte, 128)
	if _, err := rand.Read(b); err != nil {
		return ""
	}
	return hex.EncodeToString(b)
}
