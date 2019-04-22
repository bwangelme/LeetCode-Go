package l151

func reverse2(word string) string {
	runeString := []rune(word)
	var reverseString string
	for i := len(runeString) - 1; i >= 0; i-- {
		reverseString += string(runeString[i])
	}
	return reverseString
}

func reverseWords2(s string) string {
	return ""
}
