package main

import (
	"fmt"
	"io/ioutil"
	"math/rand"
)

const secretFile = ".secret"

var secret []byte

func generateSecret() {
	if bytes, err := ioutil.ReadFile(secretFile); err == nil {
		secret = bytes
	} else {
		fmt.Println("Could not read secret bytes from file, generating new ones")
		secret = make([]byte, 32)
		rand.Read(secret)
		if err := ioutil.WriteFile(secretFile, secret, 0664); err != nil {
			fmt.Println("Could not save secret bytes to file")
		}
	}
}
