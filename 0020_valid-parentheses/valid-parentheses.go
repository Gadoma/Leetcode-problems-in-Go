package goleetcode0020

func validParentheses(parentheses string) bool {
	//special cases
	switch {
	case len(parentheses)%2 == 1:
		return false
	case len(parentheses) == 0:
		return false
	}

	pairs := map[string]string{
		"(": ")",
		"{": "}",
		"[": "]",
	}
	stack := []string{}

	for _, r := range parentheses {
		item := string(r)
		if item == "(" || item == "{" || item == "[" {
			// push
			stack = append(stack, item)
		} else {
			if len(stack) == 0 {
				return false
			}

			n := len(stack) - 1
			top := stack[n]

			if item == pairs[top] {
				// pop
				stack = stack[:n]
			} else {
				return false
			}
		}
	}

	return len(stack) == 0
}
