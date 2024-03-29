package main

import (
  "net/http"
  "strings"
  "fmt"
)
func sayHello(w http.ResponseWriter, r *http.Request) {
  message := r.URL.Path
  message = strings.TrimPrefix(message, "/")
  message = "Hello " + message  
  w.Write([]byte(message))
}

func main() {
  fmt.Println("Run HTTP server")
  http.HandleFunc("/", sayHello)
  if err := http.ListenAndServe(":8090", nil); err != nil {
    panic(err)
  }
}