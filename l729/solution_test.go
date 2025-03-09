package l729

import (
	"github.com/stretchr/testify/assert"
	"testing"
)

func Test_MyCalendar(t *testing.T) {
	c := Constructor()
	assert.Equal(t, true, c.Book(10, 20))
	assert.Equal(t, false, c.Book(15, 25))
	assert.Equal(t, true, c.Book(20, 30))
}
