// Package generation2nativetrialityrestrictionk7minusr3tracebodyintertwinerrepairaudit implements
// Gate 956: Native Triality Restriction to K7Minus and R3 Tracebody Intertwiner Repair Audit.
//
// Gate 956 follows Gate 955's abstract C3 action model. It audits whether that
// model can be replaced by a native Clifford/triality-derived operator transported
// to Lambda^4 R^8, preserving K7 and K7^-, and then coupled to the dual-sealed
// aggregate R3 tracebody. The gate deliberately refuses to treat an abstract C3
// action as native triality or to use R3 trace rows as generation labels.
package generation2nativetrialityrestrictionk7minusr3tracebodyintertwinerrepairaudit

import (
	"fmt"
	"strings"
)

const (
	AuditID         = "GATE956-GENERATION2NATIVETRIALITYRESTRICTIONK7MINUSR3TRACEBODYINTERTWINERREPAIRAUDIT"
	InheritedStatus = "R4_ABSTRACT_K7_MINUS_C3_ACTION_NO_GENERATION_MAP"
	Verdict         = "NO_NATIVE_TRIALITY_TRANSPORT_TO_K7_MINUS_CERTIFIED_GATE955_C3_ACTION_REMAINS_ABSTRACT"
	Classification  = "R4_K7_MINUS_TRIALITY_ROUTE_BLOCKED_IN_CURRENT_CERTIFICATE"
	ShortStatus     = "R4_GENERATION_CARRIER_STILL_MISSING"
	NextGate        = "NEXT_GATE957_ALTERNATIVE_GENERATION_CARRIER_SEARCH_AUDIT_OR_NATIVE_TRIALITY_TRANSPORT_CONSTRUCTION"
)

const (
	K7Dim      = 7
	K7PlusDim  = 4
	K7MinusDim = 3
	TraceRows  = "3,3,1"
)

type AuditItem struct {
	Name      string
	Status    string
	Supports  []string
	Firewalls []string
}

type TrialityTransportChain struct {
	NativeTrialityOperatorConstructed bool
	OrderThree                        bool
	Nontrivial                        bool
	TransportToLambda4Certified       bool
	PreservesLambda4                  bool
	PreservesK7ContactCarrier         bool
	PreservesK7Minus                  bool
	LeakageToK7PlusUnknown            bool
	AbstractC3FromGate955Realized     bool
}

type TracebodyIntertwiner struct {
	CandidateName                string
	Certified                    bool
	RequiresNativeK7MinusAction  bool
	UsesArbitraryBasisFit        bool
	UsesR3RowsAsGenerationLabels bool
	PreservesR3DualSeal          bool
	UsesFlavorBacksolve          bool
	UsesObservedMassesOrYukawas  bool
	UsesCKMPMNSInput             bool
}

type Analysis struct {
	AuditID                     string
	Inherited                   string
	Verdict                     string
	Classification              string
	ShortStatus                 string
	NativeR3                    bool
	R3DualSealRequired          bool
	GenerationCarrierCertified  bool
	FlavorOrientationCertified  bool
	IndividualYukawaCertified   bool
	PhysicalAssignmentCertified bool
	OfficialLedgerUpdate        bool
	K7MinusDimension            int
	K7PlusDimension             int
	R3TraceRows                 string
	Transport                   TrialityTransportChain
	Intertwiner                 TracebodyIntertwiner
	Items                       []AuditItem
	Supports                    []string
	Failures                    []string
	Final                       string
}

func BuildDefault() (Analysis, error) {
	transport := TrialityTransportChain{
		NativeTrialityOperatorConstructed: false,
		OrderThree:                        false,
		Nontrivial:                        false,
		TransportToLambda4Certified:       false,
		PreservesLambda4:                  false,
		PreservesK7ContactCarrier:         false,
		PreservesK7Minus:                  false,
		LeakageToK7PlusUnknown:            true,
		AbstractC3FromGate955Realized:     false,
	}
	intertwiner := TracebodyIntertwiner{
		CandidateName:                "K7MinusTrialityR3TracebodyIntertwiner",
		Certified:                    false,
		RequiresNativeK7MinusAction:  true,
		UsesArbitraryBasisFit:        true,
		UsesR3RowsAsGenerationLabels: true,
		PreservesR3DualSeal:          true,
		UsesFlavorBacksolve:          false,
		UsesObservedMassesOrYukawas:  false,
		UsesCKMPMNSInput:             false,
	}
	items := DefaultItems(transport, intertwiner)
	if len(items) != 7 {
		return Analysis{}, fmt.Errorf("expected 7 audit items, got %d", len(items))
	}
	a := Analysis{
		AuditID:                     AuditID,
		Inherited:                   InheritedStatus,
		Verdict:                     Verdict,
		Classification:              Classification,
		ShortStatus:                 ShortStatus,
		NativeR3:                    false,
		R3DualSealRequired:          true,
		GenerationCarrierCertified:  false,
		FlavorOrientationCertified:  false,
		IndividualYukawaCertified:   false,
		PhysicalAssignmentCertified: false,
		OfficialLedgerUpdate:        false,
		K7MinusDimension:            K7MinusDim,
		K7PlusDimension:             K7PlusDim,
		R3TraceRows:                 TraceRows,
		Transport:                   transport,
		Intertwiner:                 intertwiner,
		Items:                       items,
		Supports:                    Supports(),
		Failures:                    Failures(),
		Final:                       "Gate 956 audits the repair route demanded by Gate 955: native triality operator -> Lambda^4 transport -> K7 preservation -> K7^- restriction -> R3 aggregate tracebody intertwiner. The current certificate does not supply a native triality operator or transport to Lambda^4, so the Gate 955 abstract C3 action remains a noncanonical model. The R4 generation carrier map is still missing, and flavor/Yukawa/particle claims remain firewalled.",
	}
	if err := validate(a); err != nil {
		return Analysis{}, err
	}
	return a, nil
}

func validate(a Analysis) error {
	if a.NativeR3 || !a.R3DualSealRequired || a.OfficialLedgerUpdate {
		return fmt.Errorf("Gate 956 must preserve R3 dual seal and avoid native/official overclaim")
	}
	if a.GenerationCarrierCertified || a.FlavorOrientationCertified || a.IndividualYukawaCertified || a.PhysicalAssignmentCertified {
		return fmt.Errorf("Gate 956 overclaimed R4 downstream theorem status")
	}
	if a.K7MinusDimension != 3 || a.K7PlusDimension != 4 || a.K7MinusDimension+a.K7PlusDimension != K7Dim {
		return fmt.Errorf("bad K7 polarity dimensions")
	}
	if a.R3TraceRows != TraceRows {
		return fmt.Errorf("R3 tracebody rows must remain aggregate %s", TraceRows)
	}
	if a.Transport.NativeTrialityOperatorConstructed || a.Transport.TransportToLambda4Certified || a.Transport.PreservesK7ContactCarrier || a.Transport.PreservesK7Minus || a.Transport.AbstractC3FromGate955Realized {
		return fmt.Errorf("native triality transport/restriction overclaimed: %#v", a.Transport)
	}
	if !a.Transport.LeakageToK7PlusUnknown {
		return fmt.Errorf("K7+ leakage should remain unknown without native transport operator")
	}
	if a.Intertwiner.Certified || !a.Intertwiner.RequiresNativeK7MinusAction || !a.Intertwiner.UsesArbitraryBasisFit || !a.Intertwiner.UsesR3RowsAsGenerationLabels || !a.Intertwiner.PreservesR3DualSeal {
		return fmt.Errorf("intertwiner obstruction flags incorrect: %#v", a.Intertwiner)
	}
	if a.Intertwiner.UsesFlavorBacksolve || a.Intertwiner.UsesObservedMassesOrYukawas || a.Intertwiner.UsesCKMPMNSInput {
		return fmt.Errorf("forbidden empirical/flavor input used")
	}
	return nil
}

func DefaultItems(t TrialityTransportChain, g TracebodyIntertwiner) []AuditItem {
	return []AuditItem{
		{
			Name:   "native triality operator location",
			Status: statusIf(t.NativeTrialityOperatorConstructed, "NATIVE_TRIALITY_OPERATOR_CONSTRUCTED", "BLOCKED_NO_NATIVE_TRIALITY_OPERATOR_ON_ACTIVE_ASHA_BOARD"),
			Supports: []string{
				"CONDITIONAL_SUPPORT_NATIVE_TRIALITY_OPERATOR_IS_CORRECT_REPAIR_TARGET",
				"CONDITIONAL_SUPPORT_GATE955_ABSTRACT_C3_MODEL_IDENTIFIES_REQUIRED_NATIVE_SOURCE",
			},
			Firewalls: []string{
				"FAILED_ROUTE_NO_NATIVE_TRIALITY_OPERATOR_ON_ACTIVE_ASHA_BOARD",
				"FAILED_ROUTE_CLIFFORD_TRIALITY_SOURCE_NOT_TRANSPORTED_TO_CURRENT_K7_BOARD",
			},
		},
		{
			Name:   "triality transport to Lambda4 chamber",
			Status: statusIf(t.TransportToLambda4Certified, "TRIALITY_TRANSPORTS_TO_LAMBDA4_CHAMBER", "BLOCKED_NO_TRIALITY_TO_LAMBDA4_TRANSPORT_MAP"),
			Supports: []string{
				"CONDITIONAL_SUPPORT_TRIALITY_TO_LAMBDA4_TRANSPORT_IS_DEEPEST_CURRENT_WOUND",
			},
			Firewalls: []string{
				"FAILED_ROUTE_NO_TRIALITY_TO_LAMBDA4_TRANSPORT_MAP",
				"FAILED_ROUTE_NATIVE_TRIALITY_NOT_REALIZED_AS_ENDOMORPHISM_OF_LAMBDA4_R8",
			},
		},
		{
			Name:   "preserve Boolean-octonionic K7 contact carrier",
			Status: statusIf(t.PreservesK7ContactCarrier, "TRIALITY_PRESERVES_K7_CONTACT_CARRIER", "BLOCKED_NO_TRIALITY_ACTION_ON_K7_CONTACT_CARRIER"),
			Supports: []string{
				"CONDITIONAL_SUPPORT_K7_CONTACT_CARRIER_IS_THE_REQUIRED_NATIVE_SUPPORT_FOR_RESTRICTION",
			},
			Firewalls: []string{
				"FAILED_ROUTE_TRIALITY_DOES_NOT_PRESERVE_K7_CONTACT_CARRIER_IN_CURRENT_CERTIFICATE",
				"FAILED_ROUTE_NO_P_K7_T_P_K7_EQUALS_T_P_K7_CERTIFICATE",
			},
		},
		{
			Name:   "restrict native triality to K7-minus",
			Status: statusIf(t.PreservesK7Minus, "NATIVE_TRIALITY_RESTRICTS_TO_K7_MINUS", "BLOCKED_NO_NATIVE_TRIALITY_RESTRICTION_TO_K7_MINUS"),
			Supports: []string{
				"CONDITIONAL_SUPPORT_K7_MINUS_REMAINS_STRONGEST_GENERATION_CARRIER_SHAPE",
			},
			Firewalls: []string{
				"FAILED_ROUTE_NATIVE_TRIALITY_DOES_NOT_RESTRICT_TO_K7_MINUS_IN_CURRENT_CERTIFICATE",
				"FAILED_ROUTE_K7_PLUS_LEAKAGE_NOT_EXCLUDED_WITHOUT_NATIVE_OPERATOR",
			},
		},
		{
			Name:   "compare Gate955 abstract C3 with native triality",
			Status: statusIf(t.AbstractC3FromGate955Realized, "GATE955_ABSTRACT_C3_MODEL_NATIVE_REALIZABLE", "GATE955_C3_ACTION_REMAINS_ABSTRACT_NONCANONICAL_MODEL"),
			Supports: []string{
				"CONDITIONAL_SUPPORT_GATE955_ABSTRACT_MODEL_PROVIDES_REPAIR_TARGET_BUT_NOT_NATIVE_SOURCE",
			},
			Firewalls: []string{
				"FAILED_ROUTE_GATE955_C3_ACTION_REMAINS_ABSTRACT_NONCANONICAL_MODEL",
				"FAILED_ROUTE_ABSTRACT_C3_ACTION_CANNOT_BE_PROMOTED_BY_BASIS_FIT",
			},
		},
		{
			Name:   "R3 aggregate tracebody intertwiner repair",
			Status: statusIf(g.Certified, "NATIVE_K7_MINUS_TRIALITY_ACTION_INTERTWINES_WITH_R3_TRACEBODY", "BLOCKED_NO_NATIVE_K7_MINUS_TRIALITY_R3_TRACEBODY_INTERTWINER"),
			Supports: []string{
				"CONDITIONAL_SUPPORT_R3_TRACEBODY_CAN_BE_TARGET_ONLY_AS_DUALSEALED_AGGREGATE",
			},
			Firewalls: []string{
				"FAILED_ROUTE_NO_NATIVE_K7_MINUS_TRIALITY_R3_TRACEBODY_INTERTWINER",
				"FAILED_ROUTE_INTERTWINER_REQUIRES_ARBITRARY_R3_ROW_IDENTIFICATION",
				"FAILED_ROUTE_R3_TRACE_ROWS_USED_AS_GENERATION_LABELS",
			},
		},
		{
			Name:   "noncircularity and flavor firewall",
			Status: statusIf(!g.UsesFlavorBacksolve && !g.UsesObservedMassesOrYukawas && !g.UsesCKMPMNSInput, "NONCIRCULARITY_FIREWALL_PRESERVED", "FORBIDDEN_FLAVOR_INPUT_USED"),
			Supports: []string{
				"CONDITIONAL_SUPPORT_REPAIR_AUDIT_USES_NO_FLAVOR_BACKSOLVE_OR_OBSERVED_YUKAWA_DATA",
				"CONDITIONAL_SUPPORT_R4_GENERATION_ROUTE_REMAINS_SEPARATE_FROM_FLAVOR_ORIENTATION",
			},
			Firewalls: []string{
				"FAILED_ROUTE_R3_TRACEBRIDGE_REMAINS_DUALSEALED_NOT_NATIVE",
				"FAILED_ROUTE_NO_FLAVOR_ORIENTATION_MAP",
				"FAILED_ROUTE_NO_INDIVIDUAL_YUKAWA_VALUES",
				"FAILED_ROUTE_NO_PHYSICAL_PARTICLE_ASSIGNMENT",
				"FAILED_ROUTE_NO_CKM_PMNS_THEOREM",
				"FAILED_ROUTE_NO_PMNS_THEOREM",
				"FAILED_ROUTE_NO_R4_NATIVE_YUKAWA_SPECTRUM_THEOREM",
				"FAILED_ROUTE_NO_OFFICIAL_LEDGER_UPDATE_ALLOWED",
			},
		},
	}
}

func Supports() []string {
	return []string{
		"CONDITIONAL_SUPPORT_NATIVE_TRIALITY_OPERATOR_IS_CORRECT_REPAIR_TARGET",
		"CONDITIONAL_SUPPORT_GATE955_ABSTRACT_C3_MODEL_IDENTIFIES_REQUIRED_NATIVE_SOURCE",
		"CONDITIONAL_SUPPORT_TRIALITY_TO_LAMBDA4_TRANSPORT_IS_DEEPEST_CURRENT_WOUND",
		"CONDITIONAL_SUPPORT_K7_CONTACT_CARRIER_IS_THE_REQUIRED_NATIVE_SUPPORT_FOR_RESTRICTION",
		"CONDITIONAL_SUPPORT_K7_MINUS_REMAINS_STRONGEST_GENERATION_CARRIER_SHAPE",
		"CONDITIONAL_SUPPORT_GATE955_ABSTRACT_MODEL_PROVIDES_REPAIR_TARGET_BUT_NOT_NATIVE_SOURCE",
		"CONDITIONAL_SUPPORT_R3_TRACEBODY_CAN_BE_TARGET_ONLY_AS_DUALSEALED_AGGREGATE",
		"CONDITIONAL_SUPPORT_REPAIR_AUDIT_USES_NO_FLAVOR_BACKSOLVE_OR_OBSERVED_YUKAWA_DATA",
		"CONDITIONAL_SUPPORT_R4_GENERATION_ROUTE_REMAINS_SEPARATE_FROM_FLAVOR_ORIENTATION",
	}
}

func Failures() []string {
	return []string{
		"FAILED_ROUTE_NO_GENERATION_CARRIER_MAP",
		"FAILED_ROUTE_NO_NATIVE_TRIALITY_OPERATOR_ON_ACTIVE_ASHA_BOARD",
		"FAILED_ROUTE_NO_TRIALITY_TO_LAMBDA4_TRANSPORT_MAP",
		"FAILED_ROUTE_NATIVE_TRIALITY_NOT_REALIZED_AS_ENDOMORPHISM_OF_LAMBDA4_R8",
		"FAILED_ROUTE_TRIALITY_DOES_NOT_PRESERVE_K7_CONTACT_CARRIER_IN_CURRENT_CERTIFICATE",
		"FAILED_ROUTE_NATIVE_TRIALITY_DOES_NOT_RESTRICT_TO_K7_MINUS_IN_CURRENT_CERTIFICATE",
		"FAILED_ROUTE_K7_PLUS_LEAKAGE_NOT_EXCLUDED_WITHOUT_NATIVE_OPERATOR",
		"FAILED_ROUTE_GATE955_C3_ACTION_REMAINS_ABSTRACT_NONCANONICAL_MODEL",
		"FAILED_ROUTE_NO_NATIVE_K7_MINUS_TRIALITY_R3_TRACEBODY_INTERTWINER",
		"FAILED_ROUTE_INTERTWINER_REQUIRES_ARBITRARY_R3_ROW_IDENTIFICATION",
		"FAILED_ROUTE_R3_TRACE_ROWS_USED_AS_GENERATION_LABELS",
		"FAILED_ROUTE_R3_TRACEBRIDGE_REMAINS_DUALSEALED_NOT_NATIVE",
		"FAILED_ROUTE_NO_FLAVOR_ORIENTATION_MAP",
		"FAILED_ROUTE_NO_INDIVIDUAL_YUKAWA_VALUES",
		"FAILED_ROUTE_NO_PHYSICAL_PARTICLE_ASSIGNMENT",
		"FAILED_ROUTE_NO_CKM_PMNS_THEOREM",
		"FAILED_ROUTE_NO_PMNS_THEOREM",
		"FAILED_ROUTE_NO_R4_NATIVE_YUKAWA_SPECTRUM_THEOREM",
		"FAILED_ROUTE_NO_OFFICIAL_LEDGER_UPDATE_ALLOWED",
	}
}

func RequiredSupports() []string {
	return []string{
		"CONDITIONAL_SUPPORT_NATIVE_TRIALITY_OPERATOR_IS_CORRECT_REPAIR_TARGET",
		"CONDITIONAL_SUPPORT_TRIALITY_TO_LAMBDA4_TRANSPORT_IS_DEEPEST_CURRENT_WOUND",
		"CONDITIONAL_SUPPORT_K7_MINUS_REMAINS_STRONGEST_GENERATION_CARRIER_SHAPE",
		"CONDITIONAL_SUPPORT_R3_TRACEBODY_CAN_BE_TARGET_ONLY_AS_DUALSEALED_AGGREGATE",
		"CONDITIONAL_SUPPORT_REPAIR_AUDIT_USES_NO_FLAVOR_BACKSOLVE_OR_OBSERVED_YUKAWA_DATA",
	}
}

func RequiredFailures() []string {
	return []string{
		"FAILED_ROUTE_NO_GENERATION_CARRIER_MAP",
		"FAILED_ROUTE_NO_NATIVE_TRIALITY_OPERATOR_ON_ACTIVE_ASHA_BOARD",
		"FAILED_ROUTE_NO_TRIALITY_TO_LAMBDA4_TRANSPORT_MAP",
		"FAILED_ROUTE_NATIVE_TRIALITY_DOES_NOT_RESTRICT_TO_K7_MINUS_IN_CURRENT_CERTIFICATE",
		"FAILED_ROUTE_GATE955_C3_ACTION_REMAINS_ABSTRACT_NONCANONICAL_MODEL",
		"FAILED_ROUTE_NO_NATIVE_K7_MINUS_TRIALITY_R3_TRACEBODY_INTERTWINER",
		"FAILED_ROUTE_R3_TRACE_ROWS_USED_AS_GENERATION_LABELS",
		"FAILED_ROUTE_NO_FLAVOR_ORIENTATION_MAP",
		"FAILED_ROUTE_NO_INDIVIDUAL_YUKAWA_VALUES",
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
