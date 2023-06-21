package dev

import (
	"fmt"

	"golang.org/x/crypto/bcrypt"
)

func check_hash(mdp string, hash string) bool {
	err := bcrypt.CompareHashAndPassword([]byte(hash), []byte(mdp))
	if err != nil {
		fmt.Println("Le mot de passe est incorrect.")
		return false
	}

	return true
}
