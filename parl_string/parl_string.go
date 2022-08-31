package main

import (
	"encoding/json"
	"fmt"
)

const (
	Askq = "askq test"
	Skq  = "skq"
)

type ParalStruct struct {
	Name string
}

func main() {
	// runtime.GOMAXPROCS() //只使用1个物理处理器
	testStruct1 := ParalStruct{
		Name: "b",
	}

	go func() {
		for {
			testStruct1.Name = Askq
			//time.Sleep(10)
		}
	}()

	go func() {
		for {
			testStruct1.Name = Skq
			//time.Sleep(10)
		}
	}()
	for {
		tempStr, _ := json.Marshal(testStruct1)
		fmt.Println(string(tempStr))
		//time.Sleep(10)
	}
}
