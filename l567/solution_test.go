package l567

import (
	"github.com/stretchr/testify/assert"
	"testing"
)

func Test_checkInclusion(t *testing.T) {
	for _, tt := range []struct {
		S1     string
		S2     string
		Wanted bool
	}{
		{"ab", "eidbaooo", true},
		{"ab", "eidboaoo", false},
	} {
		assert.Equal(t, tt.Wanted, checkInclusion(tt.S1, tt.S2))
	}
}
