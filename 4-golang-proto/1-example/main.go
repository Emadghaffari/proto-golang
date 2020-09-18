package main

import (
	"fmt"

	Tutorial "github.com/Emadghaffari/proto-golang/4-golang-proto/1-example/tutorial"
)

func main() {
	p1 := Tutorial.Person{
		Id:          1,
		FirstName:   "Emad",
		IsValidated: true,
		SomeArray:   []string{"Tag1", "Tag2"},
		Color:       Tutorial.Person_RED,
		Date:        &Tutorial.Person_MyDate{Year: 1996, Month: 12, Day: 26},
	}
	fmt.Println(p1.GetId())
	fmt.Println(p1.GetFirstName())
}
