package main

import (
	"bytes"
	"io"
	"os"
	"strings"
	"testing"

	"github.com/jiro4989/textchat/config"
	"github.com/stretchr/testify/assert"
)

func TestMain(t *testing.T) {
	type TestData struct {
		desc   string
		config config.Config
		expect []string
	}
	tds := []TestData{
		{
			desc: "Left",
			config: config.Config{
				Width: 80,
				Icon:  "testdata/block.txt",
				Words: []string{"こんにちは"},
				Pad:   " ",
			},
			expect: []string{
				".-----.  .--------------.                                                       ",
				"| 123 | <   こんにちは  |                                                       ",
				"| 456 |  `--------------'                                                       ",
				"| 789 |                                                                         ",
				"`-----'                                                                         ",
			},
		},
		{
			desc: "Right",
			config: config.Config{
				Width: 80,
				Icon:  "testdata/block.txt",
				Words: []string{"こんにちは"},
				Pad:   " ",
				Right: true,
			},
			expect: []string{
				"                                                       .--------------.  .-----.",
				"                                                       |  こんにちは   > | 123 |",
				"                                                       `--------------'  | 456 |",
				"                                                                         | 789 |",
				"                                                                         `-----'",
			},
		},
	}
	for _, v := range tds {
		r, w, err := os.Pipe()
		if err != nil {
			t.Error(err)
		}

		stdout := os.Stdout
		os.Stdout = w

		Main(v.config)

		os.Stdout = stdout
		w.Close()

		var buf bytes.Buffer
		io.Copy(&buf, r)

		s := buf.String()
		got := strings.Split(s, "\n")
		got = got[:len(got)-1]
		assert.Equal(t, v.expect, got, v.desc)
	}
}
