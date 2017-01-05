package main

import (
    "fmt"
    "math/rand"
    "time"
)

func main() {
    rand.Seed(time.Now().Unix())
    //dow := rand.Intn(6) + 1 // value from 1 - 7
    //fmt.Println("Day", dow)
    
    result := ""
    
    // default in switch, there is no need to use break
    switch dow := rand.Intn(6) + 1 ; dow {
        case 1:
            result = "It's Sunday!"
        case 7:
            result = "It's Saturday!"
        default:
            result = "It's a weekday"
    }
    fmt.Println(result)
    //fmt.Printf("Day %v, %v\n", dow, result)
    
    x := 42
    
    switch {
        case x < 0:
            result = "Less than zero"
        //    fallthrough       //to make it like regular switch
        case x == 0:
            result = "Equals zero"
        default:
            result = "Greater than zero"
    }
    fmt.Println(result)
}