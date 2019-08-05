package balloon

import (
	"fmt"
	"strings"

	"github.com/mattn/go-runewidth"
)

const bodyPaddingWidth = 7

// Left returns left side balloon.
func Left(text string, width int) []string {
	// Substitue left padding (4) and right padding (3)
	bodyWidth := width - bodyPaddingWidth
	if textWidth := runewidth.StringWidth(text); textWidth <= bodyWidth {
		width = textWidth + bodyPaddingWidth
	}

	var bln []string

	// Substitue left padding (2) and right padding (1)
	lineWidth := width - 3
	top := " ." + strings.Repeat("-", lineWidth) + "."
	bln = append(bln, top)

	for i, t := range splitRuneWidth(text, bodyWidth) {
		var middle string
		if i == 0 {
			middle += "<   "
		} else {
			middle += " |  "
		}
		diff := width - bodyPaddingWidth - runewidth.StringWidth(t)
		pad := strings.Repeat(" ", diff)
		middle += fmt.Sprintf("%s%s  |", t, pad)
		bln = append(bln, middle)
	}

	bottom := " `" + strings.Repeat("-", lineWidth) + "'"
	bln = append(bln, bottom)

	return bln
}

// Right returns right side balloon.
func Right(text string, width int) []string {
	srcWidth := width

	// Substitue left padding (4) and right padding (3)
	bodyWidth := width - bodyPaddingWidth
	if textWidth := runewidth.StringWidth(text); textWidth <= bodyWidth {
		width = textWidth + bodyPaddingWidth
	}

	var bln []string

	// Substitue left padding (2) and right padding (1)
	lineWidth := width - 3
	pad := strings.Repeat(" ", srcWidth-width)
	top := pad + "." + strings.Repeat("-", lineWidth) + ". "
	bln = append(bln, top)

	for i, t := range splitRuneWidth(text, bodyWidth) {
		middle := pad
		diff := width - bodyPaddingWidth - runewidth.StringWidth(t)
		pad := strings.Repeat(" ", diff)
		middle += fmt.Sprintf("|  %s%s", t, pad)
		if i == 0 {
			middle += "   >"
		} else {
			middle += "  | "
		}
		bln = append(bln, middle)
	}

	bottom := pad + "`" + strings.Repeat("-", lineWidth) + "' "
	bln = append(bln, bottom)

	return bln
}

func splitRuneWidth(text string, width int) []string {
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
