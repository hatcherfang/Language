package main

import (
	"fmt"
	"math/rand"
	"time"
	"bufio"
	"os"
	"strconv"
)

func main() {
	var list []int
	for i := 0; i < 1000000; i++ {
		list = append(list, rand.Intn(1000000))
	}
	
	save_data("D:\\sort_dat", list)
	
    t1 := time.Now()
	quicksort(list)
	t2 := time.Now()
	
	fmt.Println(t2.Sub(t1))
	
	save_data("D:\\sort_result", list)
}

func quicksort(list []int) {
	if len(list) <= 1 {
		return
	}

	i, j := 0, len(list)-1
	index := 1 
	key := list[0]

	for i < j {
		if list[index] > key {
			list[index], list[j] = list[j], list[index]
			j--
		} else {
			list[index], list[i] = list[i], list[index]
		i++
		index++
		}
	}
	
	list[i] = key
	quicksort(list[:i])
	quicksort(list[i+1:])
}


func save_data(file_name string, vals []int) error{
	file, err := os.Create(file_name)
	if err != nil {
		 fmt.Println("writer",err)
		return err
	}
	defer file.Close()
 
	writer := bufio.NewWriter(file)
	for _, v := range vals {
		writer.WriteString(strconv.Itoa(v))
		writer.WriteString("\n")
		writer.Flush()
	}
	return err
}