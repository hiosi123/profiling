package fibonacci

func FibonacciOriginal(n int) int {
	if n <= 0 {
		return 0
	} else if n == 1 {
		return 1
	}
	return FibonacciOriginal(n-1) + FibonacciOriginal(n-2)
}

func FibonacciOptimized(n int) int {
	if n <= 0 {
		return 0
	} else if n == 1 {
		return 1
	}

	fib := make([]int, n+1)
	fib[0], fib[1] = 0, 1

	for i := 2; i <= n; i++ {
		fib[i] = fib[i-1] + fib[i-2]
	}

	return fib[n]
}
