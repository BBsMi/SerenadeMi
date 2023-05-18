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

### Create your functions

Create your functions in `./menu` using some sort of standardize naming scheme (examples: `F(function number).go`, `(function short name).go`, `(function number).go`) using the following template. Replace `FXXX` with F and that functions number.

```golang
package SerenadeMi

func (i *Inits) FXXX_Init() *Entry { return nil }

func FXXX(attrib interface{}) error { return nil }

```

## Support

Please [open an issue](https://github.com/BBsMi/SerenadeMi/issues/new) for support.

## Contributing

Please contribute using [Github Flow](https://guides.github.com/introduction/flow/). Create a branch, add commits, and [open a pull request](https://github.com/fraction/readme-boilerplate/compare/).

## License

[(c) 2023 GNU AFFERO GENERAL PUBLIC LICENSE v3.0](https://github.com/BBsMi/SerenadeMi/blob/trunk/LICENSE)
