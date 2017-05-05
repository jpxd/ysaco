package main

import (
	"github.com/asdine/storm"
	"github.com/asdine/storm/codec/sereal"
)

const dbfile = "xyz.db"

var db *storm.DB

func initDB() {
	db, _ = storm.Open(dbfile, storm.Codec(sereal.Codec))
}

func getSamples() []SampleEntry {
	var samples []SampleEntry
	db.All(&samples)
	return samples
}

func saveSample(sample SampleEntry) {
	db.Save(&sample)
}

func deleteSample(id int64) {
	db.Delete("SampleEntry", id)
}
