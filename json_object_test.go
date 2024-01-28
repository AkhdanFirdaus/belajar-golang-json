package main

import (
	"encoding/json"
	"fmt"
	"testing"
)

type Customer struct {
	FirstName  string
	MiddleName string
	LastName   string
	age        int
	isMarried  bool
	Hobbies    []string
}

func TestJSONObject(t *testing.T) {
	customer := Customer{
		FirstName:  "Akhdan",
		MiddleName: "Musyaffa",
		LastName:   "Firdaus",
		age:        22,
		isMarried:  false,
	}

	bytes, _ := json.Marshal(customer)
	fmt.Println(string(bytes))
}
