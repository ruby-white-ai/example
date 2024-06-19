package main

import (
	"fmt"
	"io/ioutil"
	"math"
	"math/rand"
	"os"
	"runtime"
	"sync"
	"time"
)

func SimulationSwitch(op string) {
	switch op {
	case "cpu":
		SimulationCPU()
	case "io":
		SimulationIO()
	case "wait":
		SimulationWaitForCPU()
	case "short":
		SimulationShortTasks()
	default:
		SimulationCPU()
	}
}

func SimulationCPU() {
	// 使用所有可用的CPU核心
	runtime.GOMAXPROCS(runtime.NumCPU())

	start := 1
	end := 1000000 // 计算100万以内的素数

	fmt.Printf("Calculating primes between %d and %d...\n", start, end)
	startTime := time.Now()

	primes := calculatePrimes(start, end)

	duration := time.Since(startTime)
	fmt.Printf("Found %d primes.\n", len(primes))
	fmt.Printf("Calculation took %s\n", duration)
}

func isPrime(n int) bool {
	if n <= 1 {
		return false
	}
	for i := 2; i <= int(math.Sqrt(float64(n))); i++ {
		if n%i == 0 {
			return false
		}
	}
	return true
}

// calculatePrimes 计算给定范围内的所有素数
func calculatePrimes(start, end int) []int {
	var primes []int
	for i := start; i <= end; i++ {
		if isPrime(i) {
			primes = append(primes, i)
		}
	}
	return primes
}

func SimulationIO() {
	var wg sync.WaitGroup

	// 生成随机数据
	data := make([]byte, fileSize)
	for i := range data {
		data[i] = byte(i % 256)
	}

	startTime := time.Now()

	// 启动多个并发的文件操作
	for i := 0; i < numFiles; i++ {
		wg.Add(1)
		go createAndWriteFile(fmt.Sprintf("file-%d.tmp", i), data, &wg)
	}

	// 等待所有文件操作完成
	wg.Wait()

	duration := time.Since(startTime)
	fmt.Printf("I/O operations completed in %s\n", duration)

	// 清理临时文件
	for i := 0; i < numFiles; i++ {
		os.Remove(fmt.Sprintf("file-%d.tmp", i))
	}
}

const (
	numFiles = 100         // 要创建和写入的文件数量
	fileSize = 1024 * 1024 // 每个文件的大小（1MB)

	numGoroutines = 100000 // 要创建的 goroutine 数量

	numTasks  = 1000000 // 要创建的任务数量
	taskBatch = 10000   // 每批次创建的任务数量
)

func createAndWriteFile(filename string, data []byte, wg *sync.WaitGroup) {
	defer wg.Done()

	// 创建文件
	file, err := os.Create(filename)
	if err != nil {
		fmt.Printf("Error creating file %s: %v\n", filename, err)
		return
	}
	defer file.Close()

	// 写入数据
	_, err = file.Write(data)
	if err != nil {
		fmt.Printf("Error writing to file %s: %v\n", filename, err)
		return
	}

	// 读取数据
	_, err = ioutil.ReadFile(filename)
	if err != nil {
		fmt.Printf("Error reading file %s: %v\n", filename, err)
		return
	}
}

func SimulationWaitForCPU() {
	var wg sync.WaitGroup

	// 设置最大可用的 CPU 核心数
	runtime.GOMAXPROCS(runtime.NumCPU())

	startTime := time.Now()

	// 启动大量的 goroutine
	for i := 0; i < numGoroutines; i++ {
		wg.Add(1)
		go worker(i, &wg)
	}

	// 等待所有 goroutine 完成
	wg.Wait()

	duration := time.Since(startTime)
	fmt.Printf("All goroutines finished in %s\n", duration)
}

func worker(id int, wg *sync.WaitGroup) {
	defer wg.Done()

	// 模拟一些工作
	workTime := rand.Intn(100) + 10 // 随机工作时间，范围 10-110 毫秒
	time.Sleep(time.Duration(workTime) * time.Millisecond)

	// 模拟等待
	waitTime := rand.Intn(1000) + 100 // 随机等待时间，范围 100-1100 毫秒
	time.Sleep(time.Duration(waitTime) * time.Millisecond)

	// 再次执行一些工作
	workTime = rand.Intn(100) + 10 // 随机工作时间，范围 10-110 毫秒
	time.Sleep(time.Duration(workTime) * time.Millisecond)

	fmt.Printf("Goroutine %d finished\n", id)
}

func SimulationShortTasks() {
	var wg sync.WaitGroup

	// 设置最大可用的 CPU 核心数
	runtime.GOMAXPROCS(runtime.NumCPU())

	startTime := time.Now()

	// 使用无缓冲通道来限制并发任务数量
	ch := make(chan struct{}, taskBatch)

	// 启动大量的短任务
	for i := 0; i < numTasks; i++ {
		wg.Add(1)
		go shortTask(&wg, ch)
		// 等待通道有空位
		ch <- struct{}{}
	}

	// 等待所有任务完成
	wg.Wait()

	// 关闭通道
	close(ch)

	duration := time.Since(startTime)
	fmt.Printf("All tasks finished in %s\n", duration)
}

func shortTask(wg *sync.WaitGroup, ch chan struct{}) {
	defer wg.Done()

	// 模拟短时间的计算任务
	sum := 0
	for i := 0; i < 1000; i++ {
		sum += i
	}

	// 向通道发送完成信号
	ch <- struct{}{}
}
