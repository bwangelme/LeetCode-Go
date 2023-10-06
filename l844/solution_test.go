package l844

import (
	"github.com/stretchr/testify/assert"
	"testing"
)

func Test_backspaceCompare(t *testing.T) {
	for _, tt := range []struct {
		s     string
		t     string
		equal bool
	}{
		{"ab#c", "ad#c", true},
		{"ab##", "c#d#", true},
		{"##", "##", true},
		{"a#c", "b", false},
	} {
		assert.Equal(t, tt.equal, backspaceCompare(tt.s, tt.t))
	}
}
