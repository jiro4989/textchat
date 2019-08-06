package icon

import (
	"testing"

	"github.com/jiro4989/textchat/config"
	"github.com/stretchr/testify/assert"
)

func TestAA(t *testing.T) {
	type TestData struct {
		desc   string
		config config.Config
		expect []string
	}
	tds := []TestData{
		{
			desc: "Simple text",
			config: config.Config{
				Icon: "../testdata/block.txt",
			},
			expect: []string{
				"123",
				"456",
				"789",
			},
		},
		{
			desc: "Simple text",
			config: config.Config{
				Name: "John",
				Icon: "../testdata/block.txt",
			},
			expect: []string{
				"John",
			},
		},
	}
	for _, v := range tds {
		got, err := AA(v.config)
		assert.Equal(t, v.expect, got, v.desc)
		assert.Nil(t, err)
	}

	// 異常系
	c := config.Config{
		Icon: "hogefuga",
	}
	_, err := AA(c)
	assert.NotNil(t, err)
}
