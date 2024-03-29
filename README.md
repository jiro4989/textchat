# textchat

[![Build Status](https://travis-ci.org/jiro4989/textchat.svg?branch=develop)](https://travis-ci.org/jiro4989/textchat)

textchat is a terminal chat cli.

## Usage

Simple example is below.

```bash
$ textchat Hello
.------.  .---------.                                                           
| 福本 | <   hello  |                                                           
`------'  `---------'                                                           

$ textchat -r Hello
                                                           .---------.  .------.
                                                           |  Hello   > | 野中 |
                                                           `---------'  `------'
```

You can set icon AA file.

```bash
$ textchat -i testdata/block.txt Hello
.-----.  .--------------.                                                       
| 123 | <   こんにちは  |                                                       
| 456 |  `--------------'                                                       
| 789 |                                                                         
`-----'                                                                         
```

You can set a chat text from stdin.

```bash
$ seq 5 | textchat
.------.  .-----.                                                               
| 村田 | <   1  |                                                               
`------'  |  2  |                                                               
          |  3  |                                                               
          |  4  |                                                               
          |  5  |                                                               
          `-----'                                                               
```

You can set a your name.

```bash
$ textchat -n jiro こんにちは
.------.  .--------------.                                                      
| jiro | <   こんにちは  |                                                      
`------'  `--------------'                                                      
```

Other examples.

```bash
$ textchat -i <(cat testdata/block.txt) Hello world
$ textchat -w 100 -r Right
```

## Installation

```bash
go get github.com/jiro4989/textchat
```

or

Download binary from [Releases](https://github.com/jiro4989/textchat/releases).

## Development

go version go1.12 linux/amd64

### How to build

You run below.

```bash
make build
```

You run below if you want to do cross compiling.

```bash
make bootstrap
make xbuild
```

### Testing

You run below.

```bash
make test
```

## Help

```
textchat is a terminal chat cli.

Usage:
	textchat [options]
	textchat [options] <word>...
	textchat -h | --help
	textchat -v | --version

Options:
	-h --help               Show this screen.
	-v --version            Show version.
	-r --right              Say word on the right side.
	-i --icon=<textfile>    Set icon AA file.
	-w --width=<width>      Set chat screen width. [default: 80]
	-p --pad=<pad>          Padding string. [default:  ]
	-n --name=<name>        Set speakers name.
```
