// Package generation2r4generationcarriercomparisonflavorformulafirewallaudit implements
// Gate 953: R4 GenerationCarrier Candidate Comparison and FlavorFormula Firewall Audit.
//
// This audit is bridge-layer only. It preserves the R3 dual seals and blocks
// physical particle, individual Yukawa, and official-ledger claims.
package generation2r4generationcarriercomparisonflavorformulafirewallaudit

import (
	"fmt"
	"strings"
)

const (
	AuditID         = "GATE953-GENERATION2R4GENERATIONCARRIERCOMPARISONFLAVORFORMULAFIREWALLAUDIT"
	InheritedStatus = "R4_TRIALITY_CANDIDATE_NO_TRACEBODY_ACTION"
	Verdict         = "R4_GENERATION_FRONTIER_GROUNDED_ON_K7_MINUS_AND_TRIALITY_CANDIDATES_BUT_GENERATION_CARRIER_MAP_AND_FLAVOR_ORIENTATION_REMAIN_MISSING"
	Classification  = "R4_STRONG_GROUND_CANDIDATES_IDENTIFIED_FLAVOR_FIREWALL_PRESERVED"
	ShortStatus     = "R4_STRONG_GROUND_NO_GENERATION_OR_FLAVOR_MAP"
	NextGate        = "NEXT_GATE954_K7MINUS_TRIALITY_TRACEBODY_COUPLING_PRECONDITION_AUDIT"
)

const (
	Ssplit  = 0.0012924448188162962
	AlphaB  = 0.0003878958469680527
	NEff    = 3.002327375081808
	CYukawa = 0.9992248096922658
	CHiggs  = 1.037220510866514
)

type AuditItem struct {
	Name      string
	Status    string
	Supports  []string
	Firewalls []string
}

type Analysis struct {
	AuditID                     string
	Inherited                   string
	Verdict                     string
	Classification              string
	ShortStatus                 string
	DualSealRequired            bool
	NativeR3                    bool
	OfficialLedgerUpdate        bool
	GenerationCarrierCertified  bool
	FlavorOrientationCertified  bool
	IndividualYukawaCertified   bool
	PhysicalAssignmentCertified bool
	Items                       []AuditItem
	Supports                    []string
	Failures                    []string
	Final                       string
}

func BuildDefault() (Analysis, error) {
	items := DefaultItems()
	if len(items) != 5 {
		return Analysis{}, fmt.Errorf("expected 5 audit items, got %d", len(items))
	}
	a := Analysis{
		AuditID:                     AuditID,
		Inherited:                   InheritedStatus,
		Verdict:                     Verdict,
		Classification:              Classification,
		ShortStatus:                 ShortStatus,
		DualSealRequired:            true,
		NativeR3:                    false,
		OfficialLedgerUpdate:        false,
		GenerationCarrierCertified:  false,
		FlavorOrientationCertified:  false,
		IndividualYukawaCertified:   false,
		PhysicalAssignmentCertified: false,
		Items:                       items,
		Supports:                    Supports(),
		Failures:                    Failures(),
		Final:                       "Gate 953 establishes strong R4 ground: K7^- is the best finite dimension-three carrier shape, triality is the best origin threefold action candidate, and flavor formulas remain downstream targets. The next lawful step is a K7^-/triality-to-tracebody coupling precondition audit, not Yukawa-spectrum derivation.",
	}
	if !a.DualSealRequired {
		return Analysis{}, fmt.Errorf("R4 rail must remain under explicit R3 dual seal")
	}
	if a.NativeR3 || a.OfficialLedgerUpdate || a.GenerationCarrierCertified || a.FlavorOrientationCertified || a.IndividualYukawaCertified || a.PhysicalAssignmentCertified {
		return Analysis{}, fmt.Errorf("audit overclaimed downstream theorem status: %#v", a)
	}
	return a, nil
}

func DefaultItems() []AuditItem {
	return []AuditItem{
		{
			Name:   "candidate ranking",
			Status: "STRONG_GROUND",
			Supports: []string{
				"CONDITIONAL_SUPPORT_K7_MINUS_IS_STRONGEST_FINITE_DIMENSION_THREE_CARRIER_SHAPE",
				"CONDITIONAL_SUPPORT_TRIALITY_IS_STRONGEST_ORIGIN_THREEFOLD_ACTION_CANDIDATE",
				"CONDITIONAL_SUPPORT_K7_MINUS_TRIALITY_SYNTHESIS_IS_NEXT_PRECONDITION",
			},
		},
		{
			Name:   "R3 aggregate tracebody boundary",
			Status: "DUAL_SEALED_INPUT_ONLY",
			Supports: []string{
				"CONDITIONAL_SUPPORT_R3_TRACEBODY_AVAILABLE_ONLY_AS_AGGREGATE_DUALSEALED_INPUT",
			},
			Firewalls: []string{
				"FAILED_ROUTE_R3_TRACEBODY_NOT_GENERATION_LABELS",
				"FAILED_ROUTE_R3_TRACEBRIDGE_REMAINS_SCALAR_SOURCE_SEALED",
				"FAILED_ROUTE_R3_TRACEBRIDGE_REMAINS_POST_ORIENTATION_SEALED",
			},
		},
		{
			Name:   "previous flavor formula memory",
			Status: "FIREWALLED_PRECONDITION_ONLY",
			Supports: []string{
				"CONDITIONAL_SUPPORT_EXISTING_FLAVOR_FORMULAS_CAN_BE_USED_ONLY_AFTER_GENERATION_CARRIER_AND_FLAVOR_ORIENTATION_MAPS",
				"CONDITIONAL_SUPPORT_FLAVOR_WALL_FORMULAS_ARE_LEDGER_TARGETS_NOT_CARRIER_SOURCES",
			},
			Firewalls: []string{
				"FAILED_ROUTE_EPSILON_E_OR_KAPPA_E_FORMULAS_CANNOT_SOURCE_GENERATION_CARRIER",
				"FAILED_ROUTE_CKM_PMNS_LEDGER_NOT_GENERATION_CARRIER",
				"FAILED_ROUTE_FLAVOR_FORMULA_BACKSOLVE_IS_CIRCULAR",
			},
		},
		{
			Name:   "external family factor",
			Status: "SEAL_ONLY",
			Supports: []string{
				"CONDITIONAL_SUPPORT_EXTERNAL_C3_FAMILY_FACTOR_REMAINS_OPTIONAL_SEALED_COMPARATOR_ONLY",
			},
			Firewalls: []string{
				"FAILED_ROUTE_EXTERNAL_C3_FAMILY_FACTOR_NOT_NATIVE_ASHA_GENERATION_CARRIER",
			},
		},
		{
			Name:   "R4 next object",
			Status: "EXPLICIT",
			Supports: []string{
				"CONDITIONAL_SUPPORT_R4_REQUIRES_K7_MINUS_TRIALITY_TRACEBODY_COUPLING_PRECONDITION",
			},
			Firewalls: []string{
				"FAILED_ROUTE_NO_GENERATION_CARRIER_MAP",
				"FAILED_ROUTE_NO_FLAVOR_ORIENTATION_MAP",
				"FAILED_ROUTE_NO_R4_NATIVE_YUKAWA_SPECTRUM_THEOREM",
			},
		},
	}
}

func Supports() []string {
	return []string{
		"CONDITIONAL_SUPPORT_K7_MINUS_IS_STRONGEST_FINITE_DIMENSION_THREE_CARRIER_SHAPE",
		"CONDITIONAL_SUPPORT_TRIALITY_IS_STRONGEST_ORIGIN_THREEFOLD_ACTION_CANDIDATE",
		"CONDITIONAL_SUPPORT_K7_MINUS_TRIALITY_SYNTHESIS_IS_NEXT_PRECONDITION",
		"CONDITIONAL_SUPPORT_R3_TRACEBODY_AVAILABLE_ONLY_AS_AGGREGATE_DUALSEALED_INPUT",
		"CONDITIONAL_SUPPORT_EXISTING_FLAVOR_FORMULAS_CAN_BE_USED_ONLY_AFTER_GENERATION_CARRIER_AND_FLAVOR_ORIENTATION_MAPS",
		"CONDITIONAL_SUPPORT_FLAVOR_WALL_FORMULAS_ARE_LEDGER_TARGETS_NOT_CARRIER_SOURCES",
		"CONDITIONAL_SUPPORT_R4_REQUIRES_K7_MINUS_TRIALITY_TRACEBODY_COUPLING_PRECONDITION",
	}
}

func Failures() []string {
	return []string{
		"FAILED_ROUTE_NO_GENERATION_CARRIER_MAP",
		"FAILED_ROUTE_NO_FLAVOR_ORIENTATION_MAP",
		"FAILED_ROUTE_NO_INDIVIDUAL_YUKAWA_VALUES",
		"FAILED_ROUTE_NO_PHYSICAL_PARTICLE_ASSIGNMENT",
		"FAILED_ROUTE_NO_CKM_PMNS_THEOREM",
		"FAILED_ROUTE_NO_R4_NATIVE_YUKAWA_SPECTRUM_THEOREM",
		"FAILED_ROUTE_EPSILON_E_OR_KAPPA_E_FORMULAS_CANNOT_SOURCE_GENERATION_CARRIER",
		"FAILED_ROUTE_CKM_PMNS_LEDGER_NOT_GENERATION_CARRIER",
		"FAILED_ROUTE_FLAVOR_FORMULA_BACKSOLVE_IS_CIRCULAR",
		"FAILED_ROUTE_R3_TRACEBODY_NOT_GENERATION_LABELS",
	}
}

func ItemSupports(items []AuditItem) []string {
	var out []string
	for _, it := range items {
		out = append(out, it.Supports...)
	}
	return out
}

func ItemFailures(items []AuditItem) []string {
	var out []string
	for _, it := range items {
		out = append(out, it.Firewalls...)
	}
	return out
}

func appendAll(parts ...[]string) []string {
	var out []string
	for _, p := range parts {
		out = append(out, p...)
	}
	return out
}

func containsAll(hay []string, needles []string) bool {
	set := map[string]bool{}
	for _, h := range hay {
		set[h] = true
	}
	for _, n := range needles {
		if !set[n] {
			return false
		}
	}
	return true
}

func stringsJoin(v []string) string { return strings.Join(v, "; ") }

func ItemNotes(items []AuditItem) []string {
	notes := make([]string, 0, len(items))
	for _, it := range items {
		notes = append(notes, it.Name+" => "+it.Status)
	}
	return notes
}
