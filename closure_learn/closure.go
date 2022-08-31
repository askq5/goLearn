package main

import (
	"fmt"
)

func main() {
	a := 0
	b := 0
	//retry(func() error {
	//	b = a + 1
	//	a++
	//	fmt.Println(a, b)
	//	return errors.New("test")
	//})
	fmt.Println(a, b)
	test1 := add5(a)
	fmt.Println(a, test1(), a)
	test2 := add5ByPointer(&a)
	fmt.Println(a, test2(), a)
	test3 := getMore5(a)
	fmt.Println(a, test3(), a)
}

func retry(target func() error) error {
	var err error
	for i := 0; i < 3; i++ {
		err = target()
		if err != nil {
			fmt.Printf("retry time :%d\n", i)
		} else {
			fmt.Println("success")
			break
		}

	}
	if err != nil {
		fmt.Println("failed")
	}
	return err
}

func add5(a int) func() int {
	return func() int {
		a += 5
		return a
	}
}

func add5ByPointer(a *int) func() int {
	return func() int {
		*a += 5
		return *a
	}
}

func getMore5(a int) func() int {
	return func() int {
		return a + 5
	}
}
