package main

import "fmt"

type Handler interface {
	opt(interface{}) (interface{}, error)
}

type numAdd struct {
}

func (s numAdd) opt(p interface{}) (interface{}, error) {
	arr, ok := p.([2]int)
	if !ok {
		return nil, fmt.Errorf("wrong param: %#v", p)
	}
	return arr[0] + arr[1], nil
}

// 实现了接口的函数类型，简称为接口型函数。
type HandlerFunc func(interface{}) (interface{}, error)

func (funcAdap HandlerFunc) opt(p interface{}) (interface{}, error) {
	return funcAdap(p)
}

func numSub(p interface{}) (interface{}, error) {
	arr, ok := p.([2]int)
	if !ok {
		return nil, fmt.Errorf("wrong param: %#v", p)
	}
	return arr[0] - arr[1], nil
}

var funcHandlers map[string]Handler

func eachFunc(params map[string]interface{}, ops map[string]Handler) {
	for k, op := range params {
		handler, ok := ops[k]
		if !ok {
			fmt.Printf("no %s op\n", k)
			continue
		}
		str, err := handler.opt(op)
		if err != nil {
			fmt.Printf("%s false: params is %#v\n", k, params[k])
			continue
		}
		fmt.Printf("%v\n", str)
	}
	return
}

func main() {
	funcHandlers = make(map[string]Handler)
	funcHandlers["add"] = HandlerFunc(numSub)
	funcHandlers["sub"] = numAdd{}

	funcParams := make(map[string]interface{})
	funcParams["add"] = [2]int{1, 2}
	funcParams["sub"] = [2]int{1, 2}
	funcParams["multi"] = [2]int{1, 2}
	eachFunc(funcParams, funcHandlers)
}
