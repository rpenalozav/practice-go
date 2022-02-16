package main

import (
	"fmt"
	"sort"
)

func MySort(l []int, c chan []int) {
	sort.Ints(l)
	c <- l
}

func main() {
	c := make(chan []int)
	num := make([]int, 0)
	aux := make([]int, 0)
	var a, count int
	fmt.Println("Enter amount of numbers: ")
	fmt.Scan(&count)
	if count < 4 {
		panic("min 4 numbers")
	}

	for i := 0; i < count; i++ {
		fmt.Scan(&a)
		num = append(num, a)
	}
	size := len(num) / 4
	go MySort(num[:size], c)
	go MySort(num[size:size*2], c)
	go MySort(num[size*2:size*3], c)
	go MySort(num[size*3:], c)

	for i := 0; i < 4; i++ {
		x := <-c
		aux = append(aux, x...)
	}
	sort.Ints(aux)
	fmt.Println(aux)
}
