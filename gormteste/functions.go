package gormteste

import (
	"strconv"

	"github.com/gin-gonic/gin"

	"net/http"
)

func GetExpenses(c *gin.Context) {
	expenses := FindAll()

	c.JSON(http.StatusOK, expenses)
}

func GetExpenseByID(c *gin.Context) {
	id, err := strconv.Atoi(c.Param("id"))
	if err != nil {
		return
	}

	expense, err := GetbyID(id)
	if err != nil {
		c.JSON(http.StatusNotFound, gin.H{"message": "Expense not found"})
		return
	}

	c.JSON(http.StatusFound, expense)
}

func GetExpenseByType(c *gin.Context) {
	typeInt, err := strconv.Atoi(c.Param("type"))
	if err != nil {
		return
	}

	list := GetByType(typeInt)

	c.JSON(http.StatusOK, list)
}

func RegisterExpense(c *gin.Context) {
	var expense Expense

	if err := c.BindJSON(&expense); err != nil {
		return
	}

	created := InsertDB(expense)
	c.JSON(http.StatusCreated, created)
}

func UpdateExpense(c *gin.Context) {
	id, err := strconv.Atoi(c.Param("id"))
	if err != nil {
		return
	}

	var nExpense Expense

	if err := c.BindJSON(&nExpense); err != nil {
	}

	expense, err := GetbyID(id)
	if err != nil {
		c.JSON(http.StatusNotFound, gin.H{"message": "Expense not found"})
		return
	}
	expense = Update(expense, nExpense)

	c.JSON(http.StatusOK, expense)
}

func DeleteExpense(c *gin.Context) {
	id, err := strconv.Atoi(c.Param("id"))
	if err != nil {
		return
	}
	var exp Expense
	Delete(exp, id)

	c.JSON(http.StatusNoContent, gin.H{"message": "Expense deleted."})
}
