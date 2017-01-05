package main

import (
    "fmt"
    "math/big"
    "math"
)

func main(){
    
    i1, i2, i3 := 12, 45, 68
    var intSum int = i1 + i2 + i3
    fmt.Println("intSum: " , intSum)
    
    f1, f2, f3 := 23.5, 65.1, 76.3
    floatSum := f1 + f2 + f3
    fmt.Printf("floatSum: %v, %T\n", floatSum, floatSum) //default float64
    
    var b1, b2, b3, bigSum big.Float
    b1.SetFloat64(23.5)
    b2.SetFloat64(65.1)
    b3.SetFloat64(76.3)
    
    bigSum.Add(&b1, &b2).Add(&bigSum, &b3)
    fmt.Printf("BigSum: %.10g, %T\n", &bigSum, bigSum)
    
    radius := 15.5
    circumference := 2 * radius * math.Pi
    fmt.Printf("circumference: %3.2f\n", circumference)
}