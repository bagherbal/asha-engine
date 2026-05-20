// Package generation2tracemagnitudeoperatorsealandneffsourceminimalityaudit implements
// Gate 807: TraceMagnitudeOperatorSeal and N_eff Source Audit.
//
// Gate 807 sharpens the Yukawa magnitude side of the scalar-Higgs bridge:
// N_eff needs only positive Hermitian trace spectra H_f=Y_f†Y_f, while the
// flavor-orientation term needs sector-frame misalignment and phases.
package generation2tracemagnitudeoperatorsealandneffsourceminimalityaudit

import (
	"fmt"
	"math"
	"strings"
)

const (
	AuditID = "GATE807-TRACE-MAGNITUDE-OPERATOR-SEAL-AND-N-EFF-SOURCE-AUDIT"

	NEff      = 3.0023273474722147
	CYukawa   = 0.9992248188812008
	CHistory  = 1.038025177923625
	CHiggs    = 1.0372205204048603
	NEffDelta = NEff - 3.0

	StatusGate806Inherited      = "PASS_GATE806_GENERATION_OPERATOR_MINIMALITY_INHERITED"
	StatusNEffSubproblem        = "PASS_N_EFF_SELECTED_AS_TRACE_MAGNITUDE_SUBPROBLEM"
	StatusUnitLeverageInherited = "PASS_N_EFF_UNIT_RELATIVE_LEVERAGE_INHERITED"
	StatusSealDefined           = "PASS_TRACE_MAGNITUDE_OPERATOR_SEAL_DEFINED"
	StatusTraceFormulas         = "PASS_TRACE_MAGNITUDE_FORMULAS_RECORDED"
	StatusParticipationIdentity = "PASS_INVERSE_PARTICIPATION_IDENTITY_RECORDED"
	StatusOrientationAudited    = "PASS_ORIENTATION_INVISIBILITY_AUDITED"
	StatusRankThreeBlock        = "PASS_RANK_THREE_TOP_COLOR_BLOCK_DEFINED"
	StatusTopLimit              = "PASS_TOP_COLOR_LIMIT_RECONFIRMED"
	StatusRestPressure          = "PASS_REST_PRESSURE_DECOMPOSITION_DERIVED"
	StatusNonIdentifiability    = "PASS_AGGREGATE_NON_IDENTIFIABILITY_REAFFIRMED"
	StatusFiniteTripleAudited   = "PASS_FINITE_TRIPLE_TRACE_TEMPLATE_AUDITED"
	StatusExternalAudited       = "PASS_EXTERNAL_LEDGER_MAGNITUDE_SOURCE_AUDITED"
	StatusD4Audited             = "PASS_D4_TRILINEAR_MAGNITUDE_SOURCE_AUDITED"
	StatusK7ProjectiveAudited   = "PASS_K7_PROJECTIVE_MAGNITUDE_SOURCE_AUDITED"
	StatusScaleAudited          = "PASS_SCALE_LOCALITY_AUDITED"
	StatusScaleDifferential     = "PASS_N_EFF_SCALE_DIFFERENTIAL_RECORDED"
	StatusCHiggsAudited         = "PASS_C_HIGGS_IMPACT_AUDITED"
	StatusOutcomeRecorded       = "PASS_OUTCOME_CLASSIFICATION_RECORDED"
	StatusBranchDecision        = "PASS_BRANCH_DECISION_RECORDED"
	StatusPhysicalFirewalls     = "PASS_PHYSICAL_FIREWALLS_ENFORCED"

	StatusNeedsSpectraNotMixing      = "CONDITIONAL_SUPPORT_N_EFF_REQUIRES_HERMITIAN_SPECTRA_NOT_FULL_MIXING_DATA"
	StatusEffectiveTraceAtomCount    = "CONDITIONAL_SUPPORT_N_EFF_IS_EFFECTIVE_TRACE_ATOM_COUNT"
	StatusSpectralMagnitudeOnly      = "CONDITIONAL_SUPPORT_N_EFF_IS_SPECTRAL_MAGNITUDE_READOUT_ONLY"
	StatusNearThreeTopColorBlock     = "CONDITIONAL_SUPPORT_CURRENT_CERTIFIED_NEAR_THREE_SOURCE_IS_RANK_THREE_TOP_COLOR_BLOCK"
	StatusRestPressureAboveTop       = "CONDITIONAL_SUPPORT_N_EFF_MINUS_THREE_IS_REST_SPECTRAL_PRESSURE_ABOVE_TOP_COLOR_LIMIT"
	StatusFSTSuppliesTraceShape      = "CONDITIONAL_SUPPORT_FINITE_TRIPLE_SUPPLIES_TRACE_SHAPE"
	StatusExternalCanPopulate        = "CONDITIONAL_SUPPORT_EXTERNAL_LEDGER_CAN_POPULATE_TRACE_MAGNITUDE_OPERATOR_SEAL"
	StatusImprovesCYukawaTestability = "CONDITIONAL_SUPPORT_TRACE_MAGNITUDE_OPERATOR_SEAL_WOULD_DIRECTLY_IMPROVE_C_YUKAWA_TESTABILITY"
	StatusMagnitudeSharper           = "CONDITIONAL_SUPPORT_TRACE_MAGNITUDE_SIDE_IS_SHARPER_THAN_FULL_GENERATION_OPERATOR_FOR_N_EFF"
	StatusNextTopRestPressure        = "CONDITIONAL_SUPPORT_NEXT_NATIVE_GATE_SHOULD_AUDIT_TOP_COLOR_BLOCK_AND_REST_PRESSURE"

	StatusSealNotNative                 = "FAILED_ROUTE_TRACE_MAGNITUDE_OPERATOR_SEAL_NOT_CURRENTLY_NATIVE"
	StatusNoPMNSCKMFromNEff             = "FAILED_ROUTE_N_EFF_CANNOT_SOURCE_PMNS_OR_CKM_BY_ITSELF"
	StatusNoKappaOrientFromTrace        = "FAILED_ROUTE_TRACE_MAGNITUDE_LEDGER_DOES_NOT_SOURCE_KAPPA_ORIENT"
	StatusTopColorNotGenerationTriality = "FAILED_ROUTE_TOP_COLOR_THREE_NOT_GENERATION_TRIALITY_THEOREM"
	StatusTopBlockNoT                   = "FAILED_ROUTE_TOP_COLOR_BLOCK_DOES_NOT_DERIVE_T_VALUE"
	StatusNoAlphaBetaWithoutT           = "FAILED_ROUTE_NO_NUMERICAL_ALPHA_BETA_WITHOUT_TYPED_T_CHANNEL"
	StatusNoSectorRestAssignment        = "FAILED_ROUTE_REST_PRESSURE_NOT_SECTOR_ASSIGNED_WITHOUT_DECOMPOSED_LEDGER"
	StatusABNoOperators                 = "FAILED_ROUTE_A_B_ALONE_CANNOT_IDENTIFY_TRACE_MAGNITUDE_OPERATORS"
	StatusABNoTopChannel                = "FAILED_ROUTE_A_B_ALONE_CANNOT_IDENTIFY_TOP_CHANNEL"
	StatusABNoRestSectors               = "FAILED_ROUTE_A_B_ALONE_CANNOT_ASSIGN_REST_PRESSURE_TO_SECTORS"
	StatusFSTNoOperators                = "FAILED_ROUTE_FINITE_TRIPLE_DOES_NOT_SUPPLY_TRACE_MAGNITUDE_OPERATORS"
	StatusExternalNotNative             = "FAILED_ROUTE_EXTERNAL_LEDGER_NOT_NATIVE_TRACE_MAGNITUDE_THEOREM"
	StatusTD4NoMagnitudes               = "FAILED_ROUTE_T_D4_DOES_NOT_SUPPLY_HERMITIAN_TRACE_MAGNITUDES"
	StatusTD4NoNEff                     = "FAILED_ROUTE_T_D4_DOES_NOT_SOURCE_N_EFF"
	StatusK7NotMagnitudeOperator        = "FAILED_ROUTE_K7_MINUS_THREE_NOT_TRACE_MAGNITUDE_OPERATOR"
	StatusProjectiveNotNEff             = "FAILED_ROUTE_PROJECTIVE_ONE_PLUS_THREE_NOT_N_EFF_SOURCE"
	StatusNoScaleStability              = "FAILED_ROUTE_NO_NATIVE_N_EFF_SCALE_STABILITY_THEOREM"
	StatusMZScaleSealed                 = "FAILED_ROUTE_MZ_TRACE_MAGNITUDE_LEDGER_REMAINS_SCALE_SEALED"
	StatusSealAloneNoNativeCHiggs       = "FAILED_ROUTE_TRACE_MAGNITUDE_OPERATOR_SEAL_ALONE_DOES_NOT_MAKE_C_HIGGS_NATIVE"
	StatusCHiggsLevelB                  = "FAILED_ROUTE_C_HIGGS_REMAINS_LEVEL_B_WITH_EXTERNAL_OR_SEALED_SPECTRA"
	StatusFirewallGate807               = "FIREWALL_PRESERVED_GATE807_TRACE_MAGNITUDE_OPERATOR_BOUNDARY"
)

type Inheritance struct {
	Gate806Inherited bool
	Formula          string
	NEff             float64
	CYukawa          float64
	UnitLeverage     string
	Verdicts         []string
}

type TraceMagnitudeOperatorSeal struct {
	Defined    bool
	Name       string
	Components []string
	Chain      []string
	Verdict    string
	Supports   []string
	Failures   []string
}

type TraceFormulas struct {
	Recorded bool
	A        string
	B        string
	AtomForm string
	Identity string
	Verdicts []string
	Supports []string
}

type OrientationInvisibility struct {
	Audited         bool
	Transformation  string
	InvariantTraces []string
	Verdict         string
	Supports        []string
	Failures        []string
}

type TopColorBlock struct {
	Defined  bool
	TName    string
	ATop     string
	BTop     string
	NEffTop  float64
	Verdicts []string
	Supports []string
	Failures []string
}

type RestPressure struct {
	Derived        bool
	Definitions    []string
	NEffFormula    string
	RatioFormula   string
	DeltaFormula   string
	SmallFormula   string
	CurrentDelta   float64
	Interpretation string
	Verdict        string
	Supports       []string
	Failures       []string
}

type NonIdentifiability struct {
	Reaffirmed bool
	Reason     string
	CannotFind []string
	Verdict    string
	Failures   []string
}

type SourceAudit struct {
	Audited  bool
	Source   string
	Supplies []string
	Missing  []string
	Verdict  string
	Supports []string
	Failures []string
}

type ScaleAudit struct {
	Audited      bool
	ScaleLocal   string
	Differential string
	Needed       []string
	Verdicts     []string
	Failures     []string
}

type CHiggsImpact struct {
	Audited   bool
	Formula   string
	Upgrade   string
	Unchanged []string
	Verdict   string
	Supports  []string
	Failures  []string
}

type Outcome struct {
	Recorded bool
	Items    []string
	Verdict  string
	Supports []string
}

type BranchDecision struct {
	Recorded    bool
	Next        string
	Alternative string
	Reason      string
	Verdict     string
	Supports    []string
}

type Firewalls struct {
	Enforced      bool
	NoPMNSCKM     bool
	NoYukawa      bool
	NoEigenvalues bool
	NoFlavor      bool
	NoScalar      bool
	NoPoleMass    bool
	NoVEVGF       bool
	NoGJ          bool
	NoTriality    bool
	NoHistoryLoop bool
	Verdict       string
}

type Analysis struct {
	Inheritance        Inheritance
	Seal               TraceMagnitudeOperatorSeal
	Formulas           TraceFormulas
	Orientation        OrientationInvisibility
	TopColor           TopColorBlock
	Rest               RestPressure
	NonIdentifiability NonIdentifiability
	FiniteTriple       SourceAudit
	External           SourceAudit
	TD4                SourceAudit
	K7Projective       SourceAudit
	Scale              ScaleAudit
	CHiggs             CHiggsImpact
	Outcome            Outcome
	Branch             BranchDecision
	Firewalls          Firewalls
	Truth              string
	Final              string
}

func TopColorNEff(T float64) (float64, error) {
	if T <= 0 {
		return 0, fmt.Errorf("top-like Hermitian eigenvalue T must be positive")
	}
	aTop := 3 * T
	bTop := 3 * T * T
	return aTop * aTop / bTop, nil
}

func RestPressureNEff(alpha, beta float64) (float64, error) {
	if 1+beta <= 0 {
		return 0, fmt.Errorf("1+beta must be positive")
	}
	return 3 * math.Pow(1+alpha, 2) / (1 + beta), nil
}

func RestPressureDelta(alpha, beta float64) (float64, error) {
	neff, err := RestPressureNEff(alpha, beta)
	if err != nil {
		return 0, err
	}
	return neff - 3, nil
}

func BuildDefault() (Analysis, error) {
	neffTop, err := TopColorNEff(0.9471025365183062)
	if err != nil {
		return Analysis{}, err
	}
	if math.Abs(neffTop-3) > 1e-12 {
		return Analysis{}, fmt.Errorf("top-color limit failed: got %.17g", neffTop)
	}

	inheritance := Inheritance{
		Gate806Inherited: true,
		Formula:          "C_Higgs = (3/N_eff) C_History",
		NEff:             NEff,
		CYukawa:          CYukawa,
		UnitLeverage:     "delta C_Higgs / C_Higgs = - delta N_eff / N_eff",
		Verdicts:         []string{StatusGate806Inherited, StatusNEffSubproblem, StatusUnitLeverageInherited},
	}
	if !inheritance.Gate806Inherited || !strings.Contains(inheritance.Formula, "N_eff") || math.Abs(inheritance.NEff-NEff) > 1e-15 {
		return Analysis{}, fmt.Errorf("Gate 807 requires Gate 806 inheritance and active N_eff ledger")
	}

	seal := TraceMagnitudeOperatorSeal{
		Defined: true,
		Name:    "TraceMagnitudeOperatorSeal",
		Components: []string{
			"sector Hermitian operators H_u,H_d,H_e,H_nu", "positive spectra Spec(H_f)", "color multiplicity rule",
			"sector trace ledger", "trace atom ledger", "top-dominant block selector", "rest-pressure spectral measure",
			"scale and scheme convention", "neutrino convention", "noncircularity proof",
		},
		Chain:    []string{"H_f=Y_f†Y_f >= 0", "Spec(H_f)", "x_i >= 0", "a,b,N_eff"},
		Verdict:  StatusSealDefined,
		Supports: []string{StatusNeedsSpectraNotMixing},
		Failures: []string{StatusSealNotNative},
	}

	formulas := TraceFormulas{
		Recorded: true,
		A:        "a = Tr(H_e + H_nu + 3H_u + 3H_d)",
		B:        "b = Tr(H_e² + H_nu² + 3H_u² + 3H_d²)",
		AtomForm: "a = sum_i x_i; b = sum_i x_i²; N_eff = a²/b",
		Identity: "w_i=x_i/a; sum_i w_i=1; b/a²=sum_i w_i²; N_eff=1/sum_i w_i²",
		Verdicts: []string{StatusTraceFormulas, StatusParticipationIdentity},
		Supports: []string{StatusEffectiveTraceAtomCount},
	}

	orientation := OrientationInvisibility{
		Audited:         true,
		Transformation:  "H_f -> U_f H_f U_f†",
		InvariantTraces: []string{"Tr(H_f)", "Tr(H_f²)", "a", "b", "N_eff"},
		Verdict:         StatusOrientationAudited,
		Supports:        []string{StatusSpectralMagnitudeOnly},
		Failures:        []string{StatusNoPMNSCKMFromNEff, StatusNoKappaOrientFromTrace},
	}

	top := TopColorBlock{
		Defined:  true,
		TName:    "T = h_t",
		ATop:     "a_top = 3T",
		BTop:     "b_top = 3T²",
		NEffTop:  neffTop,
		Verdicts: []string{StatusRankThreeBlock, StatusTopLimit},
		Supports: []string{StatusNearThreeTopColorBlock},
		Failures: []string{StatusTopColorNotGenerationTriality, StatusTopBlockNoT},
	}

	rest := RestPressure{
		Derived:        true,
		Definitions:    []string{"a_rest = a - 3T", "b_rest = b - 3T²", "alpha = a_rest/(3T)", "beta = b_rest/(3T²)"},
		NEffFormula:    "N_eff = 3(1+alpha)²/(1+beta)",
		RatioFormula:   "b/a² = (1/3)(1+beta)/(1+alpha)²",
		DeltaFormula:   "N_eff - 3 = 3(2alpha + alpha² - beta)/(1+beta)",
		SmallFormula:   "N_eff - 3 ≈ 3(2alpha - beta)",
		CurrentDelta:   NEffDelta,
		Interpretation: "N_eff > 3 means rest spectrum increases quadratic trace participation more than quartic concentration",
		Verdict:        StatusRestPressure,
		Supports:       []string{StatusRestPressureAboveTop},
		Failures:       []string{StatusNoAlphaBetaWithoutT, StatusNoSectorRestAssignment},
	}

	nonID := NonIdentifiability{
		Reaffirmed: true,
		Reason:     "aggregate a,b impose only two scalar constraints on an unknown positive trace-atom list",
		CannotFind: []string{"TraceMagnitudeOperatorSeal", "top channel T", "alpha,beta", "sector fractions", "generation fractions", "neutrino contribution", "bottom/tau/charm pressure", "D4/triality structure"},
		Verdict:    StatusNonIdentifiability,
		Failures:   []string{StatusABNoOperators, StatusABNoTopChannel, StatusABNoRestSectors},
	}

	fst := SourceAudit{Audited: true, Source: "Finite spectral triple", Supplies: []string{"trace-form templates", "Tr(H_e + H_nu + 3H_u + 3H_d)", "Tr(H_e² + H_nu² + 3H_u² + 3H_d²)"}, Missing: []string{"H_f spectra", "positive trace atoms"}, Verdict: StatusFiniteTripleAudited, Supports: []string{StatusFSTSuppliesTraceShape}, Failures: []string{StatusFSTNoOperators}}
	ext := SourceAudit{Audited: true, Source: "External Yukawa ledger", Supplies: []string{"H_f spectra or equivalent singular values", "trace atoms", "sector fractions", "top/rest decomposition"}, Missing: []string{"native theorem"}, Verdict: StatusExternalAudited, Supports: []string{StatusExternalCanPopulate}, Failures: []string{StatusExternalNotNative}}
	td4 := SourceAudit{Audited: true, Source: "Complex D4 trilinear", Supplies: []string{"airlocked edge-kernel shape"}, Missing: []string{"positive Hermitian spectra", "top-dominant block", "rest-pressure operator"}, Verdict: StatusD4Audited, Failures: []string{StatusTD4NoMagnitudes, StatusTD4NoNEff}}
	k7 := SourceAudit{Audited: true, Source: "K7 / Fock / projective structures", Supplies: []string{"dim K7- = 3", "projective 4 = 1 + 3", "future carrier-search resonances"}, Missing: []string{"positive sector spectra", "TraceMagnitudeOperatorSeal"}, Verdict: StatusK7ProjectiveAudited, Failures: []string{StatusK7NotMagnitudeOperator, StatusProjectiveNotNEff}}

	scale := ScaleAudit{
		Audited:      true,
		ScaleLocal:   "N_eff(mu) = a(mu)² / b(mu); current value is an M_Z ledger value",
		Differential: "d ln N_eff = 2 d ln a - d ln b",
		Needed:       []string{"exact scale-invariance", "controlled RG transport of H_f", "preferred readout scale theorem", "multi-scale ledger"},
		Verdicts:     []string{StatusScaleAudited, StatusScaleDifferential},
		Failures:     []string{StatusNoScaleStability, StatusMZScaleSealed},
	}

	chiggs := CHiggsImpact{
		Audited:   true,
		Formula:   "C_Yukawa = 3/N_eff; C_Higgs = C_Yukawa C_History",
		Upgrade:   "validated TraceMagnitudeOperatorSeal would upgrade N_eff from aggregate trace seal to sector/atom-auditable trace-magnitude seal",
		Unchanged: []string{"C_Higgs remains Level B unless spectra are native and noncircular", "lambda_H_bridge", "m_H_tree_proxy"},
		Verdict:   StatusCHiggsAudited,
		Supports:  []string{StatusImprovesCYukawaTestability},
		Failures:  []string{StatusSealAloneNoNativeCHiggs, StatusCHiggsLevelB},
	}

	outcome := Outcome{Recorded: true, Items: []string{"N_eff depends only on Hermitian trace spectra, not PMNS/CKM orientation", "exact N_eff = 3 is sourced by a rank-three top-color dominant block", "N_eff - 3 measures unresolved rest spectral pressure", "aggregate a,b do not identify the trace-magnitude operators", "current ASHA does not supply native H_f spectra", "C_Higgs remains Level B"}, Verdict: StatusOutcomeRecorded, Supports: []string{StatusMagnitudeSharper}}
	branch := BranchDecision{Recorded: true, Next: "Gate 808 — RankThreeTopColorBlock and RestPressureOperator Source Audit", Alternative: "Gate 808 — External Yukawa Trace Magnitude Ledger Validation and Sector Contribution Audit", Reason: "audit the minimal spectral object explaining N_eff = 3 + small deviation: dominant color-tripled top block plus suppressed positive rest spectrum", Verdict: StatusBranchDecision, Supports: []string{StatusNextTopRestPressure}}
	firewalls := Firewalls{Enforced: true, NoPMNSCKM: true, NoYukawa: true, NoEigenvalues: true, NoFlavor: true, NoScalar: true, NoPoleMass: true, NoVEVGF: true, NoGJ: true, NoTriality: true, NoHistoryLoop: true, Verdict: StatusFirewallGate807}

	return Analysis{Inheritance: inheritance, Seal: seal, Formulas: formulas, Orientation: orientation, TopColor: top, Rest: rest, NonIdentifiability: nonID, FiniteTriple: fst, External: ext, TD4: td4, K7Projective: k7, Scale: scale, CHiggs: chiggs, Outcome: outcome, Branch: branch, Firewalls: firewalls, Truth: "Gate 807 sharpens N_eff: the scalar-Higgs bridge needs positive Hermitian spectra H_f=Y_f†Y_f, not the full PMNS/CKM orientation package.", Final: "The strongest current near-three source is one dominant top-like eigenvalue times color multiplicity three; N_eff-3 is unresolved rest spectral pressure. The next native target is RankThreeTopColorBlockSeal plus RestPressureOperatorSeal."}, nil
}

func Statuses() []string {
	return []string{
		StatusGate806Inherited, StatusNEffSubproblem, StatusUnitLeverageInherited, StatusSealDefined,
		StatusTraceFormulas, StatusParticipationIdentity, StatusOrientationAudited, StatusRankThreeBlock,
		StatusTopLimit, StatusRestPressure, StatusNonIdentifiability, StatusFiniteTripleAudited,
		StatusExternalAudited, StatusD4Audited, StatusK7ProjectiveAudited, StatusScaleAudited,
		StatusScaleDifferential, StatusCHiggsAudited, StatusOutcomeRecorded, StatusBranchDecision,
		StatusPhysicalFirewalls, StatusNeedsSpectraNotMixing, StatusEffectiveTraceAtomCount,
		StatusSpectralMagnitudeOnly, StatusNearThreeTopColorBlock, StatusRestPressureAboveTop,
		StatusFSTSuppliesTraceShape, StatusExternalCanPopulate, StatusImprovesCYukawaTestability,
		StatusMagnitudeSharper, StatusNextTopRestPressure, StatusSealNotNative, StatusNoPMNSCKMFromNEff,
		StatusNoKappaOrientFromTrace, StatusTopColorNotGenerationTriality, StatusTopBlockNoT,
		StatusNoAlphaBetaWithoutT, StatusNoSectorRestAssignment, StatusABNoOperators, StatusABNoTopChannel,
		StatusABNoRestSectors, StatusFSTNoOperators, StatusExternalNotNative, StatusTD4NoMagnitudes,
		StatusTD4NoNEff, StatusK7NotMagnitudeOperator, StatusProjectiveNotNEff, StatusNoScaleStability,
		StatusMZScaleSealed, StatusSealAloneNoNativeCHiggs, StatusCHiggsLevelB, StatusFirewallGate807,
	}
}

func FormatSeal(s TraceMagnitudeOperatorSeal) string {
	return fmt.Sprintf("%s components=[%s] chain=[%s] supports=[%s] failures=[%s]", s.Name, strings.Join(s.Components, "; "), strings.Join(s.Chain, " -> "), strings.Join(s.Supports, "; "), strings.Join(s.Failures, "; "))
}

func FormatFormulas(f TraceFormulas) string {
	return fmt.Sprintf("%s | %s | %s | %s supports=[%s]", f.A, f.B, f.AtomForm, f.Identity, strings.Join(f.Supports, "; "))
}

func FormatOrientation(o OrientationInvisibility) string {
	return fmt.Sprintf("%s invariants=[%s] supports=[%s] failures=[%s]", o.Transformation, strings.Join(o.InvariantTraces, "; "), strings.Join(o.Supports, "; "), strings.Join(o.Failures, "; "))
}

func FormatTop(t TopColorBlock) string {
	return fmt.Sprintf("%s; %s; %s; N_eff_top=%.12g supports=[%s] failures=[%s]", t.TName, t.ATop, t.BTop, t.NEffTop, strings.Join(t.Supports, "; "), strings.Join(t.Failures, "; "))
}

func FormatRest(r RestPressure) string {
	return fmt.Sprintf("defs=[%s] %s; %s; %s; small=%s; currentDelta=%.16g; supports=[%s] failures=[%s]", strings.Join(r.Definitions, "; "), r.NEffFormula, r.RatioFormula, r.DeltaFormula, r.SmallFormula, r.CurrentDelta, strings.Join(r.Supports, "; "), strings.Join(r.Failures, "; "))
}

func FormatNonID(n NonIdentifiability) string {
	return fmt.Sprintf("reason=%q cannotFind=[%s] failures=[%s]", n.Reason, strings.Join(n.CannotFind, "; "), strings.Join(n.Failures, "; "))
}

func FormatSource(s SourceAudit) string {
	return fmt.Sprintf("%s supplies=[%s] missing=[%s] supports=[%s] failures=[%s]", s.Source, strings.Join(s.Supplies, "; "), strings.Join(s.Missing, "; "), strings.Join(s.Supports, "; "), strings.Join(s.Failures, "; "))
}

func FormatScale(s ScaleAudit) string {
	return fmt.Sprintf("%s; %s; needed=[%s] failures=[%s]", s.ScaleLocal, s.Differential, strings.Join(s.Needed, "; "), strings.Join(s.Failures, "; "))
}

func FormatCHiggs(c CHiggsImpact) string {
	return fmt.Sprintf("%s upgrade=%q unchanged=[%s] supports=[%s] failures=[%s]", c.Formula, c.Upgrade, strings.Join(c.Unchanged, "; "), strings.Join(c.Supports, "; "), strings.Join(c.Failures, "; "))
}

func FormatOutcome(o Outcome) string {
	return fmt.Sprintf("items=[%s] supports=[%s]", strings.Join(o.Items, "; "), strings.Join(o.Supports, "; "))
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
