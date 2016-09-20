package main

import (
	"golang.org/x/crypto/bcrypt"
	"testing"
)

var bcryptTests = []string{"abc", "h89j34t", "89034jgnvčćž"}

func TestBCrypt(t *testing.T) {
	reverse := new(BCrypt)
	for i, test := range bcryptTests {
		var out string
		reverse.Run(test, &out)

		err := bcrypt.CompareHashAndPassword([]byte(out), []byte(test))

		if err != nil {
			t.Errorf("BCrypt hash was wrong for test case %d (%s)\ngot:\n\t%s", i, test, out)
		}
	}
}
