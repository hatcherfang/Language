package main

import (
"fmt"
"math/rand"
"time"
"sync/atomic"
"bufio"
"os"
)

const SORT_CNT = 1000000

func main() {
    list := make([]int32, SORT_CNT)
    r := rand.New(rand.NewSource(time.Now().UnixNano()))
    for i:=0; i<len(list); i++ {
        list[i] = int32(r.Intn(SORT_CNT))
    }
    org := make([]int32, SORT_CNT)
    copy(org, list)

    time1 := time.Now()
    routine_hash_sort(list)
    time2 := time.Now()
    fmt.Println("With routine time:", time2.Sub(time1))

    save_to_file("org.txt", org)
    save_to_file("sorted.txt", list)

    a := time.Now()
    hash_sort(org)
    b := time.Now()
    fmt.Println("No routine time:", b.Sub(a))
}

func hash_sort(src []int32) {
    var temp [2][SORT_CNT]int32
    for i:=0; i < len(src); i++ {
        tempPos := src[i]
        temp[0][tempPos] = src[i]
        temp[1][tempPos]++
    }
    src = src[:0]
    for i := int32(0); i < int32(len(temp[0])); i++ {
        for j:=int32(0); j< temp[1][i]; j++ {
            src = append(src, temp[0][i])
        }
    }
}

func hash_sort_m(dest [][]int32, src []int32, channle chan int32, routineCnt int32, i int32){
    singleLen := int32(len(src))/routineCnt
    temp := src[i*singleLen: i*singleLen + singleLen]

    for j := int32(0); j < int32(len(temp)); j++ {
        tempPos := temp[j]
        dest[0][tempPos] = temp[j]
        //dest[1][tempPos]++
        atomic.AddInt32(&dest[1][tempPos], 1)
    }
    channle <- i
}

func routine_hash_sort(src []int32) {
    routineCnt := int32(10)
    channle := make(chan int32, routineCnt)

    dest := make([][]int32, 2)
    for i := int32(0); i < int32(2); i++ {
        dest[i] = make([]int32, SORT_CNT)
    }

    for i := int32(0); i < routineCnt; i++ {
        go hash_sort_m(dest, src, channle, routineCnt, i)
    }

    for i := int32(0); i < routineCnt; i++ {
        <- channle
    }

    src = src[:0]
    for i := int32(0); i < int32(len(dest[0])); i++ {
        for j:=int32(0); j< int32(dest[1][i]); j++ {
            src = append(src, dest[0][i])
        }
    }
}

func save_to_file(fileName string, data []int32) {
    var file, _ = os.Create(fileName)
    defer file.Close()

    var write = bufio.NewWriter(file)
    for i := 0; i < len(data); i++ {
        fmt.Fprintf(write, "%d\n", data[i])
    }
    write.Flush()
}
