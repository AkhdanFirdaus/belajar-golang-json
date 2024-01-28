package main

import (
	"encoding/json"
	"fmt"
	"testing"
)

func TestJSONArray(t *testing.T) {
	customer := Customer{
		FirstName:  "Akhdan",
		MiddleName: "Musyaffa",
		LastName:   "Firdaus",
		Hobbies:    []string{"Sleeping", "Doing Sport", "Writing"},
	}

	bytes, _ := json.Marshal(customer)
	fmt.Println(string(bytes))
}

func TestJSONArrayDecode(t *testing.T) {
	jsonString := `{"FirstName":"Akhdan","MiddleName":"Musyaffa","LastName":"Firdaus","Hobbies":["Sleeping","Doing Sport","Writing"]}`
	jsonBytes := []byte(jsonString)

	customer := &Customer{}
	err := json.Unmarshal(jsonBytes, customer)
	if err != nil {
		panic(err)
	}

	fmt.Println(customer)
	fmt.Println(customer.FirstName)
	fmt.Println(customer.MiddleName)
	fmt.Println(customer.LastName)
	fmt.Println(customer.Hobbies)
}

func TestOnlyJSONArray(t *testing.T) {
	customers := []Customer{
		{
			FirstName:  "Akhdan",
			MiddleName: "Musyaffa",
			LastName:   "Firdaus",
		},
		{
			FirstName:  "Test",
			MiddleName: "123",
			LastName:   "321",
		},
		{
			FirstName:  "Set",
			MiddleName: "Est",
			LastName:   "Tes",
		},
	}

	bytes, _ := json.Marshal(customers)
	fmt.Println(string(bytes))
}

func TestOnlyJSONArrayDecode(t *testing.T) {
	jsonString := `[{"FirstName":"Akhdan","MiddleName":"Musyaffa","LastName":"Firdaus","Hobbies":null},{"FirstName":"Test","MiddleName":"123","LastName":"321","Hobbies":null},{"FirstName":"Set","MiddleName":"Est","LastName":"Tes","Hobbies":null}]`
	jsonBytes := []byte(jsonString)

	customers := &[]Customer{}

	err := json.Unmarshal(jsonBytes, customers)
	if err != nil {
		panic(err)
	}

	fmt.Println(customers)
}
