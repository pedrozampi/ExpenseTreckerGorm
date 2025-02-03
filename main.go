package main

import (
	"github.com/gin-gonic/gin"
	gormteste "pzampi.gorm/gormteste"
)

func main() {
	router := gin.Default()

	router.GET("/expenses", gormteste.GetExpenses)
	router.GET("/expense/:id", gormteste.GetExpenseByID)
	router.GET("/expense/type/:type", gormteste.GetExpenseByType)

	router.POST("/expense", gormteste.RegisterExpense)

	router.PUT("/expense/:id", gormteste.UpdateExpense)

	router.DELETE("/expense/:id", gormteste.DeleteExpense)

	router.Run("localhost:8080")
}
