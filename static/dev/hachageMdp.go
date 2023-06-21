package dev

import (
	"fmt"

	"golang.org/x/crypto/bcrypt"
)

func Hash(password string) string {
	bytes, errGenerate := bcrypt.GenerateFromPassword([]byte(password), 14)
	if errGenerate != nil {
		fmt.Println("l'erreur est:", errGenerate)
	}
	return string(bytes)
}
