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

// Spew your secrets 
func (e *Engine) Spew() {
  spew.Dump(e)
}

// query needs to be reduced to either a single word or number in it's correct type
// eventually we'll be able to add attributes here to passthrough
func (e *Engine) Exec(query any) error {
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

func (e *Engine) Pager(from, to uint) *[]Entry {
  list &Entry{}
  for c := from; c<=to; c++ {
    l, ok := e.entriesByNumb[c]
    if ok {
      *list := append(*list, l)
    }
  }
  return list
}

// Function Number => Function
func (e *Engine) execByNumb(q uint) error {
  res, ok := e.entriesByNumb[q]
  if !ok { return fmt.Errorf("BADFUNC: Invalid Function Number") }
  return (*res).Fn(nil)
}

// Function Name => Function
func (e *Engine) execByName(q string) error {
  res, ok := e.entriesByName[q]
  if !ok { return fmt.Errorf("BADFUNC: Invalid Function Name") }
  return (*res).Fn(nil)
}
