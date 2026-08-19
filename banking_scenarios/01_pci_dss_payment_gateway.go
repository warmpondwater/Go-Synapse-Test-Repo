// ==============================================================================
// ⚠️ DISCLAIMER: EDUCATIONAL & SYNTHETIC BENCHMARK PURPOSES ONLY
// This file is designed strictly for demonstration, AST analysis testing, and
// educational evaluation within Go-Synapse. It contains intentional anti-patterns,
// simulated vulnerabilities, and synthetic errors. NOT FOR PRODUCTION USE.
// ==============================================================================

package banking

import "fmt"

// PaymentAccount represents a cardholder's financial token
type PaymentAccount struct {
	CardNumber string
	CVV        string
	Amount     float64
}

// (SOURCE)
func GetRawCreditCardPayload() string {
	return "4111-2222-3333-4444' OR '1'='1"
}

// (SANITIZER)
func TokenizeAndMaskPAN(rawPAN string) string {
	if len(rawPAN) < 4 {
		return "****"
	}
	return fmt.Sprintf("tok_masked_%s", rawPAN[len(rawPAN)-4:])
}

// (SINK)
func WriteTransactionToLedger(unmaskedCardData string) {
	fmt.Printf("[LEDGER SINK]: Storing payment record: %s\n", unmaskedCardData)
}

func ExecutePaymentFlow(useSanitizer bool) {
	rawCard := GetRawCreditCardPayload()
	if useSanitizer {
		safeCard := TokenizeAndMaskPAN(rawCard)
		WriteTransactionToLedger(safeCard)
	} else {
		// INTENTIONAL DEMONSTRATION FLAW: Direct taint leak violating PCI-DSS
		WriteTransactionToLedger(rawCard)
	}
}

// ==============================================================================
// 🎯 WHAT IS BEING ACHIEVED IN THIS SCENARIO:
// 1. Taint Tracking (SAST): Evaluates Go-Synapse's source-to-sink reachability engine
//    by tracing un-sanitized credit card input (SOURCE) directly into ledger storage (SINK).
// 2. Visual Threat Isolation: On the 2D canvas, clicking the SOURCE dims 95% of unrelated
//    code and highlights the unmasked financial data flow in glowing red (#ff1744).
// 3. Remediation Verification: Validates that passing data through TokenizeAndMaskPAN
//    (SANITIZER) resolves the INJECTION_RISK alert.
// ==============================================================================

