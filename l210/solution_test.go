package l210

import (
	"github.com/stretchr/testify/assert"
	"testing"
)

func Test_findOrder(t *testing.T) {
	for _, tt := range []struct {
		NumCourses    int
		Prerequisites [][]int
		Wanted        []int
	}{
		{
			2,
			[][]int{{1, 0}},
			[]int{0, 1},
		},
	} {
		assert.Equal(t, tt.Wanted, findOrder(tt.NumCourses, tt.Prerequisites))
	}
}
