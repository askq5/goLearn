package main

import "testing"

func TestRedis(t *testing.T) {
	// 测试set 是val用"" 和 int 在incr和get时区别
	// ""  get  不明确
	// ""  incr get
	// int get 明确预期
	// int incr get 明确预期

}
