# tchat

tchat is a terminal chat cli.

## Installation

```bash
go get github.com/jiro4989/tchat
```

## Usage

```bash
tchat Hello
tchat -r Hello
tchat -i icon.txt こんにちは
tchat -i icon.txt -r こんにちは
tchat -i <(cat icon.txt) Hello world
seq 5 | tchat
tchat -w 100 -r Right
```
