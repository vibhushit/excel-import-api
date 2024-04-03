package main

import (
	"github.com/gin-gonic/gin"
	"excel-import-api/controllers"
	"excel-import-api/services"
	"excel-import-api/utils"
	"github.com/jinzhu/gorm"
	_ "github.com/jinzhu/gorm/dialects/mysql"
	"excel-import-api/models"
)

func main() {
	// Initialize Gin router
	router := gin.Default()

	// Initialize MySQL database connection
	db, err := gorm.Open("mysql", "your_mysql_user:your_mysql_password@tcp(172.21.0.2:3306)/your_database_name?charset=utf8&parseTime=True&loc=Local")
	if err != nil {
		// Log error
		utils.LogError("Failed to connect to MySQL: " + err.Error())
		panic("Failed to connect to MySQL")
	}
	defer db.Close()

	// Automatically create the table
	db.AutoMigrate(&models.Employee{})

	// Initialize services
	employeeService := services.NewEmployeeService(db)
	excelService := services.NewExcelService()

	// Initialize controllers
	employeeController := controllers.NewEmployeeController(employeeService)

	// Routes
	v1 := router.Group("/api/v1")
	{
		v1.GET("/employees", employeeController.GetEmployees)
		v1.POST("/employees", employeeController.AddEmployee)
		v1.PUT("/employees/:id", employeeController.UpdateEmployee)
		v1.DELETE("/employees/:id", employeeController.DeleteEmployee)
	}

	// Define the path to the Excel file
	filepath := "Sample_Employee_data_xlsx.xlsx"

	// Parse Excel data
	employees, err := excelService.ParseExcelData(filepath)
	if err != nil {
		// Log error
		utils.LogError("Failed to parse Excel data: " + err.Error())
	} else {
		// Log successful parsing
		utils.LogInfo("Successfully parsed Excel data")

		// Add employees to the database
		for _, emp := range employees {
			err := employeeService.AddEmployee(&emp)
			if err != nil {
				// Log error adding employee to the database
				utils.LogError("Failed to add employee to database: " + err.Error())
			} else {
				// Log successful addition of employee
				utils.LogInfo("Added employee to database: " + emp.FirstName + " " + emp.LastName)
			}
		}
	}

	// Run the server
	if err := router.Run(":8080"); err != nil {
		// Log error starting server
		utils.LogError("Failed to start server: " + err.Error())
	}
}
