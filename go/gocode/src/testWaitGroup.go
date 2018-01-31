/*
出现问题场景: 一个函数run()中包含多个goroutine函数并发, 这些goroutine函数会生成中间文件, 被run()函数运行结束后的check()函数检查. 当goroutine并发时, 并不会阻塞run()的上下文, 可能导致的情况为run()函数执行完毕(但其中的goroutine并发函数没有执行完毕), 导致check()函数执行失败.

所以我们需要一种操作, 直到当前所有goroutine没有执行完毕, 才进行下一步操作

所以需要goroutine同步, go提供了sync包和channel机制来解决goroutine之间的同步问题

sync.WaitGroup

A WaitGroup waits for a collection of goroutines to finish. The main goroutine calls Add to set the number of goroutines to wait for. Then each of the goroutines runs and calls Done when finished. At the same time, Wait can be used to block until all goroutines have finished.  -- 出自官方文档

大概意思是: WaitGroup等待一组goroutinue执行完毕. 主goroutinue调用Add设置等待的goroutinue数量. 每个goroutinue应该在执行结束时调用Done. Wait会阻塞知道所有goroutinue执行完毕.

WaitGroup的用于某个地方需要创建多个goroutine，并且一定要等它们都执行完毕后再继续执行接下来的操作.

可以把WaitGroup看作一个类似任务队列的结构. Add想队列增加任务, Done完成任务, Wait在队列不空的时候阻塞在哪里.
*/

// 官方文档中的example
package main

import (
	"fmt"
	"net/http"
	"sync"
)

func main() {
	var wg sync.WaitGroup // 声明一个WaitGroup变量
	var urls = []string{
		"http://www.baidu.org/",
		"http://www.alibaba.com/",
		"http://www.qq.com/",
	}
	for _, url := range urls {
		wg.Add(1) // WaitGroup的计数加1
		// Launch a goroutine to fetch the URL.
		go func(url string) {
			defer wg.Done() //  goroutinue完成后, WaitGroup的计数-1
			// Fetch the URL.
			http.Get(url)
			fmt.Println(url)
		}(url)
	}
	// Wait for all HTTP fetches to complete.
	wg.Wait() // 等待所有goroutinue完成
}
