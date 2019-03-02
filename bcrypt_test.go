package crypto

import (
	"github.com/geniehunter/io/util"
	"golang.org/x/crypto/bcrypt"
	"testing"
)

func TestBcrypt(t *testing.T) {
	password := util.RandString(32)
	hash, err := Bcrypt(password, bcrypt.DefaultCost)
	if err != nil {
		t.Fatal(err)
	}
	if !BcryptCheck(password, hash) {
		t.Fatal("BcryptCheck failed")
	}
}
