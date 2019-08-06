package icon

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestAA(t *testing.T) {
	type TestData struct {
		desc   string
		file   string
		expect []string
	}
	tds := []TestData{
		{
			desc: "Simple text",
			file: "../testdata/block.txt",
			expect: []string{
				"123",
				"456",
				"789",
			},
		},
	}
	for _, v := range tds {
		got, err := AA(v.file)
		assert.Equal(t, v.expect, got, v.desc)
		assert.Nil(t, err)
	}

	// 異常系
	_, err := AA("hogefuga")
	assert.NotNil(t, err)
}
