package SerenadeMiFunctions

import (
  "fmt"
  "github.com/BBsMi/SerenadeMi"
)

func (i *Inits) F1_Init() SerenadeMi.Entry { 
  return SerenadeMi.Entry{1, "HELLO", "Print Hello", "Test Print Hello", F1}
}

func F1(attrib *interface{}) error { 
  fmt.Println("Hello")
  return nil
}

