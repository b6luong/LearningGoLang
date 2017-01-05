package main

import (
    "fmt"
    "net/http"
    "io/ioutil"
    "encoding/json"
    "strings"
    "math/big"
)

// Creating a struct with the values you want from JSON
type Tour struct {
    Name, Price string
}
func main(){
    
    url := "http://services.explorecalifornia.org/json/tours.php"
    content := contentFromServer(url)
    
    //fmt.Println(content)
    
    tours := toursFromJson(content)
    //fmt.Println(tours)
    
    for _, tour := range tours {
        price, _, _ := big.ParseFloat(tour.Price, 10, 2, big.ToZero)
        fmt.Printf("%v ($%.2f)\n", tour.Name, price)
    }
}

func checkError(err error) {
    if err != nil {
        panic(err)
    }
}

func contentFromServer(url string) string {
    
    response, err := http.Get(url)
    checkError(err)
    defer response.Body.Close()
    
    bytes, err := ioutil.ReadAll(response.Body)
    checkError(err)
    
    return string(bytes)
}

func toursFromJson(content string) []Tour {
    tours := make([]Tour, 0, 20)
    
    decoder := json.NewDecoder(strings.NewReader(content))
    _, err := decoder.Token() // read the first '[' since decoder cannot handle array of objects
    checkError(err)
    
    var tour Tour
    for decoder.More() {
        // This method parses the json and assigns the values that matches the struct data fields
        err := decoder.Decode(&tour)
        checkError(err)
        tours = append(tours, tour)
    }
    
    return tours
}