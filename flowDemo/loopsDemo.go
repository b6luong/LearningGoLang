package main

import "fmt"

func main() {
    sum := 1
    fmt.Println("sum:", sum)
    
    for i := 0; i < 10; i++ {
        sum += i
    }
    fmt.Println("sum:", sum)
    
    colors := []string{"Red", "Green", "Blue"}
    fmt.Println(colors)
    
    for i := 0; i < len(colors); i++ {
        fmt.Println(colors[i])
    }
    
    for i := range colors {
        fmt.Println("i:", i, " Color:", colors[i])

    }
    
    // This is how you make a while loop in Go
    sum = 1
    for sum < 1000 {
        sum += sum
        fmt.Println("sum:", sum)
        if sum > 200 {
            goto endofprogram   // Going to a label
        }
        if sum > 500 {
            break
        }
    }
    
    // adding label
    endofprogram : fmt.Println("end of program") 
    
}