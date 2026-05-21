// Package generation2flavororientationmapconstructionauditundersealedr4 implements
// Gate 962: FlavorOrientationMap Construction Audit Under Sealed R4.
//
// Gate 962 follows the Gate 961 typing phase. It attempts to construct the
// flavor-orientation map from the sealed C^3 generation carrier into the
// post-orientation ledger interface, but preserves the core U(3) ambiguity: an
// external C^3 supplies three family slots, not a canonical flavor basis.
package generation2flavororientationmapconstructionauditundersealedr4

import (
	"fmt"
	"strings"
)

const (
	AuditID         = "GATE962-GENERATION2FLAVORORIENTATIONMAPCONSTRUCTIONAUDITUNDERSEALEDR4"
	InheritedStatus = "R4_SEALED_GENERATION_CARRIER_NO_FLAVOR_ORIENTATION"
	Verdict         = "SEALED_C3_DEFINES_FAMILY_SLOT_ORBIT_BUT_NO_CANONICAL_FLAVOR_ORIENTATION_MAP_CERTIFIED"
	Classification  = "R4_FLAVOR_ORIENTATION_CONSTRUCTION_FAILED_CANONICAL_SELECTOR_MISSING"
	ShortStatus     = "R4_EXTERNAL_C3_HAS_U3_FAMILY_ORBIT_NO_FLAVOR_BASIS"
	NextGate        = "NEXT_GATE963_CANONICAL_FLAVOR_SELECTOR_OR_EXTERNAL_FLAVOR_ORIENTATION_SEAL_DECISION_AUDIT"
)

type CandidateStatus string

const (
	StatusBasisBlocked     CandidateStatus = "FAILED_NO_NATIVE_C3_BASIS_OR_ORDERING"
	StatusRowsBlocked      CandidateStatus = "FAILED_R3_ROW_MATCHING_NOT_ORIENTATION"
	StatusInterfaceBlocked CandidateStatus = "VALID_INTERFACE_TARGET_BUT_NO_FAMILY_SELECTOR"
	StatusSocketBlocked    CandidateStatus = "FAILED_SOCKET_SELECTOR_NOT_FAMILY_ORIENTATION"
	StatusBacksolveBlocked CandidateStatus = "FAILED_FLAVOR_BACKSOLVE_CIRCULAR"
	StatusOrbitOnly        CandidateStatus = "PARTIAL_U3_FAMILY_ORIENTATION_ORBIT_ONLY"
	StatusSealPreserved    CandidateStatus = "SEALS_PRESERVED_NO_DOWNSTREAM_CLAIMS"
)

type ConstructionCandidate struct {
	Name                      string
	Route                     string
	Status                    CandidateStatus
	Attempted                 bool
	ConstructionCertified     bool
	CanonicalBasisSelected    bool
	ProvidesU3OrbitOnly       bool
	UsesR3RowsAsGeneration    bool
	UsesAFOrientAsInterface   bool
	UsesSocketAsFamily        bool
	UsesFlavorFormulaAsSource bool
	UsesObservedFlavorData    bool
	AssignsPhysicalParticles  bool
	ProvidesIndividualYukawa  bool
	ProvidesCKMPMNS           bool
	OfficialLedgerUpdate      bool
	Supports                  []string
	Firewalls                 []string
}

type Decision struct {
	ExternalC3DomainAvailable       bool
	ExternalC3NativeGeneration      bool
	U3FamilyGaugeFreedomDetected    bool
	U3OrbitClassAvailable           bool
	CanonicalFlavorBasisCertified   bool
	CanonicalRepresentativeSelected bool
	FlavorOrientationMapCertified   bool
	AFOrientIsValidInterfaceTarget  bool
	AFOrientSuppliesFamilySelector  bool
	R3TracebodyAggregateTargetValid bool
	R3RowsUsedAsGenerationLabels    bool
	FlavorFormulaBacksolveAllowed   bool
	ObservedFlavorInputAllowed      bool
	PhysicalAssignmentAllowed       bool
	IndividualYukawaAllowed         bool
	CKMPMNSAllowed                  bool
	OfficialLedgerUpdateAllowed     bool
	R3DualSealPreserved             bool
	ExternalC3SealPreserved         bool
	RequiresCanonicalSelectorOrSeal bool
}

type Analysis struct {
	AuditID        string
	Inherited      string
	Verdict        string
	Classification string
	ShortStatus    string
	NextGate       string
	Domain         string
	Target         string
	MissingObject  string
	Decision       Decision
	Candidates     []ConstructionCandidate
	Supports       []string
	Failures       []string
	Final          string
}

func BuildDefault() (Analysis, error) {
	candidates := DefaultCandidates()
	if len(candidates) != 7 {
		return Analysis{}, fmt.Errorf("expected 7 construction candidates, got %d", len(candidates))
	}
	decision := Decision{
		ExternalC3DomainAvailable:       true,
		ExternalC3NativeGeneration:      false,
		U3FamilyGaugeFreedomDetected:    true,
		U3OrbitClassAvailable:           true,
		CanonicalFlavorBasisCertified:   false,
		CanonicalRepresentativeSelected: false,
		FlavorOrientationMapCertified:   false,
		AFOrientIsValidInterfaceTarget:  true,
		AFOrientSuppliesFamilySelector:  false,
		R3TracebodyAggregateTargetValid: true,
		R3RowsUsedAsGenerationLabels:    false,
		FlavorFormulaBacksolveAllowed:   false,
		ObservedFlavorInputAllowed:      false,
		PhysicalAssignmentAllowed:       false,
		IndividualYukawaAllowed:         false,
		CKMPMNSAllowed:                  false,
		OfficialLedgerUpdateAllowed:     false,
		R3DualSealPreserved:             true,
		ExternalC3SealPreserved:         true,
		RequiresCanonicalSelectorOrSeal: true,
	}
	a := Analysis{
		AuditID:        AuditID,
		Inherited:      InheritedStatus,
		Verdict:        Verdict,
		Classification: Classification,
		ShortStatus:    ShortStatus,
		NextGate:       NextGate,
		Domain:         "G_gen^seal = C^3 under ExternalGenerationCarrierSeal(C3), with intrinsic U(3) family-gauge ambiguity",
		Target:         "Orient(A_F^orient, R3_tracebody) as a ledger interface, not physical particles or mass eigenstates",
		MissingObject:  "CanonicalFlavorSelector or ExternalFlavorOrientationSeal",
		Decision:       decision,
		Candidates:     candidates,
		Supports:       Supports(),
		Failures:       Failures(),
		Final:          "Gate 962 attempts the FlavorOrientationMap construction under R3DualSeal plus ExternalGenerationCarrierSeal(C3). The sealed C^3 carrier supplies a family-slot orbit only up to U(3); no native basis, ordering, or canonical representative is selected. A_F^orient is a valid interface target but supplies socket/gauge structure, not family orientation. R3 trace rows and flavor formulas remain downstream or aggregate-only. The construction fails honestly: Phi_flav is not certified, and the next wound is a CanonicalFlavorSelector or an explicit ExternalFlavorOrientationSeal.",
	}
	if err := validate(a); err != nil {
		return Analysis{}, err
	}
	return a, nil
}

func validate(a Analysis) error {
	if a.Inherited != InheritedStatus || a.Verdict != Verdict || a.Classification != Classification || a.ShortStatus != ShortStatus {
		return fmt.Errorf("bad Gate 962 identity")
	}
	if !a.Decision.ExternalC3DomainAvailable || a.Decision.ExternalC3NativeGeneration {
		return fmt.Errorf("Gate 962 must inherit external C3 as a seal only: %#v", a.Decision)
	}
	if !a.Decision.U3FamilyGaugeFreedomDetected || !a.Decision.U3OrbitClassAvailable || a.Decision.CanonicalFlavorBasisCertified || a.Decision.CanonicalRepresentativeSelected || a.Decision.FlavorOrientationMapCertified {
		return fmt.Errorf("Gate 962 must expose U3 orbit but no canonical flavor basis/map: %#v", a.Decision)
	}
	if !a.Decision.AFOrientIsValidInterfaceTarget || a.Decision.AFOrientSuppliesFamilySelector {
		return fmt.Errorf("Gate 962 must allow A_F^orient only as interface, not family selector: %#v", a.Decision)
	}
	if !a.Decision.R3TracebodyAggregateTargetValid || a.Decision.R3RowsUsedAsGenerationLabels {
		return fmt.Errorf("Gate 962 must preserve R3 tracebody as aggregate target only: %#v", a.Decision)
	}
	if a.Decision.FlavorFormulaBacksolveAllowed || a.Decision.ObservedFlavorInputAllowed || a.Decision.PhysicalAssignmentAllowed || a.Decision.IndividualYukawaAllowed || a.Decision.CKMPMNSAllowed || a.Decision.OfficialLedgerUpdateAllowed {
		return fmt.Errorf("Gate 962 overclaimed flavor data, physical assignments, spectrum, or ledger update: %#v", a.Decision)
	}
	if !a.Decision.R3DualSealPreserved || !a.Decision.ExternalC3SealPreserved || !a.Decision.RequiresCanonicalSelectorOrSeal {
		return fmt.Errorf("Gate 962 must preserve seals and require selector/seal: %#v", a.Decision)
	}
	for _, c := range a.Candidates {
		if !c.Attempted {
			return fmt.Errorf("candidate not attempted: %#v", c)
		}
		if c.ConstructionCertified || c.AssignsPhysicalParticles || c.ProvidesIndividualYukawa || c.ProvidesCKMPMNS || c.OfficialLedgerUpdate {
			return fmt.Errorf("candidate overclaimed downstream construction: %#v", c)
		}
		if c.UsesObservedFlavorData || c.UsesFlavorFormulaAsSource {
			return fmt.Errorf("candidate used forbidden empirical/flavor source: %#v", c)
		}
	}
	return nil
}

func DefaultCandidates() []ConstructionCandidate {
	return []ConstructionCandidate{
		{
			Name:                   "canonical basis of external C3",
			Route:                  "e_i in C^3_gen,seal -> flavor slot i",
			Status:                 StatusBasisBlocked,
			Attempted:              true,
			ConstructionCertified:  false,
			CanonicalBasisSelected: false,
			Supports: []string{
				"CONDITIONAL_SUPPORT_EXTERNAL_C3_SUPPLIES_THREE_FAMILY_SLOT_DOMAIN",
			},
			Firewalls: []string{
				"FAILED_ROUTE_EXTERNAL_C3_HAS_NO_NATIVE_BASIS_OR_ORDERING",
				"FAILED_ROUTE_CANONICAL_FLAVOR_BASIS_REQUIRES_EXTRA_ORIENTATION_SEAL",
			},
		},
		{
			Name:                   "orient using R3 trace rows",
			Route:                  "match C^3 slots to R3 row multiset 3,3,1",
			Status:                 StatusRowsBlocked,
			Attempted:              true,
			ConstructionCertified:  false,
			UsesR3RowsAsGeneration: false,
			Supports: []string{
				"CONDITIONAL_SUPPORT_R3_TRACEBODY_REMAINS_VALID_AGGREGATE_COMPATIBILITY_TARGET",
			},
			Firewalls: []string{
				"FAILED_ROUTE_R3_TRACE_ROWS_ARE_NOT_GENERATION_LABELS",
				"FAILED_ROUTE_R3_TRACEBODY_CANNOT_ORIENT_EXTERNAL_C3_BY_ROW_MATCHING",
				"FAILED_ROUTE_3_PLUS_3_PLUS_1_NOT_FLAVOR_ORIENTATION",
			},
		},
		{
			Name:                    "orient using A_F^orient",
			Route:                   "C^3_gen,seal -> post-orientation finite algebra interface",
			Status:                  StatusInterfaceBlocked,
			Attempted:               true,
			ConstructionCertified:   false,
			UsesAFOrientAsInterface: true,
			Supports: []string{
				"CONDITIONAL_SUPPORT_A_F_ORIENT_IS_VALID_INTERFACE_TARGET",
			},
			Firewalls: []string{
				"FAILED_ROUTE_A_F_ORIENT_DOES_NOT_SUPPLY_FAMILY_ORIENTATION_SELECTOR",
				"FAILED_ROUTE_SOCKET_ORIENTATION_NOT_FAMILY_ORIENTATION",
			},
		},
		{
			Name:                  "orient using Fock 1+3 / B-L",
			Route:                 "use P3 or B-L as family selector",
			Status:                StatusSocketBlocked,
			Attempted:             true,
			ConstructionCertified: false,
			UsesSocketAsFamily:    false,
			Firewalls: []string{
				"FAILED_ROUTE_FOCK_P3_RANK_THREE_NOT_FLAVOR_ORIENTATION",
				"FAILED_ROUTE_B_MINUS_L_SELECTOR_NOT_FAMILY_ORIENTATION",
			},
		},
		{
			Name:                      "orient using flavor-wall formulas",
			Route:                     "epsilon_e, kappa_e, Koide, CKM/PMNS residuals, or observed ordering -> Phi_flav",
			Status:                    StatusBacksolveBlocked,
			Attempted:                 true,
			ConstructionCertified:     false,
			UsesFlavorFormulaAsSource: false,
			UsesObservedFlavorData:    false,
			Firewalls: []string{
				"FAILED_ROUTE_FLAVOR_FORMULA_BACKSOLVE_IS_CIRCULAR",
				"FAILED_ROUTE_OBSERVED_FLAVOR_DATA_CANNOT_ORIENT_EXTERNAL_C3",
				"FAILED_ROUTE_EPSILON_E_OR_KAPPA_E_CANNOT_SOURCE_FLAVOR_ORIENTATION",
				"FAILED_ROUTE_KOIDE_BRANCH_CANNOT_SOURCE_FLAVOR_ORIENTATION",
			},
		},
		{
			Name:                  "equivariant family orbit only",
			Route:                 "define [Phi_flav]_{U(3)} without selecting a representative",
			Status:                StatusOrbitOnly,
			Attempted:             true,
			ConstructionCertified: false,
			ProvidesU3OrbitOnly:   true,
			Supports: []string{
				"CONDITIONAL_SUPPORT_EXTERNAL_C3_DEFINES_FAMILY_ORIENTATION_ORBIT_UP_TO_U3",
				"CONDITIONAL_SUPPORT_FLAVOR_ORIENTATION_REQUIRES_CANONICAL_SELECTOR_OR_ADDITIONAL_SEAL",
			},
			Firewalls: []string{
				"FAILED_ROUTE_U3_ORBIT_CLASS_IS_NOT_FLAVOR_ORIENTATION_MAP",
				"FAILED_ROUTE_NO_CANONICAL_REPRESENTATIVE_SELECTED",
			},
		},
		{
			Name:                  "seal preservation and downstream firewall",
			Route:                 "R3DualSeal + ExternalGenerationCarrierSeal(C3) remain inherited",
			Status:                StatusSealPreserved,
			Attempted:             true,
			ConstructionCertified: false,
			Supports: []string{
				"CONDITIONAL_SUPPORT_R3_DUALSEAL_AND_EXTERNAL_C3_SEAL_REMAIN_VISIBLE",
			},
			Firewalls: []string{
				"FAILED_ROUTE_NO_INDIVIDUAL_YUKAWA_VALUES",
				"FAILED_ROUTE_NO_PHYSICAL_PARTICLE_ASSIGNMENT",
				"FAILED_ROUTE_NO_CKM_PMNS_THEOREM",
				"FAILED_ROUTE_NO_PMNS_THEOREM",
				"FAILED_ROUTE_NO_R4_NATIVE_YUKAWA_SPECTRUM_THEOREM",
				"FAILED_ROUTE_R3_TRACEBRIDGE_REMAINS_DUALSEALED_NOT_NATIVE",
				"FAILED_ROUTE_EXTERNAL_C3_NOT_NATIVE_GENERATION_CARRIER",
				"FAILED_ROUTE_NO_OFFICIAL_LEDGER_UPDATE_ALLOWED",
			},
		},
	}
}

func Supports() []string {
	return []string{
		"CONDITIONAL_SUPPORT_EXTERNAL_C3_SUPPLIES_THREE_FAMILY_SLOT_DOMAIN",
		"CONDITIONAL_SUPPORT_A_F_ORIENT_IS_VALID_INTERFACE_TARGET",
		"CONDITIONAL_SUPPORT_R3_TRACEBODY_REMAINS_VALID_AGGREGATE_COMPATIBILITY_TARGET",
		"CONDITIONAL_SUPPORT_EXTERNAL_C3_DEFINES_FAMILY_ORIENTATION_ORBIT_UP_TO_U3",
		"CONDITIONAL_SUPPORT_FLAVOR_ORIENTATION_REQUIRES_CANONICAL_SELECTOR_OR_ADDITIONAL_SEAL",
		"CONDITIONAL_SUPPORT_R3_DUALSEAL_AND_EXTERNAL_C3_SEAL_REMAIN_VISIBLE",
	}
}

func Failures() []string {
	return []string{
		"FAILED_ROUTE_NO_FLAVOR_ORIENTATION_MAP_CERTIFIED",
		"FAILED_ROUTE_EXTERNAL_C3_HAS_NO_NATIVE_BASIS_OR_ORDERING",
		"FAILED_ROUTE_U3_ORBIT_CLASS_IS_NOT_FLAVOR_ORIENTATION_MAP",
		"FAILED_ROUTE_NO_CANONICAL_REPRESENTATIVE_SELECTED",
		"FAILED_ROUTE_R3_TRACE_ROWS_ARE_NOT_GENERATION_LABELS",
		"FAILED_ROUTE_R3_TRACEBODY_CANNOT_ORIENT_EXTERNAL_C3_BY_ROW_MATCHING",
		"FAILED_ROUTE_A_F_ORIENT_DOES_NOT_SUPPLY_FAMILY_ORIENTATION_SELECTOR",
		"FAILED_ROUTE_SOCKET_ORIENTATION_NOT_FAMILY_ORIENTATION",
		"FAILED_ROUTE_FOCK_P3_RANK_THREE_NOT_FLAVOR_ORIENTATION",
		"FAILED_ROUTE_B_MINUS_L_SELECTOR_NOT_FAMILY_ORIENTATION",
		"FAILED_ROUTE_FLAVOR_FORMULA_BACKSOLVE_IS_CIRCULAR",
		"FAILED_ROUTE_OBSERVED_FLAVOR_DATA_CANNOT_ORIENT_EXTERNAL_C3",
		"FAILED_ROUTE_NO_INDIVIDUAL_YUKAWA_VALUES",
		"FAILED_ROUTE_NO_PHYSICAL_PARTICLE_ASSIGNMENT",
		"FAILED_ROUTE_NO_CKM_PMNS_THEOREM",
		"FAILED_ROUTE_NO_PMNS_THEOREM",
		"FAILED_ROUTE_NO_R4_NATIVE_YUKAWA_SPECTRUM_THEOREM",
		"FAILED_ROUTE_R3_TRACEBRIDGE_REMAINS_DUALSEALED_NOT_NATIVE",
		"FAILED_ROUTE_EXTERNAL_C3_NOT_NATIVE_GENERATION_CARRIER",
		"FAILED_ROUTE_NO_OFFICIAL_LEDGER_UPDATE_ALLOWED",
	}
}

func RequiredSupports() []string {
	return []string{
		"CONDITIONAL_SUPPORT_EXTERNAL_C3_SUPPLIES_THREE_FAMILY_SLOT_DOMAIN",
		"CONDITIONAL_SUPPORT_A_F_ORIENT_IS_VALID_INTERFACE_TARGET",
		"CONDITIONAL_SUPPORT_R3_TRACEBODY_REMAINS_VALID_AGGREGATE_COMPATIBILITY_TARGET",
		"CONDITIONAL_SUPPORT_EXTERNAL_C3_DEFINES_FAMILY_ORIENTATION_ORBIT_UP_TO_U3",
		"CONDITIONAL_SUPPORT_FLAVOR_ORIENTATION_REQUIRES_CANONICAL_SELECTOR_OR_ADDITIONAL_SEAL",
	}
}

func RequiredFailures() []string {
	return []string{
		"FAILED_ROUTE_NO_FLAVOR_ORIENTATION_MAP_CERTIFIED",
		"FAILED_ROUTE_EXTERNAL_C3_HAS_NO_NATIVE_BASIS_OR_ORDERING",
		"FAILED_ROUTE_U3_ORBIT_CLASS_IS_NOT_FLAVOR_ORIENTATION_MAP",
		"FAILED_ROUTE_NO_CANONICAL_REPRESENTATIVE_SELECTED",
		"FAILED_ROUTE_R3_TRACE_ROWS_ARE_NOT_GENERATION_LABELS",
		"FAILED_ROUTE_A_F_ORIENT_DOES_NOT_SUPPLY_FAMILY_ORIENTATION_SELECTOR",
		"FAILED_ROUTE_FLAVOR_FORMULA_BACKSOLVE_IS_CIRCULAR",
		"FAILED_ROUTE_OBSERVED_FLAVOR_DATA_CANNOT_ORIENT_EXTERNAL_C3",
		"FAILED_ROUTE_NO_INDIVIDUAL_YUKAWA_VALUES",
	}
}

func CandidateSupports(candidates []ConstructionCandidate) []string {
	var out []string
	for _, c := range candidates {
		out = append(out, c.Supports...)
	}
	return out
}

func CandidateFailures(candidates []ConstructionCandidate) []string {
	var out []string
	for _, c := range candidates {
		out = append(out, c.Firewalls...)
	}
	return out
}

func CandidateNotes(candidates []ConstructionCandidate) []string {
	out := make([]string, 0, len(candidates))
	for _, c := range candidates {
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
