package main

import (
	"fmt"

	"gorm.io/driver/mysql"
	"gorm.io/gorm"
)

type User struct {
	gorm.Model
	Name      string
	Age       int
	Posts     []Post `gorm:"foreignKey:UserId"`
	PostCount int
}

type Post struct {
	gorm.Model
	UserId   int
	Comments []Comment `gorm:"foreignKey:PostId"`
	Content  string
	Status   string
}

type Comment struct {
	gorm.Model
	PostId  int
	Content string
}

func (u *Post) AfterSave(tx *gorm.DB) (err error) {
	var count int64
	tx.Debug().Model(&Post{}).Where("user_id = ?", u.UserId).Count(&count)
	tx.Debug().Model(&User{}).Where("id = ?", u.UserId).Update("post_count", count)
	return
}

func (u *Comment) AfterDelete(tx *gorm.DB) (err error) {
	post := Post{}
	tx.Debug().Model(&Post{}).Preload("Comments").Where("id = ?", u.PostId).Find(&post)
	if len(post.Comments) == 0 {
		post.Status = "无评论"
		tx.Debug().Where("id=?", u.PostId).Updates(&post)
	}
	return
}

func main() {
	dsn := "remote:mysql_remote@tcp(192.168.40.10:3306)/web3?charset=utf8mb4&parseTime=True&loc=Local"
	db, err := gorm.Open(mysql.Open(dsn), &gorm.Config{})
	if err != nil {
		fmt.Println("database connection failed")
		panic(err)
	}

	// create table
	// 	题目1：模型定义
	// 假设你要开发一个博客系统，有以下几个实体： User （用户）、 Post （文章）、 Comment （评论）。
	// 要求 ：
	// 使用Gorm定义 User 、 Post 和 Comment 模型，其中 User 与 Post 是一对多关系（一个用户可以发布多篇文章）， Post 与 Comment 也是一对多关系（一篇文章可以有多个评论）。
	// 编写Go代码，使用Gorm创建这些模型对应的数据库表。
	db.AutoMigrate(&User{}, &Post{}, &Comment{})

	// create simple
	// db.Create(&User{Name: "Tom", Age: 20})
	// db.Create(&Post{UserId: 1, Content: "Hello, World!"})
	// db.Create(&Post{UserId: 1, Content: "Hello, GO!"})
	// db.Create(&Comment{PostId: 1, Content: "Great post1!"})
	// db.Create(&Comment{PostId: 1, Content: "Great post2!"})
	// db.Create(&Comment{PostId: 2, Content: "Great post3!"})

	// 	题目2：关联查询
	// 基于上述博客系统的模型定义。
	// 要求 ：
	// 编写Go代码，使用Gorm查询某个用户发布的所有文章及其对应的评论信息。
	user := User{Name: "John"}
	db.Model(&user).Preload("Posts").Preload("Posts.Comments").Find(&user)
	fmt.Println(user)
	// 编写Go代码，使用Gorm查询评论数量最多的文章信息。
	post := Post{}
	db.Model(&Post{}).Select("posts.*, count(comments.id) as comment_count").Joins("left join comments on comments.post_id = posts.id").Group("posts.id").Order("comment_count desc").Limit(1).Scan(&post)
	fmt.Println(post)

	// 	题目3：钩子函数
	// 继续使用博客系统的模型。
	// 要求 ：
	// 为 Post 模型添加一个钩子函数，在文章创建时自动更新用户的文章数量统计字段。
	// 为 Comment 模型添加一个钩子函数，在评论删除时检查文章的评论数量，如果评论数量为 0，则更新文章的评论状态为 "无评论"。
	// db.Create(&Post{UserId: 2, Content: "Hello, World!"})
	comment := Comment{}
	db.Debug().Where("id = ?", 100).Find(&comment).Delete(&comment)
}
