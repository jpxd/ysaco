package main

import (
	"net/http"

	"regexp"

	"strconv"

	"time"

	"github.com/labstack/echo"
)

var nameValidationRegex = regexp.MustCompile("[a-zA-Z0-9 \\.,\\?\\!]+")
var idValidationRegex = regexp.MustCompile("[a-zA-Z0-9]+")
var numberValidationRegex = regexp.MustCompile("[0-9]+")

func submitEntry(c echo.Context) error {
	name := c.FormValue("name")
	id := c.FormValue("id")
	start := c.FormValue("start")
	duration := c.FormValue("duration")

	if !nameValidationRegex.MatchString(name) ||
		!idValidationRegex.MatchString(id) ||
		!numberValidationRegex.MatchString(start) ||
		!numberValidationRegex.MatchString(duration) {

		return c.NoContent(http.StatusBadRequest)
	}

	startInt, _ := strconv.ParseInt(start, 10, 32)
	durationInt, _ := strconv.ParseInt(duration, 10, 32)

	sample := SampleEntry{
		Name:         name,
		SecondsStart: int32(startInt),
		Duration:     int32(durationInt),
		YoutubeID:    id,
		Timestamp:    time.Now(),
	}

	saveSample(sample)

	return c.NoContent(http.StatusOK)
}

func getEntries(c echo.Context) error {
	return c.JSON(http.StatusOK, getSamples())
}

func deleteEntry(c echo.Context) error {
	id := c.FormValue("id")

	if !numberValidationRegex.MatchString(id) {
		return c.NoContent(http.StatusBadRequest)
	}

	idInt, _ := strconv.ParseInt(id, 10, 64)
	deleteSample(idInt)

	return c.NoContent(http.StatusOK)
}
