package utils

import "log"

// LogInfo logs an informational message
func LogInfo(message string) {
    log.Println("Info:", message)
}

// LogError logs an error message
func LogError(message string) {
    log.Println("Error:", message)
}
