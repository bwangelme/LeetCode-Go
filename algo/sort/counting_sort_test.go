package sort

import (
	"github.com/stretchr/testify/assert"
	"testing"
)

func Test_CountingSort(t *testing.T) {
	for _, tt := range []struct {
		Input  []int64
		Wanted []int64
	}{
		{
			Input:  []int64{2, 5, 3, 0, 2, 3, 0, 3},
			Wanted: []int64{0, 0, 2, 2, 3, 3, 3, 5},
		},
	} {

		assert.Equal(t, tt.Wanted, CountingSort(tt.Input, int64(len(tt.Input))))
	}
}
