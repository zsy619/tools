package xcrypto

import (
	"crypto/cipher"
	"crypto/sha256"
	"encoding/hex"
	"fmt"
	"testing"

	"github.com/tjfoc/gmsm/sm4"
)

func TestSm4Password1(t *testing.T) {
	key := sha256.Sum256([]byte("1234567890abcdef"))
	plaintext := "hello world"
	fmt.Printf("明文: %s\n", plaintext)
	block, err := sm4.NewCipher(key[0:16])
	if err != nil {
		fmt.Println(err.Error())
		return
	}

	bsize := block.BlockSize()
	plaintextBytes := []byte(plaintext)
	plaintextBytes = PKCS7Padding(plaintextBytes, bsize)

	ciphertext := make([]byte, len(plaintextBytes))
	iv := make([]byte, bsize)
	mode := cipher.NewCBCEncrypter(block, iv)
	mode.CryptBlocks(ciphertext, plaintextBytes)

	fmt.Printf("密文: %s\n", hex.EncodeToString(ciphertext))

	decrypted := make([]byte, len(ciphertext))
	mode = cipher.NewCBCDecrypter(block, iv)
	mode.CryptBlocks(decrypted, ciphertext)

	decrypted = PKCS7UnPadding(decrypted)

	fmt.Printf("解密后的明文: %s\n", decrypted)
}

func TestSm4Pkcs7Encrypt(t *testing.T) {
	pwd, err := Sm4Pkcs7Encrypt("Hello, World!", "123456")
	if err != nil {
		t.Fatal(err)
	} else {
		t.Log(pwd)
	}
}

func TestSm4Pkcs7Decrypt(t *testing.T) {
	pwd, err := Sm4Pkcs7Decrypt("a8e27956e447a206cd043d2dbf5ec3de", "123456")
	if err != nil {
		t.Fatal(err)
	} else {
		t.Log(pwd)
	}
}

func TestSm4Encrypt(t *testing.T) {
	pwd, err := Sm4Encrypt("Hello, World!", "123456")
	if err != nil {
		t.Fatal(err)
	} else {
		t.Log(pwd)
	}
}

func TestSm4Decrypt(t *testing.T) {
	pwd, err := Sm4Decrypt("e3a51f2e5774a57722c6a131721d2375", "123456")
	if err != nil {
		t.Fatal(err)
	} else {
		t.Log(pwd)
	}
}

func TestSm4EncryptReverse(t *testing.T) {
	data := "admin@1245"
	key := "admin@2023"
	pwd, err := Sm4EncryptReverse(data, key)
	if err != nil {
		t.Fatal(err)
	} else {
		fmt.Println(pwd)
	}
}

func TestSm4DecryptReverse(t *testing.T) {
	data := "cc7f5c0865c24adb044b7cdc0b95a6bd"
	key := "admin@2023"
	pwd, err := Sm4DecryptReverse(data, key)
	if err != nil {
		t.Fatal(err)
	} else {
		fmt.Println(pwd)
	}
}
