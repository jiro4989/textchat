package chat

import (
	"fmt"

	"github.com/jiro4989/align/align"
	"github.com/jiro4989/textchat/balloon"
	"github.com/jiro4989/textchat/config"
	"github.com/mattn/go-runewidth"
)

func Left(aa, chatText []string, config config.Config) []string {
	// アイコンのAAを取得
	height := len(chatText)
	if height-2 < len(aa) {
		height = len(aa) + 2
	}
	aa = balloon.Balloon(aa, -1)

	var ret []string
	for i := 0; i < height; i++ {
		var left, right, t string
		if len(chatText) <= i {
			left = aa[i]
			t = align.AlignLeft([]string{left}, config.Width, config.Pad)[0]
		} else if len(aa) <= i {
			right = chatText[i]
			lw := runewidth.StringWidth(aa[0])
			rw := runewidth.StringWidth(right)
			left := align.AlignRight([]string{right}, lw+1+rw, config.Pad)
			t = align.AlignLeft(left, config.Width, config.Pad)[0]
		} else {
			left = aa[i]
			right = chatText[i]
			lr := fmt.Sprintf("%s %s", left, right)
			t = align.AlignLeft([]string{lr}, config.Width, config.Pad)[0]
		}
		ret = append(ret, t)
	}
	return ret
}

func Right(aa, chatText []string, config config.Config) []string {
	// アイコンのAAを取得
	height := len(chatText)
	if height-2 < len(aa) {
		height = len(aa) + 2
	}
	aa = balloon.Balloon(aa, -1)

	var ret []string
	for i := 0; i < height; i++ {
		var left, right, t string
		if len(chatText) <= i {
			right = aa[i]
			t = align.AlignRight([]string{right}, config.Width, config.Pad)[0]
		} else if len(aa) <= i {
			left = chatText[i]
			lw := runewidth.StringWidth(left)
			rw := runewidth.StringWidth(aa[0])
			right := align.AlignLeft([]string{left}, lw+1+rw, config.Pad)
			t = align.AlignRight(right, config.Width, config.Pad)[0]
		} else {
			left = chatText[i]
			right = aa[i]
			lr := fmt.Sprintf("%s %s", left, right)
			t = align.AlignRight([]string{lr}, config.Width, config.Pad)[0]
		}
		ret = append(ret, t)
	}
	return ret
}
