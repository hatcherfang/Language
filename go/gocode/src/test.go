package main

import (
	"crypto/rand"
	"fmt"
	"log"
	"os"
	"path/filepath"
	"strings"
	"time"
)

func test() {
	var rand_data []byte = make([]byte, 10)
	rand.Read(rand_data)
	for i := 0; i < len(rand_data); i++ {
		println(rand_data[i])
	}
}

func getCurrentDirectory() string {
	dir, err := filepath.Abs(filepath.Dir(os.Args[0]))
	if err != nil {
		log.Fatal(err)
	}
	return strings.Replace(dir, "\\", "/", -1)
}

func saveAsFile(filePath string, array []uint) {
	fp, _ := os.Create(filePath)
	defer fp.Close()
	for i := 0; i < len(array); i++ {
		fmt.Fprintf(fp, "%d\n", array[i])
	}
	fp.Close()
}

// func readFile(filePath string) {
// 	fp, _ := os.Open(filePath)
// 	fp.ReadString()
//
// }

func fmtTest() {
	fmt.Println("f", "a", "n", "g", "haiqun")
}
func main() {
	//test()
	filePath := filepath.Join(getCurrentDirectory(), "filename.txt")
	array := []uint{11, 22, 33}
	println(array[0])
	saveAsFile(filePath, array)
	println(time.Now().UnixNano())
	fmtTest()
}
