package chat

import (
	"testing"

	"github.com/jiro4989/textchat/config"
	"github.com/stretchr/testify/assert"
)

func TestLeft(t *testing.T) {
	type TestData struct {
		desc     string
		aa       []string
		chatText []string
		config   config.Config
		expect   []string
	}
	tds := []TestData{
		{
			desc: "Simple text",
			aa: []string{
				"太郎",
			},
			chatText: []string{
				"--------------",
				"| こんにちは |",
				"--------------",
			},
			config: config.Config{
				Width: 30,
				Pad:   " ",
			},
			expect: []string{
				".------. --------------       ",
				"| 太郎 | | こんにちは |       ",
				"`------' --------------       ",
			},
		},
		{
			desc: "Simple text",
			aa: []string{
				"太郎",
			},
			chatText: []string{
				"--------------",
				"| こんにちは |",
				"--------------",
			},
			config: config.Config{
				Width: 30,
				Pad:   "　",
			},
			expect: []string{
				".------. -------------- 　　　",
				"| 太郎 | | こんにちは | 　　　",
				"`------' -------------- 　　　",
			},
		},
		{
			desc: "Simple text",
			aa: []string{
				"太郎",
				"山田",
			},
			chatText: []string{
				"--------------",
				"| こんにちは |",
				"--------------",
			},
			config: config.Config{
				Width: 30,
				Pad:   "　",
			},
			expect: []string{
				".------. -------------- 　　　",
				"| 太郎 | | こんにちは | 　　　",
				"| 山田 | -------------- 　　　",
				"`------'　　　　　　　　　　　",
			},
		},
		// {
		// 	desc: "TODO 全角の位置",
		// 	aa: []string{
		// 		"太郎",
		// 	},
		// 	chatText: []string{
		// 		"--------------",
		// 		"| こんにちは |",
		// 		"| こんばんは |",
		// 		"--------------",
		// 	},
		// 	config: config.Config{
		// 		Width: 30,
		// 		Pad:   "　",
		// 	},
		// 	expect: []string{
		// 		".------. -------------- 　　　",
		// 		"| 太郎 | | こんにちは | 　　　",
		// 		"`------' | こんばんは | 　　　",
		// 		"　　　　 -------------- 　　　",
		// 	},
		// },
	}
	for _, v := range tds {
		got := Left(v.aa, v.chatText, v.config)
		assert.Equal(t, v.expect, got, v.desc)
	}
}

func TestRight(t *testing.T) {
	type TestData struct {
		desc     string
		aa       []string
		chatText []string
		config   config.Config
		expect   []string
	}
	tds := []TestData{
		{
			desc: "Simple text",
			aa: []string{
				"太郎",
			},
			chatText: []string{
				"--------------",
				"| こんにちは |",
				"--------------",
			},
			config: config.Config{
				Width: 30,
				Pad:   " ",
			},
			expect: []string{
				"       -------------- .------.",
				"       | こんにちは | | 太郎 |",
				"       -------------- `------'",
			},
		},
		{
			desc: "Simple text",
			aa: []string{
				"太郎",
			},
			chatText: []string{
				"--------------",
				"| こんにちは |",
				"--------------",
			},
			config: config.Config{
				Width: 30,
				Pad:   "　",
			},
			expect: []string{
				"　　　 -------------- .------.",
				"　　　 | こんにちは | | 太郎 |",
				"　　　 -------------- `------'",
			},
		},
		{
			desc: "Simple text",
			aa: []string{
				"太郎",
				"山田",
			},
			chatText: []string{
				"--------------",
				"| こんにちは |",
				"--------------",
			},
			config: config.Config{
				Width: 30,
				Pad:   "　",
			},
			expect: []string{
				"　　　 -------------- .------.",
				"　　　 | こんにちは | | 太郎 |",
				"　　　 -------------- | 山田 |",
				"　　　　　　　　　　　`------'",
			},
		},
		// {
		// 	desc: "TODO 全角の位置",
		// 	aa: []string{
		// 		"太郎",
		// 	},
		// 	chatText: []string{
		// 		"--------------",
		// 		"| こんにちは |",
		// 		"| こんばんは |",
		// 		"--------------",
		// 	},
		// 	config: config.Config{
		// 		Width: 30,
		// 		Pad:   "　",
		// 	},
		// 	expect: []string{
		// 		".------. -------------- 　　　",
		// 		"| 太郎 | | こんにちは | 　　　",
		// 		"`------' | こんばんは | 　　　",
		// 		"　　　　 -------------- 　　　",
		// 	},
		// },
	}
	for _, v := range tds {
		got := Right(v.aa, v.chatText, v.config)
		assert.Equal(t, v.expect, got, v.desc)
	}
}
