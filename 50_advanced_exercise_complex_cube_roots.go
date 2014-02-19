package main

import "fmt"
import "math/cmplx"

/*
 * Let's explore Go's built-in support for complex numbers via the complex64 and complex128 types.
 * For cube roots, Newton's method amounts to repeating:
 *
 * Find the cube root of 2, just to make sure the algorithm works. There is a Pow function in the math/cmplx package.
 */
func Cbrt(x complex128) complex128 {
  var z complex128 = x * 0.33
  for i := 0; i < 10; i++ {
    z -= (z*z*z - x) / 3*z*z
    fmt.Println(z)
  }
  return z
}

func main() {
  fmt.Println(cmplx.Pow(2, 0.33))
  fmt.Println(Cbrt(2))
}
