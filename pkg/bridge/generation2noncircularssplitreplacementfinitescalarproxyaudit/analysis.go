// Package generation2noncircularssplitreplacementfinitescalarproxyaudit implements
// Gate 946: NonCircular S_split Replacement and FiniteScalar Proxy Audit.
//
// Gate 946 follows Gate 945's addend/circularity firewall. It tests whether
// the bridge/history scalar S_split can be replaced by a noncircular finite
// scalar proxy sourced from native ASHA finite data. The audit intentionally
// rejects reparameterizations, output inversions, raw rank coefficients, and
// bridge-history residuals as native R3 inputs.
package generation2noncircularssplitreplacementfinitescalarproxyaudit

import (
	"fmt"
	"math"
	"strings"
)

const (
	AuditID         = "GATE946-GENERATION2NONCIRCULARSSPLITREPLACEMENTFINITESCALARPROXYAUDIT"
	InheritedStatus = "R3_S_SPLIT_NATIVE_SOURCE_BLOCKED_BY_ADDEND_ORIGIN_AND_CIRCULARITY"
	Verdict         = "NO_NONCIRCULAR_FINITE_SCALAR_PROXY_FOUND_FOR_S_SPLIT_NATIVE_REPLACEMENT"
	Classification  = "R3_S_SPLIT_NONCIRCULAR_REPLACEMENT_AUDIT_FAILED_SCALAR_SEAL_REMAINS"
	ShortStatus     = "R3_S_SPLIT_SCALAR_SOURCE_SEAL_CONFIRMED"
	NextGate        = "NEXT_PRESSURE_GATE947_SCALAR_SOURCE_SEAL_CLASSIFICATION_AND_R3_BRIDGE_THEOREM_BOUNDARY_AUDIT"
)

const (
	Ssplit      = 0.0012924448188162962
	SevenOver72 = 7.0 / 72.0
	XiBoundary  = 0.0503471644870914
	LoopUnit    = 1.0 / (8.0 * math.Pi)
)

type CandidateStatus string

const (
	Rejected                   CandidateStatus = "REJECTED_AS_NATIVE_REPLACEMENT"
	SupportedButNotReplacement CandidateStatus = "SUPPORTED_TYPE_BUT_NOT_REPLACEMENT"
)

type ReplacementCriterion struct {
	Name   string
	Passed bool
	Detail string
}

type ProxyCandidateAudit struct {
	Name                   string
	CandidateExpression    string
	FiniteSource           bool
	NonCircular            bool
	ScalarType             bool
	B2Compatible           bool
	MagnitudeCompatible    bool
	NoArbitraryCoefficient bool
	Status                 CandidateStatus
	Verdict                string
	Supports               []string
	Failures               []string
}

type Analysis struct {
	AuditID        string
	Inherited      string
	Verdict        string
	Classification string
	ShortStatus    string
	TargetValue    float64
	Criteria       []ReplacementCriterion
	Candidates     []ProxyCandidateAudit
	Supports       []string
	Failures       []string
	Final          string
}

func BuildDefault() (Analysis, error) {
	criteria := ReplacementCriteria()
	candidates := DefaultCandidates()
	if len(candidates) != 7 {
		return Analysis{}, fmt.Errorf("expected seven replacement candidates, got %d", len(candidates))
	}
	for _, c := range candidates {
		if c.Status != Rejected && c.Status != SupportedButNotReplacement {
			return Analysis{}, fmt.Errorf("candidate %s has invalid status %s", c.Name, c.Status)
		}
		if IsValidNativeReplacement(c) {
			return Analysis{}, fmt.Errorf("Gate 946 must not find native replacement candidate: %#v", c)
		}
	}
	if HasSuccessfulReplacement(candidates) {
		return Analysis{}, fmt.Errorf("unexpected successful S_split replacement")
	}
	return Analysis{
		AuditID:        AuditID,
		Inherited:      InheritedStatus,
		Verdict:        Verdict,
		Classification: Classification,
		ShortStatus:    ShortStatus,
		TargetValue:    Ssplit,
		Criteria:       criteria,
		Candidates:     candidates,
		Supports:       Supports(),
		Failures:       Failures(),
		Final:          "Gate 946 confirms the scalar-source seal: S_split transports to B2 once accepted, but no tested noncircular finite scalar proxy reproduces its magnitude and type without reparameterization, bridge/history input, arbitrary rank algebra, or output circularity.",
	}, nil
}

func ReplacementCriteria() []ReplacementCriterion {
	return []ReplacementCriterion{
		{Name: "finite-source", Passed: false, Detail: "candidate must be built from native finite ASHA objects"},
		{Name: "noncircular", Passed: false, Detail: "candidate must not use R3, alpha_B, N_eff, C_Yukawa, or trace outputs as source"},
		{Name: "scalar-type", Passed: true, Detail: "candidate must live in the central H72 scalar lane"},
		{Name: "B2-compatible", Passed: true, Detail: "candidate must descend to the B2 boundary-response parameter"},
		{Name: "magnitude-compatible", Passed: false, Detail: fmt.Sprintf("candidate must reproduce S_split=%.19g", Ssplit)},
		{Name: "no arbitrary fitted coefficient", Passed: false, Detail: "candidate must not require a fitted functional choice"},
	}
}

func DefaultCandidates() []ProxyCandidateAudit {
	return []ProxyCandidateAudit{
		{
			Name:                   "D_base rescaling",
			CandidateExpression:    "S_proxy=(72/7)D_base with D_base=(7/72)S_split",
			FiniteSource:           false,
			NonCircular:            false,
			ScalarType:             true,
			B2Compatible:           true,
			MagnitudeCompatible:    true,
			NoArbitraryCoefficient: true,
			Status:                 SupportedButNotReplacement,
			Verdict:                "FAILED_ROUTE_D_BASE_REPLACEMENT_IS_REPARAMETERIZATION_OF_S_SPLIT",
			Supports: []string{
				"CONDITIONAL_SUPPORT_D_BASE_HAS_H72_DEFECT_RESPONSE_TYPE",
				"CONDITIONAL_SUPPORT_D_BASE_IS_COMPATIBLE_WITH_BOUNDARY_RESPONSE_SCALAR_LANE",
			},
			Failures: []string{
				"FAILED_ROUTE_D_BASE_REPLACEMENT_IS_REPARAMETERIZATION_OF_S_SPLIT",
				"FAILED_ROUTE_D_BASE_DOES_NOT_SUPPLY_NONCIRCULAR_SCALAR_SOURCE",
			},
		},
		{
			Name:                   "pure rank ratio 7/72",
			CandidateExpression:    "S_proxy=rank(K7)/rank(H72)=7/72",
			FiniteSource:           true,
			NonCircular:            true,
			ScalarType:             false,
			B2Compatible:           false,
			MagnitudeCompatible:    almostEqual(SevenOver72, Ssplit, 1e-15),
			NoArbitraryCoefficient: true,
			Status:                 Rejected,
			Verdict:                "FAILED_ROUTE_7_OVER_72_IS_NORMALIZATION_COEFFICIENT_NOT_S_SPLIT_SCALAR",
			Supports: []string{
				"CONDITIONAL_SUPPORT_7_OVER_72_IS_NATIVE_SHAPE_DEFECT_RESPONSE_COEFFICIENT",
			},
			Failures: []string{
				"FAILED_ROUTE_7_OVER_72_IS_NORMALIZATION_COEFFICIENT_NOT_S_SPLIT_SCALAR",
				"FAILED_ROUTE_PURE_RANK_RATIO_HAS_WRONG_MAGNITUDE_FOR_BOUNDARY_RESPONSE_PARAMETER",
			},
		},
		{
			Name:                   "finite rank/rational chamber expression",
			CandidateExpression:    "rational expression of ranks {1,2,3,4,7,8,10,70,72}",
			FiniteSource:           true,
			NonCircular:            true,
			ScalarType:             false,
			B2Compatible:           false,
			MagnitudeCompatible:    false,
			NoArbitraryCoefficient: false,
			Status:                 Rejected,
			Verdict:                "FAILED_ROUTE_FINITE_RANK_DATA_DO_NOT_CANONICALLY_GENERATE_S_SPLIT_MAGNITUDE",
			Failures: []string{
				"FAILED_ROUTE_FINITE_RANK_DATA_DO_NOT_CANONICALLY_GENERATE_S_SPLIT_MAGNITUDE",
				"FAILED_ROUTE_RANK_EXPRESSIONS_EXPLAIN_COEFFICIENTS_NOT_SCALAR_INPUT",
				"FAILED_ROUTE_NO_NATIVE_RATIONAL_CHAMBER_EXPRESSION_FOR_S_SPLIT_FOUND",
			},
		},
		{
			Name:                   "closure-measure fixed point",
			CandidateExpression:    "s recovered from mu_B, alpha_B, N_eff, or C_Yukawa output",
			FiniteSource:           false,
			NonCircular:            false,
			ScalarType:             true,
			B2Compatible:           true,
			MagnitudeCompatible:    true,
			NoArbitraryCoefficient: false,
			Status:                 Rejected,
			Verdict:                "FAILED_ROUTE_FIXED_POINT_SCALAR_RECOVERY_IS_CIRCULAR_WITH_TRACEBRIDGE_OUTPUT",
			Failures: []string{
				"FAILED_ROUTE_FIXED_POINT_SCALAR_RECOVERY_IS_CIRCULAR_WITH_TRACEBRIDGE_OUTPUT",
				"FAILED_ROUTE_N_EFF_OR_ALPHA_INVERSION_CANNOT_SOURCE_S_SPLIT_FOR_NATIVE_R3",
			},
		},
		{
			Name:                   "xi_boundary boundary-stress residual proxy",
			CandidateExpression:    "small residual from +xi_boundary and -xi_boundary bridge cancellation",
			FiniteSource:           false,
			NonCircular:            false,
			ScalarType:             true,
			B2Compatible:           true,
			MagnitudeCompatible:    false,
			NoArbitraryCoefficient: false,
			Status:                 SupportedButNotReplacement,
			Verdict:                "FAILED_ROUTE_XI_BOUNDARY_RESIDUAL_PROXY_NOT_NATIVE_H72_SCALAR_SOURCE",
			Supports: []string{
				"CONDITIONAL_SUPPORT_S_SPLIT_IS_SMALL_BOUNDARY_RESIDUAL_CANDIDATE",
				"CONDITIONAL_SUPPORT_XI_BOUNDARY_CANCELLATION_PATTERN_EXPLAINS_SCALE_PRESSURE",
			},
			Failures: []string{
				"FAILED_ROUTE_XI_BOUNDARY_RESIDUAL_PROXY_NOT_NATIVE_H72_SCALAR_SOURCE",
				"FAILED_ROUTE_BOUNDARY_STRESS_CANCELLATION_STILL_USES_BRIDGE_HISTORY_INPUTS",
			},
		},
		{
			Name:                   "HistoryLoopUnit or phase-normalized constants",
			CandidateExpression:    "L=1/(8pi) or phase-normalized variants",
			FiniteSource:           false,
			NonCircular:            true,
			ScalarType:             false,
			B2Compatible:           false,
			MagnitudeCompatible:    almostEqual(LoopUnit, Ssplit, 1e-15),
			NoArbitraryCoefficient: false,
			Status:                 Rejected,
			Verdict:                "FAILED_ROUTE_HISTORY_LOOP_UNIT_NOT_TYPED_TO_S_SPLIT_BOUNDARY_RESPONSE_SCALAR",
			Failures: []string{
				"FAILED_ROUTE_HISTORY_LOOP_UNIT_NOT_TYPED_TO_S_SPLIT_BOUNDARY_RESPONSE_SCALAR",
				"FAILED_ROUTE_PHASE_CONSTANTS_DO_NOT_GENERATE_S_SPLIT_WITHOUT_EXTRA_MAP",
			},
		},
		{
			Name:                   "native zero scalar",
			CandidateExpression:    "S_proxy=0",
			FiniteSource:           true,
			NonCircular:            true,
			ScalarType:             true,
			B2Compatible:           true,
			MagnitudeCompatible:    almostEqual(0, Ssplit, 1e-15),
			NoArbitraryCoefficient: true,
			Status:                 Rejected,
			Verdict:                "FAILED_ROUTE_ZERO_SCALAR_IS_NATIVE_BUT_DOES_NOT_REPRODUCE_TRACEBRIDGE",
			Failures: []string{
				"FAILED_ROUTE_ZERO_SCALAR_DOES_NOT_REPRODUCE_TRACEBRIDGE",
			},
		},
	}
}

func Supports() []string {
	return []string{
		"CONDITIONAL_SUPPORT_S_SPLIT_TRANSPORT_TO_B2_REMAINS_STRONGLY_SUPPORTED",
		"CONDITIONAL_SUPPORT_S_SPLIT_HAS_H72_COMPATIBLE_SCALAR_TYPE",
		"CONDITIONAL_SUPPORT_D_BASE_HAS_DEFECT_RESPONSE_TYPE_BUT_REPARAMETERIZES_S_SPLIT",
		"CONDITIONAL_SUPPORT_7_OVER_72_IS_NATIVE_SHAPE_COEFFICIENT_NOT_SCALAR",
		"CONDITIONAL_SUPPORT_FINITE_RANK_DATA_EXPLAIN_ALPHA_COEFFICIENTS",
		"CONDITIONAL_SUPPORT_XI_BOUNDARY_RESIDUAL_PATTERN_EXPLAINS_SMALLNESS_PRESSURE",
		"CONDITIONAL_SUPPORT_NO_TESTED_NONCIRCULAR_PROXY_REPRODUCES_S_SPLIT",
		"CONDITIONAL_SUPPORT_CERTIFICATE_II_REDUCES_TO_SCALAR_SOURCE_SEAL",
	}
}

func Failures() []string {
	return []string{
		"FAILED_ROUTE_NO_NONCIRCULAR_NATIVE_S_SPLIT_PROXY_FOUND",
		"FAILED_ROUTE_D_BASE_REPLACEMENT_IS_REPARAMETERIZATION_OF_S_SPLIT",
		"FAILED_ROUTE_7_OVER_72_IS_NORMALIZATION_COEFFICIENT_NOT_S_SPLIT_SCALAR",
		"FAILED_ROUTE_FINITE_RANK_DATA_DO_NOT_CANONICALLY_GENERATE_S_SPLIT_MAGNITUDE",
		"FAILED_ROUTE_FIXED_POINT_SCALAR_RECOVERY_IS_CIRCULAR_WITH_TRACEBRIDGE_OUTPUT",
		"FAILED_ROUTE_XI_BOUNDARY_RESIDUAL_PROXY_NOT_NATIVE_H72_SCALAR_SOURCE",
		"FAILED_ROUTE_HISTORY_LOOP_UNIT_NOT_TYPED_TO_S_SPLIT_BOUNDARY_RESPONSE_SCALAR",
		"FAILED_ROUTE_ZERO_SCALAR_DOES_NOT_REPRODUCE_TRACEBRIDGE",
		"FAILED_ROUTE_S_SPLIT_REMAINS_BRIDGE_HISTORY_SCALAR_INPUT",
		"FAILED_ROUTE_CERTIFICATE_II_NOT_PASSED",
		"FAILED_ROUTE_NATIVE_R3_NOT_GRANTED",
		"FAILED_ROUTE_NO_OFFICIAL_LEDGER_UPDATE_ALLOWED",
		"FAILED_ROUTE_NO_INDIVIDUAL_YUKAWA_VALUES",
		"FAILED_ROUTE_NO_PHYSICAL_PARTICLE_ASSIGNMENT",
		"FAILED_ROUTE_NO_GENERATION_CARRIER_MAP",
		"FAILED_ROUTE_NO_FLAVOR_ORIENTATION_MAP",
		"FAILED_ROUTE_NO_R4_NATIVE_YUKAWA_SPECTRUM_THEOREM",
	}
}

func IsValidNativeReplacement(c ProxyCandidateAudit) bool {
	return c.FiniteSource && c.NonCircular && c.ScalarType && c.B2Compatible && c.MagnitudeCompatible && c.NoArbitraryCoefficient
}

func HasSuccessfulReplacement(candidates []ProxyCandidateAudit) bool {
	for _, c := range candidates {
		if IsValidNativeReplacement(c) {
			return true
		}
	}
	return false
}

func CandidateNames(candidates []ProxyCandidateAudit) []string {
	out := make([]string, 0, len(candidates))
	for _, c := range candidates {
		out = append(out, c.Name)
	}
	return out
}

func CandidateVerdicts(candidates []ProxyCandidateAudit) []string {
	out := make([]string, 0, len(candidates))
	for _, c := range candidates {
		out = append(out, c.Verdict)
	}
	return out
}

func CandidateFailures(candidates []ProxyCandidateAudit) []string {
	var out []string
	for _, c := range candidates {
		out = append(out, c.Failures...)
	}
	return out
}

func CandidateSupports(candidates []ProxyCandidateAudit) []string {
	var out []string
	for _, c := range candidates {
		out = append(out, c.Supports...)
	}
	return out
}

func FormatCandidate(c ProxyCandidateAudit) string {
	return fmt.Sprintf("%s: expr=%s finite=%v noncircular=%v scalar=%v b2=%v magnitude=%v noFit=%v status=%s verdict=%s", c.Name, c.CandidateExpression, c.FiniteSource, c.NonCircular, c.ScalarType, c.B2Compatible, c.MagnitudeCompatible, c.NoArbitraryCoefficient, c.Status, c.Verdict)
}

func FormatCriterion(c ReplacementCriterion) string {
	return fmt.Sprintf("%s=%v: %s", c.Name, c.Passed, c.Detail)
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

func appendAll(items ...[]string) []string {
	var out []string
	for _, group := range items {
		out = append(out, group...)
	}
	return out
}

func stringsJoin(items []string) string { return strings.Join(items, " | ") }

func almostEqual(a, b, eps float64) bool { return math.Abs(a-b) <= eps }
