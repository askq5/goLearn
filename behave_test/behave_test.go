package main

import (
	"fmt"
	"log"
	"testing"
	"unicode/utf8"
)

type structTest struct {
	mapTest   map[string]string
	sliceTest []int
}

func TestStructMapBehave(t *testing.T) {
	var test1 structTest
	fmt.Println(test1)
	if _, ok := test1.mapTest["fa"]; !ok {
		fmt.Println("ok")
	}
	if test1.mapTest == nil {
		test1.mapTest = make(map[string]string)
	}
	test1.mapTest["as"] = "as"
	test1.sliceTest = append(test1.sliceTest, 1)
	fmt.Println(test1)
}

func TestMapBehave(t *testing.T) {
	// map拷贝测试
	mapCopy1 := map[string]string{
		"test1": "map1",
	}
	mapCopy2 := mapCopy1
	mapCopy2["test2"] = "map2"
	mapCopy1["test2"] = "map1"

	fmt.Println(mapCopy1)
	fmt.Println(mapCopy2)

	mapTest1 := map[int]string{}
	mapTest2 := map[int][]int{}

	log.Printf("mapTest1[0]:%v\n", mapTest1[0])

	if mapTest2[0] == nil {
		log.Printf("nil mapTest2[0]:%v\n", mapTest2[0])
	} else {
		log.Printf("not nil mapTest2[0]:%v\n", mapTest2[0])
	}

	var t1 []int
	if t1 == nil {
		log.Printf("nil t1:%v\n", t1)
	} else {
		log.Printf("not nil t1:%v\n", t1)
	}
	t2 := []int{}
	//t2 := make([]int, 0)
	if t2 == nil {
		log.Printf("nil t2:%v\n", t2)
	} else {
		log.Printf("not nil t2:%v\n", t2)
	}
}

func TestSliceBehave(t *testing.T) {
	var s []string
	if s == nil {
		log.Println("s is nil")
	}
	log.Println(len(s))

	s2 := s[0:0:0]
	if s2 == nil {
		log.Println("s2 is nil")
	}
	log.Println(len(s2))

	s1 := make([]string, 0)
	if s1 == nil {
		log.Println("s1 is nil")
	}
	log.Println(len(s1))
}

func TestRuneBehave(t *testing.T) {
	str := "宋克强skq"
	log.Println(len(([]rune)(str)))
	log.Println(utf8.RuneCountInString(str))
}

func TestMakeAndNewBehave(t *testing.T) {
	a1 := new([]int32)
	a2 := make([]int32, 0)
	fmt.Println(a1 == nil)
	fmt.Println(*a1 == nil)
	fmt.Println(a2 == nil)
}
