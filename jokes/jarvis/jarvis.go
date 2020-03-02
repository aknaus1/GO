package main

import (
  "fmt"
  "strings"
)

func contains(input string, vari string) bool {
  temp := strings.Split(input, " ")
  for i := 0; i < len(temp); i++ {
    if vari == temp[i]{
      return true
    }
  }
  return false
}

func main() {
  var input []string
  fmt.Println("Hello Sir")
  fmt.Println("How may I help you?")
  for true {
    i := 0
    repeat := fmt.Scan(&input[i])
    i++
    
    if !repeat {
      break
    }
  }
  if contains(input, "time") {
    fmt.Println("The time is ")
  } else if contains(input, "weather") {
    fmt.Println("The weather outside is")
  } else if contains(input, "jokes") {
    fmt.Println("PENIS")
  } else if contains(input, "math") {
    fmt.Println("1 + 1 = 3")
  } else if contains(input, "help") {
    fmt.Println("Mert Can Selcuk")
  } else {
    fmt.Println("You just stupid")
  }
}
