package main

import "fmt"

type IPerson interface {

	talk()
	walk()
	eat()

}

type Human struct {
}

type Martian struct {
}

func (Human) talk() {
	fmt.Println("talking")
}

func (Human) walk() {
	fmt.Println("walking")
}

func (Human) eat() {
	fmt.Println("eating sushi")
}

func (Martian) walk() {
	fmt.Println("walking")
}

func (Martian) talk() {
	fmt.Println("walking")
}



func (Martian) eat() {
	fmt.Println("eating moon rocks")
}

func invite2Party(partygoers ...IPerson) {
	for _, person := range partygoers {
		person.eat()
	}
}

func main() {
	var Bob Human
	var Sally Human
	var Jingle Martian
	var Susan Human

	invite2Party(Bob, Sally, Jingle, Susan)
}