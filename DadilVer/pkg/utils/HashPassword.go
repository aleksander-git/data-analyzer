package utils

import (
	"crypto/rand"
	"crypto/sha256"
	"encoding/hex"
)

//ЧИСТО ПРИМЕР ИСПОЛЬЗОВАНИЯ

// HashPassword хэширует пароль с использованием SHA-256
func HashPassword(password string) string {
	hash := sha256.Sum256([]byte(password))
	return hex.EncodeToString(hash[:])
}

// GenerateRandomString генерирует случайную строку заданной длины
func GenerateRandomString(n int) (string, error) {
	const letters = "abcdefghijklmnopqrstuvwxyzABCDEFGHIJKLMNOPQRSTUVWXYZ0123456789"
	bytes, err := generateRandomBytes(n)
	if err != nil {
		return "", err
	}
	for i, b := range bytes {
		bytes[i] = letters[b%byte(len(letters))]
	}
	return string(bytes), nil
}

// generateRandomBytes генерирует случайный байтовый массив заданной длины
func generateRandomBytes(n int) ([]byte, error) {
	bytes := make([]byte, n)
	_, err := rand.Read(bytes)
	if err != nil {
		return nil, err
	}
	return bytes, nil
}
