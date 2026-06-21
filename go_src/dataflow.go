package main

import "fmt"

// (SOURCE)
func GetUserInput() string {
	return "SELECT * FROM users WHERE id = 1"
}

// (SANITIZER)
func SanitizeInput(input string) string {
	// Escape unsafe characters
	return fmt.Sprintf("escaped(%s)", input)
}

// (SINK)
func ExecuteDatabaseQuery(query string) {
	fmt.Printf("Executing query in database: %s\n", query)
}

func RunDataFlow() {
	raw := GetUserInput()
	clean := SanitizeInput(raw)
	ExecuteDatabaseQuery(clean)
}
