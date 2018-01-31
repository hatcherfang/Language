package main

import (
	"bufio"
	"crypto/rand"
	"encoding/binary"
	"fmt"
	"os"
	"time"
)

func print_ary(ary []uint) {
	for i := 0; i < len(ary); i++ {
		fmt.Printf("%d,", ary[i])
	}
	fmt.Printf("\n")
}

func save_to_file(file_path string, ary []uint, cost int64) {

	var file, _ = os.Create(file_path)
	var w = bufio.NewWriter(file)

	for i := 0; i < len(ary); i++ {
		fmt.Fprintf(w, "%d\n", ary[i])
	}

	fmt.Fprintf(w, "\ncost: %d ms\n", cost)

	w.Flush()
	file.Close()
}

func sort_1(ary []uint) int64 {

	var now = time.Now()
	var start = now.UnixNano()

	for i := 0; i < len(ary)-1; i++ {
		for j := i + 1; j < len(ary); j++ {
			if ary[i] > ary[j] {
				var val = ary[i]
				ary[i] = ary[j]
				ary[j] = val
			}
		}
	}

	now = time.Now()
	var end = now.UnixNano()

	fmt.Printf("%d, cost: %d ms", len(ary), (end-start)/1000/1000)

	return (end - start) / 1000 / 1000
}

func sort_quick(ary []uint) {

	var ary_len = len(ary)
	if ary_len <= 1 {
		return
	}

	var mid int = -1
	var mid_val = ary[0]

	var j = ary_len - 1

	for j > mid {

		if ary[j] < mid_val {
			mid++

			var val = ary[j]
			ary[j] = ary[mid]
			ary[mid] = val
		} else {
			j--
		}
	}

	if mid < 0 {
		sort_quick(ary[1:])
	} else {
		sort_quick(ary[0 : mid+1])
		sort_quick(ary[mid+1:])
	}

	//print_ary(ary)
}

func sort_2(ary []uint) int64 {

	var now = time.Now()
	var start = now.UnixNano()

	sort_quick(ary)

	now = time.Now()
	var end = now.UnixNano()

	fmt.Printf("%d, cost: %d ms", len(ary), (end-start)/1000/1000)

	return (end - start) / 1000 / 1000
}

func do_sort(size int) {
	var ary []uint = make([]uint, size)

	for i := 0; i < len(ary); i++ {
		var rand_data []byte = make([]byte, 4)
		rand.Read(rand_data)

		ary[i] = uint(binary.LittleEndian.Uint32(rand_data))
	}

	save_to_file(fmt.Sprintf("d:\\temp\\sort_data_%d.txt", size), ary, 0)

	var cost = sort_2(ary)

	save_to_file(fmt.Sprintf("d:\\temp\\result_%d.txt", size), ary, cost)

	for i := 0; i < len(ary)-1; i++ {
		if ary[i] > ary[i+1] {
			panic("sort failed")
		}
	}
}

func main() {
	do_sort(1000000)
}
