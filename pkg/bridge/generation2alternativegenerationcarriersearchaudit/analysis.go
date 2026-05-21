// Package generation2alternativegenerationcarriersearchaudit implements
// Gate 958: Alternative GenerationCarrier Search Audit.
//
// Gate 958 follows Gate 957's route bifurcation: the K7^-/triality route is
// suspended in the active certificate because no TrialityAirlock to the
// Lambda^4/K7 chamber is certified. This gate searches the already-active ASHA
// board for other generation-carrier candidates. It is deliberately strict:
// a candidate needs carrier shape, native action/selector, and a typed map to
// the dual-sealed R3 tracebody. Rank-three coincidences are not enough.
package generation2alternativegenerationcarriersearchaudit

import (
	"fmt"
	"strings"
)

const (
	AuditID         = "GATE958-GENERATION2ALTERNATIVEGENERATIONCARRIERSEARCHAUDIT"
	InheritedStatus = "R4_TRIALITY_AIRLOCK_MISSING_K7_MINUS_ROUTE_BLOCKED_SEARCH_ALTERNATIVE_CARRIER"
	Verdict         = "NO_ALTERNATIVE_NATIVE_GENERATION_CARRIER_FOUND_EXTERNAL_C3_SEAL_OR_NEW_PARENT_AIRLOCK_REQUIRED"
	Classification  = "R4_ALTERNATIVE_GENERATION_CARRIER_SEARCH_FAILED_NO_NATIVE_CARRIER_MAP"
	ShortStatus     = "R4_GENERATION_CARRIER_REMAINS_MISSING"
	NextGate        = "NEXT_GATE959_GENERATION_CARRIER_SEAL_OR_PARENT_AIRLOCK_DECISION_AUDIT"
)

const (
	K7PlusDim       = 4
	K7MinusDim      = 3
	K7Dim           = 7
	BoundaryPairDim = 2
	R3TraceRows     = "3,3,1"
)

type CandidateStatus string

const (
	CandidateRejected      CandidateStatus = "REJECTED_NO_GENERATION_CARRIER_MAP"
	CandidateShapeOnly     CandidateStatus = "SHAPE_ONLY_NO_ACTION_OR_TRACEBODY_MAP"
	CandidateSealOnly      CandidateStatus = "SEAL_ONLY_NOT_NATIVE"
	CandidateNoThree       CandidateStatus = "NO_NATIVE_THREE_CARRIER_FOUND"
	CandidateStronglyValid CandidateStatus = "NATIVE_CARRIER_MAP_CERTIFIED"
)

type CandidateAudit struct {
	Name                   string
	Candidate              string
	Status                 CandidateStatus
	CarrierShape           bool
	NativeActionOrSelector bool
	TypedTracebodyMap      bool
	NonCircular            bool
	UsesFlavorBacksolve    bool
	UsesR3RowsAsLabels     bool
	AllowsOnlyAsSeal       bool
	Supports               []string
	Firewalls              []string
}

type SearchDecision struct {
	Decision                    string
	NativeGenerationCarrier     bool
	AlternativeCandidateFound   bool
	ExternalSealAvailable       bool
	ParentAirlockRequired       bool
	FlavorOrientationCertified  bool
	IndividualYukawaCertified   bool
	PhysicalAssignmentCertified bool
	OfficialLedgerUpdateAllowed bool
}

type Analysis struct {
	AuditID                    string
	Inherited                  string
	Verdict                    string
	Classification             string
	ShortStatus                string
	NextGate                   string
	R3DualSealRequired         bool
	NativeR3                   bool
	GenerationCarrierCertified bool
	K7PlusDimension            int
	K7MinusDimension           int
	BoundaryPairDimension      int
	R3TraceRows                string
	Candidates                 []CandidateAudit
	Decision                   SearchDecision
	Supports                   []string
	Failures                   []string
	Final                      string
}

func BuildDefault() (Analysis, error) {
	candidates := DefaultCandidates()
	if len(candidates) != 5 {
		return Analysis{}, fmt.Errorf("expected 5 candidate classes, got %d", len(candidates))
	}
	decision := SearchDecision{
		Decision:                    Verdict,
		NativeGenerationCarrier:     false,
		AlternativeCandidateFound:   false,
		ExternalSealAvailable:       true,
		ParentAirlockRequired:       true,
		FlavorOrientationCertified:  false,
		IndividualYukawaCertified:   false,
		PhysicalAssignmentCertified: false,
		OfficialLedgerUpdateAllowed: false,
	}
	a := Analysis{
		AuditID:                    AuditID,
		Inherited:                  InheritedStatus,
		Verdict:                    Verdict,
		Classification:             Classification,
		ShortStatus:                ShortStatus,
		NextGate:                   NextGate,
		R3DualSealRequired:         true,
		NativeR3:                   false,
		GenerationCarrierCertified: false,
		K7PlusDimension:            K7PlusDim,
		K7MinusDimension:           K7MinusDim,
		BoundaryPairDimension:      BoundaryPairDim,
		R3TraceRows:                R3TraceRows,
		Candidates:                 candidates,
		Decision:                   decision,
		Supports:                   Supports(),
		Failures:                   Failures(),
		Final:                      "Gate 958 searches the active ASHA board after the K7^-/triality route bifurcates. It audits K7 Hodge polarity without triality, the Fock/projective 1+3 selector, B2 boundary response, Boolean-octonionic projector ranks, and external C3 family seal. No candidate supplies all three required objects: native carrier, native action/selector, and typed R3 tracebody map. The generation carrier remains missing; the next lawful choice is either an explicit external family seal or a new parent airlock/certificate, not flavor backsolve.",
	}
	if err := validate(a); err != nil {
		return Analysis{}, err
	}
	return a, nil
}

func validate(a Analysis) error {
	if a.NativeR3 || !a.R3DualSealRequired || a.GenerationCarrierCertified || a.Decision.NativeGenerationCarrier || a.Decision.AlternativeCandidateFound {
		return fmt.Errorf("Gate 958 must not certify native R3 or generation carrier: %#v", a.Decision)
	}
	if a.Decision.FlavorOrientationCertified || a.Decision.IndividualYukawaCertified || a.Decision.PhysicalAssignmentCertified || a.Decision.OfficialLedgerUpdateAllowed {
		return fmt.Errorf("Gate 958 overclaimed flavor/Yukawa/physical/official status")
	}
	if !a.Decision.ExternalSealAvailable || !a.Decision.ParentAirlockRequired {
		return fmt.Errorf("Gate 958 must leave external seal or parent airlock as remaining routes")
	}
	if a.K7PlusDimension != 4 || a.K7MinusDimension != 3 || a.K7PlusDimension+a.K7MinusDimension != K7Dim {
		return fmt.Errorf("bad K7 dimensions")
	}
	if a.BoundaryPairDimension != 2 {
		return fmt.Errorf("B2 must remain rank two")
	}
	if a.R3TraceRows != R3TraceRows {
		return fmt.Errorf("R3 trace rows must remain aggregate %s", R3TraceRows)
	}
	if len(a.Candidates) != 5 {
		return fmt.Errorf("expected 5 candidates, got %d", len(a.Candidates))
	}
	for _, c := range a.Candidates {
		if c.Status == CandidateStronglyValid || c.TypedTracebodyMap || c.UsesFlavorBacksolve || c.UsesR3RowsAsLabels {
			return fmt.Errorf("candidate overclaimed or used forbidden route: %#v", c)
		}
	}
	return nil
}

func DefaultCandidates() []CandidateAudit {
	return []CandidateAudit{
		{
			Name:                   "K7 Hodge polarity without triality",
			Candidate:              "K7 = K7^+ plus K7^- with dim(K7^-)=3",
			Status:                 CandidateShapeOnly,
			CarrierShape:           true,
			NativeActionOrSelector: false,
			TypedTracebodyMap:      false,
			NonCircular:            true,
			Supports: []string{
				"CONDITIONAL_SUPPORT_K7_MINUS_REMAINS_NATIVE_DIMENSION_THREE_CARRIER_SHAPE",
				"CONDITIONAL_SUPPORT_K7_HODGE_POLARITY_REMAINS_NONEMPIRICAL_ACTIVE_BOARD_SOURCE",
			},
			Firewalls: []string{
				"FAILED_ROUTE_NO_NATIVE_ACTION_OR_TRACEBODY_MAP_FROM_HODGE_POLARITY_ALONE",
				"FAILED_ROUTE_DIM_K7_MINUS_EQUALS_THREE_NOT_GENERATION_THEOREM",
				"FAILED_ROUTE_NO_K7_POLARITY_TO_GENERATION_CARRIER_MAP",
			},
		},
		{
			Name:                   "Fock/projective 1+3 selector",
			Candidate:              "P1 plus P3 and B-L selector",
			Status:                 CandidateRejected,
			CarrierShape:           true,
			NativeActionOrSelector: true,
			TypedTracebodyMap:      false,
			NonCircular:            true,
			Supports: []string{
				"CONDITIONAL_SUPPORT_FOCK_PROJECTIVE_SELECTOR_IS_MATURE_SOCKET_INTERNAL_CHARGE_STRUCTURE",
			},
			Firewalls: []string{
				"FAILED_ROUTE_FOCK_P3_RANK_THREE_NOT_GENERATION_CARRIER",
				"FAILED_ROUTE_B_MINUS_L_SELECTOR_NOT_GENERATION_SELECTOR",
				"FAILED_ROUTE_P3_MULTIPLICITY_NOT_FAMILY_MULTIPLICITY",
			},
		},
		{
			Name:                   "Boundary pair B2 plus rank-three response",
			Candidate:              "B2 boundary response and R3 trace rows",
			Status:                 CandidateRejected,
			CarrierShape:           false,
			NativeActionOrSelector: true,
			TypedTracebodyMap:      false,
			NonCircular:            true,
			Supports: []string{
				"CONDITIONAL_SUPPORT_BOUNDARY_RESPONSE_REMAINS_VALID_R3_TRACEBRIDGE_INPUT_UNDER_DUALSEAL",
			},
			Firewalls: []string{
				"FAILED_ROUTE_BOUNDARY_PAIR_B2_NOT_THREE_GENERATION_CARRIER",
				"FAILED_ROUTE_R3_TRACE_RESPONSE_NOT_GENERATION_LABEL",
				"FAILED_ROUTE_R3_TRACE_ROWS_ARE_NOT_GENERATION_LABELS",
			},
		},
		{
			Name:                   "Boolean-octonionic projector search",
			Candidate:              "P_B, P_G, K7, K7 complement, W7, V0, U0",
			Status:                 CandidateNoThree,
			CarrierShape:           false,
			NativeActionOrSelector: false,
			TypedTracebodyMap:      false,
			NonCircular:            true,
			Supports: []string{
				"CONDITIONAL_SUPPORT_BOOLEAN_OCTONIONIC_SYSTEM_IS_VALID_NATIVE_SEARCH_SPACE",
				"CONDITIONAL_SUPPORT_ACTIVE_BOARD_SEARCH_AUDITED_BEYOND_K7_MINUS_TRIALITY_ROUTE",
			},
			Firewalls: []string{
				"FAILED_ROUTE_NO_CANONICAL_ALTERNATIVE_THREE_CARRIER_FOUND",
				"FAILED_ROUTE_BOOLEAN_OCTONIONIC_RANKS_56_14_7_DO_NOT_BY_THEMSELVES_YIELD_GENERATION_CARRIER",
				"FAILED_ROUTE_NO_TYPED_MAP_FROM_BOOLEAN_OCTONIONIC_ALTERNATIVE_TO_R3_TRACEBODY",
			},
		},
		{
			Name:                   "external C3 family factor seal",
			Candidate:              "C3_family as external generation carrier seal",
			Status:                 CandidateSealOnly,
			CarrierShape:           true,
			NativeActionOrSelector: false,
			TypedTracebodyMap:      false,
			NonCircular:            true,
			AllowsOnlyAsSeal:       true,
			Supports: []string{
				"CONDITIONAL_SUPPORT_EXTERNAL_C3_CAN_CLOSE_MODEL_AS_SEAL",
				"CONDITIONAL_SUPPORT_EXTERNAL_GENERATION_CARRIER_SEAL_REMAINS_AVAILABLE_IF_NATIVE_SEARCH_FAILS",
			},
			Firewalls: []string{
				"FAILED_ROUTE_EXTERNAL_C3_NOT_NATIVE_GENERATION_CARRIER",
				"FAILED_ROUTE_INSERTED_FAMILY_MULTIPLICITY_NOT_R4_THEOREM",
			},
		},
	}
}

func Supports() []string {
	return []string{
		"CONDITIONAL_SUPPORT_R4_MAY_PROCEED_UNDER_EXPLICIT_R3_DUALSEAL",
		"CONDITIONAL_SUPPORT_R3_TRACEBODY_CAN_BE_USED_ONLY_AS_AGGREGATE_INPUT",
		"CONDITIONAL_SUPPORT_K7_MINUS_REMAINS_NATIVE_DIMENSION_THREE_CARRIER_SHAPE",
		"CONDITIONAL_SUPPORT_BOOLEAN_OCTONIONIC_SYSTEM_IS_VALID_NATIVE_SEARCH_SPACE",
		"CONDITIONAL_SUPPORT_EXTERNAL_C3_CAN_CLOSE_MODEL_AS_SEAL",
		"CONDITIONAL_SUPPORT_NO_TESTED_ACTIVE_BOARD_CANDIDATE_SUPPLIES_CARRIER_ACTION_AND_TRACEBODY_MAP",
		"CONDITIONAL_SUPPORT_GENERATION_CARRIER_SEARCH_MUST_REQUIRE_TYPED_TRACEBODY_MAP",
	}
}

func Failures() []string {
	return []string{
		"FAILED_ROUTE_NO_GENERATION_CARRIER_MAP",
		"FAILED_ROUTE_R3_TRACE_ROWS_ARE_NOT_GENERATION_CARRIER",
		"FAILED_ROUTE_COLOR_RANK_THREE_NOT_GENERATION_CARRIER",
		"FAILED_ROUTE_DIM_K7_MINUS_EQUALS_THREE_NOT_GENERATION_THEOREM",
		"FAILED_ROUTE_NO_K7_POLARITY_TO_GENERATION_CARRIER_MAP",
		"FAILED_ROUTE_FOCK_P3_RANK_THREE_NOT_GENERATION_CARRIER",
		"FAILED_ROUTE_B_MINUS_L_SELECTOR_NOT_GENERATION_SELECTOR",
		"FAILED_ROUTE_BOUNDARY_PAIR_B2_NOT_THREE_GENERATION_CARRIER",
		"FAILED_ROUTE_R3_TRACE_RESPONSE_NOT_GENERATION_LABEL",
		"FAILED_ROUTE_NO_CANONICAL_ALTERNATIVE_THREE_CARRIER_FOUND",
		"FAILED_ROUTE_EXTERNAL_C3_NOT_NATIVE_GENERATION_CARRIER",
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

func RequiredSupports() []string {
	return []string{
		"CONDITIONAL_SUPPORT_R4_MAY_PROCEED_UNDER_EXPLICIT_R3_DUALSEAL",
		"CONDITIONAL_SUPPORT_K7_MINUS_REMAINS_NATIVE_DIMENSION_THREE_CARRIER_SHAPE",
		"CONDITIONAL_SUPPORT_BOOLEAN_OCTONIONIC_SYSTEM_IS_VALID_NATIVE_SEARCH_SPACE",
		"CONDITIONAL_SUPPORT_EXTERNAL_C3_CAN_CLOSE_MODEL_AS_SEAL",
		"CONDITIONAL_SUPPORT_GENERATION_CARRIER_SEARCH_MUST_REQUIRE_TYPED_TRACEBODY_MAP",
	}
}

func RequiredFailures() []string {
	return []string{
		"FAILED_ROUTE_NO_GENERATION_CARRIER_MAP",
		"FAILED_ROUTE_R3_TRACE_ROWS_ARE_NOT_GENERATION_CARRIER",
		"FAILED_ROUTE_COLOR_RANK_THREE_NOT_GENERATION_CARRIER",
		"FAILED_ROUTE_DIM_K7_MINUS_EQUALS_THREE_NOT_GENERATION_THEOREM",
		"FAILED_ROUTE_FOCK_P3_RANK_THREE_NOT_GENERATION_CARRIER",
		"FAILED_ROUTE_BOUNDARY_PAIR_B2_NOT_THREE_GENERATION_CARRIER",
		"FAILED_ROUTE_NO_CANONICAL_ALTERNATIVE_THREE_CARRIER_FOUND",
		"FAILED_ROUTE_EXTERNAL_C3_NOT_NATIVE_GENERATION_CARRIER",
		"FAILED_ROUTE_NO_FLAVOR_ORIENTATION_MAP",
	}
}

func CandidateSupports(cands []CandidateAudit) []string {
	var out []string
	for _, c := range cands {
		out = append(out, c.Supports...)
	}
	return out
}

func CandidateFailures(cands []CandidateAudit) []string {
	var out []string
	for _, c := range cands {
		out = append(out, c.Firewalls...)
	}
	return out
}

func CandidateNotes(cands []CandidateAudit) []string {
	out := make([]string, 0, len(cands))
	for _, c := range cands {
		out = append(out, c.Name+" => "+string(c.Status))
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
