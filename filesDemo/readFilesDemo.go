package main

import (
    "fmt"
    "io/ioutil"
)

func main() {
    
    content, err := ioutil.ReadFile("./GoTest.txt")
    checkError(err)
    
    fmt.Println("Read from File (bytes):", content)
    
    str := string(content)
    fmt.Println("Read from File (string):", str)
}

func checkError(err error){
    if err != nil {
        panic(err)
    }
}