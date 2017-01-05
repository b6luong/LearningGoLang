package main

import (
    "fmt"
    "io"
    "io/ioutil"
    "os"
)

func main() {
    content := "Hello from Go!\n"
    
    file, err := os.Create("./GoTest.txt")
    checkError(err)
    defer file.Close()
    
    ln, err := io.WriteString(file, content)
    checkError(err)
    
    fmt.Printf("File written with %v characters\n", ln)
    
    bytes := []byte(content)
    ioutil.WriteFile("./fromBytes.txt", bytes, 0644) //0644 is the file permissions
}

func checkError(err error){
    if err != nil {
        panic(err)
    }
}