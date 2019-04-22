package l151

func reverseWord(word string) string {
	runeString := []rune(word)
	var reverseString string
	for i := len(runeString) - 1; i >= 0; i-- {
		reverseString += string(runeString[i])
	}
	return reverseString
}

func reverseWords(s string) string {
	result := ""
	word := ""
	length := len(s)

	if s == "" {
		return ""
	}

	for i := length - 1; i >= 0; i-- {
		if s[i] == ' ' {
			if len(word) != 0 {
				result += reverseWord(word) + " "
				word = ""
			}
			continue
		}

		word += string(s[i])
	}

	if len(word) != 0 {
		result += reverseWord(word) + " "
	}

	if len(result) != 0 && result[len(result)-1] == ' ' {
		return result[:len(result)-1]
	}

	return result

}
