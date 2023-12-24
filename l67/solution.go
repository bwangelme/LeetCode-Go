package l67

func addBinary(a string, b string) string {
	lenA := len(a) - 1
	lenB := len(b) - 1
	carry := 0
	result := make([]byte, 0)

	for lenA >= 0 || lenB >= 0 {
		var digitA = 0
		if lenA >= 0 {
			digitA = int(a[lenA] - '0')
		}
		var digitB = 0
		if lenB >= 0 {
			digitB = int(b[lenB] - '0')
		}

		res := digitA + digitB + carry
		if res >= 2 {
			carry = 1
			res = res - 2
		} else {
			carry = 0
		}

		result = append([]byte{uint8(res) + '0'}, result...)
		lenA--
		lenB--
	}

	if carry == 1 {
		result = append([]byte{'1'}, result...)
	}

	return string(result)
}
