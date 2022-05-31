package utils

import (
	"crypto/sha1"
	"encoding/base64"
	"math/rand"
	"strings"
	"time"
)

func RandString(n int) string {
	rand.Seed(time.Now().UnixNano())
	var letterRunes = []rune("abcdefghijklmnopqrstuvwxyzABCDEFGHIJKLMNOPQRSTUVWXYZ")

	random := make([]rune, n)
	for i := range random {
		random[i] = letterRunes[rand.Intn(len(letterRunes))]
	}

	return string(random)
}

func GenerateHash(apiKey, secretKey, rnd, pki string) string {
	hashString := apiKey + rnd + secretKey + pki
	hasher := sha1.New()
	hasher.Write([]byte(hashString))
	hash := base64.URLEncoding.EncodeToString(hasher.Sum(nil))

	hash = strings.Replace(hash, "_", "/", -1)
	hash = strings.Replace(hash, "-", "+", -1)

	return hash
}
