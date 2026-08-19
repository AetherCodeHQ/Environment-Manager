package main

import (
	"fmt"
	"os"
)

// environment_manager - Manage dev environments
func environment_manager(path string) {
	fmt.Println("========================================")
	fmt.Println("  Environment-Manager")
	fmt.Println("  Manage dev environments")
	fmt.Println("========================================")
	fmt.Println()
	fmt.Println("Target:", path)
	fmt.Println("Processing...")
	fmt.Println("Done!")
}

func main() {
	path := "."
	if len(os.Args) > 1 {
		path = os.Args[1]
	}
	environment_manager(path)
}
