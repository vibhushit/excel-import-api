package utils

import "log"

// HandleError logs the error message
func HandleError(err error) {
    if err != nil {
        log.Println("Error:", err.Error())
    }
}
