package main

import (
	"bufio"
	"crypto/rand"
	"encoding/binary"
	"fmt"
	"log"
	"os"
	"path/filepath"
	"strings"
	"time"
)

func sort(size int) {
	// create an array
	var array []uint = make([]uint, size)
	for i := 0; i < len(array); i++ {
		// byte <=> uint8
		var rand_data []byte = make([]byte, 4)
		rand.Read(rand_data)
		// 4 * uint8 <=> uint32
		array[i] = uint(binary.LittleEndian.Uint32(rand_data))
	}
	// writedown raw data
	curDir := getCurrentDirectory()
	rawDataFilePath := filepath.Join(curDir, "dataBeforeSort.txt")
	saveFile(rawDataFilePath, array)
	// sort array
	start := time.Now()
	sortData(array)
	println(time.Since(start)/1000/1000, "ms")
	// writedown sorted data
	sortedDataFilePath := filepath.Join(curDir, "dataAfterSort.txt")
	saveFile(sortedDataFilePath, array)
}

func getCurrentDirectory() string {
	dir, err := filepath.Abs(filepath.Dir(os.Args[0]))
	if err != nil {
		log.Fatal(err)
		return ""
	}
	return strings.Replace(dir, "\\", "/", -1)
}

func saveFile(filePath string, array []uint) {
	fp, err := os.Create(filePath)
	if err != nil {
		log.Fatal(err)
		return
	}
	defer fp.Close()
	fpb := bufio.NewWriter(fp)
	for i := 0; i < len(array); i++ {
		fmt.Fprintf(fpb, "%d\n", array[i])
	}
	fpb.Flush()
	fp.Close()
}

func adjustArray(low int, high int, array []uint) int {
	if low >= high {
		return -1
	}
	pivot := array[low]
	//println(pivot, low, high)
	for low < high {
		for low < high && array[high] > pivot {
			high--
		}
		//println("high", low, high, array[low], array[high])
		if low < high {
			array[low] = array[high]
			low++
		}

		for low < high && array[low] <= pivot {
			low++
		}
		//println("low", low, high, array[low], array[high])
		if low < high {
			array[high] = array[low]
			high--
		}
	}
	array[low] = pivot
	return low
}

func quickSort(low int, high int, array []uint) {
	if low < high {
		mid := adjustArray(low, high, array)
		quickSort(low, mid-1, array)
		quickSort(mid+1, high, array)

	}
}

func sortData(array []uint) {
	quickSort(0, len(array)-1, array)
}

func main() {
	// write 100w random data into file
	dataNum := 1000000
	println("nums:", dataNum)
	sort(dataNum)
}
