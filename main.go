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

	// チャットのテキストを取得t
	var chatText []string
	var lines []string
	if words := config.Words; len(words) < 1 {
		lines = readStdin()
	} else {
		lines = []string{strings.Join(words, " ")}
	}
	if config.Right {
		chatText = balloon.RightSlice(lines, width)
	} else {
		chatText = balloon.LeftSlice(lines, width)
	}

	// アイコンのAAを取得
	height := len(chatText)
	if height-2 < len(aa) {
		height = len(aa) + 2
	}
	aa = balloon.Balloon(aa, -1)

	const whiteSpace = " "

	for i := 0; i < height; i++ {
		var left, right, t string
		if config.Right {
			if len(chatText) <= i {
				right = aa[i]
				t = align.AlignRight([]string{right}, config.Width, whiteSpace)[0]
			} else if len(aa) <= i {
				left = chatText[i]
				lw := runewidth.StringWidth(left)
				rw := runewidth.StringWidth(aa[0])
				right := align.AlignLeft([]string{left}, lw+1+rw, whiteSpace)
				t = align.AlignRight(right, config.Width, whiteSpace)[0]
			} else {
				left = chatText[i]
				right = aa[i]
				lr := fmt.Sprintf("%s %s", left, right)
				t = align.AlignRight([]string{lr}, config.Width, whiteSpace)[0]
			}
		} else {
			if len(chatText) <= i {
				left = aa[i]
				t = align.AlignLeft([]string{left}, config.Width, whiteSpace)[0]
			} else if len(aa) <= i {
				right = chatText[i]
				lw := runewidth.StringWidth(aa[0])
				rw := runewidth.StringWidth(right)
				left := align.AlignRight([]string{right}, lw+1+rw, whiteSpace)
				t = align.AlignLeft(left, config.Width, whiteSpace)[0]
			} else {
				left = aa[i]
				right = chatText[i]
				lr := fmt.Sprintf("%s %s", left, right)
				t = align.AlignLeft([]string{lr}, config.Width, whiteSpace)[0]
			}
		}
		fmt.Println(t)
	}
}
