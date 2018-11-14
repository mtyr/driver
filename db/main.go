package main

import (
	"encoding/json"
	"fmt"

	"github.com/nanobox-io/golang-scribble"
)

type Article struct {
	Title string `json:"title"`
	Body  string `json:"body"`
}

func main() {
	db, err := scribble.New("./db", nil)
	if err != nil {
		fmt.Println("Error", err)
	}
	db.Write("test", "2", Article{"title", "body"})
	db.Write("test", "3", Article{"title", "body"})

	records, err := db.ReadAll("test")
	if err != nil {
		fmt.Println("Error", err)
	}

	fishies := []Article{}
	for _, f := range records {
		fishFound := Article{}
		if err := json.Unmarshal([]byte(f), &fishFound); err != nil {
			fmt.Println("Error", err)
		}
		fishies = append(fishies, fishFound)
	}
	fmt.Println(fishies)
	//var article Article
	//err = db.Read("test", "1", &article)
	//if err != nil {
	//	panic(err)
	//}
	//fmt.Println(article.Title)
}
