// Package generation2yukawamatrixadmissibilityauditundertripleseal implements Gate 971: Yukawa Matrix Admissibility Audit Under Triple Seal.
//
// This gate continues the sealed R4 flavor/Yukawa rail. It is deliberately
// firewall-preserving: it records what is lawful under explicit seals and what
// remains forbidden as native ASHA theorem, physical-particle assignment, CKM/PMNS
// theorem, or official ledger update.
package generation2yukawamatrixadmissibilityauditundertripleseal

import (
	"fmt"
	"strings"
)

const (
	AuditID         = "GATE971-GENERATION2-GENERATION2YUKAWAMATRIXADMISSIBILITYAUDITUNDERTRIPLESEAL"
	InheritedStatus = "R4_DIAGNOSTICS_SYNTHESIZED_MATRIX_SOURCE_MISSING"
	Verdict         = "YUKAWA_MATRIX_REQUIRES_GENERATION_FLAVOR_ORIENTATION_SPECTRUM_AND_SECTOR_FRAME_SEALS_EXTERNAL_MATRIX_SEAL_REQUIRED"
	Classification  = "R4_YUKAWA_MATRIX_ADMISSIBILITY_AUDIT_EXTERNAL_MATRIX_SEAL_REQUIRED"
	ShortStatus     = "R4_YUKAWA_MATRIX_REQUIRES_EXTERNAL_MATRIX_SEAL"
	NextGate        = "NEXT_GATE972_EXTERNAL_YUKAWA_MATRIX_SEAL_INSTALLATION_AUDIT"
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
		Final:            "Gate 971 concludes R4_YUKAWA_MATRIX_REQUIRES_EXTERNAL_MATRIX_SEAL. It preserves all inherited seals and does not derive native flavor, physical particles, CKM/PMNS, individual Yukawa values, official ledgers, or an R4 native Yukawa spectrum theorem.",
	}
	if err := validate(a); err != nil {
		return Analysis{}, err
	}
	return a, nil
}

func validate(a Analysis) error {
	if a.AuditID != AuditID || a.Inherited != InheritedStatus || a.Verdict != Verdict || a.Classification != Classification || a.ShortStatus != ShortStatus || a.NextGate != NextGate {
		return fmt.Errorf("bad Gate 971 identity")
	}
	d := a.Decision
	if !d.InheritedSealedRail || !d.R3DualSealPreserved || !d.ScalarSourceSealPreserved || !d.PostOrientationSealPreserved || !d.ExternalGenerationCarrierSealPreserved || !d.ExternalFlavorOrientationSealPreserved {
		return fmt.Errorf("Gate 971 must preserve inherited seals: %#v", d)
	}
	if !d.AllowsSealedOperation {
		return fmt.Errorf("Gate 971 must allow only the sealed operation under audit")
	}
	if d.DerivesNativeFlavor || d.DerivesNativeYukawaMatrix || d.DerivesIndividualYukawas || d.DerivesCKMPMNS || d.AssignsPhysicalParticles || d.UpdatesOfficialLedger {
		return fmt.Errorf("Gate 971 overclaimed native/physical result: %#v", d)
	}
	if !containsAll(a.Supports, RequiredSupports()) {
		return fmt.Errorf("Gate 971 missing required supports")
	}
	if !containsAll(a.Failures, RequiredFailures()) {
		return fmt.Errorf("Gate 971 missing required firewalls")
	}
	return nil
}

func Allowed() []string {
	return []string{
		"Yukawa matrix may be typed as sealed 3x3 sector operator after generation and flavor orientation seals",
		"lawful matrix construction requires sector singular values and left/right frame data",
	}
}

func Forbidden() []string {
	return []string{
		"aggregate R3 tracebody cannot determine 3x3 entries",
		"ExternalFlavorOrientationSeal alone cannot determine singular values",
		"diagnostics cannot backsolve matrix entries",
		"no native Yukawa matrix theorem",
	}
}

func Supports() []string {
	return []string{
		"CONDITIONAL_SUPPORT_YUKAWA_MATRIX_OPERATOR_TYPE_IS_NOW_ADMISSIBLE_UNDER_SEALS",
		"CONDITIONAL_SUPPORT_GENERATION_AND_FLAVOR_ORIENTATION_SEALS_SUPPLY_MATRIX_DOMAIN_AND_BASIS_QUARANTINE",
		"CONDITIONAL_SUPPORT_EXTERNAL_YUKAWA_MATRIX_SEAL_IDENTIFIED_AS_REQUIRED_OBJECT",
	}
}

func Failures() []string {
	return []string{
		"FAILED_ROUTE_AGGREGATE_R3_TRACEBODY_DOES_NOT_DETERMINE_YUKAWA_MATRIX_ENTRIES",
		"FAILED_ROUTE_EXTERNAL_FLAVOR_ORIENTATION_SEAL_DOES_NOT_SUPPLY_SINGULAR_VALUES",
		"FAILED_ROUTE_DIAGNOSTIC_BACKSOLVE_CANNOT_SOURCE_YUKAWA_MATRIX",
		"FAILED_ROUTE_NO_NATIVE_YUKAWA_MATRIX_THEOREM",
		"FAILED_ROUTE_EXTERNAL_YUKAWA_MATRIX_SEAL_REQUIRED",
	}
}

func RequiredSupports() []string {
	return []string{
		"CONDITIONAL_SUPPORT_YUKAWA_MATRIX_OPERATOR_TYPE_IS_NOW_ADMISSIBLE_UNDER_SEALS",
		"CONDITIONAL_SUPPORT_GENERATION_AND_FLAVOR_ORIENTATION_SEALS_SUPPLY_MATRIX_DOMAIN_AND_BASIS_QUARANTINE",
		"CONDITIONAL_SUPPORT_EXTERNAL_YUKAWA_MATRIX_SEAL_IDENTIFIED_AS_REQUIRED_OBJECT",
	}
}

func RequiredFailures() []string {
	return []string{
		"FAILED_ROUTE_AGGREGATE_R3_TRACEBODY_DOES_NOT_DETERMINE_YUKAWA_MATRIX_ENTRIES",
		"FAILED_ROUTE_EXTERNAL_FLAVOR_ORIENTATION_SEAL_DOES_NOT_SUPPLY_SINGULAR_VALUES",
		"FAILED_ROUTE_DIAGNOSTIC_BACKSOLVE_CANNOT_SOURCE_YUKAWA_MATRIX",
		"FAILED_ROUTE_NO_NATIVE_YUKAWA_MATRIX_THEOREM",
		"FAILED_ROUTE_EXTERNAL_YUKAWA_MATRIX_SEAL_REQUIRED",
	}
}

func MatrixNormalForm() []string {
	return []string{
		"Y_f admissibility requires: domain C^3_gen,seal; flavor representative Phi_flav^seal; sector singular values; left/right frame data; scale and scheme metadata",
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
