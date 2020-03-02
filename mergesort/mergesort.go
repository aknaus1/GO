package main

import (
	"fmt"
  "math/rand"
  "time"
)

/*
	Rand num function: Returns a random int of max value "max"
	Parameters: max -> upper bound random int
*/
func rand_num(max int) int{
	a1 := rand.NewSource(time.Now().UnixNano())
	a2 := rand.New(a1)
	num := a2.Intn(max)
	return num
}

/*
	Generate_list function: Returns an array of randomly generated ints
	Parameters: n -> number of ints, m -> max value for ints
*/
func generate_list(n int, m int) []int{
  var list []int
  for i := 0; i < n; i++ {
    list = append(list, rand_num(m))
  }
  return list
}

/*
	Merge_sort function: Returns a array of ints sorted from smallest to largest
	Parameters: list -> array of ints
*/
func merge_sort(list []int) []int{
	if len(list) == 1 {
		return list
	}
	mid := int(len(list)/2)
	var left []int
	var right []int
	for i := 0; i < mid; i++ {
		left = append(left, list[i])
	}
	for j := mid; j < len(list); j++ {
		right = append(right, list[j])
	}
	return merge(merge_sort(left), merge_sort(right))
}

/*
	Merge function: Partially sorts 2 arrays,
	and combines them into a single array.
	Parameters: left -> array of int, right -> array of ints
*/
func merge(left []int, right []int) []int {
	var result []int
	i := 0
	for (len(left) > 0) && (len(right) > 0) {
		if left[0] < right[0] {
			result = append(result, left[0])
			left = left[1:]
			i++
		} else{
			result = append(result, right[0])
			right = right[1:]
			i++
		}
	}
	for j := 0; j < len(left); j++ {
		result = append(result, left[j])
	}
	for j := 0; j < len(right); j++ {
		result = append(result, right[j])
	}
	return result
}

/*
	Main funciton: Runs on startup
*/
func main() {
	n := 20
	m := 100
  list := generate_list(n, m)
  fmt.Println("Unsorted:", list)
  list  = merge_sort(list)
  fmt.Println("Sorted:", list)
}
