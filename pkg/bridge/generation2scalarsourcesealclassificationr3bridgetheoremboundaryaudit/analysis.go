// Package generation2scalarsourcesealclassificationr3bridgetheoremboundaryaudit implements
// Gate 947: ScalarSourceSeal Classification and R3 Bridge-Theorem Boundary Audit.
//
// Gate 947 follows Gate 946's failed noncircular replacement search. It freezes
// the honest boundary of the R3 trace bridge: the algebraic closure/measure
// surface is test-passed, but native R3 promotion is blocked by the scalar-source
// seal on S_split and by the still-open full-A_F / spontaneous-orientation status.
package generation2scalarsourcesealclassificationr3bridgetheoremboundaryaudit

import (
	"fmt"
	"strings"
)

const (
	AuditID         = "GATE947-GENERATION2SCALARSOURCESEALCLASSIFICATIONR3BRIDGETHEOREMBOUNDARYAUDIT"
	InheritedStatus = "R3_S_SPLIT_SCALAR_SOURCE_SEAL_CONFIRMED"
	Verdict         = "R3_TRACE_BRIDGE_CLASSIFIED_AS_SCALAR_SOURCE_SEALED_BRIDGE_THEOREM_CANDIDATE_NOT_NATIVE"
	Classification  = "R3_SCALAR_SOURCE_SEALED_TRACEBRIDGE_BOUNDARY_CLASSIFIED_NOT_NATIVE"
	ShortStatus     = "R3_TRACEBRIDGE_TEST_PASSED_SCALAR_SOURCE_SEALED"
	NextGateA       = "NEXT_OPTION_GATE948A_FULL_AF_SPONTANEOUS_ORIENTATION_AUDIT"
	NextGateB       = "NEXT_OPTION_GATE948B_R4_PRECONDITION_UNDER_EXPLICIT_SCALAR_SEAL_AUDIT"
)

const (
	Ssplit  = 0.0012924448188162962
	AlphaB  = 0.0003878958469680527
	NEff    = 3.002327375081808
	CYukawa = 0.9992248096922658
	CHiggs  = 1.037220510866514
)

type BoundaryItem struct {
	Name      string
	Status    string
	Evidence  []string
	Firewalls []string
}

type TheoremBoundary struct {
	BridgeTestPassed     bool
	ClosureFactored      bool
	MeasureFactored      bool
	ScalarSourceSealed   bool
	OrientationStillOpen bool
	NativeR3             bool
	OfficialLedgerUpdate bool
	Status               string
}

type Analysis struct {
	AuditID        string
	Inherited      string
	Verdict        string
	Classification string
	ShortStatus    string
	Ssplit         float64
	AlphaB         float64
	NEff           float64
	CYukawa        float64
	CHiggs         float64
	Boundary       TheoremBoundary
	Items          []BoundaryItem
	Supports       []string
	Failures       []string
	Final          string
}

func BuildDefault() (Analysis, error) {
	items := DefaultItems()
	boundary := DefaultBoundary()
	if !boundary.BridgeTestPassed || !boundary.ClosureFactored || !boundary.MeasureFactored {
		return Analysis{}, fmt.Errorf("Gate 947 must inherit test-passed closure/measure bridge surface: %#v", boundary)
	}
	if !boundary.ScalarSourceSealed || boundary.NativeR3 || boundary.OfficialLedgerUpdate {
		return Analysis{}, fmt.Errorf("Gate 947 must classify scalar-sealed non-native boundary: %#v", boundary)
	}
	if len(items) != 6 {
		return Analysis{}, fmt.Errorf("expected six boundary items, got %d", len(items))
	}
	return Analysis{
		AuditID:        AuditID,
		Inherited:      InheritedStatus,
		Verdict:        Verdict,
		Classification: Classification,
		ShortStatus:    ShortStatus,
		Ssplit:         Ssplit,
		AlphaB:         AlphaB,
		NEff:           NEff,
		CYukawa:        CYukawa,
		CHiggs:         CHiggs,
		Boundary:       boundary,
		Items:          items,
		Supports:       Supports(),
		Failures:       Failures(),
		Final:          "Gate 947 freezes the honest boundary: the R3 Z2 alpha/Yukawa trace bridge is pre-test-passed and closure-factored, but it remains a scalar-source-sealed bridge theorem candidate. Native R3, official ledger updates, particles, generation/flavor carriers, and individual Yukawa spectrum claims remain blocked.",
	}, nil
}

func DefaultBoundary() TheoremBoundary {
	return TheoremBoundary{
		BridgeTestPassed:     true,
		ClosureFactored:      true,
		MeasureFactored:      true,
		ScalarSourceSealed:   true,
		OrientationStillOpen: true,
		NativeR3:             false,
		OfficialLedgerUpdate: false,
		Status:               "TEST_PASSED_SCALAR_SOURCE_SEALED_BRIDGE_THEOREM_CANDIDATE",
	}
}

func DefaultItems() []BoundaryItem {
	return []BoundaryItem{
		{
			Name:   "validated closure/measure surface",
			Status: "BRIDGE_TEST_PASSED",
			Evidence: []string{
				"CONDITIONAL_SUPPORT_R3_TRACEBRIDGE_PRETEST_PASSED",
				"CONDITIONAL_SUPPORT_ALPHA_B_RECONSTRUCTED_FROM_CLOSURE_FACTORED_BOUNDARY_MEASURE",
				"CONDITIONAL_SUPPORT_Y_DAGGER_Y_TRACE_ROWS_RECONSTRUCT_OPERATOR_N_EFF",
			},
		},
		{
			Name:   "negative-route protection",
			Status: "FALSE_ROUTES_REJECTED",
			Evidence: []string{
				"CONDITIONAL_SUPPORT_FALSE_ALPHA_ROUTES_REJECTED_BY_NEGATIVE_TESTS",
				"CONDITIONAL_SUPPORT_NO_TESTED_NONCIRCULAR_PROXY_REPRODUCES_S_SPLIT",
			},
		},
		{
			Name:   "S_split scalar source",
			Status: "SEALED_BRIDGE_HISTORY_INPUT",
			Evidence: []string{
				"CONDITIONAL_SUPPORT_S_SPLIT_TRANSPORT_TO_B2_REMAINS_STRONGLY_SUPPORTED",
				"CONDITIONAL_SUPPORT_S_SPLIT_HAS_H72_COMPATIBLE_SCALAR_TYPE",
			},
			Firewalls: []string{
				"FAILED_ROUTE_S_SPLIT_REMAINS_BRIDGE_HISTORY_SCALAR_INPUT",
				"FAILED_ROUTE_NO_NONCIRCULAR_NATIVE_S_SPLIT_PROXY_FOUND",
			},
		},
		{
			Name:   "native R3 promotion",
			Status: "BLOCKED",
			Firewalls: []string{
				"FAILED_ROUTE_NATIVE_R3_NOT_GRANTED",
				"FAILED_ROUTE_CERTIFICATE_II_NOT_PASSED",
			},
		},
		{
			Name:   "full A_F / spontaneous orientation",
			Status: "STILL_OPEN",
			Firewalls: []string{
				"FAILED_ROUTE_FULL_A_F_DESCENT_OR_SPONTANEOUS_ORIENTATION_THEOREM_REQUIRED",
				"FAILED_ROUTE_HIGGS_ORIENTATION_REMAINS_SEALED_AS_ORIENTATION_CLASS",
			},
		},
		{
			Name:   "R4/yukawa spectrum boundary",
			Status: "OUT_OF_SCOPE_BLOCKED",
			Firewalls: []string{
				"FAILED_ROUTE_NO_INDIVIDUAL_YUKAWA_VALUES",
				"FAILED_ROUTE_NO_PHYSICAL_PARTICLE_ASSIGNMENT",
				"FAILED_ROUTE_NO_GENERATION_CARRIER_MAP",
				"FAILED_ROUTE_NO_FLAVOR_ORIENTATION_MAP",
				"FAILED_ROUTE_NO_R4_NATIVE_YUKAWA_SPECTRUM_THEOREM",
			},
		},
	}
}

func Supports() []string {
	return []string{
		"CONDITIONAL_SUPPORT_R3_TRACEBRIDGE_PRETEST_PASSED",
		"CONDITIONAL_SUPPORT_ALPHA_B_RECONSTRUCTED_FROM_CLOSURE_FACTORED_BOUNDARY_MEASURE",
		"CONDITIONAL_SUPPORT_Y_DAGGER_Y_TRACE_ROWS_RECONSTRUCT_OPERATOR_N_EFF",
		"CONDITIONAL_SUPPORT_FALSE_ALPHA_ROUTES_REJECTED_BY_NEGATIVE_TESTS",
		"CONDITIONAL_SUPPORT_S_SPLIT_TRANSPORT_TO_B2_REMAINS_STRONGLY_SUPPORTED",
		"CONDITIONAL_SUPPORT_S_SPLIT_HAS_H72_COMPATIBLE_SCALAR_TYPE",
		"CONDITIONAL_SUPPORT_NO_TESTED_NONCIRCULAR_PROXY_REPRODUCES_S_SPLIT",
		"CONDITIONAL_SUPPORT_R3_TRACEBRIDGE_CAN_BE_CLASSIFIED_AS_SCALAR_SOURCE_SEALED_BRIDGE_THEOREM_CANDIDATE",
	}
}

func Failures() []string {
	return []string{
		"FAILED_ROUTE_S_SPLIT_REMAINS_BRIDGE_HISTORY_SCALAR_INPUT",
		"FAILED_ROUTE_NO_NONCIRCULAR_NATIVE_S_SPLIT_PROXY_FOUND",
		"FAILED_ROUTE_CERTIFICATE_II_NOT_PASSED",
		"FAILED_ROUTE_NATIVE_R3_NOT_GRANTED",
		"FAILED_ROUTE_NO_OFFICIAL_LEDGER_UPDATE_ALLOWED",
		"FAILED_ROUTE_FULL_A_F_DESCENT_OR_SPONTANEOUS_ORIENTATION_THEOREM_REQUIRED",
		"FAILED_ROUTE_HIGGS_ORIENTATION_REMAINS_SEALED_AS_ORIENTATION_CLASS",
		"FAILED_ROUTE_NO_INDIVIDUAL_YUKAWA_VALUES",
		"FAILED_ROUTE_NO_PHYSICAL_PARTICLE_ASSIGNMENT",
		"FAILED_ROUTE_NO_GENERATION_CARRIER_MAP",
		"FAILED_ROUTE_NO_FLAVOR_ORIENTATION_MAP",
		"FAILED_ROUTE_NO_R4_NATIVE_YUKAWA_SPECTRUM_THEOREM",
	}
}

func FormatBoundary(b TheoremBoundary) string {
	return fmt.Sprintf("%s: bridgeTest=%v closure=%v measure=%v scalarSeal=%v orientationOpen=%v nativeR3=%v officialLedger=%v", b.Status, b.BridgeTestPassed, b.ClosureFactored, b.MeasureFactored, b.ScalarSourceSealed, b.OrientationStillOpen, b.NativeR3, b.OfficialLedgerUpdate)
}

func FormatItem(i BoundaryItem) string {
	return fmt.Sprintf("%s: status=%s evidence=[%s] firewalls=[%s]", i.Name, i.Status, strings.Join(i.Evidence, ","), strings.Join(i.Firewalls, ","))
}

func ItemNotes(items []BoundaryItem) []string {
	out := make([]string, 0, len(items))
	for _, i := range items {
		out = append(out, FormatItem(i))
	}
	return out
}

func ItemSupports(items []BoundaryItem) []string {
	var out []string
	for _, i := range items {
		out = append(out, i.Evidence...)
	}
	return out
}

func ItemFailures(items []BoundaryItem) []string {
	var out []string
	for _, i := range items {
		out = append(out, i.Firewalls...)
	}
	return out
}

func containsAll(haystack, needles []string) bool {
	set := map[string]struct{}{}
	for _, h := range haystack {
		set[h] = struct{}{}
	}
	for _, n := range needles {
		if _, ok := set[n]; !ok {
			return false
		}
	}
	return true
}

func appendAll(items ...[]string) []string {
	var out []string
	for _, group := range items {
		out = append(out, group...)
	}
	return out
}

func stringsJoin(items []string) string { return strings.Join(items, " | ") }
