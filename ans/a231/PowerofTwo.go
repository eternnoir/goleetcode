package a231

func IsPowerOfTwo(n int) bool {
	target := 1
	for {
		if target > n {
			return false
		}

		if target == n {
			return true
		}
		target = target * 2
	}
}
