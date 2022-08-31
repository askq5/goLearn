package main

import (
	"log"
	"runtime"
)

func main() {
	defer RecoverLearn()
	for {
		panic("recover test1")
	}
	log.Println("panic test1")
}

func RecoverLearn() {
	if e := recover(); e != nil {
		const size = 64 << 10
		buf := make([]byte, size)
		buf = buf[:runtime.Stack(buf, false)]
		log.Println("recover test2")
		log.Printf("ruttime.stack:%s", buf)
	}
}
