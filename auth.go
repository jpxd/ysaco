package main

import (
	"fmt"
	"io/ioutil"
	"math/rand"

	"github.com/dgrijalva/jwt-go"
	"github.com/labstack/echo"
	"github.com/labstack/echo/middleware"
)

const secretFile = ".secret"

var secretCache []byte

func secretBytes() []byte {
	if secretCache != nil {
		return secretCache
	}
	if bytes, err := ioutil.ReadFile(secretFile); err != nil {
		return bytes
	}
	fmt.Println("Could not read secret bytes from file, generating new ones")
	secretCache = make([]byte, 32)
	rand.Read(secretCache)
	if err := ioutil.WriteFile(secretFile, secretCache, 0664); err != nil {
		fmt.Println("Could not save secret bytes to file")
	}
	return secretCache
}

func generateToken(admin bool, ownerOf []int64) string {
	token := jwt.New(jwt.SigningMethodHS256)

	claims := token.Claims.(jwt.MapClaims)
	claims["admin"] = admin
	claims["ownerOf"] = ownerOf

	t, err := token.SignedString(secretBytes())
	if err != nil {
		fmt.Println("Could not sign token")
		panic(err)
	}
	return t
}

func skipWithoutToken(c echo.Context) bool {
	cookie, err := c.Cookie("token")
	return err != nil || len(cookie.Value) < 10
}

func jwtAuth() echo.MiddlewareFunc {
	c := middleware.JWTConfig{
		Skipper:       skipWithoutToken,
		SigningMethod: middleware.AlgorithmHS256,
		ContextKey:    "user",
		TokenLookup:   "cookie:token",
		Claims:        jwt.MapClaims{},
	}
	c.SigningKey = secretBytes()
	return middleware.JWTWithConfig(c)
}
