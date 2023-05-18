package SerenadeMiFunctions

import (
  "fmt"
  "github.com/BBsMi/SerenadeMi"
)

func (i *Inits) F2_Init() SerenadeMi.Entry { 
  return SerenadeMi.Entry{2, "WORLD", "Print World", "Test Print World", F2}
}

func F2(attrib *interface{}) error { 
  fmt.Println("World")
  return nil
}

