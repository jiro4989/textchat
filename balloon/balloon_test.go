package balloon

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestBalloon(t *testing.T) {
	type TestData struct {
		desc   string
		texts  []string
		height int
		expect []string
	}
	tds := []TestData{
		{
			desc: "Simple text",
			texts: []string{
				"太郎",
			},
			height: -1,
			expect: []string{
				".------.",
				"| 太郎 |",
				"`------'",
			},
		},
		{
			desc: "Simple text",
			texts: []string{
				"太郎",
				"a",
			},
			height: -1,
			expect: []string{
				".------.",
				"| 太郎 |",
				"| a    |",
				"`------'",
			},
		},
		{
			desc: "Simple text",
			texts: []string{
				"太郎",
			},
			height: 2,
			expect: []string{
				".------.",
				"| 太郎 |",
				"|      |",
				"`------'",
			},
		},
	}
	for _, v := range tds {
		got := Balloon(v.texts, v.height)
		assert.Equal(t, v.expect, got, v.desc)
	}
}

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
			desc:  "Wrap border line",
			text:  "012345678901234567890123456789012",
			width: 40,
			expect: []string{
				" .-------------------------------------.",
				"<   012345678901234567890123456789012  |",
				" `-------------------------------------'",
			},
		},
		{
			desc:  "Wrap text",
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
			desc:  "Multibyte string",
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

func TestLeftSlice(t *testing.T) {
	type TestData struct {
		desc   string
		texts  []string
		width  int
		expect []string
	}
	tds := []TestData{
		{
			desc: "Simple text",
			texts: []string{
				"a",
				"bb",
				"ccc",
			},
			width: 40,
			expect: []string{
				" .-------.",
				"<   a    |",
				" |  bb   |",
				" |  ccc  |",
				" `-------'",
			},
		},
		{
			desc: "Simple text",
			texts: []string{
				"a",
				"bb",
				"あい",
			},
			width: 40,
			expect: []string{
				" .--------.",
				"<   a     |",
				" |  bb    |",
				" |  あい  |",
				" `--------'",
			},
		},
		{
			desc: "Simple text",
			texts: []string{
				"あい",
				"a",
				"bb",
			},
			width: 40,
			expect: []string{
				" .--------.",
				"<   あい  |",
				" |  a     |",
				" |  bb    |",
				" `--------'",
			},
		},
		{
			desc: "Simple text",
			texts: []string{
				"あいうえお",
			},
			width: 40,
			expect: []string{
				" .--------------.",
				"<   あいうえお  |",
				" `--------------'",
			},
		},
	}
	for _, v := range tds {
		got := LeftSlice(v.texts, v.width)
		assert.Equal(t, v.expect, got, v.desc)
	}
}

func TestRight(t *testing.T) {
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
				"                      .---------------. ",
				"                      |  Hello world   >",
				"                      `---------------' ",
			},
		},
		{
			desc:  "Wrap border line",
			text:  "012345678901234567890123456789012",
			width: 40,
			expect: []string{
				".-------------------------------------. ",
				"|  012345678901234567890123456789012   >",
				"`-------------------------------------' ",
			},
		},
		{
			desc:  "Wrap text",
			text:  "0123456789012345678901234567890123",
			width: 40,
			expect: []string{
				".-------------------------------------. ",
				"|  012345678901234567890123456789012   >",
				"|  3                                  | ",
				"`-------------------------------------' ",
			},
		},
		{
			desc:  "Multibyte string",
			text:  "あいうえお",
			width: 40,
			expect: []string{
				"                       .--------------. ",
				"                       |  あいうえお   >",
				"                       `--------------' ",
			},
		},
	}
	for _, v := range tds {
		got := Right(v.text, v.width)
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

func TestRightSlice(t *testing.T) {
	type TestData struct {
		desc   string
		texts  []string
		width  int
		expect []string
	}
	tds := []TestData{
		{
			desc: "Simple text",
			texts: []string{
				"a",
				"bb",
				"ccc",
			},
			width: 40,
			expect: []string{
				"                              .-------. ",
				"                              |  a     >",
				"                              |  bb   | ",
				"                              |  ccc  | ",
				"                              `-------' ",
			},
		},
		{
			desc: "Simple text",
			texts: []string{
				"a",
				"bb",
				"あい",
			},
			width: 40,
			expect: []string{
				"                             .--------. ",
				"                             |  a      >",
				"                             |  bb    | ",
				"                             |  あい  | ",
				"                             `--------' ",
			},
		},
		{
			desc: "Simple text",
			texts: []string{
				"あい",
				"a",
				"bb",
			},
			width: 40,
			expect: []string{
				"                             .--------. ",
				"                             |  あい   >",
				"                             |  a     | ",
				"                             |  bb    | ",
				"                             `--------' ",
			},
		},
		{
			desc: "Simple text",
			texts: []string{
				"あいうえお",
			},
			width: 40,
			expect: []string{
				"                       .--------------. ",
				"                       |  あいうえお   >",
				"                       `--------------' ",
			},
		},
	}
	for _, v := range tds {
		got := RightSlice(v.texts, v.width)
		assert.Equal(t, v.expect, got, v.desc)
	}
}
