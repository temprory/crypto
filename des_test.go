package crypto

import (
	"github.com/geniehunter/io/util"
	"testing"
)

var cMyKey = "A39%&rtyyuuKA39%&rtyyuuK"

func TestTripleDesEncrypt(t *testing.T) {
	origData := util.RandString(32)
	enc, err := TripleDesEncrypt([]byte(origData), []byte(cMyKey))
	if err != nil {
		t.Fatal(err)
	}
	t.Log("encoded", enc)
	decrypt, err := TripleDesDecrypt(enc, []byte(cMyKey))
	if err != nil {
		t.Fatal(err)
	}
	if string(decrypt) != origData {
		t.Fatal("decrypt failed")
	}
}

func TestTripleDesB64Encrypt(t *testing.T) {
	origData := util.RandString(32)
	enc, err := TripleDesB64Encrypt(origData, cMyKey)
	if err != nil {
		t.Fatal(err)
	}
	t.Log("encoded", enc)
	decrypt, err := TripleDesB64Decrypt(enc, cMyKey)
	if err != nil {
		t.Fatal(err)
	}
	if string(decrypt) != origData {
		t.Fatal("decrypt failed")
	}
}
