# textchat

textchat is a terminal chat cli.

## Installation

```bash
go get github.com/jiro4989/textchat
```

## Usage

```bash
textchat Hello
textchat -r Hello
textchat -i icon.txt こんにちは
textchat -i icon.txt -r こんにちは
textchat -i <(cat icon.txt) Hello world
seq 5 | textchat
textchat -w 100 -r Right
```
