package main

import "fmt"

// fibonacci is a function that returns
// a function that retutns an int.
func fibonacci() func() int {
  var f func(i int) int

  f = func(i int) int {
    if (i == 0) {
      return 0
    } else if (i == 1) {
      return 1
    } else {
      return f(i-1) + f(i-2)
    }
  }

  n := 0
  return func() int {
    fib := f(n)
    n++
    return fib
  }
}

func main() {
  f := fibonacci()
  for i := 0; i < 20; i++ {
    fmt.Println(f())
  }
}
