package main

import (
	"fmt"

	"github.com/pablitovicente/interfaces_002/pkg/animals"
)

type Movable interface {
	Move() string
}

func main() {
	animals := SetupAnimals()
	MoveAll(animals)
}

func MoveAll(m []Movable) {
	for _, animal := range m {
		fmt.Println(animal.Move())
	}

}

func SetupAnimals() []Movable {
	aCat := &animals.Cat{
		Name: "Top Cat",
	}

	aDog := &animals.Dog{
		Name: "Scooby",
	}

	aKangaroo := &animals.Kangaroo{
		Name: "Oi Mate",
	}

	anEagle := &animals.Eagle{
		Name: "Angry Bird",
	}

	anOrca := &animals.Orca{
		Name: "Willy",
	}

	return []Movable{aCat, aDog, aKangaroo, anEagle, anOrca}
}
