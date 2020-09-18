package main

import (
	"fmt"

	PersonDB "github.com/Emadghaffari/proto-golang/3-golang-proto/2-example/tutorial"
)

func main() {
	p1 := PersonDB.Person{
		Id:          1,
		FirstName:   "Emad",
		IsValidated: true,
		SomeArray:   []string{"Tag1", "Tag2"},
		Color:       PersonDB.Person_RED,
		Date:        &PersonDB.Person_MyDate{Year: 1996, Month: 12, Day: 26},
	}
	fmt.Println(p1.GetId())
	fmt.Println(p1.GetFirstName())
}
