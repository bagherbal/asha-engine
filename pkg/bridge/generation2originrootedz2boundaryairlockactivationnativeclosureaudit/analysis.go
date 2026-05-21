// Package generation2originrootedz2boundaryairlockactivationnativeclosureaudit implements
// Gate 939: OriginRooted Z2 BoundaryAirlock Activation Native Closure Audit.
//
// Gate 939 follows the Gate 938A explicit native-promotion blocker ledger. It
// attempts an ambitious consolidation: the four native-R3 blockers are treated
// as certificate clauses of one origin-rooted Z2 BoundaryAirlock activation
// functor. The audit is deliberately conservative: it collapses the gaps into a
// single certificate chain, but it does not grant native/post-orientation R3
// unless all four native certificates are actually present.
package generation2originrootedz2boundaryairlockactivationnativeclosureaudit

import (
	"fmt"
	"strings"
)

const (
	AuditID         = "GATE939-GENERATION2ORIGINROOTEDZ2BOUNDARYAIRLOCKACTIVATIONNATIVECLOSUREAUDIT"
	InheritedStatus = "R3_PRETEST_PASSED_NATIVE_R3_BLOCKERS_EXPLICIT"
	Classification  = "R3_ORIGIN_ROOTED_NATIVE_PROMOTION_CANDIDATE_GAPS_COLLAPSED_NOT_FULLY_CERTIFIED"
	ShortStatus     = "R3_NATIVE_GAPS_COLLAPSED_TO_ORIGIN_ROOTED_AIRLOCK_FUNCTOR"
	FinalTruth      = "R3_ORIGIN_ROOTED_NATIVE_CLOSURE_ATTEMPT_FOUR_GAPS_COLLAPSED_NATIVE_CERTIFICATES_PENDING"
	FullPassTruth   = "R3_NATIVE_POST_ORIENTATION_Z2_TRACE_LEDGER_CERTIFIED"
	PartialTruth    = "R3_ORIGIN_ROOTED_TRACEBRIDGE_NATIVE_PROMOTION_PARTIAL_SUPPORT"
	NextGate        = "NEXT_PRESSURE_GATE940_ORIGIN_ROOTED_AIRLOCK_FUNCTOR_CERTIFICATE_SOURCE_PRIORITY_AUDIT"

	MasterFunctorName = "OriginRootedZ2BoundaryAirlockActivationFunctor"
	MasterFunctorID   = "A_R3^Z2"
	MasterFunctorFlow = "Cl(1,7)->Lambda4V8->B2->[p]Z2->A_airlockZ2->Cl_airlockZ2->Theta_BZ2->mu_B->alpha_B->YdaggerY->N_eff_operator"
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

type Diagnostics struct {
	Ssplit          float64
	AlphaLinear     float64
	AlphaQuadratic  float64
	AlphaB          float64
	NEffOperator    float64
	CYukawaOperator float64
	CHiggsOperator  float64
}

type NativeClause struct {
	Name             string
	RequiredTheorem  string
	CandidateReading string
	PassCondition    []string
	Failure          []string
	NativeCertified  bool
	ClauseCollapsed  bool
	BlocksFullNative bool
}

type RetiredFalseRoute struct {
	Name   string
	Reason string
}

type R4BoundaryItem struct {
	Name    string
	Failure string
}

type Analysis struct {
	AuditID          string
	Inherited        string
	MasterFunctor    string
	MasterFunctorID  string
	Classification   string
	ShortStatus      string
	Truth            string
	FullPassEligible bool
	Diagnostics      Diagnostics
	Clauses          []NativeClause
	RetiredRoutes    []RetiredFalseRoute
	R4Boundary       []R4BoundaryItem
	Statuses         []string
	Supports         []string
	Failures         []string
	Final            string
}

func BuildDefault() (Analysis, error) {
	clauses := NativeClauses()
	if len(clauses) != 4 {
		return Analysis{}, fmt.Errorf("Gate 939 must track exactly four native-promotion clauses")
	}
	if !allCollapsed(clauses) {
		return Analysis{}, fmt.Errorf("Gate 939 must collapse all native-promotion clauses into the origin-rooted functor")
	}
	fullPass := allNativeCertified(clauses)
	truth := FinalTruth
	classification := Classification
	short := ShortStatus
	if fullPass {
		truth = FullPassTruth
		classification = "R3_NATIVE_POST_ORIENTATION_Z2_TRACE_LEDGER_CERTIFIED"
		short = "R3_NATIVE_POST_ORIENTATION_TRACE_LEDGER_CERTIFIED"
	}
	return Analysis{
		AuditID:          AuditID,
		Inherited:        InheritedStatus,
		MasterFunctor:    MasterFunctorName,
		MasterFunctorID:  MasterFunctorID,
		Classification:   classification,
		ShortStatus:      short,
		Truth:            truth,
		FullPassEligible: fullPass,
		Diagnostics:      DefaultDiagnostics(),
		Clauses:          clauses,
		RetiredRoutes:    RetiredFalseRoutes(),
		R4Boundary:       R4Boundary(),
		Statuses:         Statuses(),
		Supports:         Supports(),
		Failures:         Failures(),
		Final:            "Gate 939 collapses the four native-R3 promotion gaps into one origin-rooted Z2 BoundaryAirlock activation functor certificate chain. It does not grant native/post-orientation R3 because the four native certificates remain pending.",
	}, nil
}

func DefaultDiagnostics() Diagnostics {
	return Diagnostics{Ssplit: Ssplit, AlphaLinear: AlphaLinear, AlphaQuadratic: AlphaQuadratic, AlphaB: AlphaB, NEffOperator: NEffOperator, CYukawaOperator: CYukawaOperator, CHiggsOperator: CHiggsOperator}
}

func NativeClauses() []NativeClause {
	return []NativeClause{
		{
			Name:             "Native admissible airlock support lattice from finite projector support",
			RequiredTheorem:  "NativeAdmissibleAirlockSupportLatticeTheorem",
			CandidateReading: "admissible supports are projector-generated finite submodules built from e_phase, e_opposite, P_1, P_3, W=P_1+P_3, and C_R^2=e_phase+e_opposite, not arbitrary rank-compatible subspaces",
			PassCondition: []string{
				"PASS_NATIVE_SUPPORTS_ARE_PROJECTOR_GENERATED_SUBMODULES",
				"PASS_TENSOR_STRUCTURED_COMPLETIONS_ARE_ONLY_ADMISSIBLE_AIRLOCK_SUPPORTS",
				"PASS_AIRLOCK_SUPPORT_LATTICE_IS_NATIVE_PROJECTOR_SUPPORT_LATTICE",
			},
			Failure: []string{
				"FAILED_ROUTE_NO_NATIVE_ADMISSIBLE_AIRLOCK_SUPPORT_LATTICE_THEOREM",
				"FAILED_ROUTE_TENSOR_STRUCTURED_ADMISSIBILITY_REMAINS_BRIDGE_RULE",
			},
			NativeCertified:  false,
			ClauseCollapsed:  true,
			BlocksFullNative: true,
		},
		{
			Name:             "Native S_split response-parameter source",
			RequiredTheorem:  "NativeSsplitBoundaryResponseParameterTheorem",
			CandidateReading: "S_split is the unique boundary-pair response scalar and is inserted uniformly into every boundary factor by Z2/boundary-pair equivariance",
			PassCondition: []string{
				"PASS_S_SPLIT_IS_NATIVE_BOUNDARY_SPLIT_SCALAR",
				"PASS_S_SPLIT_TRANSPORTS_TO_REDUCED_B2_RESPONSE_PARAMETER",
				"PASS_BOUNDARY_PAIR_EQUIVARIANCE_FORCES_UNIFORM_INSERTION",
				"PASS_S2_TERM_ARISES_FROM_EXTERIOR_MULTIPLICATION_NOT_SECOND_TRANSPORT",
			},
			Failure: []string{
				"FAILED_ROUTE_NO_NATIVE_S_SPLIT_RESPONSE_PARAMETER_THEOREM",
				"FAILED_ROUTE_NO_NATIVE_TRANSPORT_OF_S_SPLIT_TO_BOUNDARY_PAIR_RESPONSE",
				"FAILED_ROUTE_S_SPLIT_REMAINS_TYPED_BRIDGE_SCALAR_NOT_NATIVE_INPUT",
			},
			NativeCertified:  false,
			ClauseCollapsed:  true,
			BlocksFullNative: true,
		},
		{
			Name:             "Native BoundaryActivationMeasure as finite normalized trace response",
			RequiredTheorem:  "NativeBoundaryActivationMeasureTheorem",
			CandidateReading: "mu_B is read as a finite normalized trace response Tr_Hk(P_Theta(k))/Tr_Hk(I) times S_split^k on boundary-augmented response chambers H_10 and H_72",
			PassCondition: []string{
				"PASS_BOUNDARY_ACTIVATION_MEASURE_IS_FINITE_NORMALIZED_TRACE_RESPONSE",
				"PASS_H10_AND_H72_ARE_NATIVE_RESPONSE_CHAMBERS",
				"PASS_DEGREE_LANE_NORMALIZATION_BY_TRACE_STATE_IS_NATIVE",
				"PASS_MU_B_IS_NOT_ARBITRARY_BRIDGE_MEASURE",
			},
			Failure: []string{
				"FAILED_ROUTE_NO_NATIVE_BOUNDARY_ACTIVATION_MEASURE_THEOREM",
				"FAILED_ROUTE_H10_H72_NORMALIZATION_REMAINS_BRIDGE_RESPONSE_CHAMBER_RULE",
				"FAILED_ROUTE_MU_B_REMAINS_BRIDGE_MEASURE",
			},
			NativeCertified:  false,
			ClauseCollapsed:  true,
			BlocksFullNative: true,
		},
		{
			Name:             "Full A_F descent or lawful spontaneous-orientation status",
			RequiredTheorem:  "NativeFullAFDescentTheorem or LawfulSpontaneousOrientationTheorem",
			CandidateReading: "A_F -> finite one-form/Higgs-edge orientation -> A_F^orient -> Z2 airlock trace ledger may certify native post-orientation R3 without claiming native unbroken R3",
			PassCondition: []string{
				"PASS_A_F_ORIENT_IS_LAWFUL_STABILIZER_OF_FINITE_ONE_FORM_ORIENTATION",
				"PASS_POST_ORIENTATION_LEDGER_IS_NATIVE_RELATIVE_TO_SPONTANEOUS_ORIENTATION",
				"PASS_FULL_A_F_DESCENT_NOT_REQUIRED_FOR_ORIENTED_TRACE_LEDGER",
				"PASS_HIGGS_ORIENTATION_IS_LAWFUL_NATIVE_BREAKING_NOT_EXTERNAL_SEAL",
			},
			Failure: []string{
				"FAILED_ROUTE_FULL_A_F_DESCENT_STILL_BLOCKED",
				"FAILED_ROUTE_HIGGS_ORIENTATION_REMAINS_SEALED_AS_ORIENTATION_CLASS",
				"FAILED_ROUTE_POST_ORIENTATION_LEDGER_NOT_NATIVE_FINITE_SECTOR_LEDGER",
			},
			NativeCertified:  false,
			ClauseCollapsed:  true,
			BlocksFullNative: true,
		},
	}
}

func RetiredFalseRoutes() []RetiredFalseRoute {
	return []RetiredFalseRoute{
		{Name: "lambda versus barlambda representative", Reason: "absorbed into Z2 class ledger and preserved as a negative-test witness"},
		{Name: "+Q_phi versus -Q_phi", Reason: "global phase sign is no longer a primary native-R3 blocker after Z2 quotienting"},
		{Name: "representative alpha", Reason: "alpha_B is formulated by the class functor and pre-tested as representative-independent"},
		{Name: "cross-lane pollution", Reason: "pre-test rejects (7/72)s and (3/10)s^2 false lanes"},
		{Name: "Theta(2)=F_2/F_1", Reason: "associated-graded top target rejected by closure over fixed puncture base"},
		{Name: "bare denominators", Reason: "8 and 70 rejected in favor of boundary-augmented 10 and 72"},
		{Name: "orphan support fragments", Reason: "not tensor-structured airlock support completions"},
		{Name: "arbitrary rank-compatible subspaces", Reason: "not projector-generated finite support submodules under the candidate reading"},
	}
}

func R4Boundary() []R4BoundaryItem {
	return []R4BoundaryItem{
		{Name: "individual Yukawa values", Failure: "FAILED_ROUTE_NO_INDIVIDUAL_YUKAWA_VALUES"},
		{Name: "physical particle assignment", Failure: "FAILED_ROUTE_NO_PHYSICAL_PARTICLE_ASSIGNMENT"},
		{Name: "generation carrier map", Failure: "FAILED_ROUTE_NO_GENERATION_CARRIER_MAP"},
		{Name: "flavor orientation map", Failure: "FAILED_ROUTE_NO_FLAVOR_ORIENTATION_MAP"},
		{Name: "R4 native Yukawa spectrum theorem", Failure: "FAILED_ROUTE_NO_R4_NATIVE_YUKAWA_SPECTRUM_THEOREM"},
	}
}

func Statuses() []string {
	return []string{
		"PASS_GATE939_R3_PRETEST_PASSED_TRACEBRIDGE_INHERITED",
		"PASS_GATE939_FOUR_NATIVE_GAPS_COLLAPSED_TO_ORIGIN_ROOTED_FUNCTOR",
		"PASS_GATE939_ALL_OR_NOTHING_CERTIFICATE_CHAIN_EXPLICIT",
		"FIREWALL_GATE939_NATIVE_R3_NOT_GRANTED_WITHOUT_ALL_FOUR_CERTIFICATES",
		"FIREWALL_GATE939_R4_FLAVOR_AND_INDIVIDUAL_YUKAWA_EXCLUDED",
	}
}

func Supports() []string {
	return []string{
		"CONDITIONAL_SUPPORT_R3_PRETEST_PASSED_TRACEBRIDGE_INHERITED",
		"CONDITIONAL_SUPPORT_FOUR_NATIVE_R3_GAPS_COLLAPSE_TO_ORIGIN_ROOTED_AIRLOCK_ACTIVATION_FUNCTOR",
		"CONDITIONAL_SUPPORT_ADMISSIBLE_SUPPORTS_CAN_BE_READ_AS_PROJECTOR_GENERATED_FINITE_SUBMODULES",
		"CONDITIONAL_SUPPORT_S_SPLIT_CAN_BE_READ_AS_UNIQUE_BOUNDARY_PAIR_RESPONSE_SCALAR_IF_SOURCE_CERTIFIED",
		"CONDITIONAL_SUPPORT_BOUNDARY_ACTIVATION_MEASURE_CAN_BE_READ_AS_FINITE_NORMALIZED_TRACE_RESPONSE",
		"CONDITIONAL_SUPPORT_A_F_ORIENT_CAN_BE_READ_AS_LAWFUL_FINITE_ONE_FORM_STABILIZER_IF_ORIENTATION_CERTIFIED",
		"CONDITIONAL_SUPPORT_NATIVE_R3_PROMOTION_HAS_EXPLICIT_ALL_OR_NOTHING_CERTIFICATE_CHAIN",
	}
}

func Failures() []string {
	return []string{
		"FAILED_ROUTE_NATIVE_R3_NOT_GRANTED_WITHOUT_ALL_FOUR_CERTIFICATES",
		"FAILED_ROUTE_NO_NATIVE_BOUNDARY_ACTIVATION_MEASURE_THEOREM_UNLESS_TRACE_RESPONSE_CERTIFIED",
		"FAILED_ROUTE_NO_NATIVE_S_SPLIT_RESPONSE_PARAMETER_THEOREM_UNLESS_SOURCE_TRANSPORT_CERTIFIED",
		"FAILED_ROUTE_NO_NATIVE_ADMISSIBLE_AIRLOCK_SUPPORT_LATTICE_THEOREM_UNLESS_PROJECTOR_SUBMODULE_RULE_CERTIFIED",
		"FAILED_ROUTE_FULL_A_F_DESCENT_OR_SPONTANEOUS_ORIENTATION_THEOREM_REQUIRED",
		"FAILED_ROUTE_NO_OFFICIAL_LEDGER_UPDATE_ALLOWED_WITHOUT_NATIVE_R3_CERTIFICATION",
		"FAILED_ROUTE_NO_INDIVIDUAL_YUKAWA_VALUES",
		"FAILED_ROUTE_NO_PHYSICAL_PARTICLE_ASSIGNMENT",
		"FAILED_ROUTE_NO_GENERATION_CARRIER_MAP",
		"FAILED_ROUTE_NO_FLAVOR_ORIENTATION_MAP",
		"FAILED_ROUTE_NO_R4_NATIVE_YUKAWA_SPECTRUM_THEOREM",
	}
}

func allCollapsed(cs []NativeClause) bool {
	for _, c := range cs {
		if !c.ClauseCollapsed || c.RequiredTheorem == "" || len(c.Failure) == 0 || len(c.PassCondition) == 0 {
			return false
		}
	}
	return true
}

func allNativeCertified(cs []NativeClause) bool {
	for _, c := range cs {
		if !c.NativeCertified {
			return false
		}
	}
	return true
}

func allBlockFullNative(cs []NativeClause) bool {
	for _, c := range cs {
		if !c.BlocksFullNative {
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

func clauseFailures(cs []NativeClause) []string {
	var out []string
	for _, c := range cs {
		out = append(out, c.Failure...)
	}
	return out
}

func clausePassConditions(cs []NativeClause) []string {
	var out []string
	for _, c := range cs {
		out = append(out, c.PassCondition...)
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

func FormatClauses(cs []NativeClause) string {
	parts := make([]string, 0, len(cs))
	for _, c := range cs {
		state := "certificate pending"
		if c.NativeCertified {
			state = "native certified"
		}
		parts = append(parts, fmt.Sprintf("%s -> %s (%s)", c.Name, c.RequiredTheorem, state))
	}
	return strings.Join(parts, "; ")
}

func FormatClauseFailures(cs []NativeClause) string { return strings.Join(clauseFailures(cs), "; ") }

func FormatRetired(rs []RetiredFalseRoute) string {
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
