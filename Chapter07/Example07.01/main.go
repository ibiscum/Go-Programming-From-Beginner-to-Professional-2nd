package main

import (
	"encoding/json"
	"fmt"
	"io"
	"strings"
)

type Person struct {
	Name string `json:"name"`
	Age  int    `json:"age"`
}

func main() {
	s1 := `{"Name":"Joe","Age":18}`
	p1, err := loadPerson1(strings.NewReader(s1))
	if err != nil {
		fmt.Println(err)
	}
	fmt.Println(p1)

	s2 := `{"Name":"Jane","Age":21}`
	p2, err := loadPerson2(s2)
	if err != nil {
		fmt.Println(err)
	}
	fmt.Println(p2)
}

func loadPerson1(r io.Reader) (Person, error) {
	var p Person
	err := json.NewDecoder(r).Decode(&p)
	if err != nil {
		return p, err
	}
	return p, err
}

func loadPerson2(s string) (Person, error) {
	var p Person
	err := json.NewDecoder(strings.NewReader(s)).Decode(&p)
	if err != nil {
		return p, err
	}
	return p, nil
}
