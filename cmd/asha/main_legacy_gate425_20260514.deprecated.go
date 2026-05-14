//go:build legacyasha
// +build legacyasha

// Deprecated after Gate 425 runtime consolidation on 2026-05-14.
// Build with -tags legacyasha to run the old theorem-registry CLI.
package main

import (
	"log"

	"github.com/bagherbal/asha-engine/internal/app"
)

func main() {
	if err := app.Run(); err != nil {
		log.Fatal(err)
	}
}
