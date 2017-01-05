package main

import (
    "fmt"
)

type Animal interface {
    Speak() string
}

type Dog struct {}

func (d Dog) Speak() string {
    return "Woof!"
}

type Cat struct {}

func (c Cat) Speak() string {
    return "Meow!"
}

type Sheep struct {}

func (s Sheep) Speak() string {
    return "Baah!"
}



// As long as the struct implements all the methods in the interface, then it is 
// using that interface. No need to say "implements"
func main() {
    corgi := Animal(Dog{})
    fmt.Println(corgi)
    
    garfield := Cat{}
    var animals []Animal = []Animal{corgi, garfield, Sheep{}}
    for _, animal := range animals {
        fmt.Println(animal.Speak())
    }
    
}