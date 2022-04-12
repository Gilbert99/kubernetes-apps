package main

import (
   "fmt"
   "net/http"
   "log"
)

func hello(w http.ResponseWriter, r *http.Request) {
	fmt.Fprint(w, "hello from golang")
}

func main() {
    log.Println("Server is starting....")
    mux := http.NewServeMux()
    mux.HandleFunc("/", hello)
    err := http.ListenAndServe(":4000", mux)
    if err != nil {
          log.Fatal(err)
    }
}
