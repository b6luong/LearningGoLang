package main

import (
    "fmt"
    "net/http"
    "io/ioutil"
)

func main(){
    
    //url := "http://headers.jsontest.com"
    url := "http://www.scs.ryerson.ca/~b6luong/index.html"
    
    response, err := http.Get(url)
    if err != nil {
        panic(err)
    }
    
    fmt.Printf("Response type: %T\n", response)
    defer response.Body.Close()
    
    bytes, err := ioutil.ReadAll(response.Body)
    if err != nil {
        panic(err)
    }
    
    content := string(bytes)
    fmt.Print(content)
}