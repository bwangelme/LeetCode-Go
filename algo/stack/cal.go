package stack

import (
	"fmt"
	"math"
	"strconv"
	"unicode"

	"github.com/pkg/errors"
)

var priorityMap = map[string]int{
	"+": 0,
	"-": 0,
	"*": 1,
	"/": 1,
	"^": 2,
}

type OP struct {
	op string
	//priority 运算符的优先级
	priority int
}

func (o OP) String() string {
	return fmt.Sprintf("%s", o.op)
}

func NewOP(op string) (*OP, error) {
	pri, ok := priorityMap[op]
	if !ok {
		return nil, fmt.Errorf("invalid operator %v", pri)
	}

	return &OP{
		op:       op,
		priority: pri,
	}, nil
}

/*GT

o 是否比 another 的优先级高
*/
func (o *OP) GT(another *OP) bool {
	if another == nil {
		return true
	}

	return o.priority > another.priority
}

func (o *OP) Cal(num1, num2 int64) (int64, error) {
	if o.op == "+" {
		return num1 + num2, nil
	} else if o.op == "-" {
		return num1 - num2, nil
	} else if o.op == "*" {
		return num1 * num2, nil
	} else if o.op == "/" {
		if num2 == 0 {
			return 0, fmt.Errorf("divide 0")
		}

		return num1 / num2, nil
	} else {
		// ^ 运算
		return int64(math.Pow(float64(num1), float64(num2))), nil
	}
}

/*Cal
计算四则运算表达式

## 思路

创建两个栈，一个保存操作数，另一个保存运算符

我们从左向右遍历表达式，当遇到数字，我们就直接压入操作数栈；

当遇到运算符，就与运算符栈的栈顶元素进行比较。
	如果比运算符栈顶元素的优先级高，就将当前运算符压入栈；
	如果比运算符栈顶元素的优先级低或者相同，从运算符栈中取栈顶运算符，从操作数栈的栈顶取 2 个操作数，然后进行计算，再把计算完的结果压入操作数栈，继续比较。

## 复杂度分析

CalFullStack 的复杂度是 O(M)，M 是 opStack 的长度
Cal 的复杂度 O(N * M)， N 是表达式中字符的个数

所以 Cal 的复杂度就是 O(N^2)
*/
func Cal(expression string) (int64, error) {
	digitStack := NewStack[int64]()
	opStack := NewStack[OP]()

	for i := 0; i < len(expression); {
		c := expression[i]

		if unicode.IsDigit(rune(c)) {
			digitStart := i
			v, digitEnd, err := GetDigitFromExpression(expression, digitStart)
			if err != nil {
				return 0, err
			}
			fmt.Printf("Push %v on %d:%d\n", v, digitStart, digitEnd)
			digitStack.Push(v)

			i = digitEnd
			continue
		}

		op, err := NewOP(string(c))
		if err != nil {
			return 0, err
		}
		fmt.Printf("Parse op %v\n", op)

		top := opStack.Top()
		if op.GT(top) {
			opStack.Push(*op)
			i++
			continue
		} else {
			res, err := CalFullStack(&opStack, &digitStack)
			if err != nil {
				return 0, err
			}

			digitStack.Push(res)
			opStack.Push(*op)
			i++
		}
	}

	res, err := CalFullStack(&opStack, &digitStack)
	if err != nil {
		return 0, err
	}

	if opStack.Len() != 0 {
		return 0, fmt.Errorf("opStack %v is not empty", opStack)
	}

	if digitStack.Len() != 0 {
		return 0, fmt.Errorf("digitStack %v is not empty", digitStack)
	}

	return res, nil
}

/*CalFullStack

计算满状态的栈的表达式结果

满状态栈：opStack 有N个操作符，digitStack 有2N个数字
*/
func CalFullStack(opStack *Stack[OP], digitStack *Stack[int64]) (int64, error) {
	for {
		top := opStack.Pop()
		if top == nil {
			break
		}

		num2 := digitStack.Pop()
		num1 := digitStack.Pop()
		res, err := top.Cal(*num1, *num2)
		fmt.Printf("cal %v %v %v = %v\n", *num1, top, *num2, res)
		if err != nil {
			return 0, errors.WithMessagef(err, "cal %v %v %v error", num1, top, num2)
		}

		digitStack.Push(res)
	}

	res := digitStack.Pop()

	return *res, nil
}

/*GetDigitFromExpression

@description: 从表达式中解析数字
@param: expression 表达式
@param: digitStart 从这一位开始是数字
*/
func GetDigitFromExpression(expression string, digitStart int) (int64, int, error) {
	// 外部已经判断过，digitStart 指向的字符一定是数字
	// 所以从 digitStart+1 开始判断字符是否是数字
	idx := digitStart + 1

	for idx < len(expression) {
		c := rune(expression[idx])
		if !unicode.IsDigit(c) {
			break
		}
		idx++
	}

	digitEnd := idx
	digit := expression[digitStart:digitEnd]
	v, err := strconv.ParseInt(digit, 10, 0)
	if err != nil {
		return 0, 0, errors.WithMessagef(err, "covert %v failed", digit)
	}

	return v, digitEnd, nil
}
