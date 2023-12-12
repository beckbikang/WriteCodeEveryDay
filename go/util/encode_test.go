package util

import (
	"strconv"
	"testing"
)
import "github.com/stretchr/testify/assert"

func TestGetMd5Str(t *testing.T) {
	assertObj := assert.New(t)
	assertObj.Equal("25d55ad283aa400af464c76d713c07ad", GetMd5Str("12345678"), "12345678 to md5")
}
func TestGetMd5Str2(t *testing.T) {
	assertObj := assert.New(t)
	assertObj.Equal("25d55ad283aa400af464c76d713c07ad", GetMd5Str2("12345678"), "12345678 to md5")
}

func BenchmarkGetMd5Str(b *testing.B) {
	for i := 0; i < b.N; i++ {
		GetMd5Str(strconv.Itoa(i))
	}
}

func BenchmarkGetMd5Str2(b *testing.B) {
	for i := 0; i < b.N; i++ {
		GetMd5Str2(strconv.Itoa(i))
	}
}
