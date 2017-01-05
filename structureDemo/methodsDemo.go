package main

import (
    "fmt"
)

type Dog struct {
    Breed string
    Weight int
    Sound string
}

// This function is an example of pass by value
func (d Dog) Speak() {
    fmt.Println(d.Sound)
}

//This function will make changes to the calling object since it takes the pointer of the object.
func (d *Dog) SpeakThreeTimes(){
    d.Sound = fmt.Sprintf("%v! %v! %v!", d.Sound, d.Sound, d.Sound)
    fmt.Println(d.Sound)
}

func main() {
    pug := Dog{"Pug", 40, "Woof"}
    fmt.Println(pug)
    pug.Speak();
    
    pug.Sound = "Ark"
    pug.Speak();
    
    pug.SpeakThreeTimes();
    pug.SpeakThreeTimes();
    
}