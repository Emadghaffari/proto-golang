package main

import (
	"fmt"
	"io/ioutil"
	"log"

	PersonDB "github.com/Emadghaffari/proto-golang/3-golang-proto/4-example/tutorial"
	"github.com/golang/protobuf/jsonpb"
	"google.golang.org/protobuf/proto"
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
	// write("out.bin", &p1)

	js := toJSON(p1)
	p2 := &PersonDB.Person{}
	fromJSON("out.bin", p2)

	fmt.Println(js)
	fmt.Println(p2)
}
func toJSON(pm PersonDB.Person) string {
	marshaler := jsonpb.Marshaler{}
	out, err := marshaler.MarshalToString(&pm)
	if err != nil {
		log.Fatalln(err)
	}
	return out
}

func fromJSON(in string, pm *PersonDB.Person) {
	err := jsonpb.UnmarshalString(in, pm)
	if err != nil {
		log.Fatalln(err)
	}
}

func write(fname string, pm proto.Message) {
	out, err := proto.Marshal(pm)
	if err != nil {
		log.Fatalln(err)
	}
	if err = ioutil.WriteFile(fname, out, 0644); err != nil {
		log.Fatalln(err)
	}
}

func read(fname string, pm proto.Message) {
	in, err := ioutil.ReadFile(fname)
	if err != nil {
		log.Fatalln(err)
	}
	err = proto.Unmarshal(in, pm)
	if err != nil {
		log.Fatalln(err)
	}
}
