// Package generation2externalflavororientationsealinstallationaudit implements
// Gate 964: ExternalFlavorOrientationSeal Installation Audit.
//
// Gate 964 follows Gate 963's decision that no CanonicalFlavorSelector exists in
// the current sealed R4 certificate. It installs an explicit non-native flavor
// orientation seal so downstream flavor-ledger diagnostics may proceed under
// quarantine while preserving all native-theorem firewalls.
package generation2externalflavororientationsealinstallationaudit

import (
	"fmt"
	"strings"
)

const (
	AuditID         = "GATE964-GENERATION2EXTERNALFLAVORORIENTATIONSEALINSTALLATIONAUDIT"
	InheritedStatus = "R4_REQUIRES_EXTERNAL_FLAVOR_ORIENTATION_SEAL"
	Verdict         = "EXTERNAL_FLAVOR_ORIENTATION_SEAL_INSTALLED_DOWNSTREAM_FLAVOR_LEDGER_TESTS_ALLOWED_SEALED_NOT_NATIVE"
	Classification  = "R4_EXTERNAL_FLAVOR_ORIENTATION_SEALED_NO_NATIVE_FLAVOR_THEOREM"
	ShortStatus     = "R4_SEALED_FLAVOR_ORIENTATION_AVAILABLE_FOR_LEDGER_TESTS"
	NextGate        = "NEXT_GATE965_FLAVOR_LEDGER_DIAGNOSTIC_PRETEST_UNDER_TRIPLE_SEAL"
)

type SealStatus string

const (
	StatusInstalled       SealStatus = "EXTERNAL_FLAVOR_ORIENTATION_SEAL_INSTALLED"
	StatusInherited       SealStatus = "INHERITED_SEAL_PRESERVED"
	StatusDownstreamOnly  SealStatus = "DOWNSTREAM_DIAGNOSTIC_PERMISSION_ONLY"
	StatusNativeForbidden SealStatus = "NATIVE_FLAVOR_CLAIM_FORBIDDEN"
)

type SealComponent struct {
	Name                         string
	Role                         string
	Status                       SealStatus
	Installed                    bool
	Preserved                    bool
	Native                       bool
	SelectsRepresentative        bool
	AllowsDownstreamDiagnostics  bool
	DerivesYukawaEigenvalues     bool
	DerivesCKMPMNS               bool
	AssignsPhysicalParticles     bool
	UpdatesOfficialLedger        bool
	PromotesGenerationNative     bool
	PromotesR3Native             bool
	BreaksU3GaugeNatively        bool
	UsesObservedFlavorDataSource bool
	Supports                     []string
	Firewalls                    []string
}

type Decision struct {
	InheritedRequiresExternalFlavorOrientationSeal bool
	ExternalFlavorOrientationSealInstalled         bool
	ExternalFlavorOrientationSealNative            bool
	RepresentativeChosenForDiagnostics             bool
	U3OrbitAcknowledged                            bool
	CanonicalFlavorSelectorCertified               bool
	DownstreamFlavorLedgerTestsAllowed             bool
	R3DualSealPreserved                            bool
	ExternalGenerationCarrierSealPreserved         bool
	TripleSealLaneActive                           bool
	NativeFlavorTheoremCertified                   bool
	YukawaEigenvaluesDerived                       bool
	CKMPMNSDerived                                 bool
	PMNSDerived                                    bool
	PhysicalParticlesAssigned                      bool
	OfficialLedgerUpdateAllowed                    bool
}

type Analysis struct {
	AuditID        string
	Inherited      string
	Verdict        string
	Classification string
	ShortStatus    string
	NextGate       string
	InstalledSeal  string
	TripleSealLane string
	Decision       Decision
	Components     []SealComponent
	Allowed        []string
	Forbidden      []string
	Supports       []string
	Failures       []string
	Final          string
}

func BuildDefault() (Analysis, error) {
	components := DefaultComponents()
	if len(components) != 6 {
		return Analysis{}, fmt.Errorf("expected 6 seal components, got %d", len(components))
	}
	decision := Decision{
		InheritedRequiresExternalFlavorOrientationSeal: true,
		ExternalFlavorOrientationSealInstalled:         true,
		ExternalFlavorOrientationSealNative:            false,
		RepresentativeChosenForDiagnostics:             true,
		U3OrbitAcknowledged:                            true,
		CanonicalFlavorSelectorCertified:               false,
		DownstreamFlavorLedgerTestsAllowed:             true,
		R3DualSealPreserved:                            true,
		ExternalGenerationCarrierSealPreserved:         true,
		TripleSealLaneActive:                           true,
		NativeFlavorTheoremCertified:                   false,
		YukawaEigenvaluesDerived:                       false,
		CKMPMNSDerived:                                 false,
		PMNSDerived:                                    false,
		PhysicalParticlesAssigned:                      false,
		OfficialLedgerUpdateAllowed:                    false,
	}
	a := Analysis{
		AuditID:        AuditID,
		Inherited:      InheritedStatus,
		Verdict:        Verdict,
		Classification: Classification,
		ShortStatus:    ShortStatus,
		NextGate:       NextGate,
		InstalledSeal:  "ExternalFlavorOrientationSeal selecting Phi_flav^seal in [Phi_flav]_{U(3)} for downstream flavor-ledger diagnostics only",
		TripleSealLane: "R3DualSeal + ExternalGenerationCarrierSeal(C3) + ExternalFlavorOrientationSeal",
		Decision:       decision,
		Components:     components,
		Allowed:        AllowedDiagnostics(),
		Forbidden:      ForbiddenClaims(),
		Supports:       Supports(),
		Failures:       Failures(),
		Final:          "Gate 964 installs the ExternalFlavorOrientationSeal required by Gate 963. It chooses a representative of the U(3) flavor-orientation orbit only for downstream sealed diagnostics. The result activates the triple-sealed R4 lane: R3DualSeal, ExternalGenerationCarrierSeal(C3), and ExternalFlavorOrientationSeal. It does not certify a native flavor theorem, derive Yukawa eigenvalues, CKM/PMNS, particle assignments, or permit official ledger updates.",
	}
	if err := validate(a); err != nil {
		return Analysis{}, err
	}
	return a, nil
}

func validate(a Analysis) error {
	if a.Inherited != InheritedStatus || a.Verdict != Verdict || a.Classification != Classification || a.ShortStatus != ShortStatus {
		return fmt.Errorf("bad Gate 964 identity")
	}
	if !a.Decision.InheritedRequiresExternalFlavorOrientationSeal || !a.Decision.ExternalFlavorOrientationSealInstalled || a.Decision.ExternalFlavorOrientationSealNative {
		return fmt.Errorf("Gate 964 must install non-native external flavor orientation seal: %#v", a.Decision)
	}
	if !a.Decision.RepresentativeChosenForDiagnostics || !a.Decision.U3OrbitAcknowledged || a.Decision.CanonicalFlavorSelectorCertified {
		return fmt.Errorf("Gate 964 must select only a diagnostic representative of the U3 orbit, not a canonical selector: %#v", a.Decision)
	}
	if !a.Decision.DownstreamFlavorLedgerTestsAllowed || !a.Decision.TripleSealLaneActive {
		return fmt.Errorf("Gate 964 must enable downstream tests only under triple seal: %#v", a.Decision)
	}
	if !a.Decision.R3DualSealPreserved || !a.Decision.ExternalGenerationCarrierSealPreserved {
		return fmt.Errorf("Gate 964 must preserve inherited seals: %#v", a.Decision)
	}
	if a.Decision.NativeFlavorTheoremCertified || a.Decision.YukawaEigenvaluesDerived || a.Decision.CKMPMNSDerived || a.Decision.PMNSDerived || a.Decision.PhysicalParticlesAssigned || a.Decision.OfficialLedgerUpdateAllowed {
		return fmt.Errorf("Gate 964 overclaimed downstream flavor theorem/value: %#v", a.Decision)
	}
	for _, c := range a.Components {
		if !c.Installed && !c.Preserved && c.Status != StatusDownstreamOnly && c.Status != StatusNativeForbidden {
			return fmt.Errorf("component neither installed nor preserved: %#v", c)
		}
		if c.Native || c.DerivesYukawaEigenvalues || c.DerivesCKMPMNS || c.AssignsPhysicalParticles || c.UpdatesOfficialLedger || c.PromotesGenerationNative || c.PromotesR3Native || c.BreaksU3GaugeNatively || c.UsesObservedFlavorDataSource {
			return fmt.Errorf("component overclaimed native/downstream result: %#v", c)
		}
	}
	return nil
}

func DefaultComponents() []SealComponent {
	return []SealComponent{
		{
			Name:                        "ExternalFlavorOrientationSeal",
			Role:                        "choose Phi_flav^seal in [Phi_flav]_{U(3)} for downstream diagnostics only",
			Status:                      StatusInstalled,
			Installed:                   true,
			Native:                      false,
			SelectsRepresentative:       true,
			AllowsDownstreamDiagnostics: true,
			Supports: []string{
				"CONDITIONAL_SUPPORT_EXTERNAL_FLAVOR_ORIENTATION_SEAL_INSTALLED",
				"CONDITIONAL_SUPPORT_PHI_FLAV_SEAL_SELECTS_REPRESENTATIVE_FOR_DOWNSTREAM_LEDGER_TESTS",
			},
			Firewalls: []string{
				"FAILED_ROUTE_EXTERNAL_FLAVOR_ORIENTATION_SEAL_NOT_NATIVE_FLAVOR_THEOREM",
				"FAILED_ROUTE_EXTERNAL_FLAVOR_ORIENTATION_SEAL_NOT_CANONICAL_FLAVOR_SELECTOR",
			},
		},
		{
			Name:      "R3DualSeal",
			Role:      "preserve scalar-source and post-orientation seals inherited from R3 tracebridge",
			Status:    StatusInherited,
			Preserved: true,
			Native:    false,
			Supports: []string{
				"CONDITIONAL_SUPPORT_R3_DUALSEAL_PRESERVED_UNDER_FLAVOR_ORIENTATION_SEAL",
			},
			Firewalls: []string{
				"FAILED_ROUTE_R3_TRACEBRIDGE_REMAINS_DUALSEALED_NOT_NATIVE",
				"FAILED_ROUTE_R3_TRACEBRIDGE_REMAINS_SCALAR_SOURCE_SEALED",
				"FAILED_ROUTE_R3_TRACEBRIDGE_REMAINS_POST_ORIENTATION_SEALED",
			},
		},
		{
			Name:      "ExternalGenerationCarrierSeal(C3)",
			Role:      "preserve sealed C^3 family-slot carrier; do not promote generation multiplicity to native",
			Status:    StatusInherited,
			Preserved: true,
			Native:    false,
			Supports: []string{
				"CONDITIONAL_SUPPORT_EXTERNAL_GENERATION_CARRIER_SEAL_C3_PRESERVED",
			},
			Firewalls: []string{
				"FAILED_ROUTE_EXTERNAL_C3_NOT_NATIVE_GENERATION_CARRIER",
				"FAILED_ROUTE_GENERATION_MULTIPLICITY_REMAINS_EXTERNAL_SEAL",
			},
		},
		{
			Name:                        "downstream flavor-ledger diagnostics",
			Role:                        "allow sealed diagnostics such as epsilon_e/kappa_e consistency, Koide-shadow compatibility, CKM/PMNS ledger compatibility",
			Status:                      StatusDownstreamOnly,
			AllowsDownstreamDiagnostics: true,
			Supports: []string{
				"CONDITIONAL_SUPPORT_DOWNSTREAM_FLAVOR_LEDGER_TESTS_ALLOWED_UNDER_TRIPLE_SEAL",
				"CONDITIONAL_SUPPORT_FLAVOR_LEDGER_DIAGNOSTICS_MAY_PROCEED_ONLY_AS_SEALED_TESTS",
			},
			Firewalls: []string{
				"FAILED_ROUTE_DOWNSTREAM_LEDGER_TESTS_ARE_NOT_NATIVE_DERIVATIONS",
			},
		},
		{
			Name:                     "forbidden native flavor claims",
			Role:                     "block Yukawa spectrum, CKM/PMNS, mass hierarchy, particle assignment, and ledger updates",
			Status:                   StatusNativeForbidden,
			DerivesYukawaEigenvalues: false,
			DerivesCKMPMNS:           false,
			AssignsPhysicalParticles: false,
			UpdatesOfficialLedger:    false,
			Firewalls: []string{
				"FAILED_ROUTE_NO_INDIVIDUAL_YUKAWA_VALUES",
				"FAILED_ROUTE_NO_PHYSICAL_PARTICLE_ASSIGNMENT",
				"FAILED_ROUTE_NO_CKM_PMNS_THEOREM",
				"FAILED_ROUTE_NO_PMNS_THEOREM",
				"FAILED_ROUTE_NO_R4_NATIVE_YUKAWA_SPECTRUM_THEOREM",
				"FAILED_ROUTE_NO_OFFICIAL_LEDGER_UPDATE_ALLOWED",
			},
		},
		{
			Name:                         "noncircular source firewall",
			Role:                         "prevent observed masses, CKM/PMNS, epsilon_e, kappa_e, Koide branch, or flavor-wall backsolve from installing the seal as native",
			Status:                       StatusNativeForbidden,
			UsesObservedFlavorDataSource: false,
			Supports: []string{
				"CONDITIONAL_SUPPORT_EXTERNAL_ORIENTATION_SEAL_IS_EXPLICIT_QUARANTINE_NOT_BACKSOLVE",
			},
			Firewalls: []string{
				"FAILED_ROUTE_FLAVOR_FORMULA_BACKSOLVE_IS_CIRCULAR",
				"FAILED_ROUTE_OBSERVED_FLAVOR_DATA_CANNOT_SOURCE_EXTERNAL_FLAVOR_ORIENTATION_SEAL",
				"FAILED_ROUTE_EPSILON_E_KAPPA_E_KOIDE_CKM_PMNS_REMAIN_DOWNSTREAM_TARGETS_ONLY",
			},
		},
	}
}

func AllowedDiagnostics() []string {
	return []string{
		"sealed epsilon_e ledger consistency test",
		"sealed kappa_e ledger consistency test",
		"sealed Koide-shadow compatibility audit",
		"sealed CKM/PMNS ledger compatibility audit",
		"sealed flavor-wall residual consistency audit",
	}
}

func ForbiddenClaims() []string {
	return []string{
		"native flavor theorem",
		"individual Yukawa eigenvalues",
		"physical particle assignment",
		"observed mass hierarchy derivation",
		"CKM/PMNS theorem",
		"official ledger update",
	}
}

func Supports() []string {
	return []string{
		"CONDITIONAL_SUPPORT_EXTERNAL_FLAVOR_ORIENTATION_SEAL_INSTALLED",
		"CONDITIONAL_SUPPORT_PHI_FLAV_SEAL_SELECTS_REPRESENTATIVE_FOR_DOWNSTREAM_LEDGER_TESTS",
		"CONDITIONAL_SUPPORT_R3_DUALSEAL_EXTERNAL_GENERATION_SEAL_AND_EXTERNAL_FLAVOR_ORIENTATION_SEAL_ACTIVE",
		"CONDITIONAL_SUPPORT_DOWNSTREAM_FLAVOR_LEDGER_TESTS_ALLOWED_UNDER_TRIPLE_SEAL",
		"CONDITIONAL_SUPPORT_FLAVOR_LEDGER_DIAGNOSTICS_MAY_PROCEED_ONLY_AS_SEALED_TESTS",
		"CONDITIONAL_SUPPORT_EXTERNAL_ORIENTATION_SEAL_IS_EXPLICIT_QUARANTINE_NOT_BACKSOLVE",
	}
}

func Failures() []string {
	return []string{
		"FAILED_ROUTE_EXTERNAL_FLAVOR_ORIENTATION_SEAL_NOT_NATIVE_FLAVOR_THEOREM",
		"FAILED_ROUTE_EXTERNAL_FLAVOR_ORIENTATION_SEAL_NOT_CANONICAL_FLAVOR_SELECTOR",
		"FAILED_ROUTE_NO_CANONICAL_FLAVOR_SELECTOR_IN_CURRENT_CERTIFICATE",
		"FAILED_ROUTE_U3_FAMILY_GAUGE_NOT_BROKEN_BY_CURRENT_ASHA_DATA",
		"FAILED_ROUTE_EXTERNAL_C3_NOT_NATIVE_GENERATION_CARRIER",
		"FAILED_ROUTE_R3_TRACEBRIDGE_REMAINS_DUALSEALED_NOT_NATIVE",
		"FAILED_ROUTE_DOWNSTREAM_LEDGER_TESTS_ARE_NOT_NATIVE_DERIVATIONS",
		"FAILED_ROUTE_FLAVOR_FORMULA_BACKSOLVE_IS_CIRCULAR",
		"FAILED_ROUTE_OBSERVED_FLAVOR_DATA_CANNOT_SOURCE_EXTERNAL_FLAVOR_ORIENTATION_SEAL",
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
		"CONDITIONAL_SUPPORT_EXTERNAL_FLAVOR_ORIENTATION_SEAL_INSTALLED",
		"CONDITIONAL_SUPPORT_PHI_FLAV_SEAL_SELECTS_REPRESENTATIVE_FOR_DOWNSTREAM_LEDGER_TESTS",
		"CONDITIONAL_SUPPORT_DOWNSTREAM_FLAVOR_LEDGER_TESTS_ALLOWED_UNDER_TRIPLE_SEAL",
	}
}

func RequiredFailures() []string {
	return []string{
		"FAILED_ROUTE_EXTERNAL_FLAVOR_ORIENTATION_SEAL_NOT_NATIVE_FLAVOR_THEOREM",
		"FAILED_ROUTE_EXTERNAL_FLAVOR_ORIENTATION_SEAL_NOT_CANONICAL_FLAVOR_SELECTOR",
		"FAILED_ROUTE_NO_INDIVIDUAL_YUKAWA_VALUES",
		"FAILED_ROUTE_NO_PHYSICAL_PARTICLE_ASSIGNMENT",
		"FAILED_ROUTE_NO_CKM_PMNS_THEOREM",
		"FAILED_ROUTE_NO_OFFICIAL_LEDGER_UPDATE_ALLOWED",
	}
}

func ComponentSupports(components []SealComponent) []string {
	var out []string
	for _, c := range components {
		out = append(out, c.Supports...)
	}
	return out
}

func ComponentFailures(components []SealComponent) []string {
	var out []string
	for _, c := range components {
		out = append(out, c.Firewalls...)
	}
	return out
}

func ComponentNotes(components []SealComponent) []string {
	out := make([]string, 0, len(components))
	for _, c := range components {
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
