package main

import (
//  "io"
	//"log"
	"net/http"
)

func main() {
  http.HandleFunc("/hello", func(writer http.ResponseWriter, request *http.Request) {
    writer.Write([]byte("Hello World \n"))
  })

  err := http.ListenAndServe(":8080", nil)
  if err != nil {
    panic(err)
  }
}
