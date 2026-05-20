// Package generation2boundaryfnrestpressurespurionandneffminusthreeclosureaudit implements
// Gate 810: Boundary-FN RestPressure Spurion and N_eff-Minus-Three Closure Audit.
//
// Gate 810 tests the Gate 809 blink N_eff-3 ≈ (9/5)s as a boundary-driven
// Froggatt-Nielsen-style rest-pressure candidate. It computes residuals,
// fourth-root spurions, top/rest positivity, rest-concentration regimes,
// alternative coefficient controls, and preserves all firewalls preventing a
// numerical closure from becoming a native hierarchy theorem.
package generation2boundaryfnrestpressurespurionandneffminusthreeclosureaudit

import (
	"fmt"
	"math"
	"strings"
)

const (
	AuditID = "GATE810-BOUNDARY-FN-RESTPRESSURE-SPURION-NEFF-MINUS-THREE-CLOSURE-AUDIT"

	NEff      = 3.0023273474722147
	DeltaN    = NEff - 3.0
	SBoundary = 0.0012924448188162962
	CYukawa   = 0.9992248188812008
	CHistory  = 1.038025177923625
	CHiggs    = 1.0372205204048603

	NineOver5   = 9.0 / 5.0
	ThreeOver10 = 3.0 / 10.0

	StatusGate809Inherited        = "PASS_GATE809_HIERARCHY_BREAKING_AUDIT_INHERITED"
	StatusBoundaryFNSelected      = "PASS_BOUNDARY_FN_REST_PRESSURE_CANDIDATE_SELECTED_AS_CURRENT_TEST_TARGET"
	StatusDirectClosure           = "PASS_DIRECT_BOUNDARY_CLOSURE_RESIDUAL_COMPUTED"
	StatusSpurionDefined          = "PASS_BOUNDARY_FN_SPURION_DEFINED"
	StatusEpsilonsComputed        = "PASS_EPSILON_N_AND_EPSILON_B_COMPUTED"
	StatusNineFifthsSourceAudited = "PASS_NINE_OVER_FIVE_COEFFICIENT_SOURCE_CANDIDATES_AUDITED"
	StatusAlphaClosure            = "PASS_TOP_REST_ALPHA_BOUNDARY_CLOSURE_COMPUTED"
	StatusExactPositivity         = "PASS_EXACT_TOP_REST_POSITIVITY_AUDIT_COMPLETED"
	StatusAlphaBetaPositivity     = "PASS_ALPHA_EQUALS_THREE_OVER_TEN_S_TESTED_AGAINST_BETA_POSITIVITY"
	StatusConcentrationRegimes    = "PASS_REST_CONCENTRATION_REGIMES_AUDITED"
	StatusMapRequirement          = "PASS_BOUNDARY_TO_REST_PRESSURE_MAP_REQUIREMENT_DEFINED"
	StatusCoeffControls           = "PASS_ALTERNATIVE_COEFFICIENT_CONTROLS_DEFINED"
	StatusCYukawaRewrite          = "PASS_C_YUKAWA_BOUNDARY_FN_REWRITE_CANDIDATE_DEFINED"
	StatusCHiggsImpact            = "PASS_C_HIGGS_IMPACT_OF_BOUNDARY_FN_CANDIDATE_AUDITED"
	StatusPhysicalFirewalls       = "PASS_PHYSICAL_FIREWALLS_ENFORCED"
	StatusBranchDecision          = "PASS_BRANCH_DECISION_RECORDED"

	StatusDeltaNApproxNineFifths     = "CONDITIONAL_SUPPORT_DELTA_N_APPROXIMATES_NINE_OVER_FIVE_TIMES_BOUNDARY_SPLIT"
	StatusBoundaryScaleCapable       = "CONDITIONAL_SUPPORT_BOUNDARY_SPLIT_IS_NUMERICALLY_CAPABLE_OF_SOURCING_REST_PRESSURE_SCALE"
	StatusBoundarySpurionScale       = "CONDITIONAL_SUPPORT_BOUNDARY_SPLIT_CAN_DEFINE_FN_STYLE_SPURION_SCALE"
	StatusEpsilonClose               = "CONDITIONAL_SUPPORT_EPSILON_B_APPROXIMATES_EPSILON_N_AT_ONEE_MINUS_FOUR_RELATIVE_LEVEL"
	StatusNineFifthsTypedCandidate   = "CONDITIONAL_SUPPORT_NINE_OVER_FIVE_HAS_COLOR_THREE_TIMES_INVERSE_HYPERCHARGE_NORMALIZATION_SOURCE_CANDIDATE"
	StatusFiveThirdsNonarbitrary     = "CONDITIONAL_SUPPORT_EXISTING_FIVE_OVER_THREE_COEFFICIENT_MAKES_NINE_OVER_FIVE_NONARBITRARY_ENOUGH_TO_TEST"
	StatusAlphaApproxThreeTenths     = "CONDITIONAL_SUPPORT_ALPHA_APPROXIMATES_THREE_OVER_TEN_TIMES_BOUNDARY_SPLIT"
	StatusThreeTenthsTypedCandidate  = "CONDITIONAL_SUPPORT_THREE_OVER_TEN_HAS_HALF_TIMES_INVERSE_HYPERCHARGE_NORMALIZATION_SOURCE_CANDIDATE"
	StatusThreeTenthsCloseNotExact   = "CONDITIONAL_SUPPORT_THREE_OVER_TEN_S_IS_CLOSE_BUT_NOT_EXACTLY_POSITIVE_REST_COMPATIBLE"
	StatusCorrectionAboveThreeTenths = "CONDITIONAL_SUPPORT_EXACT_POSITIVE_REST_MODEL_REQUIRES_SMALL_CORRECTION_ABOVE_THREE_OVER_TEN_S"
	StatusNarrowCorridor             = "CONDITIONAL_SUPPORT_POSITIVE_REST_CORRIDOR_REMAINS_NARROW_AROUND_ALPHA_OVER_S_APPROX_0_300"
	StatusConcentrationControls      = "CONDITIONAL_SUPPORT_REST_CONCENTRATION_CONTROLS_SMALL_CORRECTION_TO_THREE_OVER_TEN"
	StatusMapWouldReduceNEff         = "CONDITIONAL_SUPPORT_BOUNDARY_FN_MAP_WOULD_REDUCE_N_EFF_IF_CERTIFIED"
	StatusMapPreciseMissing          = "CONDITIONAL_SUPPORT_BOUNDARY_FN_MAP_IS_NOW_THE_PRECISE_MISSING_OBJECT"
	StatusBestTypedCandidate         = "CONDITIONAL_SUPPORT_NINE_OVER_FIVE_IS_CURRENT_BEST_TYPED_LOW_COMPLEXITY_CANDIDATE_IF_RESIDUAL_AND_SOURCE_BOTH_PASS"
	StatusCertifiedMapReducesSeal    = "CONDITIONAL_SUPPORT_CERTIFIED_BOUNDARY_FN_MAP_WOULD_REDUCE_YUKAWA_SEAL_DEPENDENCE"
	StatusNextHyperchargeColor       = "CONDITIONAL_SUPPORT_NEXT_GATE_SHOULD_AUDIT_HYPERCHARGE_COLOR_POSITIVE_REST_CORRECTION"

	StatusNineFifthsNotExact        = "FAILED_ROUTE_NINE_OVER_FIVE_CLOSURE_NOT_EXACT_AT_CURRENT_LEDGER_PRECISION"
	StatusNumericalNotTheorem       = "FAILED_ROUTE_NUMERICAL_CLOSURE_NOT_NATIVE_HIERARCHY_BREAKING_THEOREM"
	StatusEpsilonBNotNative         = "FAILED_ROUTE_EPSILON_B_NOT_NATIVE_FN_SPURION_WITHOUT_OPERATOR"
	StatusNoBoundarySpurionMap      = "FAILED_ROUTE_NO_NATIVE_BOUNDARY_TO_YUKAWA_SPURION_MAP"
	StatusCabibboNotTheorem         = "FAILED_ROUTE_EPSILON_CLOSE_TO_CABIBBO_SCALE_NOT_YUKAWA_THEOREM"
	StatusNoColorHyperchargeTheorem = "FAILED_ROUTE_NO_NATIVE_COLOR_HYPERCHARGE_TO_REST_PRESSURE_COEFFICIENT_THEOREM"
	StatusInverseHyperchargeNotAuto = "FAILED_ROUTE_INVERSE_HYPERCHARGE_NORMALIZATION_NOT_AUTOMATICALLY_YUKAWA_HIERARCHY_OPERATOR"
	StatusNoRationalFit             = "FAILED_ROUTE_NINE_OVER_FIVE_MUST_NOT_BE_ACCEPTED_BY_RATIONAL_FIT_ALONE"
	StatusThreeTenthsNotNative      = "FAILED_ROUTE_THREE_OVER_TEN_NOT_NATIVE_ALPHA_THEOREM"
	StatusNoAlphaReadoutMap         = "FAILED_ROUTE_NO_BOUNDARY_TO_ALPHA_READOUT_MAP"
	StatusAlphaApproxNotExact       = "FAILED_ROUTE_ALPHA_APPROXIMATION_NOT_EXACT_REST_PRESSURE_THEOREM"
	StatusAlphaThreeTenthsNotExact  = "FAILED_ROUTE_ALPHA_EQUALS_THREE_OVER_TEN_S_NOT_EXACT_WITH_BETA_NONNEGATIVE"
	StatusFirstOrderNotTheorem      = "FAILED_ROUTE_FIRST_ORDER_ALPHA_CLOSURE_MUST_NOT_BE_PROMOTED_TO_EXACT_THEOREM"
	StatusPositiveBlocksExact       = "FAILED_ROUTE_POSITIVE_REST_SPECTRUM_BLOCKS_EXACT_THREE_OVER_TEN_ALPHA_WITHOUT_CORRECTION"
	StatusAggregateNoQRest          = "FAILED_ROUTE_AGGREGATE_LEDGER_DOES_NOT_FIX_Q_REST"
	StatusNoRestAtoms               = "FAILED_ROUTE_NO_REST_ATOM_COUNT_OR_SECTOR_COMPOSITION_FROM_THIS_CLOSURE"
	StatusConcentrationNotNative    = "FAILED_ROUTE_REST_CONCENTRATION_NOT_NATIVE_WITHOUT_TRACE_ATOM_LEDGER"
	StatusNoBoundaryFNMap           = "FAILED_ROUTE_NO_NATIVE_BOUNDARY_FN_REST_PRESSURE_MAP"
	StatusNoSectorTraceRule         = "FAILED_ROUTE_NO_SECTOR_SUM_TRACE_RULE_FROM_BOUNDARY_SPLIT"
	StatusNoPositiveConcentration   = "FAILED_ROUTE_NO_POSITIVE_REST_CONCENTRATION_LAW"
	StatusNoScaleStability          = "FAILED_ROUTE_NO_SCALE_STABILITY_FOR_BOUNDARY_FN_RELATION"
	StatusLowDenominatorNotTheorem  = "FAILED_ROUTE_LOW_DENOMINATOR_RATIONAL_FITTING_IS_NOT_THEOREM"
	StatusBestNumericalNeedsType    = "FAILED_ROUTE_BEST_NUMERICAL_RATIONAL_NOT_ACCEPTED_WITHOUT_TYPED_SOURCE"
	StatusNoCYukawaUpdate           = "FAILED_ROUTE_GATE810_DOES_NOT_UPDATE_C_YUKAWA_WITHOUT_CERTIFICATION"
	StatusCHiggsLevelB              = "FAILED_ROUTE_C_HIGGS_REMAINS_LEVEL_B"
	StatusCandidateNotLevelC        = "FAILED_ROUTE_CANDIDATE_REWRITE_NOT_LEVEL_C_PREDICTION"
	StatusTreeProxyNotPole          = "FAILED_ROUTE_TREE_PROXY_NOT_POLE_MASS"
	StatusFirewallGate810           = "FIREWALL_PRESERVED_GATE810_BOUNDARY_FN_REST_PRESSURE_CLOSURE_BOUNDARY"
)

type Inheritance struct {
	Gate809Inherited                           bool
	CandidateSelected                          bool
	NEff, DeltaN, S, CYukawa, CHistory, CHiggs float64
	Verdicts                                   []string
}
type DirectClosure struct {
	Computed                                                         bool
	CObs, CandidateCoeff, CandidateDelta, Residual, RelativeResidual float64
	Supports, Failures                                               []string
	Verdict                                                          string
}
type Spurion struct {
	Defined                                            bool
	EpsilonN, EpsilonB, Difference, RelativeDifference float64
	Supports, Failures                                 []string
	Verdicts                                           []string
}
type CoefficientSource struct {
	Audited                           bool
	Coefficient                       string
	CandidateSources, Interpretations []string
	Supports, Failures                []string
	Verdict                           string
}
type AlphaClosure struct {
	Computed                             bool
	AlphaApprox, AlphaBoundary, Residual float64
	CoefficientSource                    string
	Supports, Failures                   []string
	Verdict                              string
}
type PositivityAudit struct {
	Completed                                                              bool
	AlphaBoundary, BetaRequired, AlphaMin, AlphaMinOverS, CorrectionNeeded float64
	Supports, Failures                                                     []string
	Verdicts                                                               []string
}
type RestRegime struct {
	Name                           string
	Alpha, AlphaOverS, Beta, QRest float64
	Description                    string
}
type ConcentrationAudit struct {
	Audited            bool
	Regimes            []RestRegime
	Supports, Failures []string
	Verdict            string
}
type MapRequirement struct {
	Defined            bool
	Name               string
	Chain, Components  []string
	Supports, Failures []string
	Verdict            string
}
type CoefficientControl struct {
	Name                                                 string
	Coefficient, Residual, AbsResidual, RelativeResidual float64
	TypedSource                                          string
}
type CoefficientControls struct {
	Defined            bool
	Controls           []CoefficientControl
	BestTyped          string
	Supports, Failures []string
	Verdict            string
}
type CHiggsImpact struct {
	Defined                                                            bool
	CYukawaCandidate, CYukawaResidual, CHiggsCandidate, CHiggsResidual float64
	Formula, CandidateRewrite                                          string
	Supports, Failures                                                 []string
	Verdicts                                                           []string
}
type BranchDecision struct {
	Recorded               bool
	Next, Purpose, Verdict string
	Supports               []string
}
type Firewalls struct {
	Enforced                                                                                                                       bool
	NoNativeNEff, NoNativeFN, NoRationalFit, NoAlphaExact, NoYukawaSpectrum, NoCertifiedRewrite, NoLevelC, NoPoleMass, NoD4KoideGJ bool
	Verdict                                                                                                                        string
}
type Analysis struct {
	Inheritance   Inheritance
	Direct        DirectClosure
	Spurion       Spurion
	Coefficient   CoefficientSource
	Alpha         AlphaClosure
	Positivity    PositivityAudit
	Concentration ConcentrationAudit
	Map           MapRequirement
	Controls      CoefficientControls
	CHiggs        CHiggsImpact
	Branch        BranchDecision
	Firewalls     Firewalls
	Truth, Final  string
}

func EpsilonFromDelta(delta float64) (float64, error) {
	if delta <= 0 {
		return 0, fmt.Errorf("delta must be positive")
	}
	return math.Pow(delta, 0.25), nil
}
func EpsilonFromBoundary(s float64) (float64, error) {
	if s <= 0 {
		return 0, fmt.Errorf("boundary split must be positive")
	}
	return math.Pow(NineOver5*s, 0.25), nil
}
func BoundaryClosureResidual(delta, s float64) float64 { return delta - NineOver5*s }
func AlphaApproxFromDelta(delta float64) float64       { return delta / 6.0 }
func AlphaBoundaryFromS(s float64) float64             { return ThreeOver10 * s }
func BetaRequired(alpha, neff float64) float64         { return 3.0*math.Pow(1.0+alpha, 2.0)/neff - 1.0 }
func AlphaMinBetaZero(neff float64) float64            { return math.Sqrt(neff/3.0) - 1.0 }
func CYukawaBoundaryFN(s float64) float64              { return 3.0 / (3.0 + NineOver5*s) }
func CHiggsBoundaryFN(s float64) float64               { return CYukawaBoundaryFN(s) * CHistory }
func QRest(alpha, beta float64) float64                { return beta / (3.0 * alpha * alpha) }
func SolveAlphaForQ(neff, q float64) float64 { // N = 3(1+a)^2/(1+3a^2q), use small positive root.
	A := 3.0*neff*q - 3.0
	B := -6.0
	C := neff - 3.0
	disc := B*B - 4*A*C
	if disc < 0 {
		return math.NaN()
	}
	r1 := (-B + math.Sqrt(disc)) / (2 * A)
	r2 := (-B - math.Sqrt(disc)) / (2 * A)
	if r1 > 0 && (r1 < r2 || r2 <= 0) {
		return r1
	}
	return r2
}

func BuildDefault() (Analysis, error) {
	epsN, err := EpsilonFromDelta(DeltaN)
	if err != nil {
		return Analysis{}, err
	}
	epsB, err := EpsilonFromBoundary(SBoundary)
	if err != nil {
		return Analysis{}, err
	}
	directResidual := BoundaryClosureResidual(DeltaN, SBoundary)
	cObs := DeltaN / SBoundary
	alphaApprox := AlphaApproxFromDelta(DeltaN)
	alphaB := AlphaBoundaryFromS(SBoundary)
	betaReq := BetaRequired(alphaB, NEff)
	alphaMin := AlphaMinBetaZero(NEff)
	betaDiag := BetaRequired(alphaApprox, NEff)
	alphaQ1 := SolveAlphaForQ(NEff, 1.0)
	controls := []CoefficientControl{
		{Name: "c = 2", Coefficient: 2, Residual: DeltaN - 2*SBoundary, AbsResidual: math.Abs(DeltaN - 2*SBoundary), RelativeResidual: (DeltaN - 2*SBoundary) / DeltaN, TypedSource: "simple boundary double response; weaker residual"},
		{Name: "c = 7/4", Coefficient: 7.0 / 4.0, Residual: DeltaN - (7.0/4.0)*SBoundary, AbsResidual: math.Abs(DeltaN - (7.0/4.0)*SBoundary), RelativeResidual: (DeltaN - (7.0/4.0)*SBoundary) / DeltaN, TypedSource: "K7/four-chamber style candidate; weaker residual"},
		{Name: "c = 13/7", Coefficient: 13.0 / 7.0, Residual: DeltaN - (13.0/7.0)*SBoundary, AbsResidual: math.Abs(DeltaN - (13.0/7.0)*SBoundary), RelativeResidual: (DeltaN - (13.0/7.0)*SBoundary) / DeltaN, TypedSource: "prior K7/Hodge obstruction ratio candidate; weaker residual"},
		{Name: "c = 9/5", Coefficient: NineOver5, Residual: directResidual, AbsResidual: math.Abs(directResidual), RelativeResidual: directResidual / DeltaN, TypedSource: "color-three × inverse hypercharge normalization candidate"},
	}
	return Analysis{
		Inheritance:   Inheritance{Gate809Inherited: true, CandidateSelected: true, NEff: NEff, DeltaN: DeltaN, S: SBoundary, CYukawa: CYukawa, CHistory: CHistory, CHiggs: CHiggs, Verdicts: []string{StatusGate809Inherited, StatusBoundaryFNSelected}},
		Direct:        DirectClosure{Computed: true, CObs: cObs, CandidateCoeff: NineOver5, CandidateDelta: NineOver5 * SBoundary, Residual: directResidual, RelativeResidual: directResidual / DeltaN, Verdict: StatusDirectClosure, Supports: []string{StatusDeltaNApproxNineFifths, StatusBoundaryScaleCapable}, Failures: []string{StatusNineFifthsNotExact, StatusNumericalNotTheorem}},
		Spurion:       Spurion{Defined: true, EpsilonN: epsN, EpsilonB: epsB, Difference: epsN - epsB, RelativeDifference: (epsN - epsB) / epsN, Verdicts: []string{StatusSpurionDefined, StatusEpsilonsComputed}, Supports: []string{StatusBoundarySpurionScale, StatusEpsilonClose}, Failures: []string{StatusEpsilonBNotNative, StatusNoBoundarySpurionMap, StatusCabibboNotTheorem}},
		Coefficient:   CoefficientSource{Audited: true, Coefficient: "9/5 = 3 × (3/5)", CandidateSources: []string{"3: color multiplicity / top-color baseline", "3/5: inverse hypercharge normalization from the active 5/3 coefficient"}, Interpretations: []string{"color/hypercharge-normalized boundary split response", "nonarbitrary candidate only, not theorem"}, Verdict: StatusNineFifthsSourceAudited, Supports: []string{StatusNineFifthsTypedCandidate, StatusFiveThirdsNonarbitrary}, Failures: []string{StatusNoColorHyperchargeTheorem, StatusInverseHyperchargeNotAuto, StatusNoRationalFit}},
		Alpha:         AlphaClosure{Computed: true, AlphaApprox: alphaApprox, AlphaBoundary: alphaB, Residual: alphaApprox - alphaB, CoefficientSource: "3/10 = (1/2)(3/5)", Verdict: StatusAlphaClosure, Supports: []string{StatusAlphaApproxThreeTenths, StatusThreeTenthsTypedCandidate}, Failures: []string{StatusThreeTenthsNotNative, StatusNoAlphaReadoutMap, StatusAlphaApproxNotExact}},
		Positivity:    PositivityAudit{Completed: true, AlphaBoundary: alphaB, BetaRequired: betaReq, AlphaMin: alphaMin, AlphaMinOverS: alphaMin / SBoundary, CorrectionNeeded: alphaMin - alphaB, Verdicts: []string{StatusExactPositivity, StatusAlphaBetaPositivity}, Supports: []string{StatusThreeTenthsCloseNotExact, StatusCorrectionAboveThreeTenths}, Failures: []string{StatusAlphaThreeTenthsNotExact, StatusFirstOrderNotTheorem, StatusPositiveBlocksExact}},
		Concentration: ConcentrationAudit{Audited: true, Regimes: []RestRegime{{Name: "maximally dilute rest", Alpha: alphaMin, AlphaOverS: alphaMin / SBoundary, Beta: 0, QRest: 0, Description: "beta -> 0 positivity boundary"}, {Name: "first-order diagnostic", Alpha: alphaApprox, AlphaOverS: alphaApprox / SBoundary, Beta: betaDiag, QRest: QRest(alphaApprox, betaDiag), Description: "alpha=Delta_N/6 gives q_rest≈1/N_eff"}, {Name: "single-rest-concentrated corridor", Alpha: alphaQ1, AlphaOverS: alphaQ1 / SBoundary, Beta: 3 * alphaQ1 * alphaQ1, QRest: 1, Description: "q_rest≈1 upper concentration"}}, Verdict: StatusConcentrationRegimes, Supports: []string{StatusNarrowCorridor, StatusConcentrationControls}, Failures: []string{StatusAggregateNoQRest, StatusNoRestAtoms, StatusConcentrationNotNative}},
		Map:           MapRequirement{Defined: true, Name: "BoundaryFNRestPressureMap", Chain: []string{"s", "epsilon_B^4=(9/5)s", "Delta_N=N_eff-3", "N_eff", "C_Yukawa=3/N_eff"}, Components: []string{"boundary split coordinate", "hypercharge/color coefficient source", "FN-style spurion", "top/rest alpha-beta map", "positive-rest concentration law", "sector-summed trace rule", "scale/scheme convention", "noncircularity proof"}, Verdict: StatusMapRequirement, Supports: []string{StatusMapWouldReduceNEff, StatusMapPreciseMissing}, Failures: []string{StatusNoBoundaryFNMap, StatusNoSectorTraceRule, StatusNoPositiveConcentration, StatusNoScaleStability}},
		Controls:      CoefficientControls{Defined: true, Controls: controls, BestTyped: "c=9/5 combines best residual with color-three × inverse hypercharge typing", Verdict: StatusCoeffControls, Supports: []string{StatusBestTypedCandidate}, Failures: []string{StatusLowDenominatorNotTheorem, StatusBestNumericalNeedsType}},
		CHiggs:        CHiggsImpact{Defined: true, CYukawaCandidate: CYukawaBoundaryFN(SBoundary), CYukawaResidual: CYukawaBoundaryFN(SBoundary) - CYukawa, CHiggsCandidate: CHiggsBoundaryFN(SBoundary), CHiggsResidual: CHiggsBoundaryFN(SBoundary) - CHiggs, Formula: "C_Higgs=(3/N_eff)C_History", CandidateRewrite: "if certified: C_Yukawa=3/[3+(9/5)s] and C_Higgs=C_History·3/[3+(9/5)s]", Verdicts: []string{StatusCYukawaRewrite, StatusCHiggsImpact}, Supports: []string{StatusCertifiedMapReducesSeal}, Failures: []string{StatusNoCYukawaUpdate, StatusCHiggsLevelB, StatusCandidateNotLevelC}},
		Branch:        BranchDecision{Recorded: true, Next: "Gate 811 — Hypercharge-Color Boundary Coefficient and Positive-Rest Correction Audit", Purpose: "audit 9/5=3×(3/5), 3/10=(1/2)(3/5), and the positivity-compatible correction above (3/10)s", Verdict: StatusBranchDecision, Supports: []string{StatusNextHyperchargeColor}},
		Firewalls:     Firewalls{Enforced: true, NoNativeNEff: true, NoNativeFN: true, NoRationalFit: true, NoAlphaExact: true, NoYukawaSpectrum: true, NoCertifiedRewrite: true, NoLevelC: true, NoPoleMass: true, NoD4KoideGJ: true, Verdict: StatusFirewallGate810},
		Truth:         "Gate 810 finds that N_eff-3≈(9/5)s is a strong boundary-FN candidate but not exact and not native.",
		Final:         "The missing object is BoundaryFNRestPressureMap: coefficient source, spurion map, positivity correction, rest concentration law, sector trace rule, scale convention, and noncircularity proof.",
	}, nil
}

func Statuses() []string {
	return []string{StatusGate809Inherited, StatusBoundaryFNSelected, StatusDirectClosure, StatusSpurionDefined, StatusEpsilonsComputed, StatusNineFifthsSourceAudited, StatusAlphaClosure, StatusExactPositivity, StatusAlphaBetaPositivity, StatusConcentrationRegimes, StatusMapRequirement, StatusCoeffControls, StatusCYukawaRewrite, StatusCHiggsImpact, StatusPhysicalFirewalls, StatusBranchDecision, StatusDeltaNApproxNineFifths, StatusBoundaryScaleCapable, StatusBoundarySpurionScale, StatusEpsilonClose, StatusNineFifthsTypedCandidate, StatusFiveThirdsNonarbitrary, StatusAlphaApproxThreeTenths, StatusThreeTenthsTypedCandidate, StatusThreeTenthsCloseNotExact, StatusCorrectionAboveThreeTenths, StatusNarrowCorridor, StatusConcentrationControls, StatusMapWouldReduceNEff, StatusMapPreciseMissing, StatusBestTypedCandidate, StatusCertifiedMapReducesSeal, StatusNextHyperchargeColor, StatusNineFifthsNotExact, StatusNumericalNotTheorem, StatusEpsilonBNotNative, StatusNoBoundarySpurionMap, StatusCabibboNotTheorem, StatusNoColorHyperchargeTheorem, StatusInverseHyperchargeNotAuto, StatusNoRationalFit, StatusThreeTenthsNotNative, StatusNoAlphaReadoutMap, StatusAlphaApproxNotExact, StatusAlphaThreeTenthsNotExact, StatusFirstOrderNotTheorem, StatusPositiveBlocksExact, StatusAggregateNoQRest, StatusNoRestAtoms, StatusConcentrationNotNative, StatusNoBoundaryFNMap, StatusNoSectorTraceRule, StatusNoPositiveConcentration, StatusNoScaleStability, StatusLowDenominatorNotTheorem, StatusBestNumericalNeedsType, StatusNoCYukawaUpdate, StatusCHiggsLevelB, StatusCandidateNotLevelC, StatusTreeProxyNotPole, StatusFirewallGate810}
}

func FormatDirect(d DirectClosure) string {
	return fmt.Sprintf("c_obs=%.17g c_B=%.17g candidate=%.17g residual=%.17g relative=%.17g supports=[%s] failures=[%s]", d.CObs, d.CandidateCoeff, d.CandidateDelta, d.Residual, d.RelativeResidual, strings.Join(d.Supports, "; "), strings.Join(d.Failures, "; "))
}
func FormatSpurion(s Spurion) string {
	return fmt.Sprintf("epsilon_N=%.17g epsilon_B=%.17g diff=%.17g relDiff=%.17g supports=[%s] failures=[%s]", s.EpsilonN, s.EpsilonB, s.Difference, s.RelativeDifference, strings.Join(s.Supports, "; "), strings.Join(s.Failures, "; "))
}
func FormatCoeff(c CoefficientSource) string {
	return fmt.Sprintf("%s sources=[%s] interpretations=[%s] supports=[%s] failures=[%s]", c.Coefficient, strings.Join(c.CandidateSources, "; "), strings.Join(c.Interpretations, "; "), strings.Join(c.Supports, "; "), strings.Join(c.Failures, "; "))
}
func FormatAlpha(a AlphaClosure) string {
	return fmt.Sprintf("alphaApprox=%.17g alphaB=%.17g residual=%.17g source=%s supports=[%s] failures=[%s]", a.AlphaApprox, a.AlphaBoundary, a.Residual, a.CoefficientSource, strings.Join(a.Supports, "; "), strings.Join(a.Failures, "; "))
}
func FormatPositivity(p PositivityAudit) string {
	return fmt.Sprintf("alphaB=%.17g betaRequired=%.17g alphaMin=%.17g alphaMin/s=%.17g correction=%.17g supports=[%s] failures=[%s]", p.AlphaBoundary, p.BetaRequired, p.AlphaMin, p.AlphaMinOverS, p.CorrectionNeeded, strings.Join(p.Supports, "; "), strings.Join(p.Failures, "; "))
}
func FormatConcentration(c ConcentrationAudit) string {
	parts := []string{}
	for _, r := range c.Regimes {
		parts = append(parts, fmt.Sprintf("%s alpha=%.17g alpha/s=%.17g beta=%.17g q=%.17g", r.Name, r.Alpha, r.AlphaOverS, r.Beta, r.QRest))
	}
	return fmt.Sprintf("regimes=[%s] supports=[%s] failures=[%s]", strings.Join(parts, "; "), strings.Join(c.Supports, "; "), strings.Join(c.Failures, "; "))
}
func FormatMap(m MapRequirement) string {
	return fmt.Sprintf("%s chain=[%s] components=[%s] supports=[%s] failures=[%s]", m.Name, strings.Join(m.Chain, " -> "), strings.Join(m.Components, "; "), strings.Join(m.Supports, "; "), strings.Join(m.Failures, "; "))
}
func FormatControls(c CoefficientControls) string {
	parts := []string{}
	for _, x := range c.Controls {
		parts = append(parts, fmt.Sprintf("%s coeff=%.17g residual=%.17g abs=%.17g type=%s", x.Name, x.Coefficient, x.Residual, x.AbsResidual, x.TypedSource))
	}
	return fmt.Sprintf("controls=[%s] best=%s supports=[%s] failures=[%s]", strings.Join(parts, "; "), c.BestTyped, strings.Join(c.Supports, "; "), strings.Join(c.Failures, "; "))
}
func FormatCHiggs(c CHiggsImpact) string {
	return fmt.Sprintf("CYukawaCandidate=%.17g dCY=%.17g CHiggsCandidate=%.17g dCH=%.17g supports=[%s] failures=[%s]", c.CYukawaCandidate, c.CYukawaResidual, c.CHiggsCandidate, c.CHiggsResidual, strings.Join(c.Supports, "; "), strings.Join(c.Failures, "; "))
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
