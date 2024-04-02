package lt

import (
	"fmt"
	"strings"
)

type Queue[T any] []T

func NewQueue[T any]() Queue[T] {
	return Queue[T]{}
}

func (q *Queue[T]) Peek() *T {
	if q.Len() <= 0 {
		return nil
	}

	v := (*q)[0]

	return &v
}

func (q *Queue[T]) Len() int {
	return len(*q)
}

func (q *Queue[T]) Poll() *T {
	if q.Len() <= 0 {
		return nil
	}

	top := (*q)[0]
	*q = (*q)[1:]

	return &top
}

func (q *Queue[T]) Offer(elem T) {
	*q = append(*q, elem)
}

func (q *Queue[T]) String() string {
	if q.Len() <= 0 {
		return "Empty"
	}

	var res = strings.Builder{}
	for i := 0; i < q.Len(); i++ {
		res.WriteString(fmt.Sprintf("%s ", (*q)[i]))
	}

	return res.String()
}
