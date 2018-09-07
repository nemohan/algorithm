package search

//import "fmt"

// binary search
func BinarySearch(src []int, key int) bool {
	start := 0
	end := len(src) - 1
	for start <= end {
		mid := (start + end) / 2
		tmp := src[mid]
		if tmp == key {
			return true
		} else if key < tmp {
			end = mid - 1
		} else {
			start = mid + 1
		}
	}
	return false
}

func lumotoPartition(src []int) int {
	p := src[0]
	s := 0
	num := len(src)
	for i := s + 1; i < num; i++ {
		if src[i] < p {
			s++
			src[s], src[i] = src[i], src[s]
		}
	}
	src[s], src[0] = src[0], src[s]
	return s
}

// search the i'th max element
func QuickSelect(src []int, k int) int {
	start := 0
	end := len(src)
	for {
		s := lumotoPartition(src[start:end])
		if s == start+k-1 {
			return src[k]
		} else if s < start+k-1 {
			start = s + 1
		} else {
			end = start - 1
		}
	}
}

//interpolation search
func InterpolationSearch(src []int, key int) {

}
