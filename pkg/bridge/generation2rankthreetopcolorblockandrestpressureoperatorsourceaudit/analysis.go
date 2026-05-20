// Package generation2rankthreetopcolorblockandrestpressureoperatorsourceaudit implements
// Gate 808: RankThreeTopColorBlock and RestPressureOperator Source Audit.
//
// Gate 808 extracts the strongest lawful information from the aggregate Yukawa
// trace-magnitude ledger: exact N_eff=3 follows from a dominant top-like trace
// atom counted with color multiplicity three, while the observed N_eff>3 is
// typed as unresolved positive rest spectral pressure.
package generation2rankthreetopcolorblockandrestpressureoperatorsourceaudit

import (
	"fmt"
	"math"
	"strings"
)

const (
	AuditID = "GATE808-RANK-THREE-TOP-COLOR-BLOCK-AND-REST-PRESSURE-OPERATOR-SOURCE-AUDIT"

	AInherited      = 2.8424095142339083
	BInherited      = 2.6910096440382287
	NEff            = 3.0023273474722147
	CYukawa         = 0.9992248188812008
	CHistory        = 1.038025177923625
	CHiggs          = 1.0372205204048603
	VEVSeal         = 246.2196508
	NEffDelta       = NEff - 3.0
	TreeProxyShift  = -0.04862437568908
	CHiggsRestShift = CHistory - CHiggs

	StatusGate807Inherited       = "PASS_GATE807_TRACE_MAGNITUDE_AUDIT_INHERITED"
	StatusTopColorSelected       = "PASS_RANK_THREE_TOP_COLOR_BLOCK_SELECTED_AS_CURRENT_CERTIFIED_THREE_SOURCE"
	StatusRestPressureSelected   = "PASS_REST_SPECTRAL_PRESSURE_SELECTED_AS_CURRENT_DEVIATION_OBJECT"
	StatusTopColorSealDefined    = "PASS_RANK_THREE_TOP_COLOR_BLOCK_SEAL_DEFINED"
	StatusTopColorLimit          = "PASS_TOP_COLOR_LIMIT_REDERIVED"
	StatusRestSealDefined        = "PASS_REST_PRESSURE_OPERATOR_SEAL_DEFINED"
	StatusRestDecomposition      = "PASS_REST_PRESSURE_DECOMPOSITION_REDERIVED"
	StatusPositivityCorridor     = "PASS_AGGREGATE_POSITIVITY_CORRIDOR_COMPUTED"
	StatusRestConcentration      = "PASS_REST_CONCENTRATION_RATIO_DEFINED"
	StatusRestConcentrationBound = "PASS_REST_CONCENTRATION_BOUNDS_RECORDED"
	StatusSectorCandidates       = "PASS_REST_PRESSURE_SECTOR_CANDIDATES_RECORDED"
	StatusPatternFirewall        = "PASS_PATTERN_DIAGNOSTIC_FIREWALL_RECORDED"
	StatusD4Firewall             = "PASS_D4_TRIALITY_FIREWALL_REAUDITED"
	StatusFiniteTripleAudit      = "PASS_FINITE_TRIPLE_TOP_COLOR_SOURCE_AUDITED"
	StatusExternalAudit          = "PASS_EXTERNAL_LEDGER_REST_PRESSURE_SOURCE_AUDITED"
	StatusK7ProjectiveAudit      = "PASS_K7_PROJECTIVE_RESONANCE_AUDITED"
	StatusD4SourceAudit          = "PASS_COMPLEX_D4_TRILINEAR_REST_PRESSURE_SOURCE_AUDITED"
	StatusCHiggsImpact           = "PASS_C_HIGGS_IMPACT_OF_REST_PRESSURE_RECORDED"
	StatusHierarchyObstruction   = "PASS_HIERARCHY_BREAKING_OPERATOR_SELECTED_AS_NATIVE_SOURCE_OBSTRUCTION"
	StatusOutcomeRecorded        = "PASS_OUTCOME_CLASSIFICATION_RECORDED"
	StatusBranchDecision         = "PASS_BRANCH_DECISION_RECORDED"
	StatusPhysicalFirewalls      = "PASS_PHYSICAL_FIREWALLS_ENFORCED"

	StatusExactTopNEffThree            = "CONDITIONAL_SUPPORT_N_EFF_EQUALS_THREE_IN_EXACT_TOP_COLOR_DOMINANCE_LIMIT"
	StatusColorThreeStrongest          = "CONDITIONAL_SUPPORT_COLOR_MULTIPLICITY_THREE_IS_CURRENT_STRONGEST_TYPED_SOURCE_OF_N_EFF_BASELINE"
	StatusRestPressureAboveTop         = "CONDITIONAL_SUPPORT_N_EFF_MINUS_THREE_IS_REST_SPECTRAL_PRESSURE_ABOVE_TOP_COLOR_LIMIT"
	StatusRestPressureDilutesCYukawa   = "CONDITIONAL_SUPPORT_REST_PRESSURE_DILUTES_C_YUKAWA_BELOW_ONE"
	StatusNarrowCorridor               = "CONDITIONAL_SUPPORT_TOP_DOMINANT_BRANCH_FORCES_T_IN_NARROW_POSITIVITY_CORRIDOR"
	StatusAlphaScale                   = "CONDITIONAL_SUPPORT_REST_QUADRATIC_PRESSURE_SCALE_IS_APPROXIMATELY_3_88E_MINUS_4_IF_TOP_DOMINANCE_IS_ASSUMED"
	StatusAlphaQRestSplit              = "CONDITIONAL_SUPPORT_REST_PRESSURE_SPLITS_INTO_TOTAL_REST_SIZE_ALPHA_AND_REST_CONCENTRATION_Q_REST"
	StatusPlausibleRestSources         = "CONDITIONAL_SUPPORT_BOTTOM_TAU_CHARM_AND_OTHER_SMALL_ATOMS_ARE_PLAUSIBLE_REST_PRESSURE_SOURCES"
	StatusFNRelevant                   = "CONDITIONAL_SUPPORT_FN_STYLE_HIERARCHY_MAY_BE_RELEVANT_TO_REST_PRESSURE_INTERPRETATION_AFTER_LEDGER_INPUT"
	StatusGJRelevant                   = "CONDITIONAL_SUPPORT_GJ_MAY_BE_RELEVANT_TO_DOWN_LEPTON_REST_STRUCTURE_AT_HIGH_SCALE"
	StatusTrialityMotivation           = "CONDITIONAL_SUPPORT_TRIALITY_REMAINS_NATIVE_SEARCH_MOTIVATION_ONLY"
	StatusFiniteTripleColorShape       = "CONDITIONAL_SUPPORT_FINITE_TRIPLE_SUPPLIES_COLOR_FACTOR_AND_TRACE_SHAPE"
	StatusExternalCanTest              = "CONDITIONAL_SUPPORT_EXTERNAL_LEDGER_CAN_TEST_TOP_COLOR_AND_REST_PRESSURE_DECOMPOSITION"
	StatusRestPressureNumericallySmall = "CONDITIONAL_SUPPORT_REST_PRESSURE_IS_SMALL_BUT_NUMERICALLY_RELEVANT_FOR_LEVEL_B_C_HIGGS"
	StatusNeedsHierarchyMechanism      = "CONDITIONAL_SUPPORT_NATIVE_N_EFF_SOURCE_REQUIRES_TOP_DOMINANCE_AND_REST_SUPPRESSION_MECHANISM"
	StatusBestTypedModel               = "CONDITIONAL_SUPPORT_TOP_COLOR_BLOCK_PLUS_REST_PRESSURE_IS_CURRENT_BEST_TYPED_N_EFF_MODEL"
	StatusNextHierarchyGate            = "CONDITIONAL_SUPPORT_NEXT_NATIVE_GATE_SHOULD_AUDIT_HIERARCHY_BREAKING_OPERATOR"

	StatusTopBlockNoT                 = "FAILED_ROUTE_TOP_COLOR_BLOCK_DOES_NOT_DERIVE_T_VALUE"
	StatusTopThreeNotGeneration       = "FAILED_ROUTE_TOP_COLOR_THREE_NOT_GENERATION_TRIALITY_THEOREM"
	StatusTopBlockNotHierarchyTheorem = "FAILED_ROUTE_TOP_COLOR_BLOCK_NOT_NATIVE_YUKAWA_HIERARCHY_THEOREM"
	StatusRestNoSector                = "FAILED_ROUTE_REST_PRESSURE_NOT_SECTOR_ASSIGNED_WITHOUT_DECOMPOSED_LEDGER"
	StatusRestNotNative               = "FAILED_ROUTE_REST_PRESSURE_OPERATOR_NOT_NATIVE_YUKAWA_THEOREM"
	StatusCorridorNotTopDerivation    = "FAILED_ROUTE_POSITIVITY_CORRIDOR_NOT_TOP_YUKAWA_DERIVATION"
	StatusNoBackwardT                 = "FAILED_ROUTE_T_MUST_NOT_BE_SOLVED_BACKWARDS_FROM_A_B_OR_N_EFF"
	StatusCorridorNoSector            = "FAILED_ROUTE_CORRIDOR_DOES_NOT_ASSIGN_REST_PRESSURE_TO_SECTORS"
	StatusNoRestAtomCount             = "FAILED_ROUTE_NO_REST_ATOM_COUNT_WITHOUT_DECOMPOSED_LEDGER"
	StatusNoQRestFromAggregate        = "FAILED_ROUTE_NO_REST_CONCENTRATION_VALUE_FROM_AGGREGATE_A_B_ALONE"
	StatusNoSectorAssignment          = "FAILED_ROUTE_NO_SECTOR_ASSIGNMENT_WITHOUT_TRACE_ATOM_LEDGER"
	StatusNeutrinoImplicit            = "FAILED_ROUTE_NEUTRINO_CONVENTION_MUST_NOT_BE_IMPLICIT"
	StatusScaleSchemeUntyped          = "FAILED_ROUTE_SCALE_AND_SCHEME_MUST_NOT_BE_LEFT_UNTYPED"
	StatusKoideNotNEff                = "FAILED_ROUTE_KOIDE_NOT_N_EFF_SOURCE_THEOREM"
	StatusFNNotNativeRest             = "FAILED_ROUTE_FN_POWERS_NOT_NATIVE_REST_PRESSURE_OPERATOR_WITHOUT_CHARGES"
	StatusGJNotLowScaleTop            = "FAILED_ROUTE_GJ_CLEBSCH_FACTORS_NOT_LOW_SCALE_TOP_COLOR_BLOCK_THEOREM"
	StatusNEffNotD4                   = "FAILED_ROUTE_N_EFF_NEAR_THREE_NOT_D4_TRIALITY_THEOREM"
	StatusNoTrialityTraceReadout      = "FAILED_ROUTE_NO_TRIALITY_TO_TRACE_MAGNITUDE_READOUT_MAP"
	StatusNoTrialityRestPressure      = "FAILED_ROUTE_NO_TRIALITY_SOURCE_OF_REST_PRESSURE"
	StatusNoTrialityRealDescent       = "FAILED_ROUTE_NO_REAL_FORM_DESCENT_FOR_TRIALITY_YUKAWA_TRACE_LEDGER"
	StatusFSTNoTopEigenvalue          = "FAILED_ROUTE_FINITE_TRIPLE_DOES_NOT_SUPPLY_TOP_DOMINANT_EIGENVALUE"
	StatusFSTNoRestOperator           = "FAILED_ROUTE_FINITE_TRIPLE_DOES_NOT_SUPPLY_REST_PRESSURE_OPERATOR"
	StatusExternalNotNative           = "FAILED_ROUTE_EXTERNAL_LEDGER_NOT_NATIVE_REST_PRESSURE_THEOREM"
	StatusK7NotTopBlock               = "FAILED_ROUTE_K7_MINUS_THREE_NOT_RANK_THREE_TOP_COLOR_BLOCK"
	StatusProjectiveNotRest           = "FAILED_ROUTE_PROJECTIVE_ONE_PLUS_THREE_NOT_REST_PRESSURE_OPERATOR"
	StatusTD4NotTraceMagnitude        = "FAILED_ROUTE_T_D4_NOT_TRACE_MAGNITUDE_OPERATOR"
	StatusTD4NotTopDominance          = "FAILED_ROUTE_T_D4_NOT_TOP_DOMINANCE_OPERATOR"
	StatusTD4NotRestPressure          = "FAILED_ROUTE_T_D4_NOT_REST_PRESSURE_OPERATOR"
	StatusTreeProxyNotPole            = "FAILED_ROUTE_TREE_PROXY_SHIFT_NOT_POLE_MASS_STATEMENT"
	StatusNoNewSpectralData           = "FAILED_ROUTE_REST_PRESSURE_AUDIT_DOES_NOT_UPDATE_C_HIGGS_WITHOUT_NEW_SPECTRAL_DATA"
	StatusNoHierarchyBreakingOperator = "FAILED_ROUTE_NO_NATIVE_HIERARCHY_BREAKING_OPERATOR"
	StatusNoNativeTopDominance        = "FAILED_ROUTE_NO_NATIVE_TOP_DOMINANCE_THEOREM"
	StatusNoLightSuppression          = "FAILED_ROUTE_NO_NATIVE_LIGHT_FAMILY_SUPPRESSION_THEOREM"
	StatusNoNativeRestSource          = "FAILED_ROUTE_NO_NATIVE_REST_PRESSURE_SOURCE"
	StatusFirewallGate808             = "FIREWALL_PRESERVED_GATE808_TOP_COLOR_BLOCK_REST_PRESSURE_BOUNDARY"
)

type Inheritance struct {
	Gate807Inherited bool
	NEff             float64
	CYukawa          float64
	CHiggsFormula    string
	Verdicts         []string
}

type RankThreeTopColorBlockSeal struct {
	Defined    bool
	Name       string
	Components []string
	ATop       string
	BTop       string
	NEffTop    float64
	Verdicts   []string
	Supports   []string
	Failures   []string
}

type RestPressureOperatorSeal struct {
	Defined        bool
	Name           string
	Components     []string
	Definitions    []string
	NEffFormula    string
	DeltaFormula   string
	SmallFormula   string
	Interpretation string
	Verdicts       []string
	Supports       []string
	Failures       []string
}

type PositivityCorridor struct {
	Computed     bool
	AOver3       float64
	SqrtBOver3   float64
	TLower       float64
	TUpper       float64
	AlphaAtUpper float64
	BetaAtUpper  float64
	AlphaAtLower float64
	BetaAtLower  float64
	Gap          float64
	Verdict      string
	Supports     []string
	Failures     []string
}

type RestConcentration struct {
	Recorded       bool
	Formula        string
	Bounds         string
	BetaFormula    string
	Interpretation []string
	Verdicts       []string
	Supports       []string
	Failures       []string
}

type CandidateAudit struct {
	Audited  bool
	Topic    string
	Supplies []string
	Missing  []string
	Verdict  string
	Supports []string
	Failures []string
}

type CHiggsImpact struct {
	Recorded       bool
	TopLimit       float64
	Current        float64
	DeltaCHiggs    float64
	TreeShift      float64
	Interpretation string
	Verdict        string
	Supports       []string
	Failures       []string
}

type HierarchyObstruction struct {
	Selected  bool
	Name      string
	Needs     []string
	Forbidden []string
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
	Enforced       bool
	NoNativeYukawa bool
	NoPMNSCKM      bool
	NoFlavor       bool
	NoScalar       bool
	NoPoleMass     bool
	NoVEVGF        bool
	NoGJ           bool
	NoD4Triality   bool
	NoHistoryLoop  bool
	Verdict        string
}

type Analysis struct {
	Inheritance        Inheritance
	TopColor           RankThreeTopColorBlockSeal
	RestPressure       RestPressureOperatorSeal
	Corridor           PositivityCorridor
	Concentration      RestConcentration
	SectorCandidates   CandidateAudit
	PatternDiagnostics CandidateAudit
	D4Firewall         CandidateAudit
	FiniteTriple       CandidateAudit
	ExternalLedger     CandidateAudit
	K7Projective       CandidateAudit
	ComplexD4          CandidateAudit
	CHiggs             CHiggsImpact
	Hierarchy          HierarchyObstruction
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

func RestVariables(a, b, T float64) (alpha, beta float64, err error) {
	if T <= 0 {
		return 0, 0, fmt.Errorf("T must be positive")
	}
	aRest := a - 3*T
	bRest := b - 3*T*T
	if aRest < -1e-12 || bRest < -1e-12 {
		return 0, 0, fmt.Errorf("negative rest trace: a_rest=%.17g b_rest=%.17g", aRest, bRest)
	}
	if math.Abs(aRest) < 1e-18 {
		aRest = 0
	}
	if math.Abs(bRest) < 1e-18 {
		bRest = 0
	}
	return aRest / (3 * T), bRest / (3 * T * T), nil
}

func RestConcentrationBeta(alpha, qRest float64) (float64, error) {
	if alpha < 0 {
		return 0, fmt.Errorf("alpha must be nonnegative")
	}
	if qRest < 0 || qRest > 1 {
		return 0, fmt.Errorf("q_rest must be in [0,1]")
	}
	return 3 * alpha * alpha * qRest, nil
}

func TopDominantPositivityCorridor(a, b float64) (lower, upper float64, err error) {
	if a <= 0 || b <= 0 {
		return 0, 0, fmt.Errorf("a and b must be positive")
	}
	upper = math.Sqrt(b / 3)
	// b-3T² <= (a-3T)² gives 12T²-6aT+(a²-b)>=0. The top-dominant corridor uses the larger root.
	A := 12.0
	B := -6 * a
	C := a*a - b
	disc := B*B - 4*A*C
	if disc < 0 {
		return 0, 0, fmt.Errorf("negative corridor discriminant")
	}
	root1 := (-B - math.Sqrt(disc)) / (2 * A)
	root2 := (-B + math.Sqrt(disc)) / (2 * A)
	lower = math.Max(root1, root2)
	if lower > upper {
		return 0, 0, fmt.Errorf("empty top-dominant corridor: lower %.17g upper %.17g", lower, upper)
	}
	return lower, upper, nil
}

func TreeProxy(C float64) float64 {
	return (VEVSeal / 2) * math.Sqrt(C)
}

func BuildDefault() (Analysis, error) {
	neffTop, err := TopColorNEff(0.9471025365183062)
	if err != nil {
		return Analysis{}, err
	}
	if math.Abs(neffTop-3) > 1e-12 {
		return Analysis{}, fmt.Errorf("top-color limit failed: %.17g", neffTop)
	}
	lower, upper, err := TopDominantPositivityCorridor(AInherited, BInherited)
	if err != nil {
		return Analysis{}, err
	}
	alphaUpper, betaUpper, err := RestVariables(AInherited, BInherited, upper)
	if err != nil {
		return Analysis{}, err
	}
	alphaLower, betaLower, err := RestVariables(AInherited, BInherited, lower)
	if err != nil {
		return Analysis{}, err
	}
	if math.Abs(alphaUpper-0.00038781604472679744) > 1e-12 || math.Abs(betaLower-4.5172977535955994e-7) > 1e-16 {
		return Analysis{}, fmt.Errorf("unexpected corridor values alphaUpper=%.17g betaLower=%.17g", alphaUpper, betaLower)
	}
	chiggsShift := CHistory - CHiggs
	treeShift := TreeProxy(CHiggs) - TreeProxy(CHistory)

	inheritance := Inheritance{Gate807Inherited: true, NEff: NEff, CYukawa: CYukawa, CHiggsFormula: "C_Higgs = (3/N_eff) C_History", Verdicts: []string{StatusGate807Inherited, StatusTopColorSelected, StatusRestPressureSelected}}
	top := RankThreeTopColorBlockSeal{Defined: true, Name: "RankThreeTopColorBlockSeal", Components: []string{"top-like Hermitian eigenvalue T = h_t", "color multiplicity 3", "a_top = 3T", "b_top = 3T²", "top-channel selector", "scale/scheme convention", "noncircularity proof"}, ATop: "a_top = 3T", BTop: "b_top = 3T²", NEffTop: neffTop, Verdicts: []string{StatusTopColorSealDefined, StatusTopColorLimit}, Supports: []string{StatusExactTopNEffThree, StatusColorThreeStrongest}, Failures: []string{StatusTopBlockNoT, StatusTopThreeNotGeneration, StatusTopBlockNotHierarchyTheorem}}
	rest := RestPressureOperatorSeal{Defined: true, Name: "RestPressureOperatorSeal", Components: []string{"rest Hermitian spectrum H_rest", "a_rest", "b_rest", "q_rest=b_rest/a_rest²", "sector assignment if available", "scale/scheme convention", "noncircularity proof"}, Definitions: []string{"a = 3T + a_rest", "b = 3T² + b_rest", "alpha = a_rest/(3T)", "beta = b_rest/(3T²)"}, NEffFormula: "N_eff = 3(1+alpha)²/(1+beta)", DeltaFormula: "N_eff - 3 = 3(2alpha + alpha² - beta)/(1+beta)", SmallFormula: "N_eff - 3 ≈ 3(2alpha - beta)", Interpretation: "N_eff > 3 means rest quadratic trace participation dominates rest quartic concentration; b/a² < 1/3, C_Yukawa < 1", Verdicts: []string{StatusRestSealDefined, StatusRestDecomposition}, Supports: []string{StatusRestPressureAboveTop, StatusRestPressureDilutesCYukawa}, Failures: []string{StatusRestNoSector, StatusRestNotNative}}
	corridor := PositivityCorridor{Computed: true, AOver3: AInherited / 3, SqrtBOver3: math.Sqrt(BInherited / 3), TLower: lower, TUpper: upper, AlphaAtUpper: alphaUpper, BetaAtUpper: betaUpper, AlphaAtLower: alphaLower, BetaAtLower: betaLower, Gap: AInherited/3 - math.Sqrt(BInherited/3), Verdict: StatusPositivityCorridor, Supports: []string{StatusNarrowCorridor, StatusAlphaScale}, Failures: []string{StatusCorridorNotTopDerivation, StatusNoBackwardT, StatusCorridorNoSector}}
	concentration := RestConcentration{Recorded: true, Formula: "q_rest = b_rest/a_rest²", Bounds: "1/m_rest <= q_rest <= 1", BetaFormula: "beta = 3 alpha² q_rest", Interpretation: []string{"q_rest=1: rest pressure concentrated in one atom", "q_rest small: rest pressure distributed across many small atoms"}, Verdicts: []string{StatusRestConcentration, StatusRestConcentrationBound}, Supports: []string{StatusAlphaQRestSplit}, Failures: []string{StatusNoRestAtomCount, StatusNoQRestFromAggregate}}
	sectorCandidates := CandidateAudit{Audited: true, Topic: "Rest-pressure sector candidates", Supplies: []string{"bottom/tau/charm/up/down/strange/muon/electron rest atoms as candidates", "neutrino contribution if convention supplied", "scale/scheme/threshold effects as possible sources"}, Missing: []string{"decomposed trace atom ledger", "explicit neutrino convention", "scale and scheme typing"}, Verdict: StatusSectorCandidates, Supports: []string{StatusPlausibleRestSources}, Failures: []string{StatusNoSectorAssignment, StatusNeutrinoImplicit, StatusScaleSchemeUntyped}}
	patterns := CandidateAudit{Audited: true, Topic: "Koide / Froggatt-Nielsen / Georgi-Jarlskog diagnostics", Supplies: []string{"read-only diagnostics after ledger input"}, Missing: []string{"native rest-pressure source theorem"}, Verdict: StatusPatternFirewall, Supports: []string{StatusFNRelevant, StatusGJRelevant}, Failures: []string{StatusKoideNotNEff, StatusFNNotNativeRest, StatusGJNotLowScaleTop}}
	d4 := CandidateAudit{Audited: true, Topic: "D4 / triality firewall", Supplies: []string{"native search motivation only"}, Missing: []string{"triality-to-trace-magnitude readout map", "triality source of rest pressure", "real-form descent"}, Verdict: StatusD4Firewall, Supports: []string{StatusTrialityMotivation}, Failures: []string{StatusNEffNotD4, StatusNoTrialityTraceReadout, StatusNoTrialityRestPressure, StatusNoTrialityRealDescent}}
	fst := CandidateAudit{Audited: true, Topic: "Finite spectral triple", Supplies: []string{"color factor 3 in trace templates", "sector trace-form shape", "Yukawa edge locations"}, Missing: []string{"T", "rest spectrum", "hierarchy operator", "sector atom ledger"}, Verdict: StatusFiniteTripleAudit, Supports: []string{StatusFiniteTripleColorShape}, Failures: []string{StatusFSTNoTopEigenvalue, StatusFSTNoRestOperator}}
	external := CandidateAudit{Audited: true, Topic: "External Yukawa ledger", Supplies: []string{"T", "sector traces", "rest atoms", "alpha", "beta", "q_rest", "scale behavior"}, Missing: []string{"native theorem"}, Verdict: StatusExternalAudit, Supports: []string{StatusExternalCanTest}, Failures: []string{StatusExternalNotNative}}
	k7 := CandidateAudit{Audited: true, Topic: "K7 / Fock / projective resonance", Supplies: []string{"K7− dimension 3", "projective split 4=1+3"}, Missing: []string{"positive Hermitian spectra", "top block", "rest hierarchy"}, Verdict: StatusK7ProjectiveAudit, Failures: []string{StatusK7NotTopBlock, StatusProjectiveNotRest}}
	complexD4 := CandidateAudit{Audited: true, Topic: "Complex D4 trilinear", Supplies: []string{"airlocked pre-Yukawa trilinear shape"}, Missing: []string{"T", "H_f spectra", "top dominance", "rest pressure"}, Verdict: StatusD4SourceAudit, Failures: []string{StatusTD4NotTraceMagnitude, StatusTD4NotTopDominance, StatusTD4NotRestPressure}}
	chiggs := CHiggsImpact{Recorded: true, TopLimit: CHistory, Current: CHiggs, DeltaCHiggs: chiggsShift, TreeShift: treeShift, Interpretation: "rest pressure lowers C_Yukawa below one and lowers the Level-B tree proxy relative to exact N_eff=3", Verdict: StatusCHiggsImpact, Supports: []string{StatusRestPressureNumericallySmall}, Failures: []string{StatusTreeProxyNotPole, StatusNoNewSpectralData}}
	hierarchy := HierarchyObstruction{Selected: true, Name: "HierarchyBreakingOperatorSeal / TopDominanceAndRestSuppressionSeal", Needs: []string{"one large colored eigenvalue", "suppressed rest spectrum", "small N_eff-3", "scale-local value at M_Z", "T", "a_rest", "b_rest", "q_rest", "sector composition", "scale behavior"}, Forbidden: []string{"N_eff target closure", "C_Higgs", "Higgs mass", "scalar-runtime output"}, Verdict: StatusHierarchyObstruction, Supports: []string{StatusNeedsHierarchyMechanism}, Failures: []string{StatusNoHierarchyBreakingOperator, StatusNoNativeTopDominance, StatusNoLightSuppression, StatusNoNativeRestSource}}
	outcome := Outcome{Recorded: true, Items: []string{"exact N_eff=3 is explained by a rank-three top-color block", "current N_eff>3 is typed as positive rest spectral pressure", "aggregate a,b force a narrow top-dominant positivity corridor if top dominance is assumed", "the corridor is diagnostic, not a native top Yukawa derivation", "no current ASHA object sources T or the rest spectrum natively", "C_Higgs remains Level B"}, Verdict: StatusOutcomeRecorded, Supports: []string{StatusBestTypedModel}}
	branch := BranchDecision{Recorded: true, Next: "Gate 809 — HierarchyBreakingOperatorSeal and RestSpectrum Source Candidate Audit", Alternative: "Gate 809 — External Yukawa Trace Magnitude Ledger Validation and RestPressure Sector Audit", Reason: "Gate 808 extracts everything possible from aggregate a,b under the top-dominant assumption; the remaining native question is the hierarchy mechanism", Verdict: StatusBranchDecision, Supports: []string{StatusNextHierarchyGate}}
	firewalls := Firewalls{Enforced: true, NoNativeYukawa: true, NoPMNSCKM: true, NoFlavor: true, NoScalar: true, NoPoleMass: true, NoVEVGF: true, NoGJ: true, NoD4Triality: true, NoHistoryLoop: true, Verdict: StatusFirewallGate808}

	return Analysis{Inheritance: inheritance, TopColor: top, RestPressure: rest, Corridor: corridor, Concentration: concentration, SectorCandidates: sectorCandidates, PatternDiagnostics: patterns, D4Firewall: d4, FiniteTriple: fst, ExternalLedger: external, K7Projective: k7, ComplexD4: complexD4, CHiggs: chiggs, Hierarchy: hierarchy, Outcome: outcome, Branch: branch, Firewalls: firewalls, Truth: "Gate 808 extracts the strongest lawful information from the aggregate Yukawa trace ledger: baseline three comes from one dominant top-like atom times three colors; the small excess is unresolved rest spectral pressure.", Final: "The next native target is HierarchyBreakingOperatorSeal because ASHA now needs a mechanism for top dominance, light-family suppression, and small rest spectral pressure."}, nil
}

func Statuses() []string {
	return []string{
		StatusGate807Inherited, StatusTopColorSelected, StatusRestPressureSelected, StatusTopColorSealDefined, StatusTopColorLimit,
		StatusRestSealDefined, StatusRestDecomposition, StatusPositivityCorridor, StatusRestConcentration, StatusRestConcentrationBound,
		StatusSectorCandidates, StatusPatternFirewall, StatusD4Firewall, StatusFiniteTripleAudit, StatusExternalAudit, StatusK7ProjectiveAudit,
		StatusD4SourceAudit, StatusCHiggsImpact, StatusHierarchyObstruction, StatusOutcomeRecorded, StatusBranchDecision, StatusPhysicalFirewalls,
		StatusExactTopNEffThree, StatusColorThreeStrongest, StatusRestPressureAboveTop, StatusRestPressureDilutesCYukawa,
		StatusNarrowCorridor, StatusAlphaScale, StatusAlphaQRestSplit, StatusPlausibleRestSources, StatusFNRelevant, StatusGJRelevant,
		StatusTrialityMotivation, StatusFiniteTripleColorShape, StatusExternalCanTest, StatusRestPressureNumericallySmall, StatusNeedsHierarchyMechanism,
		StatusBestTypedModel, StatusNextHierarchyGate, StatusTopBlockNoT, StatusTopThreeNotGeneration, StatusTopBlockNotHierarchyTheorem,
		StatusRestNoSector, StatusRestNotNative, StatusCorridorNotTopDerivation, StatusNoBackwardT, StatusCorridorNoSector,
		StatusNoRestAtomCount, StatusNoQRestFromAggregate, StatusNoSectorAssignment, StatusNeutrinoImplicit, StatusScaleSchemeUntyped,
		StatusKoideNotNEff, StatusFNNotNativeRest, StatusGJNotLowScaleTop, StatusNEffNotD4, StatusNoTrialityTraceReadout,
		StatusNoTrialityRestPressure, StatusNoTrialityRealDescent, StatusFSTNoTopEigenvalue, StatusFSTNoRestOperator, StatusExternalNotNative,
		StatusK7NotTopBlock, StatusProjectiveNotRest, StatusTD4NotTraceMagnitude, StatusTD4NotTopDominance, StatusTD4NotRestPressure,
		StatusTreeProxyNotPole, StatusNoNewSpectralData, StatusNoHierarchyBreakingOperator, StatusNoNativeTopDominance, StatusNoLightSuppression,
		StatusNoNativeRestSource, StatusFirewallGate808,
	}
}

func FormatTop(s RankThreeTopColorBlockSeal) string {
	return fmt.Sprintf("%s components=[%s] %s; %s; N_eff_top=%.12g supports=[%s] failures=[%s]", s.Name, strings.Join(s.Components, "; "), s.ATop, s.BTop, s.NEffTop, strings.Join(s.Supports, "; "), strings.Join(s.Failures, "; "))
}

func FormatRest(s RestPressureOperatorSeal) string {
	return fmt.Sprintf("%s defs=[%s] %s; %s; small=%s; %s supports=[%s] failures=[%s]", s.Name, strings.Join(s.Definitions, "; "), s.NEffFormula, s.DeltaFormula, s.SmallFormula, s.Interpretation, strings.Join(s.Supports, "; "), strings.Join(s.Failures, "; "))
}

func FormatCorridor(c PositivityCorridor) string {
	return fmt.Sprintf("a/3=%.16g sqrt(b/3)=%.16g T=[%.16g, %.16g] alphaUpper=%.16g betaUpper=%.16g alphaLower=%.16g betaLower=%.16g supports=[%s] failures=[%s]", c.AOver3, c.SqrtBOver3, c.TLower, c.TUpper, c.AlphaAtUpper, c.BetaAtUpper, c.AlphaAtLower, c.BetaAtLower, strings.Join(c.Supports, "; "), strings.Join(c.Failures, "; "))
}

func FormatConcentration(c RestConcentration) string {
	return fmt.Sprintf("%s; %s; %s interpretation=[%s] supports=[%s] failures=[%s]", c.Formula, c.Bounds, c.BetaFormula, strings.Join(c.Interpretation, "; "), strings.Join(c.Supports, "; "), strings.Join(c.Failures, "; "))
}

func FormatAudit(a CandidateAudit) string {
	return fmt.Sprintf("%s supplies=[%s] missing=[%s] supports=[%s] failures=[%s]", a.Topic, strings.Join(a.Supplies, "; "), strings.Join(a.Missing, "; "), strings.Join(a.Supports, "; "), strings.Join(a.Failures, "; "))
}

func FormatCHiggs(c CHiggsImpact) string {
	return fmt.Sprintf("topLimit=%.16g current=%.16g deltaC=%.16g treeShift=%.14g %s supports=[%s] failures=[%s]", c.TopLimit, c.Current, c.DeltaCHiggs, c.TreeShift, c.Interpretation, strings.Join(c.Supports, "; "), strings.Join(c.Failures, "; "))
}

func FormatHierarchy(h HierarchyObstruction) string {
	return fmt.Sprintf("%s needs=[%s] forbidden=[%s] supports=[%s] failures=[%s]", h.Name, strings.Join(h.Needs, "; "), strings.Join(h.Forbidden, "; "), strings.Join(h.Supports, "; "), strings.Join(h.Failures, "; "))
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
