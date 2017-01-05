package main

import (
    "fmt"
    "sort"
)

func main() {
    provinces := make(map[string]string) // allocate and initialize memory
    fmt.Println(provinces)
    
    provinces["ON"] = "Ontario"
    provinces["BC"] = "British Columbia"
    provinces["AL"] = "Alberta"
    fmt.Println(provinces)

    westside := provinces["BC"]
    fmt.Println(westside)

    delete(provinces, "AL")
    fmt.Println(provinces)
    
    provinces["Sa"] = "Saskatchewan"
    provinces["NS"] = "Nova Scotia"
    provinces["MA"] = "Manitoba"
    
    for k, v := range provinces {
        fmt.Printf("%v : %v\n", k , v)
    }
    
    // Sort a map
    
    keys := make([]string, len(provinces))
    i := 0
    for k := range provinces{
        keys[i] = k
        i++
    }
    
    sort.Strings(keys)
    fmt.Println("\nSorted")
    
    for i := range keys {
        fmt.Println(provinces[keys[i]])
    }
}