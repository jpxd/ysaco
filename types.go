package main

import "time"

// SampleEntry holds all data for an entry
type SampleEntry struct {
	ID           int64 `storm:"id,increment"`
	Name         string
	YoutubeID    string
	SecondsStart int32
	Duration     int32
	Timestamp    time.Time
}
