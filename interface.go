package main

import "fmt"

type Animal interface {
    Eat()
    Move()
    Speak()
}


type Cow struct{
    eat string
    move string
    speak string
}
func (a Cow) Eat(){
    fmt.Println(a.eat)
}

func (a Cow) Move(){
    fmt.Println(a.move)
}

func (a Cow) Speak(){
    fmt.Println(a.speak)
}


type Bird struct{
    eat string
    move string
    speak string
}
func (a Bird) Eat(){
    fmt.Println(a.eat)
}

func (a Bird) Move(){
    fmt.Println(a.move)
}

func (a Bird) Speak(){
    fmt.Println(a.speak)
}

type Snake struct{
    eat string
    move string
    speak string
}
func (a Snake) Eat(){
    fmt.Println(a.eat)
}

func (a Snake) Move(){
    fmt.Println(a.move)
}

func (a Snake) Speak(){
    fmt.Println(a.speak)
}


func main() {
    bd := make(map[string]Animal)
    var command,animal,action string
    var a Animal
    for {
        fmt.Print(">")
        fmt.Scan(&command, &animal, &action)

        if command == "query" {
            a = bd[animal]
            switch action {
                case "eat":
                    a.Eat()
                case "move":
                    a.Move()
                case "speak":
                    a.Speak()
            }
        }else if command == "newanimal"{
            var anim Animal
            switch action{
                case "cow":
                    anim = Cow{eat: "grass", move: "walk", speak: "moo"}
                case "bird":
                    anim = Bird{eat: "worms", move: "fly", speak: "peep"}
                case "snake":
                    anim = Snake{eat:"mice", move: "slither", speak:"hsss"}
            }
            bd[animal]=anim
            fmt.Println("Created it!")
        }
    }
}
