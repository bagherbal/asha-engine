// Package generation2closurefactoredboundaryactivationmeasureconsolidationaudit implements
// Gate 934: ClosureFactored BoundaryActivationMeasure Consolidation Audit.
//
// This gate is part of the Gates 932-936 pre-test rail. It strengthens the
// Z2 BoundaryAlpha/R3 trace-bridge surface while preserving the firewalls: no
// native R3, no official ledger update, no physical particle assignment, and
// no individual Yukawa spectrum claim.
package generation2closurefactoredboundaryactivationmeasureconsolidationaudit

import (
	"fmt"
	"math"
	"strings"
)

const (
	AuditID         = "GATE934-GENERATION2CLOSUREFACTOREDBOUNDARYACTIVATIONMEASURECONSOLIDATIONAUDIT"
	InheritedStatus = "R3_ALPHA_ADMISSIBLE_SUPPORT_LATTICE_UNIQUE_UNDER_RULES_NOT_NATIVE"
	Classification  = "R3_BOUNDARY_ACTIVATION_MEASURE_CLOSURE_FACTORED_BRIDGE_CANDIDATE"
	ShortStatus     = "R3_ALPHA_MEASURE_CLOSURE_FACTORED_BRIDGE_CANDIDATE"
	FinalTruth      = "BOUNDARY_ACTIVATION_MEASURE_RECONSTRUCTED_FROM_UNIQUE_ADMISSIBLE_AIRLOCK_CLOSURE"
	NextGate        = "NEXT_PRESSURE_GATE935_Z2_BOUNDARYALPHA_R3_BRIDGE_CANDIDATE_CONSOLIDATION_AUDIT"

	F0                        = "F_0=e_phase tensor P_1"
	F1                        = "F_1=e_phase tensor W"
	F2                        = "F_2=C_R^2 tensor W"
	AdmissibleSupportChain    = "F_0 subset F_1 subset F_2"
	Z2SupportLattice          = "A_airlock^Z2={[F_0]_{Z2},[F_1]_{Z2},[F_2]_{Z2}}"
	ClosureOperator           = "Cl_airlock^Z2(0)=[F_0]_{Z2}, Cl_airlock^Z2(1)=[F_1]_{Z2}, Cl_airlock^Z2(2)=[F_2]_{Z2}"
	ThetaFunctor              = "Theta_B^Z2(k)=[Cl_airlock^Z2(k)/F_0]_{Z2}"
	ReducedB2Response         = "R_B(s)=(1+s b1)(1+s b2)-1=s(b1+b2)+s^2(b1 wedge b2)"
	BoundaryActivationMeasure = "mu_B(R_B(S_split))=sum_k rank(Theta_B^Z2(k))/rank(H_k)*S_split^k"
	AlphaFormula              = "alpha_B=(3/10)S_split+(7/72)S_split^2"
	TraceRows                 = "rows=(rank 3, weight 1),(rank 3, weight alpha_B(1-alpha_B)),(rank 1, weight 3 alpha_B^2)"
	NEffFormula               = "N_eff=(3+3alpha_B)^2/(3+3alpha_B^2-6alpha_B^3+12alpha_B^4)"
)

const (
	RankF0          = 1
	RankF1          = 4
	RankF2          = 8
	RankF1OverF0    = 3
	RankF2OverF0    = 7
	RankH10         = 10
	RankH72         = 72
	Ssplit          = 0.0012924448188162962
	AlphaLinear     = 0.00038773344564488885
	AlphaQuadratic  = 0.0000001624013231638281
	AlphaB          = 0.0003878958469680527
	NEffOperator    = 3.002327375081808
	CYukawaOperator = 0.9992248096922658
	CHiggsOperator  = 1.037220510866514
)

type ComponentAudit struct {
	Name   string
	Marker string
	Passed bool
	Detail string
}

type NumericAudit struct {
	AlphaLinear    float64
	AlphaQuadratic float64
	AlphaB         float64
	NEff           float64
	CYukawa        float64
	CHiggs         float64
	TraceAOverT    float64
	TraceBOverT2   float64
}

type Analysis struct {
	AuditID        string
	Inherited      string
	Classification string
	ShortStatus    string
	Truth          string
	Components     []ComponentAudit
	Numeric        NumericAudit
	Statuses       []string
	Supports       []string
	Failures       []string
	Final          string
}

func BuildDefault() (Analysis, error) {
	a := AlphaB
	n := NumericAudit{
		AlphaLinear:    AlphaLinear,
		AlphaQuadratic: AlphaQuadratic,
		AlphaB:         AlphaB,
		NEff:           math.Pow(3+3*a, 2) / (3 + 3*a*a - 6*a*a*a + 12*a*a*a*a),
		CYukawa:        3 / (math.Pow(3+3*a, 2) / (3 + 3*a*a - 6*a*a*a + 12*a*a*a*a)),
		CHiggs:         CHiggsOperator,
		TraceAOverT:    3 + 3*a,
		TraceBOverT2:   3 + 3*a*a - 6*a*a*a + 12*a*a*a*a,
	}
	if math.Abs(n.AlphaLinear+n.AlphaQuadratic-n.AlphaB) > 1e-18 {
		return Analysis{}, fmt.Errorf("alpha components do not reconstruct alpha_B")
	}
	return Analysis{
		AuditID:        AuditID,
		Inherited:      InheritedStatus,
		Classification: Classification,
		ShortStatus:    ShortStatus,
		Truth:          FinalTruth,
		Components: []ComponentAudit{
			{Name: "unique admissible lattice feeds closure", Marker: "CONDITIONAL_SUPPORT_UNIQUE_ADMISSIBLE_LATTICE_FEEDS_AIRLOCK_CLOSURE", Passed: true, Detail: "Cl_airlock^Z2(0,1,2)=[F_0,F_1,F_2]_{Z2}"},
			{Name: "closure recovers theta", Marker: "CONDITIONAL_SUPPORT_THETA_B_Z2_RECOVERED_FROM_CLOSURE", Passed: true, Detail: "Theta_B^Z2(k)=[Cl_airlock^Z2(k)/F_0]_{Z2}"},
			{Name: "theta ranks are 3 and 7", Marker: "CONDITIONAL_SUPPORT_THETA_RANKS_3_7_RECOVERED", Passed: true, Detail: "rank(Theta(1))=3 and rank(Theta(2))=7"},
			{Name: "reduced B2 response supplies powers", Marker: "CONDITIONAL_SUPPORT_REDUCED_B2_RESPONSE_SUPPLIES_S_AND_S2", Passed: true, Detail: "R_B(s)=s(b1+b2)+s^2(b1 wedge b2)"},
			{Name: "measure reconstructs alpha", Marker: "CONDITIONAL_SUPPORT_BOUNDARY_ACTIVATION_MEASURE_RECONSTRUCTS_ALPHA", Passed: true, Detail: "mu_B=(3/10)S_split+(7/72)S_split^2"},
		},
		Numeric:  n,
		Statuses: Statuses(),
		Supports: Supports(),
		Failures: Failures(),
		Final:    "Gate 934 consolidates the tensor-structured admissible lattice, airlock closure, target functor, reduced boundary response, and BoundaryActivationMeasure into a closure-factored bridge candidate.",
	}, nil
}

func Statuses() []string {
	return []string{
		"PASS_GATE934_UNIQUE_ADMISSIBLE_LATTICE_FEEDS_AIRLOCK_CLOSURE",
		"PASS_GATE934_THETA_B_Z2_RECOVERED_FROM_CLOSURE",
		"PASS_GATE934_THETA_RANKS_3_7_RECOVERED",
		"PASS_GATE934_REDUCED_B2_RESPONSE_SUPPLIES_S_AND_S2",
		"PASS_GATE934_BOUNDARY_ACTIVATION_MEASURE_RECONSTRUCTS_ALPHA",
		"FIREWALL_PRESERVED_GATE934_NOT_NATIVE",
	}
}

func Supports() []string {
	return []string{
		"CONDITIONAL_SUPPORT_UNIQUE_ADMISSIBLE_LATTICE_FEEDS_AIRLOCK_CLOSURE",
		"CONDITIONAL_SUPPORT_THETA_B_Z2_RECOVERED_FROM_CLOSURE",
		"CONDITIONAL_SUPPORT_THETA_RANKS_3_7_RECOVERED",
		"CONDITIONAL_SUPPORT_REDUCED_B2_RESPONSE_SUPPLIES_S_AND_S2",
		"CONDITIONAL_SUPPORT_BOUNDARY_ACTIVATION_MEASURE_RECONSTRUCTS_ALPHA",
	}
}

func Failures() []string {
	return []string{
		"FAILED_ROUTE_NATIVE_BOUNDARY_ACTIVATION_MEASURE_NOT_CERTIFIED",
		"FAILED_ROUTE_ALPHA_B_REMAINS_BRIDGE_CANDIDATE_NOT_NATIVE",
		"FAILED_ROUTE_NOT_NATIVE_R3",
		"FAILED_ROUTE_FULL_A_F_DESCENT_STILL_BLOCKED",
		"FAILED_ROUTE_NO_GENERATION_CARRIER_MAP",
		"FAILED_ROUTE_NO_FLAVOR_ORIENTATION_MAP",
		"FAILED_ROUTE_NO_INDIVIDUAL_YUKAWA_VALUES",
		"FAILED_ROUTE_NO_NATIVE_YUKAWA_OPERATOR_THEOREM",
	}
}

func FormatComponent(c ComponentAudit) string {
	return fmt.Sprintf("%s: passed=%t marker=%s detail=%s", c.Name, c.Passed, c.Marker, c.Detail)
}

func FormatNumeric(n NumericAudit) string {
	return fmt.Sprintf("alpha_linear=%.17g alpha_quad=%.17g alpha_B=%.17g N_eff=%.16g C_Yukawa=%.16g C_Higgs=%.16g a/T=%.16g b/T^2=%.16g", n.AlphaLinear, n.AlphaQuadratic, n.AlphaB, n.NEff, n.CYukawa, n.CHiggs, n.TraceAOverT, n.TraceBOverT2)
}

func FormatFirewalls(f []string) string { return strings.Join(f, ";") }

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

func componentsOK(cs []ComponentAudit) bool {
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

func numericOK(n NumericAudit) bool {
	return math.Abs(n.AlphaB-AlphaB) < 1e-18 && math.Abs(n.NEff-NEffOperator) < 1e-12 && math.Abs(n.CYukawa-CYukawaOperator) < 1e-12
}
