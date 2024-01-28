package main

import (
	"encoding/json"
	"fmt"
	"os"
	"testing"
)

func TestJsonDecode(t *testing.T) {
	reader, _ := os.Open("sample.json")
	decoder := json.NewDecoder(reader)

	post := &Post{}
	decoder.Decode(post)

	fmt.Println(post)
}
