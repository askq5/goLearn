package main

import "fmt"

// escape to heap
func deferLearn1() {
	fmt.Println("defer learn1")
	var whatever [5]int
	for i, _ := range whatever {
		defer func() {
			fmt.Println(i)
		}()
	}
}

func deferLearn2() {
	fmt.Println("defer learn2")
	var whatever [5]int
	for i, _ := range whatever {
		defer func(i int) {
			fmt.Println(i)
		}(i)
	}
}

func main() {
	//deferLearn1()
	//deferLearn2()
	//defer 符合先入后出的规则
	fmt.Println(deferRet1())
	fmt.Println(deferRet2())
	// 无论是否在返回值中显示声明返回值，return都会通过值拷贝的形式填充返回参数，而且defer在return后面，最后参会把参数返回
}

func f1() (r int) {
	defer func() {
		r++
	}()
	return 0
}

// 提前声明返回值
func deferRet1() (r int) {
	fmt.Printf("r1:%d\n", r)
	t := 5
	defer func() {
		t = t + 5
		fmt.Printf("t:%d\n", t)
		fmt.Printf("r2:%d\n", r)
	}()
	return t
}

// 不提前声明返回值
func deferRet2() int {
	t := 5
	defer func() {
		t = t + 5
		fmt.Printf("t:%d", t)
	}()
	return t
}

func f3() (r int) {
	defer func(r int) {
		r = r + 5
	}(r)
	return 1
}
