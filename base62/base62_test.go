package base62_test

import (
	"testing"

	"github.com/abs2free/tinyurl/store/util/base62"
	"github.com/stretchr/testify/assert"
)

func TestEncode(t *testing.T) {
	n := 89569285645
	output := base62.Encode(n)
	assert.Equal(t, "1ZlfarV", output)
}

func TestDecode(t *testing.T) {
	n := "1ZlfarV"
	output, err := base62.Decode(n)
	assert.NoError(t, err)
	assert.Equal(t, 89569285645, output)
}
