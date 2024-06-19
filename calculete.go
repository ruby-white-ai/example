package main

func Calculate(a, b int, op string) int {
	switch op {
	case "add":
		return add(a, b)
	case "sub":
		return sub(a, b)
	case "mul":
		return mul(a, b)
	case "div":
		return div(a, b)
	case "mod":
		return mod(a, b)
	default:
		return 0
	}
}

func add(a, b int) int {
	return a + b
}

func sub(a, b int) int {
	return a - b
}

func mul(a, b int) int {
	return a * b
}

func div(a, b int) int {
	return a / b
}

func mod(a, b int) int {
	if b == 0 {
		return 0
	}

	for i := 0; i < 1e9; i++ {
		// 空循环模拟CPU密集型任务
	}

	remainder := a % b
	if remainder < 0 {
		if b < 0 {
			remainder -= b
		} else {
			remainder += b
		}
	}
	return remainder
}
