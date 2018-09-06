package sort

func SelectSort(src []int) []int {
	num := len(src)
	for i := 0; i < num-1; i++ {
		min := i
		for j := i + 1; j < num; j++ {
			if src[j] < src[min] {
				min = j
			}
		}
		src[i], src[min] = src[min], src[i]
	}
	return src
}

func BubbleSort(src []int) []int {
	num := len(src)
	for i := 0; i < num-1; i++ {
		for j := 1; j < num-i; j++ {
			if src[j-1] > src[j] {
				src[j-1], src[j] = src[j], src[j-1]
			}
		}
	}
	return src
}

func InsertSort(src []int) []int {
	num := len(src)
	for i := 1; i <= num-1; i++ {
		tmp := src[i]
		j := i - 1
		for ; j >= 0 && tmp < src[j]; j-- {
			src[j+1] = src[j]
		}
		src[j+1] = tmp
	}
	return src
}
