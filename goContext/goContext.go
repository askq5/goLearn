package main

import (
	"context"
	"fmt"
	"log"
	"strconv"
	"time"
)

func contextTree() {
	root := context.Background()
	ctx1 := context.WithValue(root, "k1", 1111)
	ctx2, cancel1 := context.WithTimeout(ctx1, 4*time.Second)
	defer cancel1()
	ctx3, cancel2 := context.WithTimeout(ctx2, 2*time.Second)
	defer cancel2()
	ctx4 := context.WithValue(ctx2, "k2", "22222")
	ctx5 := context.WithValue(ctx4, "k3", "33333")
	ctx6 := context.WithValue(ctx3, "k4", "4444")
	ctx7, cancel3 := context.WithDeadline(ctx4, time.Now().Add(3*time.Second))
	defer cancel3()
	ctx8, cancel4 := context.WithCancel(ctx6)
	defer cancel4()

	log.Println(ctx7.Value("k1"))
	fmt.Println(ctx6.Value("k4"))
	fmt.Println(ctx5.Value("k2"))

	fmt.Println(ctx7)

	go func() {
		select {
		case <-ctx2.Done():
			fmt.Println("ctx2 canceled")
		}
	}()

	go func() {
		select {
		case <-ctx3.Done():
			fmt.Println("ctx3 canceled")
		}
	}()

	go func() {
		select {
		case <-ctx1.Done():
			fmt.Println("ctx1 canceled")
		}
	}()

	go func() {
		select {
		case <-ctx4.Done():
			fmt.Println("ctx4 canceled")
		}
	}()

	go func() {
		select {
		case <-ctx7.Done():
			fmt.Println("ctx7 canceled")
		}
	}()

	go func() {
		select {
		// ctx8 canceled before ctx7, because ctx3 canceled
		case <-ctx8.Done():
			fmt.Println("ctx8 canceled")
		}
	}()

	time.Sleep(8 * time.Second)
}

func contextTimeout() {
	// 创建一个子节点的context,3秒后自动超时
	ctxRoot := context.Background()
	ctx1, cancel1 := context.WithTimeout(ctxRoot, time.Second*3)
	defer cancel1()
	ctx2 := context.WithValue(ctxRoot, "k-2", "22")
	ctx3 := context.WithValue(ctx1, "k-3", "33")
	go watch(ctxRoot, "监控0")
	go watch(ctx1, "监控1")
	go watch(ctx2, "监控2")
	go watch(ctx3, "监控3")
	fmt.Println("现在开始等待8秒,time=", time.Now().Unix())
	time.Sleep(8 * time.Second)

	fmt.Println("等待8秒结束,准备调用cancel()函数，发现两个子协程已经结束了，time=", time.Now().Unix())

}

func main() {
	//contextTree()

	//contextTimeout()
	root := context.Background()
	ctx1 := context.WithValue(root, HeaderStressStarlightReward, "3")
	fmt.Println(GetStressStarlightReward(ctx1))

	ctx2 := context.WithValue(root, CtxKeyStressLotteryExtra, "0")
	fmt.Println(GetStressLotteryExtra(ctx2))
}

const (
	HeaderStressStarlightReward = "stress_starlight_reward"
	CtxKeyStressLotteryExtra    = "ctx_starlight_reward"
)

// getStrCtx read the value of key in ctx, return it in string type.
func getStrCtx(ctx context.Context, key string) string {
	if ctx == nil {
		return ""
	}
	v := ctx.Value(key)
	switch v := v.(type) {
	case string:
		return v
	case *string:
		return *v
	}
	return ""
}

// GetStressStarlightReward 获取指定的星光奖励在tcc列表的位置
func GetStressStarlightReward(ctx context.Context) int {
	stressStr := getStrCtx(ctx, HeaderStressStarlightReward)
	if stressStr != "" {
		rewardPos, err := strconv.ParseInt(stressStr, 10, 64)
		if err == nil && rewardPos > 0 && rewardPos < 9 {
			return int(rewardPos - 1)
		}
	}
	return 0
}

// GetStressLotteryExtra 指定是否获得抽奖额外任务
func GetStressLotteryExtra(ctx context.Context) bool {
	stressStr := getStrCtx(ctx, CtxKeyStressLotteryExtra)
	if stressStr != "" {
		extraFlag, err := strconv.ParseInt(stressStr, 10, 64)
		if err == nil && extraFlag == 1 {
			return true
		}
	}
	return false
}

// 单独的监控协程
func watch(ctx context.Context, name string) {
	for {
		select {
		case <-ctx.Done():
			fmt.Println(name, "收到信号，监控退出,time=", time.Now().Unix())
			return
		default:
			fmt.Println(name, "goroutine监控中,time=", time.Now().Unix())
			time.Sleep(1 * time.Second)
		}
	}
}
