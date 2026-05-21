// Package generation2externalyukawamatrixsealinstallationaudit implements Gate 972: ExternalYukawaMatrixSeal Installation Audit.
//
// This gate continues the sealed R4 flavor/Yukawa rail. It is deliberately
// firewall-preserving: it records what is lawful under explicit seals and what
// remains forbidden as native ASHA theorem, physical-particle assignment, CKM/PMNS
// theorem, or official ledger update.
package generation2externalyukawamatrixsealinstallationaudit

import (
	"fmt"
	"strings"
)

const (
	AuditID         = "GATE972-GENERATION2-GENERATION2EXTERNALYUKAWAMATRIXSEALINSTALLATIONAUDIT"
	InheritedStatus = "R4_YUKAWA_MATRIX_REQUIRES_EXTERNAL_MATRIX_SEAL"
	Verdict         = "EXTERNAL_YUKAWA_MATRIX_SEAL_INSTALLED_MATRIX_OPERATOR_LEDGER_ALLOWED_SEALED_NOT_NATIVE"
	Classification  = "R4_EXTERNAL_YUKAWA_MATRIX_SEALED_NO_NATIVE_YUKAWA_THEOREM"
	ShortStatus     = "R4_EXTERNAL_YUKAWA_MATRIX_SEAL_INSTALLED"
	NextGate        = "NEXT_GATE973_SEALED_YUKAWA_MATRIX_OPERATOR_CONSTRUCTION_AUDIT"
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
			ExternalYukawaMatrixSealPreserved:      true,
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
		Final:            "Gate 972 concludes R4_EXTERNAL_YUKAWA_MATRIX_SEAL_INSTALLED. It preserves all inherited seals and does not derive native flavor, physical particles, CKM/PMNS, individual Yukawa values, official ledgers, or an R4 native Yukawa spectrum theorem.",
	}
	if err := validate(a); err != nil {
		return Analysis{}, err
	}
	return a, nil
}

func validate(a Analysis) error {
	if a.AuditID != AuditID || a.Inherited != InheritedStatus || a.Verdict != Verdict || a.Classification != Classification || a.ShortStatus != ShortStatus || a.NextGate != NextGate {
		return fmt.Errorf("bad Gate 972 identity")
	}
	d := a.Decision
	if !d.InheritedSealedRail || !d.R3DualSealPreserved || !d.ScalarSourceSealPreserved || !d.PostOrientationSealPreserved || !d.ExternalGenerationCarrierSealPreserved || !d.ExternalFlavorOrientationSealPreserved {
		return fmt.Errorf("Gate 972 must preserve inherited seals: %#v", d)
	}
	if !d.AllowsSealedOperation {
		return fmt.Errorf("Gate 972 must allow only the sealed operation under audit")
	}
	if d.DerivesNativeFlavor || d.DerivesNativeYukawaMatrix || d.DerivesIndividualYukawas || d.DerivesCKMPMNS || d.AssignsPhysicalParticles || d.UpdatesOfficialLedger {
		return fmt.Errorf("Gate 972 overclaimed native/physical result: %#v", d)
	}
	if !containsAll(a.Supports, RequiredSupports()) {
		return fmt.Errorf("Gate 972 missing required supports")
	}
	if !containsAll(a.Failures, RequiredFailures()) {
		return fmt.Errorf("Gate 972 missing required firewalls")
	}
	return nil
}

func Allowed() []string {
	return []string{
		"ExternalYukawaMatrixSeal may provide quarantined sector 3x3 matrices",
		"sealed matrices may be used for downstream diagnostic validation",
		"matrix data must carry scale/scheme/sector/neutrino convention metadata",
	}
}

func Forbidden() []string {
	return []string{
		"external matrix seal is not native Yukawa theorem",
		"external matrix seal does not derive CKM/PMNS",
		"external matrix seal does not assign particles as ASHA theorem",
		"external matrix seal cannot update official ledger",
	}
}

func Supports() []string {
	return []string{
		"CONDITIONAL_SUPPORT_EXTERNAL_YUKAWA_MATRIX_SEAL_INSTALLED",
		"CONDITIONAL_SUPPORT_SEALED_3X3_SECTOR_OPERATORS_AVAILABLE_FOR_LEDGER_TESTS",
		"CONDITIONAL_SUPPORT_MATRIX_SEAL_REQUIRES_METADATA_SCALE_SCHEME_SECTOR_AND_NEUTRINO_CONVENTION",
	}
}

func Failures() []string {
	return []string{
		"FAILED_ROUTE_EXTERNAL_YUKAWA_MATRIX_SEAL_NOT_NATIVE_YUKAWA_THEOREM",
		"FAILED_ROUTE_EXTERNAL_YUKAWA_MATRIX_SEAL_NOT_CKM_PMNS_THEOREM",
		"FAILED_ROUTE_EXTERNAL_YUKAWA_MATRIX_SEAL_NOT_PHYSICAL_PARTICLE_ASSIGNMENT_THEOREM",
		"FAILED_ROUTE_NO_OFFICIAL_LEDGER_UPDATE_ALLOWED",
		"FAILED_ROUTE_R3_TRACEBRIDGE_REMAINS_DUALSEALED_NOT_NATIVE",
	}
}

func RequiredSupports() []string {
	return []string{
		"CONDITIONAL_SUPPORT_EXTERNAL_YUKAWA_MATRIX_SEAL_INSTALLED",
		"CONDITIONAL_SUPPORT_SEALED_3X3_SECTOR_OPERATORS_AVAILABLE_FOR_LEDGER_TESTS",
		"CONDITIONAL_SUPPORT_MATRIX_SEAL_REQUIRES_METADATA_SCALE_SCHEME_SECTOR_AND_NEUTRINO_CONVENTION",
	}
}

func RequiredFailures() []string {
	return []string{
		"FAILED_ROUTE_EXTERNAL_YUKAWA_MATRIX_SEAL_NOT_NATIVE_YUKAWA_THEOREM",
		"FAILED_ROUTE_EXTERNAL_YUKAWA_MATRIX_SEAL_NOT_CKM_PMNS_THEOREM",
		"FAILED_ROUTE_EXTERNAL_YUKAWA_MATRIX_SEAL_NOT_PHYSICAL_PARTICLE_ASSIGNMENT_THEOREM",
		"FAILED_ROUTE_NO_OFFICIAL_LEDGER_UPDATE_ALLOWED",
		"FAILED_ROUTE_R3_TRACEBRIDGE_REMAINS_DUALSEALED_NOT_NATIVE",
	}
}

func MatrixNormalForm() []string {
	return []string{
		"ExternalYukawaMatrixSeal fields: sector, scale_mu, scheme, Y_f in Mat_3(C), left_frame, right_frame, singular_values, uncertainty, neutrino_convention",
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
