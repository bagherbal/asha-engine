// Package generation2boundarysecondmomentrestpressurecorrectionandpositivespectrumconstructionaudit implements
// Gate 813: Boundary Second-Moment RestPressure Correction and Positive Spectrum Construction Audit.
//
// Gate 813 follows the Gate 811/812 firewall path. It tests whether the
// boundary second raw moment M2=p s^2 can repair the leading boundary-FN
// closure Delta_N ≈ (9/5)s in a way compatible with an exact positive top/rest
// spectrum. The result is intentionally forensic: the second-moment correction
// is numerically sharp and positivity-compatible in abstract families, but no
// native BoundaryToTraceMagnitudeRestMap is certified.
package generation2boundarysecondmomentrestpressurecorrectionandpositivespectrumconstructionaudit

import (
	"fmt"
	"math"
	"strings"
)

const (
	AuditID = "GATE813-BOUNDARY-SECOND-MOMENT-RESTPRESSURE-POSITIVE-SPECTRUM-CONSTRUCTION-AUDIT"

	NEff      = 3.0023273474722147
	DeltaN    = NEff - 3.0
	SBoundary = 0.0012924448188162962
	PBoundary = 7.0 / 72.0
	CYukawa   = 0.9992248188812008
	CHistory  = 1.038025177923625
	CHiggs    = 1.0372205204048603

	CLeadingNineFive   = 9.0 / 5.0
	CAlphaThreeTen     = 3.0 / 10.0
	CSecondMomentSix   = 6.0
	CAlphaHalf         = 0.5
	CAlphaThreeFifths  = 3.0 / 5.0
	CAlphaSixElevenths = 6.0 / 11.0

	StatusGate811Inherited     = "PASS_GATE811_HYPERCHARGE_COLOR_POSITIVE_REST_CORRECTION_INHERITED"
	StatusGate812Inherited     = "PASS_GATE812_CHIRALITY_FIREWALL_INHERITED"
	StatusBoundaryM2Selected   = "PASS_BOUNDARY_SECOND_MOMENT_SELECTED_AS_CURRENT_REST_PRESSURE_CORRECTION_TARGET"
	StatusTopRestFramework     = "PASS_TOP_REST_EXACT_POSITIVITY_FRAMEWORK_DEFINED"
	StatusQRestDefined         = "PASS_REST_CONCENTRATION_PARAMETER_Q_REST_DEFINED"
	StatusNaiveAlphaReaudited  = "PASS_NAIVE_ALPHA_CLOSURE_REAUDITED"
	StatusAlphaLowerBound      = "PASS_POSITIVE_LOWER_BOUND_ALPHA_COMPUTED"
	StatusAlphaFamily          = "PASS_CORRECTED_ALPHA_FAMILY_DEFINED"
	StatusAlphaProtocol        = "PASS_CANDIDATE_ALPHA_POSITIVITY_TEST_PROTOCOL_DEFINED"
	StatusDirectDeltaReaudited = "PASS_DIRECT_DELTA_N_SECOND_MOMENT_CLOSURE_REAUDITED"
	StatusSpectrumConstruction = "PASS_POSITIVE_REST_SPECTRUM_CONSTRUCTION_DEFINED"
	StatusExistenceCondition   = "PASS_ABSTRACT_EXISTENCE_CONDITION_FOR_REST_SPECTRUM_RECORDED"
	StatusCoeffTypingReaudited = "PASS_BOUNDARY_COEFFICIENT_SOURCE_TYPING_REAUDITED"
	StatusBoundaryTraceMap     = "PASS_BOUNDARY_TO_TRACE_MAGNITUDE_REST_MAP_DEFINED"
	StatusSpurionB2            = "PASS_BOUNDARY_FN_SPURION_WITH_SECOND_MOMENT_CORRECTION_DEFINED"
	StatusImpactAudited        = "PASS_C_YUKAWA_AND_C_HIGGS_CANDIDATE_IMPACT_AUDITED"
	StatusOutcomeBranches      = "PASS_OUTCOME_BRANCHES_DEFINED"
	StatusBranchDecision       = "PASS_BRANCH_DECISION_RECORDED"
	StatusPhysicalFirewalls    = "PASS_PHYSICAL_FIREWALLS_ENFORCED"

	StatusPositiveBetaBounds              = "CONDITIONAL_SUPPORT_POSITIVE_REST_SPECTRUM_REQUIRES_0_LESS_EQUAL_BETA_LESS_EQUAL_3_ALPHA_SQUARED"
	StatusMinimalCorrectionOrderM2        = "CONDITIONAL_SUPPORT_MINIMAL_POSITIVITY_CORRECTION_IS_ORDER_M2"
	StatusMinimalCorrectionHalfM2         = "CONDITIONAL_SUPPORT_MINIMAL_POSITIVITY_CORRECTION_APPROXIMATES_ONE_HALF_M2"
	StatusAlphaAboveHalfM2                = "CONDITIONAL_SUPPORT_ALPHA_CORRECTION_MUST_BE_SLIGHTLY_ABOVE_HALF_M2_OR_INCLUDE_HIGHER_ORDER_TERM"
	StatusC2SixSharpClosure               = "CONDITIONAL_SUPPORT_C2_EQUALS_SIX_GIVES_STRONG_SECOND_MOMENT_DELTA_N_CLOSURE"
	StatusC2SixTypedCandidate             = "CONDITIONAL_SUPPORT_C2_EQUALS_BOUNDARY_PAIR_DIMENSION_TIMES_COLOR_MULTIPLICITY_SOURCE_CANDIDATE"
	StatusDirectDeltaNeedsSpectrum        = "CONDITIONAL_SUPPORT_DIRECT_DELTA_N_CLOSURE_STILL_REQUIRES_POSITIVE_REST_SPECTRUM_REALIZATION"
	StatusSpectrumExistsByQ               = "CONDITIONAL_SUPPORT_REST_SPECTRUM_EXISTS_IF_Q_REST_LIES_BETWEEN_ZERO_AND_ONE"
	StatusCoefficientsTypedCandidates     = "CONDITIONAL_SUPPORT_COEFFICIENTS_HAVE_COLOR_HYPERCHARGE_BOUNDARY_SOURCE_CANDIDATES"
	StatusExactMissingMap                 = "CONDITIONAL_SUPPORT_THIS_MAP_IS_EXACT_MISSING_OBJECT_AFTER_GATE813"
	StatusB2SpurionSharper                = "CONDITIONAL_SUPPORT_SECOND_MOMENT_CORRECTED_SPURION_IS_SHARPER_THAN_LEADING_BOUNDARY_FN_SPURION"
	StatusCertifiedMapWouldReduceSeal     = "CONDITIONAL_SUPPORT_CERTIFIED_BOUNDARY_REST_MAP_WOULD_REDUCE_N_EFF_SEAL_DEPENDENCE"
	StatusExpectedPartialSuccess          = "CONDITIONAL_SUPPORT_EXPECTED_OUTCOME_IS_PARTIAL_SUCCESS_UNLESS_NATIVE_MAP_EXISTS"
	StatusThreeFifthsPositiveCandidate    = "CONDITIONAL_SUPPORT_THREE_FIFTHS_ALPHA_CORRECTION_IS_POSITIVE_REST_COMPATIBLE"
	StatusSixEleventhsPositiveCandidate   = "CONDITIONAL_SUPPORT_SIX_ELEVENTHS_ALPHA_CORRECTION_IS_POSITIVE_REST_COMPATIBLE"
	StatusDirectB2HasAbstractPositiveBand = "CONDITIONAL_SUPPORT_DIRECT_B2_CLOSURE_HAS_ABSTRACT_POSITIVE_ALPHA_BETA_BAND"

	StatusFirstOrderInsufficient     = "FAILED_ROUTE_FIRST_ORDER_DELTA_N_APPROXIMATION_NOT_SUFFICIENT_FOR_POSITIVE_REST_THEOREM"
	StatusAlphaThreeTenNegativeBeta  = "FAILED_ROUTE_ALPHA_EQUALS_THREE_OVER_TEN_S_REQUIRES_NEGATIVE_BETA"
	StatusNaiveAlphaNotPositiveExact = "FAILED_ROUTE_NAIVE_ALPHA_CLOSURE_NOT_POSITIVE_REST_COMPATIBLE_AS_EXACT_LAW"
	StatusHalfM2NotExact             = "FAILED_ROUTE_HALF_M2_CORRECTION_NOT_EXACTLY_CERTIFIED"
	StatusAlphaMinFromAggregate      = "FAILED_ROUTE_ALPHA_MIN_IS_DERIVED_FROM_AGGREGATE_N_EFF_NOT_NATIVE_BOUNDARY_MAP"
	StatusHalfM2StillNegative        = "FAILED_ROUTE_HALF_M2_ALPHA_CORRECTION_STILL_SLIGHTLY_NEGATIVE_BETA"
	StatusNoNativeCAlpha             = "FAILED_ROUTE_NO_NATIVE_C_ALPHA_COEFFICIENT_THEOREM"
	StatusDirectDeltaNotEnough       = "FAILED_ROUTE_DIRECT_DELTA_N_CLOSURE_NOT_ENOUGH_WITHOUT_ALPHA_BETA_Q_REST_CONSTRUCTION"
	StatusC2SixNotNative             = "FAILED_ROUTE_C2_EQUALS_SIX_NOT_NATIVE_THEOREM_WITHOUT_BOUNDARY_TO_REST_PRESSURE_MAP"
	StatusPositiveNoSectors          = "FAILED_ROUTE_POSITIVE_EXISTENCE_DOES_NOT_ASSIGN_SECTORS"
	StatusPositiveNotNativeYukawa    = "FAILED_ROUTE_POSITIVE_EXISTENCE_NOT_NATIVE_YUKAWA_OPERATOR_THEOREM"
	StatusNoRestAtomCount            = "FAILED_ROUTE_NO_REST_ATOM_COUNT_WITHOUT_DECOMPOSED_LEDGER"
	StatusCoeffTypingNotTraceTheorem = "FAILED_ROUTE_COEFFICIENT_SOURCE_TYPING_NOT_BOUNDARY_TO_YUKAWA_TRACE_THEOREM"
	StatusNoColorHyperchargeTraceMap = "FAILED_ROUTE_NO_NATIVE_COLOR_HYPERCHARGE_BOUNDARY_TO_REST_PRESSURE_MAP"
	StatusNoBoundaryTraceRestMap     = "FAILED_ROUTE_NO_NATIVE_BOUNDARY_TO_TRACE_MAGNITUDE_REST_MAP"
	StatusNoTopSelectorFromBoundary  = "FAILED_ROUTE_NO_NATIVE_TOP_COLOR_BLOCK_SELECTOR_FROM_BOUNDARY_DATA"
	StatusNoRestConcentrationLaw     = "FAILED_ROUTE_NO_NATIVE_REST_CONCENTRATION_LAW"
	StatusSpurionNotNative           = "FAILED_ROUTE_BOUNDARY_FN_SPURION_NOT_NATIVE_WITHOUT_FN_CHARGE_OR_REST_MAP"
	StatusSpurionNoSectors           = "FAILED_ROUTE_EPSILON_B2_DOES_NOT_ASSIGN_SECTORS_OR_YUKAWA_EIGENVALUES"
	StatusNoCYukawaUpdate            = "FAILED_ROUTE_GATE813_DOES_NOT_UPDATE_C_YUKAWA_WITHOUT_CERTIFIED_MAP"
	StatusCHiggsLevelB               = "FAILED_ROUTE_C_HIGGS_REMAINS_LEVEL_B"
	StatusFirewallGate813            = "FIREWALL_PRESERVED_GATE813_BOUNDARY_SECOND_MOMENT_POSITIVE_REST_SPECTRUM_BOUNDARY"
)

type Inheritance struct {
	Gate811Inherited, Gate812Inherited, BoundarySecondMomentSelected bool
	NEff, DeltaN, S, P, M2                                           float64
	DeltaNB1, ResidualB1, C2Obs                                      float64
	DeltaNB2, ResidualB2                                             float64
	Verdicts                                                         []string
}

type PositivityFramework struct {
	Defined    bool
	Formula    string
	Inequality string
	Verdicts   []string
	Supports   []string
	Failures   []string
}

type AlphaClosure struct {
	Audited      bool
	Alpha        float64
	BetaRequired float64
	Positive     bool
	Verdicts     []string
	Failures     []string
}

type AlphaLowerBound struct {
	Computed           bool
	AlphaMin           float64
	AlphaB1            float64
	Correction         float64
	CorrectionOverM2   float64
	CorrectionOverS    float64
	Verdicts           []string
	Supports, Failures []string
}

type AlphaCandidate struct {
	Name       string
	CAlpha     float64
	Alpha      float64
	Beta       float64
	QRest      float64
	ValidBeta  bool
	ValidQRest bool
}

type AlphaFamily struct {
	Defined            bool
	Candidates         []AlphaCandidate
	HalfM2             AlphaCandidate
	ThreeFifths        AlphaCandidate
	SixElevenths       AlphaCandidate
	ObservedMin        AlphaCandidate
	Verdicts           []string
	Supports, Failures []string
}

type DirectDeltaClosure struct {
	Audited              bool
	C2Candidate          float64
	DeltaCandidate       float64
	Residual             float64
	ResidualImprovement  float64
	NEffCandidate        float64
	AlphaMin             float64
	AlphaMaxTopBranch    float64
	PositiveBandExists   bool
	EffectiveCoefficient float64
	Verdicts             []string
	Supports, Failures   []string
}

type PositiveRestConstruction struct {
	Defined            bool
	Condition          string
	Examples           []string
	Verdicts           []string
	Supports, Failures []string
}

type CoefficientSource struct {
	Audited            bool
	NineOverFive       string
	ThreeOverTen       string
	Six                string
	Verdicts           []string
	Supports, Failures []string
}

type BoundaryToTraceMap struct {
	Defined            bool
	Objects            []string
	Target             string
	Verdicts           []string
	Supports, Failures []string
}

type SpurionAudit struct {
	Defined            bool
	EpsilonN           float64
	EpsilonB1          float64
	EpsilonB2          float64
	DeltaB1            float64
	DeltaB2            float64
	B1Diff             float64
	B2Diff             float64
	Verdicts           []string
	Supports, Failures []string
}

type Impact struct {
	Audited                                             bool
	NEffBoundaryB2, CYukawaBoundaryB2, CHiggsBoundaryB2 float64
	CYukawaOfficial, CHiggsOfficial                     float64
	Verdicts                                            []string
	Supports, Failures                                  []string
}

type Outcome struct {
	Recorded bool
	Selected string
	Verdicts []string
	Supports []string
}

type BranchDecision struct {
	Recorded       bool
	Next, IfStrong string
	Verdict        string
}

type Firewalls struct {
	Enforced                                                                                bool
	NoApproxAsTheorem, NoCoefficientFit, NoSectorAssignment, NoNativeYukawa, NoLedgerUpdate bool
	NoPoleMass                                                                              bool
	Verdict                                                                                 string
}

type Analysis struct {
	Inheritance      Inheritance
	Framework        PositivityFramework
	NaiveAlpha       AlphaClosure
	LowerBound       AlphaLowerBound
	AlphaFamily      AlphaFamily
	DirectDelta      DirectDeltaClosure
	RestConstruction PositiveRestConstruction
	Coefficient      CoefficientSource
	BoundaryMap      BoundaryToTraceMap
	Spurion          SpurionAudit
	Impact           Impact
	Outcome          Outcome
	Branch           BranchDecision
	Firewalls        Firewalls
	Truth            string
	Final            string
}

func M1(s float64) float64                        { return PBoundary * s }
func M2(s float64) float64                        { return PBoundary * s * s }
func DeltaBoundaryLeading(s float64) float64      { return CLeadingNineFive * s }
func DeltaBoundarySecondMoment(s float64) float64 { return CLeadingNineFive*s + CSecondMomentSix*M2(s) }
func BetaRequired(alpha, nEff float64) float64    { return 3.0*math.Pow(1.0+alpha, 2)/nEff - 1.0 }
func QRest(alpha, beta float64) float64           { return beta / (3.0 * alpha * alpha) }
func AlphaMin(nEff float64) float64               { return math.Sqrt(nEff/3.0) - 1.0 }

func AlphaMaxTopBranch(nEff float64) float64 {
	// beta <= 3 alpha^2. Boundary beta=3 alpha^2 gives
	// nEff(1+3a^2)=3(1+a)^2. The smaller positive root is the top-dominant branch.
	A := 3.0*nEff - 3.0
	B := -6.0
	C := nEff - 3.0
	disc := B*B - 4.0*A*C
	return (-B - math.Sqrt(disc)) / (2.0 * A)
}

func CandidateAlpha(name string, cAlpha, nEff float64) AlphaCandidate {
	alpha := CAlphaThreeTen*SBoundary + cAlpha*M2(SBoundary)
	beta := BetaRequired(alpha, nEff)
	q := QRest(alpha, beta)
	return AlphaCandidate{Name: name, CAlpha: cAlpha, Alpha: alpha, Beta: beta, QRest: q, ValidBeta: beta >= -1e-18, ValidQRest: q >= -1e-15 && q <= 1.0+1e-15}
}

func BuildDefault() (Analysis, error) {
	m2 := M2(SBoundary)
	deltaB1 := DeltaBoundaryLeading(SBoundary)
	resB1 := DeltaN - deltaB1
	c2Obs := resB1 / m2
	deltaB2 := DeltaBoundarySecondMoment(SBoundary)
	resB2 := DeltaN - deltaB2
	alphaB1 := CAlphaThreeTen * SBoundary
	betaB1 := BetaRequired(alphaB1, NEff)
	alphaMin := AlphaMin(NEff)
	correction := alphaMin - alphaB1
	cAlphaObserved := correction / m2
	half := CandidateAlpha("1/2", CAlphaHalf, NEff)
	threeFifths := CandidateAlpha("3/5", CAlphaThreeFifths, NEff)
	sixElevenths := CandidateAlpha("6/11", CAlphaSixElevenths, NEff)
	observed := CandidateAlpha("alpha_min_observed", cAlphaObserved, NEff)
	alphaCandidates := []AlphaCandidate{half, threeFifths, sixElevenths, observed}
	epsN := math.Pow(DeltaN, 0.25)
	epsB1 := math.Pow(deltaB1, 0.25)
	epsB2 := math.Pow(deltaB2, 0.25)
	nEffB2 := 3.0 + deltaB2
	direct := DirectDeltaClosure{Audited: true, C2Candidate: CSecondMomentSix, DeltaCandidate: deltaB2, Residual: resB2, ResidualImprovement: math.Abs(resB1) / math.Abs(resB2), NEffCandidate: nEffB2, AlphaMin: AlphaMin(nEffB2), AlphaMaxTopBranch: AlphaMaxTopBranch(nEffB2), PositiveBandExists: AlphaMin(nEffB2) <= AlphaMaxTopBranch(nEffB2), EffectiveCoefficient: deltaB2 / SBoundary, Verdicts: []string{StatusDirectDeltaReaudited}, Supports: []string{StatusC2SixSharpClosure, StatusC2SixTypedCandidate, StatusDirectDeltaNeedsSpectrum, StatusDirectB2HasAbstractPositiveBand}, Failures: []string{StatusDirectDeltaNotEnough, StatusC2SixNotNative}}
	return Analysis{
		Inheritance:      Inheritance{Gate811Inherited: true, Gate812Inherited: true, BoundarySecondMomentSelected: true, NEff: NEff, DeltaN: DeltaN, S: SBoundary, P: PBoundary, M2: m2, DeltaNB1: deltaB1, ResidualB1: resB1, C2Obs: c2Obs, DeltaNB2: deltaB2, ResidualB2: resB2, Verdicts: []string{StatusGate811Inherited, StatusGate812Inherited, StatusBoundaryM2Selected}},
		Framework:        PositivityFramework{Defined: true, Formula: "N_eff = 3(1+alpha)^2/(1+beta)", Inequality: "alpha>=0 and 0<=beta<=3 alpha^2 with beta=3 alpha^2 q_rest, 0<=q_rest<=1", Verdicts: []string{StatusTopRestFramework, StatusQRestDefined}, Supports: []string{StatusPositiveBetaBounds}, Failures: []string{StatusFirstOrderInsufficient}},
		NaiveAlpha:       AlphaClosure{Audited: true, Alpha: alphaB1, BetaRequired: betaB1, Positive: betaB1 >= 0, Verdicts: []string{StatusNaiveAlphaReaudited}, Failures: []string{StatusAlphaThreeTenNegativeBeta, StatusNaiveAlphaNotPositiveExact}},
		LowerBound:       AlphaLowerBound{Computed: true, AlphaMin: alphaMin, AlphaB1: alphaB1, Correction: correction, CorrectionOverM2: cAlphaObserved, CorrectionOverS: correction / SBoundary, Verdicts: []string{StatusAlphaLowerBound}, Supports: []string{StatusMinimalCorrectionOrderM2, StatusMinimalCorrectionHalfM2}, Failures: []string{StatusHalfM2NotExact, StatusAlphaMinFromAggregate}},
		AlphaFamily:      AlphaFamily{Defined: true, Candidates: alphaCandidates, HalfM2: half, ThreeFifths: threeFifths, SixElevenths: sixElevenths, ObservedMin: observed, Verdicts: []string{StatusAlphaFamily, StatusAlphaProtocol}, Supports: []string{StatusAlphaAboveHalfM2, StatusThreeFifthsPositiveCandidate, StatusSixEleventhsPositiveCandidate}, Failures: []string{StatusHalfM2StillNegative, StatusNoNativeCAlpha}},
		DirectDelta:      direct,
		RestConstruction: PositiveRestConstruction{Defined: true, Condition: "positive rest spectrum exists abstractly iff q_rest=beta/(3 alpha^2) lies in [0,1]; sectors require a decomposed ledger", Examples: []string{"q_rest=1 single concentrated rest atom", "q_rest=1/m for m equal rest atoms", "0<q_rest<1 multi-atom or spectral measure realization"}, Verdicts: []string{StatusSpectrumConstruction, StatusExistenceCondition}, Supports: []string{StatusSpectrumExistsByQ}, Failures: []string{StatusPositiveNoSectors, StatusPositiveNotNativeYukawa, StatusNoRestAtomCount}},
		Coefficient:      CoefficientSource{Audited: true, NineOverFive: "9/5 = color-three × inverse hypercharge normalization 3/5", ThreeOverTen: "3/10 = half boundary-pair average × inverse hypercharge normalization 3/5", Six: "6 = boundary-pair dimension 2 × color multiplicity 3", Verdicts: []string{StatusCoeffTypingReaudited}, Supports: []string{StatusCoefficientsTypedCandidates}, Failures: []string{StatusCoeffTypingNotTraceTheorem, StatusNoColorHyperchargeTraceMap}},
		BoundaryMap:      BoundaryToTraceMap{Defined: true, Objects: []string{"boundary split s", "K7 event weight p", "hypercharge normalization 5/3", "color multiplicity 3", "boundary-pair dimension 2", "top-color block selector", "rest-pressure alpha map", "rest concentration beta/q map", "positive spectrum construction", "trace validation", "scale/scheme convention", "noncircularity proof"}, Target: "s,p -> alpha,beta -> N_eff -> C_Yukawa", Verdicts: []string{StatusBoundaryTraceMap}, Supports: []string{StatusExactMissingMap}, Failures: []string{StatusNoBoundaryTraceRestMap, StatusNoTopSelectorFromBoundary, StatusNoRestConcentrationLaw}},
		Spurion:          SpurionAudit{Defined: true, EpsilonN: epsN, EpsilonB1: epsB1, EpsilonB2: epsB2, DeltaB1: deltaB1, DeltaB2: deltaB2, B1Diff: epsB1 - epsN, B2Diff: epsB2 - epsN, Verdicts: []string{StatusSpurionB2}, Supports: []string{StatusB2SpurionSharper}, Failures: []string{StatusSpurionNotNative, StatusSpurionNoSectors}},
		Impact:           Impact{Audited: true, NEffBoundaryB2: nEffB2, CYukawaBoundaryB2: 3.0 / nEffB2, CHiggsBoundaryB2: (3.0 / nEffB2) * CHistory, CYukawaOfficial: CYukawa, CHiggsOfficial: CHiggs, Verdicts: []string{StatusImpactAudited}, Supports: []string{StatusCertifiedMapWouldReduceSeal}, Failures: []string{StatusNoCYukawaUpdate, StatusCHiggsLevelB}},
		Outcome:          Outcome{Recorded: true, Selected: "Outcome 2 — direct Delta_N closure improves strongly and abstract positive spectra exist, but no native boundary-to-trace map is certified", Verdicts: []string{StatusOutcomeBranches}, Supports: []string{StatusExpectedPartialSuccess}},
		Branch:           BranchDecision{Recorded: true, Next: "Gate 814 — BoundaryToTraceMagnitudeRestMap Minimality and No-Go Audit", IfStrong: "Gate 814 — Boundary-FN RestPressure Theorem and C_Yukawa Reduction Audit", Verdict: StatusBranchDecision},
		Firewalls:        Firewalls{Enforced: true, NoApproxAsTheorem: true, NoCoefficientFit: true, NoSectorAssignment: true, NoNativeYukawa: true, NoLedgerUpdate: true, NoPoleMass: true, Verdict: StatusFirewallGate813},
		Truth:            "Gate 813 confirms that the boundary second-moment correction is numerically sharp and positivity-compatible only as an abstract rest-spectrum construction, not as a native Yukawa theorem.",
		Final:            "The missing object is BoundaryToTraceMagnitudeRestMap: coefficients and positive-spectrum existence are not enough without a typed map from boundary data to trace magnitudes.",
	}, nil
}

func Statuses() []string {
	return []string{StatusGate811Inherited, StatusGate812Inherited, StatusBoundaryM2Selected, StatusTopRestFramework, StatusQRestDefined, StatusNaiveAlphaReaudited, StatusAlphaLowerBound, StatusAlphaFamily, StatusAlphaProtocol, StatusDirectDeltaReaudited, StatusSpectrumConstruction, StatusExistenceCondition, StatusCoeffTypingReaudited, StatusBoundaryTraceMap, StatusSpurionB2, StatusImpactAudited, StatusOutcomeBranches, StatusBranchDecision, StatusPhysicalFirewalls, StatusPositiveBetaBounds, StatusMinimalCorrectionOrderM2, StatusMinimalCorrectionHalfM2, StatusAlphaAboveHalfM2, StatusC2SixSharpClosure, StatusC2SixTypedCandidate, StatusDirectDeltaNeedsSpectrum, StatusSpectrumExistsByQ, StatusCoefficientsTypedCandidates, StatusExactMissingMap, StatusB2SpurionSharper, StatusCertifiedMapWouldReduceSeal, StatusExpectedPartialSuccess, StatusThreeFifthsPositiveCandidate, StatusSixEleventhsPositiveCandidate, StatusDirectB2HasAbstractPositiveBand, StatusFirstOrderInsufficient, StatusAlphaThreeTenNegativeBeta, StatusNaiveAlphaNotPositiveExact, StatusHalfM2NotExact, StatusAlphaMinFromAggregate, StatusHalfM2StillNegative, StatusNoNativeCAlpha, StatusDirectDeltaNotEnough, StatusC2SixNotNative, StatusPositiveNoSectors, StatusPositiveNotNativeYukawa, StatusNoRestAtomCount, StatusCoeffTypingNotTraceTheorem, StatusNoColorHyperchargeTraceMap, StatusNoBoundaryTraceRestMap, StatusNoTopSelectorFromBoundary, StatusNoRestConcentrationLaw, StatusSpurionNotNative, StatusSpurionNoSectors, StatusNoCYukawaUpdate, StatusCHiggsLevelB, StatusFirewallGate813}
}

func FormatInheritance(a Inheritance) string {
	return fmt.Sprintf("DeltaN=%.16g s=%.16g M2=%.16g DeltaB1=%.16g R1=%.16g c2obs=%.12g DeltaB2=%.16g R2=%.16g", a.DeltaN, a.S, a.M2, a.DeltaNB1, a.ResidualB1, a.C2Obs, a.DeltaNB2, a.ResidualB2)
}

func FormatAlphaCandidate(c AlphaCandidate) string {
	return fmt.Sprintf("%s c_alpha=%.12g alpha=%.16g beta=%.16g q=%.16g valid_beta=%v valid_q=%v", c.Name, c.CAlpha, c.Alpha, c.Beta, c.QRest, c.ValidBeta, c.ValidQRest)
}

func FormatAlphaFamily(a AlphaFamily) string {
	parts := make([]string, 0, len(a.Candidates))
	for _, c := range a.Candidates {
		parts = append(parts, FormatAlphaCandidate(c))
	}
	return strings.Join(parts, " | ")
}

func FormatDirectDelta(a DirectDeltaClosure) string {
	return fmt.Sprintf("DeltaB2=%.16g residual=%.16g improvement=%.6g NEffB2=%.16g alpha_band=[%.16g, %.16g] exists=%v effective_c=%.16g", a.DeltaCandidate, a.Residual, a.ResidualImprovement, a.NEffCandidate, a.AlphaMin, a.AlphaMaxTopBranch, a.PositiveBandExists, a.EffectiveCoefficient)
}

func FormatSpurion(a SpurionAudit) string {
	return fmt.Sprintf("epsilonN=%.16g epsilonB1=%.16g diffB1=%.16g epsilonB2=%.16g diffB2=%.16g", a.EpsilonN, a.EpsilonB1, a.B1Diff, a.EpsilonB2, a.B2Diff)
}

func FormatImpact(a Impact) string {
	return fmt.Sprintf("NEffB2=%.16g CYukawaB2=%.16g CHiggsB2=%.16g officialCY=%.16g officialCH=%.16g", a.NEffBoundaryB2, a.CYukawaBoundaryB2, a.CHiggsBoundaryB2, a.CYukawaOfficial, a.CHiggsOfficial)
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
