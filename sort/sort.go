package sort

import ()

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

func partition(src []int) int {
	i := 0
	num := len(src)
	j := num - 1
	p := src[0]
	for i < j {
		for k := i + 1; k < num && src[k] < p; k++ {
			i++
		}
		for src[j] > p {
			j--
		}
		if i == num {
			i--
		}
		src[j], src[i] = src[i], src[j]
	}
	src[i], src[j] = src[j], src[i]
	src[0], src[j] = src[j], src[0]
	return j
}

//unstable
func QuickSort(src []int) []int {
	if len(src) > 1 {
		p := partition(src)
		QuickSort(src[:p])
		QuickSort(src[p+1:])
	}
	return src
}

func merge(src1 []int, src2 []int, dst []int) {
	num := len(src1)
	num2 := len(src2)
	k := 0
	i := 0
	j := 0
	for i < num && j < num2 {
		if src1[i] < src2[j] {
			dst[k] = src1[i]
			i++
		} else {
			dst[k] = src2[j]
			j++
		}
		k++
	}

	if i < num {
		copy(dst[k:], src1[i:])
	} else {
		copy(dst[k:], src2[j:])
	}
}

func MergeSort(src []int) []int {
	if len(src) > 1 {
		m := len(src) / 2
		src1 := MergeSort(src[:m])
		src2 := MergeSort(src[m:])
		merge(src1, src2, src)
	}
	tmp := make([]int, 0, len(src))
	tmp = append(tmp, src[:]...)
	return tmp
}

func HeapSort(src []int) []int {
	return nil
}
