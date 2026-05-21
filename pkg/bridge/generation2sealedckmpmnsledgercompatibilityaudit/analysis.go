// Package generation2sealedckmpmnsledgercompatibilityaudit implements Gate 969: Sealed CKM/PMNS Ledger Compatibility Audit.
//
// This gate continues the sealed R4 flavor/Yukawa rail. It is deliberately
// firewall-preserving: it records what is lawful under explicit seals and what
// remains forbidden as native ASHA theorem, physical-particle assignment, CKM/PMNS
// theorem, or official ledger update.
package generation2sealedckmpmnsledgercompatibilityaudit

import (
	"fmt"
	"strings"
)

const (
	AuditID         = "GATE969-GENERATION2-GENERATION2SEALEDCKMPMNSLEDGERCOMPATIBILITYAUDIT"
	InheritedStatus = "R4_KOIDE_SHADOW_ALLOWED_NO_LEPTON_THEOREM"
	Verdict         = "CKM_PMNS_LEDGER_COMPATIBILITY_ADMISSIBLE_UNDER_TRIPLE_SEAL_BUT_NOT_MIXING_THEOREM"
	Classification  = "R4_SEALED_CKM_PMNS_LEDGER_DIAGNOSTIC_NOT_NATIVE"
	ShortStatus     = "R4_CKM_PMNS_COMPATIBILITY_ALLOWED_NO_MIXING_THEOREM"
	NextGate        = "NEXT_GATE970_TRIPLE_SEALED_FLAVOR_LEDGER_DIAGNOSTIC_SYNTHESIS_AUDIT"
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
		Final:            "Gate 969 concludes R4_CKM_PMNS_COMPATIBILITY_ALLOWED_NO_MIXING_THEOREM. It preserves all inherited seals and does not derive native flavor, physical particles, CKM/PMNS, individual Yukawa values, official ledgers, or an R4 native Yukawa spectrum theorem.",
	}
	if err := validate(a); err != nil {
		return Analysis{}, err
	}
	return a, nil
}

func validate(a Analysis) error {
	if a.AuditID != AuditID || a.Inherited != InheritedStatus || a.Verdict != Verdict || a.Classification != Classification || a.ShortStatus != ShortStatus || a.NextGate != NextGate {
		return fmt.Errorf("bad Gate 969 identity")
	}
	d := a.Decision
	if !d.InheritedSealedRail || !d.R3DualSealPreserved || !d.ScalarSourceSealPreserved || !d.PostOrientationSealPreserved || !d.ExternalGenerationCarrierSealPreserved || !d.ExternalFlavorOrientationSealPreserved {
		return fmt.Errorf("Gate 969 must preserve inherited seals: %#v", d)
	}
	if !d.AllowsSealedOperation {
		return fmt.Errorf("Gate 969 must allow only the sealed operation under audit")
	}
	if d.DerivesNativeFlavor || d.DerivesNativeYukawaMatrix || d.DerivesIndividualYukawas || d.DerivesCKMPMNS || d.AssignsPhysicalParticles || d.UpdatesOfficialLedger {
		return fmt.Errorf("Gate 969 overclaimed native/physical result: %#v", d)
	}
	if !containsAll(a.Supports, RequiredSupports()) {
		return fmt.Errorf("Gate 969 missing required supports")
	}
	if !containsAll(a.Failures, RequiredFailures()) {
		return fmt.Errorf("Gate 969 missing required firewalls")
	}
	return nil
}

func Allowed() []string {
	return []string{
		"CKM ledger may be tested as sealed compatibility target",
		"PMNS ledger may be tested as sealed compatibility target",
		"mixing compatibility may be used only after ExternalFlavorOrientationSeal",
	}
}

func Forbidden() []string {
	return []string{
		"CKM cannot source flavor orientation",
		"PMNS cannot source flavor orientation",
		"CKM/PMNS cannot become native mixing theorem",
		"mixing data cannot derive matrix entries",
	}
}

func Supports() []string {
	return []string{
		"CONDITIONAL_SUPPORT_CKM_PMNS_LEDGER_COMPATIBILITY_TEST_ALLOWED_UNDER_TRIPLE_SEAL",
		"CONDITIONAL_SUPPORT_CKM_PMNS_ARE_DOWNSTREAM_LEDGER_TARGETS_ONLY",
		"CONDITIONAL_SUPPORT_MIXING_COMPATIBILITY_INHERITS_EXTERNAL_FLAVOR_ORIENTATION_SEAL",
	}
}

func Failures() []string {
	return []string{
		"FAILED_ROUTE_CKM_PMNS_LEDGER_NOT_NATIVE_MIXING_THEOREM",
		"FAILED_ROUTE_CKM_PMNS_CANNOT_SOURCE_FLAVOR_ORIENTATION",
		"FAILED_ROUTE_CKM_PMNS_DO_NOT_DERIVE_YUKAWA_MATRIX_ENTRIES",
		"FAILED_ROUTE_NO_CKM_PMNS_THEOREM",
		"FAILED_ROUTE_NO_PMNS_THEOREM",
		"FAILED_ROUTE_NO_OFFICIAL_LEDGER_UPDATE_ALLOWED",
	}
}

func RequiredSupports() []string {
	return []string{
		"CONDITIONAL_SUPPORT_CKM_PMNS_LEDGER_COMPATIBILITY_TEST_ALLOWED_UNDER_TRIPLE_SEAL",
		"CONDITIONAL_SUPPORT_CKM_PMNS_ARE_DOWNSTREAM_LEDGER_TARGETS_ONLY",
		"CONDITIONAL_SUPPORT_MIXING_COMPATIBILITY_INHERITS_EXTERNAL_FLAVOR_ORIENTATION_SEAL",
	}
}

func RequiredFailures() []string {
	return []string{
		"FAILED_ROUTE_CKM_PMNS_LEDGER_NOT_NATIVE_MIXING_THEOREM",
		"FAILED_ROUTE_CKM_PMNS_CANNOT_SOURCE_FLAVOR_ORIENTATION",
		"FAILED_ROUTE_CKM_PMNS_DO_NOT_DERIVE_YUKAWA_MATRIX_ENTRIES",
		"FAILED_ROUTE_NO_CKM_PMNS_THEOREM",
		"FAILED_ROUTE_NO_PMNS_THEOREM",
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
