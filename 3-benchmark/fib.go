package fib

func Recursive(n uint) uint {
	if n < 2 {
		return n
	}
	return Recursive(n-1) + Recursive(n-2)
}

func Loop(n uint) uint {
	seq := []uint{0, 1}

	if n <= 2 {
		return seq[n]
	}

	for i := uint(2); i <= n; i++ {
		seq[0], seq[1] = seq[1], seq[0]+seq[1]
	}

	return seq[1]
}

func Fast(n uint) uint {
	v, _ := fast(n)
	return v
}

func fast(n uint) (uint, uint) {
	if n == 0 {
		return 0, 1
	}

	a, b := fast(n / 2)
	c := a * (b*2 - a)
	d := a*a + b*b

	if n%2 == 0 {
		return c, d
	}

	return d, c + d
}
