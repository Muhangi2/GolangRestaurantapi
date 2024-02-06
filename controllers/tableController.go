package controllers

import (
	"context"
	"time"

	"github.com/gin-gonic/gin"
)


func GetTables() gin.HandlerFunc {
	return func(c *gin.Context) {
		var ctx, cancel = context.WithTimeout(context.Background(), 100*time.Second)
     
	}
}
func GetTable() gin.HandlerFunc {
	return func(c *gin.Context) {

	}
}
func CreateTable() gin.HandlerFunc {
	return func(c *gin.Context) {

	}
}
func UpdateTable() gin.HandlerFunc {
	return func(c *gin.Context) {

	}
}
func DeleteTable() gin.HandlerFunc {
	return func(c *gin.Context) {

	}
}
