package main

import (
	"context"
	"fmt"
	"io"
	"net/http"
	"os"
	"strings"
	"time"
)

func main() {
	start := time.Now()
	ch := make(chan string)
	urls := os.Args[1:]
	for _, url := range urls {
		go fetch(url, ch)
	}

	for range urls {
		fmt.Println(<-ch)
	}
	fmt.Printf("%.2fs elapsed\n", time.Since(start).Seconds())
}

func fetch(url string, ch chan<- string) {
	start := time.Now()
	if !(len(url) >= 7 && (strings.HasPrefix(url, "http://") || strings.HasPrefix(url, "https://"))) {
		url = "https://" + url
	}

	// 步骤1: 创建一个10秒超时的 context
	ctx, cancel := context.WithTimeout(context.Background(), 10*time.Second)
	defer cancel() // 步骤2: 确保函数结束时释放资源

	// 步骤3: 创建带 context 的请求
	req, err := http.NewRequestWithContext(ctx, "GET", url, nil)
	if err != nil {
		ch <- fmt.Sprintf("create request: %v", err)
		return
	}

	// 步骤4: 使用 http.DefaultClient.Do 发送请求
	res, err := http.DefaultClient.Do(req)
	if err != nil {
		ch <- fmt.Sprintf("fetch: %v", err)
		return
	}
	nbytes, err := io.Copy(io.Discard, res.Body)
	defer res.Body.Close()
	if err != nil {
		ch <- fmt.Sprintf("fetch reading %s: %v", url, err)
		return
	}
	timeElapsed := time.Since(start).Seconds()
	ch <- fmt.Sprintf("%.2fs  %7d  %s", timeElapsed, nbytes, url)
}
