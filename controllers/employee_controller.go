package controllers

import (
	"github.com/gin-gonic/gin"
	"excel-import-api/models"
	"excel-import-api/services"
	"net/http"
    "github.com/olekukonko/tablewriter"
    //"encoding/json"
    "fmt"
   //"strconv"
)

// EmployeeController handles HTTP requests related to employees
type EmployeeController struct {
	employeeService services.EmployeeService

    redisService    services.RedisService
}

// NewEmployeeController creates a new instance of EmployeeController
func NewEmployeeController(employeeService services.EmployeeService, redisService services.RedisService) *EmployeeController {
    return &EmployeeController{
        employeeService: employeeService,
        redisService:    redisService,
    }
}

// GetEmployees handles GET request to retrieve all employees
func (ec *EmployeeController) GetEmployees(c *gin.Context) {
    // Check if the data is available in Redis
    employees, err := ec.redisService.GetEmployeesFromCache()
    if err != nil {
        // Data not found in cache, fetch from MySQL database
        employees, err = ec.employeeService.GetEmployees()
        if err != nil {
            c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
            return
        }

        // Cache the fetched data in Redis
        if err := ec.redisService.CacheEmployees(employees); err != nil {
            c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
            return
        }
    }

    // Create a new table
    table := tablewriter.NewWriter(c.Writer)

    // Set table headers
    headers := []string{
        "ID",
        "First Name",
        "Last Name",
        "Company Address",
        "City",
        "Country",
        "Postal",
        "Phone",
        "Email",
        "Web",
    }
    table.SetHeader(headers)

    // Add data rows to the table
    // for _, emp := range employees {
    //     rowData := []string{
    //         strconv.Itoa(emp.ID),
    //         emp.FirstName,
    //         emp.LastName,
    //         emp.CompanyAddress,
    //         emp.City,
    //         emp.Country,
    //         emp.Postal,
    //         emp.Phone,
    //         emp.Email,
    //         emp.Web,
    //     }
    //     table.Append(rowData)
    // }
        // Add data rows to the table
    for _, emp := range employees {
        rowData := []string{
            fmt.Sprintf("%d", emp.ID),
            emp.FirstName,
            emp.LastName,
            emp.CompanyAddress,
            emp.City,
            emp.Country,
            emp.Postal,
            emp.Phone,
            emp.Email,
            emp.Web,
        }
        table.Append(rowData)
    }

    // Render the table
    table.Render()
}


// // GetEmployees handles GET request to retrieve all employees
// func (ec *EmployeeController) GetEmployees(c *gin.Context) {
//     // Check if the data is available in Redis
//     employees, err := ec.redisService.GetEmployeesFromCache()
//     if err != nil {
//         // Data not found in cache, fetch from MySQL database
//         employees, err = ec.employeeService.GetEmployees()
//         if err != nil {
//             c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
//             return
//         }

//         // Cache the fetched data in Redis
//         if err := ec.redisService.CacheEmployees(employees); err != nil {
//             c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
//             return
//         }
//     }

//     c.JSON(http.StatusOK, employees)
// }

// // GetEmployees handles GET request to retrieve all employees
// func (ec *EmployeeController) GetEmployees(c *gin.Context) {
//     // Check if the data is available in Redis
//     cachedData, err := ec.redisService.GetData("employees")
//     if err == nil {
//         // Data found in cache, return it
//         var employees []models.Employee
//         if err := json.Unmarshal(cachedData, &employees); err != nil {
//             c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
//             return
//         }
//         c.JSON(http.StatusOK, employees)
//         return
//     }

//     // Data not found in cache, fetch from MySQL database
//     employees, err := ec.employeeService.GetEmployees()
//     if err != nil {
//         c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
//         return
//     }

//     // Cache the fetched data in Redis
//     empJSON, err := json.Marshal(employees)
//     if err != nil {
//         c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
//         return
//     }
//     err = redisService.CacheData("employees", empJSON)
//     if err != nil {
//         c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
//         return
//     }

//     // Return the fetched data
//     c.JSON(http.StatusOK, employees)
// }


// // GetEmployees handles GET request to retrieve all employees
// func (ec *EmployeeController) GetEmployees(c *gin.Context) {
// 	// employees, err := ec.employeeService.GetEmployees()
// 	// if err != nil {
// 	// 	c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
// 	// 	return
// 	// }
// 	// c.JSON(http.StatusOK, employees)

//     employees, err := ec.employeeService.GetEmployees()
//     if err != nil {
//         c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
//         return
//     }

//     // Create a new table
//     table := tablewriter.NewWriter(c.Writer)

//     // Set table headers
//     headers := []string{
//         "ID",
//         "First Name",
//         "Last Name",
//         "Company Address",
//         "City",
//         "Country",
//         "Postal",
//         "Phone",
//         "Email",
//         "Web",
//     }
//     table.SetHeader(headers)

//     // Add data rows to the table
//     for _, emp := range employees {
//         rowData := []string{
//             fmt.Sprintf("%d", emp.ID),
//             emp.FirstName,
//             emp.LastName,
//             emp.CompanyAddress,
//             emp.City,
//             emp.Country,
//             emp.Postal,
//             emp.Phone,
//             emp.Email,
//             emp.Web,
//         }
//         table.Append(rowData)
//     }

//     // Render the table
//     table.Render()
// }

// AddEmployee handles POST request to add a new employee
func (ec *EmployeeController) AddEmployee(c *gin.Context) {
	var employee models.Employee
	if err := c.ShouldBindJSON(&employee); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}
	if err := ec.employeeService.AddEmployee(&employee); err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}
	c.JSON(http.StatusCreated, gin.H{"message": "Employee added successfully"})
}

// UpdateEmployee handles PUT request to update an existing employee
func (ec *EmployeeController) UpdateEmployee(c *gin.Context) {
	id := c.Param("id")
	var employee models.Employee
	if err := c.ShouldBindJSON(&employee); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}
	if err := ec.employeeService.UpdateEmployee(id, &employee); err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}
	c.JSON(http.StatusOK, gin.H{"message": "Employee updated successfully"})
}

// DeleteEmployee handles DELETE request to delete an employee
func (ec *EmployeeController) DeleteEmployee(c *gin.Context) {
	id := c.Param("id")
	if err := ec.employeeService.DeleteEmployee(id); err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}
	c.JSON(http.StatusOK, gin.H{"message": "Employee deleted successfully"})
}
