package main

import (
	"context"
	"fmt"
	"golang.org/x/sync/singleflight"
	"sync"
	"sync/atomic"
	"time"
)

var count int32

func main() {
	//time.AfterFunc(1*time.Second, func() {
	//	atomic.AddInt32(&count, -count)
	//})
	articleID := 1
	var (
		wg  sync.WaitGroup
		now = time.Now()
		n   = 900

		sg  = &singleflight.Group{}
	)
	go func() {
		for {
			time.AfterFunc(1*time.Millisecond, func() {
				sg.Forget(fmt.Sprintf("%d", articleID))
			})
		}
	}()
	//time2 := time.AfterFunc(5*time.Millisecond, func() {
	//	sg.Forget(fmt.Sprintf("%d", articleID))
	//})
	//defer time2.Stop()
	fmt.Printf("count:%d\n", count)
	for i := 0; i < n; i++ {
		wg.Add(1)
		go func(i int) {
			//time.Sleep(time.Duration(count) * time.Millisecond)
			v, err, flag := sg.Do(fmt.Sprintf("%d", articleID), func() (interface{}, error) {
				return getArticle(articleID, sg)
			})
			if !flag {
				fmt.Printf("ret from fn execute\n")
			}
			if err != nil {
				fmt.Printf("err:%+v", err)
				panic("err")
			}
			//res, _ := getArticle(1)
			if v.(string) != "article: 1" {
				panic("err")
			}
			wg.Done()
		}(i)
	}

	wg.Wait()
	fmt.Printf("同时发起 %d 次请求，耗时: %s count:%d", n, time.Since(now), count)
}

func getArticle(id int,sg *singleflight.Group) (article string, err error) {
	fmt.Println("get begin")
	// 假设这里会对数据库进行调用, 模拟不同并发下耗时不同
	atomic.AddInt32(&count, 1)
	//if count == 2 {
	//	sg.Forget(fmt.Sprintf("%d", id))
	//}
	time.Sleep(time.Duration(count) * time.Millisecond)
	fmt.Println("get end")
	return fmt.Sprintf("article: %d", id), nil
}

//func singleflightGetArticle(sg *singleflight.Group, id int) (string, error) {
//	v, err, _ := sg.Do(fmt.Sprintf("%d", id), func() (interface{}, error) {
//		return getArticle(id)
//	})
//
//	return v.(string), err
//}



