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

## Getting Started

To run this project locally:

1. Clone the repository.
2. Set up Docker for MySQL and Redis.
3. Run the Go application.

## Usage

- API endpoints:
  - GET /api/data: Retrieve imported data.
  - POST /api/data: Import data from an Excel file.
  - PUT /api/data/:id: Update a specific record.
  - DELETE /api/data/:id: Delete a specific record.

## Contributors

- [Vibhushit Tyagi](https://github.com/vibhushit)

