package main

import "fmt"

func main() {
    n1, l1 := FullName("Benjamin", "Luong")
    fmt.Println("Full Name:", n1, "| size:", l1)
    
    n2, l2 := FullNameNakedReturn("John", "Legend")
    fmt.Println("Full Name:", n2, "| size:", l2)

}

// implicitly declare f and l to be string. Here I can return 2 values, indicated
// by the parentheses after the function. The return statement must be in that same
// order as well. In Go, we don't use get in the function name, we just put whatever
// value you are expecting. 
func FullName(f, l string) (string, int) {
    full := f + " " + l;
    length := len(full)
    return full, length
}

// You can declare the return values in the function signature and just call return
// at the end of the function. This is called "Naked Return"
func FullNameNakedReturn(f, l string) (full string, length int) {
    full = f + " " + l;
    length = len(full)
    return 
}