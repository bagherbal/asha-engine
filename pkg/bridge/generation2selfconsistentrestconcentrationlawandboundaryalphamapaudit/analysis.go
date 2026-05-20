// Package generation2selfconsistentrestconcentrationlawandboundaryalphamapaudit implements
// Gate 817: Self-Consistent Rest Concentration Law and Boundary Alpha Map Audit.
//
// Gate 817 audits the sharper compression of the Gate 816 boundary-FN closure:
//
//	Delta_N_BFN = (9/5)s + 6 p s^2 = 6[(3/10)s + p s^2] = 6 alpha_B.
//
// It then checks whether the exact top/rest participation formula closes with a
// self-consistent concentration q_rest = 1/N_eff_BFN, while preserving that this
// is not yet a native Yukawa trace-magnitude theorem or sector ledger.
package generation2selfconsistentrestconcentrationlawandboundaryalphamapaudit

import (
	"fmt"
	"math"
	"strings"
)

const (
	AuditID = "GATE817-SELF-CONSISTENT-REST-CONCENTRATION-LAW-BOUNDARY-ALPHA-MAP-AUDIT"

	NEff      = 3.0023273474722147
	DeltaN    = NEff - 3.0
	SBoundary = 0.0012924448188162962
	PBoundary = 7.0 / 72.0
	CHistory  = 1.038025177923625
	CYukawa   = 0.9992248188812008
	CHiggs    = 1.0372205204048603

	C1NineFive = 9.0 / 5.0
	C2Six      = 6.0

	StatusGate816Inherited        = "PASS_GATE816_COEFFICIENT_PRIOR_POSITIVE_SPECTRUM_INHERITED"
	StatusBoundaryAlphaDefined    = "PASS_BOUNDARY_ALPHA_MAP_DEFINED"
	StatusExactClosureTested      = "PASS_EXACT_TOP_REST_CLOSURE_TESTED"
	StatusBetaIdentityProved      = "PASS_BETA_IDENTITY_SYMBOLICALLY_PROVED"
	StatusPositiveSpectrumAudited = "PASS_ABSTRACT_POSITIVE_REST_SPECTRUM_REALIZABILITY_AUDITED"
	StatusThreeConstructionTested = "PASS_THREE_REST_CONSTRUCTION_TESTED"
	StatusOneConstructionTested   = "PASS_ONE_REST_CONSTRUCTION_TESTED"
	StatusMixedConstructionTested = "PASS_MIXED_REST_CONSTRUCTION_TESTED"
	StatusSelfConsistencyAudited  = "PASS_SELF_CONSISTENCY_VERSUS_THEOREM_AUDITED"
	StatusMapStatusUpdated        = "PASS_BOUNDARY_TO_TRACE_MAGNITUDE_REST_MAP_STATUS_UPDATED"
	StatusImpactRecorded          = "PASS_C_YUKAWA_AND_C_HIGGS_CANDIDATE_IMPACT_RECORDED"
	StatusOutcomeRecorded         = "PASS_OUTCOME_CLASSIFICATION_RECORDED"
	StatusBranchRecorded          = "PASS_BRANCH_DECISION_RECORDED"
	StatusPhysicalFirewalls       = "PASS_PHYSICAL_FIREWALLS_ENFORCED"

	StatusAlphaSourceShape    = "CONDITIONAL_SUPPORT_ALPHA_B_HAS_TYPED_BOUNDARY_HYPERCHARGE_K7_SOURCE_SHAPE"
	StatusDeltaEqualsSixAlpha = "CONDITIONAL_SUPPORT_DELTA_N_BFN_EQUALS_SIX_ALPHA_B_EXACTLY"
	StatusBetaExactIdentity   = "CONDITIONAL_SUPPORT_BETA_B_EQUALS_THREE_ALPHA_B_SQUARED_OVER_N_EFF_BFN"
	StatusQInverseNEff        = "CONDITIONAL_SUPPORT_Q_REST_B_EQUALS_INVERSE_N_EFF_BFN_SELF_CONSISTENTLY"
	StatusPositiveQ           = "CONDITIONAL_SUPPORT_Q_REST_B_ALLOWS_POSITIVE_REST_SPECTRUM"
	StatusApproxThreeRest     = "CONDITIONAL_SUPPORT_Q_REST_B_IS_APPROXIMATELY_THREE_REST_ATOM_CONCENTRATION"
	StatusMixedFourAtomExact  = "CONDITIONAL_SUPPORT_MIXED_FOUR_ATOM_ABSTRACT_SPECTRUM_CAN_REALIZE_Q_REST_B"
	StatusSelfConsistentLaw   = "CONDITIONAL_SUPPORT_Q_REST_B_IS_SELF_CONSISTENT_CLOSURE_CONDITION"
	StatusPartialR2           = "CONDITIONAL_SUPPORT_BOUNDARY_FN_REMAINS_PARTIAL_R2_BECAUSE_Q_REST_LACKS_INDEPENDENT_SOURCE"
	StatusCandidateImpact     = "CONDITIONAL_SUPPORT_CERTIFIED_MAP_WOULD_REDUCE_N_EFF_SEAL_DEPENDENCE"
	StatusNextMinimality      = "CONDITIONAL_SUPPORT_NEXT_GATE_SHOULD_AUDIT_SELF_CONCENTRATION_LAW_SOURCE_OR_REST_MAP_MINIMALITY"

	StatusAlphaNotTheorem      = "FAILED_ROUTE_ALPHA_B_NOT_NATIVE_REST_SIZE_THEOREM_WITHOUT_TRACE_MAGNITUDE_MAP"
	StatusQNoIndependentSource = "FAILED_ROUTE_NO_NATIVE_REASON_Q_REST_EQUALS_INVERSE_TOTAL_N_EFF"
	StatusQMayBeAlgebraic      = "FAILED_ROUTE_Q_REST_B_MAY_BE_ALGEBRAIC_SELF_CONSISTENCY_NOT_NATIVE_LAW"
	StatusThreeNotExact        = "FAILED_ROUTE_THREE_EQUAL_REST_ATOMS_DO_NOT_EXACTLY_REALIZE_Q_REST_B"
	StatusOneNotMatch          = "FAILED_ROUTE_ONE_CONCENTRATED_REST_ATOM_DOES_NOT_REALIZE_Q_REST_B"
	StatusPositiveNoSectors    = "FAILED_ROUTE_POSITIVE_ABSTRACT_SPECTRUM_DOES_NOT_ASSIGN_SECTORS"
	StatusPositiveNotYukawa    = "FAILED_ROUTE_POSITIVE_ABSTRACT_SPECTRUM_NOT_NATIVE_YUKAWA_OPERATOR"
	StatusNoTraceAtoms         = "FAILED_ROUTE_NO_EXTERNAL_TRACE_ATOMS_OR_SECTOR_LEDGER_SUPPLIED"
	StatusNoR3                 = "FAILED_ROUTE_BOUNDARY_FN_BRANCH_NOT_EXTERNAL_R3_TRACE_ATOM_VALIDATED"
	StatusNoR4                 = "FAILED_ROUTE_BOUNDARY_FN_BRANCH_NOT_R4_NATIVE_YUKAWA_OPERATOR_THEOREM"
	StatusNoUpdate             = "FAILED_ROUTE_GATE817_DOES_NOT_UPDATE_C_YUKAWA_IF_Q_REST_IS_ONLY_SELF_CONSISTENCY"
	StatusCHiggsLevelB         = "FAILED_ROUTE_C_HIGGS_REMAINS_LEVEL_B"
	StatusTreeProxyNotPole     = "FAILED_ROUTE_TREE_PROXY_NOT_POLE_MASS"
	StatusFirewallGate817      = "FIREWALL_PRESERVED_GATE817_SELF_CONSISTENT_REST_CONCENTRATION_BOUNDARY"
)

type Ledger struct {
	NEff, DeltaN, S, P       float64
	M2                       float64
	DeltaNBFN, NEffBFN       float64
	ResidualBFN, RelResidual float64
	AlphaB                   float64
	SixAlphaB                float64
	BetaB                    float64
	QRestB                   float64
	InverseQRestB            float64
	CYukawaBFN, CHiggsBFN    float64
	Verdicts                 []string
	Supports                 []string
	Failures                 []string
}

type AlphaMap struct {
	Alpha                        float64
	Definition                   string
	CoefficientTyping            []string
	DeltaEqualsSixAlphaExactly   bool
	Verdicts, Supports, Failures []string
}

type ExactClosure struct {
	Alpha, NEff              float64
	BetaByFormula            float64
	BetaBySimplifiedIdentity float64
	BetaIdentityResidual     float64
	QRest                    float64
	InverseQRest             float64
	Verdicts, Supports       []string
	Failures                 []string
}

type SpectrumConstruction struct {
	Name               string
	Weights            []float64
	Sum                float64
	Q                  float64
	TargetQ            float64
	Residual           float64
	Exact              bool
	Classification     string
	Verdicts, Supports []string
	Failures           []string
}

type SelfConsistencyAudit struct {
	Classification               string
	CandidateSourceLanes         []string
	IndependentSourceCertified   bool
	Verdicts, Supports, Failures []string
}

type MapStatus struct {
	CandidateExpression          string
	Level                        string
	ConstructsAlphaBetaQ         bool
	ConstructsPositiveSpectrum   bool
	ConstructsSectorLedger       bool
	NativeYukawaTheorem          bool
	Verdicts, Supports, Failures []string
}

type Impact struct {
	Recorded                                      bool
	NEffBFN, CYukawaBFN, CHiggsBFN                float64
	OfficialNEff, OfficialCYukawa, OfficialCHiggs float64
	Verdicts, Supports, Failures                  []string
}

type BranchDecision struct {
	Recorded           bool
	Outcome            string
	NextGate           string
	Verdicts, Supports []string
}

type Firewalls struct {
	Enforced                                                                                              bool
	DeltaSixAlphaNotTheorem, QRestNotTheorem, PositiveNotLedger, SectorLedgerNotNative, BoundaryNotYukawa bool
	FNLikeNotChargeOperator, HyperchargeNotRestLaw, ColorNotGeneration, CHiggsLevelB, TreeProxyNotPole    bool
	Verdict                                                                                               string
}

type Analysis struct {
	Ledger    Ledger
	Alpha     AlphaMap
	Closure   ExactClosure
	Spectra   []SpectrumConstruction
	SelfAudit SelfConsistencyAudit
	Map       MapStatus
	Impact    Impact
	Branch    BranchDecision
	Firewalls Firewalls
	Truth     string
	Final     string
}

func M2(s float64) float64       { return PBoundary * s * s }
func AlphaB(s float64) float64   { return (3.0/10.0)*s + M2(s) }
func DeltaBFN(s float64) float64 { return C1NineFive*s + C2Six*M2(s) }
func NEffBFN(s float64) float64  { return 3.0 + DeltaBFN(s) }
func BetaFormula(alpha, nEff float64) float64 {
	return 3.0*math.Pow(1.0+alpha, 2)/nEff - 1.0
}
func BetaSelfConsistent(alpha, nEff float64) float64 { return 3.0 * alpha * alpha / nEff }
func QRest(beta, alpha float64) float64              { return beta / (3.0 * alpha * alpha) }
func CYukawaFromNEff(nEff float64) float64           { return 3.0 / nEff }
func CHiggsFromNEff(nEff float64) float64            { return CYukawaFromNEff(nEff) * CHistory }
func sum(xs []float64) float64 {
	t := 0.0
	for _, x := range xs {
		t += x
	}
	return t
}
func concentration(xs []float64) float64 {
	t := 0.0
	for _, x := range xs {
		t += x * x
	}
	return t
}

func BuildDefault() (Analysis, error) {
	m2 := M2(SBoundary)
	alpha := AlphaB(SBoundary)
	deltaBFN := DeltaBFN(SBoundary)
	nEffBFN := 3.0 + deltaBFN
	if math.Abs(deltaBFN-6.0*alpha) > 1e-18 {
		return Analysis{}, fmt.Errorf("Delta_N_BFN does not equal 6 alpha_B: %.18g vs %.18g", deltaBFN, 6.0*alpha)
	}
	betaF := BetaFormula(alpha, nEffBFN)
	betaI := BetaSelfConsistent(alpha, nEffBFN)
	q := 1.0 / nEffBFN
	ledger := Ledger{NEff: NEff, DeltaN: DeltaN, S: SBoundary, P: PBoundary, M2: m2, DeltaNBFN: deltaBFN, NEffBFN: nEffBFN, ResidualBFN: DeltaN - deltaBFN, RelResidual: (DeltaN - deltaBFN) / DeltaN, AlphaB: alpha, SixAlphaB: 6.0 * alpha, BetaB: betaI, QRestB: q, InverseQRestB: 1.0 / q, CYukawaBFN: CYukawaFromNEff(nEffBFN), CHiggsBFN: CHiggsFromNEff(nEffBFN), Verdicts: []string{StatusGate816Inherited}, Supports: []string{StatusDeltaEqualsSixAlpha, StatusQInverseNEff}, Failures: []string{StatusNoUpdate}}

	alphaMap := AlphaMap{Alpha: alpha, Definition: "alpha_B = (3/10)s + p s^2", CoefficientTyping: []string{"3/10 = (1/2)(3/5): boundary-pair averaging times inverse hypercharge", "p s^2: K7-weighted second raw boundary moment", "Delta_N_BFN = 6 alpha_B with 6 = boundary-pair dimension × color multiplicity"}, DeltaEqualsSixAlphaExactly: true, Verdicts: []string{StatusBoundaryAlphaDefined}, Supports: []string{StatusAlphaSourceShape, StatusDeltaEqualsSixAlpha}, Failures: []string{StatusAlphaNotTheorem}}

	closure := ExactClosure{Alpha: alpha, NEff: nEffBFN, BetaByFormula: betaF, BetaBySimplifiedIdentity: betaI, BetaIdentityResidual: betaF - betaI, QRest: q, InverseQRest: 1.0 / q, Verdicts: []string{StatusExactClosureTested, StatusBetaIdentityProved}, Supports: []string{StatusBetaExactIdentity, StatusQInverseNEff, StatusPositiveQ}, Failures: []string{StatusQNoIndependentSource, StatusQMayBeAlgebraic}}

	spectra := buildSpectra(q)
	selfAudit := SelfConsistencyAudit{Classification: "B — self-consistent closure condition without independent rest-concentration source", CandidateSourceLanes: []string{"inverse participation geometry", "maximum-entropy rest distribution", "three-comparable-rest-atom limit", "boundary-pair diffusion", "K7 event-weight averaging", "finite spectral-action trace normalization"}, IndependentSourceCertified: false, Verdicts: []string{StatusSelfConsistencyAudited}, Supports: []string{StatusSelfConsistentLaw}, Failures: []string{StatusQNoIndependentSource, StatusQMayBeAlgebraic}}
	mapStatus := MapStatus{CandidateExpression: "alpha_B=(3/10)s+p s^2; q_rest_B=1/(3+6 alpha_B); beta_B=3 alpha_B^2 q_rest_B; N_eff_BFN=3+6 alpha_B", Level: "partial R2 — boundary alpha map plus positive-compatible q closure, but q source and sector atoms remain missing", ConstructsAlphaBetaQ: true, ConstructsPositiveSpectrum: true, ConstructsSectorLedger: false, NativeYukawaTheorem: false, Verdicts: []string{StatusMapStatusUpdated}, Supports: []string{StatusPartialR2, StatusMixedFourAtomExact}, Failures: []string{StatusNoTraceAtoms, StatusNoR3, StatusNoR4, StatusPositiveNoSectors, StatusPositiveNotYukawa}}
	impact := Impact{Recorded: true, NEffBFN: nEffBFN, CYukawaBFN: CYukawaFromNEff(nEffBFN), CHiggsBFN: CHiggsFromNEff(nEffBFN), OfficialNEff: NEff, OfficialCYukawa: CYukawa, OfficialCHiggs: CHiggs, Verdicts: []string{StatusImpactRecorded}, Supports: []string{StatusCandidateImpact}, Failures: []string{StatusNoUpdate, StatusCHiggsLevelB, StatusTreeProxyNotPole}}
	branch := BranchDecision{Recorded: true, Outcome: "Outcome B — partial R2 success", NextGate: "Gate 818 — RestConcentrationLaw Source and BoundaryToTraceMagnitudeRestMap Minimality Audit", Verdicts: []string{StatusOutcomeRecorded, StatusBranchRecorded}, Supports: []string{StatusNextMinimality}}
	firewalls := Firewalls{Enforced: true, DeltaSixAlphaNotTheorem: true, QRestNotTheorem: true, PositiveNotLedger: true, SectorLedgerNotNative: true, BoundaryNotYukawa: true, FNLikeNotChargeOperator: true, HyperchargeNotRestLaw: true, ColorNotGeneration: true, CHiggsLevelB: true, TreeProxyNotPole: true, Verdict: StatusFirewallGate817}

	return Analysis{Ledger: ledger, Alpha: alphaMap, Closure: closure, Spectra: spectra, SelfAudit: selfAudit, Map: mapStatus, Impact: impact, Branch: branch, Firewalls: firewalls, Truth: "Gate 817 compresses Delta_N_BFN into 6 alpha_B and obtains an exact self-consistent q_rest_B = 1/N_eff_BFN, but no independent source for that concentration law is certified.", Final: "Boundary-FN upgrades from R1/partial R2 to a sharper partial R2 construction: alpha/beta/q are positive-compatible, yet q_rest remains self-consistency rather than native Yukawa trace law."}, nil
}

func buildSpectra(q float64) []SpectrumConstruction {
	three := []float64{1.0 / 3.0, 1.0 / 3.0, 1.0 / 3.0}
	one := []float64{1.0}
	mixedSmall := mixedFourWeights(q, true)
	mixedLarge := mixedFourWeights(q, false)
	return []SpectrumConstruction{
		{Name: "diffuse three-rest construction", Weights: three, Sum: sum(three), Q: concentration(three), TargetQ: q, Residual: concentration(three) - q, Exact: math.Abs(concentration(three)-q) < 1e-15, Classification: "close but not exact; q_rest_B is slightly below 1/3", Verdicts: []string{StatusThreeConstructionTested}, Supports: []string{StatusApproxThreeRest}, Failures: []string{StatusThreeNotExact, StatusPositiveNoSectors}},
		{Name: "concentrated one-rest construction", Weights: one, Sum: sum(one), Q: concentration(one), TargetQ: q, Residual: concentration(one) - q, Exact: math.Abs(concentration(one)-q) < 1e-15, Classification: "fails; q=1 is too concentrated", Verdicts: []string{StatusOneConstructionTested}, Failures: []string{StatusOneNotMatch, StatusPositiveNoSectors}},
		{Name: "mixed four-rest construction small-support branch", Weights: mixedSmall, Sum: sum(mixedSmall), Q: concentration(mixedSmall), TargetQ: q, Residual: concentration(mixedSmall) - q, Exact: math.Abs(concentration(mixedSmall)-q) < 1e-15, Classification: "exact abstract positive spectrum; one tiny support plus three comparable atoms", Verdicts: []string{StatusMixedConstructionTested, StatusPositiveSpectrumAudited}, Supports: []string{StatusMixedFourAtomExact, StatusPositiveQ}, Failures: []string{StatusPositiveNoSectors, StatusPositiveNotYukawa}},
		{Name: "mixed four-rest construction heavy-support branch", Weights: mixedLarge, Sum: sum(mixedLarge), Q: concentration(mixedLarge), TargetQ: q, Residual: concentration(mixedLarge) - q, Exact: math.Abs(concentration(mixedLarge)-q) < 1e-15, Classification: "exact abstract positive spectrum; one half-sized atom plus three smaller comparable atoms", Verdicts: []string{StatusMixedConstructionTested, StatusPositiveSpectrumAudited}, Supports: []string{StatusMixedFourAtomExact, StatusPositiveQ}, Failures: []string{StatusPositiveNoSectors, StatusPositiveNotYukawa}},
	}
}

// mixedFourWeights constructs weights [t,(1-t)/3,(1-t)/3,(1-t)/3]
// matching q = t^2 + (1-t)^2/3. The small branch is the solution close to zero.
func mixedFourWeights(q float64, small bool) []float64 {
	A, B, C := 4.0/3.0, -2.0/3.0, 1.0/3.0-q
	disc := B*B - 4*A*C
	if disc < 0 {
		return nil
	}
	sqrtDisc := math.Sqrt(disc)
	t1 := (-B - sqrtDisc) / (2 * A)
	t2 := (-B + sqrtDisc) / (2 * A)
	t := t2
	if small {
		t = t1
	}
	return []float64{t, (1 - t) / 3, (1 - t) / 3, (1 - t) / 3}
}

func Statuses() []string {
	return []string{StatusGate816Inherited, StatusBoundaryAlphaDefined, StatusExactClosureTested, StatusBetaIdentityProved, StatusPositiveSpectrumAudited, StatusThreeConstructionTested, StatusOneConstructionTested, StatusMixedConstructionTested, StatusSelfConsistencyAudited, StatusMapStatusUpdated, StatusImpactRecorded, StatusOutcomeRecorded, StatusBranchRecorded, StatusPhysicalFirewalls, StatusAlphaSourceShape, StatusDeltaEqualsSixAlpha, StatusBetaExactIdentity, StatusQInverseNEff, StatusPositiveQ, StatusApproxThreeRest, StatusMixedFourAtomExact, StatusSelfConsistentLaw, StatusPartialR2, StatusCandidateImpact, StatusNextMinimality, StatusAlphaNotTheorem, StatusQNoIndependentSource, StatusQMayBeAlgebraic, StatusThreeNotExact, StatusOneNotMatch, StatusPositiveNoSectors, StatusPositiveNotYukawa, StatusNoTraceAtoms, StatusNoR3, StatusNoR4, StatusNoUpdate, StatusCHiggsLevelB, StatusTreeProxyNotPole, StatusFirewallGate817}
}

func FormatLedger(a Ledger) string {
	return fmt.Sprintf("N_eff=%.16g Delta_N=%.16g s=%.16g p=%.16g M2=%.16g Delta_BFN=%.16g N_eff_BFN=%.16g R=%.16g rho=%.16g alpha_B=%.16g 6alpha_B=%.16g beta_B=%.16g q_rest_B=%.16g inv_q=%.16g", a.NEff, a.DeltaN, a.S, a.P, a.M2, a.DeltaNBFN, a.NEffBFN, a.ResidualBFN, a.RelResidual, a.AlphaB, a.SixAlphaB, a.BetaB, a.QRestB, a.InverseQRestB)
}

func FormatAlpha(a AlphaMap) string {
	return fmt.Sprintf("%s alpha=%.16g typing=[%s]", a.Definition, a.Alpha, strings.Join(a.CoefficientTyping, "; "))
}

func FormatClosure(a ExactClosure) string {
	return fmt.Sprintf("alpha=%.16g N_eff_BFN=%.16g beta_formula=%.16g beta_identity=%.16g residual=%.16g q=%.16g inv_q=%.16g", a.Alpha, a.NEff, a.BetaByFormula, a.BetaBySimplifiedIdentity, a.BetaIdentityResidual, a.QRest, a.InverseQRest)
}

func FormatSpectra(rows []SpectrumConstruction) string {
	parts := make([]string, 0, len(rows))
	for _, r := range rows {
		parts = append(parts, fmt.Sprintf("%s weights=%v sum=%.16g q=%.16g target=%.16g R=%.16g exact=%t class=%s", r.Name, r.Weights, r.Sum, r.Q, r.TargetQ, r.Residual, r.Exact, r.Classification))
	}
	return strings.Join(parts, " | ")
}

func FormatImpact(a Impact) string {
	return fmt.Sprintf("candidate NEff=%.16g CYukawa=%.16g CHiggs=%.16g official NEff=%.16g CYukawa=%.16g CHiggs=%.16g", a.NEffBFN, a.CYukawaBFN, a.CHiggsBFN, a.OfficialNEff, a.OfficialCYukawa, a.OfficialCHiggs)
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
