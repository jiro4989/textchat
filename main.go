package main

import (
	"fmt"
	"strings"

	"github.com/docopt/docopt-go"
	"github.com/jiro4989/align/align"
	"github.com/jiro4989/textchat/balloon"
	"github.com/jiro4989/textchat/icon"
	"github.com/mattn/go-runewidth"
)

type Config struct {
	Version bool     `docopt:"--version"`
	Right   bool     `docopt:"--right"`
	Icon    string   `docopt:"--icon"`
	Width   int      `docopt:"--width"`
	Words   []string `docopt:"<word>"`
}

const doc = `textchat is a terminal chat cli.

Usage:
	textchat [options]
	textchat [options] <word>...
	textchat -h | --help
	textchat -v | --version

Options:
	-h --help               Show this screen.
	-v --version            Show version.
	-r --right              Say word on right.
	-i --icon=<textfile>    Icon AA file.
	-w --width=<width>      Display width. [default: 80]`

func main() {
	opts, _ := docopt.ParseDoc(doc)
	var config Config
	if err := opts.Bind(&config); err != nil {
		panic(err)
	}

	if config.Version {
		fmt.Println(Version)
		return
	}

	aa, err := icon.AA(config.Icon)
	if err != nil {
		panic(err)
	}

	width := config.Width

	max := align.MaxStringWidth(aa)
	width -= max + 5 // 左右 + パディング

	var chatText []string
	words := config.Words
	if len(words) < 1 {
		lines := readStdin()
		if config.Right {
			chatText = balloon.RightSlice(lines, width)
		} else {
			chatText = balloon.LeftSlice(lines, width)
		}
	} else {
		text := strings.Join(words, " ")
		if config.Right {
			chatText = balloon.Right(text, width)
		} else {
			chatText = balloon.Left(text, width)
		}
	}

	height := len(chatText)
	if height-2 < len(aa) {
		height = len(aa) + 2
	}
	aa = balloon.Balloon(aa, -1)

	for i := 0; i < height; i++ {
		var left, right string
		if config.Right {
			if len(chatText) <= i {
				right = aa[i]
				diff := width
				left = strings.Repeat(" ", diff)
			} else if len(aa) <= i {
				left = chatText[i]
				diff := width - runewidth.StringWidth(left)
				right = strings.Repeat(" ", diff)
			} else {
				left = chatText[i]
				right = aa[i]
			}
		} else {
			if len(chatText) <= i {
				left = aa[i]
				diff := width - runewidth.StringWidth(left)
				right = strings.Repeat(" ", diff)
			} else if len(aa) <= i {
				right = chatText[i]
				diff := runewidth.StringWidth(aa[0])
				left = strings.Repeat(" ", diff)
			} else {
				left = aa[i]
				right = chatText[i]
			}

			w := width - runewidth.StringWidth(" "+right)
			pad := strings.Repeat(" ", w)
			right += pad
		}
		t := fmt.Sprintf("%s %s", left, right)
		fmt.Println(t)
	}
}
