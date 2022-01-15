package main

import (
    "fmt"
)

func hi(value string){
    fmt.Println(value)
}

func main(){
    var wait string

    for i:=0; i<10;i++{
        go hi("hello")
        go hi("bye")
    }

    fmt.Scan(&wait)

}

/*
The pupose of this program is to show the race condition
where the OS decides the order of execution
bye
hello
hello
bye
bye
hello
bye
hello
hello
bye
bye
hello
bye
hello
bye
hello
bye
bye
hello
hello
*/
