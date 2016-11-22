
## Important links
- [Golang中国](http://golangtc.com/)
- [Go by example](https://gobyexample.com/)
- [Go wiki](https://github.com/golang/go/wiki)
- [The Go forum](https://forum.golangbridge.org/)

### Go tools
[Golang code inspection tools](http://blog.ralch.com/tutorial/golang-tools-inspection/)
godegraph uses [Graphviz](http://graphviz.org/)
[Golang code comprehension tools](http://blog.ralch.com/tutorial/golang-tools-comprehension/)

- [godegraph - A tool for generating dependency graphs of Go code.](https://github.com/kisielk/godepgraph)

install godegrpah
`go get github.com/kisielk/godepgraph`

use godegraph
// this command will generate dependency graph as svg image
// -s ingores the builtin packages
// -horizontal layout the graphics horizontally
$ godepgraph -s -horizontal github.com/codegangsta/gin | dot -Tsvg -o gin-godepgraph.svg

- oracle 
The oralce is a source analysis tool that answers question about your Go source code. 

install oracle:
`go get golang.org/x/tools/cmd/oracle`

- Pythia
Pythia is a browser based user interface for oracle. Pythia is a browser based user interface for oracle. It is based on the following packages: oracle and godoc

install pythia 
`go get github.com/fzipp/pythia`

use pythia 
A specific package can be opened with the following command:
`pythia net/http`

- go vet
Vet does analysis on Go source code and reports suspicious constructs. It uses heuristics that do not guarantee all reports are genuine problems. Vet can find errors not caught by the compilers.

It can be invoked in three different ways:

	// for go package
	$ go tool vet package/path/name
	// for files
	$ go tool vet source/directory/*.go
	// for directory
	$ go tool vet source/directory

- golint
Golint differs from gofmt and govet. It prints out style mistakes. Golint is concerned with coding style. It is in use at Google, and it seeks to match the accepted style of the open source Go project.

Golint make suggestions regarding source code. It is not perfect, and has both false positives and false negatives. Do not consider its output as a truth. It will never be trustworthy enough to be enforced automatically as part of a build process.

Installation

`go get -u github.com/golang/lint/golint`

Usage

	// analysis a particular package
	$ golint package
	// analysis a particular directory
	$ golint directory 
	// analyses a particualr files
	$ golint files 







## go cross compiler: gox

### install gox
` go get github.com/mitchellh/gox`

## interactive debuging 
[Does any golang interactive debugger exist?](http://stackoverflow.com/questions/16492509/does-any-golang-interactive-debugger-exist)

### godebug
[godebug Github repository](https://github.com/mailgun/godebug)

#### Installation:

`$ go get github.com/mailgun/godebug`

## install golang

### Ubuntu 
version: 

1. get golang binary 
  + you can download from here: 
  []()
  + or if you have a local copy, you can use scp to copy it to the target machine
  `scp golangversion hank@www.mzworld.cn:/home/hank`
2. unzip the binary 
> `tar -C /usr/local -xzf go$VERSION.$OS-$ARCH.tar.gz` 

3. Add go to your path `export PATH=$PATH:/usr/local/go/bin`
4. Add GOPATH: `export GOPATH=~/Goworkspace`
5. Add bin to your path: `export PATH=$PATH:$GOPATH/bin`

- use gvm
[http://stackoverflow.com/questions/17480044/how-to-install-the-current-version-of-go-in-ubuntu-precise](http://stackoverflow.com/questions/17480044/how-to-install-the-current-version-of-go-in-ubuntu-precise)

> I like to use GVM for managing my Go versions in my Ubuntu box. Pretty simple to use, and if you're familiar with RVM, it's a nobrainer. It allows you to have multiple versions of Go installed in your system and switch between whichever version you want at any point in time.

Install GVM with:

> `sudo apt-get install bison mercurial
> `bash < <(curl -LSs 'https://raw.githubusercontent.com/moovweb/gvm/master/binscripts/gvm-installer') . "$HOME/.gvm/scripts/gvm"`

and then it's as easy as doing this:

    gvm install go1.1.1
    gvm use go1.1.1 --default

The default flag at the end of the second command will set go1.1.1 to be your default Go version whenever you start a new terminal session.