package main

import "github.com/labstack/echo"

func main() {
	initDB()
	defer db.Close()
	e := echo.New()
	e.GET("/entries", getEntries)
	e.POST("/submit", submitEntry)
	e.Static("/", "static")
	e.Debug = true
	e.Start(":6004")
}
