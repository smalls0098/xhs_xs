package xhs

import (
	"encoding/base64"
	"encoding/hex"
	"fmt"
)

func Payload(str string) string {
	key := []byte("0a9fw3m6qfll2dej")
	iv := []byte("mhaqhnjmr0rsoo3o")
	enc := aesCbcEncrypt([]byte(base64.StdEncoding.EncodeToString([]byte(str))), key, iv)
	return hex.EncodeToString(enc)
}
