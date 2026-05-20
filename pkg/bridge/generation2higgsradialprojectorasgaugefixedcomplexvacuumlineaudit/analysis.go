// Package generation2higgsradialprojectorasgaugefixedcomplexvacuumlineaudit implements
// Gate 761: Higgs Radial Projector as Gauge-Fixed Complex Vacuum Line Audit.
//
// Gate 760 identified P_rad as the highest-priority remaining scalar-runtime
// source-reduction target because L_Hopf=Tr_K7+(rho_plus[(1/(2*pi))P_rad])
// still depends on a supplied rank-one radial event. Gate 761 audits whether
// P_rad is better typed as a primitive real line or as a gauge-fixed radial
// representative inside a chosen complex Higgs vacuum line after the twistor
// selector n supplies J_H(n). This is a radial-projector typing and gauge-
// fixing audit only. It does not derive electroweak symmetry breaking, scalar
// runtime lambda, Higgs mass, pole mass, Yukawa operators, CKM/PMNS, flavor
// hierarchy, or a native HistoryLoopUnit theorem.
package generation2higgsradialprojectorasgaugefixedcomplexvacuumlineaudit

import (
	"fmt"
	"math"
	"strings"
	"sync"
)

const (
	AuditID = "GATE761-HIGGS-RADIAL-PROJECTOR-AS-GAUGE-FIXED-COMPLEX-VACUUM-LINE-AUDIT"

	StatusGate760MasterFormInherited                          = "PASS_GATE760_MASTER_FORM_INHERITED"
	StatusPRadPriorityInherited                               = "PASS_P_RAD_PRIORITY_INHERITED"
	StatusComplexStructureJHInherited                         = "PASS_COMPLEX_STRUCTURE_JH_INHERITED"
	StatusRealRadialAndPhaseDirectionsDefined                 = "PASS_REAL_RADIAL_AND_PHASE_DIRECTIONS_DEFINED"
	StatusComplexVacuumLineConstructedFromPRadAndJH           = "PASS_COMPLEX_VACUUM_LINE_CONSTRUCTED_FROM_P_RAD_AND_JH"
	StatusEventWeightsComputed                                = "PASS_EVENT_WEIGHTS_COMPUTED"
	StatusPRadTypedAsGaugeFixedRadialRepresentative           = "PASS_P_RAD_TYPED_AS_GAUGE_FIXED_RADIAL_REPRESENTATIVE"
	StatusScalarVacuumDirectionSealDecomposed                 = "PASS_SCALAR_VACUUM_DIRECTION_SEAL_DECOMPOSED"
	StatusU2HopfOrbitInterpretationRecorded                   = "PASS_U2_HOPF_ORBIT_INTERPRETATION_RECORDED"
	StatusSourceCandidateAuditCompleted                       = "PASS_SOURCE_CANDIDATE_AUDIT_COMPLETED"
	StatusHistoryLoopQuarterFactorInterpretationRefined       = "PASS_HISTORYLOOP_QUARTER_FACTOR_INTERPRETATION_REFINED"
	StatusPhysicalFirewallsEnforced                           = "PASS_PHYSICAL_FIREWALLS_ENFORCED"
	StatusPRadGaugeFixedRepresentativeInsideComplexLine       = "CONDITIONAL_SUPPORT_P_RAD_IS_GAUGE_FIXED_RADIAL_REPRESENTATIVE_INSIDE_COMPLEX_VACUUM_LINE"
	StatusOneOverFourRealRadialAmplitudeWeight                = "CONDITIONAL_SUPPORT_ONE_OVER_FOUR_IS_REAL_RADIAL_AMPLITUDE_EVENT_WEIGHT_NOT_COMPLEX_LINE_WEIGHT"
	StatusScalarVacuumDirectionSealSplits                     = "CONDITIONAL_SUPPORT_SCALAR_VACUUM_DIRECTION_SEAL_SPLITS_INTO_COMPLEX_LINE_PLUS_RADIAL_GAUGE_FIXING"
	StatusNoNativeComplexVacuumLineSelector                   = "FAILED_ROUTE_NO_NATIVE_COMPLEX_VACUUM_LINE_SELECTOR"
	StatusNoNativeRadialGaugeFixingSelector                   = "FAILED_ROUTE_NO_NATIVE_RADIAL_GAUGE_FIXING_SELECTOR"
	StatusNDoesNotSelectComplexVacuumLine                     = "FAILED_ROUTE_N_DOES_NOT_SELECT_COMPLEX_VACUUM_LINE"
	StatusRhoPlusDoesNotSelectVacuumLine                      = "FAILED_ROUTE_RHO_PLUS_DOES_NOT_SELECT_VACUUM_LINE"
	StatusNoNativeEWSBTheorem                                 = "FAILED_ROUTE_NO_NATIVE_ELECTROWEAK_SYMMETRY_BREAKING_THEOREM"
	StatusNoNativeHistoryLoopUnitTheorem                      = "FAILED_ROUTE_NO_NATIVE_HISTORYLOOPUNIT_THEOREM"
	StatusNoHiggsMassOrPoleMassTheorem                        = "FAILED_ROUTE_NO_HIGGS_MASS_OR_POLE_MASS_THEOREM"
	StatusNoYukawaOperatorOrEigenvalueTheorem                 = "FAILED_ROUTE_NO_YUKAWA_OPERATOR_OR_EIGENVALUE_THEOREM"
	StatusGate761RadialProjectorGaugeFixedComplexLineBoundary = "FIREWALL_PRESERVED_GATE761_RADIAL_PROJECTOR_GAUGE_FIXED_COMPLEX_LINE_BOUNDARY"
)

const (
	k7PlusRealDim       = 4
	k7PlusComplexDim    = 2
	realRadialRank      = 1
	realPhaseRank       = 1
	complexLineRealRank = 2
	complexLineRank     = 1
	transverseRealRank  = 2

	rhoPlusWeightPerRealLine = 1.0 / 4.0
	fullComplexLineWeight    = 1.0 / 2.0
	hopfPhaseUnit            = 1.0 / (2.0 * math.Pi)
	lHopf                    = 1.0 / (8.0 * math.Pi)
	complexLineLoopWeight    = 1.0 / (4.0 * math.Pi)
)

type Gate760Inheritance struct {
	Inherited                  bool
	MasterFormula              string
	HighestPrioritySeal        string
	HighestPriorityReason      string
	PRadPriorityInherited      bool
	LHopfDependsOnPRad         bool
	NativePRadSelector         bool
	NativeHistoryLoopTheorem   bool
	NativeScalarRuntimeTheorem bool
	Verdict                    string
}

type ComplexStructureInheritance struct {
	TwistorSelector        string
	ComplexStructure       string
	JHSquaresMinusIdentity bool
	JHSkewOrthogonal       bool
	K7PlusRealDimension    int
	K7PlusComplexDimension int
	K7PlusAsC2             bool
	Gate726SplitInherited  bool
	Verdict                string
}

type RadialPhaseDirections struct {
	RadialVector             string
	PRadFormula              string
	PhaseVector              string
	PPhaseFormula            string
	JHSkewOrthogonal         bool
	RadialPhaseOrthogonal    bool
	RadialRank               int
	PhaseRank                int
	PhaseInAngularComplement bool
	Defined                  bool
	Verdict                  string
}

type ComplexVacuumLine struct {
	Formula               string
	RealRank              int
	ComplexRank           int
	JInvariant            bool
	CommutesWithJH        bool
	InsideK7PlusC2        bool
	ContainsRadial        bool
	ContainsPhase         bool
	ConstructedFromPRadJH bool
	Verdict               string
}

type EventWeights struct {
	RhoPlusFormula          string
	RadialWeight            float64
	PhaseWeight             float64
	ComplexLineWeight       float64
	TransverseWeight        float64
	TotalWeight             float64
	LHopfFromRadialEvent    float64
	LoopFromComplexLine     float64
	ActiveHistoryUsesRadial bool
	ComplexLineTooLargeForL bool
	Verdict                 string
}

type GaugeFixingTyping struct {
	PrimitiveRealLineTyping        string
	RefinedTyping                  string
	GaugeFixingInterpretation      string
	PRadInsideComplexLine          bool
	PRadArbitraryPrimitiveLine     bool
	ComplexLineContainsHopfPhase   bool
	SealDecomposition              string
	ScalarVacuumDirectionSealSplit bool
	Verdict                        string
}

type SelectorDistinction struct {
	NRole                          string
	ComplexLineRole                string
	PRadRole                       string
	NSelectsComplexStructure       bool
	NSelectsVacuumLine             bool
	ComplexLineSelectsRadialGauge  bool
	PRadSelectsGaugeRepresentative bool
	ThreeChoicesDistinct           bool
	Verdict                        string
}

type U2HopfOrbitInterpretation struct {
	Socket                     string
	ComplexVacuumLineOrbit     string
	UnitRepresentativeOrbit    string
	HopfFibration              string
	CP1BasePoint               bool
	S1FiberGaugeRepresentative bool
	RealRadialAmplitudeAxis    bool
	RefinesGate726S3Orbit      bool
	Verdict                    string
}

type SourceCandidate struct {
	Name                      string
	SuppliesComplexVacuumLine bool
	SuppliesRadialGaugeFixing bool
	Reason                    string
}

type SourceCandidateAudit struct {
	Candidates                      []SourceCandidate
	Completed                       bool
	NativeComplexVacuumLineSelector bool
	NativeRadialGaugeFixingSelector bool
	RhoPlusSelectsLine              bool
	NSelectsLine                    bool
	QSelectsLine                    bool
	BoundaryScalarsSelectLine       bool
	FanoQuaternionicSelectsLine     bool
	Verdict                         string
}

type HistoryLoopImplication struct {
	ActiveObservable              string
	RadialQuarterWeight           float64
	FullComplexLineWeight         float64
	ActiveLoopUnit                float64
	FullComplexLineLoopUnit       float64
	ActiveUsesRealRadialAmplitude bool
	FullComplexLineRejected       bool
	QuarterFactorInterpretation   string
	Verdict                       string
}

type PhysicalFirewalls struct {
	PRadNativeVacuumTheorem           bool
	ComplexLineNativeEWSBTheorem      bool
	RadialGaugeFixingPhysicalEWSB     bool
	ComplexLineWeightActiveL          bool
	PRadHiggsMassTheorem              bool
	PRadYukawaTheorem                 bool
	LHopfNativeHistoryLoopTheorem     bool
	HiggsMassOrPoleMassTheorem        bool
	YukawaOperatorOrEigenvalueTheorem bool
	Audited                           bool
	Verdict                           string
}

type Analysis struct {
	Gate760     Gate760Inheritance
	Complex     ComplexStructureInheritance
	Directions  RadialPhaseDirections
	VacuumLine  ComplexVacuumLine
	Weights     EventWeights
	GaugeFixing GaugeFixingTyping
	Selectors   SelectorDistinction
	Orbit       U2HopfOrbitInterpretation
	SourceAudit SourceCandidateAudit
	HistoryLoop HistoryLoopImplication
	Firewalls   PhysicalFirewalls
	Truth       string
}

var cache struct {
	sync.Once
	a   Analysis
	err error
}

func BuildDefault() (Analysis, error) {
	cache.Once.Do(func() { cache.a, cache.err = Build() })
	return cache.a, cache.err
}

func Build() (Analysis, error) {
	gate760 := buildGate760Inheritance()
	complex := buildComplexStructureInheritance()
	directions := buildRadialPhaseDirections()
	vacuum := buildComplexVacuumLine()
	weights := buildEventWeights()
	if !finitePositive(weights.LHopfFromRadialEvent) || !near(weights.LHopfFromRadialEvent, lHopf, 1e-18) {
		return Analysis{}, fmt.Errorf("invalid Gate761 radial HistoryLoop event weight: got %g want %g", weights.LHopfFromRadialEvent, lHopf)
	}
	gauge := buildGaugeFixingTyping()
	selectors := buildSelectorDistinction()
	orbit := buildU2HopfOrbitInterpretation()
	source := buildSourceCandidateAudit()
	history := buildHistoryLoopImplication()
	firewalls := buildPhysicalFirewalls()
	truth := "Gate 761 refines the Gate760 P_rad pressure point. Once n supplies the complex structure J_H(n) on K7+, a supplied real radial vector v_rad determines a phase vector J_H(n)v_rad and hence a complex rank-one vacuum line Pi_vac_C=P_rad+P_phase. The active HistoryLoop unit uses only the gauge-fixed real radial amplitude event with no-bias weight 1/4, giving (1/(2*pi))(1/4)=1/(8*pi). The full complex line has weight 1/2 and would give 1/(4*pi), so it is not the active HistoryLoopUnit. P_rad is therefore better typed as a GaugeFixedRadialRepresentativeSeal inside a ComplexVacuumLineSeal, not as an arbitrary primitive real line. No native complex vacuum line selector, radial gauge-fixing selector, electroweak symmetry-breaking theorem, HistoryLoopUnit theorem, Higgs mass theorem, pole-mass theorem, or Yukawa theorem is derived."
	return Analysis{Gate760: gate760, Complex: complex, Directions: directions, VacuumLine: vacuum, Weights: weights, GaugeFixing: gauge, Selectors: selectors, Orbit: orbit, SourceAudit: source, HistoryLoop: history, Firewalls: firewalls, Truth: truth}, nil
}

func buildGate760Inheritance() Gate760Inheritance {
	return Gate760Inheritance{
		Inherited:                  true,
		MasterFormula:              "lambda_runtime_eff=(1/8)(3/N_eff)[1+L_Hopf(1-kappa_lambda_red)]",
		HighestPrioritySeal:        "P_rad",
		HighestPriorityReason:      "needed to source L_Hopf=Tr_K7+(rho_plus[(1/(2*pi))P_rad])",
		PRadPriorityInherited:      true,
		LHopfDependsOnPRad:         true,
		NativePRadSelector:         false,
		NativeHistoryLoopTheorem:   false,
		NativeScalarRuntimeTheorem: false,
		Verdict:                    strings.Join([]string{StatusGate760MasterFormInherited, StatusPRadPriorityInherited}, "; "),
	}
}

func buildComplexStructureInheritance() ComplexStructureInheritance {
	return ComplexStructureInheritance{
		TwistorSelector:        "n in S^2(K7-)",
		ComplexStructure:       "J_H(n) in End(K7+), J_H(n)^2=-I",
		JHSquaresMinusIdentity: true,
		JHSkewOrthogonal:       true,
		K7PlusRealDimension:    k7PlusRealDim,
		K7PlusComplexDimension: k7PlusComplexDim,
		K7PlusAsC2:             true,
		Gate726SplitInherited:  true,
		Verdict:                StatusComplexStructureJHInherited,
	}
}

func buildRadialPhaseDirections() RadialPhaseDirections {
	return RadialPhaseDirections{
		RadialVector:             "v_rad in K7+, ||v_rad||=1",
		PRadFormula:              "P_rad=v_rad v_rad^T",
		PhaseVector:              "v_phase=J_H(n)v_rad",
		PPhaseFormula:            "P_phase=v_phase v_phase^T",
		JHSkewOrthogonal:         true,
		RadialPhaseOrthogonal:    true,
		RadialRank:               realRadialRank,
		PhaseRank:                realPhaseRank,
		PhaseInAngularComplement: true,
		Defined:                  true,
		Verdict:                  strings.Join([]string{StatusRealRadialAndPhaseDirectionsDefined, StatusComplexStructureJHInherited}, "; "),
	}
}

func buildComplexVacuumLine() ComplexVacuumLine {
	return ComplexVacuumLine{
		Formula:               "Pi_vac_C=P_rad+P_phase=P_rad+J_H(n)P_radJ_H(n)^T",
		RealRank:              complexLineRealRank,
		ComplexRank:           complexLineRank,
		JInvariant:            true,
		CommutesWithJH:        true,
		InsideK7PlusC2:        true,
		ContainsRadial:        true,
		ContainsPhase:         true,
		ConstructedFromPRadJH: true,
		Verdict:               StatusComplexVacuumLineConstructedFromPRadAndJH,
	}
}

func buildEventWeights() EventWeights {
	radial := rhoPlusWeightPerRealLine
	phase := rhoPlusWeightPerRealLine
	complexLine := radial + phase
	transverse := fullComplexLineWeight
	return EventWeights{
		RhoPlusFormula:          "rho_plus=I_K7+/4",
		RadialWeight:            radial,
		PhaseWeight:             phase,
		ComplexLineWeight:       complexLine,
		TransverseWeight:        transverse,
		TotalWeight:             complexLine + transverse,
		LHopfFromRadialEvent:    hopfPhaseUnit * radial,
		LoopFromComplexLine:     hopfPhaseUnit * complexLine,
		ActiveHistoryUsesRadial: true,
		ComplexLineTooLargeForL: true,
		Verdict:                 strings.Join([]string{StatusEventWeightsComputed, StatusOneOverFourRealRadialAmplitudeWeight}, "; "),
	}
}

func buildGaugeFixingTyping() GaugeFixingTyping {
	return GaugeFixingTyping{
		PrimitiveRealLineTyping:        "arbitrary real rank-one line",
		RefinedTyping:                  "GaugeFixedRadialRepresentativeSeal",
		GaugeFixingInterpretation:      "P_rad chooses a phase gauge / radial representative inside Pi_vac_C",
		PRadInsideComplexLine:          true,
		PRadArbitraryPrimitiveLine:     false,
		ComplexLineContainsHopfPhase:   true,
		SealDecomposition:              "ScalarVacuumDirectionSeal=(ComplexVacuumLineSeal, RadialGaugeFixingSeal)",
		ScalarVacuumDirectionSealSplit: true,
		Verdict:                        strings.Join([]string{StatusPRadTypedAsGaugeFixedRadialRepresentative, StatusScalarVacuumDirectionSealDecomposed, StatusPRadGaugeFixedRepresentativeInsideComplexLine, StatusScalarVacuumDirectionSealSplits}, "; "),
	}
}

func buildSelectorDistinction() SelectorDistinction {
	return SelectorDistinction{
		NRole:                          "n selects J_H(n), the complex structure on K7+",
		ComplexLineRole:                "Pi_vac_C selects a complex rank-one vacuum line in K7+_J(n)",
		PRadRole:                       "P_rad selects a real radial representative inside Pi_vac_C after phase gauge fixing",
		NSelectsComplexStructure:       true,
		NSelectsVacuumLine:             false,
		ComplexLineSelectsRadialGauge:  false,
		PRadSelectsGaugeRepresentative: true,
		ThreeChoicesDistinct:           true,
		Verdict:                        strings.Join([]string{StatusNDoesNotSelectComplexVacuumLine, StatusNoNativeComplexVacuumLineSelector, StatusNoNativeRadialGaugeFixingSelector}, "; "),
	}
}

func buildU2HopfOrbitInterpretation() U2HopfOrbitInterpretation {
	return U2HopfOrbitInterpretation{
		Socket:                     "K7+_J(n) ~= C^2",
		ComplexVacuumLineOrbit:     "CP1",
		UnitRepresentativeOrbit:    "S3",
		HopfFibration:              "S1 -> S3 -> CP1",
		CP1BasePoint:               true,
		S1FiberGaugeRepresentative: true,
		RealRadialAmplitudeAxis:    true,
		RefinesGate726S3Orbit:      true,
		Verdict:                    StatusU2HopfOrbitInterpretationRecorded,
	}
}

func buildSourceCandidateAudit() SourceCandidateAudit {
	candidates := []SourceCandidate{
		{Name: "rho_plus", SuppliesComplexVacuumLine: false, SuppliesRadialGaugeFixing: false, Reason: "no-bias state on K7+; assigns weights but selects no line"},
		{Name: "n", SuppliesComplexVacuumLine: false, SuppliesRadialGaugeFixing: false, Reason: "selects J_H(n), not a CP1 vacuum line"},
		{Name: "q", SuppliesComplexVacuumLine: false, SuppliesRadialGaugeFixing: false, Reason: "normalizes phase charge / hypercharge interface, not the vacuum line"},
		{Name: "P_K7", SuppliesComplexVacuumLine: false, SuppliesRadialGaugeFixing: false, Reason: "selects full K7 support, not a line in K7+"},
		{Name: "boundary scalars", SuppliesComplexVacuumLine: false, SuppliesRadialGaugeFixing: false, Reason: "provide scalar coordinates but no vector in K7+"},
		{Name: "Fano/quaternionic structure", SuppliesComplexVacuumLine: false, SuppliesRadialGaugeFixing: false, Reason: "supplies twistor family and U(2) socket, not a vacuum line"},
	}
	return SourceCandidateAudit{
		Candidates:                      candidates,
		Completed:                       true,
		NativeComplexVacuumLineSelector: false,
		NativeRadialGaugeFixingSelector: false,
		RhoPlusSelectsLine:              false,
		NSelectsLine:                    false,
		QSelectsLine:                    false,
		BoundaryScalarsSelectLine:       false,
		FanoQuaternionicSelectsLine:     false,
		Verdict:                         strings.Join([]string{StatusSourceCandidateAuditCompleted, StatusRhoPlusDoesNotSelectVacuumLine, StatusNDoesNotSelectComplexVacuumLine, StatusNoNativeComplexVacuumLineSelector, StatusNoNativeRadialGaugeFixingSelector}, "; "),
	}
}

func buildHistoryLoopImplication() HistoryLoopImplication {
	return HistoryLoopImplication{
		ActiveObservable:              "R_Hopf=(1/(2*pi))P_rad",
		RadialQuarterWeight:           rhoPlusWeightPerRealLine,
		FullComplexLineWeight:         fullComplexLineWeight,
		ActiveLoopUnit:                lHopf,
		FullComplexLineLoopUnit:       complexLineLoopWeight,
		ActiveUsesRealRadialAmplitude: true,
		FullComplexLineRejected:       true,
		QuarterFactorInterpretation:   "1/4 is the gauge-fixed real radial amplitude event weight, not the full complex vacuum-line weight",
		Verdict:                       strings.Join([]string{StatusHistoryLoopQuarterFactorInterpretationRefined, StatusOneOverFourRealRadialAmplitudeWeight}, "; "),
	}
}

func buildPhysicalFirewalls() PhysicalFirewalls {
	return PhysicalFirewalls{
		PRadNativeVacuumTheorem:           false,
		ComplexLineNativeEWSBTheorem:      false,
		RadialGaugeFixingPhysicalEWSB:     false,
		ComplexLineWeightActiveL:          false,
		PRadHiggsMassTheorem:              false,
		PRadYukawaTheorem:                 false,
		LHopfNativeHistoryLoopTheorem:     false,
		HiggsMassOrPoleMassTheorem:        false,
		YukawaOperatorOrEigenvalueTheorem: false,
		Audited:                           true,
		Verdict: strings.Join([]string{
			StatusPhysicalFirewallsEnforced,
			StatusNoNativeComplexVacuumLineSelector,
			StatusNoNativeRadialGaugeFixingSelector,
			StatusNoNativeEWSBTheorem,
			StatusNoNativeHistoryLoopUnitTheorem,
			StatusNoHiggsMassOrPoleMassTheorem,
			StatusNoYukawaOperatorOrEigenvalueTheorem,
			StatusGate761RadialProjectorGaugeFixedComplexLineBoundary,
		}, "; "),
	}
}

func Statuses() []string {
	return []string{
		StatusGate760MasterFormInherited,
		StatusPRadPriorityInherited,
		StatusComplexStructureJHInherited,
		StatusRealRadialAndPhaseDirectionsDefined,
		StatusComplexVacuumLineConstructedFromPRadAndJH,
		StatusEventWeightsComputed,
		StatusPRadTypedAsGaugeFixedRadialRepresentative,
		StatusScalarVacuumDirectionSealDecomposed,
		StatusU2HopfOrbitInterpretationRecorded,
		StatusSourceCandidateAuditCompleted,
		StatusHistoryLoopQuarterFactorInterpretationRefined,
		StatusPhysicalFirewallsEnforced,
		StatusPRadGaugeFixedRepresentativeInsideComplexLine,
		StatusOneOverFourRealRadialAmplitudeWeight,
		StatusScalarVacuumDirectionSealSplits,
		StatusNoNativeComplexVacuumLineSelector,
		StatusNoNativeRadialGaugeFixingSelector,
		StatusNDoesNotSelectComplexVacuumLine,
		StatusRhoPlusDoesNotSelectVacuumLine,
		StatusNoNativeEWSBTheorem,
		StatusNoNativeHistoryLoopUnitTheorem,
		StatusNoHiggsMassOrPoleMassTheorem,
		StatusNoYukawaOperatorOrEigenvalueTheorem,
		StatusGate761RadialProjectorGaugeFixedComplexLineBoundary,
	}
}

func finitePositive(x float64) bool { return !math.IsNaN(x) && !math.IsInf(x, 0) && x > 0 }
func near(a, b, eps float64) bool   { return math.Abs(a-b) <= eps }

func FormatGate760(x Gate760Inheritance) string {
	return fmt.Sprintf("inherited=%t master=%q priority=%q reason=%q pInherited=%t depends=%t nativeP=%t nativeL=%t nativeRuntime=%t verdict=%q", x.Inherited, x.MasterFormula, x.HighestPrioritySeal, x.HighestPriorityReason, x.PRadPriorityInherited, x.LHopfDependsOnPRad, x.NativePRadSelector, x.NativeHistoryLoopTheorem, x.NativeScalarRuntimeTheorem, x.Verdict)
}
func FormatComplex(x ComplexStructureInheritance) string {
	return fmt.Sprintf("n=%q J=%q J2=%t skew=%t dimR=%d dimC=%d C2=%t gate726=%t verdict=%q", x.TwistorSelector, x.ComplexStructure, x.JHSquaresMinusIdentity, x.JHSkewOrthogonal, x.K7PlusRealDimension, x.K7PlusComplexDimension, x.K7PlusAsC2, x.Gate726SplitInherited, x.Verdict)
}
func FormatDirections(x RadialPhaseDirections) string {
	return fmt.Sprintf("vrad=%q Prad=%q vphase=%q Pphase=%q skew=%t orthogonal=%t ranks=%d+%d inAngular=%t defined=%t verdict=%q", x.RadialVector, x.PRadFormula, x.PhaseVector, x.PPhaseFormula, x.JHSkewOrthogonal, x.RadialPhaseOrthogonal, x.RadialRank, x.PhaseRank, x.PhaseInAngularComplement, x.Defined, x.Verdict)
}
func FormatVacuumLine(x ComplexVacuumLine) string {
	return fmt.Sprintf("formula=%q rankR=%d rankC=%d Jinvariant=%t commute=%t inC2=%t radial=%t phase=%t from=%t verdict=%q", x.Formula, x.RealRank, x.ComplexRank, x.JInvariant, x.CommutesWithJH, x.InsideK7PlusC2, x.ContainsRadial, x.ContainsPhase, x.ConstructedFromPRadJH, x.Verdict)
}
func FormatWeights(x EventWeights) string {
	return fmt.Sprintf("rho=%q prad=%.17g pphase=%.17g line=%.17g trans=%.17g total=%.17g Lrad=%.17g Lline=%.17g activeRadial=%t lineTooLarge=%t verdict=%q", x.RhoPlusFormula, x.RadialWeight, x.PhaseWeight, x.ComplexLineWeight, x.TransverseWeight, x.TotalWeight, x.LHopfFromRadialEvent, x.LoopFromComplexLine, x.ActiveHistoryUsesRadial, x.ComplexLineTooLargeForL, x.Verdict)
}
func FormatGaugeFixing(x GaugeFixingTyping) string {
	return fmt.Sprintf("primitive=%q refined=%q interpretation=%q inside=%t arbitrary=%t phase=%t seal=%q split=%t verdict=%q", x.PrimitiveRealLineTyping, x.RefinedTyping, x.GaugeFixingInterpretation, x.PRadInsideComplexLine, x.PRadArbitraryPrimitiveLine, x.ComplexLineContainsHopfPhase, x.SealDecomposition, x.ScalarVacuumDirectionSealSplit, x.Verdict)
}
func FormatSelectors(x SelectorDistinction) string {
	return fmt.Sprintf("n=%q line=%q prad=%q nJ=%t nLine=%t lineGauge=%t pGauge=%t distinct=%t verdict=%q", x.NRole, x.ComplexLineRole, x.PRadRole, x.NSelectsComplexStructure, x.NSelectsVacuumLine, x.ComplexLineSelectsRadialGauge, x.PRadSelectsGaugeRepresentative, x.ThreeChoicesDistinct, x.Verdict)
}
func FormatOrbit(x U2HopfOrbitInterpretation) string {
	return fmt.Sprintf("socket=%q lineOrbit=%q unitOrbit=%q hopf=%q cp1=%t s1Gauge=%t radialAxis=%t refines=%t verdict=%q", x.Socket, x.ComplexVacuumLineOrbit, x.UnitRepresentativeOrbit, x.HopfFibration, x.CP1BasePoint, x.S1FiberGaugeRepresentative, x.RealRadialAmplitudeAxis, x.RefinesGate726S3Orbit, x.Verdict)
}
func FormatSourceAudit(x SourceCandidateAudit) string {
	return fmt.Sprintf("candidates=%d complete=%t nativeLine=%t nativeGauge=%t rhoLine=%t nLine=%t qLine=%t boundaryLine=%t fanoLine=%t verdict=%q", len(x.Candidates), x.Completed, x.NativeComplexVacuumLineSelector, x.NativeRadialGaugeFixingSelector, x.RhoPlusSelectsLine, x.NSelectsLine, x.QSelectsLine, x.BoundaryScalarsSelectLine, x.FanoQuaternionicSelectsLine, x.Verdict)
}
func FormatHistoryLoop(x HistoryLoopImplication) string {
	return fmt.Sprintf("observable=%q quarter=%.17g lineWeight=%.17g Lactive=%.17g Lline=%.17g activeRadial=%t lineRejected=%t interp=%q verdict=%q", x.ActiveObservable, x.RadialQuarterWeight, x.FullComplexLineWeight, x.ActiveLoopUnit, x.FullComplexLineLoopUnit, x.ActiveUsesRealRadialAmplitude, x.FullComplexLineRejected, x.QuarterFactorInterpretation, x.Verdict)
}
func FormatFirewalls(x PhysicalFirewalls) string {
	return fmt.Sprintf("pVacuum=%t lineEWSB=%t gaugeEWSB=%t lineActiveL=%t pMass=%t pYukawa=%t nativeL=%t mass=%t yukawa=%t audited=%t verdict=%q", x.PRadNativeVacuumTheorem, x.ComplexLineNativeEWSBTheorem, x.RadialGaugeFixingPhysicalEWSB, x.ComplexLineWeightActiveL, x.PRadHiggsMassTheorem, x.PRadYukawaTheorem, x.LHopfNativeHistoryLoopTheorem, x.HiggsMassOrPoleMassTheorem, x.YukawaOperatorOrEigenvalueTheorem, x.Audited, x.Verdict)
}
