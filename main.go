package main

import (
	"github.com/gin-gonic/gin"
)

func main() {
	r := gin.Default()
	r.GET("/calculate", HandlerCalculation)
	r.GET("/simulate", HandlerSimulation)

	r.Run(":8080")
}

func HandlerCalculation(c *gin.Context) {
	// 根据入参完成加减乘除
	var query struct {
		A  int    `form:"a"`
		B  int    `form:"b"`
		Op string `form:"op"`
	}
	if err := c.ShouldBindQuery(&query); err != nil {
		c.JSON(400, gin.H{
			"error": err.Error(),
		})
		return
	}

	result := Calculate(query.A, query.B, query.Op)

	c.JSON(200, gin.H{
		"result": result,
	})
}

func HandlerSimulation(c *gin.Context) {
	// 根据入参完成加减乘除
	var query struct {
		Type string `form:"type"`
	}
	if err := c.ShouldBindQuery(&query); err != nil {
		c.JSON(400, gin.H{
			"error": err.Error(),
		})
		return
	}

	SimulationSwitch(query.Type)

	c.JSON(200, gin.H{
		"result": "success",
	})
}
