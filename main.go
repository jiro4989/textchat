package main

import (
	"fmt"
	"strconv"
	"strings"

	"github.com/docopt/docopt-go"
	"github.com/jiro4989/align/align"
	"github.com/jiro4989/tchat/balloon"
	"github.com/jiro4989/tchat/icon"
)

const doc = `tchat is a terminal chat cli.

Usage:
	tchat
	tchat [options] <word>...
	tchat -h | --help
	tchat -v | --version

Options:
	-h --help               Show this screen.
	-v --version            Show version.
	-r --right              Say word on right.
	-i --icon=<textfile>    Icon AA file.
	-w --width=<width>      Display width. [default: 80]`

func main() {
	opts, _ := docopt.ParseDoc(doc)
	if ok := opts["--version"].(bool); ok {
		fmt.Println(Version)
		return
	}

	// 文字の指定がない時はok がfalseになる
	iconfile, _ := opts["--icon"].(string)
	aa, err := icon.AA(iconfile)
	if err != nil {
		panic(err)
	}

	widthStr, _ := opts["--width"].(string)
	width, err := strconv.Atoi(widthStr)
	if err != nil {
		panic(err)
	}

	var chatText []string
	words, ok := opts["<word>"].([]string)
	if !ok {
		panic("error TODO")
	}
	if len(words) < 1 {
		lines := readStdin()
		chatText = balloon.LeftSlice(lines, width)
	} else {
		text := strings.Join(words, " ")
		fmt.Println(text, width)
		chatText = balloon.Left(text, width)
	}

	max := align.MaxStringWidth(aa)
	top := strings.Repeat(" ", max)
	aa = append([]string{top}, aa...)
	aa = append(aa, top)

	for _, text := range aa {
		fmt.Println(text)
	}

	for _, text := range chatText {
		fmt.Println(text)
	}
}
