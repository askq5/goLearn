package main

import "unsafe"

func main() {
	//s1 := []int64{1, 2, 3}
	//s2 := []int64{4, 5, 6}
	//s2[0] = s1[0] + 3
	//fmt.Println(s1)
	str := "test"
	//fmt.Printf("%s", str)
	bytes := []byte{'a', 'b'}
	bytes[0] = '1'
	testByte := []byte(str)
	testByte[0] = '2'
	testStr := *(*string)(unsafe.Pointer(&bytes))
	testByte1 := []byte(testStr)
	testByte1[0] = '1'
}
