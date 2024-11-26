package main

import (
	EmployeeController "backend/api/controller/employee"
	"github.com/gin-gonic/gin"
)

func main() {
	router := gin.Default()

	// แก้ไขการเชื่อมโยงฟังก์ชันจาก controller ให้ถูกต้อง
	router.GET("/employee", EmployeeController.GetEmployee)
	router.POST("/employee", EmployeeController.PostEmployee)
	router.PUT("/employee", EmployeeController.PutEmployee)
	router.DELETE("/employee", EmployeeController.DeleteEmployee)
	router.GET("/employee/:id", EmployeeController.GetEmployeeByID)   
	// แก้ไขการเรียกใช้ฟังก์ชัน Run ให้เป็นตัวพิมพ์ใหญ่
	router.Run()
}
