// Package generation2d4trialitycarrierpackageandcl17realformaudit implements
// Gate 800: D4 Triality Carrier Package Requirement and Cl(1,7) Real-Form Audit.
//
// Gate 800 audits the lawful real-form preconditions for using D4/Spin(8)
// triality inside the current ASHA Cl(1,7) board. It deliberately stops before
// any Yukawa, generation, PMNS/CKM, N_eff, Georgi-Jarlskog, Higgs, or pole-mass
// readout.
package generation2d4trialitycarrierpackageandcl17realformaudit

import (
	"fmt"
	"math"
	"strings"
)

const (
	AuditID = "GATE800-D4-TRIALITY-CARRIER-PACKAGE-CL17-REAL-FORM-AUDIT"

	StatusGate799Inherited              = "PASS_GATE799_NATIVE_THREE_SOURCE_RANKING_INHERITED"
	StatusD4Selected                    = "PASS_D4_TRIALITY_SELECTED_AS_NEXT_NATIVE_BRANCH"
	StatusCoreQuestionDefined           = "PASS_CORE_REAL_FORM_TRIALITY_QUESTION_DEFINED"
	StatusCL17ObjectsRequired           = "PASS_CL17_BOARD_OBJECTS_REQUIRED"
	StatusComplexD4Recorded             = "PASS_COMPLEX_D4_TRIALITY_CANDIDATE_RECORDED"
	StatusRealFormPreservationDefined   = "PASS_REAL_FORM_PRESERVATION_TEST_DEFINED"
	StatusCarrierSignatureDefined       = "PASS_CARRIER_SIGNATURE_AUDIT_DEFINED"
	StatusTrilinearRequirementDefined   = "PASS_CLIFFORD_TRILINEAR_INVARIANT_REQUIREMENT_DEFINED"
	StatusS3ActionTestDefined           = "PASS_S3_TRIALITY_ACTION_TEST_DEFINED"
	StatusOutcomeClassificationDefined  = "PASS_REAL_FORM_OUTCOME_CLASSIFICATION_DEFINED"
	StatusExistingASHAChecked           = "PASS_EXISTING_ASHA_OBJECTS_CHECKED_FOR_TRIALITY_CARRIER_ROLE"
	StatusTrialityYukawaFirewallDefined = "PASS_TRIALITY_TO_YUKAWA_READOUT_FIREWALL_DEFINED"
	StatusLanesSeparated                = "PASS_GJ_SU3_A2_AND_D4_LANES_SEPARATED"
	StatusSuccessCriteriaDefined        = "PASS_SUCCESS_CRITERIA_DEFINED"
	StatusBranchDecisionRecorded        = "PASS_BRANCH_DECISION_RECORDED"
	StatusPhysicalFirewallsEnforced     = "PASS_PHYSICAL_FIREWALLS_ENFORCED"

	StatusComplexD4Shape       = "CONDITIONAL_SUPPORT_COMPLEXIFIED_CL17_HAS_D4_OUTER_AUTOMORPHISM_SHAPE"
	StatusTrilinearPreYukawa   = "CONDITIONAL_SUPPORT_TRILINEAR_PAIRING_IS_REQUIRED_PRE_YUKAWA_OBJECT"
	StatusOutcomeARequired     = "CONDITIONAL_SUPPORT_OUTCOME_A_REQUIRED_BEFORE_NATIVE_D4_BRANCH_CAN_ADVANCE"
	StatusComplexOnlyTriality  = "CONDITIONAL_SUPPORT_COMPLEX_D4_TRIALITY_ONLY"
	StatusTrialityNeedsAirlock = "CONDITIONAL_SUPPORT_TRIALITY_REQUIRES_REAL_FORM_AIRLOCK"

	StatusComplexNotNative          = "FAILED_ROUTE_COMPLEX_D4_TRIALITY_NOT_AUTOMATICALLY_NATIVE_IN_CL17"
	StatusNoCarrierWithoutDims      = "FAILED_ROUTE_NO_TRIALITY_CARRIER_WITHOUT_REAL_MODULE_DIMENSION_AUDIT"
	StatusComplexOuterNotReal       = "FAILED_ROUTE_COMPLEX_OUTER_AUTOMORPHISM_NOT_YET_REAL_FORM_THEOREM"
	StatusNoNativeUnlessPreserves   = "FAILED_ROUTE_NO_NATIVE_TRIALITY_UNLESS_OUTER_AUTOMORPHISM_PRESERVES_CL17_REAL_FORM"
	StatusDim8NotEnough             = "FAILED_ROUTE_DIMENSION_EIGHT_ALONE_DOES_NOT_PROVE_REAL_TRIALITY"
	StatusSignatureMismatch         = "FAILED_ROUTE_SIGNATURE_MISMATCH_BLOCKS_NATIVE_CARRIER_PERMUTATION"
	StatusTrilinearNotYukawa        = "FAILED_ROUTE_TRILINEAR_INVARIANT_NOT_YET_YUKAWA_TRACE_LEDGER"
	StatusNoS3AuditNoTriality       = "FAILED_ROUTE_NO_NATIVE_D4_TRIALITY_WITHOUT_FULL_S3_RELATION_AUDIT"
	StatusAirlockRequired           = "FAILED_ROUTE_COMPLEX_OR_ALTERNATE_REAL_FORM_TRIALITY_REQUIRES_AIRLOCK"
	StatusK7Not8                    = "FAILED_ROUTE_K7_NOT_EIGHT_DIMENSIONAL_TRIALITY_FRAME"
	StatusK7HodgeNotCarrier         = "FAILED_ROUTE_K7_HODGE_43_NOT_D4_TRIALITY_CARRIER"
	StatusLambda4NotModule          = "FAILED_ROUTE_LAMBDA4_CHAMBER_NOT_AUTOMATICALLY_TRIALITY_MODULE"
	StatusAggregateTracesNotCarrier = "FAILED_ROUTE_AGGREGATE_YUKAWA_TRACES_DO_NOT_DEFINE_TRIALITY_CARRIER"
	StatusTrialityCarrierNotYukawa  = "FAILED_ROUTE_D4_TRIALITY_CARRIER_NOT_SUFFICIENT_FOR_YUKAWA_TRACE_THEOREM"
	StatusNoTraceAtomReadout        = "FAILED_ROUTE_NO_TRACE_ATOM_READOUT_FROM_TRIALITY_YET"
	StatusNoTopBreakingOperator     = "FAILED_ROUTE_NO_TOP_DOMINANCE_BREAKING_OPERATOR_FROM_TRIALITY_YET"
	StatusNoPMNSCKMReadout          = "FAILED_ROUTE_NO_PMNS_CKM_READOUT_FROM_TRIALITY_YET"
	StatusGJNotD4                   = "FAILED_ROUTE_GJ_THREE_NOT_D4_TRIALITY_THEOREM"
	StatusA2NotD4                   = "FAILED_ROUTE_A2_HEXAGON_NOT_D4_TRIALITY_THEOREM"
	StatusMotifNotEvidence          = "FAILED_ROUTE_SYMBOLIC_MOTIF_NOT_TYPED_EVIDENCE"
	StatusGenericSpin8NotEnough     = "FAILED_ROUTE_GENERIC_SPIN8_TRIALITY_STATEMENT_NOT_ENOUGH_FOR_ASHA"
	StatusNoFullNativeCL17D4        = "FAILED_ROUTE_NO_FULL_NATIVE_CL17_D4_TRIALITY_CARRIER"
	StatusFirewallPreservedGate800  = "FIREWALL_PRESERVED_GATE800_D4_TRIALITY_CL17_REAL_FORM_BOUNDARY"
)

const (
	P = 1
	Q = 7
	N = P + Q
)

type Inheritance struct {
	Gate799Inherited      bool
	D4Selected            bool
	NEff                  float64
	CYukawa               float64
	CurrentCertifiedThree string
	NotGeneration         bool
	NotD4                 bool
	Verdicts              []string
}

type CL17Board struct {
	P                            int
	Q                            int
	N                            int
	DifferenceMod8               int
	FullAlgebra                  string
	EvenAlgebra                  string
	VectorDimR                   int
	SpinAlgebra                  string
	ComplexifiedSpin             string
	VolumeSquare                 int
	RealChiralityProjectorsExist bool
	RealChiralSpinorDim          string
	ComplexHalfSpinorDimC        int
	ComplexHalfSpinorDimR        int
	MinimalRealSpinorDim         int
	Verdict                      string
}

type ComplexD4Candidate struct {
	Recorded               bool
	Complexification       string
	DynkinType             string
	OuterAutomorphism      string
	LawfulComplexCandidate bool
	RealNative             bool
	Verdict                string
	Support                string
	Failure                string
}

type RealFormTest struct {
	Defined           bool
	RealStructure     string
	TestEquation      string
	PreservesRealForm string
	Reason            string
	Outcome           string
	Verdict           string
	Failures          []string
}

type CarrierSignatureAudit struct {
	Defined                    bool
	VectorSignature            string
	SPlusSignature             string
	SMinusSignature            string
	CarrierDimensionCompatible bool
	SignatureCompatible        bool
	Verdict                    string
	Failures                   []string
}

type TrilinearAudit struct {
	Defined              bool
	Formula              string
	RealStatus           string
	ComplexStatus        string
	RequiredBeforeYukawa bool
	NotYukawa            bool
	Verdict              string
	Support              string
	Failure              string
}

type S3TrialityActionTest struct {
	Defined              bool
	ThreeCycle           string
	Swap                 string
	Relations            []string
	RealClosureCertified bool
	Reason               string
	Verdict              string
	Failure              string
}

type OutcomeClassification struct {
	Defined          bool
	Selected         string
	Options          []string
	OutcomeARequired bool
	FullNativeFound  bool
	Verdict          string
	Supports         []string
	Failures         []string
}

type ExistingObjectCheck struct {
	Object    string
	Dimension string
	Audited   bool
	Verdict   string
	Failure   string
}

type TrialityYukawaFirewall struct {
	Defined          bool
	MissingObjects   []string
	NotYukawaTheorem bool
	Verdict          string
	Failures         []string
}

type LaneSeparation struct {
	Recorded       bool
	GeorgiJarlskog string
	SU3A2          string
	D4Triality     string
	Blocked        []string
	Verdict        string
	Failures       []string
}

type SuccessCriteria struct {
	Defined              bool
	AllowedResults       []string
	GenericSpin8Rejected bool
	Verdict              string
	Failure              string
}

type BranchDecision struct {
	Recorded          bool
	SelectedOutcome   string
	Next              string
	AlternativeIfFull string
	AlternativeIfFail string
	Verdict           string
}

type Firewalls struct {
	Enforced        bool
	NoYukawa        bool
	NoPMNSCKM       bool
	NoNEff          bool
	NoGJ            bool
	NoScalarRuntime bool
	NoPoleMass      bool
	NoHistoryLoop   bool
	Verdict         string
}

type Analysis struct {
	Inheritance    Inheritance
	Board          CL17Board
	ComplexD4      ComplexD4Candidate
	RealForm       RealFormTest
	Signatures     CarrierSignatureAudit
	Trilinear      TrilinearAudit
	S3Test         S3TrialityActionTest
	Outcome        OutcomeClassification
	Existing       []ExistingObjectCheck
	YukawaFirewall TrialityYukawaFirewall
	Lanes          LaneSeparation
	Criteria       SuccessCriteria
	Branch         BranchDecision
	Firewalls      Firewalls
	Truth          string
	Final          string
}

func mod8(x int) int {
	r := x % 8
	if r < 0 {
		r += 8
	}
	return r
}

func volumeSquare(p, q int) int {
	n := p + q
	exponent := q + n*(n-1)/2
	if exponent%2 == 0 {
		return 1
	}
	return -1
}

func BuildDefault() (Analysis, error) {
	dmod := mod8(P - Q)
	omega2 := volumeSquare(P, Q)
	if N != 8 || dmod != 2 || omega2 != -1 {
		return Analysis{}, fmt.Errorf("unexpected Cl(1,7) arithmetic: n=%d mod=%d omega2=%d", N, dmod, omega2)
	}
	board := CL17Board{
		P: P, Q: Q, N: N, DifferenceMod8: dmod,
		FullAlgebra: "Cl(1,7) ≅ Mat(16,R) under the active Clifford-sign convention p-q≡2 mod 8",
		EvenAlgebra: "Cl^0(1,7) carries a commuting complex/chirality structure; real chiral idempotents are blocked because the real volume element squares to -1",
		VectorDimR:  8,
		SpinAlgebra: "spin(1,7)", ComplexifiedSpin: "spin(1,7)_C ≅ so(8,C)",
		VolumeSquare:                 omega2,
		RealChiralityProjectorsExist: false,
		RealChiralSpinorDim:          "not certified; chiral half-spinors are complex after complexification, not real 8-dimensional Cl(1,7) modules",
		ComplexHalfSpinorDimC:        8, ComplexHalfSpinorDimR: 16,
		MinimalRealSpinorDim: 16,
		Verdict:              StatusCL17ObjectsRequired,
	}
	existing := []ExistingObjectCheck{
		{Object: "K7", Dimension: "7", Audited: true, Verdict: "K7 is a contact/intersection carrier, not an eight-dimensional triality frame", Failure: StatusK7Not8},
		{Object: "K7+ and K7-", Dimension: "4|3", Audited: true, Verdict: "Hodge polarity is native but not a D4 triality carrier", Failure: StatusK7HodgeNotCarrier},
		{Object: "Lambda^4 R8", Dimension: "70", Audited: true, Verdict: "Λ⁴R⁸ is a chamber; no V,S+,S- triality module action is certified", Failure: StatusLambda4NotModule},
		{Object: "P_B and P_G", Dimension: "rank 56 and rank 14", Audited: true, Verdict: "projectors select Boolean/octonionic supports, not triality automorphisms", Failure: "FAILED_ROUTE_PB_PG_DO_NOT_DEFINE_TRIALITY_AUTOMORPHISMS"},
		{Object: "H72", Dimension: "72", Audited: true, Verdict: "augmented chamber supports boundary-response bookkeeping, not V,S+,S- carrier permutation", Failure: "FAILED_ROUTE_H72_NOT_TRIALITY_MODULE"},
		{Object: "Higgs socket K7+", Dimension: "4 real", Audited: true, Verdict: "Higgs socket is four-real-dimensional before complex structure, not an eight-dimensional triality frame", Failure: "FAILED_ROUTE_HIGGS_SOCKET_NOT_EIGHT_DIMENSIONAL_TRIALITY_FRAME"},
		{Object: "aggregate Yukawa traces", Dimension: "two scalars a,b", Audited: true, Verdict: "aggregate a,b carry no triality frame data", Failure: StatusAggregateTracesNotCarrier},
	}
	outcome := OutcomeClassification{
		Defined:          true,
		Selected:         "Outcome C — complex-only triality for the current Cl(1,7) audit; full real S3 carrier not certified on the native real board",
		Options:          []string{"Outcome A — full native real triality", "Outcome B — partial real-form automorphism", "Outcome C — complex-only triality", "Outcome D — alternate-real-form triality requiring RealFormAirlock"},
		OutcomeARequired: true,
		FullNativeFound:  false,
		Verdict:          StatusOutcomeClassificationDefined,
		Supports:         []string{StatusComplexOnlyTriality, StatusTrialityNeedsAirlock, StatusOutcomeARequired},
		Failures:         []string{StatusNoFullNativeCL17D4, StatusAirlockRequired},
	}
	a := Analysis{
		Inheritance:    Inheritance{Gate799Inherited: true, D4Selected: true, NEff: 3.0023273474722147, CYukawa: 0.9992248188812008, CurrentCertifiedThree: "color-tripled top dominance", NotGeneration: true, NotD4: true, Verdicts: []string{StatusGate799Inherited, StatusD4Selected}},
		Board:          board,
		ComplexD4:      ComplexD4Candidate{Recorded: true, Complexification: "spin(1,7)_C ≅ so(8,C)", DynkinType: "D4", OuterAutomorphism: "Out(D4)≅S3 over C", LawfulComplexCandidate: true, RealNative: false, Verdict: StatusComplexD4Recorded, Support: StatusComplexD4Shape, Failure: StatusComplexOuterNotReal},
		RealForm:       RealFormTest{Defined: true, RealStructure: "sigma_{1,7} selecting spin(1,7) inside so(8,C)", TestEquation: "tau sigma_{1,7} = sigma_{1,7} tau", PreservesRealForm: "not certified for full S3", Reason: "real volume element has square -1, so real chiral 8-dimensional S± carriers are not certified; complex triality does not automatically commute with the Cl(1,7) real structure", Outcome: "complex-only / real-form-airlock required", Verdict: StatusRealFormPreservationDefined, Failures: []string{StatusNoNativeUnlessPreserves, StatusComplexNotNative}},
		Signatures:     CarrierSignatureAudit{Defined: true, VectorSignature: "(1,7)", SPlusSignature: "no native real S+ eight-dimensional bilinear signature certified; complex half-spinor is 8 complex = 16 real", SMinusSignature: "no native real S- eight-dimensional bilinear signature certified; complex half-spinor is 8 complex = 16 real", CarrierDimensionCompatible: false, SignatureCompatible: false, Verdict: StatusCarrierSignatureDefined, Failures: []string{StatusDim8NotEnough, StatusSignatureMismatch, StatusNoCarrierWithoutDims}},
		Trilinear:      TrilinearAudit{Defined: true, Formula: "T(v,ψ+,ψ-)=<γ(v)ψ+,ψ->", RealStatus: "not a native real Cl(1,7) V×S+×S-→R object until real S± carriers are certified", ComplexStatus: "complexified pre-Yukawa invariant candidate V_C×S+_C×S-_C→C", RequiredBeforeYukawa: true, NotYukawa: true, Verdict: StatusTrilinearRequirementDefined, Support: StatusTrilinearPreYukawa, Failure: StatusTrilinearNotYukawa},
		S3Test:         S3TrialityActionTest{Defined: true, ThreeCycle: "V→S+→S-→V", Swap: "swap of two outer carriers", Relations: []string{"tau_3cycle^3=1", "tau_swap^2=1", "tau_swap tau_3cycle tau_swap=tau_3cycle^{-1}"}, RealClosureCertified: false, Reason: "relations are available as complex D4 test requirements, but no native real Cl(1,7) carrier action is certified", Verdict: StatusS3ActionTestDefined, Failure: StatusNoS3AuditNoTriality},
		Outcome:        outcome,
		Existing:       existing,
		YukawaFirewall: TrialityYukawaFirewall{Defined: true, MissingObjects: []string{"YukawaSectorAssignment", "TraceAtomReadout", "TopDominanceBreakingOperator", "RestPressureOperator", "GenerationMixingReadout"}, NotYukawaTheorem: true, Verdict: StatusTrialityYukawaFirewallDefined, Failures: []string{StatusTrialityCarrierNotYukawa, StatusNoTraceAtomReadout, StatusNoTopBreakingOperator, StatusNoPMNSCKMReadout}},
		Lanes:          LaneSeparation{Recorded: true, GeorgiJarlskog: "high-scale Clebsch-three diagnostic requiring multi-scale Yukawa ledger", SU3A2: "A2/SU(3) weight-geometry search motif, not D4 triality", D4Triality: "outer automorphism of D4-type Lie algebra requiring real-form-compatible carrier and S3 action", Blocked: []string{"GJ three = D4 triality", "A2 hexagon = D4 triality", "visual motif = theorem", "color SU(3)=flavor SU(3)"}, Verdict: StatusLanesSeparated, Failures: []string{StatusGJNotD4, StatusA2NotD4, StatusMotifNotEvidence}},
		Criteria:       SuccessCriteria{Defined: true, AllowedResults: []string{"PASS_FULL_NATIVE_CL17_D4_TRIALITY_CARRIER_FOUND", StatusNoFullNativeCL17D4, StatusComplexOnlyTriality, "CONDITIONAL_SUPPORT_PARTIAL_REAL_FORM_AUTOMORPHISM_ONLY", StatusTrialityNeedsAirlock}, GenericSpin8Rejected: true, Verdict: StatusSuccessCriteriaDefined, Failure: StatusGenericSpin8NotEnough},
		Branch:         BranchDecision{Recorded: true, SelectedOutcome: "Outcome C — complex-only / real-form-airlock required", Next: "Gate 801 — Real-Form Triality Airlock and Native-Status Firewall Audit", AlternativeIfFull: "Gate 801 — D4 Triality Trilinear Invariant and Yukawa Trace Readout Obstruction Audit", AlternativeIfFail: "Gate 801 — Alternative Native Three-Source Search: SU3/A2 and GenerationCarrier Audit", Verdict: StatusBranchDecisionRecorded},
		Firewalls:      Firewalls{Enforced: true, NoYukawa: true, NoPMNSCKM: true, NoNEff: true, NoGJ: true, NoScalarRuntime: true, NoPoleMass: true, NoHistoryLoop: true, Verdict: StatusFirewallPreservedGate800},
		Truth:          "Gate 800 finds complex D4 triality after complexification, but does not certify full native real Cl(1,7) S3 triality because the real chiral 8-dimensional carriers are not available on the active board.",
		Final:          "Gate 800 does not use triality to explain Yukawa structure. It audits the real form first: Cl(1,7) has a complexified D4 candidate, but the native real board does not certify a real S3 package permuting V,S+,S-. The result is Outcome C: complex-only triality / RealFormAirlock required before any Yukawa trace-readout audit.",
	}
	return a, nil
}

func Statuses() []string {
	return []string{
		StatusGate799Inherited, StatusD4Selected, StatusCoreQuestionDefined, StatusCL17ObjectsRequired,
		StatusComplexD4Recorded, StatusRealFormPreservationDefined, StatusCarrierSignatureDefined,
		StatusTrilinearRequirementDefined, StatusS3ActionTestDefined, StatusOutcomeClassificationDefined,
		StatusExistingASHAChecked, StatusTrialityYukawaFirewallDefined, StatusLanesSeparated,
		StatusSuccessCriteriaDefined, StatusBranchDecisionRecorded, StatusPhysicalFirewallsEnforced,
		StatusComplexD4Shape, StatusTrilinearPreYukawa, StatusOutcomeARequired, StatusComplexOnlyTriality, StatusTrialityNeedsAirlock,
		StatusComplexNotNative, StatusNoCarrierWithoutDims, StatusComplexOuterNotReal, StatusNoNativeUnlessPreserves,
		StatusDim8NotEnough, StatusSignatureMismatch, StatusTrilinearNotYukawa, StatusNoS3AuditNoTriality,
		StatusAirlockRequired, StatusK7Not8, StatusK7HodgeNotCarrier, StatusLambda4NotModule,
		StatusAggregateTracesNotCarrier, StatusTrialityCarrierNotYukawa, StatusNoTraceAtomReadout,
		StatusNoTopBreakingOperator, StatusNoPMNSCKMReadout, StatusGJNotD4, StatusA2NotD4,
		StatusMotifNotEvidence, StatusGenericSpin8NotEnough, StatusNoFullNativeCL17D4,
		StatusFirewallPreservedGate800,
	}
}

func FormatBoard(b CL17Board) string {
	return fmt.Sprintf("Cl(%d,%d): n=%d, p-q mod8=%d, omega^2=%d, full=%s, even=%s, dim_R(V)=%d, real_chiral_projectors=%v, S±=%s, complex_half_spinor_dim_C=%d", b.P, b.Q, b.N, b.DifferenceMod8, b.VolumeSquare, b.FullAlgebra, b.EvenAlgebra, b.VectorDimR, b.RealChiralityProjectorsExist, b.RealChiralSpinorDim, b.ComplexHalfSpinorDimC)
}
func FormatRealForm(r RealFormTest) string {
	return fmt.Sprintf("%s; test=%s; preserves=%s; outcome=%s", r.RealStructure, r.TestEquation, r.PreservesRealForm, r.Outcome)
}
func FormatSignatures(s CarrierSignatureAudit) string {
	return fmt.Sprintf("signature(V)=%s; signature(S+)=%s; signature(S-)=%s; dimension_compatible=%v; signature_compatible=%v", s.VectorSignature, s.SPlusSignature, s.SMinusSignature, s.CarrierDimensionCompatible, s.SignatureCompatible)
}
func FormatExisting(xs []ExistingObjectCheck) string {
	parts := []string{}
	for _, x := range xs {
		parts = append(parts, fmt.Sprintf("%s[%s]: %s", x.Object, x.Dimension, x.Failure))
	}
	return strings.Join(parts, "; ")
}
func containsAll(hay []string, needles []string) bool {
	joined := strings.Join(hay, "\n")
	for _, n := range needles {
		if !strings.Contains(joined, n) {
			return false
		}
	}
	return true
}
func closeAbs(a, b, tol float64) bool { return math.Abs(a-b) <= tol }
