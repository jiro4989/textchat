package balloon

import (
	"fmt"
	"strings"

	"github.com/jiro4989/align/align"
	"github.com/mattn/go-runewidth"
)

const bodyPaddingWidth = 7

// height -1
func Balloon(aa []string, height int) []string {
	max := align.MaxStringWidth(aa)
	var ret []string

	blnWidth := max + 4
	border := strings.Repeat("-", blnWidth-2)
	top := fmt.Sprintf(".%s.", border)
	ret = append(ret, top)

	if height == -1 {
		height = len(aa)
	}

	for i := 0; i < height; i++ {
		line := strings.Repeat(" ", max)
		if i < len(aa) {
			line = aa[i]
		}

		diff := max - runewidth.StringWidth(line)
		pad := strings.Repeat(" ", diff)
		line = fmt.Sprintf("| %s%s |", line, pad)
		ret = append(ret, line)
	}

	bottom := fmt.Sprintf("`%s'", border)
	ret = append(ret, bottom)
	return ret
}

// Left returns left side balloon.
func Left(text string, width int) []string {
	return LeftSlice([]string{text}, width)
}

func LeftSlice(texts []string, width int) []string {
	// 横幅を取得
	// widthは吹き出しこみの幅なのに対し、
	// maxは吹き出しなしの幅である点に留意する
	max := align.MaxStringWidth(texts)
	if w := width - bodyPaddingWidth; w < max {
		max = w
	}

	// 上部の吹き出し
	var ret []string
	border := strings.Repeat("-", max+bodyPaddingWidth-3)
	top := fmt.Sprintf(" .%s.", border)
	ret = append(ret, top)

	// 1行ずつ吹き出しで囲う
	// 1行がmaxを超過していた場合は次の行に折り返して追加する
	for i, text := range texts {
		for j, t := range SplitRuneWidth(text, max) {
			var middle string
			if i+j == 0 {
				middle += "<   "
			} else {
				middle += " |  "
			}
			diff := max - runewidth.StringWidth(t)
			pad := strings.Repeat(" ", diff)
			middle += fmt.Sprintf("%s%s  |", t, pad)
			ret = append(ret, middle)
		}
	}

	// 低部の吹き出し
	bottom := fmt.Sprintf(" `%s'", border)
	ret = append(ret, bottom)
	return ret
}

// Right returns right side balloon.
func Right(text string, width int) []string {
	return RightSlice([]string{text}, width)
}

func RightSlice(texts []string, width int) []string {
	// 横幅を取得
	// widthは吹き出しこみの幅なのに対し、
	// maxは吹き出しなしの幅である点に留意する
	max := align.MaxStringWidth(texts)
	if w := width - bodyPaddingWidth; w < max {
		max = w
	}

	// 上部の吹き出し
	var ret []string
	border := strings.Repeat("-", max+bodyPaddingWidth-3)
	top := fmt.Sprintf(".%s. ", border)
	top = strings.Repeat(" ", width-runewidth.StringWidth(top)) + top
	ret = append(ret, top)

	// 1行ずつ吹き出しで囲う
	// 1行がmaxを超過していた場合は次の行に折り返して追加する
	for i, text := range texts {
		for j, t := range SplitRuneWidth(text, max) {
			var middle string
			diff := max - runewidth.StringWidth(t)
			pad := strings.Repeat(" ", diff)
			middle += fmt.Sprintf("|  %s%s", t, pad)
			if i+j == 0 {
				middle += "   >"
			} else {
				middle += "  | "
			}

			w := runewidth.StringWidth(middle)
			middle = strings.Repeat(" ", width-w) + middle
			ret = append(ret, middle)
		}
	}

	// 低部の吹き出し
	bottom := fmt.Sprintf("`%s' ", border)
	bottom = strings.Repeat(" ", width-runewidth.StringWidth(bottom)) + bottom
	ret = append(ret, bottom)

	return ret
}

func SplitRuneWidth(text string, width int) []string {
	if width < 1 || len(text) < 1 {
		return []string{}
	}

	if runewidth.StringWidth(text) <= width {
		return []string{text}
	}

	var cw int
	var ct string
	var texts []string
	for _, r := range []rune(text) {
		w := runewidth.RuneWidth(r)
		if width < cw+w {
			texts = append(texts, ct)
			ct = ""
			cw = 0
		}
		ct += string(r)
		cw += w
	}
	if 0 < len(ct) {
		texts = append(texts, ct)
	}
	return texts
}
