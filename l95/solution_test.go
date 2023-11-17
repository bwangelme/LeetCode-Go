package l95

import (
	"fmt"
	"github.com/bwangelme/LeetCode-Go/lt"
	"github.com/stretchr/testify/assert"
	"os"
	"testing"
)

func Test_generateTrees(t *testing.T) {
	res := generateTrees(3)
	assert.Equal(t, 5, len(res))
	for _, tree := range res {
		fmt.Println("============")
		lt.PrintTree(tree, os.Stdout)
	}
	assert.Fail(t, "Stop")
}

func Test_genTreeCache(t *testing.T) {
	res := genTreeCache(3)
	fmt.Println(len(res))
	for _, tree := range res {
		fmt.Println("============")
		lt.PrintTree(tree, os.Stdout)
	}
}
