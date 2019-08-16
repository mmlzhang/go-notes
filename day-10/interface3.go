package main

import (
	"fmt"
)

type Animal interface{
	Speak() string
}


type Dog struct {

}

func (d Dog) Speak() string {
	return "woof!"
}


type Cat struct {}

func (c Cat) Speak() string {
	return "Momn"
}


type Llama struct {

}

func (l Llama) Speak() string {
	return "??????"
}


type JavaProgrammer struct{}

func (j JavaProgrammer) Speak() string {
	return "Desgin pattern"
}


func main () {
	animals := []Animal{Dog{}, Cat{}, Llama{}, JavaProgrammer{}}

	for _, animal := range animals {
		fmt.Println(animal.Speak())
	}
}

