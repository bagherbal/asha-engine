// Package generation2boundaryalphaoneplusthreerestsimplexandconcentrationsourceaudit implements
// Gate 818: Boundary-Alpha 1+3 Rest Simplex and Concentration Source Audit.
//
// Gate 818 tests whether Gate 817's self-consistent rest concentration
// q_rest_B = 1/N_eff_BFN can be independently sourced by a normalized 1+3 rest
// simplex rather than merely defined from the target effective count.
package generation2boundaryalphaoneplusthreerestsimplexandconcentrationsourceaudit

import (
	"fmt"
	"math"
	"strings"
)

const (
	AuditID = "GATE818-BOUNDARY-ALPHA-ONE-PLUS-THREE-REST-SIMPLEX-CONCENTRATION-SOURCE-AUDIT"

	NEff      = 3.0023273474722147
	DeltaN    = NEff - 3.0
	SBoundary = 0.0012924448188162962
	PBoundary = 7.0 / 72.0
	CHistory  = 1.038025177923625
	CYukawa   = 0.9992248188812008
	CHiggs    = 1.0372205204048603

	C1NineFive = 9.0 / 5.0
	C2Six      = 6.0

	StatusGate817Inherited         = "PASS_GATE817_SELF_CONSISTENT_REST_CONCENTRATION_INHERITED"
	StatusBoundaryAlphaSimplex     = "PASS_BOUNDARY_ALPHA_ONE_PLUS_THREE_SIMPLEX_DEFINED"
	StatusQSimplexFormula          = "PASS_Q_SIMPLEX_FORMULA_RECORDED"
	StatusPriorAlphaSimplexAudited = "PASS_PRIOR_SOURCED_T_EQUALS_ALPHA_B_SIMPLEX_AUDITED"
	StatusSymbolicResidualRecorded = "PASS_SYMBOLIC_FIFTH_ORDER_RESIDUAL_RECORDED"
	StatusTStarBranchAudited       = "PASS_EXACT_T_STAR_BRANCH_AUDITED"
	StatusThreeControlReaudited    = "PASS_THREE_EQUAL_REST_ATOMS_CONTROL_REAUDITED"
	StatusOneControlReaudited      = "PASS_ONE_REST_ATOM_CONTROL_REAUDITED"
	StatusPositiveAtomsRecorded    = "PASS_POSITIVE_REST_ATOM_CONSTRUCTION_FROM_SIMPLEX_RECORDED"
	StatusStructuralSourceAudited  = "PASS_STRUCTURAL_SOURCE_AUDIT_RECORDED"
	StatusRStatusUpdated           = "PASS_R_STATUS_UPDATED"
	StatusImpactRecorded           = "PASS_C_YUKAWA_AND_C_HIGGS_CANDIDATE_IMPACT_RECORDED"
	StatusOutcomeRecorded          = "PASS_OUTCOME_CLASSIFICATION_RECORDED"
	StatusBranchRecorded           = "PASS_BRANCH_DECISION_RECORDED"
	StatusPhysicalFirewalls        = "PASS_PHYSICAL_FIREWALLS_ENFORCED"

	StatusAlphaHasSourceShape   = "CONDITIONAL_SUPPORT_ALPHA_B_HAS_TYPED_BOUNDARY_HYPERCHARGE_K7_SOURCE_SHAPE"
	StatusSimplexSourcesQ       = "CONDITIONAL_SUPPORT_ONE_PLUS_THREE_REST_SIMPLEX_SOURCES_Q_REST_WITHOUT_DIRECT_Q_EQUALS_ONE_OVER_N_INPUT"
	StatusFifthOrderClosure     = "CONDITIONAL_SUPPORT_ALPHA_B_SIMPLEX_REPRODUCES_BFN_CLOSURE_TO_FIFTH_ORDER"
	StatusTStarExact            = "CONDITIONAL_SUPPORT_T_STAR_EXACTLY_REALIZES_Q_REST_B"
	StatusQPositive             = "CONDITIONAL_SUPPORT_Q_SIMPLEX_ALPHA_B_ALLOWS_POSITIVE_REST_SPECTRUM"
	StatusProjectiveRelevant    = "CONDITIONAL_SUPPORT_PROJECTIVE_ONE_PLUS_THREE_IS_NOW_RELEVANT_TO_REST_CONCENTRATION_SOURCE_CANDIDATE"
	StatusK7Resonance           = "CONDITIONAL_SUPPORT_K7_4_3_POLARITY_REMAINS_REST_SIMPLEX_RESONANCE_ONLY"
	StatusAbstractSpectrum      = "CONDITIONAL_SUPPORT_ONE_PLUS_THREE_SIMPLEX_SUPPLIES_ABSTRACT_POSITIVE_REST_SPECTRUM"
	StatusStrengthenedPartialR2 = "CONDITIONAL_SUPPORT_BOUNDARY_FN_REMAINS_PARTIAL_R2_BUT_STRENGTHENED_BY_PRIOR_SIMPLEX_CONCENTRATION"
	StatusCandidateImpact       = "CONDITIONAL_SUPPORT_CERTIFIED_SIMPLEX_MAP_WOULD_REDUCE_N_EFF_SEAL_DEPENDENCE"
	StatusNextMapMinimality     = "CONDITIONAL_SUPPORT_NEXT_GATE_SHOULD_AUDIT_ONE_PLUS_THREE_SIMPLEX_SOURCE_OR_EXTERNAL_LEDGER_TEST"

	StatusSimplexNotYukawa       = "FAILED_ROUTE_T_EQUALS_ALPHA_B_SIMPLEX_NOT_NATIVE_YUKAWA_TRACE_ATOM_THEOREM"
	StatusTStarMayBeTargetSolved = "FAILED_ROUTE_T_STAR_MAY_BE_TARGET_SOLVED_SELF_CONSISTENCY_WITHOUT_INDEPENDENT_SOURCE"
	StatusSquareRootNotNative    = "FAILED_ROUTE_SQUARE_ROOT_NORMALIZATION_CORRECTION_NOT_NATIVE_WITHOUT_TRACE_MAP"
	StatusThreeNotExact          = "FAILED_ROUTE_THREE_EQUAL_REST_ATOMS_DO_NOT_EXACTLY_REALIZE_Q_REST_B"
	StatusOneNotExact            = "FAILED_ROUTE_ONE_CONCENTRATED_REST_ATOM_DOES_NOT_REALIZE_Q_REST_B"
	StatusProjectiveNotTheorem   = "FAILED_ROUTE_PROJECTIVE_ONE_PLUS_THREE_NOT_YUKAWA_TRACE_MAGNITUDE_THEOREM"
	StatusK7NotTheorem           = "FAILED_ROUTE_K7_4_3_NOT_REST_TRACE_MAGNITUDE_THEOREM"
	StatusNoSectorAssignment     = "FAILED_ROUTE_ABSTRACT_REST_ATOMS_DO_NOT_ASSIGN_STANDARD_MODEL_SECTORS"
	StatusNotNativeYukawa        = "FAILED_ROUTE_ABSTRACT_REST_ATOMS_NOT_NATIVE_YUKAWA_EIGENVALUES"
	StatusNoExternalR3           = "FAILED_ROUTE_BOUNDARY_FN_BRANCH_NOT_EXTERNAL_R3_TRACE_ATOM_VALIDATED"
	StatusNoR4                   = "FAILED_ROUTE_BOUNDARY_FN_BRANCH_NOT_R4_NATIVE_YUKAWA_OPERATOR_THEOREM"
	StatusNoCertifiedMap         = "FAILED_ROUTE_GATE818_DOES_NOT_UPDATE_C_YUKAWA_WITHOUT_CERTIFIED_TRACE_MAGNITUDE_MAP"
	StatusCHiggsLevelB           = "FAILED_ROUTE_C_HIGGS_REMAINS_LEVEL_B"
	StatusTreeProxyNotPole       = "FAILED_ROUTE_TREE_PROXY_NOT_POLE_MASS"
	StatusFirewallGate818        = "FIREWALL_PRESERVED_GATE818_BOUNDARY_ALPHA_ONE_PLUS_THREE_REST_SIMPLEX_BOUNDARY"
)

type Ledger struct {
	NEff, DeltaN, S, P, M2 float64
	AlphaB                 float64
	DeltaNBFN, NEffBFN     float64
	QRestB                 float64
	ResidualBFN            float64
	CYukawaBFN, CHiggsBFN  float64
	Verdicts               []string
	Supports               []string
	Failures               []string
}

type SimplexFormula struct {
	Formula                      string
	Verdicts, Supports, Failures []string
}

type PriorAlphaBranch struct {
	T                            float64
	Weights                      []float64
	QSimplex                     float64
	QRestB                       float64
	QResidual                    float64
	BetaSimplex                  float64
	NEffSimplex                  float64
	NEffBFN                      float64
	NEffResidual                 float64
	SymbolicResidual             float64
	Verdicts, Supports, Failures []string
}

type TStarBranch struct {
	TStar                        float64
	DeltaFromAlpha               float64
	QStar                        float64
	QRestB                       float64
	QResidual                    float64
	Expansion                    string
	Verdicts, Supports, Failures []string
}

type Control struct {
	Name                         string
	Weights                      []float64
	Q                            float64
	QRestB                       float64
	Residual                     float64
	Exact                        bool
	Verdicts, Supports, Failures []string
}

type PositiveSpectrum struct {
	Construction                 string
	NormalizedWeights            []float64
	Sum                          float64
	Q                            float64
	Beta                         float64
	Realizable                   bool
	Verdicts, Supports, Failures []string
}

type StructuralAudit struct {
	CandidateLanes               []string
	IndependentNativeSource      bool
	Verdicts, Supports, Failures []string
}

type RStatus struct {
	Level                        string
	ConstructsConcentration      bool
	ConstructsPositiveSpectrum   bool
	ConstructsSectorLedger       bool
	NativeYukawaTheorem          bool
	Verdicts, Supports, Failures []string
}

type Impact struct {
	NEffCandidate, CYukawaCandidate, CHiggsCandidate float64
	OfficialNEff, OfficialCYukawa, OfficialCHiggs    float64
	Verdicts, Supports, Failures                     []string
}

type BranchDecision struct {
	Outcome, NextGate  string
	Verdicts, Supports []string
}

type Firewalls struct {
	Enforced                                                                                       bool
	SimplexNotTheorem, TStarNotSource, AbstractNotSector, SectorLedgerNotNative, BoundaryNotYukawa bool
	ProjectiveNotYukawa, K7NotYukawa, CHiggsLevelB, TreeProxyNotPole                               bool
	Verdict                                                                                        string
}

type Analysis struct {
	Ledger     Ledger
	Formula    SimplexFormula
	Prior      PriorAlphaBranch
	TStar      TStarBranch
	Controls   []Control
	Spectrum   PositiveSpectrum
	Structural StructuralAudit
	Status     RStatus
	Impact     Impact
	Branch     BranchDecision
	Firewalls  Firewalls
	Truth      string
	Final      string
}

func M2(s float64) float64                 { return PBoundary * s * s }
func AlphaB(s float64) float64             { return (3.0/10.0)*s + M2(s) }
func DeltaBFN(s float64) float64           { return C1NineFive*s + C2Six*M2(s) }
func NEffBFN(s float64) float64            { return 3.0 + DeltaBFN(s) }
func QSimplex(t float64) float64           { return t*t + math.Pow(1.0-t, 2)/3.0 }
func Beta(alpha, q float64) float64        { return 3.0 * alpha * alpha * q }
func NEffFrom(alpha, q float64) float64    { return 3.0 * math.Pow(1.0+alpha, 2) / (1.0 + Beta(alpha, q)) }
func TStar(alpha float64) float64          { return (1.0 - math.Sqrt((1.0-6.0*alpha)/(1.0+2.0*alpha))) / 4.0 }
func CYukawaFromNEff(nEff float64) float64 { return 3.0 / nEff }
func CHiggsFromNEff(nEff float64) float64  { return CYukawaFromNEff(nEff) * CHistory }
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
	delta := DeltaBFN(SBoundary)
	nEff := 3.0 + delta
	qRest := 1.0 / nEff
	if math.Abs(delta-6.0*alpha) > 1e-18 {
		return Analysis{}, fmt.Errorf("Delta_N_BFN != 6 alpha_B: %.18g vs %.18g", delta, 6*alpha)
	}
	ledger := Ledger{NEff: NEff, DeltaN: DeltaN, S: SBoundary, P: PBoundary, M2: m2, AlphaB: alpha, DeltaNBFN: delta, NEffBFN: nEff, QRestB: qRest, ResidualBFN: DeltaN - delta, CYukawaBFN: CYukawaFromNEff(nEff), CHiggsBFN: CHiggsFromNEff(nEff), Verdicts: []string{StatusGate817Inherited}, Supports: []string{StatusAlphaHasSourceShape}, Failures: []string{StatusNoCertifiedMap}}

	formula := SimplexFormula{Formula: "q_simplex(t) = t^2 + (1-t)^2/3 for w=[t,(1-t)/3,(1-t)/3,(1-t)/3]", Verdicts: []string{StatusBoundaryAlphaSimplex, StatusQSimplexFormula}, Supports: []string{StatusProjectiveRelevant}, Failures: []string{StatusSimplexNotYukawa}}

	weights := []float64{alpha, (1 - alpha) / 3, (1 - alpha) / 3, (1 - alpha) / 3}
	qPrior := QSimplex(alpha)
	nEffPrior := NEffFrom(alpha, qPrior)
	symbolicResidual := -24.0 * math.Pow(alpha, 5) / (1.0 + alpha*alpha - 2.0*math.Pow(alpha, 3) + 4.0*math.Pow(alpha, 4))
	prior := PriorAlphaBranch{T: alpha, Weights: weights, QSimplex: qPrior, QRestB: qRest, QResidual: qPrior - qRest, BetaSimplex: Beta(alpha, qPrior), NEffSimplex: nEffPrior, NEffBFN: nEff, NEffResidual: nEffPrior - nEff, SymbolicResidual: symbolicResidual, Verdicts: []string{StatusPriorAlphaSimplexAudited, StatusSymbolicResidualRecorded}, Supports: []string{StatusSimplexSourcesQ, StatusFifthOrderClosure, StatusQPositive}, Failures: []string{StatusSimplexNotYukawa}}

	ts := TStar(alpha)
	qStar := QSimplex(ts)
	tStar := TStarBranch{TStar: ts, DeltaFromAlpha: ts - alpha, QStar: qStar, QRestB: qRest, QResidual: qStar - qRest, Expansion: "t_star = alpha_B + 4 alpha_B^3 + 8 alpha_B^4 + O(alpha_B^5)", Verdicts: []string{StatusTStarBranchAudited}, Supports: []string{StatusTStarExact}, Failures: []string{StatusTStarMayBeTargetSolved, StatusSquareRootNotNative}}

	controls := []Control{
		{Name: "three equal rest atoms", Weights: []float64{0, 1.0 / 3.0, 1.0 / 3.0, 1.0 / 3.0}, Q: 1.0 / 3.0, QRestB: qRest, Residual: 1.0/3.0 - qRest, Exact: false, Verdicts: []string{StatusThreeControlReaudited}, Failures: []string{StatusThreeNotExact}},
		{Name: "one concentrated rest atom", Weights: []float64{1}, Q: 1, QRestB: qRest, Residual: 1 - qRest, Exact: false, Verdicts: []string{StatusOneControlReaudited}, Failures: []string{StatusOneNotExact}},
	}

	spectrum := PositiveSpectrum{Construction: "abstract normalized rest atoms r_i = a_rest w_i with w=[alpha_B,(1-alpha_B)/3,(1-alpha_B)/3,(1-alpha_B)/3]", NormalizedWeights: weights, Sum: sum(weights), Q: qPrior, Beta: Beta(alpha, qPrior), Realizable: qPrior > 0 && qPrior < 1, Verdicts: []string{StatusPositiveAtomsRecorded}, Supports: []string{StatusAbstractSpectrum, StatusQPositive}, Failures: []string{StatusNoSectorAssignment, StatusNotNativeYukawa}}

	structural := StructuralAudit{CandidateLanes: []string{"Projective/Fock 1+3 selector: one distinguished rest atom plus triplet rest chamber", "K7 Hodge 4|3 polarity: native resonance, not trace theorem", "Boundary alpha_B: small boundary/FN dust weight", "External Yukawa ledger: can test actual rest atoms, not native"}, IndependentNativeSource: false, Verdicts: []string{StatusStructuralSourceAudited}, Supports: []string{StatusProjectiveRelevant, StatusK7Resonance}, Failures: []string{StatusProjectiveNotTheorem, StatusK7NotTheorem}}

	status := RStatus{Level: "Outcome B — strengthened partial R2: alpha_B simplex gives an independent concentration candidate, but no native trace map or sector ledger is certified", ConstructsConcentration: true, ConstructsPositiveSpectrum: true, ConstructsSectorLedger: false, NativeYukawaTheorem: false, Verdicts: []string{StatusRStatusUpdated, StatusOutcomeRecorded}, Supports: []string{StatusStrengthenedPartialR2}, Failures: []string{StatusNoExternalR3, StatusNoR4}}

	impact := Impact{NEffCandidate: nEffPrior, CYukawaCandidate: CYukawaFromNEff(nEffPrior), CHiggsCandidate: CHiggsFromNEff(nEffPrior), OfficialNEff: NEff, OfficialCYukawa: CYukawa, OfficialCHiggs: CHiggs, Verdicts: []string{StatusImpactRecorded}, Supports: []string{StatusCandidateImpact}, Failures: []string{StatusNoCertifiedMap, StatusCHiggsLevelB, StatusTreeProxyNotPole}}

	branch := BranchDecision{Outcome: status.Level, NextGate: "Gate 819 — OnePlusThree RestSimplex Source Minimality and External Ledger Falsification Audit", Verdicts: []string{StatusBranchRecorded}, Supports: []string{StatusNextMapMinimality}}

	firewalls := Firewalls{Enforced: true, SimplexNotTheorem: true, TStarNotSource: true, AbstractNotSector: true, SectorLedgerNotNative: true, BoundaryNotYukawa: true, ProjectiveNotYukawa: true, K7NotYukawa: true, CHiggsLevelB: true, TreeProxyNotPole: true, Verdict: StatusFirewallGate818}

	truth := "Gate 818 sources q_rest through a 1+3 simplex at t=alpha_B to fifth-order accuracy, but preserves that the source is structural and not yet a native Yukawa trace map."
	final := "Gate 818 strengthens partial R2: the 1+3 rest simplex gives a prior concentration candidate without directly setting q=1/N_eff, yet projective 1+3 and K7 4|3 remain resonances until a typed trace-magnitude map or external atom ledger exists."

	return Analysis{Ledger: ledger, Formula: formula, Prior: prior, TStar: tStar, Controls: controls, Spectrum: spectrum, Structural: structural, Status: status, Impact: impact, Branch: branch, Firewalls: firewalls, Truth: truth, Final: final}, nil
}

func Statuses() []string {
	return []string{StatusGate817Inherited, StatusBoundaryAlphaSimplex, StatusQSimplexFormula, StatusPriorAlphaSimplexAudited, StatusSymbolicResidualRecorded, StatusTStarBranchAudited, StatusThreeControlReaudited, StatusOneControlReaudited, StatusPositiveAtomsRecorded, StatusStructuralSourceAudited, StatusRStatusUpdated, StatusImpactRecorded, StatusOutcomeRecorded, StatusBranchRecorded, StatusPhysicalFirewalls, StatusAlphaHasSourceShape, StatusSimplexSourcesQ, StatusFifthOrderClosure, StatusTStarExact, StatusQPositive, StatusProjectiveRelevant, StatusK7Resonance, StatusAbstractSpectrum, StatusStrengthenedPartialR2, StatusCandidateImpact, StatusNextMapMinimality, StatusSimplexNotYukawa, StatusTStarMayBeTargetSolved, StatusSquareRootNotNative, StatusThreeNotExact, StatusOneNotExact, StatusProjectiveNotTheorem, StatusK7NotTheorem, StatusNoSectorAssignment, StatusNotNativeYukawa, StatusNoExternalR3, StatusNoR4, StatusNoCertifiedMap, StatusCHiggsLevelB, StatusTreeProxyNotPole, StatusFirewallGate818}
}

func FormatLedger(a Ledger) string {
	return fmt.Sprintf("N_eff=%.16g Delta_N=%.16g s=%.16g p=%.16g M2=%.16g alpha_B=%.16g Delta_BFN=%.16g N_eff_BFN=%.16g q_rest_B=%.16g R=%.16g CYukawa_BFN=%.16g CHiggs_BFN=%.16g", a.NEff, a.DeltaN, a.S, a.P, a.M2, a.AlphaB, a.DeltaNBFN, a.NEffBFN, a.QRestB, a.ResidualBFN, a.CYukawaBFN, a.CHiggsBFN)
}

func FormatPrior(a PriorAlphaBranch) string {
	return fmt.Sprintf("t=alpha_B=%.16g weights=%v q_simplex=%.16g q_rest_B=%.16g q_R=%.16g beta=%.16g N_eff_simplex=%.16g N_eff_BFN=%.16g N_R=%.16g symbolic_R=%.16g", a.T, a.Weights, a.QSimplex, a.QRestB, a.QResidual, a.BetaSimplex, a.NEffSimplex, a.NEffBFN, a.NEffResidual, a.SymbolicResidual)
}

func FormatTStar(a TStarBranch) string {
	return fmt.Sprintf("t_star=%.16g t_star-alpha_B=%.16g q_star=%.16g q_rest_B=%.16g q_R=%.16g expansion=%s", a.TStar, a.DeltaFromAlpha, a.QStar, a.QRestB, a.QResidual, a.Expansion)
}

func FormatControls(rows []Control) string {
	out := make([]string, 0, len(rows))
	for _, r := range rows {
		out = append(out, fmt.Sprintf("%s weights=%v q=%.16g target=%.16g R=%.16g exact=%t", r.Name, r.Weights, r.Q, r.QRestB, r.Residual, r.Exact))
	}
	return strings.Join(out, " | ")
}

func FormatSpectrum(a PositiveSpectrum) string {
	return fmt.Sprintf("%s weights=%v sum=%.16g q=%.16g beta=%.16g realizable=%t", a.Construction, a.NormalizedWeights, a.Sum, a.Q, a.Beta, a.Realizable)
}

func FormatImpact(a Impact) string {
	return fmt.Sprintf("candidate NEff=%.16g CYukawa=%.16g CHiggs=%.16g official NEff=%.16g CYukawa=%.16g CHiggs=%.16g", a.NEffCandidate, a.CYukawaCandidate, a.CHiggsCandidate, a.OfficialNEff, a.OfficialCYukawa, a.OfficialCHiggs)
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
