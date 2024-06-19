package main

func Calculate(a, b int, op string) int {
	switch op {
	case "add":
		return Add(a, b)
	case "sub":
		return Sub(a, b)
	case "mul":
		return Mul(a, b)
	case "div":
		return Div(a, b)
	case "mod":
		return Mod(a, b)
	default:
		return 0
	}
}

func Add(a, b int) int {
	return a + b
}

func Sub(a, b int) int {
	return a - b
}

func Mul(a, b int) int {
	return a * b
}

func Div(a, b int) int {
	return a / b
}

func Mod(a, b int) int {
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
