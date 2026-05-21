// Package generation2fourcertificatenativeclauseexecutionaudit implements
// Gate 940: FourCertificate Native Clause Execution and FailureLocalization Audit.
//
// Gate 940 follows the Gate 939 origin-rooted collapse. It executes the four
// native-promotion certificates directly and localizes the remaining failures.
// The audit is intentionally conservative: partial native/bridge support is
// recorded where available, but native R3 is not granted unless all four
// certificates pass.
package generation2fourcertificatenativeclauseexecutionaudit

import (
	"fmt"
	"strings"
)

const (
	AuditID         = "GATE940-GENERATION2FOURCERTIFICATENATIVECLAUSEEXECUTIONAUDIT"
	InheritedStatus = "R3_NATIVE_GAPS_COLLAPSED_TO_ORIGIN_ROOTED_AIRLOCK_FUNCTOR"
	Verdict         = "R3_NATIVE_CERTIFICATE_CHAIN_PARTIALLY_EXECUTED_NATIVE_PROMOTION_BLOCKED"
	Classification  = "R3_NATIVE_PROMOTION_CERTIFICATE_AUDIT_PARTIAL_NATIVE_SUPPORT_NOT_FULL_R3"
	ShortStatus     = "R3_NATIVE_CLAUSE_EXECUTION_PARTIAL_PASS_NOT_NATIVE_R3"
	FullPassVerdict = "R3_NATIVE_POST_ORIENTATION_Z2_TRACE_LEDGER_CERTIFIED"
	NextGate        = "NEXT_PRESSURE_GATE941_S_SPLIT_NATIVE_SOURCE_AND_BOUNDARY_RESPONSE_SCALAR_ORIGIN_AUDIT"
)

const (
	Ssplit          = 0.0012924448188162962
	AlphaB          = 0.0003878958469680527
	NEffOperator    = 3.002327375081808
	CYukawaOperator = 0.9992248096922658
	CHiggsOperator  = 1.037220510866514
)

type CertificateStatus string

const (
	CertificateNativePass     CertificateStatus = "NATIVE_PASS"
	CertificatePartialSupport CertificateStatus = "PARTIAL_SUPPORT"
	CertificateBlocked        CertificateStatus = "BLOCKED"
)

type Certificate struct {
	Index            int
	Name             string
	Question         string
	RequiredTheorem  string
	Status           CertificateStatus
	NativeCertified  bool
	PartialSupported bool
	BlocksNativeR3   bool
	PassMarkers      []string
	SupportMarkers   []string
	FailureMarkers   []string
	Localization     string
}

type DiagnosticValues struct {
	Ssplit          float64
	AlphaB          float64
	NEffOperator    float64
	CYukawaOperator float64
	CHiggsOperator  float64
}

type Analysis struct {
	AuditID            string
	Inherited          string
	Verdict            string
	Classification     string
	ShortStatus        string
	FullNativeEligible bool
	DiagnosticValues   DiagnosticValues
	Certificates       []Certificate
	PrimaryPressure    []string
	RetiredNonBlockers []string
	R4Boundary         []string
	Supports           []string
	Failures           []string
	Final              string
}

func BuildDefault() (Analysis, error) {
	certs := Certificates()
	if len(certs) != 4 {
		return Analysis{}, fmt.Errorf("Gate 940 must execute exactly four native certificates")
	}
	if allNativeCertified(certs) {
		return Analysis{
			AuditID:            AuditID,
			Inherited:          InheritedStatus,
			Verdict:            FullPassVerdict,
			Classification:     "R3_NATIVE_POST_ORIENTATION_Z2_TRACE_LEDGER_CERTIFIED",
			ShortStatus:        "R3_NATIVE_POST_ORIENTATION_TRACE_LEDGER_CERTIFIED",
			FullNativeEligible: true,
			DiagnosticValues:   DefaultDiagnostics(),
			Certificates:       certs,
			PrimaryPressure:    PrimaryPressure(),
			RetiredNonBlockers: RetiredNonBlockers(),
			R4Boundary:         R4Boundary(),
			Supports:           Supports(),
			Failures:           Failures(),
			Final:              "All four certificates passed; native/post-orientation R3 would be certified.",
		}, nil
	}
	return Analysis{
		AuditID:            AuditID,
		Inherited:          InheritedStatus,
		Verdict:            Verdict,
		Classification:     Classification,
		ShortStatus:        ShortStatus,
		FullNativeEligible: false,
		DiagnosticValues:   DefaultDiagnostics(),
		Certificates:       certs,
		PrimaryPressure:    PrimaryPressure(),
		RetiredNonBlockers: RetiredNonBlockers(),
		R4Boundary:         R4Boundary(),
		Supports:           Supports(),
		Failures:           Failures(),
		Final:              "Gate 940 executes the four native-promotion clauses. The projector support lattice and finite trace-measure form receive partial support; S_split source and post-orientation/full-A_F status remain blocking. Native R3 is not granted.",
	}, nil
}

func DefaultDiagnostics() DiagnosticValues {
	return DiagnosticValues{Ssplit: Ssplit, AlphaB: AlphaB, NEffOperator: NEffOperator, CYukawaOperator: CYukawaOperator, CHiggsOperator: CHiggsOperator}
}

func Certificates() []Certificate {
	return []Certificate{
		{
			Index:            1,
			Name:             "Projector-generated admissible support lattice",
			Question:         "Are admissible supports native finite projector submodules rather than arbitrary subspaces?",
			RequiredTheorem:  "NativeAdmissibleAirlockSupportLatticeTheorem",
			Status:           CertificatePartialSupport,
			NativeCertified:  false,
			PartialSupported: true,
			BlocksNativeR3:   true,
			PassMarkers: []string{
				"PASS_PROJECTOR_GENERATED_SUPPORTS_ARE_NATIVE_FINITE_SUBMODULES_RELATIVE_TO_ORIENTED_PROJECTOR_CATEGORY",
				"PASS_ARBITRARY_VECTOR_SUBSPACES_ARE_NOT_REPRESENTED_FINITE_SUPPORTS",
			},
			SupportMarkers: []string{
				"CONDITIONAL_SUPPORT_PROJECTOR_GENERATED_SUPPORT_LATTICE_CAN_BE_NATIVE_IN_FINITE_PROJECTOR_CATEGORY",
				"CONDITIONAL_SUPPORT_AIRLOCK_LATTICE_IS_NATIVE_RELATIVE_TO_ORIENTED_FINITE_PROJECTOR_CATEGORY",
			},
			FailureMarkers: []string{
				"FAILED_ROUTE_PROJECTOR_SUPPORT_LATTICE_NATIVE_ONLY_RELATIVE_TO_ORIENTED_FINITE_PROJECTOR_CATEGORY",
				"FAILED_ROUTE_FULL_A_F_STABILITY_NOT_SOLVED_BY_SUPPORT_LATTICE_CERTIFICATE",
			},
			Localization: "strongest certificate, but only relative to the represented/post-orientation projector category",
		},
		{
			Index:            2,
			Name:             "Native S_split response-parameter source",
			Question:         "Why is S_split the scalar inserted into the boundary-pair response?",
			RequiredTheorem:  "NativeSsplitBoundaryResponseParameterTheorem",
			Status:           CertificateBlocked,
			NativeCertified:  false,
			PartialSupported: false,
			BlocksNativeR3:   true,
			PassMarkers: []string{
				"PASS_S_SPLIT_SQUARED_ARISES_FROM_EXTERIOR_PRODUCT",
			},
			SupportMarkers: []string{
				"CONDITIONAL_SUPPORT_S_SPLIT_UNIFORM_INSERTION_IS_REQUIRED_FOR_BOUNDARY_PAIR_EQUIVARIANCE",
				"CONDITIONAL_SUPPORT_S_SPLIT_POWER_STRUCTURE_FOLLOWS_FROM_EXTERIOR_MULTIPLICATION",
			},
			FailureMarkers: []string{
				"FAILED_ROUTE_NO_NATIVE_S_SPLIT_RESPONSE_PARAMETER_THEOREM_IF_SOURCE_NOT_CERTIFIED",
				"FAILED_ROUTE_NO_NATIVE_TRANSPORT_OF_S_SPLIT_TO_BOUNDARY_PAIR_RESPONSE",
			},
			Localization: "weakest clause; the scalar insertion rule is bridge-valid but the native origin of S_split is not certified",
		},
		{
			Index:            3,
			Name:             "BoundaryActivationMeasure as finite normalized trace response",
			Question:         "Is mu_B a native finite normalized trace response?",
			RequiredTheorem:  "NativeBoundaryActivationMeasureTheorem",
			Status:           CertificatePartialSupport,
			NativeCertified:  false,
			PartialSupported: true,
			BlocksNativeR3:   true,
			PassMarkers: []string{
				"PASS_RANK_OVER_CHAMBER_RANK_IS_FINITE_TRACE_NORMALIZATION",
				"PASS_BOUNDARY_ACTIVATION_MEASURE_NOT_ARBITRARY_AS_FORMAL_TRACE_RESPONSE",
			},
			SupportMarkers: []string{
				"CONDITIONAL_SUPPORT_BOUNDARY_ACTIVATION_MEASURE_CAN_BE_READ_AS_FINITE_NORMALIZED_TRACE_RESPONSE",
				"CONDITIONAL_SUPPORT_TRACE_MEASURE_NATIVE_FORM_SUPPORTED_IF_CHAMBER_ASSIGNMENT_CERTIFIED",
			},
			FailureMarkers: []string{
				"FAILED_ROUTE_NO_NATIVE_BOUNDARY_ACTIVATION_MEASURE_IF_CHAMBER_ASSIGNMENT_NOT_CERTIFIED",
				"FAILED_ROUTE_H10_H72_RESPONSE_CHAMBER_ASSIGNMENT_NOT_NATIVE",
			},
			Localization: "finite trace form is strong; native status still depends on H10/H72 chamber assignment",
		},
		{
			Index:            4,
			Name:             "Full A_F descent or lawful spontaneous-orientation status",
			Question:         "Is the R3 trace ledger native under full A_F, or native only after a lawful finite one-form orientation?",
			RequiredTheorem:  "NativeFullAFDescentTheorem or LawfulSpontaneousOrientationTheorem",
			Status:           CertificateBlocked,
			NativeCertified:  false,
			PartialSupported: false,
			BlocksNativeR3:   true,
			PassMarkers:      []string{},
			SupportMarkers: []string{
				"CONDITIONAL_SUPPORT_POST_ORIENTATION_R3_LEDGER_IS_CORRECT_TARGET_IF_SPONTANEOUS_ORIENTATION_CERTIFIED",
			},
			FailureMarkers: []string{
				"FAILED_ROUTE_NO_NATIVE_POST_ORIENTATION_R3_IF_FINITE_ONE_FORM_ORIENTATION_NOT_CERTIFIED",
				"FAILED_ROUTE_FULL_A_F_DESCENT_STILL_BLOCKED_IF_NO_SPONTANEOUS_ORIENTATION_THEOREM",
			},
			Localization: "orientation/full-A_F status remains blocked; trace bridge stays post-orientation bridge diagnostic",
		},
	}
}

func PrimaryPressure() []string {
	return []string{
		"Certificate II: S_split native response-parameter source",
		"Certificate IV: lawful finite one-form/spontaneous orientation or full A_F descent",
		"Certificate III: H10/H72 response-chamber assignment for native trace measure",
		"Certificate I: full native status of projector support lattice beyond oriented category",
	}
}

func RetiredNonBlockers() []string {
	return []string{
		"lambda versus barlambda representative",
		"+Q_phi versus -Q_phi",
		"representative alpha",
		"cross-lane pollution",
		"Theta(2)=F_2/F_1",
		"bare denominators",
		"orphan support fragments",
		"arbitrary rank-compatible subspaces",
	}
}

func R4Boundary() []string {
	return []string{
		"NO_INDIVIDUAL_YUKAWA_VALUES",
		"NO_PHYSICAL_PARTICLE_ASSIGNMENT",
		"NO_GENERATION_CARRIER_MAP",
		"NO_FLAVOR_ORIENTATION_MAP",
		"NO_R4_NATIVE_YUKAWA_SPECTRUM_THEOREM",
	}
}

func Supports() []string {
	return []string{
		"CONDITIONAL_SUPPORT_PROJECTOR_GENERATED_SUPPORT_LATTICE_CAN_BE_NATIVE_IN_FINITE_PROJECTOR_CATEGORY",
		"CONDITIONAL_SUPPORT_BOUNDARY_ACTIVATION_MEASURE_CAN_BE_READ_AS_FINITE_NORMALIZED_TRACE_RESPONSE",
		"CONDITIONAL_SUPPORT_S_SPLIT_UNIFORM_INSERTION_IS_REQUIRED_FOR_BOUNDARY_PAIR_EQUIVARIANCE",
		"CONDITIONAL_SUPPORT_POST_ORIENTATION_R3_LEDGER_IS_CORRECT_TARGET_IF_SPONTANEOUS_ORIENTATION_CERTIFIED",
		"CONDITIONAL_SUPPORT_NATIVE_PROMOTION_REDUCED_TO_EXPLICIT_CERTIFICATE_EXECUTION",
		"CONDITIONAL_SUPPORT_R3_PRETEST_PASSED_TRACEBRIDGE_INHERITED",
		"CONDITIONAL_SUPPORT_NATIVE_R3_NOW_DEPENDS_MAINLY_ON_S_SPLIT_SOURCE_AND_LAWFUL_SPONTANEOUS_ORIENTATION",
	}
}

func Failures() []string {
	return []string{
		"FAILED_ROUTE_NATIVE_R3_NOT_GRANTED_WITHOUT_ALL_FOUR_CERTIFICATES",
		"FAILED_ROUTE_NO_NATIVE_S_SPLIT_RESPONSE_PARAMETER_THEOREM_IF_SOURCE_NOT_CERTIFIED",
		"FAILED_ROUTE_NO_NATIVE_BOUNDARY_ACTIVATION_MEASURE_IF_CHAMBER_ASSIGNMENT_NOT_CERTIFIED",
		"FAILED_ROUTE_NO_NATIVE_POST_ORIENTATION_R3_IF_FINITE_ONE_FORM_ORIENTATION_NOT_CERTIFIED",
		"FAILED_ROUTE_FULL_A_F_DESCENT_STILL_BLOCKED_IF_NO_SPONTANEOUS_ORIENTATION_THEOREM",
		"FAILED_ROUTE_NO_OFFICIAL_LEDGER_UPDATE_ALLOWED",
		"FAILED_ROUTE_NO_INDIVIDUAL_YUKAWA_VALUES",
		"FAILED_ROUTE_NO_PHYSICAL_PARTICLE_ASSIGNMENT",
		"FAILED_ROUTE_NO_GENERATION_CARRIER_MAP",
		"FAILED_ROUTE_NO_FLAVOR_ORIENTATION_MAP",
		"FAILED_ROUTE_NO_R4_NATIVE_YUKAWA_SPECTRUM_THEOREM",
	}
}

func allNativeCertified(certs []Certificate) bool {
	if len(certs) == 0 {
		return false
	}
	for _, c := range certs {
		if !c.NativeCertified {
			return false
		}
	}
	return true
}

func certificateFailures(certs []Certificate) []string {
	var out []string
	for _, c := range certs {
		out = append(out, c.FailureMarkers...)
	}
	return out
}

func certificateSupports(certs []Certificate) []string {
	var out []string
	for _, c := range certs {
		out = append(out, c.SupportMarkers...)
	}
	return out
}

func countStatus(certs []Certificate, status CertificateStatus) int {
	count := 0
	for _, c := range certs {
		if c.Status == status {
			count++
		}
	}
	return count
}

func FormatDiagnostics(d DiagnosticValues) string {
	return fmt.Sprintf("S_split=%.19g alpha_B=%.19g N_eff_operator=%.16g C_Yukawa_operator=%.16g C_Higgs_operator=%.16g", d.Ssplit, d.AlphaB, d.NEffOperator, d.CYukawaOperator, d.CHiggsOperator)
}

func FormatCertificates(certs []Certificate) string {
	parts := make([]string, 0, len(certs))
	for _, c := range certs {
		parts = append(parts, fmt.Sprintf("%d:%s:%s:%s", c.Index, c.Name, c.Status, c.Localization))
	}
	return strings.Join(parts, " | ")
}

func containsAll(haystack, needles []string) bool {
	set := map[string]struct{}{}
	for _, h := range haystack {
		set[h] = struct{}{}
	}
	for _, n := range needles {
		if _, ok := set[n]; !ok {
			return false
		}
	}
	return true
}
