package main

import (
	"context"
	"log"
	"os"
	"time"

	"gorm.io/gorm/logger"

	"gorm.io/driver/sqlite"
	"gorm.io/gorm"
)

type TableTest struct {
	gorm.Model
	Code  string
	Price uint
	Title string
}

func main() {
	ctx := context.Background()
	logger := logger.New(
		log.New(os.Stdout, "\r\n", log.LstdFlags), // io writer（日志输出的目标，前缀和日志包含的内容——译者注）
		logger.Config{
			SlowThreshold:             time.Second, // 慢 SQL 阈值
			LogLevel:                  logger.Info, // 日志级别 Silent、Error、Warn、Info
			IgnoreRecordNotFoundError: false,       // 忽略ErrRecordNotFound（记录未找到）错误
			Colorful:                  false,       // 禁用彩色打印
		},
	)
	option := &gorm.Config{
		Logger: logger,
	}
	//dsn := "host=localhost user=gorm password=gorm dbname=gorm port=9920 sslmode=disable TimeZone=Asia/Shanghai"
	db, _ := gorm.Open(sqlite.Open("gorm.db"), option)
	var err error

	//db.AutoMigrate(&TableTest{})
	//user1 := TableTest{
	//	Code:  "aaa",
	//	Price: 100,
	//}
	//user2 := TableTest{
	//	Code:  "bbb",
	//	Price: 200,
	//}
	//db.FirstOrCreate(&user1)
	//db.FirstOrCreate(&user2)
	//var test []TableTest

	//err = db.Table("table_tests").WithContext(ctx).Where("code = ?", "'aaa' or '1'='1'").Find(&test).Error
	//if err != nil {
	//	log.Printf("err:%+v\n", err)
	//} else {
	//	log.Printf("test:%+v\n", test)
	//}
	//
	//err = db.WithContext(ctx).Where("code = ?", "ccc").Find(&test).Error
	//if err != nil {
	//	log.Printf("err:%+v\n", err)
	//} else {
	//	log.Printf("test:%+v\n", test)
	//}
	//
	//db.Model(&test).Update("Price", 2000)
	//err = db.WithContext(ctx).Where("code = ?", "ccc").Updates(TableTest{Price: 1000, Title: "ccc"}).Error

	type TempTableTest struct {
		gorm.Model
		Code string
	}
	tempUser := TempTableTest{
		Code: "eee",
	}
	db.Table("table_tests").AutoMigrate(&TempTableTest{})
	err = db.WithContext(ctx).Table("table_tests").FirstOrCreate(&tempUser, "Code", "eee").Error
	log.Println(err)
	err = db.Table("table_tests").Create(&tempUser).Error
}
