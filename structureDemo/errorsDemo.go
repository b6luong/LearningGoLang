package main

import (
    "fmt"
    "os"
    "errors"
)

func main() {
    file, err := os.Open("doesnotexist.ext")
    
    if err == nil {
        fmt.Println(file)
    } else {
        fmt.Println(err.Error())
        // OR
        // fmt.Println(err)
    }
    
    myError := errors.New("My error string")
    fmt.Println(myError)
    
    attendence := map[string]bool {
        "Mike":true,
        "Ann":true,
        "Tony":false}
    
    // Comma-OK syntax
    // In this example, attended is the value corresponding to the key, and ok is true if there is an entry for that key, otherwise false if there isn't an entry
    attended, ok := attendence["Mike"]
    if ok {
        fmt.Println("Mike attended?", attended)
    } else {
        fmt.Println("No intel on Mike")
    }
    
    attended, ok = attendence["Michelle"]
    if ok {
        fmt.Println("Michelle attended?", attended)
    } else {
        fmt.Println("No intel on Michelle")
    }
}