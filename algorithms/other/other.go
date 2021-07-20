package other

func ContainString(s []string, target string) bool {
	for _, v := range s {
		if v == target {
			return true
		}
	}

	return false
}

func CheckBrackets(str string) bool {
	// mapping of opening and closing brackets.
	m := map[string]string{
		"(": ")",
		"{": "}",
		"[": "]",
	}

	stack := []string{}

	for _, char := range str {
		// if it is an opener, push to the stack.
		if ContainString([]string{"(", "{", "["}, string(char)) {
			stack = append(stack, string(char))
		} else if ContainString([]string{")", "}", "]"}, string(char)) {
			// return false if the stack is empty.
			if len(stack) == 0 {
				return false
			}

			// pop to get the last item from the stack.
			lastUnclosed := stack[len(stack)-1]
			stack = stack[:len(stack)-1]

			// return false if the last opening bracket doesn't match the
			// current closing one.
			if m[lastUnclosed] != string(char) {
				return false
			}
		}
	}

	return len(stack) == 0
}
