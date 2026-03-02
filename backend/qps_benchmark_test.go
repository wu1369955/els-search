package main

import (
	"fmt"
	"net/http"
	"sync"
	"testing"
	"time"
)

func TestQPSBenchmark(t *testing.T) {
	// 测试配置
	url := "http://localhost/api/search?index=web_text_zh&q=测试&fields=title,content&page=1&size=10"
	concurrency := 100     // 并发数
	totalRequests := 10000 // 总请求数

	// 统计变量
	var wg sync.WaitGroup
	var mu sync.Mutex
	var completedRequests int
	var failedRequests int
	var totalResponseTime time.Duration

	// 开始时间
	startTime := time.Now()

	// 启动并发请求
	for i := 0; i < concurrency; i++ {
		wg.Add(1)
		go func() {
			defer wg.Done()
			client := &http.Client{Timeout: 10 * time.Second}

			// 每个goroutine处理的请求数
			requestsPerGoroutine := totalRequests / concurrency

			for j := 0; j < requestsPerGoroutine; j++ {
				reqStart := time.Now()
				req, err := http.NewRequest("GET", url, nil)
				if err != nil {
					mu.Lock()
					failedRequests++
					mu.Unlock()
					continue
				}

				resp, err := client.Do(req)
				if err != nil {
					mu.Lock()
					failedRequests++
					mu.Unlock()
					continue
				}

				if resp.StatusCode != http.StatusOK {
					mu.Lock()
					failedRequests++
					mu.Unlock()
					resp.Body.Close()
					continue
				}

				resp.Body.Close()

				reqDuration := time.Since(reqStart)
				mu.Lock()
				completedRequests++
				totalResponseTime += reqDuration
				mu.Unlock()
			}
		}()
	}

	// 等待所有请求完成
	wg.Wait()

	// 计算结果
	totalTime := time.Since(startTime)
	qps := float64(completedRequests) / totalTime.Seconds()

	// 输出结果
	fmt.Printf("QPS测试结果:\n")
	fmt.Printf("总请求数: %d\n", totalRequests)
	fmt.Printf("成功请求数: %d\n", completedRequests)
	fmt.Printf("失败请求数: %d\n", failedRequests)
	fmt.Printf("总耗时: %v\n", totalTime)
	fmt.Printf("QPS: %.2f\n", qps)

	if completedRequests > 0 {
		averageResponseTime := totalResponseTime / time.Duration(completedRequests)
		fmt.Printf("平均响应时间: %v\n", averageResponseTime)
	} else {
		fmt.Printf("平均响应时间: N/A (无成功请求)\n")
	}
}
