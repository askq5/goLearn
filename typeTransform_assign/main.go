package main

//
////type strStruct struct {
////	Str string
////}
//
//type strStruct []string
//
//type A strStruct
//type B strStruct
//
//func main() {
//	var a1 A
//	var pa *A
//	var b1 B
//	var pb *B
//	//str := strStruct{"abc"}
//	//b1 = strStruct{"abc"}
//
//	str := []string{"abc"}
//	b1 = []string{"abc"}
//
//	pb = &str
//	// a1 = b1  // x
//	a1 = A(b1)
//}
//
//func stringTest() {
//	type testStruct string
//	type A testStruct
//	type B testStruct
//	var a1 A
//	var pa *A
//	var b1 B
//	var pb *B
//	str := "abc"
//	b1 = str   //Cannot use 'str' (type string) as the type B
//	b1 = "abc" // ok 常量的一种转换
//}
//
//func sliceTest() {
//	type testStruct []string
//	type A testStruct
//	type B testStruct
//	var a1 A
//	var pa *A
//	var b1 B
//	var pb *B
//	str := []string{"abc"}
//	b1 = str
//	b1 = []string{"abc"}
//	pb = (*B)(&[]string{"abc"})
//}
//
//func mapTest() {
//
//	type testStruct map[int]string
//	type A testStruct
//	type B testStruct
//	var a1 A
//	var pa *A
//	var b1 B
//	var pb *B
//	str := make(map[int]string)
//	b1 = str
//	b1 = make(map[int]string)
//	pb = (*B)(&[]string{"abc"})
//}
//
//func chanTest() {
//
//	type testStruct chan int
//	type A testStruct
//	type B testStruct
//	var a1 A
//	var pa *A
//	var b1 B
//	var pb *B
//	str := make(chan int)
//	b1 = str
//	b1 = make(chan int)
//	pb = (*B)(&[]string{"abc"})
//}
//
//func intTest() {
//
//	type testStruct int
//	type A testStruct
//	type B testStruct
//	var a1 A
//	var pa *A
//	var b1 B
//	var pb *B
//	str := 5
//	b1 = str
//	b1 = make(int)
//	pb = (*B)(&[]string{"abc"})
//}
//
//func structTest() {
//	type testStruct struct {
//		Str string
//	}
//	type A testStruct
//	type B testStruct
//	var a1 A
//	var pa *A
//	var b1 B
//	var pb *B
//	str := testStruct{"abc"}
//	b1 = str               //Cannot use 'str' (type testStruct) as the type B
//	b1 = testStruct{"abc"} //Cannot use 'str' (type testStruct) as the type B
//	a1 = A(b1)
//	pa = (*A)(pb)
//}
