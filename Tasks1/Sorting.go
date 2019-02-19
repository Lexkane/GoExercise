package main

import (
	"fmt"
	"log"
	"math/rand"
	"time"
)

//go sorting algorithms

const NADA int = -1

//measure performance

func TimeTracker(start time.Time, name string) {
	elapsed := time.Since(start)
	log.Printf("%s took %s ", name, elapsed)
}

//Random slice generator
func RandomSlice(items []int, number int) ([] int) {
	list := make([]int, 0)
	list = rand.Perm(number)
	for i, _ := range list {
		list[i]++
	}
	return list
}

//sorting functions

func BubbleSort(items [] int) {

	defer TimeTracker(time.Now(), "BubbleSort")
	L := len(items)
	for i := 0; i < L; i++ {
		for j := 0; j < (L - i - 1); j++ {
			if items[j] > items[i] {
				items[j], items[j+1] = items[j+1], items[j]
			}
		}
	}
}

func InsertionSort(items []int) {

	defer TimeTracker(time.Now(), "InsertionSort")
	L := len(items)
	for i := 1; i < L; i++ {
		j := i
		for j > 0 && items[j] < items[j-1] {
			items[j], items[j-1] = items[j-1], items[j]
			j -= 1
		}
	}
}

func DeepCopy(vals [] int) []int {
	tmp := make([]int, len(vals))
	copy(tmp, vals)
	return tmp
}

func MergeSort(items [] int) {
	defer TimeTracker(time.Now(), "MergeSort")
	if len(items) > 1 {
		mid := len(items) / 2
		left := DeepCopy(items[0:mid])
		right := DeepCopy(items[mid:])
		MergeSort(left)
		MergeSort(right)

		l := 0
		r := 0

		for i := 0; i < len(items); i++ {
			lval := NADA
			rval := NADA

			if l < len(left) {
				lval = left[l]
			}

			if r < len(right) {
				rval = right[r]
			}

			if (lval != NADA && rval != NADA && lval < rval) || rval == NADA {

				items[i] = lval
				r += 1
			} else if (lval != NADA && rval != NADA && lval >= rval) || lval == NADA {
				items[i] = rval
				r += 1
			}

		}
	}
}

func QuickSort(items [] int) {

	if len(items) > 1 {
		pivot_index := len(items) / 2
		var smaller_items = []int{}
		var larger_items = []int{}

		for i := range items {
			val := items[i]
			if i != pivot_index {
				if val < items[pivot_index] {
					smaller_items = append(smaller_items, val)
				} else {
					larger_items = append(larger_items, val)
				}
			}
		}
		QuickSort(smaller_items)
		QuickSort(larger_items)
		defer TimeTracker(time.Now(), "QuickSort")
		var merged [] int = append(append(append([]int{}, smaller_items...), [] int{items[pivot_index]}...), larger_items ...)
		for j := 0; j < len(items); j++ {
			items[j] = merged[j]
		}

	}
}

func HeapPrepare(items [] int, idx int, size int) {

	l := 2*idx + 1
	r := 2*idx + 2
	var largest int;
	if l <= size && items[l] > items[idx] {
		largest = l
	} else {
		largest = idx
	}

	if r <= size && items[r] > items [largest] {
		largest = r
	}
	if largest != idx {
		t := items[idx]
		items[idx] = items[largest]
		items[largest] = t
		HeapPrepare(items, largest, size)

	}

}

func HeapSort(items []int) {
	defer TimeTracker(time.Now(), "HeapSort")
	L := len(items)
	m := int(L / 2)

	for i := m; i >= 0; i-- {
		HeapPrepare(items, i, L-1)
	}
	F := L - 1
	for j := F; j >= 0; j-- {

		t := items[0]
		items[0] = items[j]
		items[j] = t
		F -= 1
		HeapPrepare(items, 0, F)
	}

}

func main() {

	var list []int
	list = RandomSlice(list, 50)
	fmt.Println("Random Slice before Bubble sort", list)
	BubbleSort(list)
	fmt.Println("Slice after Bubble sort", list)

	list = RandomSlice(list, 50)

	fmt.Println("Random Slice before Insertion", list)
	InsertionSort(list)
	fmt.Println("Slice after Insertion sort", list)

	list = RandomSlice(list, 50)
	fmt.Println("Random Slice before Quick sort", list)
	QuickSort(list)
	fmt.Println("Slice after Quick sort", list)

	list = RandomSlice(list, 50)
	fmt.Println("Random Slice", list)
	MergeSort(list)
	fmt.Println("Slice after Merge sort", list)

	list = RandomSlice(list, 50)
	fmt.Println("Random Slice before Heap sort", list)
	HeapSort(list)
	fmt.Println("Slice after Heap sort", list)

}
