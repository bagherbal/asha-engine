// Package asha is the standalone runtime board for the ASHA engine.
//
// It is deliberately not a theorem-gate package.  It is a compact, runtime-safe
// API for the final post-Gate-535 result: native finite geometry, bridge
// coefficients, quarantined family axioms, environmental/PACS metadata, and
// scenario reports.  Earlier theorem packages remain as audit history; this
// package is the endpoint that a user or CI job can rely on for deterministic
// calculation and reporting.
package asha
