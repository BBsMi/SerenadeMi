package SerenadeMi

type Inits struct {}

type Entry struct {
	fn_num  uint                            // Function Number
	fn_name string                          // Function Name
	s_desc  string                          // Function Short Description
	l_desc  string                          // Function Long Description
	fn      func(attrib *interface{}) error  // Function itself
}

