package models

// Employee represents the structure of employee data
type Employee struct {
    ID            uint   `gorm:"primaryKey"`
    FirstName     string `gorm:"column:first_name"`
    LastName      string `gorm:"column:last_name"`
    CompanyName   string `gorm:"column:company_name"`
    Address       string `gorm:"column:address"`
    City          string `gorm:"column:city"`
    Country       string `gorm:"column:country"`
    Postal        string `gorm:"column:postal"`
    Phone         string `gorm:"column:phone"`
    Email         string `gorm:"column:email"`
    Web           string `gorm:"column:web"`
}


// TableName returns the name of the database table associated with the Employee model
func (Employee) TableName() string {
    return "employees"
}
