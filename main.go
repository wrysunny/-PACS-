package main

import (
	"crypto/cipher"
	"crypto/des"
	"encoding/base64"
	"flag"
	"fmt"
	"os"
)

func main() {
	var (
		iv  = []uint8{18, 52, 86, 120, 144, 171, 205, 239}
		key = "88916830"
		enc string
	)

	flag.StringVar(&enc, "enc", "", "Input encryption text.")
	flag.Parse()
	if enc == "" {
		flag.Usage()
		os.Exit(0)
	}

	v1, _ := base64.StdEncoding.DecodeString(enc)
	block, err := des.NewCipher([]byte(key))
	if err != nil {
		panic(err)
	}
	mode := cipher.NewCBCDecrypter(block, iv)
	plaintext := make([]byte, len(v1))
	mode.CryptBlocks(plaintext, v1)
	plaintext = PKCS5UnPadding(plaintext)
	fmt.Println(string(plaintext))
}

func PKCS5UnPadding(origData []byte) []byte {
	length := len(origData)
	unpadding := int(origData[length-1])
	return origData[:(length - unpadding)]
}
