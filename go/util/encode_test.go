package util

import "testing"
import "github.com/stretchr/testify/assert"

func TestGetMd5Str(t *testing.T) {
	assertObj := assert.New(t)
	assertObj.Equal("25d55ad283aa400af464c76d713c07ad", GetMd5Str("12345678"), "12345678 to md5")
}
