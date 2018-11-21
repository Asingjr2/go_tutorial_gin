// Logic for specific route urls seperated out from main.
package main

// Relies on router to list specific route control logic.
func initializeRoutes() {
	router.GET("/", showIndexPage)
}
