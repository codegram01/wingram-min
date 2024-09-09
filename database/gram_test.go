package database

import (
	"github.com/codegram01/wingram-min/gram"
	"testing"
)

func TestCreate(t *testing.T) {
	db := DbTest()

	gram, err := db.GramCreate(&gram.Gram{
		Title:       "Hello",
		Description: "Hello World",
		Content:     "Content here",
	})

	if err != nil {
		t.Fatal(err)
	}

	t.Log(gram)
}

func TestList(t *testing.T) {
	db := DbTest()

	grams, err := db.GramList()

	if err != nil {
		t.Fatal(err)
	}

	t.Log(grams)
}

func TestDetail(t *testing.T) {
	db := DbTest()

	gram, err := db.GramDetail(1)

	if err != nil {
		t.Fatal(err)
	}

	t.Log(gram)
}

func TestDelete(t *testing.T) {
	db := DbTest()

	err := db.GramDelete(1)

	if err != nil {
		t.Fatal(err)
	}

	t.Log("success")
}

func TestUpdate(t *testing.T) {
	db := DbTest()

	gram, err := db.GramUpdate(2, &gram.Gram{
		Title:       "Hello 2",
		Description: "Hello World",
		Content:     "Content here",
	})

	if err != nil {
		t.Fatal(err)
	}

	t.Log(gram)
}
