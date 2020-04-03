package l65

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestNumberStr(t *testing.T) {
	for _, tt := range []struct {
		numberStr string
		want      bool
	}{
		{" 3.", true},
		{"+3.", true},
		{"+.8", true},
		{".2", true},
		{".2E5", false},  // 指数只支持 e 形式
		{".2e5", true},
		{"46.e3", true},
		{"46.e-3", true},
		{"46.e-", false},
		{"0.e", false},
		{".", false},
		{".1.", false},
		{" 99e2.5", false},
		{"0", true},
		{" 0.1", true},
		{"abc", false},
		{"1 a", false},
		{"2e10", true},
		{" -90e3", true},
		{" 1e", false},
		{"e3", false},
		{" 6e-1", true},
		{"53.5e93", true},
		{" --6", false},
		{"-+3", false},
		{"95a54e53", false},
		{"6e-1-1", false},
	} {
		assert.Equal(t, tt.want, isNumber(tt.numberStr), tt.numberStr)
	}
}

func TestSpecial(t *testing.T) {
	for _, tt := range []struct {
		numberStr string
		want      bool
	}{
		{"46.e-3", true},
	} {
		assert.Equal(t, tt.want, isNumber(tt.numberStr), tt.numberStr)
	}
}
