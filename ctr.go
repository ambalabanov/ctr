package ctr

import (
	"crypto/aes"
	"crypto/cipher"
)

//Crypt AES-CTR-256
func Crypt(data, key, iv []byte) (cdata []byte) {

	block, err := aes.NewCipher(key)
	if err != nil {
		panic(err)
	}
	cdata = make([]byte, len(data))
	stream := cipher.NewCTR(block, iv)
	stream.XORKeyStream(cdata, data)
	return cdata
}
