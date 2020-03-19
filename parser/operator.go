package parser

func precedence(op string) int {
	switch op {
	case "in", "not_in":
		return 7
	case "*", "/", "%":
		return 6
	case "+", "-":
		return 5
	case "<", "<=", ">", ">=":
		return 4
	case "==", "!=":
		return 3
	case "&&", "and":
		return 2
	case "||", "or":
		return 1
	}
	return 0
}
