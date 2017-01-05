package main

import(
    "fmt"
    "bufio"
    "os"
    "strings"
)

func main() {
    reader := bufio.NewReader(os.Stdin)
    
    //Read a string from stdin
    fmt.Print("Enter a string: ")
    input, _ := reader.ReadString('\n')
    
    // Remove the \n at the end of input
    input = strings.TrimSpace(input)
    
    // Remove all the spaces in the string
    input = strings.ToLower(strings.Replace(input, " ", "", -1))
    fmt.Println("Iterative: ", palindromeInterative(input))
    fmt.Println("Recursive: ", palindromeRecursive(input))
}

func palindromeInterative(str string) bool {
    i := 0
    j := len(str) - 1
    for i < j {
        if str[i] != str[j] {
            return false
        }
        i++
        j--
    }
    return true
}

func palindromeRecursive(str string) bool {
    if len(str) <= 1 {
        return true;
    } else if str[0] != str[len(str) - 1] {
        return false;
    }
    return palindromeRecursive(str[1:len(str) - 1])
}