package api

import (
	"encoding/json"
	"net/http"
)

type Book struct {
	Title  string `json:"title"`
	Author string `json:"author"`
	ISBN   string `json:"isbn"`
}

func (b Book) ToJson() []byte {
	ToJson, err := json.Marshal(b)
	if err != nil {
		panic(err)
	}
	return ToJson
}

func fromJSON(data []byte) Book {
	book := Book{}
	err := json.Unmarshal(data, &book)

	if err != nil {
		panic(err)
	}
	return book
}

var Books = []Book{
	{Title: "a", Author: "b", ISBN: "123"},
	{Title: "b", Author: "c", ISBN: "345"},
}

func BookHandlerFunc(w http.ResponseWriter, r *http.Request) {
	b, err := json.Marshal(Books)

	if err != nil {
		panic(err)
	}

	w.Header().Add("Content-type", "application/json; charset-utf-8")
	w.Write(b)
}
