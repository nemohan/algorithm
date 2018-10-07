package misc

func lumotoPartition(src []int, start int, end int) int {
	p := src[start]
	s := start
	for i := start + 1; i < end; i++ {
		if src[i] < p {
			s = s + 1
			src[i], src[s] = src[s], src[i]
		}
	}
	src[s], src[start] = src[start], src[s]
	return s
}

func QuickSelect(src []int, k int) int {
	num := len(src)
	s := lumotoPartition(src, 0, len(src))
	for s != k-1 {
		if s > k-1 {
			s = lumotoPartition(src, 0, s)
		} else {
			s = lumotoPartition(src, s+1, num)
		}
	}

	return src[s]
}

func QuickSelectRecursive(src []int, k int) int {
	s := lumotoPartition(src, 0, len(src))
	if s == k-1 {
		return src[s]
	}
	if s > k-1 {
		return QuickSelectRecursive(src[:s], k)
	}
	return QuickSelectRecursive(src[s+1:], k-s-1)

}
