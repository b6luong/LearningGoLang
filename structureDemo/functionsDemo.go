package main

import (
    "fmt"
)

func main() {
    doSomething()
    
    sum := addValues(23, 54)
    fmt.Println("sum:", sum)
    
    sum = addAllValues(12, 54, 79)
    fmt.Println("sum:", sum)
}


// No function overloading!!! Cannot have 2 functions with the same name and same
// number of arugment. Also, functions that start with a lowercase are set to private
// funtions that start with an uppercase are set to public and can be exported
func doSomething(){
    fmt.Println("Doing Something!")
}

func addValues(v1 int, v2 int) int {
    return v1 + v2;
}

func addAllValues(values ...int) int {
    sum := 0
    for i := range values{
        sum += values[i]
    }
    fmt.Printf("%T\n", values)
    return sum;
}