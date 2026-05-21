// Package generation2nativer3promotiongapaudit implements
// Gate 938A: Native R3 Promotion Gap Audit.
//
// Gate 938A follows the Gate 937 bridge-level pre-test pass. Its purpose is
// protective rather than promotional: it records exactly which native theorem
// gaps still prevent the validated Z2 alpha/Yukawa trace bridge from being
// called native R3, and it keeps R4 generation/flavor/Yukawa-spectrum work out
// of the R3 promotion lane.
package generation2nativer3promotiongapaudit

import (
	"fmt"
	"strings"
)

const (
	AuditID         = "GATE938A-GENERATION2NATIVER3PROMOTIONGAPAUDIT"
	InheritedStatus = "R3_TRACEBRIDGE_PRETEST_PASSED"
	Classification  = "R3_TRACEBRIDGE_PRETEST_PASSED_NATIVE_PROMOTION_GAPS_EXPLICIT"
	ShortStatus     = "R3_PRETEST_PASSED_NATIVE_R3_BLOCKERS_EXPLICIT"
	FinalTruth      = "NATIVE_R3_PROMOTION_BLOCKED_BY_BOUNDARY_MEASURE_SOURCE_S_SPLIT_SOURCE_ADMISSIBLE_LATTICE_SOURCE_AND_FULL_AF_DESCENT"
	NextGate        = "NEXT_PRESSURE_GATE939_NATIVE_BOUNDARY_ACTIVATION_MEASURE_SOURCE_VS_S_SPLIT_SOURCE_PRIORITY_AUDIT"

	BridgeFormula = "alpha_B=(3/10)S_split+(7/72)S_split^2"
	TraceRows     = "rows=(rank 3, weight 1),(rank 3, weight alpha_B(1-alpha_B)),(rank 1, weight 3 alpha_B^2)"
	TraceFormula  = "N_eff=(3+3alpha_B)^2/(3+3alpha_B^2-6alpha_B^3+12alpha_B^4)"
)

const (
	Ssplit          = 0.0012924448188162962
	AlphaLinear     = 0.00038773344564488885
	AlphaQuadratic  = 0.0000001624013231638281
	AlphaB          = 0.0003878958469680527
	NEffOperator    = 3.002327375081808
	CYukawaOperator = 0.9992248096922658
	CHiggsOperator  = 1.037220510866514
)

type PromotionBlocker struct {
	Name           string
	RequiredGate   string
	CurrentSupport string
	Failure        string
	Primary        bool
}

type RetiredWound struct {
	Name   string
	Reason string
}

type R4BoundaryItem struct {
	Name    string
	Failure string
}

type Diagnostics struct {
	Ssplit          float64
	AlphaLinear     float64
	AlphaQuadratic  float64
	AlphaB          float64
	NEffOperator    float64
	CYukawaOperator float64
	CHiggsOperator  float64
}

type Analysis struct {
	AuditID        string
	Inherited      string
	Classification string
	ShortStatus    string
	Truth          string
	Diagnostics    Diagnostics
	Blockers       []PromotionBlocker
	RetiredWounds  []RetiredWound
	R4Boundary     []R4BoundaryItem
	Statuses       []string
	Supports       []string
	Failures       []string
	Final          string
}

func BuildDefault() (Analysis, error) {
	blockers := NativeR3Blockers()
	if len(blockers) != 4 || !allPrimary(blockers) {
		return Analysis{}, fmt.Errorf("native R3 blocker ledger must contain four primary blockers")
	}
	retired := RetiredPrimaryWounds()
	if len(retired) == 0 {
		return Analysis{}, fmt.Errorf("retired wound ledger is empty")
	}
	r4 := R4Boundary()
	if len(r4) == 0 {
		return Analysis{}, fmt.Errorf("R4 boundary ledger is empty")
	}
	return Analysis{
		AuditID:        AuditID,
		Inherited:      InheritedStatus,
		Classification: Classification,
		ShortStatus:    ShortStatus,
		Truth:          FinalTruth,
		Diagnostics:    DefaultDiagnostics(),
		Blockers:       blockers,
		RetiredWounds:  retired,
		R4Boundary:     r4,
		Statuses:       Statuses(),
		Supports:       Supports(),
		Failures:       Failures(),
		Final:          "Gate 938A prevents false native promotion after the Gate 937 pre-test pass: the R3 trace bridge is bridge-tested, while native R3 remains blocked by BoundaryActivationMeasure source, S_split response-parameter source, admissible support lattice source, and full A_F descent/orientation status.",
	}, nil
}

func DefaultDiagnostics() Diagnostics {
	return Diagnostics{Ssplit: Ssplit, AlphaLinear: AlphaLinear, AlphaQuadratic: AlphaQuadratic, AlphaB: AlphaB, NEffOperator: NEffOperator, CYukawaOperator: CYukawaOperator, CHiggsOperator: CHiggsOperator}
}

func NativeR3Blockers() []PromotionBlocker {
	return []PromotionBlocker{
		{
			Name:           "Native BoundaryActivationMeasure theorem",
			RequiredGate:   "NativeBoundaryActivationMeasureTheorem",
			CurrentSupport: "BoundaryActivationMeasure passed the closure-factored bridge pre-test and rejected false measures",
			Failure:        "FAILED_ROUTE_NO_NATIVE_BOUNDARY_ACTIVATION_MEASURE_THEOREM",
			Primary:        true,
		},
		{
			Name:           "Native S_split response-parameter theorem",
			RequiredGate:   "SsplitResponseParameterTheorem",
			CurrentSupport: "single insertion S_split -> s into (1+s b1)(1+s b2) is pre-test valid and s^2 follows from exterior multiplication",
			Failure:        "FAILED_ROUTE_NO_NATIVE_S_SPLIT_RESPONSE_PARAMETER_THEOREM",
			Primary:        true,
		},
		{
			Name:           "Native admissible airlock support lattice theorem",
			RequiredGate:   "NativeAdmissibleAirlockSupportLatticeTheorem",
			CurrentSupport: "tensor-structured admissibility and unique support lattice passed bridge pre-test with orphan and arbitrary subspace routes rejected",
			Failure:        "FAILED_ROUTE_NO_NATIVE_ADMISSIBLE_AIRLOCK_SUPPORT_LATTICE_THEOREM",
			Primary:        true,
		},
		{
			Name:           "Full A_F descent or lawful spontaneous-orientation theorem",
			RequiredGate:   "FullAFDescentOrSpontaneousOrientationTheorem",
			CurrentSupport: "post-orientation Z2 trace ledger is stable and representative-independent",
			Failure:        "FAILED_ROUTE_FULL_A_F_DESCENT_STILL_BLOCKED",
			Primary:        true,
		},
	}
}

func RetiredPrimaryWounds() []RetiredWound {
	return []RetiredWound{
		{Name: "lambda versus barlambda representative", Reason: "absorbed by Z2 orientation-class trace ledger"},
		{Name: "+Q_phi versus -Q_phi", Reason: "phase sign is no longer a primary R3 blocker at bridge level"},
		{Name: "representative-dependent alpha", Reason: "alpha_B is formulated on the Z2 class"},
		{Name: "cross-lane pollution", Reason: "negative pre-test rejects false degree-target lanes"},
		{Name: "Theta(2)=F_2/F_1", Reason: "associated-graded top target rejected; fixed-base cumulative target F_2/F_0 passed"},
		{Name: "bare denominators", Reason: "negative pre-test rejects 8 and 70 in favor of boundary-augmented 10 and 72"},
		{Name: "orphan support fragments", Reason: "tensor-structured admissible lattice rejects opposite-socket orphan fragments"},
		{Name: "arbitrary rank-compatible subspaces", Reason: "admissible supports are tensor-structured completions"},
	}
}

func R4Boundary() []R4BoundaryItem {
	return []R4BoundaryItem{
		{Name: "generation carrier", Failure: "FAILED_ROUTE_NO_GENERATION_CARRIER_MAP"},
		{Name: "flavor orientation", Failure: "FAILED_ROUTE_NO_FLAVOR_ORIENTATION_MAP"},
		{Name: "individual Yukawa eigenvalues", Failure: "FAILED_ROUTE_NO_INDIVIDUAL_YUKAWA_VALUES"},
		{Name: "physical particle assignment", Failure: "FAILED_ROUTE_NO_PHYSICAL_PARTICLE_ASSIGNMENT"},
		{Name: "native Yukawa operator theorem", Failure: "FAILED_ROUTE_NO_NATIVE_YUKAWA_OPERATOR_THEOREM"},
	}
}

func Statuses() []string {
	return []string{
		"PASS_GATE938A_R3_TRACEBRIDGE_PRETEST_INHERITED",
		"PASS_GATE938A_NATIVE_PROMOTION_BLOCKERS_EXPLICIT",
		"PASS_GATE938A_OLD_WOUNDS_RETIRED_FROM_PRIMARY_STATUS",
		"PASS_GATE938A_R4_BOUNDARY_PRESERVED",
		"FIREWALL_PRESERVED_GATE938A_NOT_NATIVE_R3",
	}
}

func Supports() []string {
	return []string{
		"CONDITIONAL_SUPPORT_R3_TRACEBRIDGE_PRETEST_PASSED",
		"CONDITIONAL_SUPPORT_ALPHA_B_VALIDATED_AS_Z2_CLOSURE_FACTORED_BRIDGE_MEASURE",
		"CONDITIONAL_SUPPORT_Y_DAGGER_Y_TRACE_ROWS_RECONSTRUCT_OPERATOR_N_EFF",
		"CONDITIONAL_SUPPORT_FALSE_ALPHA_ROUTES_REJECTED_BY_NEGATIVE_TESTS",
		"CONDITIONAL_SUPPORT_PHASE_SIGN_AND_REPRESENTATIVE_ALPHA_NO_LONGER_PRIMARY_BLOCKERS",
		"CONDITIONAL_SUPPORT_NATIVE_R3_PROMOTION_REQUIREMENTS_NOW_EXPLICIT",
		"CONDITIONAL_SUPPORT_R4_GENERATION_FLAVOR_BRANCH_REMAINS_SEPARATE",
	}
}

func Failures() []string {
	return []string{
		"FAILED_ROUTE_NOT_NATIVE_R3",
		"FAILED_ROUTE_NO_NATIVE_BOUNDARY_ACTIVATION_MEASURE_THEOREM",
		"FAILED_ROUTE_NO_NATIVE_S_SPLIT_RESPONSE_PARAMETER_THEOREM",
		"FAILED_ROUTE_NO_NATIVE_TRANSPORT_OF_S_SPLIT_TO_BOUNDARY_PAIR_RESPONSE",
		"FAILED_ROUTE_NO_NATIVE_ADMISSIBLE_AIRLOCK_SUPPORT_LATTICE_THEOREM",
		"FAILED_ROUTE_TENSOR_STRUCTURED_ADMISSIBILITY_NOT_NATIVE_ASHA_THEOREM",
		"FAILED_ROUTE_FULL_A_F_DESCENT_STILL_BLOCKED",
		"FAILED_ROUTE_HIGGS_ORIENTATION_REMAINS_SEALED_AS_ORIENTATION_CLASS",
		"FAILED_ROUTE_POST_ORIENTATION_LEDGER_NOT_FULL_NATIVE_A_F_SECTOR_LEDGER",
		"FAILED_ROUTE_NO_OFFICIAL_LEDGER_UPDATE_ALLOWED",
		"FAILED_ROUTE_NO_INDIVIDUAL_YUKAWA_VALUES",
		"FAILED_ROUTE_NO_PHYSICAL_PARTICLE_ASSIGNMENT",
		"FAILED_ROUTE_NO_GENERATION_CARRIER_MAP",
		"FAILED_ROUTE_NO_FLAVOR_ORIENTATION_MAP",
		"FAILED_ROUTE_NO_NATIVE_YUKAWA_OPERATOR_THEOREM",
	}
}

func allPrimary(bs []PromotionBlocker) bool {
	for _, b := range bs {
		if !b.Primary || b.Failure == "" || b.RequiredGate == "" {
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

func blockerFailures(bs []PromotionBlocker) []string {
	out := make([]string, 0, len(bs))
	for _, b := range bs {
		out = append(out, b.Failure)
	}
	return out
}

func r4Failures(rs []R4BoundaryItem) []string {
	out := make([]string, 0, len(rs))
	for _, r := range rs {
		out = append(out, r.Failure)
	}
	return out
}

func FormatDiagnostics(d Diagnostics) string {
	return fmt.Sprintf("S_split=%.17g alpha_linear=%.17g alpha_quad=%.17g alpha_B=%.17g N_eff=%.17g C_Yukawa=%.17g C_Higgs=%.17g", d.Ssplit, d.AlphaLinear, d.AlphaQuadratic, d.AlphaB, d.NEffOperator, d.CYukawaOperator, d.CHiggsOperator)
}

func FormatBlockers(bs []PromotionBlocker) string {
	parts := make([]string, 0, len(bs))
	for _, b := range bs {
		parts = append(parts, fmt.Sprintf("%s -> %s", b.Name, b.Failure))
	}
	return strings.Join(parts, "; ")
}

func FormatRetired(rs []RetiredWound) string {
	parts := make([]string, 0, len(rs))
	for _, r := range rs {
		parts = append(parts, r.Name+": "+r.Reason)
	}
	return strings.Join(parts, "; ")
}

func FormatR4(rs []R4BoundaryItem) string {
	parts := make([]string, 0, len(rs))
	for _, r := range rs {
		parts = append(parts, r.Name+" -> "+r.Failure)
	}
	return strings.Join(parts, "; ")
}

func FormatFirewalls(f []string) string { return strings.Join(f, ";") }
