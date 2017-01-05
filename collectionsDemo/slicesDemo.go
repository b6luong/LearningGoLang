package main

import (
    "fmt"
    "sort"
)

func main() {
    // Slices are dynamic arrays, they don't have a fixed length
    var colors = []string{"Red", "Green", "Blue"}
    fmt.Println(colors)
    
    // Add to the end of the array
    colors = append(colors, "Purple")
    fmt.Println(colors)

    // Remove the first element from the array
    colors = append(colors[1:]) // append(colors[1:len(colors)]) will do the same thing
    fmt.Println(colors)
    
    // Remove the last element from the array
    colors = append(colors[:len(colors) - 1]) 
    fmt.Println(colors)

    // Creating a slice with initial size and capacity
    numbers := make([]int, 5, 5)
    numbers[0] = 1
    numbers[1] = 796
    numbers[2] = 32
    numbers[3] = 12
    numbers[4] = 156
    
    fmt.Println(numbers)
    
    numbers = append(numbers, 235)
    fmt.Println(numbers)
    fmt.Println(cap(numbers)) //OUTPUT: 10
    
    // Sorting
    sort.Ints(numbers)
    fmt.Println(numbers)
}