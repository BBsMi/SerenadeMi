package main

import (
  smf "myapp/menu"
  "github.com/BBsMi/SerenadeMi"
)

func main() {
  // Init our functions
  om := smf.Init()

  // Load menu into engine
  smi := SerenadeMi.Init(&om)

  smi.Spew()
  
  smi.Exec(uint(1))
  smi.Exec("WORLD")
  // smi.Pager(uint(0), uint(100))
}

