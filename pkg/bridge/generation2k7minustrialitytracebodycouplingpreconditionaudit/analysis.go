// Package generation2k7minustrialitytracebodycouplingpreconditionaudit implements
// Gate 954: K7Minus Triality Tracebody Coupling Precondition Audit.
//
// This audit is bridge-layer only. It tests the precondition for a lawful R4
// generation carrier coupling above the dual-sealed R3 tracebody. It does not
// derive generations, flavor, individual Yukawa values, CKM/PMNS, physical
// particles, or a native R4 spectrum theorem.
package generation2k7minustrialitytracebodycouplingpreconditionaudit

import (
	"fmt"
	"strings"
)

const (
	AuditID         = "GATE954-GENERATION2K7MINUSTRIALITYTRACEBODYCOUPLINGPRECONDITIONAUDIT"
	InheritedStatus = "R4_STRONG_GROUND_NO_GENERATION_OR_FLAVOR_MAP"
	Verdict         = "K7_MINUS_AND_TRIALITY_HAVE_COMPATIBLE_NATIVE_THREEFOLD_SHAPES_BUT_NO_TRACEBODY_COUPLING_MAP_CERTIFIED"
	Classification  = "R4_K7_MINUS_TRIALITY_COUPLING_PRECONDITION_SUPPORTED_NO_MAP"
	ShortStatus     = "R4_K7_MINUS_TRIALITY_COUPLING_CANDIDATE_NO_INTERTWINER"
	NextGate        = "NEXT_GATE955_K7MINUS_TRIALITY_INTERTWINER_CONSTRUCTION_AUDIT"
)

const (
	K7Dim      = 7
	K7PlusDim  = 4
	K7MinusDim = 3

	TraceRowRankA = 3
	TraceRowRankB = 3
	TraceRowRankC = 1

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

type CouplingShape struct {
	Carrier                 string
	CarrierDimension        int
	ActionShape             string
	ActionLanes             []string
	Target                  string
	TraceRowRanks           []int
	DualSealRequired        bool
	GenerationMapCertified  bool
	FlavorMapCertified      bool
	IntertwinerCertified    bool
	IndividualYukawaAllowed bool
	FlavorBacksolveAllowed  bool
}

type Analysis struct {
	AuditID                     string
	Inherited                   string
	Verdict                     string
	Classification              string
	ShortStatus                 string
	NativeR3                    bool
	DualSealRequired            bool
	GenerationCarrierCertified  bool
	FlavorOrientationCertified  bool
	IndividualYukawaCertified   bool
	PhysicalAssignmentCertified bool
	OfficialLedgerUpdate        bool
	Coupling                    CouplingShape
	Items                       []AuditItem
	Supports                    []string
	Failures                    []string
	Final                       string
}

func BuildDefault() (Analysis, error) {
	items := DefaultItems()
	if len(items) != 4 {
		return Analysis{}, fmt.Errorf("expected 4 audit items, got %d", len(items))
	}
	coupling := CouplingShape{
		Carrier:                 "K7^-",
		CarrierDimension:        K7MinusDim,
		ActionShape:             "triality threefold action-shape",
		ActionLanes:             []string{"vector lane", "left spinor lane", "right spinor lane"},
		Target:                  "dual-sealed aggregate R3 tracebody",
		TraceRowRanks:           []int{TraceRowRankA, TraceRowRankB, TraceRowRankC},
		DualSealRequired:        true,
		GenerationMapCertified:  false,
		FlavorMapCertified:      false,
		IntertwinerCertified:    false,
		IndividualYukawaAllowed: false,
		FlavorBacksolveAllowed:  false,
	}
	a := Analysis{
		AuditID:                     AuditID,
		Inherited:                   InheritedStatus,
		Verdict:                     Verdict,
		Classification:              Classification,
		ShortStatus:                 ShortStatus,
		NativeR3:                    false,
		DualSealRequired:            true,
		GenerationCarrierCertified:  false,
		FlavorOrientationCertified:  false,
		IndividualYukawaCertified:   false,
		PhysicalAssignmentCertified: false,
		OfficialLedgerUpdate:        false,
		Coupling:                    coupling,
		Items:                       items,
		Supports:                    Supports(),
		Failures:                    Failures(),
		Final:                       "Gate 954 supports the K7^-/triality synthesis as the next lawful R4 precondition: K7^- supplies the finite dimension-three carrier shape, triality supplies the native threefold action-shape, and the dual-sealed R3 tracebody is only an aggregate target. The missing object is the typed K7MinusTrialityTracebodyIntertwiner / GenerationCarrierMap.",
	}
	if err := validate(a); err != nil {
		return Analysis{}, err
	}
	return a, nil
}

func validate(a Analysis) error {
	if !a.DualSealRequired || !a.Coupling.DualSealRequired {
		return fmt.Errorf("R4 coupling audit must stay under explicit R3 dual seal")
	}
	if a.NativeR3 || a.OfficialLedgerUpdate || a.GenerationCarrierCertified || a.FlavorOrientationCertified || a.IndividualYukawaCertified || a.PhysicalAssignmentCertified {
		return fmt.Errorf("audit overclaimed forbidden theorem status: %#v", a)
	}
	if a.Coupling.GenerationMapCertified || a.Coupling.FlavorMapCertified || a.Coupling.IntertwinerCertified || a.Coupling.IndividualYukawaAllowed || a.Coupling.FlavorBacksolveAllowed {
		return fmt.Errorf("coupling shape overclaimed downstream map or flavor permission: %#v", a.Coupling)
	}
	if a.Coupling.CarrierDimension != 3 || K7MinusDim != 3 || K7PlusDim != 4 || K7Dim != K7PlusDim+K7MinusDim {
		return fmt.Errorf("bad K7 polarity dimensions")
	}
	if len(a.Coupling.ActionLanes) != 3 {
		return fmt.Errorf("triality candidate must record three action lanes")
	}
	if len(a.Coupling.TraceRowRanks) != 3 || a.Coupling.TraceRowRanks[0] != 3 || a.Coupling.TraceRowRanks[1] != 3 || a.Coupling.TraceRowRanks[2] != 1 {
		return fmt.Errorf("R3 tracebody row ranks must be aggregate 3,3,1")
	}
	return nil
}

func DefaultItems() []AuditItem {
	return []AuditItem{
		{
			Name:   "carrier/action compatibility",
			Status: "COMPATIBLE_SHAPE_NO_COUPLING",
			Supports: []string{
				"CONDITIONAL_SUPPORT_K7_MINUS_AND_TRIALITY_HAVE_COMPATIBLE_NATIVE_THREEFOLD_SHAPES",
				"CONDITIONAL_SUPPORT_K7_MINUS_SUPPLIES_FINITE_DIMENSION_THREE_CARRIER_SHAPE",
				"CONDITIONAL_SUPPORT_TRIALITY_SUPPLIES_NATIVE_THREEFOLD_ACTION_SHAPE",
				"CONDITIONAL_SUPPORT_K7_MINUS_TRIALITY_SYNTHESIS_IS_LAWFUL_NEXT_R4_PRECONDITION",
			},
			Firewalls: []string{
				"FAILED_ROUTE_SHAPE_COMPATIBILITY_NOT_COUPLING_THEOREM",
				"FAILED_ROUTE_DIM_K7_MINUS_EQUALS_THREE_NOT_GENERATION_THEOREM",
				"FAILED_ROUTE_TRIALITY_THREEFOLD_SHAPE_NOT_GENERATION_THEOREM",
			},
		},
		{
			Name:   "typed action/intertwiner requirement",
			Status: "MISSING_OBJECT_EXPLICIT",
			Supports: []string{
				"CONDITIONAL_SUPPORT_REQUIRED_OBJECT_IS_K7_MINUS_TRIALITY_TRACEBODY_INTERTWINER",
				"CONDITIONAL_SUPPORT_GENERATION_CARRIER_MAP_MUST_BE_TYPED_MAP_NOT_NUMBER",
			},
			Firewalls: []string{
				"FAILED_ROUTE_NO_GENERATION_CARRIER_MAP",
				"FAILED_ROUTE_NO_K7_MINUS_TRIALITY_ACTION_MAP",
				"FAILED_ROUTE_NO_K7_MINUS_TRIALITY_TO_R3_TRACEBODY_INTERTWINER",
			},
		},
		{
			Name:   "R3 tracebody boundary",
			Status: "AGGREGATE_TARGET_ONLY",
			Supports: []string{
				"CONDITIONAL_SUPPORT_R3_TRACEBODY_CAN_BE_USED_ONLY_AS_DUALSEALED_AGGREGATE_TARGET",
			},
			Firewalls: []string{
				"FAILED_ROUTE_R3_TRACE_ROWS_ARE_NOT_GENERATION_LABELS",
				"FAILED_ROUTE_3_PLUS_3_PLUS_1_NOT_FLAVOR_DECOMPOSITION",
				"FAILED_ROUTE_R3_TRACEBRIDGE_REMAINS_SCALAR_SOURCE_SEALED",
				"FAILED_ROUTE_R3_TRACEBRIDGE_REMAINS_POST_ORIENTATION_SEALED",
			},
		},
		{
			Name:   "flavor firewall",
			Status: "DOWNSTREAM_ONLY",
			Supports: []string{
				"CONDITIONAL_SUPPORT_FLAVOR_FORMULAS_REMAIN_DOWNSTREAM_LEDGER_TARGETS_ONLY",
			},
			Firewalls: []string{
				"FAILED_ROUTE_FLAVOR_FORMULA_BACKSOLVE_IS_CIRCULAR",
				"FAILED_ROUTE_NO_FLAVOR_ORIENTATION_MAP",
				"FAILED_ROUTE_NO_INDIVIDUAL_YUKAWA_VALUES",
				"FAILED_ROUTE_NO_PHYSICAL_PARTICLE_ASSIGNMENT",
				"FAILED_ROUTE_NO_CKM_PMNS_THEOREM",
				"FAILED_ROUTE_NO_R4_NATIVE_YUKAWA_SPECTRUM_THEOREM",
				"FAILED_ROUTE_NO_OFFICIAL_LEDGER_UPDATE_ALLOWED",
			},
		},
	}
}

func Supports() []string {
	return []string{
		"CONDITIONAL_SUPPORT_K7_MINUS_AND_TRIALITY_HAVE_COMPATIBLE_NATIVE_THREEFOLD_SHAPES",
		"CONDITIONAL_SUPPORT_K7_MINUS_SUPPLIES_FINITE_DIMENSION_THREE_CARRIER_SHAPE",
		"CONDITIONAL_SUPPORT_TRIALITY_SUPPLIES_NATIVE_THREEFOLD_ACTION_SHAPE",
		"CONDITIONAL_SUPPORT_K7_MINUS_TRIALITY_SYNTHESIS_IS_LAWFUL_NEXT_R4_PRECONDITION",
		"CONDITIONAL_SUPPORT_R3_TRACEBODY_CAN_BE_USED_ONLY_AS_DUALSEALED_AGGREGATE_TARGET",
		"CONDITIONAL_SUPPORT_FLAVOR_FORMULAS_REMAIN_DOWNSTREAM_LEDGER_TARGETS_ONLY",
	}
}

func Failures() []string {
	return []string{
		"FAILED_ROUTE_NO_GENERATION_CARRIER_MAP",
		"FAILED_ROUTE_NO_K7_MINUS_TRIALITY_ACTION_MAP",
		"FAILED_ROUTE_NO_K7_MINUS_TRIALITY_TO_R3_TRACEBODY_INTERTWINER",
		"FAILED_ROUTE_SHAPE_COMPATIBILITY_NOT_COUPLING_THEOREM",
		"FAILED_ROUTE_DIM_K7_MINUS_EQUALS_THREE_NOT_GENERATION_THEOREM",
		"FAILED_ROUTE_TRIALITY_THREEFOLD_SHAPE_NOT_GENERATION_THEOREM",
		"FAILED_ROUTE_R3_TRACE_ROWS_ARE_NOT_GENERATION_LABELS",
		"FAILED_ROUTE_3_PLUS_3_PLUS_1_NOT_FLAVOR_DECOMPOSITION",
		"FAILED_ROUTE_FLAVOR_FORMULA_BACKSOLVE_IS_CIRCULAR",
		"FAILED_ROUTE_NO_FLAVOR_ORIENTATION_MAP",
		"FAILED_ROUTE_NO_INDIVIDUAL_YUKAWA_VALUES",
		"FAILED_ROUTE_NO_PHYSICAL_PARTICLE_ASSIGNMENT",
		"FAILED_ROUTE_NO_CKM_PMNS_THEOREM",
		"FAILED_ROUTE_NO_R4_NATIVE_YUKAWA_SPECTRUM_THEOREM",
		"FAILED_ROUTE_R3_TRACEBRIDGE_REMAINS_SCALAR_SOURCE_SEALED",
		"FAILED_ROUTE_R3_TRACEBRIDGE_REMAINS_POST_ORIENTATION_SEALED",
		"FAILED_ROUTE_NO_OFFICIAL_LEDGER_UPDATE_ALLOWED",
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
