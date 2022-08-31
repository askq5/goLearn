package main

import (
	"fmt"
	"sort"
)

type people struct {
	id   int32
	name string
}

func main() {
	array := []people{{1, "1"}, {4, "4"}, {2, "2"}}
	sort.Slice(array, func(i, j int) bool {
		return array[i].id < array[j].id
	})
	fmt.Println(array)
}
