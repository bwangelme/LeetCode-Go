package l310

import (
	"github.com/stretchr/testify/assert"
	"testing"
)

func Test_findMinHeightTrees(t *testing.T) {
	assert.Equal(t, []int{0}, findMinHeightTrees(1, [][]int{}))
	assert.Equal(t, []int{1}, findMinHeightTrees(4, [][]int{{1, 0}, {1, 2}, {1, 3}}))
	// 返回的结果因为遍历 map 的原因，顺序不确定
	//assert.Equal(t, []int{3, 4}, findMinHeightTrees(6, [][]int{{3, 0}, {3, 1}, {3, 2}, {3, 4}, {5, 4}}))
	assert.Equal(t, []int{6}, findMinHeightTrees(7, [][]int{{3, 0}, {3, 1}, {3, 2}, {5, 4}, {6, 3}, {6, 4}}))
}
