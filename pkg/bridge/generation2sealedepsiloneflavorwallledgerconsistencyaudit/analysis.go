// Package generation2sealedepsiloneflavorwallledgerconsistencyaudit implements Gate 966: Sealed Epsilon_e Flavor-Wall Ledger Consistency Audit.
//
// This gate continues the sealed R4 flavor/Yukawa rail. It is deliberately
// firewall-preserving: it records what is lawful under explicit seals and what
// remains forbidden as native ASHA theorem, physical-particle assignment, CKM/PMNS
// theorem, or official ledger update.
package generation2sealedepsiloneflavorwallledgerconsistencyaudit

import (
	"fmt"
	"strings"
)

const (
	AuditID         = "GATE966-GENERATION2-GENERATION2SEALEDEPSILONEFLAVORWALLLEDGERCONSISTENCYAUDIT"
	InheritedStatus = "R4_FLAVOR_LEDGER_TESTS_ALLOWED_SEALED_NOT_NATIVE"
	Verdict         = "SEALED_EPSILON_E_FLAVOR_WALL_LEDGER_CONSISTENCY_ADMISSIBLE_BUT_NOT_NATIVE_FLAVOR_THEOREM"
	Classification  = "R4_SEALED_EPSILON_E_LEDGER_DIAGNOSTIC_NOT_NATIVE"
	ShortStatus     = "R4_EPSILON_E_DIAGNOSTIC_ALLOWED_NO_SOURCE_ROLE"
	NextGate        = "NEXT_GATE967_SEALED_KAPPA_FLAVOR_LEDGER_CONSISTENCY_AUDIT"
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
		Final:            "Gate 966 concludes R4_EPSILON_E_DIAGNOSTIC_ALLOWED_NO_SOURCE_ROLE. It preserves all inherited seals and does not derive native flavor, physical particles, CKM/PMNS, individual Yukawa values, official ledgers, or an R4 native Yukawa spectrum theorem.",
	}
	if err := validate(a); err != nil {
		return Analysis{}, err
	}
	return a, nil
}

func validate(a Analysis) error {
	if a.AuditID != AuditID || a.Inherited != InheritedStatus || a.Verdict != Verdict || a.Classification != Classification || a.ShortStatus != ShortStatus || a.NextGate != NextGate {
		return fmt.Errorf("bad Gate 966 identity")
	}
	d := a.Decision
	if !d.InheritedSealedRail || !d.R3DualSealPreserved || !d.ScalarSourceSealPreserved || !d.PostOrientationSealPreserved || !d.ExternalGenerationCarrierSealPreserved || !d.ExternalFlavorOrientationSealPreserved {
		return fmt.Errorf("Gate 966 must preserve inherited seals: %#v", d)
	}
	if !d.AllowsSealedOperation {
		return fmt.Errorf("Gate 966 must allow only the sealed operation under audit")
	}
	if d.DerivesNativeFlavor || d.DerivesNativeYukawaMatrix || d.DerivesIndividualYukawas || d.DerivesCKMPMNS || d.AssignsPhysicalParticles || d.UpdatesOfficialLedger {
		return fmt.Errorf("Gate 966 overclaimed native/physical result: %#v", d)
	}
	if !containsAll(a.Supports, RequiredSupports()) {
		return fmt.Errorf("Gate 966 missing required supports")
	}
	if !containsAll(a.Failures, RequiredFailures()) {
		return fmt.Errorf("Gate 966 missing required firewalls")
	}
	return nil
}

func Allowed() []string {
	return []string{
		"epsilon_e may be tested as sealed flavor-wall ledger consistency target",
		"epsilon_e may compare against flavor-wall residual ledger under triple seal",
	}
}

func Forbidden() []string {
	return []string{
		"epsilon_e cannot source generation carrier",
		"epsilon_e cannot source flavor orientation",
		"epsilon_e cannot become native flavor theorem",
		"epsilon_e cannot update official ledger",
	}
}

func Supports() []string {
	return []string{
		"CONDITIONAL_SUPPORT_EPSILON_E_FLAVOR_WALL_LEDGER_TEST_ALLOWED_UNDER_TRIPLE_SEAL",
		"CONDITIONAL_SUPPORT_EPSILON_E_CAN_BE_USED_AS_DOWNSTREAM_SEALED_DIAGNOSTIC",
		"CONDITIONAL_SUPPORT_EPSILON_E_TEST_INHERITS_R3_DUALSEAL_EXTERNAL_C3_AND_EXTERNAL_FLAVOR_ORIENTATION_SEAL",
	}
}

func Failures() []string {
	return []string{
		"FAILED_ROUTE_EPSILON_E_NOT_NATIVE_FLAVOR_THEOREM",
		"FAILED_ROUTE_EPSILON_E_CANNOT_SOURCE_GENERATION_OR_ORIENTATION",
		"FAILED_ROUTE_EPSILON_E_CANNOT_SOURCE_CANONICAL_FLAVOR_SELECTOR",
		"FAILED_ROUTE_EPSILON_E_LEDGER_TEST_DOES_NOT_DERIVE_YUKAWA_MATRIX",
		"FAILED_ROUTE_NO_OFFICIAL_LEDGER_UPDATE_ALLOWED",
	}
}

func RequiredSupports() []string {
	return []string{
		"CONDITIONAL_SUPPORT_EPSILON_E_FLAVOR_WALL_LEDGER_TEST_ALLOWED_UNDER_TRIPLE_SEAL",
		"CONDITIONAL_SUPPORT_EPSILON_E_CAN_BE_USED_AS_DOWNSTREAM_SEALED_DIAGNOSTIC",
		"CONDITIONAL_SUPPORT_EPSILON_E_TEST_INHERITS_R3_DUALSEAL_EXTERNAL_C3_AND_EXTERNAL_FLAVOR_ORIENTATION_SEAL",
	}
}

func RequiredFailures() []string {
	return []string{
		"FAILED_ROUTE_EPSILON_E_NOT_NATIVE_FLAVOR_THEOREM",
		"FAILED_ROUTE_EPSILON_E_CANNOT_SOURCE_GENERATION_OR_ORIENTATION",
		"FAILED_ROUTE_EPSILON_E_CANNOT_SOURCE_CANONICAL_FLAVOR_SELECTOR",
		"FAILED_ROUTE_EPSILON_E_LEDGER_TEST_DOES_NOT_DERIVE_YUKAWA_MATRIX",
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
