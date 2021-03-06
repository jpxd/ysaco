package main

import (
	"github.com/labstack/echo"
)

func main() {
	initDB()
	defer db.Close()
	generateSecret()
	e := echo.New()
	e.GET("/entries", getEntries)
	e.POST("/submit", submitEntry)
	e.POST("/delete", deleteEntry)
	e.GET("/g3tr00t", getRoot)
	e.Static("/", "static")
	e.Debug = true
	e.Start(":6004")
}
