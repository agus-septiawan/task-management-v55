package utils

import (
	"crypto/rand"
	"crypto/sha256"
	"encoding/base64"
	"fmt"
	"log"
)

func generateSecureToken() (string, error) {
	// Buat byte acak dengan panjang 32 byte
	randomBytes := make([]byte, 32)
	_, err := rand.Read(randomBytes)
	if err != nil {
		return "", fmt.Errorf("gagal menghasilkan byte acak: %v", err)
	}

	// Hash byte acak menggunakan SHA-256
	hash := sha256.Sum256(randomBytes)

	// Encode hasil hash ke dalam format base64 untuk token yang lebih aman
	token := base64.URLEncoding.EncodeToString(hash[:])

	return token, nil
}

func main() {
	// Generate token
	token, err := generateSecureToken()
	if err != nil {
		log.Fatalf("Error: %v", err)
	}

	// Tampilkan token
	fmt.Printf("Access Token: %s\n", token)
}
