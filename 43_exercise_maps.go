package main

import (
  "code.google.com/p/go-tour/wc"
  "strings"
)

func WordCount(s string) map[string]int {
  words := strings.Fields(s)
  ret := make(map[string]int)
  
  for i := range words {
    v := ret[words[i]]
    ret[words[i]] = v + 1
  }

  return ret
}

func main() {
  wc.Test(WordCount)
}
