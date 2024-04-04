package services

import (
	"fmt"
	"excel-import-api/models"
	"github.com/jinzhu/gorm"
)

// MySQLService handles MySQL database operations
type MySQLService struct {
	DB *gorm.DB
}

// NewMySQLService creates a new instance of MySQLService
func NewMySQLService() (*MySQLService, error) {
	// MySQL connection string
	dbURI := "your_mysql_user:your_mysql_password@tcp(172.19.0.2:3306)/your_database_name?charset=utf8&parseTime=True&loc=Local"

	// Connect to MySQL database
	db, err := gorm.Open("mysql", dbURI)
	if err != nil {
		return nil, err
	}

	// Disable table name's pluralization
	db.SingularTable(true)

	// Automigrate the models
	db.AutoMigrate(&models.Employee{}) // Adjust model name as per your project

	fmt.Println("Connected to MySQL database")

	return &MySQLService{DB: db}, nil
}

// Close closes the MySQL connection
func (ms *MySQLService) Close() error {
	return ms.DB.Close()
}

// EmployeeService handles operations related to employees in the database
type EmployeeService interface {
	GetEmployees() ([]models.Employee, error)
	AddEmployee(employee *models.Employee) error
	UpdateEmployee(id string, employee *models.Employee) error
	DeleteEmployee(id string) error
	//GetEmployeeByID(id uint) (*models.Employee, error)
	//GetEmployeeByEmail(email string) (*models.Employee, error)
	//GetEmployeeByPhoneAndEmail(phone, email string) (*models.Employee, error)
}

// employeeService implements EmployeeService
type employeeService struct {
	db *gorm.DB
}

// NewEmployeeService creates a new instance of EmployeeService
func NewEmployeeService(db *gorm.DB) EmployeeService {
	return &employeeService{db: db}
}

// GetEmployees retrieves all employees from the database
func (es *employeeService) GetEmployees() ([]models.Employee, error) {
	var employees []models.Employee
	if err := es.db.Find(&employees).Error; err != nil {
		return nil, err
	}
	return employees, nil
}

// AddEmployee adds a new employee to the database
func (es *employeeService) AddEmployee(employee *models.Employee) error {
    // Proceed to add the employee to the database
    if err := es.db.Create(employee).Error; err != nil {
        return err // Return error if unable to add employee to the database
    }
    return nil
}

// UpdateEmployee updates an existing employee in the database
func (es *employeeService) UpdateEmployee(id string, employee *models.Employee) error {
	if err := es.db.Where("id = ?", id).Save(employee).Error; err != nil {
		return err
	}
	return nil
}

// DeleteEmployee deletes an employee from the database
func (es *employeeService) DeleteEmployee(id string) error {
	if err := es.db.Where("id = ?", id).Delete(&models.Employee{}).Error; err != nil {
		return err
	}
	return nil
}
