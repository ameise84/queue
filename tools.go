package queue

func ceilToPowerOfTwo(n uint32) uint32 {
	if n <= 0 {
		return 0
	}
	if n <= 2 {
		return n
	}

	n--
	n |= n >> 1
	n |= n >> 2
	n |= n >> 4
	n |= n >> 8
	n |= n >> 16
	n++
	if n < 0 {
		n = 0
	}
	return n
}
