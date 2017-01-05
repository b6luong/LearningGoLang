// Exploring the fmt package and outputting to console

package main

import "fmt"

func main() {
    str1 := "The quick red fox"
    str2 := "jumped over"
    str3 := "the lazy brown dog"
    
    aNumber := 42
    
    isTrue := true
    
    // You can use an _ (underscore) for variables that are returned but aren't used
    stringLength, err := fmt.Println(str1, str2, str3)
    //OUTPUT: The quick red fox jumped over the lazy brown dog

    if err == nil {
        fmt.Println("String length:", stringLength);
        //OUTPUT: String length: 49
    }
    
    fmt.Printf("Value of aNumber: %v\n", aNumber)
    fmt.Printf("Value of aNumber as float: %.2f\n", float64(aNumber)) //float64(int) converts int to float
    fmt.Printf("Value of isTrue: %v\n", isTrue)
    
    fmt.Printf("Data types: %T, %T, %T, %T, and %T\n", str1, str2, str3, aNumber, isTrue)
    
    myString := fmt.Sprintf("Data types as var: %T, %T, %T, %T, and %T\n", str1, str2, str3, aNumber, isTrue)
    fmt.Print(myString)
    //OUTPUT: Data types as var: string, string, string, int, and bool

}