package api

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestBookToJson(t *testing.T) {
	book := Book{Title: "h", Author: "micel", ISBN: "111"}
	json := book.ToJson()

	assert.Equal(t, `{"title":"h","author":"micel","isbn":"111"}`, string(json), "Error")
}

func TestBookFromJson(t *testing.T) {
	book := fromJSON([]byte(`{"title":"h","author":"micel","isbn":"111"}`))

	assert.Equal(t, Book{Title: "h", Author: "micel", ISBN: "111"}, book, "Error")
}
