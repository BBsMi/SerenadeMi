package SerenadeMi

import (
  "fmt"
  "reflect"
  "github.com/davecgh/go-spew/spew"
)

func Init(i any) {
  // fmt.Println(i)
  t := reflect.TypeOf(i)
  fmt.Println(t.NumMethod())

  for m := 0; m < t.NumMethod(); m++ {
    // fmt.Println(runtime.FuncForPC(t.Method(m).Pointer()).Name())
    mm := t.Method(m)
    fmt.Println(mm.Name)
    val := reflect.ValueOf(i).MethodByName(mm.Name).Call([]reflect.Value{})
    res := val[0].Interface()
    spew.Dump(res, (res).FnName)
    // spew.Dump(res.Indirect())
    // fmt.Println(res)
    //*res.Fn()
  }
}
