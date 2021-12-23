package main

import "fmt"

type Animal struct {
    food string
    locomotion string
    noise string
}

func (a *Animal) Eat(){
    fmt.Println("eat:", a.food)
}

func (a *Animal) Move(){
    fmt.Println("locomotion:", a.locomotion)
}

func (a *Animal) Speak(){
    fmt.Println("noise:", a.noise)
}

func main() {
    var animal string
    var action string
    var a Animal
    for {
        fmt.Print(">")
        fmt.Scan(&animal)
        if animal == "cow" {
            a = Animal{"grass", "walk", "moo"}
		} else if animal == "bird" {
			a = Animal{"worms", "fly", "peep"}
		} else if animal == "snake" {
			a = Animal{"mice", "slither", "hsss"}
        } 

        fmt.Print(">")
        fmt.Scan(&action)
        if action == "eat" {
			a.Eat()
		} else if action == "move" {
			a.Move()
		} else if action == "speak" {
			a.Speak()
        }
    }
}
