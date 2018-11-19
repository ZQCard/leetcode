package climbStairs

func climbStairs(n int) int {
	if n < 3 {
		return n
	}
	a := 1
	b := 1
	var temp int
	for i := 3; i <= n + 1; i++{
		temp = a + b
		a = b
		b = temp
	}
	return temp
}



func FibonacciSum(n int) int {
	var total int
	for i := n; i > 0; i-- {
		total += Fibonacci(n)
	}
	return total
}

func Fibonacci(n int) int {
	if (n <= 0) {
		return 0
	}
	if (n == 1) {
		return 1
	}

	return Fibonacci(n - 1) + Fibonacci(n - 2)
}

func Fibonacci2(n int) int {
	if n < 3 {
		return n
	}
	var temp int
	a := 1
	b := 1
	for i := 3; i <= n; i++ {
		temp = a + b
		a = b
		b = temp
}
	return temp
}
