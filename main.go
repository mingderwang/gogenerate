package main

//go:generate joiner $GOFILE

import (
	"fmt"
	"github.com/davecgh/go-spew/spew"
)

// @joiner
type Person struct {
	FirstName string
	LastName  string
	HairColor string
}

func main() {
	people := []Person{
		Person{"Sideshow", "Bob", "red"},
		Person{"Homer", "Simpson", "n/a"},
		Person{"Lisa", "Simpson", "blonde"},
		Person{"Marge", "Simpson", "blue"},
		Person{"Mr", "Burns", "gray"},
	}

	joined_people := JoinPerson(people).With("\n")
	fmt.Printf("My favorite Simpsons Characters:\n%s\n", joined_people)
	spew.Dump(people)
	spew.Dump(joined_people)
}
