package domain

import "fmt"

type Person struct {
	Id       string
	Name     string
	Age      int
	Type     string
	document string
}

func (person Person) InsertPerson() {
	fmt.Println(person)
}
