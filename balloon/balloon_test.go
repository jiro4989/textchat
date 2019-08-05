package balloon

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestLeft(t *testing.T) {
	type TestData struct {
		desc   string
		text   string
		width  int
		expect []string
	}
	tds := []TestData{
		{
			desc:  "Simple text",
			text:  "Hello world",
			width: 40,
			expect: []string{
				" .---------------.",
				"<   Hello world  |",
				" `---------------'",
			},
		},
		{
			desc:  "Simple text",
			text:  "012345678901234567890123456789012",
			width: 40,
			expect: []string{
				" .-------------------------------------.",
				"<   012345678901234567890123456789012  |",
				" `-------------------------------------'",
			},
		},
		{
			desc:  "Simple text",
			text:  "0123456789012345678901234567890123",
			width: 40,
			expect: []string{
				" .-------------------------------------.",
				"<   012345678901234567890123456789012  |",
				" |  3                                  |",
				" `-------------------------------------'",
			},
		},
		{
			desc:  "Simple text",
			text:  "あいうえお",
			width: 40,
			expect: []string{
				" .--------------.",
				"<   あいうえお  |",
				" `--------------'",
			},
		},
	}
	for _, v := range tds {
		got := Left(v.text, v.width)
		assert.Equal(t, v.expect, got, v.desc)
	}
}

func TestSplitRuneWidth(t *testing.T) {
	type TestData struct {
		desc   string
		text   string
		width  int
		expect []string
	}
	tds := []TestData{
		{
			desc:  "Simple text",
			text:  "HelloWorld",
			width: 5,
			expect: []string{
				"Hello",
				"World",
			},
		},
		{
			desc:  "Simple text",
			text:  "こんにちは",
			width: 5,
			expect: []string{
				"こん",
				"にち",
				"は",
			},
		},
		{
			desc:   "Simple text",
			text:   "あ",
			width:  2,
			expect: []string{"あ"},
		},
		{
			desc:   "Simple text",
			text:   "",
			width:  2,
			expect: []string{},
		},
		{
			desc:   "Simple text",
			text:   "aiueo",
			width:  0,
			expect: []string{},
		},
	}
	for _, v := range tds {
		got := splitRuneWidth(v.text, v.width)
		assert.Equal(t, v.expect, got, v.desc)
	}
}
