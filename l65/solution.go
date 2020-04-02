package l65

type State uint8
type Input uint8

func (i Input) String() string {
	strMap := map[Input]string{
		INVALID:  "INVALID",
		SPACE:    "SPACE",
		SIGN:     "SIGN",
		DIGIT:    "DIGIT",
		DOT:      "DOT",
		EXPONENT: "EXPONENT",
	}
	v, ok := strMap[i]
	if !ok {
		return "Not Exist"
	}
	return v
}

const (
	s0 State = iota // 初始无输入或只输入了空格
	s1              // 输入了数字后的状态(这个数字中可能有符号和点)
	s2              // 前面无数字，只输入了一个点
	s3              // 前面无数字，只输入了一个符号
	s4              // 前面包含数字和点
	s5              // 1或4状态后输入了指数符号 E
	s6              // 5状态后输入了符号
	s7              // 5或6状态后输入了数字
	s8              // 前面有有效输入后，输入了空格
	sfailed
)

const (
	INVALID Input = iota
	SPACE
	SIGN
	DIGIT
	DOT
	EXPONENT
)

var TransTable = map[State]map[Input]State{
	s0: {
		SPACE: s0, SIGN: s3, DOT: s2, DIGIT: s1, EXPONENT: sfailed,
	},
	s1: {
		SPACE: s8, SIGN: sfailed, DOT: s4, DIGIT: s1, EXPONENT: s5,
	},
	s2: {
		SPACE: sfailed, SIGN: sfailed, DOT: sfailed, DIGIT: s1, EXPONENT: sfailed,
	},
	s3: {
		SPACE: sfailed, SIGN: sfailed, DOT: s2, DIGIT: s1, EXPONENT: sfailed,
	},
	s4: {
		SPACE: s8, SIGN: sfailed, DOT: sfailed, DIGIT: s4, EXPONENT: s5,
	},
	s5: {
		SPACE: sfailed, SIGN: s6, DOT: sfailed, DIGIT: s7, EXPONENT: sfailed,
	},
	s6: {
		SPACE: sfailed, SIGN: sfailed, DOT: sfailed, DIGIT: s7, EXPONENT: sfailed,
	},
	s7: {
		SPACE: s8, SIGN: sfailed, DOT: sfailed, DIGIT: s7, EXPONENT: sfailed,
	},
	s8: {
		SPACE: s8, SIGN: sfailed, DOT: sfailed, DIGIT: sfailed, EXPONENT: sfailed,
	},
}

func isNumber(s string) bool {
	var state = s0
	var input = INVALID

	for _, c := range s {
		switch {
		case c == ' ':
			input = SPACE
		case c == '+' || c == '-':
			input = SIGN
		case c >= '0' && c <= '9':
			input = DIGIT
		case c == '.':
			input = DOT
		case c == 'e' || c == 'E':
			input = EXPONENT
		default:
			return false
		}
		state = TransTable[state][input]
		//fmt.Println(input, state)
		if state == sfailed {
			return false
		}
	}
	return state == s1 || state == s4 || state == s7 || state == s8
}
