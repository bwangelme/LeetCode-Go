package l20

type Stack []rune

func NewStack() Stack {
	return Stack{}
}

func (s *Stack) Len() int {
	return len(*s)
}

func (s *Stack) Push(elem rune) bool {
	*s = append(*s, elem)

	return true
}

func (s *Stack) Pop() *rune {
	stackLen := len(*s)
	if stackLen == 0 {
		return nil
	}

	// 注意: *s 必须扩起来，[] 的优先级高于 *
	top := (*s)[stackLen-1]
	*s = (*s)[:stackLen-1]

	return &top
}

var (
	parentheses = map[rune]rune{
		'(': ')',
		'[': ']',
		'{': '}',
	}
)

func isValid(s string) bool {
	stack := NewStack()

	for _, c := range s {
		switch c {
		case '(', '[', '{':
			stack.Push(c)
		case ')', ']', '}':
			if stack.Len() == 0 {
				return false
			}

			saveChar := stack.Pop()
			expectedChar, exist := parentheses[*saveChar]
			if !exist {
				return false
			}
			if c != expectedChar {
				return false
			}
		}
	}

	return stack.Len() == 0
}
