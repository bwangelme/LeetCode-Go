package l67

import (
	"github.com/stretchr/testify/assert"
	"testing"
)

func Test_addBinary(t *testing.T) {
	for _, tt := range []struct {
		A      string
		B      string
		Wanted string
	}{
		{"11", "1", "100"},
		{"1010", "1011", "10101"},
	} {
		assert.Equal(t, tt.Wanted, addBinary(tt.A, tt.B))
	}
}
