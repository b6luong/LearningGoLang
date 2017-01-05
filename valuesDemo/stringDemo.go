package main

import (
    "fmt"
    "strings"
)

func main(){
    str1 := "An implicityly typed string"
    fmt.Printf("str1: %v, %T\n", str1, str1)
    
    var str2 string = "An explicityly typed string"
    fmt.Printf("str2: %v, %T\n", str2, str2)
    
    fmt.Println("strings.ToUpper(str1): ", strings.ToUpper(str1))
    fmt.Println("strings.Title(str2): ", strings.Title(str2))
    
    uValue := "HELLO"
    lValue := "hello"
    
    fmt.Printf("uValue: %v, %T; lValue: %v, %T\n", uValue, uValue, lValue, lValue)
    fmt.Println("Equal?", uValue == lValue)
    fmt.Println("Equal Non-Case-Sensitive?", strings.EqualFold(uValue, lValue))
    
    fmt.Println("Contains exp?", strings.Contains(str1, "exp"))
    fmt.Println("Contains exp?", strings.Contains(str2, "exp"))
    
}