package SerenadeMi

type Inits struct {}

type Entry struct {
	FnNum  uint                             // Function Number
	FnName string                           // Function Name
	SDesc  string                           // Function Short Description
	Ldesc  string                           // Function Long Description
	Fn     func(attrib *interface{}) error  // Function itself
}

type Engine struct {
  entries *[]Entry
}

