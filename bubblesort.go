package main

import ("fmt"
    )

func Swap(s []int, i int){
    s[i],s[i+1]=s[i+1],s[i]
}

func BubbleSort(sli []int){
    n := len(sli)
    for i:=1; i<n; i++ {
        for j:=0; j<n-1; j++{
            if sli[j]>sli[j+1] {
                Swap(sli,j)
            }
        }
    }
}

func main() {
    var value int
    var arr []int
    for i:=0;i<10;i++{
        fmt.Scanln(&value)
        arr = append(arr, value)
    }
    BubbleSort(arr)
    fmt.Println(arr)
}
