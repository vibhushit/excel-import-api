package controllers

import (
	"github.com/gin-gonic/gin"
	"excel-import-api/models"
	"excel-import-api/services"
	"net/http"
   "github.com/olekukonko/tablewriter"
    "fmt"
   "context"
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
        "Company Name", 
        "Address",      
        "City",
        "Country",
        "Postal",
        "Phone",
        "Email",
        "Web",
    }
    table.SetHeader(headers)

    // Add data rows to the table
    for _, emp := range employees {
        rowData := []string{
            fmt.Sprintf("%d", emp.ID),
            emp.FirstName,
            emp.LastName,
            emp.CompanyName, 
            emp.Address,     
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

// AddEmployee handles POST request to add a new employee
func (ec *EmployeeController) AddEmployee(c *gin.Context) {
    var employee models.Employee
    if err := c.ShouldBindJSON(&employee); err != nil {
        c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
        return
    }

    // Add employee to MySQL database
    if err := ec.employeeService.AddEmployee(&employee); err != nil {
        c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
        return
    }

    // Cache employee data in Redis
    if err := ec.redisService.CacheEmployee(employee); err != nil {
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

    // Update employee in MySQL database
    if err := ec.employeeService.UpdateEmployee(id, &employee); err != nil {
        c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
        return
    }

    // Cache updated employee data in Redis
    if err := ec.redisService.CacheEmployee(employee); err != nil {
        c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
        return
    }

    c.JSON(http.StatusOK, gin.H{"message": "Employee updated successfully"})
}

// DeleteEmployee handles DELETE request to delete an employee
func (ec *EmployeeController) DeleteEmployee(c *gin.Context) {
    id := c.Param("id")

    // Delete employee from MySQL database
    if err := ec.employeeService.DeleteEmployee(id); err != nil {
        c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
        return
    }

    // Remove employee data from Redis cache
    ctx := context.Background()
    if err := ec.redisService.RemoveEmployeeFromCache(ctx, id); err != nil {
        c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
        return
    }

    c.JSON(http.StatusOK, gin.H{"message": "Employee deleted successfully"})
}

// ClearCache handles HTTP requests to clear the cache
func (ec *EmployeeController) ClearCache(c *gin.Context) {
    // Clear the cache
    ctx := context.Background()
    if err := ec.redisService.ClearCache(ctx); err != nil {
        c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
        return
    }

    c.JSON(http.StatusOK, gin.H{"message": "Cache cleared successfully"})
}


// UpdateCacheToRedis updates the Redis cache with data from MySQL database
func (ec *EmployeeController) UpdateCacheToRedis(c *gin.Context) {
    // Get all employee data from MySQL database
    dbEmployees, err := ec.employeeService.GetEmployees()
    if err != nil {
        c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
        return
    }

    // Cache employees data in Redis
    if err := ec.redisService.CacheEmployees(dbEmployees); err != nil {
        c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
        return
    }

    c.JSON(http.StatusOK, gin.H{"message": "Redis cache updated with MySQL data successfully"})
}



