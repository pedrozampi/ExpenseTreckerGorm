package gormteste

import (
	"errors"
	"fmt"

	"gorm.io/driver/sqlite"
	"gorm.io/gorm"
)

type Expense struct {
	gorm.Model
	Value float64
	Type  int `gorm:"default:0"`
}

const (
	Groceries  = 1
	Leisure    = 2
	Eletronics = 3
	Utilities  = 4
	Clothing   = 5
	Health     = 6
	Others     = 0
)

func (e Expense) StringType() string {
	return [...]string{
		"Groceries",
		"Leisure",
		"Electronics",
		"Utilities",
		"Clothing",
		"Health",
		"Others",
	}[e.Type-1]
}

func (e Expense) String() string {
	return fmt.Sprintf("ID: %d, Value: %.2f, Type: %s", e.ID, e.Value, e.StringType())
}

func GetDB() (db *gorm.DB) {
	db, err := gorm.Open(sqlite.Open("test.db"), &gorm.Config{})
	if err != nil {
		panic("failed to connect database")
	}
	db.AutoMigrate(&Expense{})
	return
}

func InsertDB(exp Expense) Expense {
	db := GetDB()
	db.Create(&exp)
	return exp
}

func FindAll() (exps []Expense) {
	db := GetDB()
	db.Find(&exps)

	return
}

func GetbyID(id int) (exp Expense, err error) {
	db := GetDB()
	db.First(&exp, id)
	if exp.ID == 0 {
		err = errors.New("not found")
	}
	return
}

func GetByType(tp int) (exp []Expense) {
	db := GetDB()
	db.Find(&exp, "type  = ?", tp)
	return
}

func UpdateColumn(exp Expense, column string, value interface{}) Expense {
	db := GetDB()
	db.Model(&exp).Update(column, value)

	return exp
}

func Update(exp Expense, nExp Expense) Expense {
	db := GetDB()
	db.Model(&exp).Updates(nExp)

	return exp
}

func Delete(exp Expense, id int) {
	db := GetDB()
	db.Delete(&exp, id)
}
