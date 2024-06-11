package main

import (
	"crypto/rand"
	"flag"
	"fmt"
	"log"
)

var (
	Charset = "abcdefghijklmnopqrstuvwxyzABCDEFGHIJKLMNOPQRSTUVWXYZ1234567890"

	// 6 bits to represent 64 possibilities / indexes
	IdxBits = 6

	// All 1-bits, as many as letterIdxBits
	IdxMask = 1<<IdxBits - 1
)

// randomSecureBytes returns the requested number of bytes using crypto/rand
func randomSecureBytes(length int) ([]byte, error) {
	randomBytes := make([]byte, length)
	_, err := rand.Read(randomBytes)
	if err != nil {
		return nil, fmt.Errorf("unable to generate random bytes: %w", err)
	}
	return randomBytes, nil
}

// randomString returns a secure string.
func randomSecureString(length int) (string, error) {
	secureString := make([]byte, length)
	bufferSize := int(float64(length) * 1.3)
	for i, j, randomBytes := 0, 0, []byte{}; i < length; j++ {
		if j%bufferSize == 0 {
			var err error
			randomBytes, err = randomSecureBytes(bufferSize)
			if err != nil {
				return "", err
			}
		}
		if idx := int(randomBytes[j%length] & byte(IdxMask)); idx < len(Charset) {
			secureString[i] = Charset[idx]
			i++
		}
	}
	return string(secureString), nil
}

func main() {
	n := flag.Int("n", 1, "Generate n random strings")
	length := flag.Int("len", 16, "Each string should be len characters long")
	flag.Parse()

	for i := 0; i < *n; i++ {
		secure, err := randomSecureString(*length)
		if err != nil {
			log.Fatal(err)
		}
		if i == *n-1 {
			fmt.Print(secure)
		} else {
			fmt.Println(secure)
		}
	}
}
