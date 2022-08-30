package linkedlist

import (
	"fmt"
	"testing"

	"github.com/stretchr/testify/assert"
)

func Test_removeNthFromEnd(t *testing.T) {
	for _, tt := range []struct {
		arr    []string
		n      int
		wanted string
	}{
		//{
		//	[]string{"零", "壹", "贰", "叁", "肆", "伍", "陆", "柒", "捌", "玖"},
		//	3,
		//	"零 -> 壹 -> 贰 -> 叁 -> 肆 -> 伍 -> 陆 -> 捌 -> 玖",
		//},
		{
			[]string{"1"},
			1,
			"",
		},
	} {
		head := NewList(tt.arr)
		fmt.Println(head, head.Next)
		head = removeNthFromEnd(head, tt.n)
		fmt.Println(head)
		assert.Equal(t, tt.wanted, PrintList(head))
	}

}
