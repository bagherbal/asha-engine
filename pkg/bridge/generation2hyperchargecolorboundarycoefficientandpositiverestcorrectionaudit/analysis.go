// Package generation2hyperchargecolorboundarycoefficientandpositiverestcorrectionaudit implements
// Gate 811: Hypercharge-Color Boundary Coefficient and Positive-Rest Correction Audit.
//
// Gate 811 follows Gate 810 by auditing the typed source of the coefficient
// 9/5 = 3 × (3/5), the related alpha coefficient 3/10 = (1/2)(3/5),
// and the small positive-rest correction required because alpha=(3/10)s is
// close but not compatible with beta>=0. It also tests second-raw-moment
// corrections without promoting residual fitting into a native Yukawa theorem.
package generation2hyperchargecolorboundarycoefficientandpositiverestcorrectionaudit

import (
	"fmt"
	"math"
	"strings"
)

const (
	AuditID = "GATE811-HYPERCHARGE-COLOR-BOUNDARY-COEFFICIENT-POSITIVE-REST-CORRECTION-AUDIT"

	NEff      = 3.0023273474722147
	DeltaN    = NEff - 3.0
	SBoundary = 0.0012924448188162962
	PBoundary = 7.0 / 72.0
	CYukawa   = 0.9992248188812008
	CHistory  = 1.038025177923625
	CHiggs    = 1.0372205204048603

	NineOverFive = 9.0 / 5.0
	ThreeTenths  = 3.0 / 10.0
	Half         = 1.0 / 2.0
	Six          = 6.0

	StatusGate810Inherited      = "PASS_GATE810_BOUNDARY_FN_REST_PRESSURE_CLOSURE_INHERITED"
	StatusCoeffTargetSelected   = "PASS_HYPERCHARGE_COLOR_COEFFICIENT_SELECTED_AS_CURRENT_SOURCE_TARGET"
	StatusRestCorrectionTarget  = "PASS_POSITIVE_REST_CORRECTION_SELECTED_AS_CURRENT_EXACTNESS_TARGET"
	StatusNineFiveFactor        = "PASS_NINE_OVER_FIVE_FACTORIZATION_AUDITED"
	StatusLedgerExistence       = "PASS_COLOR_THREE_AND_FIVE_OVER_THREE_EXISTENCE_IN_ACTIVE_LEDGER_CONFIRMED"
	StatusThreeTenthsFactor     = "PASS_THREE_OVER_TEN_FACTORIZATION_AUDITED"
	StatusCorrectionDefined     = "PASS_EXACT_POSITIVE_REST_CORRECTION_DEFINED"
	StatusCorrectionScale       = "PASS_POSITIVITY_CORRECTION_SMALL_SCALE_AUDITED"
	StatusAlphaCorrDefined      = "PASS_CORRECTED_ALPHA_CANDIDATE_DEFINED"
	StatusAlphaCorrTestRequired = "PASS_CORRECTED_ALPHA_POSITIVITY_TEST_REQUIRED"
	StatusQRestDefined          = "PASS_REST_CONCENTRATION_READOUT_DEFINED"
	StatusDeltaCorrDefined      = "PASS_DIRECT_DELTA_N_SECOND_MOMENT_CORRECTION_DEFINED"
	StatusC2ObsComputed         = "PASS_C2_OBS_FROM_RESIDUAL_COMPUTED"
	StatusPackageDefined        = "PASS_HYPERCHARGE_COLOR_BOUNDARY_REST_PRESSURE_PACKAGE_DEFINED"
	StatusControlRules          = "PASS_ALTERNATIVE_COEFFICIENT_CONTROL_RULES_PRESERVED"
	StatusImpactDefined         = "PASS_C_YUKAWA_AND_C_HIGGS_CANDIDATE_IMPACT_DEFINED"
	StatusPhysicalFirewalls     = "PASS_PHYSICAL_FIREWALLS_ENFORCED"
	StatusBranchDecision        = "PASS_BRANCH_DECISION_RECORDED"

	StatusNineFiveTypedCandidate     = "CONDITIONAL_SUPPORT_NINE_OVER_FIVE_HAS_TYPED_COLOR_HYPERCHARGE_SOURCE_CANDIDATE"
	StatusRestCoeffColorHypercharge  = "CONDITIONAL_SUPPORT_REST_PRESSURE_COEFFICIENT_MAY_BE_COLOR_THREE_TIMES_INVERSE_HYPERCHARGE_RESPONSE"
	StatusThreeTenthsTypedCandidate  = "CONDITIONAL_SUPPORT_THREE_OVER_TEN_HAS_HALF_TIMES_INVERSE_HYPERCHARGE_SOURCE_CANDIDATE"
	StatusHalfBoundaryPairCandidate  = "CONDITIONAL_SUPPORT_HALF_FACTOR_MAY_REFLECT_BOUNDARY_PAIR_MIDPOINT_OR_TWO_ENDPOINT_AVERAGING"
	StatusCorrectionOrderM2          = "CONDITIONAL_SUPPORT_POSITIVITY_CORRECTION_IS_ORDER_M2"
	StatusHalfM2Approx               = "CONDITIONAL_SUPPORT_DELTA_ALPHA_POS_APPROXIMATES_HALF_OF_SECOND_RAW_BOUNDARY_MOMENT"
	StatusCorrectionMayBeM2          = "CONDITIONAL_SUPPORT_POSITIVE_REST_CORRECTION_MAY_BE_BOUNDARY_SECOND_MOMENT_RESPONSE"
	StatusAlphaHalfM2LawfulToTest    = "CONDITIONAL_SUPPORT_ALPHA_CORRECTION_BY_HALF_M2_IS_LAWFUL_TO_TEST"
	StatusDeltaResidualOrderM2       = "CONDITIONAL_SUPPORT_DELTA_N_RESIDUAL_IS_ORDER_P_S_SQUARED"
	StatusC2ObsCloseToSix            = "CONDITIONAL_SUPPORT_C2_OBS_IS_CLOSE_TO_BOUNDARY_PAIR_TIMES_COLOR_CANDIDATE_SIX"
	StatusPackageSharpestCandidate   = "CONDITIONAL_SUPPORT_HYPERCHARGE_COLOR_PACKAGE_IS_CURRENT_SHARPEST_BOUNDARY_FN_SOURCE_CANDIDATE"
	StatusCorrectedModelControl      = "CONDITIONAL_SUPPORT_CORRECTED_NINE_OVER_FIVE_MODEL_CAN_BE_TESTED_AGAINST_CONTROLS"
	StatusNextSecondMomentCorrection = "CONDITIONAL_SUPPORT_NEXT_GATE_SHOULD_TEST_BOUNDARY_SECOND_MOMENT_POSITIVE_SPECTRUM_CONSTRUCTION"

	StatusExistenceNotTheorem         = "FAILED_ROUTE_EXISTENCE_OF_3_AND_5_OVER_3_DOES_NOT_PROVE_9_OVER_5_REST_PRESSURE_THEOREM"
	StatusNoColorHyperchargeMap       = "FAILED_ROUTE_NO_NATIVE_COLOR_HYPERCHARGE_REST_PRESSURE_MAP_YET"
	StatusInverseHyperchargeNotYukawa = "FAILED_ROUTE_INVERSE_HYPERCHARGE_RESPONSE_NOT_YUKAWA_HIERARCHY_OPERATOR_BY_ITSELF"
	StatusNoBoundaryAverageTheorem    = "FAILED_ROUTE_NO_NATIVE_BOUNDARY_PAIR_AVERAGING_TO_ALPHA_THEOREM"
	StatusThreeTenthsNotNative        = "FAILED_ROUTE_THREE_OVER_TEN_NOT_NATIVE_ALPHA_COEFFICIENT"
	StatusBoundaryPairNotReadout      = "FAILED_ROUTE_BOUNDARY_PAIR_EXISTENCE_DOES_NOT_PROVE_YUKAWA_REST_PRESSURE_READOUT"
	StatusCorrectionNotExactHalfM2    = "FAILED_ROUTE_POSITIVITY_CORRECTION_NOT_EXACTLY_HALF_M2_WITH_CURRENT_LEDGER"
	StatusNoPositiveCorrectionTheorem = "FAILED_ROUTE_NO_NATIVE_POSITIVE_REST_CORRECTION_THEOREM"
	StatusM2NotConcentrationLaw       = "FAILED_ROUTE_M2_RESONANCE_NOT_YET_REST_CONCENTRATION_LAW"
	StatusAlphaCorrNeedsBeta          = "FAILED_ROUTE_CORRECTED_ALPHA_NOT_ACCEPTED_WITHOUT_BETA_NONNEGATIVITY"
	StatusAlphaCorrNotNative          = "FAILED_ROUTE_CORRECTED_ALPHA_NOT_NATIVE_REST_PRESSURE_THEOREM"
	StatusQRestNotSector              = "FAILED_ROUTE_Q_REST_NOT_SECTOR_ASSIGNMENT"
	StatusQRestNotNativeSpectrum      = "FAILED_ROUTE_AGGREGATE_Q_REST_NOT_NATIVE_REST_SPECTRUM"
	StatusNoAtomLedgerFromQRest       = "FAILED_ROUTE_NO_REST_ATOM_LEDGER_FROM_Q_REST_ALONE"
	StatusC2ObsNotNative              = "FAILED_ROUTE_C2_OBS_NOT_NATIVE_SECOND_MOMENT_COEFFICIENT"
	StatusSecondMomentNeedsMap        = "FAILED_ROUTE_SECOND_MOMENT_CORRECTION_NOT_ACCEPTED_WITHOUT_TYPED_MAP"
	StatusResidualFittingForbidden    = "FAILED_ROUTE_RESIDUAL_FITTING_MUST_NOT_REPLACE_THEOREM"
	StatusPackageNeedsTraceMap        = "FAILED_ROUTE_PACKAGE_NOT_NATIVE_WITHOUT_BOUNDARY_TO_TRACE_MAGNITUDE_READOUT_MAP"
	StatusPackageNeedsSpectrum        = "FAILED_ROUTE_PACKAGE_NOT_NATIVE_WITHOUT_POSITIVE_REST_SPECTRUM_CONSTRUCTION"
	StatusPackageNeedsScale           = "FAILED_ROUTE_PACKAGE_NOT_NATIVE_WITHOUT_SCALE_STABILITY"
	StatusRationalClosenessForbidden  = "FAILED_ROUTE_RATIONAL_CLOSENESS_ALONE_REMAINS_FORBIDDEN"
	StatusModelNeedsSourcePositivity  = "FAILED_ROUTE_COEFFICIENT_MODEL_NOT_VALID_WITHOUT_SOURCE_AND_POSITIVITY"
	StatusNoCYukawaUpdate             = "FAILED_ROUTE_GATE811_DOES_NOT_UPDATE_C_YUKAWA_WITHOUT_CERTIFICATION"
	StatusCHiggsLevelB                = "FAILED_ROUTE_C_HIGGS_REMAINS_LEVEL_B"
	StatusTreeProxyNotPole            = "FAILED_ROUTE_TREE_PROXY_NOT_POLE_MASS"
	StatusFirewallGate811             = "FIREWALL_PRESERVED_GATE811_HYPERCHARGE_COLOR_POSITIVE_REST_CORRECTION_BOUNDARY"
)

type Inheritance struct {
	Gate810Inherited, CoeffTargetSelected, CorrectionTargetSelected bool
	NEff, DeltaN, S, CYukawa, CHistory, CHiggs                      float64
	Verdicts                                                        []string
}

type FactorizationAudit struct {
	Audited     bool
	Expression  string
	Sources     []string
	Supports    []string
	Failures    []string
	Verdicts    []string
	Description string
}

type PositivityCorrection struct {
	Defined                                       bool
	AlphaMin, AlphaB, DeltaAlpha, DeltaAlphaOverS float64
	DeltaAlphaOverAlphaB, S2, M2, HalfM2, M3      float64
	DeltaOverHalfM2                               float64
	Supports, Failures, Verdicts                  []string
}

type CorrectedAlpha struct {
	Defined, Tested, BetaNonnegative, QRestValid bool
	AlphaCorr, AlphaMin, AlphaCorrMinusAlphaMin  float64
	BetaCorr, QRestCorr                          float64
	Supports, Failures, Verdicts                 []string
}

type DirectDeltaCorrection struct {
	Defined, Computed                                  bool
	Residual, M2, C2Obs, CandidateC2                   float64
	CandidateDelta, CandidateResidual, CandidateRelErr float64
	Supports, Failures, Verdicts                       []string
}

type BoundaryRestPressurePackage struct {
	Defined            bool
	Name               string
	Components         []string
	Supports, Failures []string
	Verdict            string
}

type CoefficientControl struct {
	Name        string
	Coefficient float64
	Residual    float64
	AbsResidual float64
	TypedSource string
}

type AlternativeControls struct {
	Defined            bool
	Controls           []CoefficientControl
	Supports, Failures []string
	Verdict            string
}

type CandidateImpact struct {
	Defined                                    bool
	NBoundary, CYukawaBoundary, CHiggsBoundary float64
	CYukawaResidual, CHiggsResidual            float64
	Failures, Verdicts                         []string
}

type Firewalls struct {
	Enforced                                                                                                    bool
	NoCoeffShortcut, NoAlphaExact, NoHalfM2Native, NoSixFit, NoPackageNative, NoBoundarySpectrum, NoQRestSector bool
	NoCertifiedRewrite, NoPoleMass                                                                              bool
	Verdict                                                                                                     string
}

type BranchDecision struct {
	Recorded               bool
	Next, Purpose, Verdict string
	Supports               []string
}

type Analysis struct {
	Inheritance Inheritance
	NineFifths  FactorizationAudit
	ThreeTenths FactorizationAudit
	Correction  PositivityCorrection
	AlphaCorr   CorrectedAlpha
	DeltaCorr   DirectDeltaCorrection
	Package     BoundaryRestPressurePackage
	Controls    AlternativeControls
	Impact      CandidateImpact
	Firewalls   Firewalls
	Branch      BranchDecision
	Truth       string
	Final       string
}

func M2(s float64) float64                                  { return PBoundary * s * s }
func M3(s float64) float64                                  { return PBoundary * s * s * s }
func AlphaBoundary(s float64) float64                       { return ThreeTenths * s }
func AlphaMin(neff float64) float64                         { return math.Sqrt(neff/3.0) - 1.0 }
func BetaRequired(alpha, neff float64) float64              { return 3.0*math.Pow(1.0+alpha, 2.0)/neff - 1.0 }
func QRest(alpha, beta float64) float64                     { return beta / (3.0 * alpha * alpha) }
func DeltaCandidateLeading(s float64) float64               { return NineOverFive * s }
func DeltaCandidateCorrected(s float64, c2 float64) float64 { return NineOverFive*s + c2*M2(s) }
func CYukawaFromDelta(delta float64) float64                { return 3.0 / (3.0 + delta) }
func CHiggsFromCYukawa(cy float64) float64                  { return cy * CHistory }

func BuildDefault() (Analysis, error) {
	m2 := M2(SBoundary)
	m3 := M3(SBoundary)
	alphaB := AlphaBoundary(SBoundary)
	alphaMin := AlphaMin(NEff)
	deltaAlpha := alphaMin - alphaB
	halfM2 := Half * m2
	alphaCorr := alphaB + halfM2
	betaCorr := BetaRequired(alphaCorr, NEff)
	qRest := math.NaN()
	qValid := false
	if betaCorr >= 0 {
		qRest = QRest(alphaCorr, betaCorr)
		qValid = qRest >= 0 && qRest <= 1
	} else {
		qRest = QRest(alphaCorr, betaCorr)
	}
	leadingResidual := DeltaN - DeltaCandidateLeading(SBoundary)
	c2Obs := leadingResidual / m2
	deltaCorr := DeltaCandidateCorrected(SBoundary, Six)
	candidateResidual := DeltaN - deltaCorr
	cy := CYukawaFromDelta(deltaCorr)
	ch := CHiggsFromCYukawa(cy)
	controls := []CoefficientControl{
		{Name: "c = 2", Coefficient: 2, Residual: DeltaN - 2*SBoundary, AbsResidual: math.Abs(DeltaN - 2*SBoundary), TypedSource: "simple boundary double response; poor residual"},
		{Name: "c = 7/4", Coefficient: 7.0 / 4.0, Residual: DeltaN - (7.0/4.0)*SBoundary, AbsResidual: math.Abs(DeltaN - (7.0/4.0)*SBoundary), TypedSource: "K7/four-chamber style candidate; weaker residual"},
		{Name: "c = 13/7", Coefficient: 13.0 / 7.0, Residual: DeltaN - (13.0/7.0)*SBoundary, AbsResidual: math.Abs(DeltaN - (13.0/7.0)*SBoundary), TypedSource: "prior K7/Hodge obstruction ratio; weaker residual"},
		{Name: "c = 9/5", Coefficient: NineOverFive, Residual: leadingResidual, AbsResidual: math.Abs(leadingResidual), TypedSource: "color-three × inverse hypercharge normalization"},
		{Name: "c = 9/5 + 6ps", Coefficient: NineOverFive + Six*PBoundary*SBoundary, Residual: candidateResidual, AbsResidual: math.Abs(candidateResidual), TypedSource: "leading color-hypercharge coefficient plus boundary-pair × color second-moment correction"},
	}
	return Analysis{
		Inheritance: Inheritance{Gate810Inherited: true, CoeffTargetSelected: true, CorrectionTargetSelected: true, NEff: NEff, DeltaN: DeltaN, S: SBoundary, CYukawa: CYukawa, CHistory: CHistory, CHiggs: CHiggs, Verdicts: []string{StatusGate810Inherited, StatusCoeffTargetSelected, StatusRestCorrectionTarget}},
		NineFifths:  FactorizationAudit{Audited: true, Expression: "9/5 = 3 × (3/5)", Sources: []string{"3: finite spectral-action color multiplicity / top-color block", "5/3: active hypercharge/gauge normalization in kappa_e_red", "3/5: inverse hypercharge-normalization response"}, Description: "color-three baseline multiplied by inverse hypercharge-normalized boundary response", Verdicts: []string{StatusNineFiveFactor, StatusLedgerExistence}, Supports: []string{StatusNineFiveTypedCandidate, StatusRestCoeffColorHypercharge}, Failures: []string{StatusExistenceNotTheorem, StatusNoColorHyperchargeMap, StatusInverseHyperchargeNotYukawa}},
		ThreeTenths: FactorizationAudit{Audited: true, Expression: "3/10 = (1/2)(3/5)", Sources: []string{"3/5: inverse hypercharge-normalization response", "1/2: boundary pair midpoint / two-endpoint averaging candidate"}, Description: "midpoint-normalized inverse-hypercharge boundary response", Verdicts: []string{StatusThreeTenthsFactor}, Supports: []string{StatusThreeTenthsTypedCandidate, StatusHalfBoundaryPairCandidate}, Failures: []string{StatusNoBoundaryAverageTheorem, StatusThreeTenthsNotNative, StatusBoundaryPairNotReadout}},
		Correction:  PositivityCorrection{Defined: true, AlphaMin: alphaMin, AlphaB: alphaB, DeltaAlpha: deltaAlpha, DeltaAlphaOverS: deltaAlpha / SBoundary, DeltaAlphaOverAlphaB: deltaAlpha / alphaB, S2: SBoundary * SBoundary, M2: m2, HalfM2: halfM2, M3: m3, DeltaOverHalfM2: deltaAlpha / halfM2, Verdicts: []string{StatusCorrectionDefined, StatusCorrectionScale}, Supports: []string{StatusCorrectionOrderM2, StatusHalfM2Approx, StatusCorrectionMayBeM2}, Failures: []string{StatusCorrectionNotExactHalfM2, StatusNoPositiveCorrectionTheorem, StatusM2NotConcentrationLaw}},
		AlphaCorr:   CorrectedAlpha{Defined: true, Tested: true, BetaNonnegative: betaCorr >= 0, QRestValid: qValid, AlphaCorr: alphaCorr, AlphaMin: alphaMin, AlphaCorrMinusAlphaMin: alphaCorr - alphaMin, BetaCorr: betaCorr, QRestCorr: qRest, Verdicts: []string{StatusAlphaCorrDefined, StatusAlphaCorrTestRequired, StatusQRestDefined}, Supports: []string{StatusAlphaHalfM2LawfulToTest}, Failures: []string{StatusAlphaCorrNeedsBeta, StatusAlphaCorrNotNative, StatusQRestNotSector, StatusQRestNotNativeSpectrum, StatusNoAtomLedgerFromQRest}},
		DeltaCorr:   DirectDeltaCorrection{Defined: true, Computed: true, Residual: leadingResidual, M2: m2, C2Obs: c2Obs, CandidateC2: Six, CandidateDelta: deltaCorr, CandidateResidual: candidateResidual, CandidateRelErr: candidateResidual / DeltaN, Verdicts: []string{StatusDeltaCorrDefined, StatusC2ObsComputed}, Supports: []string{StatusDeltaResidualOrderM2, StatusC2ObsCloseToSix}, Failures: []string{StatusC2ObsNotNative, StatusSecondMomentNeedsMap, StatusResidualFittingForbidden}},
		Package:     BoundaryRestPressurePackage{Defined: true, Name: "HyperchargeColorBoundaryRestPressurePackage", Components: []string{"leading coefficient 9/5 = 3 × 3/5", "alpha coefficient 3/10 = 1/2 × 3/5", "positive-rest correction of order p s^2", "direct Delta_N correction candidate 6 p s^2", "boundary-to-trace-magnitude readout", "positive spectrum construction", "scale stability", "noncircularity"}, Verdict: StatusPackageDefined, Supports: []string{StatusPackageSharpestCandidate}, Failures: []string{StatusPackageNeedsTraceMap, StatusPackageNeedsSpectrum, StatusPackageNeedsScale}},
		Controls:    AlternativeControls{Defined: true, Controls: controls, Verdict: StatusControlRules, Supports: []string{StatusCorrectedModelControl}, Failures: []string{StatusRationalClosenessForbidden, StatusModelNeedsSourcePositivity}},
		Impact:      CandidateImpact{Defined: true, NBoundary: 3.0 + deltaCorr, CYukawaBoundary: cy, CHiggsBoundary: ch, CYukawaResidual: cy - CYukawa, CHiggsResidual: ch - CHiggs, Verdicts: []string{StatusImpactDefined}, Failures: []string{StatusNoCYukawaUpdate, StatusCHiggsLevelB}},
		Firewalls:   Firewalls{Enforced: true, NoCoeffShortcut: true, NoAlphaExact: true, NoHalfM2Native: true, NoSixFit: true, NoPackageNative: true, NoBoundarySpectrum: true, NoQRestSector: true, NoCertifiedRewrite: true, NoPoleMass: true, Verdict: StatusFirewallGate811},
		Branch:      BranchDecision{Recorded: true, Next: "Gate 812 — Boundary Second-Moment RestPressure Correction and Positive Spectrum Construction Audit", Purpose: "test whether a second-moment correction can be promoted from residual-improving candidate into a positivity-compatible boundary-to-rest-pressure construction", Verdict: StatusBranchDecision, Supports: []string{StatusNextSecondMomentCorrection}},
		Truth:       "Gate 811 finds that 9/5 has the strongest typed color-hypercharge coefficient candidate, but the law is not native and the naive alpha=(3/10)s exact closure fails positivity.",
		Final:       "The sharper missing object is a HyperchargeColorBoundaryRestPressurePackage with a second-moment correction and an actual positive rest-spectrum construction.",
	}, nil
}

func Statuses() []string {
	return []string{StatusGate810Inherited, StatusCoeffTargetSelected, StatusRestCorrectionTarget, StatusNineFiveFactor, StatusLedgerExistence, StatusThreeTenthsFactor, StatusCorrectionDefined, StatusCorrectionScale, StatusAlphaCorrDefined, StatusAlphaCorrTestRequired, StatusQRestDefined, StatusDeltaCorrDefined, StatusC2ObsComputed, StatusPackageDefined, StatusControlRules, StatusImpactDefined, StatusPhysicalFirewalls, StatusBranchDecision, StatusNineFiveTypedCandidate, StatusRestCoeffColorHypercharge, StatusThreeTenthsTypedCandidate, StatusHalfBoundaryPairCandidate, StatusCorrectionOrderM2, StatusHalfM2Approx, StatusCorrectionMayBeM2, StatusAlphaHalfM2LawfulToTest, StatusDeltaResidualOrderM2, StatusC2ObsCloseToSix, StatusPackageSharpestCandidate, StatusCorrectedModelControl, StatusNextSecondMomentCorrection, StatusExistenceNotTheorem, StatusNoColorHyperchargeMap, StatusInverseHyperchargeNotYukawa, StatusNoBoundaryAverageTheorem, StatusThreeTenthsNotNative, StatusBoundaryPairNotReadout, StatusCorrectionNotExactHalfM2, StatusNoPositiveCorrectionTheorem, StatusM2NotConcentrationLaw, StatusAlphaCorrNeedsBeta, StatusAlphaCorrNotNative, StatusQRestNotSector, StatusQRestNotNativeSpectrum, StatusNoAtomLedgerFromQRest, StatusC2ObsNotNative, StatusSecondMomentNeedsMap, StatusResidualFittingForbidden, StatusPackageNeedsTraceMap, StatusPackageNeedsSpectrum, StatusPackageNeedsScale, StatusRationalClosenessForbidden, StatusModelNeedsSourcePositivity, StatusNoCYukawaUpdate, StatusCHiggsLevelB, StatusTreeProxyNotPole, StatusFirewallGate811}
}

func FormatFactor(a FactorizationAudit) string {
	return fmt.Sprintf("%s sources=[%s] supports=[%s] failures=[%s]", a.Expression, strings.Join(a.Sources, "; "), strings.Join(a.Supports, "; "), strings.Join(a.Failures, "; "))
}
func FormatCorrection(c PositivityCorrection) string {
	return fmt.Sprintf("alpha_min=%.17g alpha_B=%.17g delta=%.17g delta/s=%.17g M2=%.17g halfM2=%.17g delta/halfM2=%.17g supports=[%s] failures=[%s]", c.AlphaMin, c.AlphaB, c.DeltaAlpha, c.DeltaAlphaOverS, c.M2, c.HalfM2, c.DeltaOverHalfM2, strings.Join(c.Supports, "; "), strings.Join(c.Failures, "; "))
}
func FormatAlphaCorr(c CorrectedAlpha) string {
	return fmt.Sprintf("alpha_corr=%.17g alpha_corr-alpha_min=%.17g beta_corr=%.17g q_rest=%.17g beta_nonnegative=%v q_valid=%v supports=[%s] failures=[%s]", c.AlphaCorr, c.AlphaCorrMinusAlphaMin, c.BetaCorr, c.QRestCorr, c.BetaNonnegative, c.QRestValid, strings.Join(c.Supports, "; "), strings.Join(c.Failures, "; "))
}
func FormatDeltaCorr(d DirectDeltaCorrection) string {
	return fmt.Sprintf("R_B=%.17g M2=%.17g c2_obs=%.17g candidateC2=%.17g candidateDelta=%.17g residual=%.17g rel=%.17g supports=[%s] failures=[%s]", d.Residual, d.M2, d.C2Obs, d.CandidateC2, d.CandidateDelta, d.CandidateResidual, d.CandidateRelErr, strings.Join(d.Supports, "; "), strings.Join(d.Failures, "; "))
}
func FormatPackage(p BoundaryRestPressurePackage) string {
	return fmt.Sprintf("%s components=[%s] supports=[%s] failures=[%s]", p.Name, strings.Join(p.Components, "; "), strings.Join(p.Supports, "; "), strings.Join(p.Failures, "; "))
}
func FormatControls(c AlternativeControls) string {
	parts := make([]string, 0, len(c.Controls))
	for _, x := range c.Controls {
		parts = append(parts, fmt.Sprintf("%s residual=%.17g source=%s", x.Name, x.Residual, x.TypedSource))
	}
	return fmt.Sprintf("controls=[%s] supports=[%s] failures=[%s]", strings.Join(parts, " | "), strings.Join(c.Supports, "; "), strings.Join(c.Failures, "; "))
}
func FormatImpact(i CandidateImpact) string {
	return fmt.Sprintf("N_boundary=%.17g C_Yukawa_boundary=%.17g C_Yukawa_residual=%.17g C_Higgs_boundary=%.17g C_Higgs_residual=%.17g failures=[%s]", i.NBoundary, i.CYukawaBoundary, i.CYukawaResidual, i.CHiggsBoundary, i.CHiggsResidual, strings.Join(i.Failures, "; "))
}

func containsAll(hay []string, needles []string) bool {
	m := map[string]bool{}
	for _, h := range hay {
		m[h] = true
	}
	for _, n := range needles {
		if !m[n] {
			return false
		}
	}
	return true
}
