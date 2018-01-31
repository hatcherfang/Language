package main
import ("fmt"     
    "math/rand"
    "sort"
    "time")
    
type IntSlice []int

func (c IntSlice) Len() int {
	return len(c)
}
func (c IntSlice) Swap(i, j int) {
	c[i], c[j] = c[j], c[i]
}
func (c IntSlice) Less(i, j int) bool {
	return c[i] < c[j]
}

func main() {
    const size int = 1000000
    var c IntSlice
    for i:=0; i<size; i++ {
        c = append(c,rand.Intn(size))
    }
    //fmt.Println(c)
    start := time.Now()
    sort.Sort(c)
    elapsed := time.Since(start)
    fmt.Println("App elapsed: ", elapsed)
    //fmt.Println(c)
} 
