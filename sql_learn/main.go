package main

import (
	"code.byted.org/gorm/bytedgorm"
	"gorm.io/gorm"
)

func main() {

}

func dataBase() {

}

func gormDb() {
	DB, err := gorm.Open(bytedgorm.MySQL())

	// 预编译  用完就关掉（interpolateParams=false）
	// 针对密码的加解密插件开发
	// 批量数据创建，查询，更新
	// 批量数据加速操作，关闭默认事务、skiphooks跳过默认hooks、使用prepared statement、混合使用
	// 代码复用，分库分表

	// 关联的删除，孤儿数据
	// save的歧义 检查id是否被赋值来确定实际操作是update || insert
	// 接不接的住
}
