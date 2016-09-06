package main

import (
	"encoding/json"
	"fmt"
	"io"
	"log"
	"strings"
)

type Record struct {
	Author Author `json:"author"`
	Title  string `json:"title"`
	URL    string `json:"url"`
}

type Author struct {
	ID    uint64 `json:"id"`
	Email string `json:"email"`
}

// Used to avoid recursion in UnmarshalJSON below.
type author Author

func (a *Author) UnmarshalJSON(b []byte) (err error) {
	j, s, n := author{}, "", uint64(0)
	if err = json.Unmarshal(b, &j); err == nil {
		*a = Author(j)
		return
	}
	if err = json.Unmarshal(b, &s); err == nil {
		a.Email = s
		return
	}
	if err = json.Unmarshal(b, &n); err == nil {
		a.ID = n
	}
	return
}

func Decode(r io.Reader) (x *Record, err error) {
	x = new(Record)
	err = json.NewDecoder(r).Decode(x)
	return x, err
}

type Reader interface {
	Read(p []byte) (n int, err error)
}

func main() {
	const jsonStream = `
		[
			{"Name": "Ed", "Text": "Knock knock."},
			{"Name": "Sam", "Text": "Who's there?"},
			{"Name": "Ed", "Text": "Go fmt."},
			{"Name": "Sam", "Text": "Go fmt who?"},
			{"Name": "Ed", "Text": "Go fmt yourself!"}
		]
	`
	type Message struct {
		Name, Text string
	}
	dec := json.NewDecoder(strings.NewReader(jsonStream))

	// read open bracket
	t, err := dec.Token()
	if err != nil {
		log.Fatal(err)
	}
	fmt.Printf("%T: %v\n", t, t)

	// while the array contains values
	for dec.More() {
		var m Message
		// decode an array value (Message)
		err := dec.Decode(&m)
		if err != nil {
			log.Fatal(err)
		}

		fmt.Printf("%v: %v\n", m.Name, m.Text)
	}

	// read closing bracket
	t, err = dec.Token()
	if err != nil {
		log.Fatal(err)
	}
	fmt.Printf("%T: %v\n", t, t)
}
