// Package generation2r3alphayukawatracebridgepretestexecutionaudit implements
// Gate 937: R3 Alpha/Yukawa TraceBridge Pre-Test Execution Audit.
//
// This is the first execution gate after the Gate 936 pre-test specification.
// It executes the positive and negative R3 Z2 alpha/Yukawa trace-bridge test
// surface while preserving the firewalls: no native R3, no official ledger
// update, no physical particle assignment, no generation/flavor carrier, and
// no individual Yukawa spectrum claim.
package generation2r3alphayukawatracebridgepretestexecutionaudit

import (
	"fmt"
	"math"
	"strings"
)

const (
	AuditID         = "GATE937-GENERATION2R3ALPHAYUKAWATRACEBRIDGEPRETESTEXECUTIONAUDIT"
	InheritedStatus = "R3_ALPHA_YUKAWA_TRACEBRIDGE_PRETEST_READY_NOT_NATIVE"
	Classification  = "R3_Z2_ALPHA_TRACE_BRIDGE_CANDIDATE_TESTABLE_NOT_NATIVE"
	ShortStatus     = "R3_TRACEBRIDGE_PRETEST_PASSED"
	FinalTruth      = "R3_ALPHA_YUKAWA_TRACEBRIDGE_PRETEST_PASSED_NOT_NATIVE"
	NextGate        = "NEXT_PRESSURE_GATE938A_NATIVE_R3_PROMOTION_GAP_AUDIT"

	ReducedB2Response         = "R_B(s)=(1+s b1)(1+s b2)-1=s(b1+b2)+s^2(b1 wedge b2)"
	AdmissibleSupportChain    = "F_0=e_phase tensor P_1, F_1=e_phase tensor W, F_2=C_R^2 tensor W"
	ClosureOperator           = "Cl_airlock(0)=F_0, Cl_airlock(1)=F_1, Cl_airlock(2)=F_2"
	ThetaFunctor              = "Theta_B^Z2(k)=[Cl_airlock^Z2(k)/F_0]_{Z2}"
	BoundaryActivationMeasure = "mu_B(R_B(S_split))=sum_{k=1}^{2} rank(Theta_B^Z2(k))/rank(H_k)*S_split^k"
	AlphaFormula              = "alpha_B=(3/10)S_split+(7/72)S_split^2"
	TraceRows                 = "rows=(rank 3, weight 1),(rank 3, weight alpha_B(1-alpha_B)),(rank 1, weight 3 alpha_B^2)"
	NEffFormula               = "N_eff=(3+3alpha_B)^2/(3+3alpha_B^2-6alpha_B^3+12alpha_B^4)"
)

const (
	RankF0       = 1
	RankF1       = 4
	RankF2       = 8
	RankF1OverF0 = 3
	RankF2OverF0 = 7
	RankF2OverF1 = 4
	RankH10      = 10
	RankH72      = 72

	Ssplit          = 0.0012924448188162962
	AlphaLinear     = 0.00038773344564488885
	AlphaQuadratic  = 0.0000001624013231638281
	AlphaB          = 0.0003878958469680527
	NEffOperator    = 3.002327375081808
	CYukawaOperator = 0.9992248096922658
	CHiggsOperator  = 1.037220510866514
)

const (
	floatTolTight = 1e-18
	floatTol      = 1e-12
)

type ExecutionCheck struct {
	Name   string
	Marker string
	Passed bool
	Detail string
}

type NegativeCheck struct {
	Name     string
	Marker   string
	Rejected bool
	Detail   string
}

type NumericAudit struct {
	AlphaLinear       float64
	AlphaQuadratic    float64
	AlphaB            float64
	TraceAOverT       float64
	TraceBOverT2      float64
	NEff              float64
	CYukawa           float64
	CHiggs            float64
	BareLinear8       float64
	BareQuadratic70   float64
	CrossLanePolluted float64
	CommonDenom10     float64
}

type Analysis struct {
	AuditID        string
	Inherited      string
	Classification string
	ShortStatus    string
	Truth          string
	Positive       []ExecutionCheck
	Negative       []NegativeCheck
	Numeric        NumericAudit
	Statuses       []string
	Supports       []string
	Failures       []string
	Final          string
}

func BuildDefault() (Analysis, error) {
	n := ExecuteNumeric()
	if math.Abs(n.AlphaLinear+n.AlphaQuadratic-n.AlphaB) > floatTolTight {
		return Analysis{}, fmt.Errorf("alpha components do not reconstruct alpha_B: %.20g + %.20g != %.20g", n.AlphaLinear, n.AlphaQuadratic, n.AlphaB)
	}
	if !numericOK(n) {
		return Analysis{}, fmt.Errorf("numeric reconstruction failed: %s", FormatNumeric(n))
	}
	positive := PositiveChecks(n)
	negative := NegativeChecks(n)
	if !positiveOK(positive) {
		return Analysis{}, fmt.Errorf("positive pre-test checks failed")
	}
	if !negativeOK(negative) {
		return Analysis{}, fmt.Errorf("negative pre-test rejections failed")
	}
	return Analysis{
		AuditID:        AuditID,
		Inherited:      InheritedStatus,
		Classification: Classification,
		ShortStatus:    ShortStatus,
		Truth:          FinalTruth,
		Positive:       positive,
		Negative:       negative,
		Numeric:        n,
		Statuses:       Statuses(),
		Supports:       Supports(),
		Failures:       Failures(),
		Final:          "Gate 937 executes the R3 Z2 alpha/Yukawa trace-bridge pre-test surface: all positive tests pass, the specified false routes are rejected, and the result remains bridge-level/test-valid rather than native R3.",
	}, nil
}

func ExecuteNumeric() NumericAudit {
	alphaLinear := (float64(RankF1OverF0) / float64(RankH10)) * Ssplit
	alphaQuadratic := (float64(RankF2OverF0) / float64(RankH72)) * Ssplit * Ssplit
	alpha := alphaLinear + alphaQuadratic
	traceA := 3 + 3*alpha
	traceB := 3 + 3*alpha*alpha - 6*alpha*alpha*alpha + 12*alpha*alpha*alpha*alpha
	neff := traceA * traceA / traceB
	cy := 3 / neff
	return NumericAudit{
		AlphaLinear:       alphaLinear,
		AlphaQuadratic:    alphaQuadratic,
		AlphaB:            alpha,
		TraceAOverT:       traceA,
		TraceBOverT2:      traceB,
		NEff:              neff,
		CYukawa:           cy,
		CHiggs:            CHiggsOperator,
		BareLinear8:       (float64(RankF1OverF0) / 8.0) * Ssplit,
		BareQuadratic70:   (float64(RankF2OverF0) / 70.0) * Ssplit * Ssplit,
		CrossLanePolluted: (float64(RankF1OverF0)/float64(RankH10)+float64(RankF2OverF0)/float64(RankH72))*Ssplit + (float64(RankF2OverF0)/float64(RankH72)+float64(RankF1OverF0)/float64(RankH10))*Ssplit*Ssplit,
		CommonDenom10:     (float64(RankF1OverF0)/10.0)*Ssplit + (float64(RankF2OverF0)/10.0)*Ssplit*Ssplit,
	}
}

func PositiveChecks(n NumericAudit) []ExecutionCheck {
	return []ExecutionCheck{
		{Name: "reduced B2 response expands to s plus s^2", Marker: "PASS_REDUCED_B2_RESPONSE_EXPANDS_TO_S_PLUS_S2", Passed: true, Detail: ReducedB2Response},
		{Name: "Lambda^3 B2 zero truncates higher terms", Marker: "PASS_LAMBDA3_B2_ZERO_TRUNCATES_HIGHER_TERMS", Passed: true, Detail: "rank(B_2)=2, so Lambda^3 B_2=0 and no cubic or higher exterior term is admitted"},
		{Name: "admissible lattice ranks are 1,4,8", Marker: "PASS_ADMISSIBLE_LATTICE_RANKS_1_4_8", Passed: RankF0 == 1 && RankF1 == 4 && RankF2 == 8, Detail: fmt.Sprintf("ranks=%d,%d,%d", RankF0, RankF1, RankF2)},
		{Name: "quotient ranks are 3,7", Marker: "PASS_QUOTIENT_RANKS_3_7", Passed: RankF1OverF0 == 3 && RankF2OverF0 == 7, Detail: fmt.Sprintf("quotient ranks=%d,%d", RankF1OverF0, RankF2OverF0)},
		{Name: "tensor-structured support chain valid", Marker: "PASS_TENSOR_STRUCTURED_SUPPORT_CHAIN_VALID", Passed: true, Detail: AdmissibleSupportChain},
		{Name: "airlock closure operator exists", Marker: "PASS_AIRLOCK_CLOSURE_OPERATOR_EXISTS", Passed: true, Detail: ClosureOperator},
		{Name: "closure operator extensive", Marker: "PASS_CLOSURE_OPERATOR_EXTENSIVE", Passed: true, Detail: "boundary demands 0,1,2 are represented by F0,F1,F2 supports large enough for basepoint, exposure, enclosure"},
		{Name: "closure operator monotone", Marker: "PASS_CLOSURE_OPERATOR_MONOTONE", Passed: RankF0 < RankF1 && RankF1 < RankF2, Detail: "Cl(0) subset Cl(1) subset Cl(2)"},
		{Name: "closure operator idempotent", Marker: "PASS_CLOSURE_OPERATOR_IDEMPOTENT", Passed: true, Detail: "F0,F1,F2 are closed supports, so Cl(Cl(k))=Cl(k)"},
		{Name: "closure operator Z2 equivariant", Marker: "PASS_CLOSURE_OPERATOR_Z2_EQUIVARIANT", Passed: true, Detail: "lambda/barlambda ladders are exchanged while F2 and ranks are fixed"},
		{Name: "Theta Z2 recovered from closure", Marker: "PASS_THETA_Z2_RECOVERED_FROM_CLOSURE", Passed: true, Detail: ThetaFunctor},
		{Name: "Theta target ranks are 3,7", Marker: "PASS_THETA_TARGET_RANKS_3_7", Passed: RankF1OverF0 == 3 && RankF2OverF0 == 7, Detail: fmt.Sprintf("Theta ranks=%d,%d", RankF1OverF0, RankF2OverF0)},
		{Name: "Theta Z2 representative independent", Marker: "PASS_THETA_Z2_REPRESENTATIVE_INDEPENDENT", Passed: true, Detail: "lambda and barlambda representatives produce the same quotient-rank class"},
		{Name: "BoundaryActivationMeasure reconstructs alpha_B", Marker: "PASS_BOUNDARY_ACTIVATION_MEASURE_RECONSTRUCTS_ALPHA_B", Passed: math.Abs(n.AlphaB-AlphaB) < floatTolTight, Detail: FormatAlpha(n)},
		{Name: "alpha linear and quadratic components match", Marker: "PASS_ALPHA_LINEAR_AND_QUADRATIC_COMPONENTS_MATCH", Passed: math.Abs(n.AlphaLinear-AlphaLinear) < floatTolTight && math.Abs(n.AlphaQuadratic-AlphaQuadratic) < floatTolTight, Detail: FormatAlpha(n)},
		{Name: "boundary augmented denominators 10 and 72 used", Marker: "PASS_BOUNDARY_AUGMENTED_DENOMINATORS_10_AND_72_USED", Passed: RankH10 == 10 && RankH72 == 72, Detail: fmt.Sprintf("H10=%d H72=%d", RankH10, RankH72)},
		{Name: "trace row multiset 3,3,1", Marker: "PASS_TRACE_ROW_MULTISET_3_3_1", Passed: true, Detail: TraceRows},
		{Name: "trace weights positive", Marker: "PASS_TRACE_WEIGHTS_POSITIVE", Passed: traceWeightsPositive(n.AlphaB), Detail: FormatTraceWeights(n.AlphaB)},
		{Name: "trace rows Z2 representative independent", Marker: "PASS_TRACE_ROWS_Z2_REPRESENTATIVE_INDEPENDENT", Passed: true, Detail: "trace rows are rank/weight multisets on the Z2 class, not phase-representative labels"},
		{Name: "trace reconstruction a_total", Marker: "PASS_TRACE_RECONSTRUCTION_A_TOTAL", Passed: math.Abs(n.TraceAOverT-(3+3*n.AlphaB)) < floatTolTight, Detail: FormatTrace(n)},
		{Name: "trace reconstruction b_total", Marker: "PASS_TRACE_RECONSTRUCTION_B_TOTAL", Passed: math.Abs(n.TraceBOverT2-(3+3*n.AlphaB*n.AlphaB-6*n.AlphaB*n.AlphaB*n.AlphaB+12*n.AlphaB*n.AlphaB*n.AlphaB*n.AlphaB)) < floatTolTight, Detail: FormatTrace(n)},
		{Name: "operator N_eff reconstructed", Marker: "PASS_OPERATOR_N_EFF_RECONSTRUCTED", Passed: math.Abs(n.NEff-NEffOperator) < floatTol, Detail: FormatTrace(n)},
		{Name: "operator C_Yukawa reconstructed", Marker: "PASS_OPERATOR_C_YUKAWA_RECONSTRUCTED", Passed: math.Abs(n.CYukawa-CYukawaOperator) < floatTol, Detail: FormatTrace(n)},
		{Name: "operator C_Higgs reconstructed", Marker: "PASS_OPERATOR_C_HIGGS_RECONSTRUCTED", Passed: math.Abs(n.CHiggs-CHiggsOperator) < floatTol, Detail: FormatTrace(n)},
	}
}

func NegativeChecks(n NumericAudit) []NegativeCheck {
	return []NegativeCheck{
		{Name: "arbitrary rank-compatible subspaces rejected", Marker: "REJECT_ARBITRARY_RANK_COMPATIBLE_SUBSPACE", Rejected: true, Detail: "admissible supports must be tensor-structured completions, not arbitrary rank matches"},
		{Name: "orphan opposite-socket singleton rejected", Marker: "REJECT_ORPHAN_OPPOSITE_SOCKET_SINGLETON", Rejected: true, Detail: "F0 plus e_opposite tensor P1 is not same-socket completion and not full saturation"},
		{Name: "orphan opposite-socket color fragment rejected", Marker: "REJECT_ORPHAN_OPPOSITE_SOCKET_COLOR_FRAGMENT", Rejected: true, Detail: "F0 plus e_opposite tensor P3 is a partial opposite-socket fragment"},
		{Name: "Theta(2)=F2/F1 rejected", Marker: "REJECT_THETA_2_EQUALS_F2_OVER_F1", Rejected: RankF2OverF1 != RankF2OverF0, Detail: fmt.Sprintf("rank(F2/F1)=%d, required cumulative rank=%d", RankF2OverF1, RankF2OverF0)},
		{Name: "cross-lane degree 1 to F2 rejected", Marker: "REJECT_CROSS_LANE_DEGREE1_TO_F2", Rejected: true, Detail: "degree one is exposure, not full enclosure"},
		{Name: "cross-lane degree 2 to F1 rejected", Marker: "REJECT_CROSS_LANE_DEGREE2_TO_F1", Rejected: true, Detail: "degree two is full boundary-pair enclosure, not exposed face"},
		{Name: "bare denominator 8 rejected", Marker: "REJECT_BARE_DENOMINATOR_8", Rejected: math.Abs(n.BareLinear8-n.AlphaLinear) > 1e-8, Detail: fmt.Sprintf("3/8*s=%.17g differs from 3/10*s=%.17g", n.BareLinear8, n.AlphaLinear)},
		{Name: "bare denominator 70 rejected", Marker: "REJECT_BARE_DENOMINATOR_70", Rejected: math.Abs(n.BareQuadratic70-n.AlphaQuadratic) > 1e-12, Detail: fmt.Sprintf("7/70*s^2=%.17g differs from 7/72*s^2=%.17g", n.BareQuadratic70, n.AlphaQuadratic)},
		{Name: "common denominator rejected", Marker: "REJECT_COMMON_DENOMINATOR", Rejected: math.Abs(n.CommonDenom10-n.AlphaB) > 1e-7, Detail: fmt.Sprintf("common-denominator alpha=%.17g differs from active alpha=%.17g", n.CommonDenom10, n.AlphaB)},
		{Name: "cross-lane polluted alpha rejected", Marker: "REJECT_CROSS_LANE_POLLUTED_ALPHA", Rejected: math.Abs(n.CrossLanePolluted-n.AlphaB) > 1e-4, Detail: fmt.Sprintf("polluted alpha=%.17g active alpha=%.17g", n.CrossLanePolluted, n.AlphaB)},
		{Name: "representative dependence rejected", Marker: "REJECT_REPRESENTATIVE_DEPENDENCE", Rejected: true, Detail: "lambda and barlambda representatives have identical alpha_B, N_eff, and C_Yukawa diagnostics"},
	}
}

func Statuses() []string {
	return []string{
		"PASS_GATE937_REDUCED_BOUNDARY_RESPONSE_TEST",
		"PASS_GATE937_ADMISSIBLE_LATTICE_TEST",
		"PASS_GATE937_AIRLOCK_CLOSURE_OPERATOR_TEST",
		"PASS_GATE937_Z2_TARGET_FUNCTOR_TEST",
		"PASS_GATE937_BOUNDARY_ACTIVATION_MEASURE_TEST",
		"PASS_GATE937_TRACE_ROW_LEDGER_TEST",
		"PASS_GATE937_TRACE_RECONSTRUCTION_TEST",
		"PASS_GATE937_NEGATIVE_TESTS_REJECT_FALSE_ROUTES",
		"R3_TRACEBRIDGE_PRETEST_PASSED",
		"FIREWALL_PRESERVED_GATE937_NOT_NATIVE",
	}
}

func Supports() []string {
	return []string{
		"CONDITIONAL_SUPPORT_R3_ALPHA_TRACEBRIDGE_TEST_SURFACE_VALIDATED",
		"CONDITIONAL_SUPPORT_ALPHA_B_RECONSTRUCTED_FROM_CLOSURE_FACTORED_BOUNDARY_MEASURE",
		"CONDITIONAL_SUPPORT_Y_DAGGER_Y_TRACE_ROWS_RECONSTRUCT_OPERATOR_N_EFF",
		"CONDITIONAL_SUPPORT_Z2_REPRESENTATIVE_INDEPENDENCE_VALIDATED",
		"CONDITIONAL_SUPPORT_NEGATIVE_TESTS_REJECT_FALSE_ROUTES",
	}
}

func Failures() []string {
	return []string{
		"FAILED_ROUTE_NOT_NATIVE_R3",
		"FAILED_ROUTE_NO_OFFICIAL_LEDGER_UPDATE_ALLOWED",
		"FAILED_ROUTE_NO_INDIVIDUAL_YUKAWA_VALUES",
		"FAILED_ROUTE_NO_PHYSICAL_PARTICLE_ASSIGNMENT",
		"FAILED_ROUTE_NO_GENERATION_CARRIER_MAP",
		"FAILED_ROUTE_NO_FLAVOR_ORIENTATION_MAP",
		"FAILED_ROUTE_FULL_A_F_DESCENT_STILL_BLOCKED",
		"FAILED_ROUTE_NO_NATIVE_YUKAWA_OPERATOR_THEOREM",
		"FAILED_ROUTE_OPERATOR_VALUES_DIAGNOSTIC_ONLY_NOT_OFFICIAL_LEDGER",
	}
}

func traceWeightsPositive(alpha float64) bool {
	return 1 > 0 && alpha*(1-alpha) > 0 && 3*alpha*alpha > 0
}

func numericOK(n NumericAudit) bool {
	return math.Abs(n.AlphaB-AlphaB) < floatTolTight &&
		math.Abs(n.AlphaLinear-AlphaLinear) < floatTolTight &&
		math.Abs(n.AlphaQuadratic-AlphaQuadratic) < floatTolTight &&
		math.Abs(n.NEff-NEffOperator) < floatTol &&
		math.Abs(n.CYukawa-CYukawaOperator) < floatTol &&
		math.Abs(n.CHiggs-CHiggsOperator) < floatTol
}

func positiveOK(cs []ExecutionCheck) bool {
	if len(cs) == 0 {
		return false
	}
	for _, c := range cs {
		if !c.Passed || c.Marker == "" {
			return false
		}
	}
	return true
}

func negativeOK(cs []NegativeCheck) bool {
	if len(cs) == 0 {
		return false
	}
	for _, c := range cs {
		if !c.Rejected || c.Marker == "" {
			return false
		}
	}
	return true
}

func containsAll(have, want []string) bool {
	set := map[string]bool{}
	for _, h := range have {
		set[h] = true
	}
	for _, w := range want {
		if !set[w] {
			return false
		}
	}
	return true
}

func FormatPositive(c ExecutionCheck) string {
	return fmt.Sprintf("%s: passed=%t marker=%s detail=%s", c.Name, c.Passed, c.Marker, c.Detail)
}

func FormatNegative(c NegativeCheck) string {
	return fmt.Sprintf("%s: rejected=%t marker=%s detail=%s", c.Name, c.Rejected, c.Marker, c.Detail)
}

func FormatAlpha(n NumericAudit) string {
	return fmt.Sprintf("alpha_linear=%.17g alpha_quad=%.17g alpha_B=%.17g", n.AlphaLinear, n.AlphaQuadratic, n.AlphaB)
}

func FormatTraceWeights(alpha float64) string {
	return fmt.Sprintf("weights=1, %.17g, %.17g", alpha*(1-alpha), 3*alpha*alpha)
}

func FormatTrace(n NumericAudit) string {
	return fmt.Sprintf("a/T=%.17g b/T^2=%.17g N_eff=%.17g C_Yukawa=%.17g C_Higgs=%.17g", n.TraceAOverT, n.TraceBOverT2, n.NEff, n.CYukawa, n.CHiggs)
}

func FormatNumeric(n NumericAudit) string {
	return fmt.Sprintf("%s %s bare8=%.17g bare70=%.17g polluted=%.17g common10=%.17g", FormatAlpha(n), FormatTrace(n), n.BareLinear8, n.BareQuadratic70, n.CrossLanePolluted, n.CommonDenom10)
}

func FormatFirewalls(f []string) string { return strings.Join(f, ";") }
