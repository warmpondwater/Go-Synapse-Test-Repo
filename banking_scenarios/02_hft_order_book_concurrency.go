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

// ==============================================================================
// 🎯 WHAT IS BEING ACHIEVED IN THIS SCENARIO:
// 1. Concurrency Analysis: Tests Go-Synapse's heuristic detection of channel
//    synchronization boundaries, goroutines, and mutex lock patterns.
// 2. Anti-Pattern Detection: Evaluates AST warning triggers for loop-scoped defers that
//    cause lock starvation and memory accumulation in high-frequency trading loops.
// 3. Visual Graph Signals: Renders color-coded concurrency and sync barrier edges
//    on the 2D spatial canvas.
// ==============================================================================

