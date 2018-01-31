package main

import "fmt"
import "math/rand"

//import "io"
import "os"

//import "strconv"
import "bufio"
import "time"

const count = 1000000

func quick_sort(para *[count]int, left int, right int) {
	temp := (*para)[left]
	k := left
	i, j := left, right
	for i <= j {
		for j >= k && (*para)[j] >= temp {
			j--
		}
		if j >= k {
			(*para)[k] = (*para)[j]
			k = j
		}
		for i <= k && (*para)[i] <= temp {
			i++
		}
		if i <= k {
			(*para)[k] = (*para)[i]
			k = i
		}
	}
	(*para)[k] = temp
	if k-left > 1 {
		quick_sort(para, left, k-1)
	}
	if right-k > 1 {
		quick_sort(para, k+1, right)
	}
}

func check(e error) {
	if nil != e {
		fmt.Print("error!\n")
	}
}

func main() {
	var arr [count]int
	var f *os.File
	//var err error
	var filename = "/root/sort_value.txt"
	start_time := time.Now()
	//fmt.Print("original data:\n")
	for i := 0; i < count; i++ {
		arr[i] = rand.Intn(count)
		//fmt.Print(arr[i],"\n")
	}
	//fmt.Print("quick sort data:\n")
	quick_sort(&arr, 0, count-1)
	f, _ = os.Create(filename)
	//check(err)
	fb := bufio.NewWriter(f)
	for i := 0; i < count; i++ {
		//str:=strconv.Itoa(arr[i])
		//num:=strconv.Itoa(i)
		//io.WriteString(f,num +" "+str+"\n")
		fmt.Fprintf(fb, "%d\n", arr[i])
		//fmt.Print(arr[i],"\n")
	}
	end_time := time.Now()
	//io.WriteString(f,end_time.Sub(start_time))
	fmt.Print("Generate ")
	fmt.Print(count)
	fmt.Print(" random numbers and sort.\n")
	fmt.Print("Cost time:")
	fmt.Print(end_time.Sub(start_time))
	fmt.Print("\n")
}
