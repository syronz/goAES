# goAES
Simple package for AES encryption and decryption. It use sha256 for hash the secret key.
Based on CTR

[![BuildStatus](https://api.travis-ci.org/syronz/goAES.svg?branch=master)](http://travis-ci.org/syronz/goAES)
[![ReportCard](https://goreportcard.com/badge/github.com/syronz/goAES)](https://goreportcard.com/report/github.com/syronz/goAES) 
[![Go Coverage](https://github.com/syronz/goAES/blob/master/coverage_badge.png)](https://gocover.io/github.com/syronz/goAES)
[![Codacy Badge](https://api.codacy.com/project/badge/Grade/2a4d66750d5047e58742349f2dfc8c8d)](https://www.codacy.com/manual/syronz/goAES?utm_source=github.com&amp;utm_medium=referral&amp;utm_content=syronz/goAES&amp;utm_campaign=Badge_Grade)
[![codebeat badge](https://codebeat.co/badges/a882cf1a-cc00-4690-b65f-e69fb74cf574)](https://codebeat.co/projects/github-com-syronz-goaes-master)
[![Maintainability](https://api.codeclimate.com/v1/badges/382db50d589346613f15/maintainability)](https://codeclimate.com/github/syronz/goAES/maintainability)
[![GolangCI](https://golangci.com/badges/github.com/gojek/darkroom.svg)](https://golangci.com/r/github.com/syronz/goAES)
[![Open Source Helpers](https://www.codetriage.com/syronz/goaes/badges/users.svg)](https://www.codetriage.com/syronz/goaes)

[![GoDoc](https://godoc.org/github.com/syronz/go-log?status.svg)](https://godoc.org/github.com/syronz/goAES)
[![FOSSA Status](https://app.fossa.io/api/projects/git%2Bgithub.com%2Fsyronz%2FgoAES.svg?type=shield)](https://app.fossa.io/projects/git%2Bgithub.com%2Fsyronz%2FgoAES?ref=badge_shield)
[![MIT license](https://img.shields.io/badge/license-MIT-brightgreen.svg)](https://github.com/syronz/goAES/blob/master/LICENSE)


## Usage

### Encryption

```go
package main

import (
	"fmt"
	"log"

	goaes "github.com/syronz/goAES"
)

func main() {
	aes, err := goaes.New().Key("secret key").IV("iv is 16 char!!!").Build()
	if err != nil {
		log.Fatal(err)
	}

	term := "Hello"
	encrypted := aes.Encrypt(term)

	fmt.Printf("Encryption of %q is %q \n", term, encrypted)
}

```

output:
```shell
Encryption of "Hello" is "257d85d55b" 
```

### Decryption

Default length is 100 character, in case the orginal term is more than 100 in AES decryption, you shold mention the Length in the build.
```go
package main

import (
	"fmt"
	"log"

	goaes "github.com/syronz/goAES"
)

func main() {
	aes, err := goaes.New().Key("secret key").IV("iv is 16 char!!!").Build()
	if err != nil {
		log.Fatal(err)
	}

	encrypted := "257d85d55b"
	decrypted := aes.Decrypt(encrypted)

	fmt.Printf("Decryption of %q is %q \n", encrypted, decrypted)
}
```

output:
```shell
Decryption of "257d85d55b" is "Hello" 
```

### Note

In case of the length of decrypted word is more than 100, we should use **.Length(int)** like here. Extra chars will be trimmed
```go
aes, err := goaes.New().Key("secret key").IV("iv is 16 char!!!").Length(10000).Build()
```

### Online test
By going to the [gchq.github.io/CyberChef](https://gchq.github.io/CyberChef/#recipe=AES_Encrypt(%7B'option':'Hex','string':'e3b0c44298fc1c149afbf4c8996fb92427ae41e4649b934ca495991b7852b855'%7D,%7B'option':'Latin1','string':'1234567890123456'%7D,'CTR','Raw','Hex')AES_Decrypt(%7B'option':'Hex','string':'e3b0c44298fc1c149afbf4c8996fb92427ae41e4649b934ca495991b7852b855'%7D,%7B'option':'Latin1','string':'1234567890123456'%7D,'CTR','Hex','Raw',%7B'option':'Latin1','string':''%7D/breakpoint)&input=QUJDREVGR0hJSktMTU5PUFFSU1RVVldYWVo) you can test it manually 

## License
[![FOSSA Status](https://app.fossa.io/api/projects/git%2Bgithub.com%2Fsyronz%2FgoAES.svg?type=large)](https://app.fossa.io/projects/git%2Bgithub.com%2Fsyronz%2FgoAES?ref=badge_large)
