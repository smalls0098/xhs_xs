package xhs

import (
	"bytes"
	"crypto/aes"
	"crypto/cipher"
)

func aesCbcDecrypt(encrypted, key, iv []byte) []byte {
	block, _ := aes.NewCipher(key)
	mode := cipher.NewCBCDecrypter(block, iv)
	mode.CryptBlocks(encrypted, encrypted)
	decrypted := pkcs7UnPadding(encrypted)
	return decrypted
}

func aesCbcEncrypt(data, key, iv []byte) []byte {
	block, _ := aes.NewCipher(key)
	paddText := pkcs7Padding(data, block.BlockSize())
	blockMode := cipher.NewCBCEncrypter(block, iv)
	result := make([]byte, len(paddText))
	blockMode.CryptBlocks(result, paddText)
	return result
}

func pkcs7UnPadding(src []byte) []byte {
	length := len(src)
	unPadding := int(src[length-1])
	return src[:(length - unPadding)]
}

func pkcs7Padding(src []byte, blockSize int) []byte {
	padding := blockSize - len(src)%blockSize
	padText := bytes.Repeat([]byte{byte(padding)}, padding)
	return append(src, padText...)
}
