// Package generation2flavorledgerdiagnosticpretestundertripleseal implements
// Gate 965: FlavorLedger Diagnostic Pretest Under Triple Seal.
//
// Gate 965 follows Gate 964's installation of ExternalFlavorOrientationSeal. It
// opens downstream flavor-ledger diagnostics only under the visible quarantine of
// R3DualSeal, ExternalGenerationCarrierSeal(C3), and ExternalFlavorOrientationSeal.
// It explicitly does not derive flavor, individual Yukawa values, CKM/PMNS,
// physical particle assignments, or official ledger updates.
package generation2flavorledgerdiagnosticpretestundertripleseal

import (
	"fmt"
	"strings"
)

const (
	AuditID         = "GATE965-GENERATION2FLAVORLEDGERDIAGNOSTICPRETESTUNDERTRIPLESEAL"
	InheritedStatus = "R4_SEALED_FLAVOR_ORIENTATION_AVAILABLE_FOR_LEDGER_TESTS"
	Verdict         = "FLAVOR_LEDGER_DIAGNOSTICS_ARE_ADMISSIBLE_UNDER_TRIPLE_SEAL_BUT_REMAIN_NON_NATIVE_TESTS"
	Classification  = "R4_TRIPLE_SEALED_FLAVOR_LEDGER_DIAGNOSTIC_PRETEST"
	ShortStatus     = "R4_FLAVOR_LEDGER_TESTS_ALLOWED_SEALED_NOT_NATIVE"
	NextGate        = "NEXT_GATE966_SEALED_EPSILON_E_FLAVOR_WALL_LEDGER_CONSISTENCY_AUDIT"
)

type DiagnosticKind string

const (
	KindEpsilon    DiagnosticKind = "EPSILON_E_LEDGER_DIAGNOSTIC"
	KindKappa      DiagnosticKind = "KAPPA_LEDGER_DIAGNOSTIC"
	KindKoide      DiagnosticKind = "KOIDE_SHADOW_COMPATIBILITY_DIAGNOSTIC"
	KindCKMPMNS    DiagnosticKind = "CKM_PMNS_LEDGER_COMPATIBILITY_DIAGNOSTIC"
	KindOfficialFW DiagnosticKind = "OFFICIAL_LEDGER_FIREWALL"
)

type Diagnostic struct {
	Name                          string
	Kind                          DiagnosticKind
	AllowedAsSealedDiagnostic     bool
	UsedAsNativeTheorem           bool
	UsedAsGenerationSource        bool
	UsedAsFlavorOrientationSource bool
	DerivesIndividualYukawas      bool
	DerivesCKMPMNS                bool
	DerivesPMNS                   bool
	AssignsPhysicalParticles      bool
	UpdatesOfficialLedger         bool
	Supports                      []string
	Firewalls                     []string
}

type Decision struct {
	InheritedTripleSealActive              bool
	R3DualSealPreserved                    bool
	ScalarSourceSealPreserved              bool
	PostOrientationSealPreserved           bool
	ExternalGenerationCarrierSealPreserved bool
	ExternalFlavorOrientationSealPreserved bool
	EpsilonLedgerDiagnosticAllowed         bool
	KappaLedgerDiagnosticAllowed           bool
	KoideShadowDiagnosticAllowed           bool
	CKMPMNSLedgerDiagnosticAllowed         bool
	AllDiagnosticsSealedOnly               bool
	NativeFlavorTheoremDerived             bool
	YukawaSpectrumDerived                  bool
	CKMPNMSTheoremDerived                  bool
	PMNSTheoremDerived                     bool
	PhysicalParticlesAssigned              bool
	OfficialLedgerUpdateAllowed            bool
}

type Analysis struct {
	AuditID        string
	Inherited      string
	Verdict        string
	Classification string
	ShortStatus    string
	NextGate       string
	SealLane       string
	Decision       Decision
	Diagnostics    []Diagnostic
	Allowed        []string
	Forbidden      []string
	Supports       []string
	Failures       []string
	Final          string
}

func BuildDefault() (Analysis, error) {
	diagnostics := DefaultDiagnostics()
	if len(diagnostics) != 5 {
		return Analysis{}, fmt.Errorf("expected 5 diagnostics/firewall entries, got %d", len(diagnostics))
	}
	decision := Decision{
		InheritedTripleSealActive:              true,
		R3DualSealPreserved:                    true,
		ScalarSourceSealPreserved:              true,
		PostOrientationSealPreserved:           true,
		ExternalGenerationCarrierSealPreserved: true,
		ExternalFlavorOrientationSealPreserved: true,
		EpsilonLedgerDiagnosticAllowed:         true,
		KappaLedgerDiagnosticAllowed:           true,
		KoideShadowDiagnosticAllowed:           true,
		CKMPMNSLedgerDiagnosticAllowed:         true,
		AllDiagnosticsSealedOnly:               true,
		NativeFlavorTheoremDerived:             false,
		YukawaSpectrumDerived:                  false,
		CKMPNMSTheoremDerived:                  false,
		PMNSTheoremDerived:                     false,
		PhysicalParticlesAssigned:              false,
		OfficialLedgerUpdateAllowed:            false,
	}
	a := Analysis{
		AuditID:        AuditID,
		Inherited:      InheritedStatus,
		Verdict:        Verdict,
		Classification: Classification,
		ShortStatus:    ShortStatus,
		NextGate:       NextGate,
		SealLane:       "R3DualSeal + ScalarSourceSeal(S_split) + PostOrientationSeal(A_F^orient) + ExternalGenerationCarrierSeal(C3) + ExternalFlavorOrientationSeal",
		Decision:       decision,
		Diagnostics:    diagnostics,
		Allowed:        AllowedDiagnostics(),
		Forbidden:      ForbiddenClaims(),
		Supports:       Supports(),
		Failures:       Failures(),
		Final:          "Gate 965 opens the downstream flavor-ledger diagnostic board under explicit triple seal. Epsilon_e, kappa ledgers, Koide-shadow compatibility, and CKM/PMNS ledger compatibility are admissible only as sealed diagnostics. None of them may source generation, source flavor orientation, derive individual Yukawa values, assign particles, prove CKM/PMNS, or update official ledgers.",
	}
	if err := validate(a); err != nil {
		return Analysis{}, err
	}
	return a, nil
}

func validate(a Analysis) error {
	if a.Inherited != InheritedStatus || a.Verdict != Verdict || a.Classification != Classification || a.ShortStatus != ShortStatus {
		return fmt.Errorf("bad Gate 965 identity")
	}
	if !a.Decision.InheritedTripleSealActive || !a.Decision.R3DualSealPreserved || !a.Decision.ScalarSourceSealPreserved || !a.Decision.PostOrientationSealPreserved || !a.Decision.ExternalGenerationCarrierSealPreserved || !a.Decision.ExternalFlavorOrientationSealPreserved {
		return fmt.Errorf("Gate 965 must preserve all inherited seals: %#v", a.Decision)
	}
	if !a.Decision.EpsilonLedgerDiagnosticAllowed || !a.Decision.KappaLedgerDiagnosticAllowed || !a.Decision.KoideShadowDiagnosticAllowed || !a.Decision.CKMPMNSLedgerDiagnosticAllowed || !a.Decision.AllDiagnosticsSealedOnly {
		return fmt.Errorf("Gate 965 must allow diagnostics only under seal: %#v", a.Decision)
	}
	if a.Decision.NativeFlavorTheoremDerived || a.Decision.YukawaSpectrumDerived || a.Decision.CKMPNMSTheoremDerived || a.Decision.PMNSTheoremDerived || a.Decision.PhysicalParticlesAssigned || a.Decision.OfficialLedgerUpdateAllowed {
		return fmt.Errorf("Gate 965 overclaimed flavor theorem/value: %#v", a.Decision)
	}
	for _, d := range a.Diagnostics {
		if d.Kind != KindOfficialFW && !d.AllowedAsSealedDiagnostic {
			return fmt.Errorf("diagnostic not allowed as sealed test: %#v", d)
		}
		if d.UsedAsNativeTheorem || d.UsedAsGenerationSource || d.UsedAsFlavorOrientationSource || d.DerivesIndividualYukawas || d.DerivesCKMPMNS || d.DerivesPMNS || d.AssignsPhysicalParticles || d.UpdatesOfficialLedger {
			return fmt.Errorf("diagnostic overclaimed native/source role: %#v", d)
		}
	}
	return nil
}

func DefaultDiagnostics() []Diagnostic {
	return []Diagnostic{
		{
			Name:                      "epsilon_e ledger consistency",
			Kind:                      KindEpsilon,
			AllowedAsSealedDiagnostic: true,
			Supports: []string{
				"CONDITIONAL_SUPPORT_EPSILON_E_LEDGER_TEST_ALLOWED_UNDER_TRIPLE_SEAL",
			},
			Firewalls: []string{
				"FAILED_ROUTE_EPSILON_E_NOT_NATIVE_FLAVOR_THEOREM",
				"FAILED_ROUTE_EPSILON_E_CANNOT_SOURCE_GENERATION_OR_ORIENTATION",
			},
		},
		{
			Name:                      "kappa_e and kappa_lambda ledger consistency",
			Kind:                      KindKappa,
			AllowedAsSealedDiagnostic: true,
			Supports: []string{
				"CONDITIONAL_SUPPORT_KAPPA_LEDGER_TEST_ALLOWED_UNDER_TRIPLE_SEAL",
			},
			Firewalls: []string{
				"FAILED_ROUTE_KAPPA_E_CANNOT_SOURCE_FLAVOR_ORIENTATION",
				"FAILED_ROUTE_KAPPA_LEDGER_NOT_NATIVE_YUKAWA_SPECTRUM",
			},
		},
		{
			Name:                      "Koide-shadow compatibility",
			Kind:                      KindKoide,
			AllowedAsSealedDiagnostic: true,
			Supports: []string{
				"CONDITIONAL_SUPPORT_KOIDE_SHADOW_COMPATIBILITY_TEST_ALLOWED_UNDER_TRIPLE_SEAL",
			},
			Firewalls: []string{
				"FAILED_ROUTE_KOIDE_SHADOW_NOT_NATIVE_CHARGED_LEPTON_THEOREM",
				"FAILED_ROUTE_KOIDE_BRANCH_CANNOT_SOURCE_FLAVOR_ORIENTATION",
			},
		},
		{
			Name:                      "CKM/PMNS ledger compatibility",
			Kind:                      KindCKMPMNS,
			AllowedAsSealedDiagnostic: true,
			Supports: []string{
				"CONDITIONAL_SUPPORT_CKM_PMNS_LEDGER_COMPATIBILITY_TEST_ALLOWED_UNDER_TRIPLE_SEAL",
			},
			Firewalls: []string{
				"FAILED_ROUTE_CKM_PMNS_LEDGER_NOT_NATIVE_MIXING_THEOREM",
				"FAILED_ROUTE_CKM_PMNS_CANNOT_SOURCE_FLAVOR_ORIENTATION",
			},
		},
		{
			Name:                      "official ledger and native theorem firewall",
			Kind:                      KindOfficialFW,
			AllowedAsSealedDiagnostic: false,
			Supports: []string{
				"CONDITIONAL_SUPPORT_FLAVOR_LEDGER_DIAGNOSTIC_BOARD_OPENED_UNDER_TRIPLE_SEAL",
			},
			Firewalls: []string{
				"FAILED_ROUTE_NO_OFFICIAL_LEDGER_UPDATE_ALLOWED",
				"FAILED_ROUTE_DOWNSTREAM_LEDGER_TESTS_ARE_NOT_NATIVE_DERIVATIONS",
				"FAILED_ROUTE_NO_R4_NATIVE_YUKAWA_SPECTRUM_THEOREM",
			},
		},
	}
}

func AllowedDiagnostics() []string {
	return []string{
		"epsilon_e ledger consistency as sealed downstream target",
		"kappa_e and kappa_lambda ledger consistency as sealed downstream targets",
		"Koide-shadow compatibility as sealed diagnostic",
		"CKM/PMNS ledger compatibility as sealed diagnostic",
		"flavor-wall residual consistency as sealed diagnostic",
	}
}

func ForbiddenClaims() []string {
	return []string{
		"epsilon_e/kappa_e as source of generation or orientation",
		"native flavor theorem",
		"native charged-lepton theorem",
		"individual Yukawa eigenvalues",
		"physical particle assignment",
		"CKM/PMNS theorem",
		"official ledger update",
	}
}

func Supports() []string {
	return []string{
		"CONDITIONAL_SUPPORT_EPSILON_E_LEDGER_TEST_ALLOWED_UNDER_TRIPLE_SEAL",
		"CONDITIONAL_SUPPORT_KAPPA_LEDGER_TEST_ALLOWED_UNDER_TRIPLE_SEAL",
		"CONDITIONAL_SUPPORT_KOIDE_SHADOW_COMPATIBILITY_TEST_ALLOWED_UNDER_TRIPLE_SEAL",
		"CONDITIONAL_SUPPORT_CKM_PMNS_LEDGER_COMPATIBILITY_TEST_ALLOWED_UNDER_TRIPLE_SEAL",
		"CONDITIONAL_SUPPORT_FLAVOR_LEDGER_DIAGNOSTIC_BOARD_OPENED_UNDER_TRIPLE_SEAL",
		"CONDITIONAL_SUPPORT_ALL_DIAGNOSTIC_OUTPUTS_INHERIT_R3_DUALSEAL_EXTERNAL_C3_AND_EXTERNAL_FLAVOR_ORIENTATION_SEAL",
	}
}

func Failures() []string {
	return []string{
		"FAILED_ROUTE_EPSILON_E_NOT_NATIVE_FLAVOR_THEOREM",
		"FAILED_ROUTE_EPSILON_E_CANNOT_SOURCE_GENERATION_OR_ORIENTATION",
		"FAILED_ROUTE_KAPPA_E_CANNOT_SOURCE_FLAVOR_ORIENTATION",
		"FAILED_ROUTE_KAPPA_LEDGER_NOT_NATIVE_YUKAWA_SPECTRUM",
		"FAILED_ROUTE_KOIDE_SHADOW_NOT_NATIVE_CHARGED_LEPTON_THEOREM",
		"FAILED_ROUTE_KOIDE_BRANCH_CANNOT_SOURCE_FLAVOR_ORIENTATION",
		"FAILED_ROUTE_CKM_PMNS_LEDGER_NOT_NATIVE_MIXING_THEOREM",
		"FAILED_ROUTE_CKM_PMNS_CANNOT_SOURCE_FLAVOR_ORIENTATION",
		"FAILED_ROUTE_NO_OFFICIAL_LEDGER_UPDATE_ALLOWED",
		"FAILED_ROUTE_DOWNSTREAM_LEDGER_TESTS_ARE_NOT_NATIVE_DERIVATIONS",
		"FAILED_ROUTE_NO_R4_NATIVE_YUKAWA_SPECTRUM_THEOREM",
	}
}

func RequiredSupports() []string {
	return []string{
		"CONDITIONAL_SUPPORT_EPSILON_E_LEDGER_TEST_ALLOWED_UNDER_TRIPLE_SEAL",
		"CONDITIONAL_SUPPORT_KAPPA_LEDGER_TEST_ALLOWED_UNDER_TRIPLE_SEAL",
		"CONDITIONAL_SUPPORT_KOIDE_SHADOW_COMPATIBILITY_TEST_ALLOWED_UNDER_TRIPLE_SEAL",
		"CONDITIONAL_SUPPORT_CKM_PMNS_LEDGER_COMPATIBILITY_TEST_ALLOWED_UNDER_TRIPLE_SEAL",
	}
}

func RequiredFailures() []string {
	return []string{
		"FAILED_ROUTE_EPSILON_E_NOT_NATIVE_FLAVOR_THEOREM",
		"FAILED_ROUTE_EPSILON_E_CANNOT_SOURCE_GENERATION_OR_ORIENTATION",
		"FAILED_ROUTE_KAPPA_E_CANNOT_SOURCE_FLAVOR_ORIENTATION",
		"FAILED_ROUTE_KOIDE_SHADOW_NOT_NATIVE_CHARGED_LEPTON_THEOREM",
		"FAILED_ROUTE_CKM_PMNS_LEDGER_NOT_NATIVE_MIXING_THEOREM",
		"FAILED_ROUTE_NO_OFFICIAL_LEDGER_UPDATE_ALLOWED",
		"FAILED_ROUTE_DOWNSTREAM_LEDGER_TESTS_ARE_NOT_NATIVE_DERIVATIONS",
		"FAILED_ROUTE_NO_R4_NATIVE_YUKAWA_SPECTRUM_THEOREM",
	}
}

func DiagnosticSupports(diagnostics []Diagnostic) []string {
	var out []string
	for _, d := range diagnostics {
		out = append(out, d.Supports...)
	}
	return out
}

func DiagnosticFailures(diagnostics []Diagnostic) []string {
	var out []string
	for _, d := range diagnostics {
		out = append(out, d.Firewalls...)
	}
	return out
}

func DiagnosticNotes(diagnostics []Diagnostic) []string {
	out := make([]string, 0, len(diagnostics))
	for _, d := range diagnostics {
		out = append(out, d.Name+" => "+string(d.Kind))
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
