// Package generation2triplesealedflavorledgerdiagnosticsynthesisaudit implements Gate 970: Triple-Sealed Flavor Ledger Diagnostic Synthesis Audit.
//
// This gate continues the sealed R4 flavor/Yukawa rail. It is deliberately
// firewall-preserving: it records what is lawful under explicit seals and what
// remains forbidden as native ASHA theorem, physical-particle assignment, CKM/PMNS
// theorem, or official ledger update.
package generation2triplesealedflavorledgerdiagnosticsynthesisaudit

import (
	"fmt"
	"strings"
)

const (
	AuditID         = "GATE970-GENERATION2-GENERATION2TRIPLESEALEDFLAVORLEDGERDIAGNOSTICSYNTHESISAUDIT"
	InheritedStatus = "R4_CKM_PMNS_COMPATIBILITY_ALLOWED_NO_MIXING_THEOREM"
	Verdict         = "TRIPLE_SEALED_FLAVOR_DIAGNOSTICS_SYNTHESIZED_AS_ALLOWED_TARGETS_YUKAWA_MATRIX_SOURCE_STILL_MISSING"
	Classification  = "R4_FLAVOR_DIAGNOSTIC_SYNTHESIS_YUKAWA_MATRIX_SOURCE_GAP"
	ShortStatus     = "R4_DIAGNOSTICS_SYNTHESIZED_MATRIX_SOURCE_MISSING"
	NextGate        = "NEXT_GATE971_YUKAWA_MATRIX_ADMISSIBILITY_AUDIT_UNDER_TRIPLE_SEAL"
)

type Decision struct {
	InheritedSealedRail                    bool
	R3DualSealPreserved                    bool
	ScalarSourceSealPreserved              bool
	PostOrientationSealPreserved           bool
	ExternalGenerationCarrierSealPreserved bool
	ExternalFlavorOrientationSealPreserved bool
	ExternalYukawaMatrixSealPreserved      bool
	AllowsSealedOperation                  bool
	DerivesNativeFlavor                    bool
	DerivesNativeYukawaMatrix              bool
	DerivesIndividualYukawas               bool
	DerivesCKMPMNS                         bool
	AssignsPhysicalParticles               bool
	UpdatesOfficialLedger                  bool
}

type Analysis struct {
	AuditID          string
	Inherited        string
	Verdict          string
	Classification   string
	ShortStatus      string
	NextGate         string
	SealLane         string
	Decision         Decision
	Allowed          []string
	Forbidden        []string
	Supports         []string
	Failures         []string
	MatrixNormalForm []string
	Final            string
}

func BuildDefault() (Analysis, error) {
	a := Analysis{
		AuditID:        AuditID,
		Inherited:      InheritedStatus,
		Verdict:        Verdict,
		Classification: Classification,
		ShortStatus:    ShortStatus,
		NextGate:       NextGate,
		SealLane:       "R3DualSeal + ScalarSourceSeal(S_split) + PostOrientationSeal(A_F^orient) + ExternalGenerationCarrierSeal(C3) + ExternalFlavorOrientationSeal + optional ExternalYukawaMatrixSeal when installed",
		Decision: Decision{
			InheritedSealedRail:                    true,
			R3DualSealPreserved:                    true,
			ScalarSourceSealPreserved:              true,
			PostOrientationSealPreserved:           true,
			ExternalGenerationCarrierSealPreserved: true,
			ExternalFlavorOrientationSealPreserved: true,
			ExternalYukawaMatrixSealPreserved:      false,
			AllowsSealedOperation:                  true,
			DerivesNativeFlavor:                    false,
			DerivesNativeYukawaMatrix:              false,
			DerivesIndividualYukawas:               false,
			DerivesCKMPMNS:                         false,
			AssignsPhysicalParticles:               false,
			UpdatesOfficialLedger:                  false,
		},
		Allowed:          Allowed(),
		Forbidden:        Forbidden(),
		Supports:         Supports(),
		Failures:         Failures(),
		MatrixNormalForm: MatrixNormalForm(),
		Final:            "Gate 970 concludes R4_DIAGNOSTICS_SYNTHESIZED_MATRIX_SOURCE_MISSING. It preserves all inherited seals and does not derive native flavor, physical particles, CKM/PMNS, individual Yukawa values, official ledgers, or an R4 native Yukawa spectrum theorem.",
	}
	if err := validate(a); err != nil {
		return Analysis{}, err
	}
	return a, nil
}

func validate(a Analysis) error {
	if a.AuditID != AuditID || a.Inherited != InheritedStatus || a.Verdict != Verdict || a.Classification != Classification || a.ShortStatus != ShortStatus || a.NextGate != NextGate {
		return fmt.Errorf("bad Gate 970 identity")
	}
	d := a.Decision
	if !d.InheritedSealedRail || !d.R3DualSealPreserved || !d.ScalarSourceSealPreserved || !d.PostOrientationSealPreserved || !d.ExternalGenerationCarrierSealPreserved || !d.ExternalFlavorOrientationSealPreserved {
		return fmt.Errorf("Gate 970 must preserve inherited seals: %#v", d)
	}
	if !d.AllowsSealedOperation {
		return fmt.Errorf("Gate 970 must allow only the sealed operation under audit")
	}
	if d.DerivesNativeFlavor || d.DerivesNativeYukawaMatrix || d.DerivesIndividualYukawas || d.DerivesCKMPMNS || d.AssignsPhysicalParticles || d.UpdatesOfficialLedger {
		return fmt.Errorf("Gate 970 overclaimed native/physical result: %#v", d)
	}
	if !containsAll(a.Supports, RequiredSupports()) {
		return fmt.Errorf("Gate 970 missing required supports")
	}
	if !containsAll(a.Failures, RequiredFailures()) {
		return fmt.Errorf("Gate 970 missing required firewalls")
	}
	return nil
}

func Allowed() []string {
	return []string{
		"epsilon/kappa/Koide/CKM/PMNS diagnostics may coexist as sealed target ledger",
		"diagnostic suite may constrain consistency of a later sealed Yukawa matrix ledger",
	}
}

func Forbidden() []string {
	return []string{
		"diagnostic suite cannot derive singular values",
		"diagnostic suite cannot derive left/right flavor frames",
		"diagnostic suite cannot become native Yukawa spectrum",
		"diagnostic suite cannot update official ledger",
	}
}

func Supports() []string {
	return []string{
		"CONDITIONAL_SUPPORT_FLAVOR_DIAGNOSTICS_SYNTHESIZED_UNDER_TRIPLE_SEAL",
		"CONDITIONAL_SUPPORT_DIAGNOSTIC_SUITE_CAN_VALIDATE_LATER_SEALED_YUKAWA_MATRIX_LEDGER",
		"CONDITIONAL_SUPPORT_EPSILON_KAPPA_KOIDE_CKM_PMNS_REMAIN_TARGETS_NOT_SOURCES",
	}
}

func Failures() []string {
	return []string{
		"FAILED_ROUTE_DIAGNOSTIC_SUITE_DOES_NOT_DERIVE_YUKAWA_SINGULAR_VALUES",
		"FAILED_ROUTE_DIAGNOSTIC_SUITE_DOES_NOT_DERIVE_LEFT_RIGHT_FLAVOR_FRAMES",
		"FAILED_ROUTE_DIAGNOSTIC_SUITE_NOT_NATIVE_YUKAWA_SPECTRUM",
		"FAILED_ROUTE_YUKAWA_MATRIX_SOURCE_STILL_MISSING",
		"FAILED_ROUTE_NO_OFFICIAL_LEDGER_UPDATE_ALLOWED",
	}
}

func RequiredSupports() []string {
	return []string{
		"CONDITIONAL_SUPPORT_FLAVOR_DIAGNOSTICS_SYNTHESIZED_UNDER_TRIPLE_SEAL",
		"CONDITIONAL_SUPPORT_DIAGNOSTIC_SUITE_CAN_VALIDATE_LATER_SEALED_YUKAWA_MATRIX_LEDGER",
		"CONDITIONAL_SUPPORT_EPSILON_KAPPA_KOIDE_CKM_PMNS_REMAIN_TARGETS_NOT_SOURCES",
	}
}

func RequiredFailures() []string {
	return []string{
		"FAILED_ROUTE_DIAGNOSTIC_SUITE_DOES_NOT_DERIVE_YUKAWA_SINGULAR_VALUES",
		"FAILED_ROUTE_DIAGNOSTIC_SUITE_DOES_NOT_DERIVE_LEFT_RIGHT_FLAVOR_FRAMES",
		"FAILED_ROUTE_DIAGNOSTIC_SUITE_NOT_NATIVE_YUKAWA_SPECTRUM",
		"FAILED_ROUTE_YUKAWA_MATRIX_SOURCE_STILL_MISSING",
		"FAILED_ROUTE_NO_OFFICIAL_LEDGER_UPDATE_ALLOWED",
	}
}

func MatrixNormalForm() []string {
	return []string{
		"no matrix normal form is certified at this gate; this gate only sharpens the lawful sealed diagnostic boundary",
	}
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
