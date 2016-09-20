package main

import (
	"golang.org/x/crypto/bcrypt"
)

type BCrypt int

func (w *BCrypt) Run(in string, result *string) error {
	hashedString, err := bcrypt.GenerateFromPassword([]byte(in), bcrypt.DefaultCost)
	if err != nil {
		return err
	}

	*result = string(hashedString[:])
	return nil
}
