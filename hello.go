package main

import "fmt"

func Sum(x int, y int) int {
  return x + y
}

func main () {
  fmt.Println("Hello world")
  fmt.Println("The total is:", Sum(4, 6));
}
