package main

import (
	"fmt"
	"net/http"
	"time"

	"github.com/dgrijalva/jwt-go"
	"github.com/labstack/echo"
)

// User ...
type User struct {
	IsAdmin bool
	OwnerOf []int64
	jwt.StandardClaims
}

const cookieName = "token"

var signingMethod = jwt.SigningMethodHS512

var defaultUser = User{IsAdmin: false, OwnerOf: []int64{}}
var adminUser = User{IsAdmin: true, OwnerOf: []int64{}}

func getUser(c echo.Context) User {
	var user User
	cookie, err := c.Cookie(cookieName)
	if err != nil {
		return defaultUser
	}
	token, err := jwt.ParseWithClaims(cookie.Value, &user, func(t *jwt.Token) (interface{}, error) {
		if t.Method.Alg() != signingMethod.Alg() {
			return nil, fmt.Errorf("Unexpected jwt signing method=%v", t.Header["alg"])
		}
		return secret, nil
	})
	if err != nil || !token.Valid {
		return defaultUser
	}
	return user
}

func updateUser(c echo.Context, user *User) {
	token := jwt.NewWithClaims(signingMethod, user)
	tokenstring, err := token.SignedString(secret)
	if err != nil {
		fmt.Println(err)
		return
	}
	expiration := time.Now().Add(365 * 24 * time.Hour)
	c.SetCookie(&http.Cookie{
		Name:    cookieName,
		Value:   tokenstring,
		Expires: expiration,
	})
}
