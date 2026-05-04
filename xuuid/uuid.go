package xuuid

import (
	"crypto/rand"
	"encoding/base64"
	"fmt"
	"log"
	"strings"
)

func NewUUID() string {
	b := make([]byte, 16)
	_, err := rand.Read(b)
	if err != nil {
		log.Fatal(err)
	}
	uuid := fmt.Sprintf("%x-%x-%x-%x-%x", b[0:4], b[4:6], b[6:8], b[8:10], b[10:])
	return uuid
}

func GetAuth() string {
	auth := NewUUID()
	auth = strings.Replace(auth, "-", "", -1)
	auth = auth[:16]
	return auth
}

func GetAes() string {
	aes := NewUUID()
	aes = strings.Replace(aes, "-", "", -1)
	return aes
}

// GenerateRandomBytes returns securely generated random bytes.
// It will return an error if the system's secure random
// number generator fails to function correctly, in which
// case the caller should not continue.
func GenerateRandomBytes(n int) ([]byte, error) {
	b := make([]byte, n)
	_, err := rand.Read(b)
	// Note that err == nil only if we read len(b) bytes.
	if err != nil {
		return nil, err
	}

	return b, nil
}

func GenerateRandomKey(keyLen int) (string, error) {
	key, err := GenerateRandomBytes(keyLen)
	if err != nil {
		return "", err
	}
	fmt.Println("GenerateRandomKey：", fmt.Sprintf("%x", key))
	return base64.StdEncoding.EncodeToString(key), nil
}

func GenerateAESSecret16Key() (string, error) {
	return GenerateRandomKey(16)
}
