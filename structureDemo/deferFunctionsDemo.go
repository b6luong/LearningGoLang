package main

// It defers execution of a block of code until everything else in the current function is finished.
import (
    "fmt"
)

// Each time you call the "defer" statement, it adds an instruction to a stack, and when the deferred statements are executed, they're handled in LIFO order
func main() {
    defer fmt.Println("Close the file")
    fmt.Println("Open the file")
    
    fmt.Println(myFunc())
    
    defer fmt.Println("Statement 1")
    defer fmt.Println("Statement 2")
    defer fmt.Println("Statement 3")
    defer fmt.Println("Statement 4")
    fmt.Println("Undefered Statement")
    
    x := 1000
    defer fmt.Println("Value of x:", x) //1000
    x++
    fmt.Println("Value of x after increment",x) //1001
    
    y := &x
    defer fmt.Println("Value of *y:", *y) //1001
    *y++ // (*y)++ give the same output
    fmt.Println("Value of *y after increment",*y) //1002
}

func myFunc() string {
    defer fmt.Println("Deferred function in MyFunc")
    fmt.Println("Not deferred function in MyFunc")
    return "Return string" // All defered codes will be called before the return statement is called
}