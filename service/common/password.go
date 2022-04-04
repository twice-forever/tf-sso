package common

import (
	"crypto/sha256"
	"fmt"

	"github.com/google/uuid"
)

func GenerateSalt() string {
	uuid := uuid.New().String()
	sum := sha256.Sum256([]byte(uuid))
	return fmt.Sprintf("%x", sum)
}

func GeneratePassHash(password, salt string) string {
	sum := sha256.Sum256([]byte(password + salt))
	return fmt.Sprintf("%x", sum)
}
