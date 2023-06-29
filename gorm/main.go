package main

import (
	"fmt"
	"log"
	"time"

	"github.com/spf13/pflag"
	"gorm.io/driver/mysql"
	"gorm.io/gorm"
)

type Product struct {
	gorm.Model
	Code  string `gorm:"column:code"`
	Price uint   `gorm:"column:price"`
	//Test  string `gorm:"column:test"`
}

// TableName maps to mysql table name.
func (p *Product) TableName() string {
	return "product"
}

var (
	//iam59!z$
	host     = pflag.StringP("host", "H", "127.0.0.1:3306", "MySQL service host address")
	username = pflag.StringP("username", "u", "root", "Username for access to mysql service")
	password = pflag.StringP("password", "p", "root", "Password for access to mysql, should be used pair with password")
	database = pflag.StringP("database", "d", "test", "Database name to use")
	help     = pflag.BoolP("help", "h", false, "Print this help message")
)

func main() {
	// Parse command line flags
	pflag.CommandLine.SortFlags = false
	pflag.Usage = func() {
		pflag.PrintDefaults()
	}
	pflag.Parse()
	if *help {
		pflag.Usage()
		return
	}

	dsn := fmt.Sprintf(`%s:%s@tcp(%s)/%s?charset=utf8&parseTime=%t&loc=%s`,
		*username,
		*password,
		*host,
		*database,
		true,
		"Local")
	db, err := gorm.Open(mysql.Open(dsn), &gorm.Config{})
	if err != nil {
		panic("failed to connect database")
	}

	sqlDB, err := db.DB()
	if err != nil {
		return
	}
	sqlDB.SetMaxIdleConns(100)          // 设置MySQL的最大空闲连接数（推荐100）
	sqlDB.SetMaxOpenConns(100)          // 设置MySQL的最大连接数（推荐100）
	sqlDB.SetConnMaxLifetime(time.Hour) // 设置MySQL的空闲连接最大存活时间（推荐10s）

	//product := &Product{}
	////db.Unscoped().Where("id = ?", "1").First(product)
	////fmt.Println("product: ", product)
	products := make([]*Product, 0)
	////db.Unscoped().Where("price = ?", "200").Offset(2).Limit(2).Find(&products)
	//db.Unscoped().Where("price = ?", "200").Find(&products)
	//db.Unscoped().Where("code = ?", "D43").Delete(&Product{})
	//for _, product := range products {
	//	fmt.Println("product: ", product)
	//}

	//first := db.First(&product)
	//fmt.Println("product: ", product)
	//fmt.Println("Error: ", first.Error)
	//fmt.Println("RowsAffected: ", first.RowsAffected)
	//fmt.Println("Statement: ", first.Statement)

	//db.Distinct("code", "price").Order("code,price desc").Find(&products)

	var count int64
	db.Distinct("code", "price").Find(&products).Count(&count)
	fmt.Println("count:", count)
	db.Unscoped().Where("price = ?", "200").Find(&products).Count(&count)
	fmt.Println("count:", count)
	for _, product := range products {
		fmt.Println("product: ", product)
	}
	//fmt.Println("count:", count)

	//product.Price = 250
	//product.Code = "D43"
	//db.Save(&product)

	//// 1. Auto migration for given models
	//fmt.Println("11111111111111111111111111111111111111")
	//db.AutoMigrate(&Product{})
	//
	//// 2. Insert the value into database
	//fmt.Println("22222222222222222222222222222")
	//if err := db.Create(&Product{Code: "D42", Price: 100}).Error; err != nil {
	//	log.Fatalf("Create error: %v", err)
	//}
	//PrintProducts(db)

	//// 3. Find first record that match given conditions
	//fmt.Println("33333333333333333333333333333333")
	//product := &Product{}
	//if err := db.Where("code= ?", "D42").First(&product).Error; err != nil {
	//	log.Fatalf("Get product error: %v", err)
	//}
	//
	//// 4. Update value in database, if the value doesn't have primary key, will insert it
	//product.Price = 200
	//fmt.Println("44444444444444444444444444")
	//if err := db.Save(product).Error; err != nil {
	//	log.Fatalf("Update product error: %v", err)
	//}
	//
	//PrintProducts(db)

	//// 5. Delete value match given conditions
	//fmt.Println("55555555555555555555555555555")
	//if err := db.Where("code = ?", "D42").Delete(&Product{}).Error; err != nil {
	//	log.Fatalf("Delete product error: %v", err)
	//}
	//
	//PrintProducts(db)
}

// List products
func PrintProducts(db *gorm.DB) {
	products := make([]*Product, 0)
	var count int64
	d := db.Where("code like ?", "%D%").Offset(0).Limit(5).Order("id desc").Find(&products).Offset(-1).Limit(-1).Count(&count)
	if d.Error != nil {
		log.Fatalf("List products error: %v", d.Error)
	}

	log.Printf("totalcount: %d", count)
	for _, product := range products {
		log.Printf("\tcode: %s, price: %d\n", product.Code, product.Price)
	}
}
