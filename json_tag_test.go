package main

import (
	"encoding/json"
	"fmt"
	"testing"
)

type Post struct {
	UserId int64  `json:"userId"`
	Id     int64  `json:"id"`
	Title  string `json:"title"`
	Body   string `json:"body"`
}

func TestJsonTag(t *testing.T) {
	post := Post{
		UserId: 1,
		Id:     1,
		Title:  "Ini Judul",
		Body:   "Ini Body",
	}

	bytes, _ := json.Marshal(post)
	fmt.Println(string(bytes))
}

func TestJsonTagDecode(t *testing.T) {
	jsonString := `{"userId": 1,"id": 1,"title": "sunt aut facere repellat provident occaecati excepturi optio reprehenderit","body": "quia et suscipit\nsuscipit recusandae consequuntur expedita et cum\nreprehenderit molestiae ut ut quas totam\nnostrum rerum est autem sunt rem eveniet architecto"}`
	jsonBytes := []byte(jsonString)

	post := &Post{}
	json.Unmarshal(jsonBytes, post)
	fmt.Println(post)
}
