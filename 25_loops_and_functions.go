package main

import (
  "fmt"
  "math"
)

func Sqrt(x float64) float64 {
  z := 1.41
  epsilon := 0.1
  
  for i := 0; i < 10; i++ {
    z -= (math.Pow(z, 2) - x) / 2*z
    fmt.Println(z, math.Abs(x - z*z),)
    if math.Abs(x - z*z) < epsilon {
      break 
    } 
  }
  return z
}

func main() {
  fmt.Println(Sqrt(2))
  fmt.Println(math.Sqrt(2))
}
