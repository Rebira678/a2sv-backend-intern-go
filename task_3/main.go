package main

import (
	"library_management/controllers"
	"library_management/services"
)

func main() {
	// 1. Create the library service (The Brain)
	libraryService := services.NewLibrary()

	// 2. Pass the service to the controller (The Front Desk)
	controller := controllers.LibraryController{
		Service: libraryService,
	}

	// 3. Start the application
	controller.Run()
}
