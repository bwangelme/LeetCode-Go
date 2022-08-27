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

//TestNewCase2022
// 2022-08-28 题目发生了变化, 指数支持 E 和 e
func TestNewCase2022(t *testing.T) {
	for _, tt := range []struct {
		numberStr string
		want      bool
	}{
		{"1E9", true},
		{".2E5", true},
	} {
		assert.Equal(t, tt.want, isNumber(tt.numberStr), tt.numberStr)
	}
}
