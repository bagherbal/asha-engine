// Package generation2boundarytotracemagnituderestmapminimalityandnogoaudit implements
// Gate 814: BoundaryToTraceMagnitudeRestMap Minimality and No-Go Audit.
//
// Gate 814 follows Gate 813's boundary second-moment rest-pressure result. It
// separates the strong aggregate scalar closure Delta_N ≈ (9/5)s + 6ps² from
// the stronger missing construction: a noncircular map from boundary data to
// positive Yukawa trace magnitudes. The result is deliberately conservative:
// the current branch reaches R1 and partial R2, but not trace atoms, sectors, or
// native Yukawa operators.
package generation2boundarytotracemagnituderestmapminimalityandnogoaudit

import (
	"fmt"
	"math"
	"strings"
)

const (
	AuditID = "GATE814-BOUNDARY-TO-TRACE-MAGNITUDE-REST-MAP-MINIMALITY-NO-GO-AUDIT"

	NEff      = 3.0023273474722147
	DeltaN    = NEff - 3.0
	SBoundary = 0.0012924448188162962
	PBoundary = 7.0 / 72.0
	CYukawa   = 0.9992248188812008
	CHistory  = 1.038025177923625
	CHiggs    = 1.0372205204048603

	CLeadingNineFive = 9.0 / 5.0
	CSecondMomentSix = 6.0

	StatusGate813Inherited     = "PASS_GATE813_BOUNDARY_SECOND_MOMENT_REST_PRESSURE_INHERITED"
	StatusMapDefined           = "PASS_BOUNDARY_TO_TRACE_MAGNITUDE_REST_MAP_DEFINED"
	StatusTargetChain          = "PASS_TARGET_CHAIN_RECORDED"
	StatusNonCircularityReq    = "PASS_NONCIRCULARITY_REQUIREMENT_RECORDED"
	StatusMinimality           = "PASS_BOUNDARY_TO_TRACE_MAGNITUDE_REST_MAP_MINIMALITY_AUDITED"
	StatusExistingObjects      = "PASS_EXISTING_ASHA_OBJECTS_AUDITED_FOR_MAP_SOURCE"
	StatusBoundaryPairSource   = "PASS_BOUNDARY_PAIR_SOURCE_AUDITED"
	StatusK7Source             = "PASS_K7_EVENT_WEIGHT_SOURCE_AUDITED"
	StatusFiniteTripleSource   = "PASS_FINITE_TRIPLE_SOURCE_AUDITED"
	StatusHyperchargeSource    = "PASS_HYPERCHARGE_NORMALIZATION_SOURCE_AUDITED"
	StatusExternalLedgerSource = "PASS_EXTERNAL_LEDGER_SOURCE_AUDITED"
	StatusD4Source             = "PASS_D4_TRIALITY_SOURCE_REAUDITED"
	StatusChiralityFirewall    = "PASS_CHIRALITY_MASS_BRIDGE_FIREWALL_RECORDED"
	StatusClosureFirewall      = "PASS_DELTA_CLOSURE_VERSUS_TRACE_CONSTRUCTION_FIREWALL_DEFINED"
	StatusSpectrumNonUnique    = "PASS_POSITIVE_SPECTRUM_NON_UNIQUENESS_AUDITED"
	StatusLevelsDefined        = "PASS_REST_MAP_STATUS_LEVELS_DEFINED"
	StatusNonCircularityAudit  = "PASS_NONCIRCULARITY_AUDIT_DEFINED"
	StatusImpactRecorded       = "PASS_C_YUKAWA_AND_C_HIGGS_IMPACT_RECORDED"
	StatusOutcomeRecorded      = "PASS_OUTCOME_CLASSIFICATION_RECORDED"
	StatusBranchRecorded       = "PASS_BRANCH_DECISION_RECORDED"
	StatusPhysicalFirewalls    = "PASS_PHYSICAL_FIREWALLS_ENFORCED"

	StatusExactMissingObject       = "CONDITIONAL_SUPPORT_THIS_MAP_IS_THE_EXACT_MISSING_OBJECT_FOR_REDUCING_N_EFF"
	StatusAllSubobjectsNonCosmetic = "CONDITIONAL_SUPPORT_ALL_SUBOBJECTS_ARE_NONCOSMETIC"
	StatusBoundaryPairSupplies     = "CONDITIONAL_SUPPORT_BOUNDARY_PAIR_SUPPLIES_SPLIT_AND_DIMENSION_TWO"
	StatusK7SuppliesM2             = "CONDITIONAL_SUPPORT_K7_SUPPLIES_SECOND_RAW_BOUNDARY_MOMENT_WEIGHT"
	StatusFiniteTripleSupplies     = "CONDITIONAL_SUPPORT_FINITE_TRIPLE_SUPPLIES_COLOR_AND_TRACE_SHAPE"
	StatusHyperchargeSupplies      = "CONDITIONAL_SUPPORT_INVERSE_HYPERCHARGE_NORMALIZATION_SOURCES_THREE_OVER_FIVE_CANDIDATE"
	StatusSecondMomentStrong       = "CONDITIONAL_SUPPORT_SECOND_MOMENT_DELTA_CLOSURE_IS_NUMERICALLY_STRONG"
	StatusPositiveWeaker           = "CONDITIONAL_SUPPORT_POSITIVE_EXISTENCE_IS_WEAKER_THAN_SECTOR_LEDGER"
	StatusR1PartialR2              = "CONDITIONAL_SUPPORT_CURRENT_BOUNDARY_FN_BRANCH_REACHES_R1_AND_PARTIAL_R2"
	StatusTypedSources             = "CONDITIONAL_SUPPORT_COEFFICIENTS_HAVE_TYPED_SOURCE_CANDIDATES"
	StatusCertifiedMapWouldReduce  = "CONDITIONAL_SUPPORT_CERTIFIED_MAP_WOULD_REDUCE_N_EFF_SEAL_DEPENDENCE"
	StatusNextFreezeProtocol       = "CONDITIONAL_SUPPORT_NEXT_GATE_SHOULD_FREEZE_TESTABLE_BOUNDARY_FN_REST_PRESSURE_PROTOCOL"

	StatusMapNotCoefficientFit      = "FAILED_ROUTE_BOUNDARY_TO_TRACE_MAGNITUDE_REST_MAP_CANNOT_BE_COMPRESSED_TO_COEFFICIENT_FIT"
	StatusNoScaleWithoutS           = "FAILED_ROUTE_NO_REST_PRESSURE_SCALE_WITHOUT_BOUNDARY_SPLIT"
	StatusNoM2WithoutP              = "FAILED_ROUTE_NO_SECOND_MOMENT_CORRECTION_WITHOUT_K7_EVENT_WEIGHT"
	StatusNoNineFiveWithoutHyper    = "FAILED_ROUTE_NO_TYPED_SOURCE_FOR_NINE_OVER_FIVE_WITHOUT_HYPERCHARGE_NORMALIZATION"
	StatusNoTopBaselineWithoutColor = "FAILED_ROUTE_NO_TOP_COLOR_BASELINE_WITHOUT_COLOR_MULTIPLICITY"
	StatusNoSixWithoutBoundaryTwo   = "FAILED_ROUTE_NO_TYPED_SOURCE_FOR_C2_EQUALS_SIX_WITHOUT_BOUNDARY_PAIR_DIMENSION"
	StatusNoAlphaBetaWithoutTop     = "FAILED_ROUTE_NO_ALPHA_BETA_REST_DECOMPOSITION_WITHOUT_TOP_BLOCK_SELECTOR"
	StatusDirectClosureNoAtoms      = "FAILED_ROUTE_DIRECT_DELTA_N_CLOSURE_DOES_NOT_CONSTRUCT_TRACE_ATOMS"
	StatusNoYukawaWithoutPositive   = "FAILED_ROUTE_NO_YUKAWA_TRACE_MAGNITUDE_WITHOUT_POSITIVE_REST_SPECTRUM"
	StatusNoScaleLocalWithoutScheme = "FAILED_ROUTE_NO_SCALE_LOCAL_N_EFF_WITHOUT_SCALE_SCHEME_CONVENTION"
	StatusNoPredictionNoNonCirc     = "FAILED_ROUTE_NO_PREDICTION_STATUS_WITHOUT_NONCIRCULARITY_PROOF"
	StatusBoundaryPairNoMap         = "FAILED_ROUTE_BOUNDARY_PAIR_DOES_NOT_SUPPLY_TRACE_MAGNITUDE_REST_MAP"
	StatusK7NoYukawaAtoms           = "FAILED_ROUTE_K7_EVENT_WEIGHT_DOES_NOT_SUPPLY_YUKAWA_REST_ATOMS"
	StatusFiniteTripleNoMap         = "FAILED_ROUTE_FINITE_TRIPLE_DOES_NOT_SUPPLY_BOUNDARY_REST_PRESSURE_MAP"
	StatusHyperchargeNotTheorem     = "FAILED_ROUTE_INVERSE_HYPERCHARGE_NOT_YET_REST_PRESSURE_THEOREM"
	StatusExternalNotNative         = "FAILED_ROUTE_EXTERNAL_LEDGER_NOT_NATIVE_BOUNDARY_REST_MAP"
	StatusD4NotMap                  = "FAILED_ROUTE_D4_TRIALITY_NOT_BOUNDARY_TO_TRACE_MAGNITUDE_MAP"
	StatusChiralityNotSource        = "FAILED_ROUTE_CHIRALITY_MASS_BRIDGE_NOT_REST_PRESSURE_SOURCE"
	StatusClosureNotTraceTheorem    = "FAILED_ROUTE_DELTA_CLOSURE_ALONE_NOT_TRACE_MAGNITUDE_OPERATOR_THEOREM"
	StatusClosureNotSpectrum        = "FAILED_ROUTE_DELTA_CLOSURE_ALONE_NOT_YUKAWA_REST_SPECTRUM"
	StatusPositiveNoSectors         = "FAILED_ROUTE_POSITIVE_SPECTRUM_EXISTENCE_DOES_NOT_ASSIGN_BOTTOM_TAU_CHARM_OR_NEUTRINO_COMPONENTS"
	StatusQNoLedger                 = "FAILED_ROUTE_Q_REST_DOES_NOT_DETERMINE_REST_ATOM_LEDGER_UNIQUELY"
	StatusNotR3                     = "FAILED_ROUTE_BOUNDARY_FN_BRANCH_NOT_R3_TRACE_ATOM_CONSTRUCTION"
	StatusNotR4                     = "FAILED_ROUTE_BOUNDARY_FN_BRANCH_NOT_R4_YUKAWA_OPERATOR_THEOREM"
	StatusPostHocRisk               = "FAILED_ROUTE_TYPED_SOURCE_CANDIDATES_DO_NOT_REMOVE_POST_HOC_SELECTION_RISK"
	StatusNoPriorCoeffTheorem       = "FAILED_ROUTE_BOUNDARY_FN_MAP_NOT_PREDICTIVE_WITHOUT_PRIOR_COEFFICIENT_THEOREM"
	StatusNoCYukawaUpdate           = "FAILED_ROUTE_GATE814_DOES_NOT_UPDATE_C_YUKAWA"
	StatusCHiggsLevelB              = "FAILED_ROUTE_C_HIGGS_REMAINS_LEVEL_B"
	StatusFirewallGate814           = "FIREWALL_PRESERVED_GATE814_BOUNDARY_TO_TRACE_MAGNITUDE_REST_MAP_BOUNDARY"
)

type Inheritance struct {
	Gate813Inherited, SecondMomentSelected, PositiveCompatibilityInherited bool
	NEff, DeltaN, S, P, M2                                                 float64
	DeltaNB1, ResidualB1                                                   float64
	DeltaNB2, ResidualB2, ResidualImprovement                              float64
	Verdicts, Failures                                                     []string
}

type MapDefinition struct {
	Defined          bool
	Objects          []string
	TargetChain      string
	ForbiddenSources []string
	Verdicts         []string
	Supports         []string
}

type Minimality struct {
	Audited            bool
	RemovalFailures    map[string]string
	Verdicts           []string
	Supports, Failures []string
}

type SourceAudit struct {
	Audited            bool
	BoundaryPair       string
	K7EventWeight      string
	FiniteTriple       string
	Hypercharge        string
	ExternalLedger     string
	D4Triality         string
	ChiralityMass      string
	Verdicts           []string
	Supports, Failures []string
}

type ClosureFirewall struct {
	Defined            bool
	WeakAchievement    string
	StrongAchievement  string
	Verdicts           []string
	Supports, Failures []string
}

type SpectrumNonUniqueness struct {
	Audited            bool
	Examples           []string
	Verdicts           []string
	Supports, Failures []string
}

type CandidateMapLevels struct {
	Defined            bool
	Levels             []string
	CurrentStatus      string
	Verdicts           []string
	Supports, Failures []string
}

type NonCircularityAudit struct {
	Defined            bool
	Rules              []string
	Verdicts           []string
	Supports, Failures []string
}

type Impact struct {
	Recorded                                            bool
	NEffBoundaryB2, CYukawaBoundaryB2, CHiggsBoundaryB2 float64
	OfficialNEff, OfficialCYukawa, OfficialCHiggs       float64
	Verdicts                                            []string
	Supports, Failures                                  []string
}

type Outcome struct {
	Recorded bool
	Items    []string
	Verdicts []string
}

type BranchDecision struct {
	Recorded    bool
	Recommended string
	Alternative string
	Verdict     string
	Support     string
}

type Firewalls struct {
	Enforced                                                                                                       bool
	NoCoefficientPromotion, NoDeltaAsTraceTheorem, NoPositiveAsSectorLedger, NoPostHoc, NoLedgerUpdate, NoPoleMass bool
	Verdict                                                                                                        string
}

type Analysis struct {
	Inheritance    Inheritance
	Map            MapDefinition
	Minimality     Minimality
	Sources        SourceAudit
	Closure        ClosureFirewall
	Spectrum       SpectrumNonUniqueness
	Levels         CandidateMapLevels
	NonCircularity NonCircularityAudit
	Impact         Impact
	Outcome        Outcome
	Branch         BranchDecision
	Firewalls      Firewalls
	Truth          string
	Final          string
}

func M2(s float64) float64                       { return PBoundary * s * s }
func DeltaBoundaryB1(s float64) float64          { return CLeadingNineFive * s }
func DeltaBoundaryB2(s float64) float64          { return CLeadingNineFive*s + CSecondMomentSix*M2(s) }
func CYukawaFromNEff(nEff float64) float64       { return 3.0 / nEff }
func CHiggsFromNEff(nEff float64) float64        { return CYukawaFromNEff(nEff) * CHistory }
func ResidualImprovement(r1, r2 float64) float64 { return math.Abs(r1) / math.Abs(r2) }

func BuildDefault() (Analysis, error) {
	m2 := M2(SBoundary)
	deltaB1 := DeltaBoundaryB1(SBoundary)
	resB1 := DeltaN - deltaB1
	deltaB2 := DeltaBoundaryB2(SBoundary)
	resB2 := DeltaN - deltaB2
	improvement := ResidualImprovement(resB1, resB2)
	nEffB2 := 3.0 + deltaB2
	objects := []string{
		"boundary split coordinate s",
		"K7 event weight p",
		"hypercharge normalization 5/3",
		"color multiplicity 3",
		"boundary-pair dimension 2",
		"top-color block selector",
		"alpha rest-size map",
		"beta rest-quartic map",
		"q_rest concentration law",
		"positive rest spectrum construction",
		"trace atom validation",
		"scale/scheme convention",
		"noncircularity proof",
	}
	removals := map[string]string{
		"s":                    StatusNoScaleWithoutS,
		"p":                    StatusNoM2WithoutP,
		"5/3":                  StatusNoNineFiveWithoutHyper,
		"color 3":              StatusNoTopBaselineWithoutColor,
		"boundary pair 2":      StatusNoSixWithoutBoundaryTwo,
		"top selector":         StatusNoAlphaBetaWithoutTop,
		"alpha beta q maps":    StatusDirectClosureNoAtoms,
		"positive spectrum":    StatusNoYukawaWithoutPositive,
		"scale scheme":         StatusNoScaleLocalWithoutScheme,
		"noncircularity proof": StatusNoPredictionNoNonCirc,
	}
	return Analysis{
		Inheritance:    Inheritance{Gate813Inherited: true, SecondMomentSelected: true, PositiveCompatibilityInherited: true, NEff: NEff, DeltaN: DeltaN, S: SBoundary, P: PBoundary, M2: m2, DeltaNB1: deltaB1, ResidualB1: resB1, DeltaNB2: deltaB2, ResidualB2: resB2, ResidualImprovement: improvement, Verdicts: []string{StatusGate813Inherited}, Failures: []string{"FAILED_ROUTE_GATE813_DID_NOT_CONSTRUCT_NATIVE_TRACE_MAGNITUDE_REST_MAP"}},
		Map:            MapDefinition{Defined: true, Objects: objects, TargetChain: "(s,p,5/3,3,2) -> alpha(s,p), beta(s,p), q_rest(s,p) -> positive rest atoms r_j >= 0 -> N_eff -> C_Yukawa=3/N_eff", ForbiddenSources: []string{"N_eff", "C_Higgs", "lambda_runtime_eff", "m_H_tree_proxy", "m_H_pole", "observed Higgs mass", "fitted Yukawa atoms"}, Verdicts: []string{StatusMapDefined, StatusTargetChain, StatusNonCircularityReq}, Supports: []string{StatusExactMissingObject}},
		Minimality:     Minimality{Audited: true, RemovalFailures: removals, Verdicts: []string{StatusMinimality}, Supports: []string{StatusAllSubobjectsNonCosmetic}, Failures: []string{StatusMapNotCoefficientFit, StatusNoScaleWithoutS, StatusNoM2WithoutP, StatusNoNineFiveWithoutHyper, StatusNoTopBaselineWithoutColor, StatusNoSixWithoutBoundaryTwo, StatusNoAlphaBetaWithoutTop, StatusDirectClosureNoAtoms, StatusNoYukawaWithoutPositive, StatusNoScaleLocalWithoutScheme, StatusNoPredictionNoNonCirc}},
		Sources:        SourceAudit{Audited: true, BoundaryPair: "supplies s, xi_boundary, two-endpoint structure, and boundary-pair dimension 2; not alpha/beta/rest atoms", K7EventWeight: "supplies p=7/72 and M2=p s^2; not Yukawa rest atoms", FiniteTriple: "supplies color factor 3, Yukawa edge templates, and trace-form shape; not T or rest spectra", Hypercharge: "supplies 5/3 and hence inverse 3/5 candidate; not a rest-pressure theorem", ExternalLedger: "can test rest-pressure candidate but remains external", D4Triality: "airlocked and not a trace-magnitude map", ChiralityMass: "blocked by Gate 812 as rest-pressure source", Verdicts: []string{StatusExistingObjects, StatusBoundaryPairSource, StatusK7Source, StatusFiniteTripleSource, StatusHyperchargeSource, StatusExternalLedgerSource, StatusD4Source, StatusChiralityFirewall}, Supports: []string{StatusBoundaryPairSupplies, StatusK7SuppliesM2, StatusFiniteTripleSupplies, StatusHyperchargeSupplies}, Failures: []string{StatusBoundaryPairNoMap, StatusK7NoYukawaAtoms, StatusFiniteTripleNoMap, StatusHyperchargeNotTheorem, StatusExternalNotNative, StatusD4NotMap, StatusChiralityNotSource}},
		Closure:        ClosureFirewall{Defined: true, WeakAchievement: "Delta_N ≈ (9/5)s + 6p s²", StrongAchievement: "s,p -> alpha,beta,q_rest -> positive rest atoms -> a,b,N_eff", Verdicts: []string{StatusClosureFirewall}, Supports: []string{StatusSecondMomentStrong}, Failures: []string{StatusClosureNotTraceTheorem, StatusClosureNotSpectrum}},
		Spectrum:       SpectrumNonUniqueness{Audited: true, Examples: []string{"q_rest=1: one concentrated rest atom", "q_rest=1/m: m equal rest atoms", "0<q_rest<1: many nonunique multi-atom or continuous distributions"}, Verdicts: []string{StatusSpectrumNonUnique}, Supports: []string{StatusPositiveWeaker}, Failures: []string{StatusPositiveNoSectors, StatusQNoLedger}},
		Levels:         CandidateMapLevels{Defined: true, Levels: []string{"R0 coefficient resonance: 9/5=3×3/5 and 6=2×3", "R1 scalar closure: Delta_N=(9/5)s+6ps²", "R2 positive top/rest model: alpha,beta,q_rest with 0<=q<=1", "R3 trace atom construction: r_j(s,p)>=0", "R4 sector/Yukawa operator theorem: H_f=Y_f†Y_f"}, CurrentStatus: "R1 with partial R2 compatibility; not R3 or R4", Verdicts: []string{StatusLevelsDefined}, Supports: []string{StatusR1PartialR2}, Failures: []string{StatusNotR3, StatusNotR4}},
		NonCircularity: NonCircularityAudit{Defined: true, Rules: []string{"derive coefficients before comparison to N_eff", "do not tune against C_Higgs, Higgs mass, or N_eff residual", "map must predict Delta_N from boundary objects", "external Yukawa ledger must test sector/rest consequences"}, Verdicts: []string{StatusNonCircularityAudit}, Supports: []string{StatusTypedSources}, Failures: []string{StatusPostHocRisk, StatusNoPriorCoeffTheorem}},
		Impact:         Impact{Recorded: true, NEffBoundaryB2: nEffB2, CYukawaBoundaryB2: CYukawaFromNEff(nEffB2), CHiggsBoundaryB2: CHiggsFromNEff(nEffB2), OfficialNEff: NEff, OfficialCYukawa: CYukawa, OfficialCHiggs: CHiggs, Verdicts: []string{StatusImpactRecorded}, Supports: []string{StatusCertifiedMapWouldReduce}, Failures: []string{StatusNoCYukawaUpdate, StatusCHiggsLevelB}},
		Outcome:        Outcome{Recorded: true, Items: []string{"second-moment closure is the sharpest aggregate rest-pressure candidate", "compatible with abstract positive rest spectra", "existing ASHA objects source color 3, hypercharge 5/3, boundary-pair 2, and K7 event weight p", "no existing object supplies the full BoundaryToTraceMagnitudeRestMap", "branch is R1/partial R2, not R3/R4", "C_Higgs remains Level B"}, Verdicts: []string{StatusOutcomeRecorded}},
		Branch:         BranchDecision{Recorded: true, Recommended: "Gate 815 — Boundary-FN RestPressure Test Protocol and External Ledger Prediction Audit", Alternative: "Gate 815 — BoundaryToTraceMagnitudeRestMap Construction Candidate Audit", Verdict: StatusBranchRecorded, Support: StatusNextFreezeProtocol},
		Firewalls:      Firewalls{Enforced: true, NoCoefficientPromotion: true, NoDeltaAsTraceTheorem: true, NoPositiveAsSectorLedger: true, NoPostHoc: true, NoLedgerUpdate: true, NoPoleMass: true, Verdict: StatusFirewallGate814},
		Truth:          "Gate 814 shows that the second-moment boundary-FN closure is a sharp aggregate candidate, but not a trace-magnitude theorem.",
		Final:          "The branch must be frozen as a testable Level-B+ hypothesis until a BoundaryToTraceMagnitudeRestMap or external decomposed ledger supplies trace atoms.",
	}, nil
}

func Statuses() []string {
	return []string{StatusGate813Inherited, StatusMapDefined, StatusTargetChain, StatusNonCircularityReq, StatusMinimality, StatusExistingObjects, StatusBoundaryPairSource, StatusK7Source, StatusFiniteTripleSource, StatusHyperchargeSource, StatusExternalLedgerSource, StatusD4Source, StatusChiralityFirewall, StatusClosureFirewall, StatusSpectrumNonUnique, StatusLevelsDefined, StatusNonCircularityAudit, StatusImpactRecorded, StatusOutcomeRecorded, StatusBranchRecorded, StatusPhysicalFirewalls, StatusExactMissingObject, StatusAllSubobjectsNonCosmetic, StatusBoundaryPairSupplies, StatusK7SuppliesM2, StatusFiniteTripleSupplies, StatusHyperchargeSupplies, StatusSecondMomentStrong, StatusPositiveWeaker, StatusR1PartialR2, StatusTypedSources, StatusCertifiedMapWouldReduce, StatusNextFreezeProtocol, StatusMapNotCoefficientFit, StatusNoScaleWithoutS, StatusNoM2WithoutP, StatusNoNineFiveWithoutHyper, StatusNoTopBaselineWithoutColor, StatusNoSixWithoutBoundaryTwo, StatusNoAlphaBetaWithoutTop, StatusDirectClosureNoAtoms, StatusNoYukawaWithoutPositive, StatusNoScaleLocalWithoutScheme, StatusNoPredictionNoNonCirc, StatusBoundaryPairNoMap, StatusK7NoYukawaAtoms, StatusFiniteTripleNoMap, StatusHyperchargeNotTheorem, StatusExternalNotNative, StatusD4NotMap, StatusChiralityNotSource, StatusClosureNotTraceTheorem, StatusClosureNotSpectrum, StatusPositiveNoSectors, StatusQNoLedger, StatusNotR3, StatusNotR4, StatusPostHocRisk, StatusNoPriorCoeffTheorem, StatusNoCYukawaUpdate, StatusCHiggsLevelB, StatusFirewallGate814}
}

func FormatInheritance(a Inheritance) string {
	return fmt.Sprintf("DeltaN=%.16g s=%.16g p=%.16g M2=%.16g DeltaB1=%.16g R1=%.16g DeltaB2=%.16g R2=%.16g improvement=%.8g", a.DeltaN, a.S, a.P, a.M2, a.DeltaNB1, a.ResidualB1, a.DeltaNB2, a.ResidualB2, a.ResidualImprovement)
}

func FormatMap(a MapDefinition) string {
	return fmt.Sprintf("objects=[%s] target=%s forbidden=[%s]", strings.Join(a.Objects, "; "), a.TargetChain, strings.Join(a.ForbiddenSources, "; "))
}

func FormatMinimality(a Minimality) string {
	keys := []string{"s", "p", "5/3", "color 3", "boundary pair 2", "top selector", "alpha beta q maps", "positive spectrum", "scale scheme", "noncircularity proof"}
	parts := make([]string, 0, len(keys))
	for _, k := range keys {
		parts = append(parts, k+" -> "+a.RemovalFailures[k])
	}
	return strings.Join(parts, " | ")
}

func FormatSources(a SourceAudit) string {
	return strings.Join([]string{a.BoundaryPair, a.K7EventWeight, a.FiniteTriple, a.Hypercharge, a.ExternalLedger, a.D4Triality, a.ChiralityMass}, " | ")
}

func FormatLevels(a CandidateMapLevels) string {
	return strings.Join(a.Levels, " | ") + " || current=" + a.CurrentStatus
}

func FormatImpact(a Impact) string {
	return fmt.Sprintf("candidate NEffB2=%.16g CYukawaB2=%.16g CHiggsB2=%.16g official NEff=%.16g CYukawa=%.16g CHiggs=%.16g", a.NEffBoundaryB2, a.CYukawaBoundaryB2, a.CHiggsBoundaryB2, a.OfficialNEff, a.OfficialCYukawa, a.OfficialCHiggs)
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
