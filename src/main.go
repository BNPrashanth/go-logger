package main

import (
	u "go-logger/src/utils"
)

func main() {
	// Strting the Application
	u.GeneralLogger.Println("Starting..")

	// Execution od Code
	u.GeneralLogger.Println("Running Process 1")
	u.GeneralLogger.Println("Running Process 2")

	u.ErrorLogger.Println("An Error Occured: Error abcd")

	u.GeneralLogger.Println("Running Process i")
	u.ErrorLogger.Println("An Error Occured: Error efgh")
	u.GeneralLogger.Println("Running Process n")

	// Ending the Application
	u.GeneralLogger.Println("Execution Completed..")
}
