package main

import (
	"fmt"
	"gorm.io/driver/mysql"
	"gorm.io/gorm"
	"log"
	"time"
)

// Product 定义结构体模型, 与数据库字段对应
type Product struct {
	// gorm.Model	// 可以自动带上四个默认字段, 个人不太建议用这个

	Id         uint64    `gorm:"primaryKey,column:id"` // 设置主键, 并指定列名
	No         string    `gorm:"column:no"`
	Name       string    `gorm:"column:name"`
	CreateTime time.Time `gorm:"column:create_time"`
	Deleted    bool      `gorm:"column:deleted"`
}

// TableName 设置表名
func (Product) TableName() string {
	return "t_product"
}

func main() {
	// 连接数据库, 此处会直接检查数据库配置, 如果连接失败则会返回 error
	dsn := "root:HongluoJdjhDB@12344@tcp(8.141.57.107:3306)/log?charset=utf8mb4&parseTime=True&loc=Local"
	db, err := gorm.Open(mysql.Open(dsn), &gorm.Config{})
	if err != nil {
		log.Panic(err)
	}

	// 连接池配置
	sqlDB, err := db.DB()
	// SetMaxIdleConns 设置空闲连接池中连接的最大数量
	sqlDB.SetMaxIdleConns(10)
	// SetMaxOpenConns 设置打开数据库连接的最大数量。
	sqlDB.SetMaxOpenConns(10)
	// SetConnMaxLifetime 设置了连接可复用的最大时间。
	sqlDB.SetConnMaxLifetime(time.Hour)

	// 创建表 自动迁移(把结构体和数据表进行对应)
	db.AutoMigrate(&Product{})

	// 新增
	p1 := Product{1, "123", "gotest1", time.Now(), false}
	// 如果结构体比较大可以传递指针 &p1
	db.Create(p1)

	// 查询第一条
	var p2 Product
	db.First(&p2)
	fmt.Printf("**********%#v", p2)

	// 更新
	db.Model(&p2).Update("name", "gotest111")

	// 删除
	db.Delete(&p2)

}
