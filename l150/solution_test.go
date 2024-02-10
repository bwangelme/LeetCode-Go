package l150

import (
	"github.com/stretchr/testify/assert"
	"testing"
)

func Test_evalRPN(t *testing.T) {
	for _, tt := range []struct {
		Input  []string
		Wanted int
	}{
		{[]string{"2", "1", "+", "3", "*"}, 9},
		{[]string{"4", "13", "5", "/", "+"}, 6},
	} {
		assert.Equal(t, tt.Wanted, evalRPN(tt.Input))
	}
}
