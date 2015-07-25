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
   1.1.0-dev (HEAD)

AUTHOR(S):
   Millipede crew <https://github.com/getmillipede/millipede-go>

COMMANDS:
   help, h	Shows a list of commands or help for one command

GLOBAL OPTIONS:
   --reverse, -r                           reverse the millipede
   --opposite, -o                          go the opposite direction
   --skin, --template, -s, -t "default"    millipede skin (default)
   --help, -h		                       show help
   --version, -v	                       print the version
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
