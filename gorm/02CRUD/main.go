package main

import (
	"database/sql"
	"fmt"
	"time"

	"gorm.io/driver/mysql"
	"gorm.io/gorm"
	"gorm.io/gorm/clause"
)

type User struct {
	gorm.Model
	Name         string         // 可传空字符串
	Email        *string        // 可传空指针
	Birthday     *time.Time     // 可传空指针
	MemberNumber sql.NullString // 可传空字符串
	ActivatedAt  sql.NullTime   // 可传空时间
	Dept         Dept           `gorm:"foreignKey:DeptId"`
	DeptId       uint
}

type Dept struct {
	gorm.Model
	Name string
}

// 创建单条记录
// 创建字符串指针的时候，要将字符串给一个数组，然后取数组第一个元素，然后在取指针
// 将字符串转时间用 time.Parse，因为这个方法有两个返回值是，所以需要先处理下
func createOne(db *gorm.DB) {
	birthday, _ := time.Parse("2006-01-02", "1993-03-23")
	activatedAt, _ := time.Parse("2006-01-02", "2025-07-29")
	user := &User{
		Name:     "张三",
		Email:    &[]string{"123@163.com"}[0],
		Birthday: &birthday,
		MemberNumber: sql.NullString{
			String: "1234567890",
			// 可以传空字符串
			Valid: true,
		},
		ActivatedAt: sql.NullTime{
			Time:  activatedAt,
			Valid: true,
		},
	}
	res := db.Create(user)
	fmt.Println("影响行数：", res.RowsAffected)
}

// 创建多条数据
// 数组对象的指针要放在【】后
func createMany(db *gorm.DB) {
	birthday, _ := time.Parse("2006-01-02", "1993-03-23")
	activatedAt, _ := time.Parse("2006-01-02", "2025-07-29")
	users := []*User{
		{
			Name:     "李四",
			Email:    &[]string{"123@163.com"}[0],
			Birthday: &birthday,
			MemberNumber: sql.NullString{
				String: "1234567890",
				// 可以传空字符串
				Valid: true,
			},
			ActivatedAt: sql.NullTime{
				Time:  activatedAt,
				Valid: true,
			},
		},
		{
			Name:     "王五",
			Email:    &[]string{"123@163.com"}[0],
			Birthday: &birthday,
			MemberNumber: sql.NullString{
				String: "1234567890",
				// 可以传空字符串
				Valid: true,
			},
			ActivatedAt: sql.NullTime{
				Time:  activatedAt,
				Valid: true,
			},
		},
	}
	res := db.Create(users)
	fmt.Println("影响行数：", res.RowsAffected)
}

// 创建/忽略部分字段
func createPart(db *gorm.DB) {
	birthday, _ := time.Parse("2006-01-02", "1993-03-23")
	activatedAt, _ := time.Parse("2006-01-02", "2025-07-29")
	user := &User{
		Name:     "张三",
		Email:    &[]string{"123@163.com"}[0],
		Birthday: &birthday,
		MemberNumber: sql.NullString{
			String: "1234567890",
			// 可以传空字符串
			Valid: true,
		},
		ActivatedAt: sql.NullTime{
			Time:  activatedAt,
			Valid: true,
		},
	}
	// 创建部分字段
	db.Select("Name", "Age").Create(user)
	// 忽略部分字段
	user.ID += 1
	db.Omit("Name", "Age").Create(user)
}

// 批次创建
func batchCreate(db *gorm.DB) {
	users := []User{
		{Name: "一"},
		{Name: "二"},
		{Name: "三"},
	}
	db.CreateInBatches(users, 2)
}

// hook钩子
// 参数是*gorm.DB 类型
func (e *User) BeforeSave(db *gorm.DB) (err error) {
	if e.Name == "张三" {
		e.Name = "马六"
	}
	return err
}

func (e *User) BeforeCreate() {

}

func (e *User) AfterSave() {

}

func (e *User) AfterCreate() {

}

// 跳过hook
func junpHook(db *gorm.DB) {
	user := &User{
		Name: "张三",
	}
	db.Session(&gorm.Session{SkipHooks: true}).Create(user)
}

// 关联创建
// 在结构体中除了创建关联对象外还要创建一个关联字段，并在gorm标签中加上外键字段
func relationCreate(db *gorm.DB) {
	dept := Dept{Name: "开发部"}
	user := User{
		Name:   "关联",
		Dept:   dept,
		DeptId: dept.ID,
	}
	// db.Create(&dept)
	db.Create(&user)
}

// 跳过关联创建
func jumpRelationCreate(db *gorm.DB) {
	dept := Dept{Name: "开发部"}
	user := User{
		Name:   "关联",
		DeptId: dept.ID,
		Dept:   dept,
	}
	// db.Create(&user)
	// db.Omit("Dept").Create(&user)
	// skip all associations
	db.Omit(clause.Associations).Create(&user)

}

// Upsert 及冲突
func conflict(db *gorm.DB) {
	db.Clauses(clause.OnConflict{})
}

// 查询
func testSelect(db *gorm.DB) {
	user := User{}
	// 获取第一条记录（主键升序）
	res := db.First(&user)
	print(res)
	// 获取一条记录，没有指定排序字段
	// 获取最后一条记录（主键降序）
	// 检查 ErrRecordNotFound 错误
}
func main() {
	// 连接数据库：方式1
	dsn := "remote:mysql_remote@tcp(192.168.40.10:3306)/web3?charset=utf8mb4&parseTime=True&loc=Local"
	db, err := gorm.Open(mysql.Open(dsn), &gorm.Config{})

	if err != nil {
		fmt.Println("数据库连接失败,失败信息：", err)
	}
	// 创建表
	// db.AutoMigrate(&User{}, &Dept{})
	// 创建单条记录
	// createOne(db)
	//创建多条数据
	// createMany(db)
	// 创建部分字段
	// createPart(db)
	// 批次创建
	// batchCreate(db)
	// 创建钩子
	// createOne(db)
	// 跳过钩子
	// junpHook(db)
	// 关联创建
	// relationCreate(db)
	// 跳过关联创建
	// jumpRelationCreate(db)
	// Upsert 及冲突
	// conflict(db)
	// 查询
	testSelect(db)
}
