// Package generation2sealedkappaflavorledgerconsistencyaudit implements Gate 967: Sealed Kappa Flavor Ledger Consistency Audit.
//
// This gate continues the sealed R4 flavor/Yukawa rail. It is deliberately
// firewall-preserving: it records what is lawful under explicit seals and what
// remains forbidden as native ASHA theorem, physical-particle assignment, CKM/PMNS
// theorem, or official ledger update.
package generation2sealedkappaflavorledgerconsistencyaudit

import (
	"fmt"
	"strings"
)

const (
	AuditID         = "GATE967-GENERATION2-GENERATION2SEALEDKAPPAFLAVORLEDGERCONSISTENCYAUDIT"
	InheritedStatus = "R4_EPSILON_E_DIAGNOSTIC_ALLOWED_NO_SOURCE_ROLE"
	Verdict         = "SEALED_KAPPA_LEDGER_CONSISTENCY_ADMISSIBLE_UNDER_TRIPLE_SEAL_BUT_NOT_NATIVE_YUKAWA_SPECTRUM"
	Classification  = "R4_SEALED_KAPPA_LEDGER_DIAGNOSTIC_NOT_NATIVE"
	ShortStatus     = "R4_KAPPA_DIAGNOSTIC_ALLOWED_NO_SPECTRUM_THEOREM"
	NextGate        = "NEXT_GATE968_SEALED_KOIDE_SHADOW_COMPATIBILITY_AUDIT"
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
		Final:            "Gate 967 concludes R4_KAPPA_DIAGNOSTIC_ALLOWED_NO_SPECTRUM_THEOREM. It preserves all inherited seals and does not derive native flavor, physical particles, CKM/PMNS, individual Yukawa values, official ledgers, or an R4 native Yukawa spectrum theorem.",
	}
	if err := validate(a); err != nil {
		return Analysis{}, err
	}
	return a, nil
}

func validate(a Analysis) error {
	if a.AuditID != AuditID || a.Inherited != InheritedStatus || a.Verdict != Verdict || a.Classification != Classification || a.ShortStatus != ShortStatus || a.NextGate != NextGate {
		return fmt.Errorf("bad Gate 967 identity")
	}
	d := a.Decision
	if !d.InheritedSealedRail || !d.R3DualSealPreserved || !d.ScalarSourceSealPreserved || !d.PostOrientationSealPreserved || !d.ExternalGenerationCarrierSealPreserved || !d.ExternalFlavorOrientationSealPreserved {
		return fmt.Errorf("Gate 967 must preserve inherited seals: %#v", d)
	}
	if !d.AllowsSealedOperation {
		return fmt.Errorf("Gate 967 must allow only the sealed operation under audit")
	}
	if d.DerivesNativeFlavor || d.DerivesNativeYukawaMatrix || d.DerivesIndividualYukawas || d.DerivesCKMPMNS || d.AssignsPhysicalParticles || d.UpdatesOfficialLedger {
		return fmt.Errorf("Gate 967 overclaimed native/physical result: %#v", d)
	}
	if !containsAll(a.Supports, RequiredSupports()) {
		return fmt.Errorf("Gate 967 missing required supports")
	}
	if !containsAll(a.Failures, RequiredFailures()) {
		return fmt.Errorf("Gate 967 missing required firewalls")
	}
	return nil
}

func Allowed() []string {
	return []string{
		"kappa_e may be tested as sealed flavor-orientation residual target",
		"kappa_lambda may be tested as sealed scalar/flavor consistency target",
	}
}

func Forbidden() []string {
	return []string{
		"kappa_e cannot source Phi_flav",
		"kappa_lambda cannot source native scalar theorem",
		"kappa ledger cannot derive Yukawa spectrum",
		"kappa ledger cannot update official constants",
	}
}

func Supports() []string {
	return []string{
		"CONDITIONAL_SUPPORT_KAPPA_LEDGER_TEST_ALLOWED_UNDER_TRIPLE_SEAL",
		"CONDITIONAL_SUPPORT_KAPPA_E_AND_KAPPA_LAMBDA_ARE_DOWNSTREAM_CONSISTENCY_TARGETS",
		"CONDITIONAL_SUPPORT_KAPPA_DIAGNOSTICS_INHERIT_TRIPLE_SEAL",
	}
}

func Failures() []string {
	return []string{
		"FAILED_ROUTE_KAPPA_E_CANNOT_SOURCE_FLAVOR_ORIENTATION",
		"FAILED_ROUTE_KAPPA_LAMBDA_NOT_NATIVE_SCALAR_THEOREM",
		"FAILED_ROUTE_KAPPA_LEDGER_NOT_NATIVE_YUKAWA_SPECTRUM",
		"FAILED_ROUTE_KAPPA_DIAGNOSTIC_DOES_NOT_DERIVE_YUKAWA_MATRIX",
		"FAILED_ROUTE_NO_OFFICIAL_LEDGER_UPDATE_ALLOWED",
	}
}

func RequiredSupports() []string {
	return []string{
		"CONDITIONAL_SUPPORT_KAPPA_LEDGER_TEST_ALLOWED_UNDER_TRIPLE_SEAL",
		"CONDITIONAL_SUPPORT_KAPPA_E_AND_KAPPA_LAMBDA_ARE_DOWNSTREAM_CONSISTENCY_TARGETS",
		"CONDITIONAL_SUPPORT_KAPPA_DIAGNOSTICS_INHERIT_TRIPLE_SEAL",
	}
}

func RequiredFailures() []string {
	return []string{
		"FAILED_ROUTE_KAPPA_E_CANNOT_SOURCE_FLAVOR_ORIENTATION",
		"FAILED_ROUTE_KAPPA_LAMBDA_NOT_NATIVE_SCALAR_THEOREM",
		"FAILED_ROUTE_KAPPA_LEDGER_NOT_NATIVE_YUKAWA_SPECTRUM",
		"FAILED_ROUTE_KAPPA_DIAGNOSTIC_DOES_NOT_DERIVE_YUKAWA_MATRIX",
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
