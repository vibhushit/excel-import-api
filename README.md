# Excel Import API

This project is a Go application that imports data from an Excel file, stores it in a MySQL database, and caches the data in Redis. It provides API endpoints using the Gin framework for viewing and managing the imported data.

## Overview

- **Import Excel Data**: Reads data from an Excel file and structures it appropriately.
- **Store Data**: Connects to MySQL to store the imported data and caches it in Redis.
- **API Endpoints**: Provides endpoints to view imported data, edit records, and update the database and cache.

## Technologies Used

- Go
- Gin (for API endpoints)
- MySQL (Dockerized)
- Redis (Dockerized)

## Installation

1. Clone the repository.
2. Navigate to the project directory:
3. run the command `go mod tidy`
4. run command `docker-compose up -d`
5. Modify the MySQL connection string in the `main.go` file with the IP address of the MySQL container:

// Initialize MySQL database connection
db, err := gorm.Open("mysql", "your_mysql_user:your_mysql_password@tcp(mysql-container_ip:3306)/your_database_name?charset=utf8&parseTime=True&loc=Local")

// Replace mysql-container_ip with the IP address of the MySQL container.
// You can fetch the IP address using the following command:
// docker inspect -f '{{range .NetworkSettings.Networks}}{{.IPAddress}}{{end}}' mysql-container


6. run the application `go run main.go`
7. once you stop the application do `docker-compose down`

# Usage

## Import Excel Data:

1. Place your Excel file with data to import in the project directory.
2. The application will automatically read and parse the data from the Excel file, storing it into the MySQL database and caching it in Redis.

## View Imported List:

- Access the API endpoints to view the imported data:
  - **GET** `/api/v1/employees`: Retrieves the list of imported employees.

## Edit Record:

- Use the provided API endpoints to edit a specific record:
  - **PUT** `/api/v1/employees/:id`: Updates the record with the specified ID in both MySQL database and Redis cache.

## Clear Cache 

- Use the provided API endpoint to clear the cache :
  - **POST** `/api/v1/clear-cache`: Clears the cache in Redis.

# API Endpoints

- **GET** `/api/v1/employees`: Get all employees.
- **POST** `/api/v1/employees`: Add a new employee.
- **PUT** `/api/v1/employees/:id`: Update an employee.
- **DELETE** `/api/v1/employees/:id`: Delete an employee.
- **POST** `/api/v1/clear-cache`: Clear cache and update database from cache.
- **POST** `/api/v1/update-cache-to-mysql`: Update MySQL database from cache.

# Explanation of Imports

- `github.com/gin-gonic/gin`: Gin web framework for building HTTP servers.
- `excel-import-api/controllers`: Custom controllers for handling HTTP requests.
- `excel-import-api/services`: Custom services for business logic.
- `excel-import-api/utils`: Utility functions for logging and error handling.
- `github.com/jinzhu/gorm`: GORM is an ORM library for Golang, used for database operations.
- `github.com/jinzhu/gorm/dialects/mysql`: MySQL dialect for GORM.
- `github.com/xuri/excelize/v2`: Excel file parser library for Go.
- `github.com/go-redis/redis/v8`: Redis client for Go.
- `github.com/olekukonko/tablewriter`: Library for generating ASCII tables.

## Contributors

- [Vibhushit Tyagi](https://github.com/vibhushit)












