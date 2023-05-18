package SerenadeMi

import (
  "fmt"
  "reflect"
  "github.com/davecgh/go-spew/spew"
)

// This function takes SerenadeMiFunctions.Inits in which is just a recast of SerenadeMi.Inits
// This function makes a list of all functions attached (one Initializer per function)

// func Init(i any) *[]Entry {
func Init(i any) *Engine {
  e := &Engine{}
  entries := &[]Entry{}
  t := reflect.TypeOf(i)
  fmt.Println("Initializers Found:", t.NumMethod())

  for m := 0; m < t.NumMethod(); m++ {
    mm := t.Method(m)
    fmt.Println(mm.Name)
    val := reflect.ValueOf(i).MethodByName(mm.Name).Call([]reflect.Value{})
    // Needs type assertion
    res := val[0].Interface().(Entry)
    *entries = append(*entries, res)
  }
  (*e).entries = entries
  return e
}

func (e *Engine) Spew() {
  spew.Dump(e)
}
