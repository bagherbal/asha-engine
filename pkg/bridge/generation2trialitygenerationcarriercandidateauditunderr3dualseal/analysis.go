// Package generation2trialitygenerationcarriercandidateauditunderr3dualseal implements
// Gate 952: Triality GenerationCarrier Candidate Audit Under R3 DualSeal.
//
// This audit is bridge-layer only. It preserves the R3 dual seals and blocks
// physical particle, individual Yukawa, and official-ledger claims.
package generation2trialitygenerationcarriercandidateauditunderr3dualseal

import (
	"fmt"
	"strings"
)

const (
	AuditID         = "GATE952-GENERATION2TRIALITYGENERATIONCARRIERCANDIDATEAUDITUNDERR3DUALSEAL"
	InheritedStatus = "R4_K7_MINUS_CANDIDATE_NO_GENERATION_MAP"
	Verdict         = "TRIALITY_THREEFOLD_SOURCE_IS_DEEP_NATIVE_CANDIDATE_BUT_NO_ACTION_ON_R3_TRACEBODY_CERTIFIED"
	Classification  = "R4_TRIALITY_GENERATION_CARRIER_CANDIDATE_SUPPORTED_NO_ACTION_MAP"
	ShortStatus     = "R4_TRIALITY_CANDIDATE_NO_TRACEBODY_ACTION"
	NextGate        = "NEXT_GATE953_R4_GENERATIONCARRIER_CANDIDATE_COMPARISON_AND_FLAVOR_FIREWALL_AUDIT"
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
		Final:                       "Gate 952 audits triality as the deepest native threefold source candidate. It remains a candidate only: no action of the triality carrier on the dual-sealed R3 tracebody is certified.",
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
			Name:   "Cl(1,7)/Spin(8)-style triality source",
			Status: "DEEP_NATIVE_CANDIDATE",
			Supports: []string{
				"CONDITIONAL_SUPPORT_CL17_ROOT_MAKES_TRIALITY_A_LAWFUL_THREEFOLD_SOURCE",
				"CONDITIONAL_SUPPORT_VECTOR_SPINOR_SPINOR_THREEFOLD_IS_STRUCTURAL_NOT_EMPIRICAL",
			},
		},
		{
			Name:   "triality not generations by itself",
			Status: "FIREWALLED",
			Firewalls: []string{
				"FAILED_ROUTE_TRIALITY_THREEFOLD_SHAPE_NOT_GENERATION_THEOREM",
				"FAILED_ROUTE_VECTOR_SPINOR_SPINOR_ORBIT_NOT_OBSERVED_FAMILY_MULTIPLICITY",
			},
		},
		{
			Name:   "action on R3 tracebody",
			Status: "MISSING",
			Firewalls: []string{
				"FAILED_ROUTE_NO_TRIALITY_ACTION_ON_R3_TRACEBODY",
				"FAILED_ROUTE_NO_TRIALITY_TO_YUKAWA_SOCKET_LEDGER_FUNCTOR",
				"FAILED_ROUTE_TRIALITY_NOT_YET_FLAVOR_SPLITTING_THEOREM",
			},
		},
		{
			Name:   "relation to K7 minus",
			Status: "POSSIBLE_SYNTHESIS",
			Supports: []string{
				"CONDITIONAL_SUPPORT_TRIALITY_MAY_SUPPLY_ACTION_PRINCIPLE_FOR_K7_MINUS_CARRIER",
			},
			Firewalls: []string{
				"FAILED_ROUTE_NO_TRIALITY_K7_MINUS_COUPLING_THEOREM",
			},
		},
		{
			Name:   "dual seal compatibility",
			Status: "COMPATIBLE_UNDER_SEAL",
			Supports: []string{
				"CONDITIONAL_SUPPORT_TRIALITY_AUDIT_RESPECTS_SCALAR_SOURCE_AND_POST_ORIENTATION_SEALS",
			},
		},
	}
}

func Supports() []string {
	return []string{
		"CONDITIONAL_SUPPORT_TRIALITY_IS_DEEP_NATIVE_THREEFOLD_SOURCE_CANDIDATE",
		"CONDITIONAL_SUPPORT_CL17_ROOT_MAKES_TRIALITY_A_LAWFUL_R4_SOURCE_TO_AUDIT",
		"CONDITIONAL_SUPPORT_VECTOR_SPINOR_SPINOR_THREEFOLD_IS_STRUCTURAL_NOT_EMPIRICAL",
		"CONDITIONAL_SUPPORT_TRIALITY_MAY_SUPPLY_ACTION_PRINCIPLE_FOR_K7_MINUS_CARRIER",
		"CONDITIONAL_SUPPORT_TRIALITY_AUDIT_RESPECTS_SCALAR_SOURCE_AND_POST_ORIENTATION_SEALS",
	}
}

func Failures() []string {
	return []string{
		"FAILED_ROUTE_TRIALITY_THREEFOLD_SHAPE_NOT_GENERATION_THEOREM",
		"FAILED_ROUTE_NO_TRIALITY_TO_GENERATION_CARRIER_MAP",
		"FAILED_ROUTE_NO_TRIALITY_ACTION_ON_R3_TRACEBODY",
		"FAILED_ROUTE_NO_TRIALITY_TO_YUKAWA_SOCKET_LEDGER_FUNCTOR",
		"FAILED_ROUTE_NO_TRIALITY_K7_MINUS_COUPLING_THEOREM",
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
