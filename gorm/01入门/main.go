package main

import (
	"database/sql"
	"time"

	"gorm.io/driver/mysql"
	"gorm.io/gorm"
)

// 模型定义
type User struct {
	Model
	Name         string         // 可传空字符串
	Email        *string        // 可传空指针
	Birthday     *time.Time     // 可传空指针
	MemberNumber sql.NullString // 可传空字符串
	ActivatedAt  sql.NullTime   // 可传空时间
}

// gorm.Model 的定义
type Model struct {
	ID        uint `gorm:"primaryKey"`
	CreatedAt time.Time
	UpdatedAt time.Time
	DeletedAt gorm.DeletedAt `gorm:"index"`
}

// 正常的结构体字段，你也可以通过标签 embedded 将其嵌入
// 使用这种模式，所有的默认设置都会失效，比如自增，默认时间，软删除等
type User2 struct {
	Model        Model          `gorm:"embedded"`
	Name         string         // 可传空字符串
	Email        *string        // 可传空指针
	Birthday     *time.Time     // 可传空指针
	MemberNumber sql.NullString // 可传空字符串
	ActivatedAt  sql.NullTime   // 可传空时间
}

func main() {
	// 连接数据库：方式1
	// dsn := "remote:mysql_remote@tcp(192.168.40.10:3306)/web3?charset=utf8mb4&parseTime=True&loc=Local"
	// db, err := gorm.Open(mysql.Open(dsn), &gorm.Config{})
	// if err == nil {
	// 	db.AutoMigrate(&User{})
	// }
	// 连接数据库：方式2
	db, err := gorm.Open(mysql.New(mysql.Config{
		DSN:                       "remote:mysql_remote@tcp(192.168.40.10:3306)/web3?charset=utf8mb4&parseTime=True&loc=Local",
		DefaultStringSize:         256,   // string 类型字段的默认长度
		DisableDatetimePrecision:  true,  // 禁用 datetime 精度，MySQL 5.6 之前的数据库不支持
		DontSupportRenameIndex:    true,  // 重命名索引时采用删除并新建的方式，MySQL 5.7 之前的数据库和 MariaDB 不支持重命名索引
		DontSupportRenameColumn:   true,  // 用 `change` 重命名列，MySQL 8 之前的数据库和 MariaDB 不支持重命名列
		SkipInitializeWithVersion: false, // 根据当前 MySQL 版本自动配置
	}), &gorm.Config{})
	// 建表
	// if err == nil {
	// 	db.AutoMigrate(&User{})
	// }

	// if err == nil {
	// 	db.AutoMigrate(&User2{})
	// }

	// 插入数据
	if err == nil {
		user := &User{}
		// sql.NullString,true的时候，有默认空值，而不是null
		user.MemberNumber.Valid = true
		db.Create(user)
	}
}
