package main

import (
    "fmt"
    "net/http"
)

type MyServer struct {}

// The definition of this function must follow the one from net/http.
// Parameter names can be anything
func (s MyServer) ServeHTTP(rw http.ResponseWriter, req *http.Request) {
    // Fprint prints to a writer object
    fmt.Fprint(rw, "<h1>Hello from the Go web server!</h1")
}
func main() {
    var server MyServer
    err := http.ListenAndServe("localhost:8080", server)
    checkError(err)
}

func checkError(err error){
    if err != nil {
        panic(err)
    }
}