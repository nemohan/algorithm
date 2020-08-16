package sort

import (
	"math/rand"
	"sort"
	"testing"
)

func genTestInput(size int) []int {
	keySet := make(map[int]bool, 0)
	keys := make([]int, 0)
	limit := size * 2
	for i := 0; i < size; i++ {
		//j := rand.Int31n(500)
		j := rand.Intn(limit)
		if _, ok := keySet[j]; ok {
			continue
		}
		keySet[j] = false
		keys = append(keys, j)
	}
	return keys
}
func checkResult(want []int, got []int) bool {
	for i, v := range want {
		if v != got[i] {
			return false
		}
	}
	return true
}

func TestQuickSortAsendingOrder(t *testing.T) {
	old := []int{1, 2, 3, 4, 5}
	input := []int{1, 2, 3, 4, 5}
	quickSort(input)
	if !checkResult(old, input) {
		t.Errorf("TestQuickSort want:%v got:%v\n", old, input)
	}

	old = []int{1, 2, 3, 4}
	input = []int{1, 2, 3, 4}
	quickSort(input)
	if !checkResult(old, input) {
		t.Errorf("TestQuickSort want:%v got:%v\n", old, input)
	}
}

func TestQuickSortDecentOrder(t *testing.T) {
	old := []int{1, 2, 3, 4, 5}
	input := []int{5, 4, 3, 2, 1}
	quickSort(input)
	if !checkResult(old, input) {
		t.Errorf("TestQuickSort want:%v got:%v\n", old, input)
	}

	old = []int{1, 2, 3, 4}
	input = []int{4, 3, 2, 1}
	quickSort(input)
	if !checkResult(old, input) {
		t.Errorf("TestQuickSort want:%v got:%v\n", old, input)
	}
}

func TestQuickSort(t *testing.T) {
	input := genTestInput(1000000)
	want := make([]int, len(input))
	copy(want, input)
	sort.Ints(want)

	quickSort(want)
	if !checkResult(want, input) {
		t.Errorf("TestQuickSort want:%v got:%v\n", want, input)
	}
}

func TestMergeSort(t *testing.T) {
	input := genTestInput(1000000)
	data := make([]int, len(input))
	copy(data, input)
	sort.Ints(data)

	got := mergeSort(input)
	if !checkResult(data, got) {
		t.Errorf("TestMergeSort want:%v got:%v\n", data, got)
	}
}

func TestShellSort(t *testing.T) {
	input := []int{1, 2, 3, 4, 5}
	data := make([]int, len(input))
	copy(data, input)
	got := shellSort(input)
	sort.Ints(data)
	if !checkResult(data, got) {
		t.Errorf("TestShellSort failed.")
	}

	input = []int{5, 4, 3, 2, 1}
	copy(data, input)
	got = shellSort(input)
	sort.Ints(data)
	if !checkResult(data, got) {
		t.Errorf("TestShellSort failed. want:%v got:%v", data, got)
	}
	input = genTestInput(1000000)
	data = make([]int, len(input))
	copy(data, input)
	sort.Ints(data)

	got = shellSort(input)
	if !checkResult(data, got) {
		t.Errorf("TestShellSort failed. want:%v got:%v", data, got)
	}
}

func TestInsertionSort(t *testing.T) {
	input := []int{1, 2, 3, 4, 5}
	data := make([]int, len(input))
	copy(data, input)
	sort.Ints(data)

	got := insertionSort(input)
	if !checkResult(data, got) {
		t.Errorf("TestInsertionSort failed.")
	}

	input = []int{5, 4, 3, 2, 1}
	copy(data, input)
	sort.Ints(data)

	got = insertionSort(input)
	if !checkResult(data, got) {
		t.Errorf("TestInsertionSort failed. want:%v got:%v", data, got)
	}
	input = genTestInput(1000000)
	data = make([]int, len(input))
	copy(data, input)
	sort.Ints(data)

	got = insertionSort(input)
	if !checkResult(data, got) {
		t.Errorf("TestInsertionSort failed. want:%v got:%v", data, got)
	}
}

func BenchmarkMergeSort(b *testing.B) {
	for i := 0; i < b.N; i++ {
		input := genTestInput(200)
		got := mergeSort(input)
		sort.Ints(input)
		if !checkResult(input, got) {
			b.Errorf("TestMergeSort want:%v got:%v\n", input, got)
		}
	}
}

func BenchmarkShellSort(b *testing.B) {
	for i := 0; i < b.N; i++ {
		input := genTestInput(2000000)
		data := make([]int, len(input))
		copy(data, input)

		got := shellSort(input)
		sort.Ints(data)
		if !checkResult(data, got) {
			b.Errorf("TestMergeSort want:%v got:%v\n", data, got)
		}
	}
}
