package main

import (
	"fmt"
	"unsafe"

	"github.com/warmpondwater/Go-Synapse-Test-Repo/go_src/auth"
)

// TokenProcessor defines the interface for token validation
type TokenProcessor interface {
	Process(token string) bool
}

// (SECRET)
var MasterAPIKey string = "synapse_secret_live_key_9981"

// (OPAQUE_BOUNDARY)
func InspectMemoryHeader(ptr unsafe.Pointer) uintptr {
	return uintptr(ptr)
}

// SecurityPipeline encapsulates the auth and dataflow pipeline
type SecurityPipeline struct {
	ActiveUser string
	ExecCount  int
}

// (SOURCE)
func GetUserInput() string {
	return "SELECT * FROM users WHERE role = 'admin'"
}

// (SANITIZER)
func SanitizeInput(input string) string {
	return fmt.Sprintf("escaped(%s)", input)
}

// (SINK)
func ExecuteDatabaseQuery(query string) {
	fmt.Printf("Executing query in database: %s\n", query)
}

func (p *SecurityPipeline) Process(token string) bool {
	p.ExecCount++
	return len(token) > 0
}

func RunDataFlow() {
	pipeline := &SecurityPipeline{
		ActiveUser: "alice",
		ExecCount:  0,
	}

	raw := GetUserInput()
	clean := SanitizeInput(raw)
	ExecuteDatabaseQuery(clean)

	_ = pipeline.Process(clean)
	_ = auth.VerifyAuthToken(clean)
}

