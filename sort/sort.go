package sort

func quickSort(a []int) {
	if len(a) <= 1 {
		return
	}
	p := partition(a)
	quickSort(a[:p])
	quickSort(a[p:])
}

func partition(a []int) int {
	p := 0
	l := 1
	r := len(a) - 1
	for l <= r {
		for l <= r && a[l] < a[p] {
			l++
		}
		for r >= l && a[r] > a[p] {
			r--
		}
		if r < l {
			break
		}
		a[l], a[r] = a[r], a[l]
	}
	if r <= p {
		p++
		return p
	}
	a[p], a[r] = a[r], a[p]
	return r
}

func mergeSort(a []int) []int {
	n := len(a)
	if n <= 1 {
		return a
	}
	l := mergeSort(a[:n/2])
	r := mergeSort(a[n/2:])
	l = merge(l, r)
	return l
}

func merge(l, r []int) []int {
	nl := len(l)
	nr := len(r)
	dst := make([]int, len(l)+len(r))
	j := 0
	i := 0
	k := 0
	for j < nl && k < nr {
		if l[j] < r[k] {
			dst[i] = l[j]
			i++
			j++
		} else {
			dst[i] = r[k]
			i++
			k++
		}
	}
	for j < nl {
		dst[i] = l[j]
		i++
		j++
	}
	for k < nr {
		dst[i] = r[k]
		i++
		k++
	}
	return dst
}

func shellSort(src []int) []int {
	n := len(src)
	for h := n / 2; h > 0; h = h / 2 {
		i := h
		for ; i < n; i++ {
			tmp := src[i]
			j := i
			for ; j >= h; j -= h {
				/*
					if src[j] < src[j-h] {
						src[j-h], src[j] = src[j], src[j-h]
					}
				*/
				if src[j-h] > tmp {
					src[j] = src[j-h]
				} else {
					break
				}
			}
			src[j] = tmp
		}
	}
	return src
}

/*
func shellSortWrong(src [])[]int{
	n := len(src)
	for h := n / 2; h > 0; h = h / 2 {
		i := h
		for ; i < n; i++ {
			if src[i] < src[i-h]{

			}
		}
	}
	return src
}
*/

func insertionSort(src []int) []int {
	n := len(src)
	for i := 1; i < n; i++ {
		j := i
		tmp := src[j]
		for ; j >= 1; j-- {
			if tmp >= src[j-1] {
				break
			}
			src[j] = src[j-1]
		}
		src[j] = tmp
	}
	return src
}
