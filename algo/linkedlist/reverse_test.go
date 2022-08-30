package linkedlist

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func Test_Reverse(t *testing.T) {
	dict := []string{"零", "壹", "贰", "叁", "肆", "伍", "陆", "柒", "捌", "玖"}

	head := NewList(dict)
	head = Reverse(head)

	res := PrintList(head)

	assert.Equal(t, "玖 -> 捌 -> 柒 -> 陆 -> 伍 -> 肆 -> 叁 -> 贰 -> 壹 -> 零", res)

}
