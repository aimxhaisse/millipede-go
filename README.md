millipede-go
============

Golang version of [getmillipede](https://github.com/getmillipede/).

## Usage

```command
$ millipede-go -h
NAME:
   millipede-go - Print a beautiful millipede

USAGE:
   millipede-go [global options] command [command options] [arguments...]

VERSION:
   1.2.0-dev (HEAD)

AUTHOR(S):
   Millipede crew <https://github.com/getmillipede/millipede-go>

COMMANDS:
   help, h     Shows a list of commands or help for one command

GLOBAL OPTIONS:
   --reverse, -r                           reverse the millipede
   --opposite, -o                          go the opposite direction
   --skin, --template, -s, -t "default"    millipede skin (default, frozen, love, corporate, musician, bocal, ascii)
   --width, -w "3"                         millipede width
   --curve, -c "4"                         millipede curve size
   --help, -h                              show help
   --version, -v                           print the version
```

## Examples

```command
$ millipede-go
    ╚⊙ ⊙╝
  ╚═(███)═╝
 ╚═(███)═╝
╚═(███)═╝
 ╚═(███)═╝
  ╚═(███)═╝
   ╚═(███)═╝
    ╚═(███)═╝
    ╚═(███)═╝
   ╚═(███)═╝
  ╚═(███)═╝
 ╚═(███)═╝
╚═(███)═╝
 ╚═(███)═╝
  ╚═(███)═╝
   ╚═(███)═╝
    ╚═(███)═╝
    ╚═(███)═╝
   ╚═(███)═╝
  ╚═(███)═╝
 ╚═(███)═╝
```

```command
$ millipede-go 42
    ╚⊙ ⊙╝
  ╚═(███)═╝
 ╚═(███)═╝
╚═(███)═╝
 ╚═(███)═╝
  ╚═(███)═╝
   ╚═(███)═╝
    ╚═(███)═╝
    ╚═(███)═╝
   ╚═(███)═╝
  ╚═(███)═╝
 ╚═(███)═╝
╚═(███)═╝
 ╚═(███)═╝
  ╚═(███)═╝
   ╚═(███)═╝
    ╚═(███)═╝
    ╚═(███)═╝
   ╚═(███)═╝
  ╚═(███)═╝
 ╚═(███)═╝
╚═(███)═╝
 ╚═(███)═╝
  ╚═(███)═╝
   ╚═(███)═╝
    ╚═(███)═╝
    ╚═(███)═╝
   ╚═(███)═╝
  ╚═(███)═╝
 ╚═(███)═╝
╚═(███)═╝
 ╚═(███)═╝
  ╚═(███)═╝
   ╚═(███)═╝
    ╚═(███)═╝
    ╚═(███)═╝
   ╚═(███)═╝
  ╚═(███)═╝
 ╚═(███)═╝
╚═(███)═╝
 ╚═(███)═╝
  ╚═(███)═╝
   ╚═(███)═╝
```

```command
$ millipede-go --reverse
 ╔═(███)═╗
  ╔═(███)═╗
   ╔═(███)═╗
    ╔═(███)═╗
    ╔═(███)═╗
   ╔═(███)═╗
  ╔═(███)═╗
 ╔═(███)═╗
╔═(███)═╗
 ╔═(███)═╗
  ╔═(███)═╗
   ╔═(███)═╗
    ╔═(███)═╗
    ╔═(███)═╗
   ╔═(███)═╗
  ╔═(███)═╗
 ╔═(███)═╗
╔═(███)═╗
 ╔═(███)═╗
  ╔═(███)═╗
    ╔⊙ ⊙╗
```

```command
$ millipede-go --opposite
    ╚⊙ ⊙╝
  ╚═(███)═╝
   ╚═(███)═╝
    ╚═(███)═╝
    ╚═(███)═╝
   ╚═(███)═╝
  ╚═(███)═╝
 ╚═(███)═╝
╚═(███)═╝
 ╚═(███)═╝
  ╚═(███)═╝
   ╚═(███)═╝
    ╚═(███)═╝
    ╚═(███)═╝
   ╚═(███)═╝
  ╚═(███)═╝
 ╚═(███)═╝
╚═(███)═╝
 ╚═(███)═╝
  ╚═(███)═╝
   ╚═(███)═╝
```

```command
$ millipede-go --skin=bocal --opposite --reverse
   ╔═(🐟🐟🐟)═╗
  ╔═(🐟🐟🐟)═╗
 ╔═(🐟🐟🐟)═╗
╔═(🐟🐟🐟)═╗
 ╔═(🐟🐟🐟)═╗
  ╔═(🐟🐟🐟)═╗
   ╔═(🐟🐟🐟)═╗
    ╔═(🐟🐟🐟)═╗
    ╔═(🐟🐟🐟)═╗
   ╔═(🐟🐟🐟)═╗
  ╔═(🐟🐟🐟)═╗
 ╔═(🐟🐟🐟)═╗
╔═(🐟🐟🐟)═╗
 ╔═(🐟🐟🐟)═╗
  ╔═(🐟🐟🐟)═╗
   ╔═(🐟🐟🐟)═╗
    ╔═(🐟🐟🐟)═╗
    ╔═(🐟🐟🐟)═╗
   ╔═(🐟🐟🐟)═╗
  ╔═(🐟🐟🐟)═╗
    ╔⊙ ⊙╗
```

```command
$ millipede-go --skin=ascii --width=7
    \o     o/
  |=(#######)=|
 |=(#######)=|
|=(#######)=|
 |=(#######)=|
  |=(#######)=|
   |=(#######)=|
    |=(#######)=|
    |=(#######)=|
   |=(#######)=|
  |=(#######)=|
 |=(#######)=|
|=(#######)=|
 |=(#######)=|
  |=(#######)=|
   |=(#######)=|
    |=(#######)=|
    |=(#######)=|
   |=(#######)=|
  |=(#######)=|
 |=(#######)=|
```

```command
$ millipede-go --skin=bocal --curve=8 --opposite
          ╚⊙ ⊙╝
        ╚═(🐟🐟🐟)═╝
       ╚═(🐟🐟🐟)═╝
      ╚═(🐟🐟🐟)═╝
     ╚═(🐟🐟🐟)═╝
    ╚═(🐟🐟🐟)═╝
   ╚═(🐟🐟🐟)═╝
  ╚═(🐟🐟🐟)═╝
 ╚═(🐟🐟🐟)═╝
╚═(🐟🐟🐟)═╝
 ╚═(🐟🐟🐟)═╝
  ╚═(🐟🐟🐟)═╝
   ╚═(🐟🐟🐟)═╝
    ╚═(🐟🐟🐟)═╝
     ╚═(🐟🐟🐟)═╝
      ╚═(🐟🐟🐟)═╝
       ╚═(🐟🐟🐟)═╝
        ╚═(🐟🐟🐟)═╝
       ╚═(🐟🐟🐟)═╝
      ╚═(🐟🐟🐟)═╝
     ╚═(🐟🐟🐟)═╝
```

```command
$ millipede-go --skin=bocal --curve=0
  ╚⊙ ⊙╝
╚═(🐟🐟🐟)═╝
╚═(🐟🐟🐟)═╝
╚═(🐟🐟🐟)═╝
╚═(🐟🐟🐟)═╝
╚═(🐟🐟🐟)═╝
╚═(🐟🐟🐟)═╝
╚═(🐟🐟🐟)═╝
╚═(🐟🐟🐟)═╝
╚═(🐟🐟🐟)═╝
╚═(🐟🐟🐟)═╝
╚═(🐟🐟🐟)═╝
╚═(🐟🐟🐟)═╝
╚═(🐟🐟🐟)═╝
╚═(🐟🐟🐟)═╝
╚═(🐟🐟🐟)═╝
╚═(🐟🐟🐟)═╝
╚═(🐟🐟🐟)═╝
╚═(🐟🐟🐟)═╝
╚═(🐟🐟🐟)═╝
╚═(🐟🐟🐟)═╝
```

## Install

Download a statically compiled binary

- Go to [github.com/getmillipede/millipede-go/releases/latest](https://github.com/getmillipede/millipede-go/releases/latest)

---

Download and build using Golang

```command
$ go get github.com/getmillipede/millipede-go
```

## License

[MIT](https://github.com/getmillipede/millipede-go/blob/master/LICENSE)
