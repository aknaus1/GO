package main

import (
  "fmt"
)

func main() {
  var word string = "Hello World"
  alt := "Hello World"
  var option = "Hello World"
  for true {
    fmt.Println(word)
    fmt.Printf("%s\n", alt)
    fmt.Printf("%T\n", option)
    break
  }
}
