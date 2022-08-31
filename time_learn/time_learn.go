package main

import (
	"fmt"
	"strconv"
	"time"
)

const (
	MaxAuditTime = 1000 * 60 * 60 * 12

	FmtMSEC          = "2006-01-02 15:04:05.000"
	FmtTime          = "2006-01-02 15:04:05"
	FmtDate          = "20060102"
	LocationShanghai = "Asia/Shanghai"
	FmtPushDate      = "2006-01-02"
	FmtPushHourMin   = "15:04"
)

// 轮询时间：1647525101
// 1647592210

// 计算时间差（天数）
func SubDays(t1, t2 time.Time) (day int) {
	day = int(t1.Sub(t2).Hours() / 24)
	return
}

// 利用Add()计算具体时间点：
func AddTimeDuration() {
	timeDuration, _ := time.ParseDuration("-10m") // 参数有多种选择，自行百度...
	beginTime := time.Now().Add(timeDuration).Format("2006-01-02 15:04:05")
	fmt.Println("前推10分钟：", beginTime)
}

func TimeDiff() {
	// (1) time.Sub() - 计算两个time类型日期相差天数：
	timeFormat := "2006-01-02 15:04:05"                             // 默认转换日期格式
	beginTime, err := time.Parse(timeFormat, "2020-09-08 10:00:00") // string -> time
	if err != nil {
		fmt.Println("time.Parse error")
	}
	beginDay, _ := strconv.Atoi(beginTime.Format("20060102"))

	endTime, err := time.Parse(timeFormat, "2020-09-18 6:00:00") // string -> time
	if err != nil {
		fmt.Println("time.Parse error")
	}

	endDay, _ := strconv.Atoi(endTime.Format("20060102"))
	diffDay := endDay - beginDay
	fmt.Println(diffDay)
	subDays := SubDays(endTime, beginTime) // 参数1：结束日期:	参数2：开始靠前
	fmt.Println("相差天数1 = ", subDays)
}

func main() {
	//1632758400
	timeInt := int64(1660665600)
	fmt.Println(time.Unix(timeInt, 0))

	timeDate := time.Unix(timeInt, 0).Format(FmtMSEC)
	fmt.Println(timeDate)
	//
	test := time.Date(2022, time.August, 29, 0, 0, 0, 0, time.Local).Unix()
	fmt.Println(test)
	fmt.Println(time.Unix(test, 0))

	cannyTimeLine := time.Date(2122, time.May, 26, 0, 0, 0, 0, time.Local).Unix()
	fmt.Println(cannyTimeLine)
	// 7.50--1643586600
	// 11---1643598000
	// 14--1643608800
	// 16--1643616000
	// 17--1643619600
	// 20--1643630400
	//test1 := time.Date(2022, 1, 31, 0, 0, 1, 0, time.Local).UnixNano() / 1e6
	//fmt.Println(test1)
	//fmt.Println(time.Unix(test1/1e3, 1e9))

	//end=2022/01/31-05:59:59&m=sum:rate{counter}:ies.dm.ploy.script.tiger_sixty.new_round&start=2022/01/30-23:49:59
	//group	result	computations
	//{ }
	//84348025.5202865
	//30 * sum(q("sum:rate{counter}:ies.dm.ploy.script.tiger_sixty.new_round{}", "1643586600", "1643608800"))	84.34803M
	//sum(q("sum:rate{counter}:ies.dm.ploy.script.tiger_sixty.new_round{}", "1643586600", "1643608800"))	2.8116M
	time1 := time.Now()
	t1 := time1.Unix()
	fmt.Println(t1)
	time2 := time.Now().AddDate(0, 0, 2)
	fmt.Println(time1.Before(time2))
	t2 := time2.UnixNano() / 1e6
	fmt.Println(time1, t1, time2, t2, (t2-t1)/MaxAuditTime, MaxAuditTime)
	now := time.Now()
	//now1  := time.Now().Location()
	location, _ := time.LoadLocation("Asia/Shanghai")
	now1 := time.Now().In(location)
	pushTime1 := time.Now().AddDate(0, 0, 2)
	min1, _ := time.ParseDuration("1m")
	pushTime2 := time.Now().Add(min1)
	pushTime3 := time.Now().Add(time.Minute * 1)
	fmt.Println(now, now1, pushTime1, pushTime2, pushTime3)
	fmt.Println("1分钟时间差", pushTime2.Sub(now))
	nowFmtMsec := now.Format(FmtMSEC)
	nowFmtDate := now.Format(FmtDate)
	fmt.Println(nowFmtMsec, nowFmtDate)
}
