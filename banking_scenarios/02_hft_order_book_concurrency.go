// ==============================================================================
// ⚠️ DISCLAIMER: EDUCATIONAL & SYNTHETIC BENCHMARK PURPOSES ONLY
// This file is designed strictly for demonstration, AST analysis testing, and
// educational evaluation within Go-Synapse. It contains intentional anti-patterns,
// simulated vulnerabilities, and synthetic errors. NOT FOR PRODUCTION USE.
// ==============================================================================

package banking

import (
	"fmt"
	"sync"
)

type TradeOrder struct {
	Symbol string
	Price  float64
	Volume int
}

type OrderMatchingEngine struct {
	OrderStream chan TradeOrder
	mu          sync.Mutex
	ActiveCount int
}

func NewOrderMatchingEngine() *OrderMatchingEngine {
	return &OrderMatchingEngine{
		OrderStream: make(chan TradeOrder), // Unbuffered channel demonstration
		ActiveCount: 0,
	}
}

func (e *OrderMatchingEngine) SubmitOrder(order TradeOrder) {
	e.mu.Lock()
	defer e.mu.Unlock()
	e.ActiveCount++
	fmt.Printf("Submitting trade order: %s volume %d\n", order.Symbol, order.Volume)
}

func (e *OrderMatchingEngine) ProcessMatchingLoop(batches int) {
	for i := 0; i < batches; i++ {
		// INTENTIONAL DEMONSTRATION FLAW: Defer inside tight execution loop
		func() {
			e.mu.Lock()
			defer e.mu.Unlock()
			fmt.Printf("Processing matching batch #%d\n", i)
		}()
	}
}
