package main

import "fmt"

func main() {
  var num int
  var primes []int
  fmt.Println("Please enter the top limit")
  fmt.Scan(&num)
  for i := 1; i < num; i++ {
    count := 0
    for j := 1; j < i; j++ {
      if i%j == 0 {
        count++
      }
    }
    if count == 1 {
      primes = append(primes, i)
    }
  }
  fmt.Println(primes)
}
