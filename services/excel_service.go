package services

import (
	"fmt"
    "github.com/xuri/excelize/v2"
    "excel-import-api/models"
)

// ExcelService handles parsing data from Excel files
type ExcelService interface {
    ParseExcelData(filepath string) ([]models.Employee, error)
}

// excelService implements ExcelService
type excelService struct{}

// NewExcelService creates a new instance of ExcelService
func NewExcelService() ExcelService {
    return &excelService{}
}

// ParseExcelData parses data from the specified Excel file
func (es *excelService) ParseExcelData(filepath string) ([]models.Employee, error) {
    var employees []models.Employee

    // Open the Excel file
    f, err := excelize.OpenFile(filepath)
    if err != nil {
        return nil, err
    }

    // Get all rows from the uk-500 sheet
    rows, err := f.GetRows("uk-500")
    if err != nil {
        return nil, err
    }

    // Print the content of the Excel file for testing
    for _, row := range rows {
        fmt.Println(row)
    }

// Iterate over rows, skipping the header row
    for _, row := range rows[1:] {
        employee := models.Employee{
            FirstName:    row[0],
            LastName:     row[1],
            CompanyName:  row[2],
            Address:      row[3], // Add Address field
            City:         row[4], // Adjust index accordingly if Address is added
            Country:      row[5],
            Postal:       row[6],
            Phone:        row[7],
            Email:        row[8],
            Web:          row[9], // Adjust index accordingly if Address is added
        }
        employees = append(employees, employee)
    }

    return employees, nil
}
