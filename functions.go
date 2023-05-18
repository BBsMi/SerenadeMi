package SerenadeMi

import (
  "fmt"
  "reflect"
)

func Init(i *Inits) {
  fmt.Println(i)
  t := reflect.TypeOf(&i)
  fmt.Println(t.NumMethod())

  for m := 0; m < t.NumMethod(); m++ {
    // fmt.Println(runtime.FuncForPC(t.Method(m).Pointer()).Name())
    mm := t.Method(m)
    fmt.Println(mm.Name)
    val := reflect.ValueOf(&i).MethodByName(mm.Name).Call([]reflect.Value{})
    fmt.Println(val)
  }
}
