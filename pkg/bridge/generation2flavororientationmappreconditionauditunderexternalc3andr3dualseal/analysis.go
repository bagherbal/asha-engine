// Package generation2flavororientationmappreconditionauditunderexternalc3andr3dualseal implements
// Gate 961: FlavorOrientationMap Precondition Audit Under ExternalC3 and R3 DualSeal.
//
// Gate 961 follows the external C3 generation-carrier seal. It does not derive
// flavor, particles, CKM/PMNS, or individual Yukawa values. It only types the
// next missing object: a FlavorOrientationMap from the sealed C^3 family-slot
// carrier into a lawful ledger/orientation interface under the inherited R3 dual
// seal and the ExternalGenerationCarrierSeal.
package generation2flavororientationmappreconditionauditunderexternalc3andr3dualseal

import (
	"fmt"
	"strings"
)

const (
	AuditID         = "GATE961-GENERATION2FLAVORORIENTATIONMAPPRECONDITIONAUDITUNDEREXTERNALC3ANDR3DUALSEAL"
	InheritedStatus = "R4_SEALED_GENERATION_CARRIER_AVAILABLE_NO_FLAVOR_MAP"
	Verdict         = "FLAVOR_ORIENTATION_MAP_IDENTIFIED_AS_NEXT_REQUIRED_OBJECT_UNDER_EXTERNAL_C3_AND_R3_DUALSEAL_BUT_NOT_CERTIFIED"
	Classification  = "R4_FLAVOR_ORIENTATION_PRECONDITION_AUDIT_MAP_MISSING"
	ShortStatus     = "R4_SEALED_GENERATION_CARRIER_NO_FLAVOR_ORIENTATION"
	NextGate        = "NEXT_GATE962_FLAVOR_ORIENTATION_MAP_CONSTRUCTION_AUDIT_UNDER_SEALED_R4"
)

type AuditStatus string

const (
	StatusTypedDomain     AuditStatus = "TYPED_DOMAIN_UNDER_SEAL"
	StatusTypedCodomain   AuditStatus = "TYPED_LEDGER_CODOMAIN_CANDIDATE"
	StatusMissingMap      AuditStatus = "MISSING_FLAVOR_ORIENTATION_MAP"
	StatusNonCircularity  AuditStatus = "NONCIRCULARITY_FIREWALL"
	StatusSealInheritance AuditStatus = "SEALS_INHERITED_AND_VISIBLE"
)

type OrientationAudit struct {
	Name                      string
	Object                    string
	Status                    AuditStatus
	AllowedRole               string
	ForbiddenRole             string
	DomainTyped               bool
	CodomainTyped             bool
	MapRequired               bool
	MapCertified              bool
	UsesObservedFlavorData    bool
	UsesFlavorFormulaAsSource bool
	AssignsPhysicalParticles  bool
	ProvidesIndividualYukawa  bool
	ProvidesCKMPMNS           bool
	OfficialLedgerUpdate      bool
	InheritedR3DualSeal       bool
	InheritedExternalC3Seal   bool
	Supports                  []string
	Firewalls                 []string
}

type Decision struct {
	ExternalC3DomainAvailable     bool
	ExternalC3NativeGeneration    bool
	LedgerCodomainTyped           bool
	FlavorOrientationMapRequired  bool
	FlavorOrientationMapCertified bool
	ObservedFlavorInputAllowed    bool
	FlavorFormulaBacksolveAllowed bool
	PhysicalAssignmentAllowed     bool
	IndividualYukawaAllowed       bool
	CKMPMNSAllowed                bool
	OfficialLedgerUpdateAllowed   bool
	R3DualSealPreserved           bool
	ExternalC3SealPreserved       bool
	CanProceedToConstructionAudit bool
}

type Analysis struct {
	AuditID        string
	Inherited      string
	Verdict        string
	Classification string
	ShortStatus    string
	NextGate       string
	Domain         string
	Codomain       string
	RequiredMap    string
	Decision       Decision
	Audits         []OrientationAudit
	Supports       []string
	Failures       []string
	Final          string
}

func BuildDefault() (Analysis, error) {
	audits := DefaultAudits()
	if len(audits) != 5 {
		return Analysis{}, fmt.Errorf("expected 5 orientation audits, got %d", len(audits))
	}
	decision := Decision{
		ExternalC3DomainAvailable:     true,
		ExternalC3NativeGeneration:    false,
		LedgerCodomainTyped:           true,
		FlavorOrientationMapRequired:  true,
		FlavorOrientationMapCertified: false,
		ObservedFlavorInputAllowed:    false,
		FlavorFormulaBacksolveAllowed: false,
		PhysicalAssignmentAllowed:     false,
		IndividualYukawaAllowed:       false,
		CKMPMNSAllowed:                false,
		OfficialLedgerUpdateAllowed:   false,
		R3DualSealPreserved:           true,
		ExternalC3SealPreserved:       true,
		CanProceedToConstructionAudit: true,
	}
	a := Analysis{
		AuditID:        AuditID,
		Inherited:      InheritedStatus,
		Verdict:        Verdict,
		Classification: Classification,
		ShortStatus:    ShortStatus,
		NextGate:       NextGate,
		Domain:         "G_gen^seal = C^3 under ExternalGenerationCarrierSeal(C3)",
		Codomain:       "ledger/interface candidates: A_F^orient, dual-sealed R3 tracebody, socket ledger, flavor-wall ledger targets",
		RequiredMap:    "Phi_flav: C^3_gen,seal -> Orient(A_F^orient, R3_tracebody)",
		Decision:       decision,
		Audits:         audits,
		Supports:       Supports(),
		Failures:       Failures(),
		Final:          "Gate 961 types the FlavorOrientationMap as the next required R4 object under R3DualSeal plus ExternalGenerationCarrierSeal(C3). The sealed C^3 carrier may serve only as a family-slot domain. The codomain may be a ledger/orientation interface, not physical particles, observed mass eigenstates, CKM/PMNS, or individual Yukawa values. No flavor orientation map is certified yet; the next lawful gate is a construction audit under seals.",
	}
	if err := validate(a); err != nil {
		return Analysis{}, err
	}
	return a, nil
}

func validate(a Analysis) error {
	if a.Inherited != InheritedStatus || a.Verdict != Verdict || a.Classification != Classification || a.ShortStatus != ShortStatus {
		return fmt.Errorf("bad Gate 961 identity")
	}
	if !a.Decision.ExternalC3DomainAvailable || a.Decision.ExternalC3NativeGeneration {
		return fmt.Errorf("Gate 961 must use external C3 as sealed domain only: %#v", a.Decision)
	}
	if !a.Decision.LedgerCodomainTyped || !a.Decision.FlavorOrientationMapRequired || a.Decision.FlavorOrientationMapCertified {
		return fmt.Errorf("Gate 961 must type codomain and require, but not certify, Phi_flav: %#v", a.Decision)
	}
	if a.Decision.ObservedFlavorInputAllowed || a.Decision.FlavorFormulaBacksolveAllowed || a.Decision.PhysicalAssignmentAllowed || a.Decision.IndividualYukawaAllowed || a.Decision.CKMPMNSAllowed || a.Decision.OfficialLedgerUpdateAllowed {
		return fmt.Errorf("Gate 961 overclaimed downstream or empirical permissions: %#v", a.Decision)
	}
	if !a.Decision.R3DualSealPreserved || !a.Decision.ExternalC3SealPreserved || !a.Decision.CanProceedToConstructionAudit {
		return fmt.Errorf("Gate 961 must preserve seals and prepare construction audit: %#v", a.Decision)
	}
	if len(a.Audits) != 5 {
		return fmt.Errorf("expected 5 audits, got %d", len(a.Audits))
	}
	for _, au := range a.Audits {
		if au.MapCertified || au.UsesObservedFlavorData || au.UsesFlavorFormulaAsSource || au.AssignsPhysicalParticles || au.ProvidesIndividualYukawa || au.ProvidesCKMPMNS || au.OfficialLedgerUpdate {
			return fmt.Errorf("audit overclaimed flavor/source/downstream status: %#v", au)
		}
	}
	return nil
}

func DefaultAudits() []OrientationAudit {
	return []OrientationAudit{
		{
			Name:                    "domain typing",
			Object:                  "G_gen^seal = C^3",
			Status:                  StatusTypedDomain,
			AllowedRole:             "sealed family-slot carrier domain for a future orientation map",
			ForbiddenRole:           "native generation theorem, observed flavor basis, mass eigenbasis, CKM/PMNS basis",
			DomainTyped:             true,
			InheritedExternalC3Seal: true,
			Supports: []string{
				"CONDITIONAL_SUPPORT_EXTERNAL_C3_CAN_SERVE_AS_DOMAIN_FOR_FLAVOR_ORIENTATION_MAP",
				"CONDITIONAL_SUPPORT_SEALED_GENERATION_CARRIER_CAN_ENTER_ONLY_AS_FAMILY_SLOT_DOMAIN",
			},
			Firewalls: []string{
				"FAILED_ROUTE_EXTERNAL_C3_NOT_NATIVE_GENERATION_CARRIER",
				"FAILED_ROUTE_EXTERNAL_C3_IS_NOT_OBSERVED_FLAVOR_BASIS",
				"FAILED_ROUTE_EXTERNAL_C3_IS_NOT_MASS_EIGENBASIS",
			},
		},
		{
			Name:          "codomain typing",
			Object:        "A_F^orient, R3 aggregate tracebody, socket ledger, flavor-wall ledger targets",
			Status:        StatusTypedCodomain,
			AllowedRole:   "ledger/interface target for future orientation structure",
			ForbiddenRole: "physical particles, observed mass eigenstates, CKM/PMNS matrices, individual Yukawa values",
			CodomainTyped: true,
			Supports: []string{
				"CONDITIONAL_SUPPORT_FLAVOR_ORIENTATION_CODOMAIN_CAN_BE_TYPED_AS_LEDGER_INTERFACE",
				"CONDITIONAL_SUPPORT_A_F_ORIENT_R3_TRACEBODY_AND_SOCKET_LEDGER_ARE_ALLOWED_INTERFACE_TARGETS",
			},
			Firewalls: []string{
				"FAILED_ROUTE_NO_PHYSICAL_PARTICLE_ASSIGNMENT",
				"FAILED_ROUTE_NO_INDIVIDUAL_YUKAWA_VALUES",
				"FAILED_ROUTE_NO_CKM_PMNS_THEOREM",
				"FAILED_ROUTE_NO_OBSERVED_MASS_EIGENSTATE_CODOMAIN_CERTIFIED",
			},
		},
		{
			Name:          "orientation-map requirement",
			Object:        "Phi_flav: C^3_gen,seal -> Orient(A_F^orient, R3_tracebody)",
			Status:        StatusMissingMap,
			AllowedRole:   "family-slot orientation, ledger compatibility, noncircularity, and seal inheritance",
			ForbiddenRole: "numerical spectrum, charged-lepton/quark theorem, CKM/PMNS theorem, particle assignment",
			MapRequired:   true,
			MapCertified:  false,
			Supports: []string{
				"CONDITIONAL_SUPPORT_FLAVOR_ORIENTATION_MAP_IS_NEXT_REQUIRED_OBJECT",
				"CONDITIONAL_SUPPORT_PHI_FLAV_DOMAIN_AND_CODOMAIN_ARE_NOW_TYPED_UNDER_SEALS",
			},
			Firewalls: []string{
				"FAILED_ROUTE_NO_FLAVOR_ORIENTATION_MAP_CERTIFIED_YET",
				"FAILED_ROUTE_GENERATION_CARRIER_NOT_FLAVOR_ORIENTATION",
				"FAILED_ROUTE_NO_R4_NATIVE_YUKAWA_SPECTRUM_THEOREM",
			},
		},
		{
			Name:                      "noncircularity firewall",
			Object:                    "flavor data, CKM/PMNS, epsilon_e, kappa_e, Koide, and flavor-wall formulas",
			Status:                    StatusNonCircularity,
			AllowedRole:               "downstream ledger targets only after a typed orientation map exists",
			ForbiddenRole:             "source or backsolve for Phi_flav",
			UsesObservedFlavorData:    false,
			UsesFlavorFormulaAsSource: false,
			Supports: []string{
				"CONDITIONAL_SUPPORT_FLAVOR_FORMULAS_REMAIN_DOWNSTREAM_TARGETS_ONLY",
				"CONDITIONAL_SUPPORT_OBSERVED_FLAVOR_DATA_EXCLUDED_FROM_ORIENTATION_SOURCE",
			},
			Firewalls: []string{
				"FAILED_ROUTE_FLAVOR_FORMULA_BACKSOLVE_IS_CIRCULAR",
				"FAILED_ROUTE_OBSERVED_FLAVOR_DATA_CANNOT_ORIENT_EXTERNAL_C3",
				"FAILED_ROUTE_EPSILON_E_OR_KAPPA_E_CANNOT_SOURCE_FLAVOR_ORIENTATION",
				"FAILED_ROUTE_KOIDE_BRANCH_CANNOT_SOURCE_FLAVOR_ORIENTATION",
			},
		},
		{
			Name:                    "seal inheritance and next frontier",
			Object:                  "R3DualSeal plus ExternalGenerationCarrierSeal(C3)",
			Status:                  StatusSealInheritance,
			AllowedRole:             "all future R4 flavor work must explicitly inherit seals",
			ForbiddenRole:           "native R3, native generation theorem, official ledger update",
			InheritedR3DualSeal:     true,
			InheritedExternalC3Seal: true,
			Supports: []string{
				"CONDITIONAL_SUPPORT_R3_DUALSEAL_AND_EXTERNAL_C3_SEAL_REMAIN_VISIBLE",
				"CONDITIONAL_SUPPORT_NEXT_LAWFUL_GATE_IS_FLAVOR_ORIENTATION_MAP_CONSTRUCTION_AUDIT_UNDER_SEALED_R4",
			},
			Firewalls: []string{
				"FAILED_ROUTE_R3_TRACEBRIDGE_REMAINS_DUALSEALED_NOT_NATIVE",
				"FAILED_ROUTE_NO_OFFICIAL_LEDGER_UPDATE_ALLOWED",
				"FAILED_ROUTE_EXTERNAL_GENERATION_SEAL_DOES_NOT_REMOVE_R3_DUALSEAL",
			},
		},
	}
}

func Supports() []string {
	return []string{
		"CONDITIONAL_SUPPORT_EXTERNAL_C3_CAN_SERVE_AS_DOMAIN_FOR_FLAVOR_ORIENTATION_MAP",
		"CONDITIONAL_SUPPORT_SEALED_GENERATION_CARRIER_CAN_ENTER_ONLY_AS_FAMILY_SLOT_DOMAIN",
		"CONDITIONAL_SUPPORT_FLAVOR_ORIENTATION_CODOMAIN_CAN_BE_TYPED_AS_LEDGER_INTERFACE",
		"CONDITIONAL_SUPPORT_FLAVOR_ORIENTATION_MAP_IS_NEXT_REQUIRED_OBJECT",
		"CONDITIONAL_SUPPORT_FLAVOR_FORMULAS_REMAIN_DOWNSTREAM_TARGETS_ONLY",
		"CONDITIONAL_SUPPORT_R3_DUALSEAL_AND_EXTERNAL_C3_SEAL_REMAIN_VISIBLE",
	}
}

func Failures() []string {
	return []string{
		"FAILED_ROUTE_EXTERNAL_C3_NOT_NATIVE_GENERATION_CARRIER",
		"FAILED_ROUTE_EXTERNAL_C3_IS_NOT_OBSERVED_FLAVOR_BASIS",
		"FAILED_ROUTE_NO_FLAVOR_ORIENTATION_MAP_CERTIFIED_YET",
		"FAILED_ROUTE_GENERATION_CARRIER_NOT_FLAVOR_ORIENTATION",
		"FAILED_ROUTE_FLAVOR_FORMULA_BACKSOLVE_IS_CIRCULAR",
		"FAILED_ROUTE_OBSERVED_FLAVOR_DATA_CANNOT_ORIENT_EXTERNAL_C3",
		"FAILED_ROUTE_NO_PHYSICAL_PARTICLE_ASSIGNMENT",
		"FAILED_ROUTE_NO_INDIVIDUAL_YUKAWA_VALUES",
		"FAILED_ROUTE_NO_CKM_PMNS_THEOREM",
		"FAILED_ROUTE_NO_PMNS_THEOREM",
		"FAILED_ROUTE_NO_R4_NATIVE_YUKAWA_SPECTRUM_THEOREM",
		"FAILED_ROUTE_R3_TRACEBRIDGE_REMAINS_DUALSEALED_NOT_NATIVE",
		"FAILED_ROUTE_R3_TRACEBRIDGE_REMAINS_SCALAR_SOURCE_SEALED",
		"FAILED_ROUTE_R3_TRACEBRIDGE_REMAINS_POST_ORIENTATION_SEALED",
		"FAILED_ROUTE_NO_OFFICIAL_LEDGER_UPDATE_ALLOWED",
	}
}

func RequiredSupports() []string {
	return []string{
		"CONDITIONAL_SUPPORT_EXTERNAL_C3_CAN_SERVE_AS_DOMAIN_FOR_FLAVOR_ORIENTATION_MAP",
		"CONDITIONAL_SUPPORT_FLAVOR_ORIENTATION_CODOMAIN_CAN_BE_TYPED_AS_LEDGER_INTERFACE",
		"CONDITIONAL_SUPPORT_FLAVOR_ORIENTATION_MAP_IS_NEXT_REQUIRED_OBJECT",
		"CONDITIONAL_SUPPORT_R3_DUALSEAL_AND_EXTERNAL_C3_SEAL_REMAIN_VISIBLE",
	}
}

func RequiredFailures() []string {
	return []string{
		"FAILED_ROUTE_EXTERNAL_C3_NOT_NATIVE_GENERATION_CARRIER",
		"FAILED_ROUTE_NO_FLAVOR_ORIENTATION_MAP_CERTIFIED_YET",
		"FAILED_ROUTE_FLAVOR_FORMULA_BACKSOLVE_IS_CIRCULAR",
		"FAILED_ROUTE_OBSERVED_FLAVOR_DATA_CANNOT_ORIENT_EXTERNAL_C3",
		"FAILED_ROUTE_NO_PHYSICAL_PARTICLE_ASSIGNMENT",
		"FAILED_ROUTE_NO_INDIVIDUAL_YUKAWA_VALUES",
	}
}

func AuditSupports(audits []OrientationAudit) []string {
	var out []string
	for _, au := range audits {
		out = append(out, au.Supports...)
	}
	return out
}

func AuditFailures(audits []OrientationAudit) []string {
	var out []string
	for _, au := range audits {
		out = append(out, au.Firewalls...)
	}
	return out
}

func AuditNotes(audits []OrientationAudit) []string {
	out := make([]string, 0, len(audits))
	for _, au := range audits {
		out = append(out, au.Name+" => "+string(au.Status))
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
