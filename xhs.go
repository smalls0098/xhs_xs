package xhs

import (
	"encoding/base64"
	"encoding/hex"
	"fmt"
)

func Payload(str string) string {
	key := []byte("fn7cxhamzqet4ltw")
	iv := []byte("3w5zacaub8dqv9zq")
	enc := aesCbcEncrypt([]byte(base64.StdEncoding.EncodeToString([]byte(str))), key, iv)
	return hex.EncodeToString(enc)
}
