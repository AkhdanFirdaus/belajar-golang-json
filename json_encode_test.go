package main

import (
	"encoding/json"
	"os"
	"testing"
)

func TestEncoder(t *testing.T) {
	writer, _ := os.Create("customer_output.json")
	encoder := json.NewEncoder(writer)

	post := Post{
		UserId: 1,
		Id:     1,
		Title:  "Ini Judul",
		Body:   "Ini Body",
	}

	encoder.Encode(post)
}
