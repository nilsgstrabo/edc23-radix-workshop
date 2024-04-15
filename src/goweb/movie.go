package main

import "time"

type Movie struct {
	Title    string    `json:"title"`
	Released time.Time `json:"released,omitempty"`
	Rating   float32   `json:"rating,omitempty"`
}
