// ==============================================================================
// ⚠️ DISCLAIMER: EDUCATIONAL & SYNTHETIC BENCHMARK PURPOSES ONLY
// This file is designed strictly for demonstration, AST analysis testing, and
// educational evaluation within Go-Synapse. It contains intentional anti-patterns,
// simulated vulnerabilities, and synthetic errors. NOT FOR PRODUCTION USE.
// ==============================================================================

package banking

import "fmt"

// (DEAD CODE)
func Legacy1998SwiftMT103Parser() {
	var isolatedBuffer string = "MT103_RAW_LEGACY_BUFFER"
	fmt.Println("Abandoned 1998 SWIFT MT103 parser", isolatedBuffer)
}

// (DEAD CODE)
func DeprecatedFedWireClearingProtocol() {
	var oldProtocolCode int = 9012
	fmt.Println("Abandoned FedWire clearing protocol", oldProtocolCode)
}

// (DEAD CODE)
type ObsoleteSettlementTable struct {
	RoutingTransitNumber int
	LegacyAccountID      string
}

// ActiveModernSettlementRouter represents modern ISO 20022 router
type ActiveModernSettlementRouter struct {
	MessageID string
}

func (r *ActiveModernSettlementRouter) RouteISO20022Payment(payload string) bool {
	fmt.Println("Routing modern ISO 20022 payment:", payload)
	return true
}
