millipede-go :bug:
==================

[![Build Status](https://travis-ci.org/getmillipede/millipede-go.svg?branch=master)](https://travis-ci.org/getmillipede/millipede-go)
[![GoDoc](https://godoc.org/github.com/getmillipede/millipede-go/millipede?status.svg)](https://godoc.org/github.com/getmillipede/millipede-go/millipede)
![License](https://img.shields.io/github/license/getmillipede/millipede-go.svg)


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
   --chameleon                             the millipede use its environment color
   --rainbow                               the millipede live with care bears
   --zalgo                                 invoke the hive-mind representing chaos
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


```command
$ millipede-go --zalgo
  ╚⊙̲͋ ⊙̴╝͇͠
╚═(̀█̴█̴█̨)̛̻═͜╝͍͡
 ╚̘═̶(ͪ́█̛ͯ█҉█͚́)̷̰═͙͜╝̴̦
  ╚̨═(͉͋█͘█͜█̲͘)̷═̢╝̴
   ╚═̯(̴█͢█҉█͟)͢═̫͝╝͏
    ╚͝═(̨█͜█̶█̜͠)̯̉́═̢╝̧͎͛
   ╚̩ͩ═̵(ͦ█̴█̡█̤͟)̧̟═̸͕̓╝̕
  ╚͟═̨(͟█̞͠█̕█́͢)̪ͪ͟═̡͎╝̸
 ╚═͞(█̴̚█̨̾█̓͜)͜═̴╝̡
╚̲═̨(█͠█҉̲█̴̮̓)̸̭═̕╝̴͎
 ╚͏═҉̬(̀█̛̬ͥ█̶͕█̷)͝═͈͠╝́
  ╚̵═(█̚͘█́█̴͚)̷̱═̶̺̎╝̛͉
   ╚═́(̙̀█̨█̢█ͧ͘)̵═̸╝҉
    ╚̛͉̔═̶̻(█̡͔█̓͘█̢͐)̡̜═͘╝̀
   ╚̷═̫(̞͗█̢͇█̴͉█̶̥)̷═̤̑͡╝̝̏͡
  ╚═͕(ͭ█̛█̸█̢͇)͞═͎͠╝͞
 ╚̧═̀(҉̙█̍͏█̟͘█͏)̵̫═̡̱╝̷
╚═̸(̉͟█̸█͝█̢)͟═̴͎╝̼̑͞
 ╚═(̂█̵͕█͏█̛)̧═̶̿╝̖̉͞
  ╚̡═̴ͦ(̦█̨█̶̥█҉)̛̮═͡╝̴
   ╚═(̨█͢█̨█̸)̡═̕╝̛̋

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
