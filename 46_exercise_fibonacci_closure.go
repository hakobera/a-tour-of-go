package main

import "fmt"

// fibonacci is a function that returns
// a function that retutns an int.
func fibonacci() func() int {
  n1, n2 := 0, 1
  return func() int {
    n1, n2 = n1 + n2, n1
    return n1
  }
}

func main() {
  f := fibonacci()
  for i := 0; i < 20; i++ {
    fmt.Println(f())
  }
}
