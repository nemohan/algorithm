package sort

import ()

func SelectSort(src []int) []int {
	num := len(src)
	for i := 0; i < num-1; i++ {
		min := src[i]
		idx := 0
		for j := i + 1; j < num; j++ {
			if src[j] < min {
				min = src[j]
				idx = j
			}
		}
		src[i], src[idx] = src[idx], src[i]
	}
	return src
}

func BubbleSort(src []int) []int {
	num := len(src)
	for i := 0; i < num; i++ {
		for j := i + 1; j < num-i; j++ {
			if src[j] > src[i] {
				src[i], src[j] = src[j], src[i]
			}
		}
	}
	return src
}
