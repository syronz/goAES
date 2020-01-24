package main

import (
	"fmt"

	goaes "github.com/syronz/goAES"
)

func main() {
	secretKey := "Hello"
	hashed := goaes.New().Key(secretKey).GetKey()

	fmt.Printf("hashed of %q is %q \n", secretKey, hashed)
}
