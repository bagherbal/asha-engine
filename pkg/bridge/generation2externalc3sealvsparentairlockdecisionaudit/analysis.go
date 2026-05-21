// Package generation2externalc3sealvsparentairlockdecisionaudit implements
// Gate 959: ExternalC3 Seal vs ParentAirlock Decision Audit.
//
// Gate 959 follows Gate 958's negative closure: no active-board generation
// carrier currently satisfies carrier shape + native action/selector + typed
// R3 tracebody map. This gate forces the lawful fork. Generation multiplicity
// is either sealed explicitly as an external C3 carrier, or native R4 remains
// open only through a new typed parent-board airlock. It is not a flavor gate.
package generation2externalc3sealvsparentairlockdecisionaudit

import (
	"fmt"
	"strings"
)

const (
	AuditID         = "GATE959-GENERATION2EXTERNALC3SEALVSPARENTAIRLOCKDECISIONAUDIT"
	InheritedStatus = "R4_GENERATION_CARRIER_REMAINS_MISSING"
	Verdict         = "GENERATION_MULTIPLICITY_NOT_NATIVE_IN_ACTIVE_BOARD_EXTERNAL_C3_SEAL_OR_PARENT_AIRLOCK_REQUIRED"
	Classification  = "R4_GENERATION_SOURCE_DECISION_EXTERNAL_SEAL_OR_PARENT_AIRLOCK"
	ShortStatus     = "R4_REQUIRES_EXTERNAL_SEAL_OR_NEW_PARENT_AIRLOCK"
	NextGate        = "NEXT_GATE960_PARENT_AIRLOCK_CONSTRUCTION_OR_EXTERNAL_GENERATION_CARRIER_SEAL_AUDIT"
)

type RouteStatus string

const (
	RouteExhausted         RouteStatus = "EXHAUSTED_IN_ACTIVE_CERTIFICATE"
	RouteSealAvailable     RouteStatus = "AVAILABLE_AS_EXPLICIT_SEAL_ONLY"
	RouteNativeOpenPending RouteStatus = "NATIVE_ROUTE_OPEN_BUT_AIRLOCK_PENDING"
	RouteForbiddenAsSource RouteStatus = "FORBIDDEN_AS_NATIVE_SOURCE"
	RouteCertifiedNative   RouteStatus = "CERTIFIED_NATIVE_GENERATION_CARRIER"
)

type RouteAudit struct {
	Name                 string
	Route                string
	Status               RouteStatus
	NativeCarrier        bool
	ExternalSeal         bool
	ParentAirlock        bool
	AllowedForFutureWork bool
	UsesFlavorBacksolve  bool
	UsesR3RowsAsLabels   bool
	Supports             []string
	Firewalls            []string
}

type Decision struct {
	Decision                         string
	ActiveBoardExhausted             bool
	NativeGenerationCarrierCertified bool
	ExternalC3SealAllowed            bool
	ParentAirlockOnlyNativeRoute     bool
	FlavorDerivationAllowed          bool
	IndividualYukawaAllowed          bool
	PhysicalAssignmentAllowed        bool
	OfficialLedgerUpdateAllowed      bool
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
	Decision                   Decision
	Routes                     []RouteAudit
	Supports                   []string
	Failures                   []string
	Final                      string
}

func BuildDefault() (Analysis, error) {
	routes := DefaultRoutes()
	if len(routes) != 4 {
		return Analysis{}, fmt.Errorf("expected 4 route audits, got %d", len(routes))
	}
	decision := Decision{
		Decision:                         Verdict,
		ActiveBoardExhausted:             true,
		NativeGenerationCarrierCertified: false,
		ExternalC3SealAllowed:            true,
		ParentAirlockOnlyNativeRoute:     true,
		FlavorDerivationAllowed:          false,
		IndividualYukawaAllowed:          false,
		PhysicalAssignmentAllowed:        false,
		OfficialLedgerUpdateAllowed:      false,
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
		Decision:                   decision,
		Routes:                     routes,
		Supports:                   Supports(),
		Failures:                   Failures(),
		Final:                      "Gate 959 freezes the post-Gate-958 fork. The active ASHA board has no certified native generation carrier under the strict criterion carrier + native action/selector + typed R3 tracebody map. External C3 may be used only as an explicit ExternalGenerationCarrierSeal. The only remaining native route is a new parent-board airlock, especially a typed D4/Spin(8) triality parent-to-Lambda4/K7 transport certificate. No flavor derivation, particle assignment, individual Yukawa spectrum, or official ledger update is allowed.",
	}
	if err := validate(a); err != nil {
		return Analysis{}, err
	}
	return a, nil
}

func validate(a Analysis) error {
	if !a.R3DualSealRequired || a.NativeR3 || a.GenerationCarrierCertified || a.Decision.NativeGenerationCarrierCertified {
		return fmt.Errorf("Gate 959 must preserve R3 dual seal and avoid native generation overclaim")
	}
	if !a.Decision.ActiveBoardExhausted || !a.Decision.ExternalC3SealAllowed || !a.Decision.ParentAirlockOnlyNativeRoute {
		return fmt.Errorf("Gate 959 must force external seal or parent airlock decision: %#v", a.Decision)
	}
	if a.Decision.FlavorDerivationAllowed || a.Decision.IndividualYukawaAllowed || a.Decision.PhysicalAssignmentAllowed || a.Decision.OfficialLedgerUpdateAllowed {
		return fmt.Errorf("Gate 959 overclaimed downstream flavor/Yukawa/physical/official status")
	}
	if len(a.Routes) != 4 {
		return fmt.Errorf("expected 4 routes, got %d", len(a.Routes))
	}
	for _, r := range a.Routes {
		if r.Status == RouteCertifiedNative || r.NativeCarrier || r.UsesFlavorBacksolve || r.UsesR3RowsAsLabels {
			return fmt.Errorf("route overclaimed or used forbidden source: %#v", r)
		}
	}
	return nil
}

func DefaultRoutes() []RouteAudit {
	return []RouteAudit{
		{
			Name:                 "active ASHA board native search",
			Route:                "K7^-, P3, B-L, B2, R3 rows, Boolean-octonionic complements",
			Status:               RouteExhausted,
			AllowedForFutureWork: false,
			Supports: []string{
				"CONDITIONAL_SUPPORT_ACTIVE_BOARD_SEARCH_EXHAUSTED_UNDER_CURRENT_CERTIFICATE",
				"CONDITIONAL_SUPPORT_NO_RETURN_TO_BEAUTIFUL_THREES_WITHOUT_NEW_MAP",
			},
			Firewalls: []string{
				"FAILED_ROUTE_ACTIVE_BOARD_NATIVE_GENERATION_CARRIER_EXHAUSTED_IN_CURRENT_CERTIFICATE",
				"FAILED_ROUTE_RANK_THREE_OBJECTS_CANNOT_BE_REUSED_AS_GENERATION_CARRIERS_WITHOUT_NEW_TYPED_MAP",
			},
		},
		{
			Name:                 "ExternalC3 generation carrier seal",
			Route:                "G_ext = C^3 as explicit family multiplicity carrier seal",
			Status:               RouteSealAvailable,
			ExternalSeal:         true,
			AllowedForFutureWork: true,
			Supports: []string{
				"CONDITIONAL_SUPPORT_EXTERNAL_C3_GENERATION_CARRIER_CAN_BE_USED_AS_EXPLICIT_SEAL",
				"CONDITIONAL_SUPPORT_EXTERNAL_GENERATION_CARRIER_SEAL_CAN_CLOSE_MODEL_ONLY_AS_QUARANTINED_INPUT",
			},
			Firewalls: []string{
				"FAILED_ROUTE_EXTERNAL_C3_NOT_NATIVE_GENERATION_CARRIER",
				"FAILED_ROUTE_EXTERNAL_C3_SEAL_DOES_NOT_DERIVE_GENERATION_MULTIPLICITY",
			},
		},
		{
			Name:                 "parent-board airlock native route",
			Route:                "T_parent = 8_v plus 8_s_plus plus 8_s_minus to End(Lambda4 R8) or End(K7)",
			Status:               RouteNativeOpenPending,
			ParentAirlock:        true,
			AllowedForFutureWork: true,
			Supports: []string{
				"CONDITIONAL_SUPPORT_PARENT_AIRLOCK_IS_ONLY_REMAINING_NATIVE_R4_ROUTE",
				"CONDITIONAL_SUPPORT_D4_SPIN8_TRIALITY_PARENT_LAYER_REMAINS_PRIMARY_PARENT_BOARD_CANDIDATE",
			},
			Firewalls: []string{
				"FAILED_ROUTE_NO_PARENT_AIRLOCK_CERTIFIED_YET",
				"FAILED_ROUTE_NO_TRIALITY_AIRLOCK_TO_ACTIVE_ASHA_CHAMBER_CERTIFIED",
			},
		},
		{
			Name:                 "flavor/Yukawa downstream branch",
			Route:                "CKM, PMNS, observed masses, flavor formulas, individual Yukawa values",
			Status:               RouteForbiddenAsSource,
			AllowedForFutureWork: false,
			Supports: []string{
				"CONDITIONAL_SUPPORT_FLAVOR_BRANCH_REMAINS_DOWNSTREAM_AFTER_GENERATION_SOURCE_DECISION",
			},
			Firewalls: []string{
				"FAILED_ROUTE_FLAVOR_FORMULA_BACKSOLVE_IS_CIRCULAR",
				"FAILED_ROUTE_NO_FLAVOR_ORIENTATION_MAP",
				"FAILED_ROUTE_NO_INDIVIDUAL_YUKAWA_VALUES",
			},
		},
	}
}

func Supports() []string {
	return []string{
		"CONDITIONAL_SUPPORT_GENERATION_MULTIPLICITY_NOT_NATIVE_IN_ACTIVE_BOARD_CURRENT_CERTIFICATE",
		"CONDITIONAL_SUPPORT_EXTERNAL_C3_GENERATION_CARRIER_CAN_BE_USED_AS_EXPLICIT_SEAL",
		"CONDITIONAL_SUPPORT_PARENT_AIRLOCK_IS_ONLY_REMAINING_NATIVE_R4_ROUTE",
		"CONDITIONAL_SUPPORT_R4_REQUIRES_EXPLICIT_CHOICE_BETWEEN_SEAL_AND_NEW_PARENT_AIRLOCK",
		"CONDITIONAL_SUPPORT_ACTIVE_BOARD_SEARCH_SHOULD_NOT_LOOP_BACK_TO_RANK_THREE_CANDIDATES",
		"CONDITIONAL_SUPPORT_NATIVE_R4_REMAINS_OPEN_ONLY_WITH_NEW_TYPED_PARENT_BOARD_TRANSPORT",
	}
}

func Failures() []string {
	return []string{
		"FAILED_ROUTE_NO_GENERATION_CARRIER_MAP",
		"FAILED_ROUTE_ACTIVE_BOARD_NATIVE_GENERATION_CARRIER_EXHAUSTED_IN_CURRENT_CERTIFICATE",
		"FAILED_ROUTE_EXTERNAL_C3_NOT_NATIVE_GENERATION_CARRIER",
		"FAILED_ROUTE_NO_PARENT_AIRLOCK_CERTIFIED_YET",
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
		"CONDITIONAL_SUPPORT_GENERATION_MULTIPLICITY_NOT_NATIVE_IN_ACTIVE_BOARD_CURRENT_CERTIFICATE",
		"CONDITIONAL_SUPPORT_EXTERNAL_C3_GENERATION_CARRIER_CAN_BE_USED_AS_EXPLICIT_SEAL",
		"CONDITIONAL_SUPPORT_PARENT_AIRLOCK_IS_ONLY_REMAINING_NATIVE_R4_ROUTE",
		"CONDITIONAL_SUPPORT_ACTIVE_BOARD_SEARCH_SHOULD_NOT_LOOP_BACK_TO_RANK_THREE_CANDIDATES",
	}
}

func RequiredFailures() []string {
	return []string{
		"FAILED_ROUTE_NO_GENERATION_CARRIER_MAP",
		"FAILED_ROUTE_ACTIVE_BOARD_NATIVE_GENERATION_CARRIER_EXHAUSTED_IN_CURRENT_CERTIFICATE",
		"FAILED_ROUTE_EXTERNAL_C3_NOT_NATIVE_GENERATION_CARRIER",
		"FAILED_ROUTE_NO_PARENT_AIRLOCK_CERTIFIED_YET",
		"FAILED_ROUTE_NO_FLAVOR_ORIENTATION_MAP",
		"FAILED_ROUTE_NO_INDIVIDUAL_YUKAWA_VALUES",
	}
}

func RouteSupports(routes []RouteAudit) []string {
	var out []string
	for _, r := range routes {
		out = append(out, r.Supports...)
	}
	return out
}

func RouteFailures(routes []RouteAudit) []string {
	var out []string
	for _, r := range routes {
		out = append(out, r.Firewalls...)
	}
	return out
}

func RouteNotes(routes []RouteAudit) []string {
	out := make([]string, 0, len(routes))
	for _, r := range routes {
		out = append(out, r.Name+" => "+string(r.Status))
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
