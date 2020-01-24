package main

import (
	"fmt"
	"github.com/syronz/goAES"
)

func main() {
	aes := goAES.New().Key("secret key").IV("iv is 16 char!!!").Build()
	encrypted := aes.Encrypt("encrypt here")

	fmt.Println("hello diako", encrypted)
}
