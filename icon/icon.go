package icon

import (
	"io/ioutil"
	"strings"

	"github.com/jiro4989/align/align"
	"github.com/jiro4989/textchat/config"
	"github.com/mattn/go-gimei"
)

func AA(config config.Config) ([]string, error) {
	if name := config.Name; name != "" {
		return []string{name}, nil
	}

	file := config.Icon
	if file == "" {
		name := gimei.NewName().Last.Kanji()
		return []string{name}, nil
	}

	b, err := ioutil.ReadFile(file)
	if err != nil {
		return nil, err
	}

	text := string(b)
	lines := strings.Split(text, "\n")
	lines = lines[:len(lines)-1]
	aa := align.AlignLeft(lines, -1, " ")
	return aa, nil
}
