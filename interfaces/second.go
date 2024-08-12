package interfaces

import "fmt"

type Animal interface {
	Speak() string
}

type Dog struct {}

func (d Dog) Speak() string {
	return "wooof woof"
}

type Cat struct {}

func (c Cat) Speak() string {
	return "meeaw meaw"
}

type Duck struct {}

func (d Duck) Speak() string {
	return "quack quack"
}

func AnimalSpeaker() {
	animals := []Animal{Dog{}, Cat{}, Duck{}}

	for _, animal := range animals {
		fmt.Println(animal.Speak())
	}
}
