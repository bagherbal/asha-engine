// Package generation2fullafdescentspontaneousorientationsealaudit implements
// Gate 948: Full A_F Descent vs SpontaneousOrientation Seal Audit.
//
// Gate 948 follows Gate 947's scalar-source-sealed R3 tracebridge boundary.
// It audits the second remaining wall: whether the tracebridge descends to the
// full unbroken finite algebra A_F, or whether it is only lawful in the
// post-orientation stabilizer layer A_F^orient under a spontaneous-orientation
// seal.
package generation2fullafdescentspontaneousorientationsealaudit

import (
	"fmt"
	"strings"
)

const (
	AuditID         = "GATE948-GENERATION2FULLAFDESCENTSPONTANEOUSORIENTATIONSEALAUDIT"
	InheritedStatus = "R3_TRACEBRIDGE_TEST_PASSED_SCALAR_SOURCE_SEALED"
	Verdict         = "FULL_AF_DESCENT_BLOCKED_BUT_TRACEBRIDGE_IS_LAWFUL_IN_POST_ORIENTATION_STABILIZER_LAYER_UNDER_SPONTANEOUS_ORIENTATION_SEAL"
	Classification  = "R3_TRACEBRIDGE_SCALAR_SOURCE_SEALED_AND_POST_ORIENTATION_SEALED_NOT_NATIVE"
	ShortStatus     = "R3_TRACEBRIDGE_TEST_PASSED_DUAL_SEALED_NOT_NATIVE"
	NextGate        = "NEXT_GATE949_R3_DUALSEAL_BRIDGE_THEOREM_FINAL_BOUNDARY_AND_R4_PRECONDITION_AUDIT"
)

const (
	Ssplit  = 0.0012924448188162962
	AlphaB  = 0.0003878958469680527
	NEff    = 3.002327375081808
	CYukawa = 0.9992248096922658
	CHiggs  = 1.037220510866514
)

type OrientationBoundary struct {
	ScalarSourceSealed     bool
	FullAFDescentBlocked   bool
	StableInAFOrient       bool
	SpontaneousOrientSeal  bool
	NativeOrientation      bool
	NativeR3               bool
	OfficialLedgerUpdate   bool
	PhysicalAssignment     bool
	GenerationFlavorClaims bool
	Status                 string
}

type AuditItem struct {
	Name      string
	Status    string
	Supports  []string
	Firewalls []string
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
	Boundary       OrientationBoundary
	Items          []AuditItem
	Supports       []string
	Failures       []string
	Final          string
}

func BuildDefault() (Analysis, error) {
	boundary := DefaultBoundary()
	items := DefaultItems()
	if !boundary.ScalarSourceSealed || !boundary.FullAFDescentBlocked || !boundary.StableInAFOrient || !boundary.SpontaneousOrientSeal {
		return Analysis{}, fmt.Errorf("Gate 948 requires scalar-sealed, full-AF-blocked, post-orientation-stable boundary: %#v", boundary)
	}
	if boundary.NativeOrientation || boundary.NativeR3 || boundary.OfficialLedgerUpdate || boundary.PhysicalAssignment || boundary.GenerationFlavorClaims {
		return Analysis{}, fmt.Errorf("Gate 948 must not grant native orientation/R3 or downstream physical claims: %#v", boundary)
	}
	if len(items) != 6 {
		return Analysis{}, fmt.Errorf("expected six audit items, got %d", len(items))
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
		Final:          "Gate 948 classifies the R3 tracebridge as test-passed, scalar-source sealed, and lawful only inside the post-orientation stabilizer layer under a spontaneous-orientation seal. Full A_F descent remains blocked, the spontaneous-orientation theorem is not native-certified, and native R3 is not granted.",
	}, nil
}

func DefaultBoundary() OrientationBoundary {
	return OrientationBoundary{
		ScalarSourceSealed:     true,
		FullAFDescentBlocked:   true,
		StableInAFOrient:       true,
		SpontaneousOrientSeal:  true,
		NativeOrientation:      false,
		NativeR3:               false,
		OfficialLedgerUpdate:   false,
		PhysicalAssignment:     false,
		GenerationFlavorClaims: false,
		Status:                 "TEST_PASSED_SCALAR_SOURCE_SEALED_POST_ORIENTATION_SEALED_TRACEBRIDGE",
	}
}

func DefaultItems() []AuditItem {
	return []AuditItem{
		{
			Name:   "inherited scalar-source-sealed tracebridge",
			Status: "BRIDGE_TEST_PASSED_SCALAR_SOURCE_SEALED",
			Supports: []string{
				"CONDITIONAL_SUPPORT_R3_TRACEBRIDGE_PRETEST_PASSED",
				"CONDITIONAL_SUPPORT_TRACEBRIDGE_CAN_BE_CLASSIFIED_AS_POST_ORIENTATION_BRIDGE_OBJECT",
			},
			Firewalls: []string{
				"FAILED_ROUTE_S_SPLIT_REMAINS_BRIDGE_HISTORY_SCALAR_INPUT",
				"FAILED_ROUTE_NO_NONCIRCULAR_NATIVE_S_SPLIT_PROXY_FOUND",
			},
		},
		{
			Name:   "full A_F descent",
			Status: "BLOCKED",
			Supports: []string{
				"CONDITIONAL_SUPPORT_FULL_A_F_DESCENT_AUDITED",
				"CONDITIONAL_SUPPORT_FULL_A_F_DESCENT_REMAINS_BLOCKED",
			},
			Firewalls: []string{
				"FAILED_ROUTE_FULL_A_F_DESCENT_STILL_BLOCKED",
				"FAILED_ROUTE_SOCKET_PROJECTORS_NOT_STABLE_UNDER_FULL_H_ACTION",
				"FAILED_ROUTE_POST_ORIENTATION_LEDGER_NOT_FULL_UNBROKEN_A_F_LEDGER",
			},
		},
		{
			Name:   "post-orientation stabilizer layer",
			Status: "LAWFUL_BRIDGE_SUPPORT",
			Supports: []string{
				"CONDITIONAL_SUPPORT_TRACEBRIDGE_STABLE_IN_A_F_ORIENT_LAYER",
				"CONDITIONAL_SUPPORT_A_F_ORIENT_IS_POST_ORIENTATION_STABILIZER",
				"CONDITIONAL_SUPPORT_R3_TRACEBRIDGE_HAS_LAWFUL_POST_ORIENTATION_SUPPORT",
			},
		},
		{
			Name:   "spontaneous orientation source candidate",
			Status: "SEALED_BRIDGE_LAYER_NOT_NATIVE",
			Supports: []string{
				"CONDITIONAL_SUPPORT_FINITE_ONE_FORM_IS_STRONGEST_ORIENTATION_SOURCE_CANDIDATE",
				"CONDITIONAL_SUPPORT_FINITE_ONE_FORM_REMAINS_STRONGEST_ORIENTATION_SOURCE_CANDIDATE",
				"CONDITIONAL_SUPPORT_NULL_EDGE_PATTERN_SOURCE_TYPES_POST_ORIENTATION_FRAME",
				"CONDITIONAL_SUPPORT_SPONTANEOUS_ORIENTATION_SEAL_IS_LAWFUL_BRIDGE_LAYER",
			},
			Firewalls: []string{
				"FAILED_ROUTE_NO_NATIVE_SPONTANEOUS_ORIENTATION_THEOREM",
				"FAILED_ROUTE_HIGGS_ORIENTATION_REMAINS_SEALED_AS_ORIENTATION_CLASS",
				"FAILED_ROUTE_FINITE_ONE_FORM_ORIENTATION_NOT_NATIVE_CERTIFIED",
			},
		},
		{
			Name:   "scalar seal interaction",
			Status: "ORIENTATION_CANNOT_RESCUE_NATIVE_R3",
			Supports: []string{
				"CONDITIONAL_SUPPORT_TRACEBRIDGE_CAN_BE_CLASSIFIED_AS_POST_ORIENTATION_BRIDGE_OBJECT",
			},
			Firewalls: []string{
				"FAILED_ROUTE_NATIVE_R3_NOT_GRANTED",
				"FAILED_ROUTE_S_SPLIT_REMAINS_BRIDGE_HISTORY_SCALAR_INPUT",
			},
		},
		{
			Name:   "R4 and physical claims boundary",
			Status: "OUT_OF_SCOPE_BLOCKED",
			Firewalls: []string{
				"FAILED_ROUTE_NO_OFFICIAL_LEDGER_UPDATE_ALLOWED",
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
		"CONDITIONAL_SUPPORT_FULL_A_F_DESCENT_AUDITED",
		"CONDITIONAL_SUPPORT_FULL_A_F_DESCENT_REMAINS_BLOCKED",
		"CONDITIONAL_SUPPORT_TRACEBRIDGE_STABLE_IN_A_F_ORIENT_LAYER",
		"CONDITIONAL_SUPPORT_A_F_ORIENT_IS_POST_ORIENTATION_STABILIZER",
		"CONDITIONAL_SUPPORT_A_F_ORIENT_IS_THE_CORRECT_POST_ORIENTATION_STABILIZER",
		"CONDITIONAL_SUPPORT_FINITE_ONE_FORM_IS_STRONGEST_ORIENTATION_SOURCE_CANDIDATE",
		"CONDITIONAL_SUPPORT_FINITE_ONE_FORM_REMAINS_STRONGEST_ORIENTATION_SOURCE_CANDIDATE",
		"CONDITIONAL_SUPPORT_SPONTANEOUS_ORIENTATION_SEAL_IS_LAWFUL_BRIDGE_LAYER",
		"CONDITIONAL_SUPPORT_TRACEBRIDGE_CAN_BE_CLASSIFIED_AS_POST_ORIENTATION_BRIDGE_OBJECT",
	}
}

func Failures() []string {
	return []string{
		"FAILED_ROUTE_NATIVE_R3_NOT_GRANTED",
		"FAILED_ROUTE_S_SPLIT_REMAINS_BRIDGE_HISTORY_SCALAR_INPUT",
		"FAILED_ROUTE_NO_NONCIRCULAR_NATIVE_S_SPLIT_PROXY_FOUND",
		"FAILED_ROUTE_FULL_A_F_DESCENT_STILL_BLOCKED",
		"FAILED_ROUTE_SOCKET_PROJECTORS_NOT_STABLE_UNDER_FULL_H_ACTION",
		"FAILED_ROUTE_POST_ORIENTATION_LEDGER_NOT_FULL_UNBROKEN_A_F_LEDGER",
		"FAILED_ROUTE_NO_NATIVE_SPONTANEOUS_ORIENTATION_THEOREM",
		"FAILED_ROUTE_HIGGS_ORIENTATION_REMAINS_SEALED_AS_ORIENTATION_CLASS",
		"FAILED_ROUTE_FINITE_ONE_FORM_ORIENTATION_NOT_NATIVE_CERTIFIED",
		"FAILED_ROUTE_NO_OFFICIAL_LEDGER_UPDATE_ALLOWED",
		"FAILED_ROUTE_NO_INDIVIDUAL_YUKAWA_VALUES",
		"FAILED_ROUTE_NO_PHYSICAL_PARTICLE_ASSIGNMENT",
		"FAILED_ROUTE_NO_GENERATION_CARRIER_MAP",
		"FAILED_ROUTE_NO_FLAVOR_ORIENTATION_MAP",
		"FAILED_ROUTE_NO_R4_NATIVE_YUKAWA_SPECTRUM_THEOREM",
	}
}

func FormatBoundary(b OrientationBoundary) string {
	return fmt.Sprintf("%s: scalarSeal=%v fullAFBlocked=%v stableInOrient=%v orientationSeal=%v nativeOrientation=%v nativeR3=%v officialLedger=%v physical=%v generationFlavor=%v", b.Status, b.ScalarSourceSealed, b.FullAFDescentBlocked, b.StableInAFOrient, b.SpontaneousOrientSeal, b.NativeOrientation, b.NativeR3, b.OfficialLedgerUpdate, b.PhysicalAssignment, b.GenerationFlavorClaims)
}

func FormatItem(i AuditItem) string {
	return fmt.Sprintf("%s: status=%s supports=[%s] firewalls=[%s]", i.Name, i.Status, strings.Join(i.Supports, ","), strings.Join(i.Firewalls, ","))
}

func ItemNotes(items []AuditItem) []string {
	out := make([]string, 0, len(items))
	for _, i := range items {
		out = append(out, FormatItem(i))
	}
	return out
}

func ItemSupports(items []AuditItem) []string {
	var out []string
	for _, i := range items {
		out = append(out, i.Supports...)
	}
	return out
}

func ItemFailures(items []AuditItem) []string {
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
