package main 

import (
    "fmt"
)

func main() {
    var x float64 = -42
    var result string
    
    if x < 0 {
        result = "x is Less than zero"
    } else if x == 0 {
        result = "x is Equal to zero"  
    } else {
        result = "x is Greater than zero"
    }
    
    fmt.Println("Result:", result)
    
    if y := 42; y < 0 { // You can use a local variable inside the if block. It only lasts during the conditional logic
        result = "y is Less than zero"
    } else if x == 0 {
        result = "y is Equal to zero"  
    } else {
        result = "y is Greater than zero"
    }

    fmt.Println("Result:", result)

}