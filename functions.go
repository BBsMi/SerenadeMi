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
  
  // entries := &[]Entry{}

  (*e).ebName := make(map[string]*Entry)
  (*e).ebNumb := make(map[uint]*Entry)

  t := reflect.TypeOf(i)
  fmt.Println("Initializers Found:", t.NumMethod())

  for m := 0; m < t.NumMethod(); m++ {
    mm := t.Method(m)
    fmt.Println(mm.Name)
    val := reflect.ValueOf(i).MethodByName(mm.Name).Call([]reflect.Value{})
    // Needs type assertion
    res := val[0].Interface().(Entry)
    // *entries = append(*entries, res)
    (*e).ebName[res.FnName] = &res
    (*e).ebNumb[res.FnNumb] = &res
  }
  // (*e).entries = entries
  return e
}

func (e *Engine) Spew() {
  spew.Dump(e)
}

func (e *Engine) Exec(query any) error {
  //fmt.Println(reflect.TypeOf(query))
  switch query.(type) {
    case int:
      fmt.Println("We have an int")
    case string:
      fmt.Println("We have a string")
    default:
      return fmt.Errorf("BADTYPE: Invalid Type")
  }
  return nil
}

