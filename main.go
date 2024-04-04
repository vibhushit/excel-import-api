package main

import (
	"github.com/gin-gonic/gin"
	"excel-import-api/controllers"
	"excel-import-api/services"
	"excel-import-api/utils"
	"github.com/jinzhu/gorm"
	_ "github.com/jinzhu/gorm/dialects/mysql"
	"excel-import-api/models"

	//"github.com/go-redis/redis/v8" 
	// "context" 
	// "strconv"
	// "encoding/json"
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
	//db.AutoMigrate(&models.Employee{})

	// // Initialize Redis client
    // rdb := redis.NewClient(&redis.Options{
    //     Addr: "localhost:6379",
    //     Password: "", // no password set
    //     DB: 0,  // use default DB
    // })

	// // Initialize services
	// employeeService := services.NewEmployeeService(db)
	// excelService := services.NewExcelService()

	// // Initialize controllers
	// employeeController := controllers.NewEmployeeController(employeeService)

	 // Initialize Redis service
	 redisService := services.NewRedisService()

	 // Automatically create the table
	 db.AutoMigrate(&models.Employee{})
 
	 // Initialize services
	 employeeService := services.NewEmployeeService(db)
	 excelService := services.NewExcelService()
 
	 // Initialize controllers
	 employeeController := controllers.NewEmployeeController(employeeService, redisService)

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

        // Add employees to the database if they don't already exist
        for _, emp := range employees {
            existingEmployee, err := employeeService.GetEmployeeByEmail(emp.Email)
            if err != nil {
                // Log error checking if employee exists
                utils.LogError("Failed to check if employee exists: " + err.Error())
                continue // Skip this employee and proceed with the next one
            }

            if existingEmployee != nil {
                // Employee already exists in the database, skip adding
                utils.LogInfo("Employee already exists in the database: " + emp.FirstName + " " + emp.LastName)
                continue // Skip this employee and proceed with the next one
            }

            // Employee does not exist in the database, add it
            if err := employeeService.AddEmployee(&emp); err != nil {
                // Log error adding employee to the database
                utils.LogError("Failed to add employee to database: " + err.Error())
            } else {
                // Log successful addition of employee
                utils.LogInfo("Added employee to database: " + emp.FirstName + " " + emp.LastName)

                // Cache employee data in Redis
                if err := redisService.CacheEmployee(emp); err != nil {
                    // Log error caching employee data
                    utils.LogError("Failed to cache employee data in Redis: " + err.Error())
                } else {
                    // Log successful caching of employee data
                    utils.LogInfo("Cached employee data in Redis: " + emp.FirstName + " " + emp.LastName)
                }
            }
        }
    }



	// // Parse Excel data
	// employees, err := excelService.ParseExcelData(filepath)
	// if err != nil {
	// 	// Log error
	// 	utils.LogError("Failed to parse Excel data: " + err.Error())
	// } else {
	// 	// Log successful parsing
	// 	utils.LogInfo("Successfully parsed Excel data")

	// 	// Add employees to the database
	// 	for _, emp := range employees {
	// 		err := employeeService.AddEmployee(&emp)
	// 		if err != nil {
	// 			// Log error adding employee to the database
	// 			utils.LogError("Failed to add employee to database: " + err.Error())
	// 		} else {
	// 			// Log successful addition of employee
	// 			utils.LogInfo("Added employee to database: " + emp.FirstName + " " + emp.LastName)

	// 			 // Serialize employee struct to JSON
	// 			 empJSON, err := json.Marshal(emp)
	// 			 if err != nil {
	// 				 // Log error marshaling employee to JSON
	// 				 utils.LogError("Failed to marshal employee to JSON: " + err.Error())
	// 			 } else {
	// 				 // Cache employee data in Redis
	// 				 ctx := context.Background() // Create context
	// 				 err := rdb.Set(ctx, "employee:"+strconv.Itoa(int(emp.ID)), empJSON, 0).Err()
	// 				 if err != nil {
	// 					 // Log error caching employee data
	// 					 utils.LogError("Failed to cache employee data in Redis: " + err.Error())
	// 				 } else {
	// 					 // Log successful caching of employee data
	// 					 utils.LogInfo("Cached employee data in Redis: " + emp.FirstName + " " + emp.LastName)
	// 				 }
	// 			 }
	// 		}
	// 	}
	// }

	// Run the server
	if err := router.Run(":8080"); err != nil {
		// Log error starting server
		utils.LogError("Failed to start server: " + err.Error())
	}
}
