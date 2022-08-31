package main

import (
	"fmt"
	"time"
	"unsafe"
	/*"reflect"
	"time"*/)

func changeString(s string) {
	fmt.Printf("inner: %v, %v\n", s, &s)
	s = "b"
	fmt.Printf("inner: %v, %v\n", s, &s)
}

func sizeofVar() {

	var s1 []int
	s2 := make([]int, 0)
	var s3 = []int{}

	println("s1", len(s1), cap(s1), s1 == nil, &s1, unsafe.Sizeof(s1), s1[:], s1[:] == nil)
	println("s2", len(s2), cap(s2), s2 == nil, &s2, unsafe.Sizeof(s2), s2[:], s2[:] == nil)
	println("s3", len(s3), cap(s3), s3 == nil, &s3, unsafe.Sizeof(s3), s3[:], s3[:] == nil)
}

//chanTicker 测试chan 多个case可用时的调度
//time.NewTicker特性 维护一个定时有time.Time类型消息可接受的通道
//读关闭通道
func chanTicker() {
	donec := make(chan bool, 1)
	defer close(donec)
	go func() {
		time.Sleep(12 * time.Second)
		donec <- true
	}()

	tick1 := time.NewTicker(1 * time.Second)
	tick2 := time.NewTicker(2 * time.Second)
	tick3 := time.NewTicker(3 * time.Second)
	var t1, t2, t3 time.Time
	var doneFlag bool
	for {
		select {
		case t1 = <-tick1.C:
			fmt.Printf("1 Second t1:%s t2:%s t3:%s\n", t1, t2, t3)
			t1 = time.Time{}
		case t2 = <-tick2.C:
			fmt.Printf("2 Second t1:%s t2:%s t3:%s\n", t1, t2, t3)
			t2 = time.Time{}
		case t3 = <-tick3.C:
			fmt.Printf("3 Second t1:%s t2:%s t3:%s\n", t1, t2, t3)
			t3 = time.Time{}
		}
		time.Sleep(500 * time.Millisecond)
		select {
		case doneFlag = <-donec:
			fmt.Println("done")
		default:
			fmt.Println("goon")
		}
		if doneFlag {
			break
		}
	}
}

//chanAfter 任务定时器
func chanLearn() {
	ch := make(chan time.Time)
	go func(ch chan time.Time) {
		for {
			tick := time.NewTicker(2 * time.Second)

			//time.After(2*time.Second)
			//time.Sleep(4 * time.Second)
			a := <-tick.C
			//fmt.Println(reflect.TypeOf(a))
			ch <- a
		}
	}(ch)
	for {
		select {
		case <-time.After(2 * time.Second):
			fmt.Println("2 second")
		case chFlag := <-ch:
			//do something
			fmt.Println(chFlag)
		}

	}
}

const (
	Sunday     = iota //0
	Monday            //1
	Tuesday           //2
	Wedenesday        //3
	Thursday          //4
	Friday            //5
	Saturday          //6

	//同Go语言的其他符号（symbol）一样，以大写字母开头的常量在包外可见。
	//以上例子中 numberOfDays 为包内私有，其他符号则可被其他包访问。
	numberOfDays
)

func enumLearn() {

	fmt.Println("Sunday=", Sunday)
	fmt.Println("Monday=", Monday)
	fmt.Println("Tuesday=", Tuesday)
	fmt.Println("Wedenesday=", Wedenesday)
	fmt.Println("Thursday=", Thursday)
	fmt.Println("Friday=", Friday)
	fmt.Println("Saturday=", Saturday)
	fmt.Println("numberOfDays=", numberOfDays)
}

func main() {

	//sizeofVar()

	chanTicker()

	//chanLearn()
	//chanLearn()
	/*s := "a"
	fmt.Printf("outer: %v, %v\n",s, &s)
	changeString(s)
	fmt.Printf("outer: %v, %v\n",s, &s)*/

	//enumLearn()

}
