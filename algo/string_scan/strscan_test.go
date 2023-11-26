package string_scan

import (
	"github.com/stretchr/testify/assert"
	"testing"
)

func Test_ScanStr(t *testing.T) {
	res := ScanInput(`[[0, 12], [12, 2]]`)
	assert.Equal(t, [][]int{{0, 12}, {12, 2}}, res)
}
