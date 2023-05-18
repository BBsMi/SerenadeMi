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
  
  (*e).entriesByName = make(map[string]*Entry)
  (*e).entriesByNumb = make(map[uint]*Entry)

  t := reflect.TypeOf(i)
  fmt.Println("Initializers Found:", t.NumMethod())

  for m := 0; m < t.NumMethod(); m++ {
    mm := t.Method(m)
    fmt.Println(mm.Name)
    val := reflect.ValueOf(i).MethodByName(mm.Name).Call([]reflect.Value{})
    // Needs type assertion
    res := val[0].Interface().(Entry)
    // *entries = append(*entries, res)
    (*e).entriesByName[res.FnName] = &res
    (*e).entriesByNumb[res.FnNumb] = &res
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
    case uint:
      return e.execByNumb(query.(uint))
    case string:
      return e.execByName(query.(string))
    default:
      return fmt.Errorf("BADTYPE: Invalid Type")
  }
  return nil
}

func (e *Engine) execByNumb(q uint) error {
  res, ok := e.entriesByNumb[q]
  if !ok { return fmt.Errorf("BADFUNC: Invalid Function Number") }
  (*res).Fn(nil)
  return nil
}
func (e *Engine) execByName(q string) error {
  /*
  res, ok := e.entriesByName[q]
  if !ok { return fmt.Errorf("BADFUNC: Invalid Function Name") }
  (*res).Fn(nil)
  */
  return nil
}
