package main

import (
	"fmt"
	"gorm.io/driver/mysql"
	//"gorm.io/driver/sqlite"
	"gorm.io/gorm"
)

/**
gorm官网：https://gorm.io/docs/
*/

type Product struct {
	gorm.Model
	Code  string
	Price uint
}

func main() {
	fmt.Println("aaaaaa")
	/**
	使用sqlite
	https://gorm.io/docs/
	*/
	//db, err := gorm.Open(sqlite.Open("test.db"), &gorm.Config{
	//	//DryRun: true, // DryRun generate sql without execute
	//})

	/**
	使用mysql
	https://gorm.io/docs/connecting_to_the_database.html#Customize-Driver
	https://github.com/go-sql-driver/mysql#dsn-data-source-name
	*/
	dsn := "root:anivia@930511@tcp(localhost:3306)/geektime-basic-go"
	db, err := gorm.Open(mysql.Open(dsn), &gorm.Config{})
	if err != nil {
		panic("failed to connect database")
	}

	fmt.Println(db)
	db = db.Debug()

	// Migrate the schema
	// 建表
	db.AutoMigrate(&Product{})

	// Create 添加行数据
	db.Create(&Product{Code: "D440", Price: 100})
	db.Create(&Product{Code: "D441", Price: 200})
	db.Create(&Product{Code: "D442", Price: 300})

	// Read 查询
	var product Product
	db.First(&product, 3)                  // find product with integer primary key
	db.First(&product, "code = ?", "D442") // find product with code D42
	fmt.Println("product:", product)

	// Update - update product's price to 200 将 product 的 price 更新为 200
	db.Model(&product).Update("Price", 2000)
	// Update - update multiple fields 更新多个字段
	// 这一句会更新 Price 和 Code 两个字段
	// SET `price`=200 `code`="F42"
	db.Model(&product).Updates(Product{Price: 3000, Code: "F42"}) // non-zero fields
	db.Model(&product).Updates(map[string]interface{}{"Price": 4000, "Code": "F43"})

	// Delete - delete product
	db.Delete(&product, 3)
}
