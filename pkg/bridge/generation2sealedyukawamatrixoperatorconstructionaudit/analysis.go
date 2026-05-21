// Package generation2sealedyukawamatrixoperatorconstructionaudit implements Gate 973: Sealed Yukawa Matrix Operator Construction Audit.
//
// This gate continues the sealed R4 flavor/Yukawa rail. It is deliberately
// firewall-preserving: it records what is lawful under explicit seals and what
// remains forbidden as native ASHA theorem, physical-particle assignment, CKM/PMNS
// theorem, or official ledger update.
package generation2sealedyukawamatrixoperatorconstructionaudit

import (
	"fmt"
	"strings"
)

const (
	AuditID         = "GATE973-GENERATION2-GENERATION2SEALEDYUKAWAMATRIXOPERATORCONSTRUCTIONAUDIT"
	InheritedStatus = "R4_EXTERNAL_YUKAWA_MATRIX_SEAL_INSTALLED"
	Verdict         = "SEALED_YUKAWA_MATRIX_OPERATOR_LEDGER_CONSTRUCTED_UNDER_R3_EXTERNAL_C3_FLAVOR_ORIENTATION_AND_YUKAWA_MATRIX_SEALS_NOT_NATIVE"
	Classification  = "R4_SEALED_YUKAWA_MATRIX_LEDGER_AVAILABLE_NOT_NATIVE"
	ShortStatus     = "R4_LAWFUL_SEALED_YUKAWA_MATRIX_AVAILABLE_NOT_NATIVE"
	NextGate        = "NEXT_GATE974_SEALED_YUKAWA_MATRIX_LEDGER_VALIDATION_AGAINST_FLAVOR_DIAGNOSTICS_AUDIT"
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
		Final:            "Gate 973 concludes R4_LAWFUL_SEALED_YUKAWA_MATRIX_AVAILABLE_NOT_NATIVE. It preserves all inherited seals and does not derive native flavor, physical particles, CKM/PMNS, individual Yukawa values, official ledgers, or an R4 native Yukawa spectrum theorem.",
	}
	if err := validate(a); err != nil {
		return Analysis{}, err
	}
	return a, nil
}

func validate(a Analysis) error {
	if a.AuditID != AuditID || a.Inherited != InheritedStatus || a.Verdict != Verdict || a.Classification != Classification || a.ShortStatus != ShortStatus || a.NextGate != NextGate {
		return fmt.Errorf("bad Gate 973 identity")
	}
	d := a.Decision
	if !d.InheritedSealedRail || !d.R3DualSealPreserved || !d.ScalarSourceSealPreserved || !d.PostOrientationSealPreserved || !d.ExternalGenerationCarrierSealPreserved || !d.ExternalFlavorOrientationSealPreserved {
		return fmt.Errorf("Gate 973 must preserve inherited seals: %#v", d)
	}
	if !d.AllowsSealedOperation {
		return fmt.Errorf("Gate 973 must allow only the sealed operation under audit")
	}
	if d.DerivesNativeFlavor || d.DerivesNativeYukawaMatrix || d.DerivesIndividualYukawas || d.DerivesCKMPMNS || d.AssignsPhysicalParticles || d.UpdatesOfficialLedger {
		return fmt.Errorf("Gate 973 overclaimed native/physical result: %#v", d)
	}
	if !containsAll(a.Supports, RequiredSupports()) {
		return fmt.Errorf("Gate 973 missing required supports")
	}
	if !containsAll(a.Failures, RequiredFailures()) {
		return fmt.Errorf("Gate 973 missing required firewalls")
	}
	return nil
}

func Allowed() []string {
	return []string{
		"Y_f^seal = U_fL^seal D_f^seal U_fR^seal_dagger is lawful as sealed sector matrix normal form",
		"H_f^seal = Y_f_dagger Y_f is lawful as sealed trace diagnostic object",
		"CKM/PMNS compatibility may be tested as sealed frame-misalignment diagnostic",
	}
}

func Forbidden() []string {
	return []string{
		"sealed Yukawa matrix is not native ASHA derivation",
		"sealed matrix entries are not computed from finite core",
		"sealed matrix cannot update official physical constants",
		"sealed matrix cannot close R4 native Yukawa theorem",
	}
}

func Supports() []string {
	return []string{
		"CONDITIONAL_SUPPORT_SEALED_YUKAWA_MATRIX_NORMAL_FORM_CONSTRUCTED",
		"CONDITIONAL_SUPPORT_SEALED_HERMITIAN_TRACEBODY_H_F_EQUALS_Y_DAGGER_Y_AVAILABLE",
		"CONDITIONAL_SUPPORT_SEALED_MATRIX_LEDGER_CAN_BE_VALIDATED_AGAINST_EPSILON_KAPPA_KOIDE_CKM_PMNS_DIAGNOSTICS",
		"CONDITIONAL_SUPPORT_LAWFUL_YUKAWA_MATRIX_EXISTS_ONLY_UNDER_EXPLICIT_SEALS",
	}
}

func Failures() []string {
	return []string{
		"FAILED_ROUTE_SEALED_YUKAWA_MATRIX_NOT_NATIVE_DERIVATION",
		"FAILED_ROUTE_YUKAWA_MATRIX_ENTRIES_NOT_COMPUTED_FROM_FINITE_CORE",
		"FAILED_ROUTE_NO_R4_NATIVE_YUKAWA_SPECTRUM_THEOREM",
		"FAILED_ROUTE_NO_OFFICIAL_LEDGER_UPDATE_ALLOWED",
		"FAILED_ROUTE_R3_TRACEBRIDGE_REMAINS_DUALSEALED_NOT_NATIVE",
		"FAILED_ROUTE_EXTERNAL_C3_AND_EXTERNAL_FLAVOR_ORIENTATION_SEALS_REMAIN_ACTIVE",
	}
}

func RequiredSupports() []string {
	return []string{
		"CONDITIONAL_SUPPORT_SEALED_YUKAWA_MATRIX_NORMAL_FORM_CONSTRUCTED",
		"CONDITIONAL_SUPPORT_SEALED_HERMITIAN_TRACEBODY_H_F_EQUALS_Y_DAGGER_Y_AVAILABLE",
		"CONDITIONAL_SUPPORT_SEALED_MATRIX_LEDGER_CAN_BE_VALIDATED_AGAINST_EPSILON_KAPPA_KOIDE_CKM_PMNS_DIAGNOSTICS",
	}
}

func RequiredFailures() []string {
	return []string{
		"FAILED_ROUTE_SEALED_YUKAWA_MATRIX_NOT_NATIVE_DERIVATION",
		"FAILED_ROUTE_YUKAWA_MATRIX_ENTRIES_NOT_COMPUTED_FROM_FINITE_CORE",
		"FAILED_ROUTE_NO_R4_NATIVE_YUKAWA_SPECTRUM_THEOREM",
		"FAILED_ROUTE_NO_OFFICIAL_LEDGER_UPDATE_ALLOWED",
		"FAILED_ROUTE_R3_TRACEBRIDGE_REMAINS_DUALSEALED_NOT_NATIVE",
	}
}

func MatrixNormalForm() []string {
	return []string{
		"Y_f^seal = U_fL^seal D_f^seal (U_fR^seal)^†",
		"D_f^seal = diag(y_f1^seal, y_f2^seal, y_f3^seal)",
		"H_f^seal = (Y_f^seal)^† Y_f^seal = U_fR^seal (D_f^seal)^2 (U_fR^seal)^†",
		"CKM^seal = (U_uL^seal)^† U_dL^seal",
		"PMNS^seal = (U_eL^seal)^† U_nuL^seal when neutrino convention permits",
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
