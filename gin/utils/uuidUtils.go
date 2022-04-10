package utils

import (
	"crypto/rand"
	"encoding/base64"
	"github.com/satori/go.uuid"
)

func GetUUID() string {
	return uuid.NewV4().String()
}

func GenUUID(s int) string {
	b, err := generateRandomBytes(s)
	CheckErr(err)
	return base64.URLEncoding.EncodeToString(b)
}

func generateRandomBytes(n int) ([]byte, error) {
	b := make([]byte, n)
	_, err := rand.Read(b)
	// Note that err == nil only if we read len(b) bytes.
	if err != nil {
		return nil, err
	}

	return b, nil
}
