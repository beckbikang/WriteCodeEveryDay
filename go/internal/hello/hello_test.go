package hello

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestHello(t *testing.T) {
	// 测试正常情况
	result := Hello("World")
	assert.Equal(t, "Hello, World!", result)

	// 测试空字符串
	result = Hello("")
	assert.Equal(t, "Hello, !", result)

	// 测试中文名称
	result = Hello("张三")
	assert.Equal(t, "Hello, 张三!", result)
}
