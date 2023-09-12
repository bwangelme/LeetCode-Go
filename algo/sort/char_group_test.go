package sort

import (
	"github.com/stretchr/testify/assert"
	"testing"
)

func Test_LetterGroup(t *testing.T) {
	for _, tt := range []struct {
		Input  []byte
		Output []byte
	}{
		{
			[]byte{'D', 'a', 'F', 'B', 'c', 'A', 'z'},
			[]byte{'z', 'a', 'c', 'B', 'F', 'A', 'D'},
		},
	} {
		LetterGroup(tt.Input, len(tt.Input))
		assert.Equal(t, tt.Output, tt.Input)
	}
}

func Test_DigitLetterGroup(t *testing.T) {
	for _, tt := range []struct {
		Input  []byte
		Output []byte
	}{
		{
			[]byte{'D', 'a', 'F', 'B', '7', 'c', '4', 'A', 'z', '3'},
			[]byte{'c', 'a', 'z', '4', '7', '3', 'B', 'A', 'F', 'D'},
		},
	} {
		DigitLetterGroup(tt.Input, len(tt.Input))
		assert.Equal(t, tt.Output, tt.Input)
	}
}
