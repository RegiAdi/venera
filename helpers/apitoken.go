package helpers

import (
	"crypto/rand"
	"encoding/base64"
	"time"

	"github.com/RegiAdi/hatchet/config"
)

// GenerateRandomBytes returns securely generated random bytes.
// It will return an error if the system's secure random
// number generator fails to function correctly, in which
// case the caller should not continue.
func generateRandomBytes(n int) ([]byte, error) {
	b := make([]byte, n)
	_, err := rand.Read(b)
	// Note that err == nil only if we read len(b) bytes.
	if err != nil {
		return nil, err
	}

	return b, nil
}

// GenerateRandomString returns a URL-safe, base64 encoded
// securely generated random string.
func GenerateAPIToken() (string, error) {
	b, err := generateRandomBytes(32)

	return base64.URLEncoding.EncodeToString(b), err
}

func GenerateAPITokenExpiration() time.Time {
	currentTime := GetCurrentTime()

	return currentTime.AddDate(0, 0, config.GetTokenDuration())
}
