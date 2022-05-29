package main

import (
	"context"
	"fmt"
	"time"

	"gorm.io/driver/mysql"
	"gorm.io/gorm"
	"gorm.io/gorm/logger"
)

type SqlLogger struct {
	logger.Interface
}

func (l SqlLogger) Trace(ctx context.Context, begin time.Time, fc func() (sql string, rowAffected int64), err error) {
	sql, _ := fc()
	fmt.Printf("%v \n=========\n", sql)
}

var db *gorm.DB

func main() {
	dsn := "root:1234@tcp(127.0.0.1:3306)/ong?parseTime=True"
	var err error
	db, err = gorm.Open(mysql.Open(dsn), &gorm.Config{
		Logger: &SqlLogger{},
		// DryRun: true,
	})

	if err != nil {
		panic(err)
	}

	// err = db.AutoMigrate(Customer{})
	if err != nil {
		panic(err)
	}
	// CreateGender("dsyasa")
	// GetGenders()
	// GetGender(1)
	// GetGenderByName("pasit")
	// CreateCustomer("pasit", 5)
	GetCustomer()
}

type Customer struct {
	ID       uint
	Name     string
	Gender   Gender
	GenderID uint
}

func CreateCustomer(name string, genderID uint) {
	customer := Customer{
		Name:     name,
		GenderID: genderID,
	}
	tx := db.Create(&customer)
	if tx.Error != nil {
		fmt.Println(tx.Error)
		return
	}
	fmt.Println(customer)
}

func GetCustomer() {
	customers := []Customer{}
	tx := db.Preload("Gender").Find(&customers)
	if tx.Error != nil {
		fmt.Println(tx.Error)
		return
	}
	fmt.Println(customers)
}

func CreateGender(name string) {
	gender := Gender{Name: name}
	tx := db.Create(&gender)
	if tx.Error != nil {
		fmt.Println(tx.Error)
		return
	}
	fmt.Println(gender)
}

func UpdateGender(id uint, name string) {
	gender := Gender{}
	tx := db.First(&gender, id)
	if tx.Error != nil {
		fmt.Println(tx.Error)
		return
	}
	gender.Name = name
	tx = db.Save(&gender)
	if tx.Error != nil {
		fmt.Println(tx.Error)
		return
	}
}

func UpdateGender2(id uint, name string) {
	gender := Gender{Name: name}
	tx := db.Model(&Gender{}).Where("id=?", id).Updates(gender)
	if tx.Error != nil {
		fmt.Println(tx.Error)
		return
	}
	fmt.Println(gender)
}

func GetGenderByName(name string) {
	gender := Gender{}
	tx := db.Where("name=?", name).Find(&gender)
	if tx.Error != nil {
		fmt.Println(tx.Error)
	}
	fmt.Println(gender)
}

func GetGender(id uint) {
	gender := Gender{}
	tx := db.First(&gender, id)
	if tx.Error != nil {
		fmt.Println(tx.Error)
	}
	fmt.Println(gender)
}

func GetGenders() {
	genders := []Gender{}
	tx := db.Order("id").Find(&genders)
	if tx.Error != nil {
		fmt.Println(tx.Error)
		return
	}
	fmt.Println(genders)
}

type Gender struct {
	ID   uint
	Name string `gorm:"unique;size:10"`
}

// func (t Test) TableName() string {
// 	return "My test"
// }

type Test struct {
	gorm.Model
	Code uint   `gorm:"comment:This is code"`
	Name string `gorm:"column:myname;size:30; not null; default:Hello"`
}
