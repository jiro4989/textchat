package main

import (
	"fmt"
	"strings"

	"github.com/docopt/docopt-go"
	"github.com/jiro4989/align/align"
	"github.com/jiro4989/textchat/balloon"
	"github.com/jiro4989/textchat/chat"
	"github.com/jiro4989/textchat/config"
	"github.com/jiro4989/textchat/icon"
)

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
	-w --width=<width>      Display width. [default: 80]
	-p --pad=<pad>          Pad string. [default:  ]`

func main() {
	opts, _ := docopt.ParseDoc(doc)
	var config config.Config
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

	// チャットのテキストを取得t
	max := align.MaxStringWidth(aa)
	width := config.Width - max - 5
	chatText := ChatText(width, config)

	var lines []string
	if config.Right {
		lines = chat.Right(aa, chatText, config)
	} else {
		lines = chat.Left(aa, chatText, config)
	}
	for _, line := range lines {
		fmt.Println(line)
	}
}

func ChatText(width int, config config.Config) []string {
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
	return chatText
}
