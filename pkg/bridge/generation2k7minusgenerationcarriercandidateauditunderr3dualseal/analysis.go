// Package generation2k7minusgenerationcarriercandidateauditunderr3dualseal implements
// Gate 951: K7Minus GenerationCarrier Candidate Audit Under R3 DualSeal.
//
// This audit is bridge-layer only. It preserves the R3 dual seals and blocks
// physical particle, individual Yukawa, and official-ledger claims.
package generation2k7minusgenerationcarriercandidateauditunderr3dualseal

import (
	"fmt"
	"strings"
)

const (
	AuditID         = "GATE951-GENERATION2K7MINUSGENERATIONCARRIERCANDIDATEAUDITUNDERR3DUALSEAL"
	InheritedStatus = "R4_GENERATION_CARRIER_NOT_YET_CERTIFIED"
	Verdict         = "K7_MINUS_HAS_NATIVE_DIMENSION_THREE_CARRIER_SHAPE_BUT_NO_TYPED_MAP_TO_R3_TRACEBODY_CERTIFIED"
	Classification  = "R4_K7_MINUS_GENERATION_CARRIER_CANDIDATE_SUPPORTED_NO_MAP"
	ShortStatus     = "R4_K7_MINUS_CANDIDATE_NO_GENERATION_MAP"
	NextGate        = "NEXT_GATE952_TRIALITY_GENERATIONCARRIER_CANDIDATE_AUDIT_UNDER_R3_DUALSEAL"
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
		Final:                       "Gate 951 makes K7^- the strongest finite nonempirical dimension-three carrier candidate, but it does not certify generation: the missing object is a typed K7^- to R3 tracebody/socket-ledger action map.",
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
			Name:   "K7 Hodge-polarity inheritance",
			Status: "NATIVE_SHAPE_STRONG",
			Supports: []string{
				"CONDITIONAL_SUPPORT_K7_HODGE_POLARITY_INHERITS_NATIVE_SPLIT_SIGNATURE_CARRIER",
				"CONDITIONAL_SUPPORT_K7_MINUS_DIMENSION_THREE_IS_NONEMPIRICAL",
			},
		},
		{
			Name:   "dimension-three carrier shape",
			Status: "CANDIDATE_ONLY",
			Supports: []string{
				"CONDITIONAL_SUPPORT_K7_MINUS_HAS_CORRECT_DIMENSION_FOR_GENERATION_CARRIER_SHAPE",
			},
			Firewalls: []string{
				"FAILED_ROUTE_DIMENSION_THREE_ALONE_NOT_GENERATION_THEOREM",
			},
		},
		{
			Name:   "map to R3 tracebody",
			Status: "MISSING",
			Firewalls: []string{
				"FAILED_ROUTE_NO_TYPED_MAP_FROM_K7_MINUS_TO_R3_TRACEBODY",
				"FAILED_ROUTE_NO_K7_MINUS_ACTION_ON_Y_DAGGER_Y_ROWS",
				"FAILED_ROUTE_NO_K7_MINUS_TO_SOCKET_LEDGER_INTERTWINER",
			},
		},
		{
			Name:   "dual seal compatibility",
			Status: "COMPATIBLE_UNDER_SEAL",
			Supports: []string{
				"CONDITIONAL_SUPPORT_K7_MINUS_CAN_BE_AUDITED_WITHOUT_BREAKING_SCALAR_SOURCE_SEAL",
				"CONDITIONAL_SUPPORT_K7_MINUS_CAN_BE_AUDITED_WITHOUT_FULL_AF_DESCENT",
			},
			Firewalls: []string{
				"FAILED_ROUTE_R3_TRACEBRIDGE_REMAINS_DUALSEALED_INPUT",
			},
		},
		{
			Name:   "flavor formula relation",
			Status: "PREMATURE",
			Supports: []string{
				"CONDITIONAL_SUPPORT_K7_MINUS_MAY_LATER_HOST_FLAVOR_ORIENTATION_IF_MAP_EXISTS",
			},
			Firewalls: []string{
				"FAILED_ROUTE_FLAVOR_FORMULA_CANNOT_SOURCE_GENERATION_CARRIER",
				"FAILED_ROUTE_K7_MINUS_NOT_YET_FLAVOR_ORIENTATION_MAP",
			},
		},
	}
}

func Supports() []string {
	return []string{
		"CONDITIONAL_SUPPORT_K7_MINUS_HAS_NATIVE_DIMENSION_THREE_CARRIER_SHAPE",
		"CONDITIONAL_SUPPORT_K7_HODGE_POLARITY_INHERITS_NATIVE_SPLIT_SIGNATURE_CARRIER",
		"CONDITIONAL_SUPPORT_K7_CANDIDATE_IS_NONEMPIRICAL_AND_FINITE",
		"CONDITIONAL_SUPPORT_K7_MINUS_CAN_BE_AUDITED_WITHOUT_BREAKING_R3_DUALSEAL",
		"CONDITIONAL_SUPPORT_K7_MINUS_MAY_LATER_HOST_FLAVOR_ORIENTATION_IF_MAP_EXISTS",
	}
}

func Failures() []string {
	return []string{
		"FAILED_ROUTE_DIM_K7_MINUS_EQUALS_THREE_NOT_GENERATION_THEOREM",
		"FAILED_ROUTE_NO_K7_POLARITY_TO_GENERATION_CARRIER_MAP",
		"FAILED_ROUTE_NO_TYPED_MAP_FROM_K7_MINUS_TO_R3_TRACEBODY",
		"FAILED_ROUTE_NO_K7_MINUS_ACTION_ON_Y_DAGGER_Y_ROWS",
		"FAILED_ROUTE_NO_K7_MINUS_TO_SOCKET_LEDGER_INTERTWINER",
		"FAILED_ROUTE_FLAVOR_FORMULA_CANNOT_SOURCE_GENERATION_CARRIER",
		"FAILED_ROUTE_NO_FLAVOR_ORIENTATION_MAP",
		"FAILED_ROUTE_NO_INDIVIDUAL_YUKAWA_VALUES",
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
