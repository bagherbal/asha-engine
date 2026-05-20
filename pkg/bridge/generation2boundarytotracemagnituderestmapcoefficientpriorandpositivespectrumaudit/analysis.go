// Package generation2boundarytotracemagnituderestmapcoefficientpriorandpositivespectrumaudit implements
// Gate 816: BoundaryToTraceMagnitudeRestMap Coefficient-Prior and Positive-Spectrum Construction Audit.
//
// Gate 816 tests whether the frozen boundary-FN scalar closure from Gate 815 can
// be upgraded into a prior-sourced alpha/beta/q_rest positive top/rest model. It
// deliberately separates coefficient-prior evidence from a true trace-magnitude
// map, and preserves that no Yukawa atoms or sector spectra are constructed.
package generation2boundarytotracemagnituderestmapcoefficientpriorandpositivespectrumaudit

import (
	"fmt"
	"math"
	"sort"
	"strings"
)

const (
	AuditID = "GATE816-BOUNDARY-TO-TRACE-MAGNITUDE-REST-MAP-COEFFICIENT-PRIOR-POSITIVE-SPECTRUM-AUDIT"

	NEff      = 3.0023273474722147
	DeltaN    = NEff - 3.0
	SBoundary = 0.0012924448188162962
	PBoundary = 7.0 / 72.0
	CHistory  = 1.038025177923625
	CYukawa   = 0.9992248188812008
	CHiggs    = 1.0372205204048603

	C1NineFive = 9.0 / 5.0
	C2Six      = 6.0

	StatusGate815Inherited     = "PASS_GATE815_BOUNDARY_FN_TEST_PROTOCOL_INHERITED"
	StatusNumericalLedger      = "PASS_NUMERICAL_LEDGER_RECOMPUTED"
	StatusCoeffPrior95         = "PASS_NINE_OVER_FIVE_COEFFICIENT_PRIOR_AUDITED"
	StatusCoeffPrior6          = "PASS_SIX_SECOND_MOMENT_COEFFICIENT_PRIOR_AUDITED"
	StatusScalarVsConstruction = "PASS_SCALAR_CLOSURE_VERSUS_TOP_REST_CONSTRUCTION_AUDITED"
	StatusAlphaCandidates      = "PASS_ALPHA_CANDIDATE_TABLE_COMPUTED"
	StatusBetaQPositivity      = "PASS_BETA_Q_REST_POSITIVITY_TABLE_COMPUTED"
	StatusDeltaRecon           = "PASS_DELTA_N_RECONSTRUCTION_TABLE_COMPUTED"
	StatusPositiveSpectrum     = "PASS_POSITIVE_SPECTRUM_REALIZABILITY_AUDITED"
	StatusCoeffPriorNoGo       = "PASS_COEFFICIENT_PRIOR_NO_GO_TEST_COMPLETED"
	StatusNonCircularity       = "PASS_NONCIRCULARITY_REQUIREMENTS_ENFORCED"
	StatusStatusLevel          = "PASS_BOUNDARY_FN_STATUS_LEVEL_CLASSIFIED"
	StatusImpactRecorded       = "PASS_C_YUKAWA_AND_C_HIGGS_IMPACT_RECORDED"
	StatusBranchRecorded       = "PASS_BRANCH_DECISION_RECORDED"
	StatusPhysicalFirewalls    = "PASS_PHYSICAL_FIREWALLS_ENFORCED"

	StatusClosureReady       = "CONDITIONAL_SUPPORT_BOUNDARY_FN_CLOSURE_REMAINS_SHARP_NUMERICALLY"
	StatusCoeff95Prior       = "CONDITIONAL_SUPPORT_NINE_OVER_FIVE_HAS_COLOR_THREE_TIMES_INVERSE_HYPERCHARGE_PRIOR"
	StatusCoeff6Prior        = "CONDITIONAL_SUPPORT_SIX_HAS_BOUNDARY_PAIR_DIMENSION_TIMES_COLOR_PRIOR"
	StatusCoeffPriorBridge   = "CONDITIONAL_SUPPORT_COEFFICIENTS_ARE_BRIDGE_LAYER_PRIORS_NOT_NATIVE_THEOREMS"
	StatusAlpha3_5Positive   = "CONDITIONAL_SUPPORT_C_ALPHA_THREE_OVER_FIVE_IS_POSITIVE_COMPATIBLE_FOR_INHERITED_AND_BFN_N_EFF"
	StatusAlpha6_11Positive  = "CONDITIONAL_SUPPORT_C_ALPHA_SIX_OVER_ELEVEN_IS_POSITIVE_COMPATIBLE_FOR_INHERITED_AND_BFN_N_EFF"
	StatusAlphaOnePositive   = "CONDITIONAL_SUPPORT_C_ALPHA_ONE_IS_POSITIVE_COMPATIBLE_AND_HAS_Q_NEAR_ONE_THIRD_FOR_BFN"
	StatusHalfAlmostBoundary = "CONDITIONAL_SUPPORT_HALF_M2_ALPHA_CORRECTION_IS_CLOSE_BUT_REMAINS_BELOW_POSITIVITY_BOUNDARY"
	StatusQRestFamily        = "CONDITIONAL_SUPPORT_Q_REST_CANDIDATES_CLASSIFY_POSSIBLE_REST_CONCENTRATIONS"
	StatusAbstractPositive   = "CONDITIONAL_SUPPORT_ABSTRACT_POSITIVE_REST_SPECTRA_EXIST_FOR_POSITIVE_COMPATIBLE_ROWS"
	StatusNoMapConstructed   = "CONDITIONAL_SUPPORT_BOUNDARY_TO_TRACE_MAGNITUDE_REST_MAP_REMAINS_MISSING"
	StatusLevelPartialR2     = "CONDITIONAL_SUPPORT_BOUNDARY_FN_STATUS_IS_R1_WITH_PARTIAL_R2_POSITIVE_COMPATIBILITY"
	StatusNoLedgerUpdate     = "CONDITIONAL_SUPPORT_CANDIDATE_VALUES_REMAIN_UNOFFICIAL_UNTIL_MAP_IS_CERTIFIED"
	StatusNextNoGo           = "CONDITIONAL_SUPPORT_NEXT_GATE_SHOULD_AUDIT_BOUNDARY_FN_COEFFICIENT_PRIOR_NO_GO_OR_EXTERNAL_LEDGER_TEST"

	StatusCoeff95NotTheorem      = "FAILED_ROUTE_EXISTENCE_OF_COLOR_THREE_AND_FIVE_OVER_THREE_DOES_NOT_PROVE_NINE_OVER_FIVE"
	StatusCoeff6NotTheorem       = "FAILED_ROUTE_SIX_MUST_NOT_BE_ACCEPTED_BY_ROUNDING_C2_OBS"
	StatusScalarNotTraceMap      = "FAILED_ROUTE_SCALAR_CLOSURE_DOES_NOT_CONSTRUCT_ALPHA_BETA_Q_REST_BY_ITSELF"
	StatusHalfM2NegativeBeta     = "FAILED_ROUTE_C_ALPHA_ONE_HALF_STILL_GIVES_NEGATIVE_BETA"
	StatusNoNativeCAlpha         = "FAILED_ROUTE_NO_NATIVE_C_ALPHA_COEFFICIENT_THEOREM"
	StatusQRestChosenNotTheorem  = "FAILED_ROUTE_Q_REST_CHOSEN_TO_FORCE_CLOSURE_IS_NOT_NATIVE_REST_LAW"
	StatusCoeffPackageNoMap      = "FAILED_ROUTE_COEFFICIENT_PRIOR_PACKAGE_DOES_NOT_CONSTRUCT_ALPHA_BETA_Q_REST"
	StatusPositiveNoSectors      = "FAILED_ROUTE_ABSTRACT_POSITIVE_SPECTRUM_DOES_NOT_ASSIGN_SECTORS"
	StatusNoTraceAtoms           = "FAILED_ROUTE_NO_TRACE_ATOM_CONSTRUCTION_SUPPLIED"
	StatusNoYukawaOperator       = "FAILED_ROUTE_NO_NATIVE_YUKAWA_OPERATOR_THEOREM"
	StatusNoRetune               = "FAILED_ROUTE_COEFFICIENT_RETUNING_REMAINS_FORBIDDEN"
	StatusNoTopSolve             = "FAILED_ROUTE_TOP_REST_VARIABLES_MUST_NOT_BE_SOLVED_BACKWARDS_FROM_N_EFF_AS_THEOREM"
	StatusNoHiggsData            = "FAILED_ROUTE_HIGGS_OR_SCALAR_RUNTIME_DATA_MUST_NOT_SOURCE_YUKAWA_REST_PRESSURE"
	StatusNoFNCharges            = "FAILED_ROUTE_FN_LIKE_FOURTH_ROOT_NOT_FN_CHARGE_OPERATOR"
	StatusGate816NoCYukawaUpdate = "FAILED_ROUTE_GATE816_DOES_NOT_UPDATE_C_YUKAWA"
	StatusCHiggsLevelB           = "FAILED_ROUTE_C_HIGGS_REMAINS_LEVEL_B"
	StatusTreeProxyNotPole       = "FAILED_ROUTE_TREE_PROXY_NOT_POLE_MASS"
	StatusFirewallGate816        = "FIREWALL_PRESERVED_GATE816_BOUNDARY_TO_TRACE_MAGNITUDE_RESTMAP_COEFFICIENT_PRIOR_BOUNDARY"
)

type NumericalLedger struct {
	NEff, DeltaN, S, P           float64
	M1, M2, M3                   float64
	DeltaNBFN, NEffBFN           float64
	ResidualBFN, RelResidual     float64
	C2Obs, EpsilonBFN            float64
	Verdicts, Supports, Failures []string
}

type CoefficientPrior struct {
	Name                         string
	Value                        float64
	Factorization                string
	CandidateTyping              []string
	Classification               string
	Verdicts, Supports, Failures []string
}

type AlphaCandidate struct {
	Name             string
	CAlpha           float64
	Source           string
	Alpha            float64
	AgainstInherited PositivityRow
	AgainstBFN       PositivityRow
	Classification   string
}

type PositivityRow struct {
	NEffName           string
	NEff               float64
	BetaRequired       float64
	QRestRequired      float64
	PositiveCompatible bool
	BoundaryLimit      bool
}

type QRestCandidate struct {
	Name   string
	QRest  float64
	Source string
}

type DeltaReconstruction struct {
	AlphaName            string
	QRestName            string
	Alpha, QRest, DeltaN float64
	ResidualToBFN        float64
	PositiveCompatible   bool
}

type PositiveSpectrumStatus struct {
	AbstractExistenceAny         bool
	BestRows                     []string
	Level                        string
	Verdicts, Supports, Failures []string
}

type CoefficientPriorNoGo struct {
	PackageDefined               bool
	Ingredients                  []string
	ProducesScalarClosure        bool
	ConstructsAlphaBetaQ         bool
	Verdicts, Supports, Failures []string
}

type NonCircularity struct {
	Enforced           bool
	Forbidden, Allowed []string
	Verdicts, Failures []string
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
	Enforced                                                                                                        bool
	CoeffPriorNotTheorem, ScalarClosureNotMap, PositiveExistenceNotLedger, SectorLedgerNotNative, BoundaryNotYukawa bool
	FNLikeNotChargeOperator, HyperchargeNotRestLaw, ColorNotGeneration, CHiggsLevelB, TreeProxyNotPole              bool
	Verdict                                                                                                         string
}

type Analysis struct {
	Ledger          NumericalLedger
	Coeff95         CoefficientPrior
	Coeff6          CoefficientPrior
	AlphaRows       []AlphaCandidate
	QRestCandidates []QRestCandidate
	DeltaRows       []DeltaReconstruction
	Positive        PositiveSpectrumStatus
	NoGo            CoefficientPriorNoGo
	Protocol        NonCircularity
	Impact          Impact
	Branch          BranchDecision
	Firewalls       Firewalls
	Truth           string
	Final           string
}

func M1(s float64) float64                         { return PBoundary * s }
func M2(s float64) float64                         { return PBoundary * s * s }
func M3(s float64) float64                         { return PBoundary * s * s * s }
func DeltaBFN(s float64) float64                   { return C1NineFive*s + C2Six*M2(s) }
func NEffBFN(s float64) float64                    { return 3.0 + DeltaBFN(s) }
func CYukawaFromNEff(nEff float64) float64         { return 3.0 / nEff }
func CHiggsFromNEff(nEff float64) float64          { return CYukawaFromNEff(nEff) * CHistory }
func C2Observed(deltaN, s float64) float64         { return (deltaN - C1NineFive*s) / M2(s) }
func Epsilon(delta float64) float64                { return math.Pow(delta, 0.25) }
func AlphaFromC(cAlpha float64, s float64) float64 { return 0.3*s + cAlpha*M2(s) }
func BetaRequired(alpha, nEff float64) float64     { return 3.0*math.Pow(1.0+alpha, 2)/nEff - 1.0 }
func QRestRequired(alpha, beta float64) float64 {
	if alpha == 0 {
		return math.NaN()
	}
	return beta / (3.0 * alpha * alpha)
}
func DeltaFromAlphaQ(alpha, qRest float64) float64 {
	beta := 3.0 * alpha * alpha * qRest
	return 3.0*math.Pow(1+alpha, 2)/(1+beta) - 3.0
}
func PositiveCompatible(beta, q float64) bool { return beta >= 0 && q >= 0 && q <= 1 }
func BoundaryLimit(beta float64) bool         { return math.Abs(beta) < 1e-10 }

func BuildDefault() (Analysis, error) {
	m1, m2, m3 := M1(SBoundary), M2(SBoundary), M3(SBoundary)
	deltaBFN := DeltaBFN(SBoundary)
	nEffBFN := 3.0 + deltaBFN
	residual := DeltaN - deltaBFN
	ledger := NumericalLedger{NEff: NEff, DeltaN: DeltaN, S: SBoundary, P: PBoundary, M1: m1, M2: m2, M3: m3, DeltaNBFN: deltaBFN, NEffBFN: nEffBFN, ResidualBFN: residual, RelResidual: residual / DeltaN, C2Obs: C2Observed(DeltaN, SBoundary), EpsilonBFN: Epsilon(deltaBFN), Verdicts: []string{StatusGate815Inherited, StatusNumericalLedger}, Supports: []string{StatusClosureReady}, Failures: []string{StatusGate816NoCYukawaUpdate}}

	coeff95 := CoefficientPrior{Name: "9/5", Value: C1NineFive, Factorization: "9/5 = 3 × 3/5", CandidateTyping: []string{"3: color multiplicity / top-color baseline", "5/3: active hypercharge normalization", "3/5: inverse hypercharge normalization", "9/5: color-three times inverse-hypercharge-normalized boundary split"}, Classification: "B — bridge-layer coefficient prior, not native theorem", Verdicts: []string{StatusCoeffPrior95}, Supports: []string{StatusCoeff95Prior, StatusCoeffPriorBridge}, Failures: []string{StatusCoeff95NotTheorem}}
	coeff6 := CoefficientPrior{Name: "6", Value: C2Six, Factorization: "6 = 2 × 3", CandidateTyping: []string{"2: boundary-pair dimension / two wall endpoints", "3: color multiplicity / top-color baseline", "6: boundary-pair dimension times color multiplicity multiplying p s^2"}, Classification: "B — bridge-layer second-moment coefficient prior, not native theorem", Verdicts: []string{StatusCoeffPrior6}, Supports: []string{StatusCoeff6Prior, StatusCoeffPriorBridge}, Failures: []string{StatusCoeff6NotTheorem}}

	cands := []struct {
		name   string
		c      float64
		source string
	}{
		{"1/2", 0.5, "half second raw boundary moment"},
		{"3/5", 0.6, "inverse hypercharge correction"},
		{"6/11", 6.0 / 11.0, "possible normalized boundary/color correction"},
		{"1", 1.0, "full second raw boundary moment"},
	}
	alphaRows := make([]AlphaCandidate, 0, len(cands))
	for _, c := range cands {
		alpha := AlphaFromC(c.c, SBoundary)
		inh := makePositivityRow("inherited", NEff, alpha)
		bfn := makePositivityRow("BFN", nEffBFN, alpha)
		class := "failed"
		if inh.PositiveCompatible && bfn.PositiveCompatible {
			class = "positive-compatible"
		} else if inh.BoundaryLimit || bfn.BoundaryLimit {
			class = "boundary-limit"
		}
		if class == "positive-compatible" && (c.name == "6/11" || c.name == "3/5" || c.name == "1") {
			// Positive-compatible does not imply native coefficient source.
			class += "; untyped until coefficient theorem exists"
		}
		alphaRows = append(alphaRows, AlphaCandidate{Name: c.name, CAlpha: c.c, Source: c.source, Alpha: alpha, AgainstInherited: inh, AgainstBFN: bfn, Classification: class})
	}

	qCands := []QRestCandidate{{"0", 0, "beta-zero diffuse boundary limit"}, {"1/3", 1.0 / 3.0, "three comparable rest-pressure atoms"}, {"1/2", 0.5, "two comparable rest-pressure atoms / boundary-pair candidate"}, {"1", 1, "single concentrated rest atom"}, {"p", PBoundary, "K7 event-weight concentration candidate"}, {"3/5", 0.6, "inverse hypercharge candidate"}}
	deltaRows := make([]DeltaReconstruction, 0, len(alphaRows)*len(qCands))
	for _, ar := range alphaRows {
		for _, q := range qCands {
			delta := DeltaFromAlphaQ(ar.Alpha, q.QRest)
			deltaRows = append(deltaRows, DeltaReconstruction{AlphaName: ar.Name, QRestName: q.Name, Alpha: ar.Alpha, QRest: q.QRest, DeltaN: delta, ResidualToBFN: delta - deltaBFN, PositiveCompatible: q.QRest >= 0 && q.QRest <= 1})
		}
	}
	sort.Slice(deltaRows, func(i, j int) bool {
		return math.Abs(deltaRows[i].ResidualToBFN) < math.Abs(deltaRows[j].ResidualToBFN)
	})

	positiveRows := []string{}
	for _, ar := range alphaRows {
		if ar.AgainstBFN.PositiveCompatible {
			positiveRows = append(positiveRows, fmt.Sprintf("c_alpha=%s q_BFN=%.16g", ar.Name, ar.AgainstBFN.QRestRequired))
		}
	}
	positive := PositiveSpectrumStatus{AbstractExistenceAny: len(positiveRows) > 0, BestRows: positiveRows, Level: "R1 scalar closure with partial R2 positive top/rest compatibility; not R3/R4", Verdicts: []string{StatusPositiveSpectrum, StatusStatusLevel}, Supports: []string{StatusAbstractPositive, StatusNoMapConstructed, StatusLevelPartialR2, StatusAlpha3_5Positive, StatusAlpha6_11Positive, StatusAlphaOnePositive, StatusHalfAlmostBoundary, StatusQRestFamily}, Failures: []string{StatusHalfM2NegativeBeta, StatusNoNativeCAlpha, StatusPositiveNoSectors, StatusNoTraceAtoms, StatusNoYukawaOperator}}

	noGo := CoefficientPriorNoGo{PackageDefined: true, Ingredients: []string{"color factor 3", "inverse hypercharge factor 3/5", "boundary-pair factor 2", "K7 event weight p", "boundary split s", "second raw moment p s^2", "top-color baseline", "positive-rest correction rule"}, ProducesScalarClosure: true, ConstructsAlphaBetaQ: false, Verdicts: []string{StatusCoeffPriorNoGo, StatusScalarVsConstruction}, Supports: []string{StatusCoeffPriorBridge}, Failures: []string{StatusCoeffPackageNoMap, StatusScalarNotTraceMap}}
	protocol := NonCircularity{Enforced: true, Forbidden: []string{"derive 9/5 by minimizing residual", "derive 6 by rounding c2_obs", "choose c_alpha to force beta>=0", "choose q_rest to force closure and call it native", "solve T,alpha,beta backward from N_eff as theorem", "use Higgs/C_Higgs/runtime data", "use Koide/FN/GJ patterns to invent rest atoms"}, Allowed: []string{"test fixed coefficient priors", "compute residuals", "classify positivity", "construct abstract positive spectra", "state missing maps", "recommend empirical validation protocol"}, Verdicts: []string{StatusNonCircularity}, Failures: []string{StatusNoRetune, StatusNoTopSolve, StatusNoHiggsData, StatusQRestChosenNotTheorem, StatusNoFNCharges}}
	impact := Impact{Recorded: true, NEffBFN: nEffBFN, CYukawaBFN: CYukawaFromNEff(nEffBFN), CHiggsBFN: CHiggsFromNEff(nEffBFN), OfficialNEff: NEff, OfficialCYukawa: CYukawa, OfficialCHiggs: CHiggs, Verdicts: []string{StatusImpactRecorded}, Supports: []string{StatusNoLedgerUpdate}, Failures: []string{StatusGate816NoCYukawaUpdate, StatusCHiggsLevelB, StatusTreeProxyNotPole}}
	branch := BranchDecision{Recorded: true, Outcome: "Outcome B — partial success", NextGate: "Gate 817 — BoundaryToTraceMagnitudeRestMap Coefficient-Prior No-Go and External Ledger Validation Branch Audit", Verdicts: []string{StatusBranchRecorded}, Supports: []string{StatusNextNoGo}}
	firewalls := Firewalls{Enforced: true, CoeffPriorNotTheorem: true, ScalarClosureNotMap: true, PositiveExistenceNotLedger: true, SectorLedgerNotNative: true, BoundaryNotYukawa: true, FNLikeNotChargeOperator: true, HyperchargeNotRestLaw: true, ColorNotGeneration: true, CHiggsLevelB: true, TreeProxyNotPole: true, Verdict: StatusFirewallGate816}

	return Analysis{Ledger: ledger, Coeff95: coeff95, Coeff6: coeff6, AlphaRows: alphaRows, QRestCandidates: qCands, DeltaRows: deltaRows, Positive: positive, NoGo: noGo, Protocol: protocol, Impact: impact, Branch: branch, Firewalls: firewalls, Truth: "Gate 816 finds typed coefficient priors and abstract positive-spectrum compatibility, but no BoundaryToTraceMagnitudeRestMap.", Final: "Boundary-FN remains a Level-B+ R1/partial-R2 hypothesis until a non-fit alpha/beta/q map or external trace atoms are supplied."}, nil
}

func makePositivityRow(name string, nEff, alpha float64) PositivityRow {
	beta := BetaRequired(alpha, nEff)
	q := QRestRequired(alpha, beta)
	return PositivityRow{NEffName: name, NEff: nEff, BetaRequired: beta, QRestRequired: q, PositiveCompatible: PositiveCompatible(beta, q), BoundaryLimit: BoundaryLimit(beta)}
}

func Statuses() []string {
	return []string{StatusGate815Inherited, StatusNumericalLedger, StatusCoeffPrior95, StatusCoeffPrior6, StatusScalarVsConstruction, StatusAlphaCandidates, StatusBetaQPositivity, StatusDeltaRecon, StatusPositiveSpectrum, StatusCoeffPriorNoGo, StatusNonCircularity, StatusStatusLevel, StatusImpactRecorded, StatusBranchRecorded, StatusPhysicalFirewalls, StatusClosureReady, StatusCoeff95Prior, StatusCoeff6Prior, StatusCoeffPriorBridge, StatusAlpha3_5Positive, StatusAlpha6_11Positive, StatusAlphaOnePositive, StatusHalfAlmostBoundary, StatusQRestFamily, StatusAbstractPositive, StatusNoMapConstructed, StatusLevelPartialR2, StatusNoLedgerUpdate, StatusNextNoGo, StatusCoeff95NotTheorem, StatusCoeff6NotTheorem, StatusScalarNotTraceMap, StatusHalfM2NegativeBeta, StatusNoNativeCAlpha, StatusQRestChosenNotTheorem, StatusCoeffPackageNoMap, StatusPositiveNoSectors, StatusNoTraceAtoms, StatusNoYukawaOperator, StatusNoRetune, StatusNoTopSolve, StatusNoHiggsData, StatusNoFNCharges, StatusGate816NoCYukawaUpdate, StatusCHiggsLevelB, StatusTreeProxyNotPole, StatusFirewallGate816}
}

func FormatLedger(a NumericalLedger) string {
	return fmt.Sprintf("N_eff=%.16g Delta_N=%.16g s=%.16g p=%.16g M1=%.16g M2=%.16g M3=%.16g Delta_BFN=%.16g R=%.16g rho=%.16g c2_obs=%.16g epsilon_BFN=%.16g", a.NEff, a.DeltaN, a.S, a.P, a.M1, a.M2, a.M3, a.DeltaNBFN, a.ResidualBFN, a.RelResidual, a.C2Obs, a.EpsilonBFN)
}

func FormatCoefficient(a CoefficientPrior) string {
	return fmt.Sprintf("%s value=%.16g %s classification=%s typing=[%s]", a.Name, a.Value, a.Factorization, a.Classification, strings.Join(a.CandidateTyping, "; "))
}

func FormatAlphaRows(rows []AlphaCandidate) string {
	parts := make([]string, 0, len(rows))
	for _, r := range rows {
		parts = append(parts, fmt.Sprintf("c=%s alpha=%.16g inherited(beta=%.16g,q=%.16g,pos=%t) BFN(beta=%.16g,q=%.16g,pos=%t) %s", r.Name, r.Alpha, r.AgainstInherited.BetaRequired, r.AgainstInherited.QRestRequired, r.AgainstInherited.PositiveCompatible, r.AgainstBFN.BetaRequired, r.AgainstBFN.QRestRequired, r.AgainstBFN.PositiveCompatible, r.Classification))
	}
	return strings.Join(parts, " | ")
}

func FormatDeltaRows(rows []DeltaReconstruction, limit int) string {
	if limit > len(rows) {
		limit = len(rows)
	}
	parts := make([]string, 0, limit)
	for i := 0; i < limit; i++ {
		r := rows[i]
		parts = append(parts, fmt.Sprintf("alpha=%s q=%s Delta=%.16g R_BFN=%.16g", r.AlphaName, r.QRestName, r.DeltaN, r.ResidualToBFN))
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
