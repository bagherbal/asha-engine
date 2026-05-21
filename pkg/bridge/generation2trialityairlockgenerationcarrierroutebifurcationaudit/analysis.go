// Package generation2trialityairlockgenerationcarrierroutebifurcationaudit implements
// Gate 957: Triality Airlock and GenerationCarrier Route Bifurcation Audit.
//
// Gate 957 follows Gate 956's hard blocker: the K7^-/triality route is not
// certified because triality has not been installed as a native operator on the
// active ASHA Lambda^4/K7 board. This gate does not try another abstract C3
// model and does not enter flavor. It audits whether the missing object is a
// TrialityAirlock from the D4/Spin(8) parent triality layer into Lambda^4 R^8
// and K7, then decides whether the K7^- route is repaired, rerouted, or blocked.
package generation2trialityairlockgenerationcarrierroutebifurcationaudit

import (
	"fmt"
	"strings"
)

const (
	AuditID         = "GATE957-GENERATION2TRIALITYAIRLOCKGENERATIONCARRIERROUTEBIFURCATIONAUDIT"
	InheritedStatus = "R4_K7_MINUS_TRIALITY_ROUTE_BLOCKED_IN_CURRENT_CERTIFICATE"
	Verdict         = "NO_TRIALITY_AIRLOCK_TO_ACTIVE_ASHA_CHAMBER_CERTIFIED_ALTERNATIVE_GENERATION_CARRIER_SEARCH_REQUIRED"
	Classification  = "R4_TRIALITY_AIRLOCK_MISSING_K7_MINUS_ROUTE_BLOCKED_SEARCH_ALTERNATIVE_CARRIER"
	ShortStatus     = "R4_TRIALITY_AIRLOCK_MISSING_ROUTE_BIFURCATES_TO_ALTERNATIVE_SEARCH"
	NextGate        = "NEXT_GATE958_ALTERNATIVE_GENERATION_CARRIER_SEARCH_AUDIT"
)

const (
	TrialityParentBoardName = "T_parent = 8_v plus 8_s_plus plus 8_s_minus"
	TrialityAirlockName     = "A_tri_to_Lambda4 or A_tri_to_K7"
	K7Dim                   = 7
	K7PlusDim               = 4
	K7MinusDim              = 3
	R3TraceRows             = "3,3,1"
)

type AuditItem struct {
	Name      string
	Status    string
	Supports  []string
	Firewalls []string
}

type ParentTrialityBoard struct {
	Identified                   bool
	HasVectorLane                bool
	HasLeftSpinorLane            bool
	HasRightSpinorLane           bool
	HasOrderThreePermutation     bool
	NativeD4Spin8SourceCertified bool
	ReplacesGate955AbstractC3    bool
}

type TrialityAirlock struct {
	ToLambda4Certified                bool
	ToK7Certified                     bool
	UsesVectorExteriorAction          bool
	UsesSpinorBilinearFierzMap        bool
	UsesOctonionicCalibration         bool
	UsesSpin8InvariantTensor          bool
	CompatibleWithBooleanProjector    bool
	CompatibleWithOctonionicProjector bool
	PreservesK7                       bool
	PreservesK7Minus                  bool
	SelectsAlternativeThreeCarrier    bool
	SelectsNoCanonicalThreeCarrier    bool
}

type RouteDecision struct {
	Decision                    string
	K7MinusRouteReopened        bool
	AlternativeCarrierSelected  bool
	AlternativeSearchRequired   bool
	GenerationCarrierCertified  bool
	FlavorOrientationCertified  bool
	IndividualYukawaCertified   bool
	PhysicalAssignmentCertified bool
	OfficialLedgerUpdate        bool
}

type Analysis struct {
	AuditID                    string
	Inherited                  string
	Verdict                    string
	Classification             string
	ShortStatus                string
	NativeR3                   bool
	R3DualSealRequired         bool
	GenerationCarrierCertified bool
	K7MinusDimension           int
	K7PlusDimension            int
	R3TraceRows                string
	Parent                     ParentTrialityBoard
	Airlock                    TrialityAirlock
	Decision                   RouteDecision
	Items                      []AuditItem
	Supports                   []string
	Failures                   []string
	Final                      string
}

func BuildDefault() (Analysis, error) {
	parent := ParentTrialityBoard{
		Identified:                   false,
		HasVectorLane:                false,
		HasLeftSpinorLane:            false,
		HasRightSpinorLane:           false,
		HasOrderThreePermutation:     false,
		NativeD4Spin8SourceCertified: false,
		ReplacesGate955AbstractC3:    false,
	}
	airlock := TrialityAirlock{
		ToLambda4Certified:                false,
		ToK7Certified:                     false,
		UsesVectorExteriorAction:          false,
		UsesSpinorBilinearFierzMap:        false,
		UsesOctonionicCalibration:         false,
		UsesSpin8InvariantTensor:          false,
		CompatibleWithBooleanProjector:    false,
		CompatibleWithOctonionicProjector: false,
		PreservesK7:                       false,
		PreservesK7Minus:                  false,
		SelectsAlternativeThreeCarrier:    false,
		SelectsNoCanonicalThreeCarrier:    true,
	}
	decision := RouteDecision{
		Decision:                    Verdict,
		K7MinusRouteReopened:        false,
		AlternativeCarrierSelected:  false,
		AlternativeSearchRequired:   true,
		GenerationCarrierCertified:  false,
		FlavorOrientationCertified:  false,
		IndividualYukawaCertified:   false,
		PhysicalAssignmentCertified: false,
		OfficialLedgerUpdate:        false,
	}
	items := DefaultItems(parent, airlock, decision)
	if len(items) != 7 {
		return Analysis{}, fmt.Errorf("expected 7 audit items, got %d", len(items))
	}
	a := Analysis{
		AuditID:                    AuditID,
		Inherited:                  InheritedStatus,
		Verdict:                    Verdict,
		Classification:             Classification,
		ShortStatus:                ShortStatus,
		NativeR3:                   false,
		R3DualSealRequired:         true,
		GenerationCarrierCertified: false,
		K7MinusDimension:           K7MinusDim,
		K7PlusDimension:            K7PlusDim,
		R3TraceRows:                R3TraceRows,
		Parent:                     parent,
		Airlock:                    airlock,
		Decision:                   decision,
		Items:                      items,
		Supports:                   Supports(),
		Failures:                   Failures(),
		Final:                      "Gate 957 bifurcates the K7^-/triality route after Gate 956. It identifies TrialityAirlock / NativeTrialityTransportCertificate as the missing object, distinguishes parent D4/Spin(8) triality from an abstract C3 model, and audits the missing airlock into Lambda^4/K7/K7^-. No airlock is certified in the current board, so the K7^- route remains blocked and the next lawful move is an alternative generation-carrier search rather than flavor or R3-row labeling.",
	}
	if err := validate(a); err != nil {
		return Analysis{}, err
	}
	return a, nil
}

func validate(a Analysis) error {
	if a.NativeR3 || !a.R3DualSealRequired {
		return fmt.Errorf("Gate 957 must preserve R3 dual seal and avoid native R3 overclaim")
	}
	if a.GenerationCarrierCertified || a.Decision.GenerationCarrierCertified || a.Decision.FlavorOrientationCertified || a.Decision.IndividualYukawaCertified || a.Decision.PhysicalAssignmentCertified || a.Decision.OfficialLedgerUpdate {
		return fmt.Errorf("Gate 957 overclaimed generation/flavor/Yukawa/official status")
	}
	if a.K7MinusDimension != 3 || a.K7PlusDimension != 4 || a.K7MinusDimension+a.K7PlusDimension != K7Dim {
		return fmt.Errorf("bad K7 polarity dimensions")
	}
	if a.R3TraceRows != R3TraceRows {
		return fmt.Errorf("R3 trace rows must remain aggregate %s", R3TraceRows)
	}
	if a.Parent.Identified || a.Parent.NativeD4Spin8SourceCertified || a.Parent.ReplacesGate955AbstractC3 || a.Parent.HasOrderThreePermutation {
		return fmt.Errorf("parent triality board overcertified: %#v", a.Parent)
	}
	if a.Airlock.ToLambda4Certified || a.Airlock.ToK7Certified || a.Airlock.PreservesK7 || a.Airlock.PreservesK7Minus || a.Airlock.SelectsAlternativeThreeCarrier {
		return fmt.Errorf("triality airlock overcertified: %#v", a.Airlock)
	}
	if !a.Airlock.SelectsNoCanonicalThreeCarrier {
		return fmt.Errorf("absence of canonical three-carrier should be recorded")
	}
	if a.Decision.K7MinusRouteReopened || a.Decision.AlternativeCarrierSelected || !a.Decision.AlternativeSearchRequired {
		return fmt.Errorf("route bifurcation decision inconsistent: %#v", a.Decision)
	}
	if a.Decision.Decision != Verdict {
		return fmt.Errorf("decision does not match verdict")
	}
	return nil
}

func DefaultItems(p ParentTrialityBoard, air TrialityAirlock, d RouteDecision) []AuditItem {
	return []AuditItem{
		{
			Name:   "parent triality board reconstruction",
			Status: statusIf(p.Identified && p.NativeD4Spin8SourceCertified, "NATIVE_TRIALITY_PARENT_BOARD_IDENTIFIED", "BLOCKED_NO_NATIVE_TRIALITY_PARENT_BOARD_IN_ACTIVE_CERTIFICATE"),
			Supports: []string{
				"CONDITIONAL_SUPPORT_TRIALITY_PARENT_BOARD_IS_CORRECT_SOURCE_LAYER_TO_AUDIT",
				"CONDITIONAL_SUPPORT_TRIALITY_LIVES_AT_D4_SPIN8_PARENT_LAYER_NOT_AUTOMATIC_LAMBDA4_ENDOMORPHISM",
			},
			Firewalls: []string{
				"FAILED_ROUTE_NO_NATIVE_TRIALITY_PARENT_BOARD_IN_ACTIVE_CERTIFICATE",
				"FAILED_ROUTE_GATE955_C3_ACTION_REMAINS_ABSTRACT_UNLESS_NATIVE_AIRLOCK_IS_CERTIFIED",
			},
		},
		{
			Name:   "triality-to-Lambda4 airlock",
			Status: statusIf(air.ToLambda4Certified, "TRIALITY_AIRLOCK_TO_LAMBDA4_EXISTS", "BLOCKED_NO_TRIALITY_TO_LAMBDA4_AIRLOCK"),
			Supports: []string{
				"CONDITIONAL_SUPPORT_TRIALITY_AIRLOCK_TO_LAMBDA4_IS_THE_REQUIRED_TRANSPORT_CERTIFICATE",
				"CONDITIONAL_SUPPORT_VECTOR_SPINOR_SPINOR_PARENT_LAYER_MUST_BE_TYPED_TO_ACTIVE_LAMBDA4_CHAMBER",
			},
			Firewalls: []string{
				"FAILED_ROUTE_NO_TRIALITY_TO_LAMBDA4_AIRLOCK",
				"FAILED_ROUTE_TRIALITY_SOURCE_LAYER_NOT_TRANSPORTED_TO_ACTIVE_LAMBDA4_CHAMBER",
			},
		},
		{
			Name:   "Boolean-octonionic K7 compatibility",
			Status: statusIf(air.PreservesK7 && air.CompatibleWithBooleanProjector && air.CompatibleWithOctonionicProjector, "TRIALITY_AIRLOCK_PRESERVES_BOOLEAN_OCTONIONIC_CONTACT_CARRIER", "BLOCKED_NO_TRIALITY_AIRLOCK_PRESERVATION_OF_K7"),
			Supports: []string{
				"CONDITIONAL_SUPPORT_K7_CONTACT_CARRIER_IS_REQUIRED_AFTER_LAMBDA4_AIRLOCK",
			},
			Firewalls: []string{
				"FAILED_ROUTE_TRIALITY_AIRLOCK_BREAKS_OR_DOES_NOT_CERTIFY_BOOLEAN_PROJECTOR_COMPATIBILITY",
				"FAILED_ROUTE_TRIALITY_AIRLOCK_BREAKS_OR_DOES_NOT_CERTIFY_OCTONIONIC_PROJECTOR_COMPATIBILITY",
				"FAILED_ROUTE_TRIALITY_AIRLOCK_DOES_NOT_PRESERVE_K7",
			},
		},
		{
			Name:   "Hodge polarity and K7-minus route",
			Status: hodgeRouteStatus(air),
			Supports: []string{
				"CONDITIONAL_SUPPORT_K7_MINUS_REMAINS_STRONG_CARRIER_SHAPE_IF_NATIVE_AIRLOCK_CAN_REACH_IT",
				"CONDITIONAL_SUPPORT_ROUTE_BIFURCATION_DISTINGUISHES_K7_MINUS_REPAIR_FROM_ALTERNATIVE_CARRIER_SELECTION",
			},
			Firewalls: []string{
				"FAILED_ROUTE_TRIALITY_AIRLOCK_DOES_NOT_SELECT_NATIVE_GENERATION_CARRIER",
				"FAILED_ROUTE_NATIVE_TRIALITY_RESTRICTION_TO_K7_MINUS_NOT_CERTIFIED",
			},
		},
		{
			Name:   "route bifurcation decision",
			Status: d.Decision,
			Supports: []string{
				"CONDITIONAL_SUPPORT_K7_MINUS_TRIALITY_FAILURE_IS_ARCHITECTURAL_UNTIL_TRIALITY_AIRLOCK_IS_TESTED",
				"CONDITIONAL_SUPPORT_ALTERNATIVE_GENERATION_CARRIER_SEARCH_IS_REQUIRED_IF_AIRLOCK_ABSENT",
			},
			Firewalls: []string{
				"FAILED_ROUTE_NO_GENERATION_CARRIER_MAP",
				"FAILED_ROUTE_NO_TRIALITY_AIRLOCK_TO_ACTIVE_ASHA_CHAMBER_CERTIFIED",
			},
		},
		{
			Name:   "R3 dual-seal and trace-row firewall",
			Status: "R3_DUALSEAL_AND_TRACE_ROW_FIREWALL_PRESERVED",
			Supports: []string{
				"CONDITIONAL_SUPPORT_R3_TRACEBODY_CAN_REMAIN_ONLY_DUALSEALED_AGGREGATE_TARGET",
			},
			Firewalls: []string{
				"FAILED_ROUTE_R3_TRACE_ROWS_ARE_NOT_GENERATION_LABELS",
				"FAILED_ROUTE_R3_TRACEBRIDGE_REMAINS_DUALSEALED_NOT_NATIVE",
				"FAILED_ROUTE_R3_TRACEBRIDGE_REMAINS_SCALAR_SOURCE_SEALED",
				"FAILED_ROUTE_R3_TRACEBRIDGE_REMAINS_POST_ORIENTATION_SEALED",
			},
		},
		{
			Name:   "flavor and empirical firewall",
			Status: "FLAVOR_AND_EMPIRICAL_FIREWALL_PRESERVED",
			Supports: []string{
				"CONDITIONAL_SUPPORT_GATE957_DOES_NOT_USE_FLAVOR_BACKSOLVE_OR_OBSERVED_YUKAWA_INPUT",
			},
			Firewalls: []string{
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

func hodgeRouteStatus(air TrialityAirlock) string {
	if air.PreservesK7Minus {
		return "TRIALITY_AIRLOCK_RESTRICTS_TO_K7_MINUS_ROUTE_REOPENED"
	}
	if air.SelectsAlternativeThreeCarrier {
		return "TRIALITY_AIRLOCK_SELECTS_ALTERNATIVE_NATIVE_THREE_CARRIER"
	}
	return "NO_CANONICAL_THREE_CARRIER_SELECTED_BY_TRIALITY_AIRLOCK"
}

func Supports() []string {
	return []string{
		"CONDITIONAL_SUPPORT_TRIALITY_PARENT_BOARD_IS_CORRECT_SOURCE_LAYER_TO_AUDIT",
		"CONDITIONAL_SUPPORT_TRIALITY_LIVES_AT_D4_SPIN8_PARENT_LAYER_NOT_AUTOMATIC_LAMBDA4_ENDOMORPHISM",
		"CONDITIONAL_SUPPORT_TRIALITY_AIRLOCK_TO_LAMBDA4_IS_THE_REQUIRED_TRANSPORT_CERTIFICATE",
		"CONDITIONAL_SUPPORT_VECTOR_SPINOR_SPINOR_PARENT_LAYER_MUST_BE_TYPED_TO_ACTIVE_LAMBDA4_CHAMBER",
		"CONDITIONAL_SUPPORT_K7_CONTACT_CARRIER_IS_REQUIRED_AFTER_LAMBDA4_AIRLOCK",
		"CONDITIONAL_SUPPORT_K7_MINUS_REMAINS_STRONG_CARRIER_SHAPE_IF_NATIVE_AIRLOCK_CAN_REACH_IT",
		"CONDITIONAL_SUPPORT_ROUTE_BIFURCATION_DISTINGUISHES_K7_MINUS_REPAIR_FROM_ALTERNATIVE_CARRIER_SELECTION",
		"CONDITIONAL_SUPPORT_K7_MINUS_TRIALITY_FAILURE_IS_ARCHITECTURAL_UNTIL_TRIALITY_AIRLOCK_IS_TESTED",
		"CONDITIONAL_SUPPORT_ALTERNATIVE_GENERATION_CARRIER_SEARCH_IS_REQUIRED_IF_AIRLOCK_ABSENT",
		"CONDITIONAL_SUPPORT_R3_TRACEBODY_CAN_REMAIN_ONLY_DUALSEALED_AGGREGATE_TARGET",
		"CONDITIONAL_SUPPORT_GATE957_DOES_NOT_USE_FLAVOR_BACKSOLVE_OR_OBSERVED_YUKAWA_INPUT",
	}
}

func Failures() []string {
	return []string{
		"FAILED_ROUTE_NO_GENERATION_CARRIER_MAP",
		"FAILED_ROUTE_NO_NATIVE_TRIALITY_PARENT_BOARD_IN_ACTIVE_CERTIFICATE",
		"FAILED_ROUTE_GATE955_C3_ACTION_REMAINS_ABSTRACT_UNLESS_NATIVE_AIRLOCK_IS_CERTIFIED",
		"FAILED_ROUTE_ABSTRACT_C3_ACTION_CANNOT_BE_PROMOTED_BY_BASIS_FIT",
		"FAILED_ROUTE_NO_TRIALITY_TO_LAMBDA4_AIRLOCK",
		"FAILED_ROUTE_TRIALITY_SOURCE_LAYER_NOT_TRANSPORTED_TO_ACTIVE_LAMBDA4_CHAMBER",
		"FAILED_ROUTE_TRIALITY_AIRLOCK_BREAKS_OR_DOES_NOT_CERTIFY_BOOLEAN_PROJECTOR_COMPATIBILITY",
		"FAILED_ROUTE_TRIALITY_AIRLOCK_BREAKS_OR_DOES_NOT_CERTIFY_OCTONIONIC_PROJECTOR_COMPATIBILITY",
		"FAILED_ROUTE_TRIALITY_AIRLOCK_DOES_NOT_PRESERVE_K7",
		"FAILED_ROUTE_TRIALITY_AIRLOCK_DOES_NOT_SELECT_NATIVE_GENERATION_CARRIER",
		"FAILED_ROUTE_NATIVE_TRIALITY_RESTRICTION_TO_K7_MINUS_NOT_CERTIFIED",
		"FAILED_ROUTE_NO_TRIALITY_AIRLOCK_TO_ACTIVE_ASHA_CHAMBER_CERTIFIED",
		"FAILED_ROUTE_R3_TRACE_ROWS_ARE_NOT_GENERATION_LABELS",
		"FAILED_ROUTE_R3_TRACEBRIDGE_REMAINS_DUALSEALED_NOT_NATIVE",
		"FAILED_ROUTE_R3_TRACEBRIDGE_REMAINS_SCALAR_SOURCE_SEALED",
		"FAILED_ROUTE_R3_TRACEBRIDGE_REMAINS_POST_ORIENTATION_SEALED",
		"FAILED_ROUTE_NO_FLAVOR_ORIENTATION_MAP",
		"FAILED_ROUTE_NO_INDIVIDUAL_YUKAWA_VALUES",
		"FAILED_ROUTE_NO_PHYSICAL_PARTICLE_ASSIGNMENT",
		"FAILED_ROUTE_NO_CKM_PMNS_THEOREM",
		"FAILED_ROUTE_NO_R4_NATIVE_YUKAWA_SPECTRUM_THEOREM",
		"FAILED_ROUTE_NO_OFFICIAL_LEDGER_UPDATE_ALLOWED",
	}
}

func RequiredSupports() []string {
	return []string{
		"CONDITIONAL_SUPPORT_TRIALITY_PARENT_BOARD_IS_CORRECT_SOURCE_LAYER_TO_AUDIT",
		"CONDITIONAL_SUPPORT_TRIALITY_AIRLOCK_TO_LAMBDA4_IS_THE_REQUIRED_TRANSPORT_CERTIFICATE",
		"CONDITIONAL_SUPPORT_K7_CONTACT_CARRIER_IS_REQUIRED_AFTER_LAMBDA4_AIRLOCK",
		"CONDITIONAL_SUPPORT_ROUTE_BIFURCATION_DISTINGUISHES_K7_MINUS_REPAIR_FROM_ALTERNATIVE_CARRIER_SELECTION",
		"CONDITIONAL_SUPPORT_ALTERNATIVE_GENERATION_CARRIER_SEARCH_IS_REQUIRED_IF_AIRLOCK_ABSENT",
	}
}

func RequiredFailures() []string {
	return []string{
		"FAILED_ROUTE_NO_GENERATION_CARRIER_MAP",
		"FAILED_ROUTE_NO_NATIVE_TRIALITY_PARENT_BOARD_IN_ACTIVE_CERTIFICATE",
		"FAILED_ROUTE_GATE955_C3_ACTION_REMAINS_ABSTRACT_UNLESS_NATIVE_AIRLOCK_IS_CERTIFIED",
		"FAILED_ROUTE_NO_TRIALITY_TO_LAMBDA4_AIRLOCK",
		"FAILED_ROUTE_TRIALITY_AIRLOCK_DOES_NOT_PRESERVE_K7",
		"FAILED_ROUTE_NATIVE_TRIALITY_RESTRICTION_TO_K7_MINUS_NOT_CERTIFIED",
		"FAILED_ROUTE_NO_TRIALITY_AIRLOCK_TO_ACTIVE_ASHA_CHAMBER_CERTIFIED",
		"FAILED_ROUTE_R3_TRACE_ROWS_ARE_NOT_GENERATION_LABELS",
		"FAILED_ROUTE_NO_FLAVOR_ORIENTATION_MAP",
	}
}

func statusIf(ok bool, yes, no string) string {
	if ok {
		return yes
	}
	return no
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

func ItemNotes(items []AuditItem) []string {
	out := make([]string, 0, len(items))
	for _, it := range items {
		out = append(out, it.Name+" => "+it.Status)
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
