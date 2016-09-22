package main

import (
	"golang.org/x/crypto/bcrypt"
	"log"
	"time"
)

type BCrypt int

func (w *BCrypt) Run(in string, result *string) error {
	hashedString, err := bcrypt.GenerateFromPassword([]byte(in), bcrypt.DefaultCost)
	if err != nil {
		return err
	}

	*result = string(hashedString[:])
	log.Printf("Input: %s; Output: %s", in, *result)

	if *slow {
		time.Sleep(time.Second)
	}

	return nil
}
