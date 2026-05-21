// Package generation2r4generationcarrierpreconditionauditunderr3dualseal implements
// Gate 950: R4 GenerationCarrier Precondition Audit Under R3 DualSeal.
//
// This audit is bridge-layer only. It preserves the R3 dual seals and blocks
// physical particle, individual Yukawa, and official-ledger claims.
package generation2r4generationcarrierpreconditionauditunderr3dualseal

import (
	"fmt"
	"strings"
)

const (
	AuditID         = "GATE950-GENERATION2R4GENERATIONCARRIERPRECONDITIONAUDITUNDERR3DUALSEAL"
	InheritedStatus = "R3_DUALSEAL_TRACEBRIDGE_FINAL_BOUNDARY"
	Verdict         = "R4_GENERATION_CARRIER_PRECONDITIONS_AUDITED_UNDER_R3_DUALSEAL_K7_MINUS_AND_TRIALITY_ARE_STRONGEST_CANDIDATES_BUT_NO_GENERATION_CARRIER_MAP_CERTIFIED"
	Classification  = "R4_GENERATION_CARRIER_PRECONDITION_AUDIT_CANDIDATES_FOUND_NO_MAP"
	ShortStatus     = "R4_GENERATION_CARRIER_NOT_YET_CERTIFIED"
	NextGate        = "NEXT_GATE951_K7MINUS_GENERATIONCARRIER_CANDIDATE_AUDIT_UNDER_R3_DUALSEAL"
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
	if len(items) != 7 {
		return Analysis{}, fmt.Errorf("expected 7 audit items, got %d", len(items))
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
		Final:                       "Gate 950 opens R4 only as a precondition audit under the explicit R3 dual seal. It blocks the rank-three-to-three-generations shortcut and identifies K7^- plus triality as the strongest candidates, with no GenerationCarrierMap certified.",
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
			Name:   "R3 trace rows are aggregate socket tracebody, not generations",
			Status: "BLOCKED_AS_GENERATION_SOURCE",
			Supports: []string{
				"CONDITIONAL_SUPPORT_R3_TRACE_ROWS_PROVIDE_AGGREGATE_YUKAWA_TRACEBODY_UNDER_SEAL",
				"CONDITIONAL_SUPPORT_R4_CAN_USE_TRACEBODY_ONLY_AS_AGGREGATE_INPUT",
			},
			Firewalls: []string{
				"FAILED_ROUTE_R3_TRACE_ROWS_ARE_NOT_GENERATION_CARRIER",
				"FAILED_ROUTE_RANK_THREE_SOCKET_ATOMS_NOT_GENERATION_THEOREM",
				"FAILED_ROUTE_3_PLUS_3_PLUS_1_NOT_PHYSICAL_FAMILY_SPLIT",
			},
		},
		{
			Name:   "color rank three is not generation rank three",
			Status: "BLOCKED_AS_FAMILY_SOURCE",
			Firewalls: []string{
				"FAILED_ROUTE_COLOR_RANK_THREE_NOT_GENERATION_CARRIER",
				"FAILED_ROUTE_P3_MULTIPLICITY_NOT_FAMILY_MULTIPLICITY",
				"FAILED_ROUTE_NO_COLOR_TO_GENERATION_TRANSMUTATION_THEOREM",
			},
		},
		{
			Name:   "external C3 family factor",
			Status: "QUARANTINED_SEAL_ONLY",
			Supports: []string{
				"CONDITIONAL_SUPPORT_EXTERNAL_C3_FAMILY_FACTOR_CAN_BE_USED_ONLY_AS_SEAL",
			},
			Firewalls: []string{
				"FAILED_ROUTE_EXTERNAL_C3_FAMILY_FACTOR_NOT_NATIVE_ASHA_GENERATION_CARRIER",
				"FAILED_ROUTE_INSERTED_FAMILY_MULTIPLICITY_NOT_R4_THEOREM",
			},
		},
		{
			Name:   "K7 minus Hodge-polarity candidate",
			Status: "STRONGEST_FINITE_CANDIDATE_SHAPE",
			Supports: []string{
				"CONDITIONAL_SUPPORT_K7_MINUS_HAS_NATIVE_DIMENSION_THREE_CARRIER_SHAPE",
				"CONDITIONAL_SUPPORT_K7_HODGE_POLARITY_IS_STRONGEST_NATIVE_GENERATION_CARRIER_CANDIDATE",
				"CONDITIONAL_SUPPORT_K7_CANDIDATE_IS_NONEMPIRICAL_AND_FINITE",
			},
			Firewalls: []string{
				"FAILED_ROUTE_DIM_K7_MINUS_EQUALS_THREE_NOT_GENERATION_THEOREM",
				"FAILED_ROUTE_NO_K7_POLARITY_TO_GENERATION_CARRIER_MAP",
				"FAILED_ROUTE_NO_TYPED_MAP_FROM_K7_MINUS_TO_YUKAWA_SOCKET_LEDGER",
				"FAILED_ROUTE_K7_HODGE_POLARITY_NOT_YET_R4_GENERATION_CARRIER",
			},
		},
		{
			Name:   "triality candidate",
			Status: "DEEP_NATIVE_THREEFOLD_CANDIDATE",
			Supports: []string{
				"CONDITIONAL_SUPPORT_TRIALITY_IS_DEEP_NATIVE_THREEFOLD_SOURCE_CANDIDATE",
				"CONDITIONAL_SUPPORT_CL17_ROOT_MAKES_TRIALITY_A_LAWFUL_R4_SOURCE_TO_AUDIT",
			},
			Firewalls: []string{
				"FAILED_ROUTE_TRIALITY_THREEFOLD_SHAPE_NOT_GENERATION_THEOREM",
				"FAILED_ROUTE_NO_TRIALITY_TO_GENERATION_CARRIER_MAP",
				"FAILED_ROUTE_NO_ACTION_OF_TRIALITY_CARRIER_ON_R3_TRACEBRIDGE",
				"FAILED_ROUTE_TRIALITY_NOT_YET_FLAVOR_OR_GENERATION_SPLITTING_THEOREM",
			},
		},
		{
			Name:   "Fock/projective 1+3 selector",
			Status: "NOT_GENERATION_CARRIER",
			Firewalls: []string{
				"FAILED_ROUTE_FOCK_PROJECTIVE_1_PLUS_3_SELECTOR_NOT_GENERATION_CARRIER",
				"FAILED_ROUTE_B_MINUS_L_SELECTOR_NOT_FAMILY_SPLITTING_THEOREM",
			},
		},
		{
			Name:   "R4 inheritance policy under dual seal",
			Status: "AGGREGATE_ONLY",
			Supports: []string{
				"CONDITIONAL_SUPPORT_R4_MAY_PROCEED_UNDER_EXPLICIT_R3_DUALSEAL",
				"CONDITIONAL_SUPPORT_R3_TRACEBODY_CAN_BE_USED_ONLY_AS_AGGREGATE_INPUT",
				"CONDITIONAL_SUPPORT_R4_REQUIRES_GENERATION_CARRIER_MAP_BEFORE_FLAVOR_OR_YUKAWA_SPECTRUM",
			},
			Firewalls: []string{
				"FAILED_ROUTE_NO_FLAVOR_ORIENTATION_MAP",
				"FAILED_ROUTE_NO_INDIVIDUAL_YUKAWA_VALUES",
				"FAILED_ROUTE_NO_PHYSICAL_PARTICLE_ASSIGNMENT",
				"FAILED_ROUTE_NO_CKM_PMNS_THEOREM",
				"FAILED_ROUTE_NO_R4_NATIVE_YUKAWA_SPECTRUM_THEOREM",
			},
		},
	}
}

func Supports() []string {
	return []string{
		"CONDITIONAL_SUPPORT_R4_MAY_PROCEED_UNDER_EXPLICIT_R3_DUALSEAL",
		"CONDITIONAL_SUPPORT_R3_TRACEBODY_CAN_BE_USED_ONLY_AS_AGGREGATE_INPUT",
		"CONDITIONAL_SUPPORT_K7_MINUS_HAS_NATIVE_DIMENSION_THREE_CARRIER_SHAPE",
		"CONDITIONAL_SUPPORT_K7_HODGE_POLARITY_IS_STRONGEST_FINITE_GENERATION_CARRIER_CANDIDATE",
		"CONDITIONAL_SUPPORT_TRIALITY_IS_DEEP_NATIVE_THREEFOLD_SOURCE_CANDIDATE",
		"CONDITIONAL_SUPPORT_EXTERNAL_C3_FAMILY_FACTOR_CAN_BE_USED_ONLY_AS_SEAL",
		"CONDITIONAL_SUPPORT_R4_REQUIRES_GENERATION_CARRIER_MAP_BEFORE_FLAVOR_OR_YUKAWA_SPECTRUM",
	}
}

func Failures() []string {
	return []string{
		"FAILED_ROUTE_NO_GENERATION_CARRIER_MAP",
		"FAILED_ROUTE_R3_TRACE_ROWS_ARE_NOT_GENERATION_CARRIER",
		"FAILED_ROUTE_COLOR_RANK_THREE_NOT_GENERATION_CARRIER",
		"FAILED_ROUTE_DIM_K7_MINUS_EQUALS_THREE_NOT_GENERATION_THEOREM",
		"FAILED_ROUTE_NO_K7_POLARITY_TO_GENERATION_CARRIER_MAP",
		"FAILED_ROUTE_TRIALITY_THREEFOLD_SHAPE_NOT_GENERATION_THEOREM",
		"FAILED_ROUTE_NO_TRIALITY_TO_GENERATION_CARRIER_MAP",
		"FAILED_ROUTE_EXTERNAL_C3_FAMILY_FACTOR_NOT_NATIVE_ASHA_GENERATION_CARRIER",
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
