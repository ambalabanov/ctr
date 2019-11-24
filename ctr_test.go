package ctr

import (
	"encoding/hex"
	"testing"
)

func TestCrypt(t *testing.T) {
	key, _ := hex.DecodeString("5F4DCC3B5AA765D61D8327DEB882CF992B95990A9151374ABD8FF8C5A7A0FE08")
	iv, _ := hex.DecodeString("B7B4372CDFBCB3D16A2631B59B509E94")
	ct, _ := hex.DecodeString("d860d4f47437cc935feb8c005616")
	pt := string(Crypt(ct, key, iv))
	if pt != "secret message" {
		t.Errorf("Decrypt error")
	}
	pt2 := []byte("Here is a string....")
	ct2 := hex.EncodeToString(Crypt(pt2, key, iv))
	if ct2 != "e360c5e3312a9fde5bb88c15431a971a171d9958" {
		t.Errorf("Encrypt error")
	}
}
