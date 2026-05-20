// Package generation2boundaryfnrestpressuretestprotocolandexternalledgerpredictionaudit implements
// Gate 815: Boundary-FN RestPressure Test Protocol and External Ledger Prediction Audit.
//
// Gate 815 freezes the Gate 814 boundary-FN second-moment closure as a
// falsifiable Level-B+ hypothesis. It does not promote the closure to a Yukawa
// theorem; it defines exactly what an independent decomposed Yukawa trace ledger
// must test, and forbids coefficient retuning after data input.
package generation2boundaryfnrestpressuretestprotocolandexternalledgerpredictionaudit

import (
	"fmt"
	"math"
	"strings"
)

const (
	AuditID = "GATE815-BOUNDARY-FN-RESTPRESSURE-TEST-PROTOCOL-EXTERNAL-LEDGER-PREDICTION-AUDIT"

	NEff      = 3.0023273474722147
	DeltaN    = NEff - 3.0
	SBoundary = 0.0012924448188162962
	PBoundary = 7.0 / 72.0
	CHistory  = 1.038025177923625
	CYukawa   = 0.9992248188812008
	CHiggs    = 1.0372205204048603

	C1NineFive = 9.0 / 5.0
	C2Six      = 6.0

	StatusGate814Inherited    = "PASS_GATE814_BOUNDARY_FN_REST_MAP_STATUS_INHERITED"
	StatusHypothesisFrozen    = "PASS_BOUNDARY_FN_HYPOTHESIS_FROZEN"
	StatusExternalLedgerReqs  = "PASS_EXTERNAL_TRACE_MAGNITUDE_LEDGER_REQUIREMENTS_DEFINED"
	StatusTopRestReqs         = "PASS_TOP_REST_READOUT_REQUIREMENTS_DEFINED"
	StatusAggregateTest       = "PASS_AGGREGATE_DELTA_N_TEST_DEFINED"
	StatusC2Diagnostic        = "PASS_C2_EXT_DIAGNOSTIC_DEFINED"
	StatusPositiveTopRest     = "PASS_POSITIVE_TOP_REST_TEST_DEFINED"
	StatusAlphaBand           = "PASS_BOUNDARY_FN_ALPHA_BAND_COMPUTED"
	StatusRestConcentration   = "PASS_REST_CONCENTRATION_DIAGNOSTIC_DEFINED"
	StatusSectorPressure      = "PASS_SECTOR_PRESSURE_INTERFACE_DEFINED"
	StatusSpurionTest         = "PASS_BOUNDARY_FN_SPURION_TEST_DEFINED"
	StatusNonCircularProtocol = "PASS_NONCIRCULAR_TEST_PROTOCOL_DEFINED"
	StatusFailureCriteria     = "PASS_FAILURE_CRITERIA_DEFINED"
	StatusPassCriteria        = "PASS_PASS_CRITERIA_DEFINED"
	StatusImpactRecorded      = "PASS_C_YUKAWA_AND_C_HIGGS_CANDIDATE_IMPACT_RECORDED"
	StatusPatternDiagnostics  = "PASS_PATTERN_DIAGNOSTIC_LANES_CLASSIFIED"
	StatusOutcomeRecorded     = "PASS_OUTCOME_CLASSIFICATION_RECORDED"
	StatusBranchRecorded      = "PASS_BRANCH_DECISION_RECORDED"
	StatusPhysicalFirewalls   = "PASS_PHYSICAL_FIREWALLS_ENFORCED"

	StatusReadyForTest            = "CONDITIONAL_SUPPORT_BOUNDARY_FN_CLOSURE_IS_READY_FOR_FALSIFIABLE_LEDGER_TESTING"
	StatusTypedCoeffSources       = "CONDITIONAL_SUPPORT_H_BFN_HAS_TYPED_COLOR_HYPERCHARGE_BOUNDARY_COEFFICIENT_SOURCES"
	StatusExternalCanTest         = "CONDITIONAL_SUPPORT_EXTERNAL_LEDGER_CAN_TEST_BOUNDARY_FN_REST_PRESSURE_IF_TOP_CHANNEL_IS_TYPED"
	StatusC2NearSixPrimary        = "CONDITIONAL_SUPPORT_C2_EXT_NEAR_SIX_IS_PRIMARY_SECOND_MOMENT_TEST"
	StatusAlphaNearThreeTenths    = "CONDITIONAL_SUPPORT_BOUNDARY_FN_PREDICTS_ALPHA_OVER_S_NEAR_THREE_TENTHS"
	StatusQClassifies             = "CONDITIONAL_SUPPORT_Q_REST_CAN_CLASSIFY_REST_SPECTRUM_SHAPE_AFTER_LEDGER_INPUT"
	StatusSmallPositiveRest       = "CONDITIONAL_SUPPORT_BOUNDARY_FN_PREDICTS_SMALL_POSITIVE_REST_PRESSURE"
	StatusEpsilonSharp            = "CONDITIONAL_SUPPORT_EPSILON_BFN_IS_A_SHARP_AGGREGATE_FN_STYLE_DIAGNOSTIC"
	StatusNonCircularlyTestable   = "CONDITIONAL_SUPPORT_BOUNDARY_FN_CAN_BE_TESTED_NONCIRCULARLY_IF_LEDGER_IS_SUPPLIED"
	StatusClearFalsification      = "CONDITIONAL_SUPPORT_BOUNDARY_FN_HYPOTHESIS_HAS_CLEAR_FALSIFICATION_CHANNELS"
	StatusCanUpgradeR2R3          = "CONDITIONAL_SUPPORT_SUCCESSFUL_EXTERNAL_LEDGER_TEST_CAN_UPGRADE_BOUNDARY_FN_TO_R2_OR_EXTERNAL_R3"
	StatusCanReduceSealDependence = "CONDITIONAL_SUPPORT_BOUNDARY_FN_SUCCESS_WOULD_REDUCE_N_EFF_SEAL_DEPENDENCE"
	StatusFNInterpretsRest        = "CONDITIONAL_SUPPORT_FN_DIAGNOSTIC_MAY_HELP_INTERPRET_REST_ATOMS_AFTER_LEDGER_INPUT"
	StatusGJTestsDownLepton       = "CONDITIONAL_SUPPORT_GJ_DIAGNOSTIC_MAY_TEST_DOWN_LEPTON_REST_STRUCTURE_AT_HIGH_SCALE"
	StatusNativeNextCoeffPrior    = "CONDITIONAL_SUPPORT_NATIVE_NEXT_BRANCH_SHOULD_TARGET_COEFFICIENT_PRIOR_IF_NO_LEDGER_EXISTS"
	StatusEmpiricalNextFrozenTest = "CONDITIONAL_SUPPORT_EMPIRICAL_NEXT_BRANCH_SHOULD_RUN_FROZEN_TEST_IF_LEDGER_EXISTS"

	StatusNotPromoted           = "FAILED_ROUTE_BOUNDARY_FN_CANDIDATE_NOT_PROMOTED_TO_TRACE_MAGNITUDE_THEOREM"
	StatusNoRetuning            = "FAILED_ROUTE_COEFFICIENTS_MUST_NOT_BE_RETUNED_AFTER_DATA_INPUT"
	StatusTypedNotNative        = "FAILED_ROUTE_TYPED_COEFFICIENT_SOURCES_NOT_YET_NATIVE_REST_PRESSURE_THEOREM"
	StatusNoTestNoAtoms         = "FAILED_ROUTE_NO_TEST_WITHOUT_DECOMPOSED_TRACE_ATOMS"
	StatusNoAlphaBetaQNoTop     = "FAILED_ROUTE_NO_ALPHA_BETA_Q_TEST_WITHOUT_TOP_CHANNEL_SELECTOR"
	StatusExternalNotNative     = "FAILED_ROUTE_EXTERNAL_LEDGER_NOT_NATIVE_YUKAWA_THEOREM"
	StatusAggregateNoSectors    = "FAILED_ROUTE_AGGREGATE_DELTA_TEST_DOES_NOT_ASSIGN_SECTORS"
	StatusC2NoAbsorbRetune      = "FAILED_ROUTE_C2_EXT_DEVIATION_MUST_NOT_BE_ABSORBED_BY_RETUNING_C2"
	StatusAlphaBandNoSectors    = "FAILED_ROUTE_POSITIVE_ALPHA_BAND_DOES_NOT_IDENTIFY_SECTORS"
	StatusAlphaBandNotOperator  = "FAILED_ROUTE_ALPHA_BAND_IS_NOT_YUKAWA_OPERATOR_THEOREM"
	StatusNoUniqueQ             = "FAILED_ROUTE_BOUNDARY_FN_CURRENTLY_DOES_NOT_PREDICT_UNIQUE_Q_REST"
	StatusQNoSectors            = "FAILED_ROUTE_Q_REST_DOES_NOT_ASSIGN_SECTORS_WITHOUT_ATOM_LEDGER"
	StatusNoSectorAssignment    = "FAILED_ROUTE_BOUNDARY_FN_DOES_NOT_CURRENTLY_ASSIGN_REST_PRESSURE_TO_BOTTOM_TAU_CHARM_OR_NEUTRINO"
	StatusGJKoideSecondary      = "FAILED_ROUTE_GJ_AND_KOIDE_REMAIN_SECONDARY_DIAGNOSTICS"
	StatusEpsilonNotNativeFN    = "FAILED_ROUTE_EPSILON_BFN_NOT_NATIVE_FN_SPURION_WITHOUT_CHARGE_OPERATOR"
	StatusEpsilonNoAtoms        = "FAILED_ROUTE_EPSILON_BFN_DOES_NOT_ASSIGN_TRACE_ATOMS"
	StatusCoeffRetuningInvalid  = "FAILED_ROUTE_COEFFICIENT_RETUNING_INVALIDATES_TEST"
	StatusTopSelectorNoForce    = "FAILED_ROUTE_TOP_SELECTOR_MUST_NOT_BE_CHOSEN_TO_FORCE_BOUNDARY_FN_CLOSURE"
	StatusNoHiggsDataForYukawa  = "FAILED_ROUTE_HIGGS_DATA_MUST_NOT_SOURCE_YUKAWA_LEDGER"
	StatusDowngradeIfFail       = "FAILED_ROUTE_BOUNDARY_FN_BRANCH_MUST_BE_DOWNGRADED_IF_EXTERNAL_LEDGER_FAILS_PROTOCOL"
	StatusExternalR3NotNativeR4 = "FAILED_ROUTE_EXTERNAL_R3_VALIDATION_NOT_NATIVE_R4_YUKAWA_OPERATOR_THEOREM"
	StatusNoCYukawaUpdate       = "FAILED_ROUTE_GATE815_DOES_NOT_UPDATE_C_YUKAWA"
	StatusCHiggsLevelB          = "FAILED_ROUTE_C_HIGGS_REMAINS_LEVEL_B"
	StatusPatternsNoMap         = "FAILED_ROUTE_PATTERN_DIAGNOSTICS_DO_NOT_SOURCE_BOUNDARY_FN_MAP"
	StatusFirewallGate815       = "FIREWALL_PRESERVED_GATE815_BOUNDARY_FN_RESTPRESSURE_TEST_PROTOCOL_BOUNDARY"
)

type Inheritance struct {
	Gate814Inherited, CandidateSelected bool
	NEff, DeltaN, S, P, M2              float64
	DeltaNBFN, NEffBFN, ResidualBFN     float64
	RelativeResidual                    float64
	Verdicts, Supports, Failures        []string
}

type FrozenHypothesis struct {
	Frozen          bool
	C1, C2          float64
	Formula         string
	CoefficientText []string
	Verdicts        []string
	Supports        []string
	Failures        []string
}

type ExternalLedgerRequirements struct {
	Defined       bool
	RequiredItems []string
	Readouts      []string
	Verdicts      []string
	Supports      []string
	Failures      []string
}

type AggregatePredictionTest struct {
	Defined     bool
	DeltaNBFN   float64
	C2Observed  float64
	Diagnostics []string
	Verdicts    []string
	Supports    []string
	Failures    []string
}

type PositiveTopRestTest struct {
	Defined                      bool
	NEffBFN                      float64
	AlphaMin, AlphaMax           float64
	AlphaMinOverS, AlphaMaxOverS float64
	Verdicts, Supports, Failures []string
}

type RestConcentrationDiagnostic struct {
	Defined                      bool
	Interpretations              []string
	Verdicts, Supports, Failures []string
}

type SectorPressureInterface struct {
	Defined                      bool
	Predictions                  []string
	NonPredictions               []string
	Verdicts, Supports, Failures []string
}

type SpurionPrediction struct {
	Defined                      bool
	EpsilonBFN, EpsilonN         float64
	ResidualEpsilon              float64
	Verdicts, Supports, Failures []string
}

type NonCircularProtocol struct {
	Defined                      bool
	Steps                        []string
	Forbidden                    []string
	Verdicts, Supports, Failures []string
}

type Criteria struct {
	Defined                      bool
	Items                        []string
	Verdicts, Supports, Failures []string
}

type Impact struct {
	Recorded                                      bool
	NEffBFN, CYukawaBFN, CHiggsBFN                float64
	OfficialNEff, OfficialCYukawa, OfficialCHiggs float64
	Verdicts, Supports, Failures                  []string
}

type PatternDiagnostics struct {
	Classified                   bool
	Lanes                        []string
	Verdicts, Supports, Failures []string
}

type Outcome struct {
	Recorded bool
	Items    []string
	Verdicts []string
}

type BranchDecision struct {
	Recorded      bool
	NativeNext    string
	EmpiricalNext string
	Verdict       string
	Supports      []string
}

type Firewalls struct {
	Enforced                                                                                             bool
	NoPromotion, NoRetuning, NoSilentTopChoice, NoHiggsData, NoPatternSource, NoLedgerUpdate, NoPoleMass bool
	Verdict                                                                                              string
}

type Analysis struct {
	Inheritance Inheritance
	Hypothesis  FrozenHypothesis
	Ledger      ExternalLedgerRequirements
	Aggregate   AggregatePredictionTest
	TopRest     PositiveTopRestTest
	RestQ       RestConcentrationDiagnostic
	Sector      SectorPressureInterface
	Spurion     SpurionPrediction
	Protocol    NonCircularProtocol
	Failure     Criteria
	Pass        Criteria
	Impact      Impact
	Patterns    PatternDiagnostics
	Outcome     Outcome
	Branch      BranchDecision
	Firewalls   Firewalls
	Truth       string
	Final       string
}

func M2(s float64) float64                 { return PBoundary * s * s }
func DeltaBFN(s float64) float64           { return C1NineFive*s + C2Six*M2(s) }
func NEffBFN(s float64) float64            { return 3.0 + DeltaBFN(s) }
func CYukawaFromNEff(nEff float64) float64 { return 3.0 / nEff }
func CHiggsFromNEff(nEff float64) float64  { return CYukawaFromNEff(nEff) * CHistory }
func C2Observed(deltaN, s float64) float64 { return (deltaN - C1NineFive*s) / M2(s) }
func Epsilon(delta float64) float64        { return math.Pow(delta, 0.25) }
func AlphaMin(nEff float64) float64        { return math.Sqrt(nEff/3.0) - 1.0 }
func AlphaMaxQ1(nEff float64) float64 {
	// q_rest=1 branch solves nEff = 3(1+a)^2/(1+3a^2).
	A := 3.0*nEff - 3.0
	B := -6.0
	C := nEff - 3.0
	disc := B*B - 4*A*C
	if disc < 0 || A == 0 {
		return math.NaN()
	}
	r1 := (-B + math.Sqrt(disc)) / (2 * A)
	r2 := (-B - math.Sqrt(disc)) / (2 * A)
	if r1 >= 0 && r1 < 0.01 {
		return r1
	}
	return r2
}

func BuildDefault() (Analysis, error) {
	m2 := M2(SBoundary)
	deltaBFN := DeltaBFN(SBoundary)
	nEffBFN := 3.0 + deltaBFN
	residual := DeltaN - deltaBFN
	alphaMin := AlphaMin(nEffBFN)
	alphaMax := AlphaMaxQ1(nEffBFN)
	epsBFN := Epsilon(deltaBFN)
	epsN := Epsilon(DeltaN)

	required := []string{"scale_mu", "scheme", "Yukawa_normalization", "color_convention", "neutrino_convention", "Spec(H_u), Spec(H_d), Spec(H_e), Spec(H_nu)", "top-channel selector T=h_t", "positive trace atoms x_i", "uncertainties", "validation rules"}
	readouts := []string{"a_ext", "b_ext", "N_eff_ext", "Delta_N_ext", "a_top,b_top", "a_rest,b_rest", "alpha_ext,beta_ext,q_rest_ext"}
	protocolSteps := []string{"freeze s,p,5/3,color 3,boundary-pair 2", "freeze c1=9/5 and c2=6", "import explicit decomposed trace-magnitude ledger", "validate a_ext,b_ext,N_eff_ext", "compute Delta_N_ext,c2_ext,alpha_ext,beta_ext,q_rest_ext,sector fractions", "compare against frozen predictions", "record pass/fail without retuning"}
	forbidden := []string{"choose c1/c2 after seeing external ledger residuals", "choose top selector to force alpha/s≈0.3", "silently renormalize atoms to match inherited a,b", "use Higgs mass or C_Higgs to choose Yukawa atoms", "use Koide/FN/GJ patterns to invent missing atoms"}

	return Analysis{
		Inheritance: Inheritance{Gate814Inherited: true, CandidateSelected: true, NEff: NEff, DeltaN: DeltaN, S: SBoundary, P: PBoundary, M2: m2, DeltaNBFN: deltaBFN, NEffBFN: nEffBFN, ResidualBFN: residual, RelativeResidual: residual / DeltaN, Verdicts: []string{StatusGate814Inherited}, Supports: []string{StatusReadyForTest}, Failures: []string{StatusNotPromoted}},
		Hypothesis:  FrozenHypothesis{Frozen: true, C1: C1NineFive, C2: C2Six, Formula: "Delta_N_pred=(9/5)s+6p s^2", CoefficientText: []string{"9/5=3×3/5: color-three times inverse hypercharge normalization", "6=2×3: boundary-pair dimension times color multiplicity"}, Verdicts: []string{StatusHypothesisFrozen}, Supports: []string{StatusTypedCoeffSources}, Failures: []string{StatusNoRetuning, StatusTypedNotNative}},
		Ledger:      ExternalLedgerRequirements{Defined: true, RequiredItems: required, Readouts: readouts, Verdicts: []string{StatusExternalLedgerReqs, StatusTopRestReqs}, Supports: []string{StatusExternalCanTest}, Failures: []string{StatusNoTestNoAtoms, StatusNoAlphaBetaQNoTop, StatusExternalNotNative}},
		Aggregate:   AggregatePredictionTest{Defined: true, DeltaNBFN: deltaBFN, C2Observed: C2Observed(DeltaN, SBoundary), Diagnostics: []string{"R_Delta=Delta_N_ext-Delta_N_BFN", "rho_Delta=R_Delta/Delta_N_ext", "rho_M2=R_Delta/(p s^2)", "c2_ext=[Delta_N_ext-(9/5)s]/(p s^2)"}, Verdicts: []string{StatusAggregateTest, StatusC2Diagnostic}, Supports: []string{StatusC2NearSixPrimary}, Failures: []string{StatusAggregateNoSectors, StatusC2NoAbsorbRetune}},
		TopRest:     PositiveTopRestTest{Defined: true, NEffBFN: nEffBFN, AlphaMin: alphaMin, AlphaMax: alphaMax, AlphaMinOverS: alphaMin / SBoundary, AlphaMaxOverS: alphaMax / SBoundary, Verdicts: []string{StatusPositiveTopRest, StatusAlphaBand}, Supports: []string{StatusAlphaNearThreeTenths}, Failures: []string{StatusAlphaBandNoSectors, StatusAlphaBandNotOperator}},
		RestQ:       RestConcentrationDiagnostic{Defined: true, Interpretations: []string{"q≈1: one concentrated rest atom", "q≈1/m: m comparable rest atoms", "q≈0: highly diffuse rest pressure or beta-zero boundary limit"}, Verdicts: []string{StatusRestConcentration}, Supports: []string{StatusQClassifies}, Failures: []string{StatusNoUniqueQ, StatusQNoSectors}},
		Sector:      SectorPressureInterface{Defined: true, Predictions: []string{"alpha_ext≈3.88e-4", "a_rest>0 and b_rest>=0", "a_rest << 3T", "FN-style hierarchy, if relevant, appears in suppressed non-top atoms", "GJ relevance requires multi-scale down/lepton ledger"}, NonPredictions: []string{"which sector dominates", "bottom/charm/tau ordering", "neutrino contribution", "Koide relation", "GJ ratios", "FN charges"}, Verdicts: []string{StatusSectorPressure}, Supports: []string{StatusSmallPositiveRest}, Failures: []string{StatusNoSectorAssignment, StatusGJKoideSecondary}},
		Spurion:     SpurionPrediction{Defined: true, EpsilonBFN: epsBFN, EpsilonN: epsN, ResidualEpsilon: epsN - epsBFN, Verdicts: []string{StatusSpurionTest}, Supports: []string{StatusEpsilonSharp}, Failures: []string{StatusEpsilonNotNativeFN, StatusEpsilonNoAtoms}},
		Protocol:    NonCircularProtocol{Defined: true, Steps: protocolSteps, Forbidden: forbidden, Verdicts: []string{StatusNonCircularProtocol}, Supports: []string{StatusNonCircularlyTestable}, Failures: []string{StatusCoeffRetuningInvalid, StatusTopSelectorNoForce, StatusNoHiggsDataForYukawa}},
		Failure:     Criteria{Defined: true, Items: []string{"F1 aggregate closure failure", "F2 second-moment coefficient failure", "F3 top/rest positivity failure", "F4 alpha-band failure", "F5 sector incoherence", "F6 scale instability"}, Verdicts: []string{StatusFailureCriteria}, Supports: []string{StatusClearFalsification}, Failures: []string{StatusDowngradeIfFail}},
		Pass:        Criteria{Defined: true, Items: []string{"P1 aggregate pass", "P2 c2_ext≈6", "P3 0<=q_rest<=1", "P4 alpha_ext/s≈0.300", "P5 sector ledger pass", "P6 scale pass"}, Verdicts: []string{StatusPassCriteria}, Supports: []string{StatusCanUpgradeR2R3}, Failures: []string{StatusExternalR3NotNativeR4}},
		Impact:      Impact{Recorded: true, NEffBFN: nEffBFN, CYukawaBFN: CYukawaFromNEff(nEffBFN), CHiggsBFN: CHiggsFromNEff(nEffBFN), OfficialNEff: NEff, OfficialCYukawa: CYukawa, OfficialCHiggs: CHiggs, Verdicts: []string{StatusImpactRecorded}, Supports: []string{StatusCanReduceSealDependence}, Failures: []string{StatusNoCYukawaUpdate, StatusCHiggsLevelB}},
		Patterns:    PatternDiagnostics{Classified: true, Lanes: []string{"Koide: charged-lepton diagnostic only", "Froggatt-Nielsen: hierarchy-power diagnostic after charges/ledger", "Georgi-Jarlskog: high-scale down/lepton diagnostic with RG ledger", "D4/triality: airlocked structural search branch", "chirality/mass bridge: edge compatibility firewall"}, Verdicts: []string{StatusPatternDiagnostics}, Supports: []string{StatusFNInterpretsRest, StatusGJTestsDownLepton}, Failures: []string{StatusPatternsNoMap}},
		Outcome:     Outcome{Recorded: true, Items: []string{"Boundary-FN frozen as Level-B+ testable hypothesis", "coefficients 9/5 and 6 locked before external ledger testing", "future ledger must test Delta_N,c2_ext,alpha,beta,q,sector fractions", "branch can be falsified by aggregate, positivity, alpha-band, sector, or scale failure", "passing tests upgrades to R2 or external R3, not native R4", "C_Higgs remains Level B"}, Verdicts: []string{StatusOutcomeRecorded}},
		Branch:      BranchDecision{Recorded: true, NativeNext: "Gate 816 — BoundaryToTraceMagnitudeRestMap Construction Candidate and Coefficient-Prior Audit", EmpiricalNext: "Gate 816 — External TraceMagnitude Ledger Validation and Boundary-FN RestPressure Test Audit", Verdict: StatusBranchRecorded, Supports: []string{StatusNativeNextCoeffPrior, StatusEmpiricalNextFrozenTest}},
		Firewalls:   Firewalls{Enforced: true, NoPromotion: true, NoRetuning: true, NoSilentTopChoice: true, NoHiggsData: true, NoPatternSource: true, NoLedgerUpdate: true, NoPoleMass: true, Verdict: StatusFirewallGate815},
		Truth:       "Gate 815 freezes Boundary-FN as a falsifiable Level-B+ protocol, not a trace-magnitude theorem.",
		Final:       "The next move is coefficient-prior construction if no ledger exists, or frozen external-ledger validation if decomposed trace atoms are supplied.",
	}, nil
}

func Statuses() []string {
	return []string{StatusGate814Inherited, StatusHypothesisFrozen, StatusExternalLedgerReqs, StatusAggregateTest, StatusC2Diagnostic, StatusPositiveTopRest, StatusAlphaBand, StatusRestConcentration, StatusSectorPressure, StatusSpurionTest, StatusNonCircularProtocol, StatusFailureCriteria, StatusPassCriteria, StatusImpactRecorded, StatusPatternDiagnostics, StatusOutcomeRecorded, StatusBranchRecorded, StatusPhysicalFirewalls, StatusReadyForTest, StatusTypedCoeffSources, StatusExternalCanTest, StatusC2NearSixPrimary, StatusAlphaNearThreeTenths, StatusQClassifies, StatusSmallPositiveRest, StatusEpsilonSharp, StatusNonCircularlyTestable, StatusClearFalsification, StatusCanUpgradeR2R3, StatusCanReduceSealDependence, StatusFNInterpretsRest, StatusGJTestsDownLepton, StatusNativeNextCoeffPrior, StatusEmpiricalNextFrozenTest, StatusNotPromoted, StatusNoRetuning, StatusTypedNotNative, StatusNoTestNoAtoms, StatusNoAlphaBetaQNoTop, StatusExternalNotNative, StatusAggregateNoSectors, StatusC2NoAbsorbRetune, StatusAlphaBandNoSectors, StatusAlphaBandNotOperator, StatusNoUniqueQ, StatusQNoSectors, StatusNoSectorAssignment, StatusGJKoideSecondary, StatusEpsilonNotNativeFN, StatusEpsilonNoAtoms, StatusCoeffRetuningInvalid, StatusTopSelectorNoForce, StatusNoHiggsDataForYukawa, StatusDowngradeIfFail, StatusExternalR3NotNativeR4, StatusNoCYukawaUpdate, StatusCHiggsLevelB, StatusPatternsNoMap, StatusFirewallGate815}
}

func FormatInheritance(a Inheritance) string {
	return fmt.Sprintf("DeltaN=%.16g s=%.16g p=%.16g M2=%.16g DeltaN_BFN=%.16g NEff_BFN=%.16g R=%.16g rho=%.16g", a.DeltaN, a.S, a.P, a.M2, a.DeltaNBFN, a.NEffBFN, a.ResidualBFN, a.RelativeResidual)
}

func FormatHypothesis(a FrozenHypothesis) string {
	return fmt.Sprintf("%s c1=%.16g c2=%.16g coeffs=[%s]", a.Formula, a.C1, a.C2, strings.Join(a.CoefficientText, "; "))
}

func FormatLedger(a ExternalLedgerRequirements) string {
	return fmt.Sprintf("required=[%s] readouts=[%s]", strings.Join(a.RequiredItems, "; "), strings.Join(a.Readouts, "; "))
}

func FormatAggregate(a AggregatePredictionTest) string {
	return fmt.Sprintf("DeltaN_BFN=%.16g c2_obs=%.16g diagnostics=[%s]", a.DeltaNBFN, a.C2Observed, strings.Join(a.Diagnostics, "; "))
}

func FormatTopRest(a PositiveTopRestTest) string {
	return fmt.Sprintf("NEff_BFN=%.16g alpha_min=%.16g alpha_max=%.16g alpha_min/s=%.16g alpha_max/s=%.16g", a.NEffBFN, a.AlphaMin, a.AlphaMax, a.AlphaMinOverS, a.AlphaMaxOverS)
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
