package main

import (
	"encoding/json"
	"net/http"
	"time"

	"encoding/base64"

	"github.com/dgrijalva/jwt-go"
	"github.com/labstack/echo"
)

// User ...
type User struct {
	IsAdmin bool
	OwnerOf []int64
}

var defaultUser = User{IsAdmin: false, OwnerOf: []int64{}}

func getUser(c echo.Context) User {
	userSmth := c.Get("user")
	if userSmth == nil {
		return defaultUser
	}
	user := userSmth.(*jwt.Token)
	if user == nil {
		return defaultUser
	}
	claims := user.Claims.(jwt.MapClaims)
	if claims == nil {
		return defaultUser
	}
	ownerOf := claims["ownerOf"].([]int64)
	isAdmin := claims["admin"].(bool)
	return User{IsAdmin: isAdmin, OwnerOf: ownerOf}
}

func updateUser(c echo.Context, u User) {
	token := generateToken(u.IsAdmin, u.OwnerOf)
	expiration := time.Now().Add(365 * 24 * time.Hour)
	c.SetCookie(&http.Cookie{
		Name:    "token",
		Value:   token,
		Expires: expiration,
	})
	userJSON, _ := json.Marshal(u)
	userBase := base64.StdEncoding.EncodeToString(userJSON)
	c.SetCookie(&http.Cookie{
		Name:    "user",
		Expires: expiration,
		Value:   userBase,
	})
}
