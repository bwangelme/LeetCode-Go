package l373

import (
	"fmt"
	"github.com/stretchr/testify/assert"
	"os"
	"sort"
	"strconv"
	"strings"
	"testing"
)

func readArrListFromFile(filename string) (res [][]int, err error) {
	// 读取文件内容
	data, err := os.ReadFile(filename)
	if err != nil {
		err = fmt.Errorf("error reading file: %v", err)
		return res, err
	}

	// 将文件内容转换为字符串
	content := string(data)

	// 去掉外层的方括号
	content = strings.Trim(content, "[]")

	// 按 "],[" 分割字符串，得到二维数组的每一部分
	pairs := strings.Split(content, "],[")

	// 遍历每一对数据
	for _, pair := range pairs {
		// 将每一对数字按逗号分割
		nums := strings.Split(pair, ",")

		// 创建一个临时切片来存储当前的数字对
		var pairArray []int

		// 将字符串转换为整数
		for _, num := range nums {
			n, err := strconv.Atoi(num)
			if err != nil {
				err = fmt.Errorf("error converting string to int: %v", err)
				return res, err
			}
			pairArray = append(pairArray, n)
		}

		// 将转换后的数字对添加到结果二维数组中
		res = append(res, pairArray)
	}

	// 输出结果
	return res, nil
}

func readArrFromFile(filename string) (res []int, err error) {
	// 读取文件内容
	data, err := os.ReadFile(filename)
	if err != nil {
		err = fmt.Errorf("error reading file: %v", err)
		return res, err
	}

	// 将文件内容转换为字符串
	content := string(data)
	// 去掉开头和结尾的方括号
	content = strings.Trim(content, "[]")
	// 使用逗号分隔每个数字
	strNumbers := strings.Split(content, ",")

	// 创建一个整型数组
	var numbers []int
	// 将每个数字字符串转换为整数并加入数组
	for _, str := range strNumbers {
		num, err := strconv.Atoi(str)
		if err != nil {
			err = fmt.Errorf("error converting string to int: %v", err)
			return res, err
		}
		numbers = append(numbers, num)
	}

	// 输出结果
	return numbers, nil
}

func Test_kSmallestPairsV3(t *testing.T) {
	var res [][]int

	//res = kSmallestPairsV3([]int{1, 7, 11}, []int{2, 4, 6}, 3)
	//assert.Equal(t, res, [][]int{{1, 2}, {1, 4}, {1, 6}})
	//
	//res = kSmallestPairsV3([]int{1, 1, 2}, []int{1, 2, 3}, 2)
	//assert.Equal(t, res, [][]int{{1, 1}, {1, 1}})

	num1, err := readArrFromFile("num1.txt")
	assert.Nil(t, err)
	num2, err := readArrFromFile("num2.txt")
	assert.Nil(t, err)
	wanted, err := readArrListFromFile("answer.txt")
	assert.Nil(t, err)

	res = kSmallestPairsV3(num1, num2, 9484)
	assert.Equal(t, len(res), len(wanted))

	// 对结果进行排序后，再比较它们是否相等
	sort.Slice(wanted, func(i, j int) bool {
		// 比较第一位（索引 0）
		if wanted[i][0] != wanted[j][0] {
			return wanted[i][0] < wanted[j][0]
		}
		// 如果第一位相等，比较第二位（索引 1）
		return wanted[i][1] < wanted[j][1]
	})
	sort.Slice(res, func(i, j int) bool {
		// 比较第一位（索引 0）
		if res[i][0] != res[j][0] {
			return res[i][0] < res[j][0]
		}
		// 如果第一位相等，比较第二位（索引 1）
		return res[i][1] < res[j][1]
	})
	n := len(res)
	for i := 0; i < n; i++ {
		if res[i][0] != wanted[i][0] || res[i][1] != wanted[i][1] {
			t.Fatalf("%v: %v != %v", i, res[i], wanted[i])
		}
	}
}
