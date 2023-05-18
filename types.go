package SerenadeMi

type Inits struct {}

type Entry struct {
	FnNumb uint                             // Function Number
	FnName string                           // Function Name
	SDesc  string                           // Function Short Description
	Ldesc  string                           // Function Long Description
	Fn     func(attrib *interface{}) error  // Function itself
}

type Engine struct {
  // entries *[]Entry
  entriesByNumb map[uint]*Entry
  entriesByName map[string]*Entry
}
