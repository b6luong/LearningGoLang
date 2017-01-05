package main

import (
    "fmt"
    "os"
    "path/filepath"
)

func main() {
    
    root, _ := filepath.Abs("..")
    fmt.Println("Parent directory path: ", root)
    
    err := filepath.Walk(root, processPath)
    if err != nil {
        fmt.Println(err)
    }
}

// Everything in this method declaration is mandatory, except the name can be whatever you want
func processPath(path string, info os.FileInfo, err error) error {
    if err != nil {
        return err
    }
    
    // make sure we are not playing in the system root folder
    if path != "." {
        if info.IsDir() {
            fmt.Println("Directory: ", path)
        } else {
            fmt.Println("File: ", path)
        }
    } 
    
    return nil
}