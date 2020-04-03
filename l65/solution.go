package l65

type State uint8
type Input uint8

func (i Input) String() string {
	strMap := map[Input]string{
		INVALID:  "INVALID",
		SPACE:    "SPACE",
		SIGN:     "SIGN",
		DOT:      "DOT",
		DIGIT:    "DIGIT",
		EXPONENT: "EXPONENT",
	}
	v, ok := strMap[i]
	if !ok {
		return "Not Exist"
	}
	return v
}

// https://passage-1253400711.cos-website.ap-beijing.myqcloud.com/2020-04-03-041334.png
// 关于状态机各个状态的说明
const (
	s0 State = iota
	s1
	s2
	_
	s4
	s5
	s6
	s7
	s8
	s9
	s10
	sf
)

const (
	INVALID Input = iota
	SPACE
	SIGN
	DOT
	DIGIT
	EXPONENT
)

var TransTable = map[State]map[Input]State{
	s0: {
		SPACE: s0, SIGN: s9, DIGIT: s1, DOT: s8, EXPONENT: sf,
	},
	s1: {
		SPACE: s4, SIGN: sf, DIGIT: s1, DOT: s2, EXPONENT: s5,
	},
	s2: {
		SPACE: s4, SIGN: sf, DIGIT: s2, DOT: sf, EXPONENT: s5,
	},
	s4: {
		SPACE: s4, SIGN: sf, DIGIT: sf, DOT: sf, EXPONENT: sf,
	},
	s5: {
		SPACE: sf, SIGN: s7, DIGIT: s6, DOT: sf, EXPONENT: sf,
	},
	s6: {
		SPACE: s4, SIGN: sf, DIGIT: s6, DOT: sf, EXPONENT: sf,
	},
	s7: {
		SPACE: sf, SIGN: sf, DIGIT: s6, DOT: sf, EXPONENT: sf,
	},
	s8: {
		SPACE: sf, SIGN: sf, DIGIT: s2, DOT: sf, EXPONENT: sf,
	},
	s9: {
		SPACE: sf, SIGN: sf, DIGIT: s1, DOT: s8, EXPONENT: sf,
	},
}

func isNumber(s string) bool {
	var state = s0
	var input = INVALID
	_ = s10

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
		case c == 'e':
			input = EXPONENT
		default:
			return false
		}
		state = TransTable[state][input]
		//fmt.Println(input, state)
		if state == sf {
			return false
		}
	}
	return state == s1 || state == s2 || state == s4 || state == s6
}
