package main

import (
	"fmt"
	"reflect"
)

//type strList []string

type strStruct struct {
	str string
}

type structList []strStruct

func main() {
	tag := &structList{}
	tempTag1 := &[]strStruct{{"abc"}}
	*tag = *tempTag1
	fmt.Println(reflect.TypeOf(*tag) == reflect.TypeOf(*tempTag1))
}

//var tag *strList
//tag := &strList{}
//tempTag1 := &[]string{"asb"}
//tag = tempTag1
//*tag = *tempTag1
//fmt.Printf("tag:%+v\n", tag)
//
//temp1 := []string{"test"}
//var temp strList
//temp = temp1
//fmt.Println(temp)
