package main

import (
	"errors"
	"fmt"
	"sync"
	"sync/atomic"

	"golang.org/x/sync/errgroup"
)

func main() {
	//test()
	//test2()
	test5()
	//test3()
	//test4()

}
func test() {
	var a = 0
	wg := sync.WaitGroup{}
	wg.Add(100)
	for i := 0; i < 100; i++ {
		go func() {
			for j := 0; j < 10000; j++ {
				a++
			}
			wg.Done()
		}()
	}
	wg.Wait()
	fmt.Println(a)
}

func test1() {
	a := int64(0)
	wg := sync.WaitGroup{}
	wg.Add(100)
	for i := 0; i < 100; i++ {
		go func() {
			for j := 0; j < 10000; j++ {
				atomic.AddInt64(&a, 1)
			}
			wg.Done()
		}()
	}
	wg.Wait()
	fmt.Println(a)
}

func test2() {

	nums := make([]int64, 10)
	for i := 0; i < 10; i++ {
		nums[i] = int64(i)
	}
	wg := sync.WaitGroup{}
	for _, num := range nums {
		wg.Add(1)
		go func() {
			defer wg.Done()
			println(num)
		}()
	}
	wg.Wait()
}

func test3() {

	nums := make([]int64, 10)
	for i := 0; i < 10; i++ {
		nums[i] = int64(i)
	}
	wg := sync.WaitGroup{}
	for _, num := range nums {
		wg.Add(1)
		tempNum := num
		go func() {
			defer wg.Done()
			println(tempNum)
		}()
	}
	wg.Wait()
}

func test4() {
	nums := make([]int64, 10)
	for i := 0; i < 10; i++ {
		nums[i] = int64(i)
	}
	wg := sync.WaitGroup{}
	for _, num := range nums {
		wg.Add(1)
		go func(num int64) {
			defer wg.Done()
			println(num)
		}(num)
	}
	wg.Wait()
}

func test5() {
	nums := make([]int64, 10)
	for i := 0; i < 10; i++ {
		nums[i] = int64(i)
	}
	eg := errgroup.Group{}
	for _, num := range nums {
		tempNum := num
		eg.Go(func() error {
			println(num)
			fmt.Printf("num:%d, tempNum:%d\n", num, tempNum)
			if tempNum != num {
				fmt.Printf("num:%d, tempNum:%d\n", num, tempNum)
				return errors.New("race detector")
			}
			return nil
		})
	}

	if err := eg.Wait(); err != nil {
		fmt.Printf("error:%+v", err)
	}
}
