package main

import (
	"encoding/json"
	"fmt"
	"testing"
)

func TestMap(t *testing.T) {
	jsonString := `{"userId": 1,"id": 1,"title": "sunt aut facere repellat provident occaecati excepturi optio reprehenderit","body": "quia et suscipit\nsuscipit recusandae consequuntur expedita et cum\nreprehenderit molestiae ut ut quas totam\nnostrum rerum est autem sunt rem eveniet architecto"}`
	jsonBytes := []byte(jsonString)

	var result map[string]interface{}
	json.Unmarshal(jsonBytes, &result)
	fmt.Println(result)
	fmt.Println(result["userId"])
	fmt.Println(result["id"])
	fmt.Println(result["post"])
	fmt.Println(result["body"])
}

func TestMapEncode(t *testing.T) {
	post := map[string]interface{}{
		"userId": 1,
		"id":     1,
		"title":  "sunt aut facere repellat provident occaecati excepturi optio reprehenderit",
		"body":   "quia et suscipit\nsuscipit recusandae consequuntur expedita et cum\nreprehenderit molestiae ut ut quas totam\nnostrum rerum est autem sunt rem eveniet architecto",
	}

	bytes, _ := json.Marshal(post)
	fmt.Println(string(bytes))
}
