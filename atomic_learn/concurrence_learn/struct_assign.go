package main

import (
	"fmt"
	"runtime"
	"sync"
	"sync/atomic"
)

var (
	total int64 = 1 << 20
)

type StructTest struct {
	X int
	Y int
}

func int64Assign() {
	var g uint32
	for i := int64(0); i < total; i++ {
		var wg sync.WaitGroup
		// 协程 1
		wg.Add(1)
		go func() {
			defer wg.Done()
			g = 1
		}()

		// 协程 2
		wg.Add(1)
		go func() {
			defer wg.Done()
			g = 2
		}()
		wg.Wait()

		// 赋值异常判断
		if !(g != 2 || g != 3) {
			fmt.Printf("concurrent int64 assignment error, i=%v g=%d\n", i, g)
			return
		}
	}
	fmt.Printf("concurrent int64 assignment success. g=%d\n", g)
}

func structAssign() {
	var g StructTest

	for i := int64(0); i < total; i++ {
		var wg sync.WaitGroup
		// 协程 1
		wg.Add(1)
		go func() {
			defer wg.Done()
			g = StructTest{1, 2}
		}()

		// 协程 2
		wg.Add(1)
		go func() {
			defer wg.Done()
			g = StructTest{3, 4}
		}()
		wg.Wait()

		// 赋值异常判断
		if !((g.X == 1 && g.Y == 2) || (g.X == 3 && g.Y == 4)) {
			fmt.Printf("concurrent struct assignment error, i=%v g=%+v", i, g)
			break
		}
	}
}

func stringAssign(str1, str2 string) {
	var s string
	for i := int64(0); i < total; i++ {
		var wg sync.WaitGroup
		// 协程 1
		wg.Add(1)
		go func() {
			defer wg.Done()
			s = str1
		}()

		// 协程 2
		wg.Add(1)
		go func() {
			defer wg.Done()
			s = str2
		}()
		wg.Wait()

		// 赋值异常判断
		if s != str1 && s != str2 {
			fmt.Printf("concurrent string assignment error, i=%v lenS=%d s=%v lenStr1=%d str1=%s lenStr2=%d str=%s\n", i, len(s), s,
				len(str1), str1, len(str2), str2)
			return
		}
	}
	fmt.Printf("concurrent string assignment success s=%s lenStr1=%d str1=%s lenStr2=%d str2=%s\n", s, len(str1), str1, len(str2), str2)
}

func atomicDisplay() {
	procNum := 1
	wNum := 100
	runtime.GOMAXPROCS(procNum)

	var w sync.WaitGroup
	count := int32(0)
	w.Add(wNum)
	for i := 0; i < wNum; i++ {
		go func() {
			for j := 0; j < 20; j++ {
				count++
			}
			w.Done()
		}()
	}
	w.Wait()
	fmt.Println(count)
}

func countWithoutAtomic() {
	var cnt int32 = 0
	wg := sync.WaitGroup{}
	for i := 0; i < 100; i++ {
		wg.Add(1)
		go func() {
			for j := 0; j < 100; j++ {
				cnt++
			}
			wg.Done()
		}()
	}
	wg.Wait()
	fmt.Printf("countWithoutAtomic:%d\n", cnt)
	return
}

func countWithAtomic() {
	var cnt int32 = 0
	wg := sync.WaitGroup{}
	for i := 0; i < 100; i++ {
		wg.Add(1)
		go func() {
			for j := 0; j < 100; j++ {
				atomic.AddInt32(&cnt, 1)
			}
			wg.Done()
		}()
	}
	wg.Wait()
	fmt.Printf("countWithAtomic:%d\n", cnt)
	return
}

// go run -race .
func main() {
	// int64 赋值原子性测试
	int64Assign()
	// 原子操作演示
	//atomicDisplay()
	// 结构体测试
	//structAssign()
	// string 测试
	//str1 := "123456789"
	//str2 := "0123456789"
	//stringAssign(str1, str2)

	// i++和 addint比较
	//countWithoutAtomic()
	//countWithAtomic()
	fmt.Printf("%x", ^uintptr(0))
}
