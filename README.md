# SerenadeMi

A dynamic menu system for writing BBsMi appllications and more

## Table of Contents

- [Installation](#installation)
- [Usage](#usage)
- [Support](#support)
- [Contributing](#contributing)
- [License](#license)

## Installation

```sh
go get github.com/BBsMi/SerenadeMi
```

## Usage

Steps to use

### Make sure your go.mod is correctly initalized for your project

```sh
go mod init (package url)
```

### Create subfolder "menu" in your current project

```sh
mkdir menu
```
### Import menu

```golang
import "(your package url)/menu"
 
```

### Link your struct

Golang will not accept new functions added to an external package. To get around this we must declare we are extending the packages struct. In `menu/types.go`:

```golang
package SerenadeMiFunctions

import "github.com/BBsMi/SerenadeMi"

type Inits SerenadeMi.Inits

func Init() Inits {
  return Inits{}
}
```

### Create your functions

Create your functions in `./menu` using some sort of standardize naming scheme (examples: `F(function number).go`, `(function short name).go`, `(function number).go`) using the following template. Replace `FXXX` with F and that functions number.

```golang
package SerenadeMiFunctions

import "github.com/BBsMi/SerenadeMi"

func (i *Inits) FXXX_Init() SerenadeMi.Entry { return nil }

func FXXX(attrib *interface{}) error { return nil }

```

### Using your new menu

Boilerplate ( `go mod init myapp` with this example )
```golang
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
  
  // smi.Exec(uint(1))
  // smi.Exec("hello")
  // smi.Pager(uint(0), uint(100))
}
```

## Support

Please [open an issue](https://github.com/BBsMi/SerenadeMi/issues/new) for support.

## Contributing

Please contribute using [Github Flow](https://guides.github.com/introduction/flow/). Create a branch, add commits, and [open a pull request](https://github.com/fraction/readme-boilerplate/compare/).

## License

[(c) 2023 GNU AFFERO GENERAL PUBLIC LICENSE v3.0](https://github.com/BBsMi/SerenadeMi/blob/trunk/LICENSE)
