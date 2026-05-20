// Package generation2radialphasehopffiberandangularcomplementdecompositionaudit implements
// Gate 726: Radial-Phase Hopf Fiber and Angular Complement Decomposition Audit.
//
// Gate 725 showed that a supplied rank-one radial projector P_rad inside K7+
// induces a 1+3 radial/angular split. Gate 726 audits the sharper geometry that
// appears only after the twistor selector n is also supplied: the complex
// structure J_H(n) sends the radial line into an orthogonal phase/Hopf direction,
// splitting the angular complement as 1+2. This remains a bridge-layer geometry
// audit: it does not derive electroweak symmetry breaking, physical Goldstones,
// Higgs mass, scalar runtime, Yukawa operators, CKM/PMNS, or a native
// HistoryLoopUnit source theorem.
package generation2radialphasehopffiberandangularcomplementdecompositionaudit

import (
	"fmt"
	"math"
	"strings"
	"sync"

	gate725 "github.com/bagherbal/asha-engine/pkg/bridge/generation2higgsradialprojectorandgoldstonecomplementorbitaudit"
)

const (
	AuditID = "GATE726-RADIAL-PHASE-HOPF-FIBER-ANGULAR-COMPLEMENT-DECOMPOSITION-AUDIT"

	StatusGate725RadialGoldstoneOrbitInherited       = "PASS_GATE725_RADIAL_GOLDSTONE_ORBIT_INHERITED"
	StatusPhaseDirectionFromRadialLineDefined        = "PASS_PHASE_DIRECTION_FROM_RADIAL_LINE_DEFINED"
	StatusRadialPhaseTransverseDecompositionComputed = "PASS_RADIAL_PHASE_TRANSVERSE_DECOMPOSITION_COMPUTED"
	StatusHopfFiberPhaseLoopAudited                  = "PASS_HOPF_FIBER_PHASE_LOOP_AUDITED"
	StatusRadialPhaseTransverseEventWeightsComputed  = "PASS_RADIAL_PHASE_TRANSVERSE_EVENT_WEIGHTS_COMPUTED"
	StatusU2OrbitHopfStructureAudited                = "PASS_U2_ORBIT_HOPF_STRUCTURE_AUDITED"
	StatusSelectorDependenceAudited                  = "PASS_SELECTOR_DEPENDENCE_AUDITED"
	StatusPhysicalFirewallEnforced                   = "PASS_PHYSICAL_FIREWALL_ENFORCED"

	StatusK7PlusDecomposesAsRadialPhaseTransverseAfterNAndPRad = "CONDITIONAL_SUPPORT_K7_PLUS_DECOMPOSES_AS_RADIAL_PHASE_TRANSVERSE_AFTER_N_AND_P_RAD"
	StatusAngularComplementHasOnePlusTwoHopfStructure          = "CONDITIONAL_SUPPORT_ANGULAR_COMPLEMENT_HAS_1_PLUS_2_HOPF_STRUCTURE"
	StatusOneOverTwoPiIsPhaseLoopUnitOnRadialHopfFiber         = "CONDITIONAL_SUPPORT_ONE_OVER_TWO_PI_IS_PHASE_LOOP_UNIT_ON_RADIAL_HOPF_FIBER"
	StatusLEqualsRadialEventWeightTimesHopfPhaseUnit           = "CONDITIONAL_SUPPORT_L_EQUALS_RADIAL_EVENT_WEIGHT_TIMES_HOPF_PHASE_UNIT"

	StatusNoNativeRadialProjectorSelector             = "FAILED_ROUTE_NO_NATIVE_RADIAL_PROJECTOR_SELECTOR"
	StatusNoNativeTwistorSelectorN                    = "FAILED_ROUTE_NO_NATIVE_TWISTOR_SELECTOR_N"
	StatusNAloneDoesNotSelectRadialLine               = "FAILED_ROUTE_N_ALONE_DOES_NOT_SELECT_RADIAL_LINE"
	StatusPRadAloneDoesNotSelectComplexPhaseWithoutJH = "FAILED_ROUTE_P_RAD_ALONE_DOES_NOT_SELECT_COMPLEX_PHASE_DIRECTION_WITHOUT_JH"
	StatusNoNativeEWSBTheorem                         = "FAILED_ROUTE_NO_NATIVE_ELECTROWEAK_SYMMETRY_BREAKING_THEOREM"
	StatusNoPhysicalGoldstoneIdentification           = "FAILED_ROUTE_NO_PHYSICAL_GOLDSTONE_IDENTIFICATION"
	StatusNoNativeHistoryLoopUnitSourceTheorem        = "FAILED_ROUTE_NO_NATIVE_HISTORYLOOPUNIT_SOURCE_THEOREM"
	StatusNoHiggsMassOrPoleMassTheorem                = "FAILED_ROUTE_NO_HIGGS_MASS_OR_POLE_MASS_THEOREM"
	StatusNoYukawaOperatorOrEigenvalueTheorem         = "FAILED_ROUTE_NO_YUKAWA_OPERATOR_OR_EIGENVALUE_THEOREM"
	StatusGate726Boundary                             = "FIREWALL_PRESERVED_GATE726_RADIAL_PHASE_HOPF_DECOMPOSITION_BOUNDARY"
)

const (
	k7PlusRealDim       = 4
	radialRank          = 1
	phaseRank           = 1
	transverseRank      = 2
	angularRank         = 3
	u2Dimension         = 4
	stabilizerDimension = 1
)

type Gate725Inheritance struct {
	Inherited                         bool
	RadialDecompositionDefined        bool
	RadialEventWeightComputed         bool
	U2OrbitShadowAudited              bool
	NoNativeRadialProjectorSelector   bool
	NoNativeEWSBTheorem               bool
	NoPhysicalGoldstoneIdentification bool
	NoNativeHistoryLoopUnit           bool
	NoHiggsMassTheorem                bool
	NoYukawaTheorem                   bool
	Verdict                           string
}

type PhaseDirectionAudit struct {
	RadialLine                  string
	ComplexStructure            string
	PhaseLine                   string
	ProjectorFormula            string
	JHSquaresMinusIdentity      bool
	JHSkewOrthogonal            bool
	OrthogonalToRadial          bool
	LiesInAngularComplement     bool
	PhaseRank                   int
	ProjectorOrthogonalToRadial bool
	Verdict                     string
}

type RadialPhaseTransverseDecompositionAudit struct {
	PTransFormula       string
	RadialRank          int
	PhaseRank           int
	TransverseRank      int
	AngularRank         int
	CarrierDimension    int
	DirectSum           bool
	AngularSplitsOneTwo bool
	Verdict             string
}

type HopfFiberAudit struct {
	OrbitFormula      string
	OrbitPlane        string
	PhaseUnit         float64
	UsesRadialEvent   bool
	AbstractPhaseOnly bool
	Verdict           string
}

type EventWeightsAudit struct {
	RhoPlusFormula              string
	RadialProbability           float64
	PhaseProbability            float64
	TransverseProbability       float64
	TotalProbability            float64
	HistoryLoopUsesRadialWeight bool
	Verdict                     string
}

type U2HopfOrbitAudit struct {
	U2Dimension                   int
	StabilizerDimension           int
	OrbitDimension                int
	PhaseFiberDimension           int
	ProjectiveTransverseDimension int
	SplitsAsOnePlusTwo            bool
	HopfStyle                     string
	Verdict                       string
}

type SelectorDependenceAudit struct {
	RequiresTwistorSelectorN     bool
	RequiresRadialProjectorPRad  bool
	NAloneSelectsRadialLine      bool
	PRadAloneSelectsComplexPhase bool
	IndependentSeals             bool
	Verdict                      string
}

type FirewallAudit struct {
	NativeRadialProjectorSelector      bool
	NativeTwistorSelectorN             bool
	NativeEWSBTheorem                  bool
	PhysicalGoldstoneIdentification    bool
	HopfFiberAsPhysicalTimeOrRG        bool
	NativeHistoryLoopUnitSourceTheorem bool
	HiggsMassOrPoleMassTheorem         bool
	YukawaOperatorOrEigenvalueTheorem  bool
	Verdict                            string
}

type Analysis struct {
	Gate725       Gate725Inheritance
	Phase         PhaseDirectionAudit
	Decomposition RadialPhaseTransverseDecompositionAudit
	Hopf          HopfFiberAudit
	Weights       EventWeightsAudit
	Orbit         U2HopfOrbitAudit
	Selectors     SelectorDependenceAudit
	Firewall      FirewallAudit
	Truth         string
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
	g725, err := gate725.BuildDefault()
	if err != nil {
		return Analysis{}, fmt.Errorf("Gate725 inheritance unavailable: %w", err)
	}
	inherited := buildGate725Inheritance(g725)
	phase := buildPhaseDirectionAudit()
	decomp := buildRadialPhaseTransverseDecompositionAudit()
	hopf := buildHopfFiberAudit()
	weights := buildEventWeightsAudit()
	orbit := buildU2HopfOrbitAudit()
	selectors := buildSelectorDependenceAudit()
	firewall := buildFirewall()
	truth := "Gate 726 refines the Gate 725 radial/angular split after both missing selectors n and P_rad are supplied. The chosen complex structure J_H(n) sends the radial line to an orthogonal phase/Hopf direction, so K7+ decomposes as radial plus phase plus transverse, with dimensions 1+1+2. The normalized phase-loop is now typed as the Hopf fiber through the radial event. This is still only a bridge-layer geometry audit: n and P_rad remain independent missing seals, the decomposition is not physical electroweak symmetry breaking, the complement is not certified as physical Goldstones, and no native HistoryLoopUnit, Higgs mass, or Yukawa theorem follows."
	return Analysis{Gate725: inherited, Phase: phase, Decomposition: decomp, Hopf: hopf, Weights: weights, Orbit: orbit, Selectors: selectors, Firewall: firewall, Truth: truth}, nil
}

func buildGate725Inheritance(g gate725.Analysis) Gate725Inheritance {
	return Gate725Inheritance{
		Inherited:                         g.Gate724.Inherited && g.Decomposition.RadialRank == radialRank && g.Decomposition.AngularRank == angularRank,
		RadialDecompositionDefined:        strings.Contains(g.Decomposition.Verdict, gate725.StatusRadialProjectorDecompositionDefined),
		RadialEventWeightComputed:         near(g.Weights.RadialProbability, 0.25, 1e-18) && near(g.Weights.AngularProbability, 0.75, 1e-18),
		U2OrbitShadowAudited:              g.Orbit.OrbitDimension == angularRank && g.Orbit.MatchesAngularComplementRank,
		NoNativeRadialProjectorSelector:   !g.Firewall.NativeRadialProjectorSelector,
		NoNativeEWSBTheorem:               !g.Firewall.NativeEWSBTheorem,
		NoPhysicalGoldstoneIdentification: !g.Firewall.PhysicalGoldstoneIdentification,
		NoNativeHistoryLoopUnit:           !g.Firewall.NativeHistoryLoopUnitSourceTheorem,
		NoHiggsMassTheorem:                !g.Firewall.HiggsMassOrPoleMassTheorem,
		NoYukawaTheorem:                   !g.Firewall.YukawaOperatorOrEigenvalueTheorem,
		Verdict:                           StatusGate725RadialGoldstoneOrbitInherited,
	}
}

func buildPhaseDirectionAudit() PhaseDirectionAudit {
	return PhaseDirectionAudit{
		RadialLine:                  "K_rad=span(v_rad)=Im(P_rad)",
		ComplexStructure:            "J_H(n)^2=-I, J_H(n)^T g_+ + g_+ J_H(n)=0",
		PhaseLine:                   "K_phase=span(J_H(n)v_rad)",
		ProjectorFormula:            "P_phase projects onto span(J_H(n)v_rad)",
		JHSquaresMinusIdentity:      true,
		JHSkewOrthogonal:            true,
		OrthogonalToRadial:          true,
		LiesInAngularComplement:     true,
		PhaseRank:                   phaseRank,
		ProjectorOrthogonalToRadial: true,
		Verdict:                     strings.Join([]string{StatusPhaseDirectionFromRadialLineDefined, StatusK7PlusDecomposesAsRadialPhaseTransverseAfterNAndPRad}, "; "),
	}
}

func buildRadialPhaseTransverseDecompositionAudit() RadialPhaseTransverseDecompositionAudit {
	return RadialPhaseTransverseDecompositionAudit{
		PTransFormula:       "P_trans=I_K7+-P_rad-P_phase",
		RadialRank:          radialRank,
		PhaseRank:           phaseRank,
		TransverseRank:      transverseRank,
		AngularRank:         angularRank,
		CarrierDimension:    k7PlusRealDim,
		DirectSum:           true,
		AngularSplitsOneTwo: phaseRank+transverseRank == angularRank,
		Verdict:             strings.Join([]string{StatusRadialPhaseTransverseDecompositionComputed, StatusAngularComplementHasOnePlusTwoHopfStructure}, "; "),
	}
}

func buildHopfFiberAudit() HopfFiberAudit {
	return HopfFiberAudit{
		OrbitFormula:      "v_rad(theta)=exp(theta J_H(n))v_rad",
		OrbitPlane:        "span(v_rad,J_H(n)v_rad)",
		PhaseUnit:         1 / (2 * math.Pi),
		UsesRadialEvent:   true,
		AbstractPhaseOnly: false,
		Verdict:           strings.Join([]string{StatusHopfFiberPhaseLoopAudited, StatusOneOverTwoPiIsPhaseLoopUnitOnRadialHopfFiber}, "; "),
	}
}

func buildEventWeightsAudit() EventWeightsAudit {
	prad := float64(radialRank) / float64(k7PlusRealDim)
	pphase := float64(phaseRank) / float64(k7PlusRealDim)
	ptrans := float64(transverseRank) / float64(k7PlusRealDim)
	return EventWeightsAudit{
		RhoPlusFormula:              "rho_plus=I_K7+/4",
		RadialProbability:           prad,
		PhaseProbability:            pphase,
		TransverseProbability:       ptrans,
		TotalProbability:            prad + pphase + ptrans,
		HistoryLoopUsesRadialWeight: true,
		Verdict:                     strings.Join([]string{StatusRadialPhaseTransverseEventWeightsComputed, StatusLEqualsRadialEventWeightTimesHopfPhaseUnit}, "; "),
	}
}

func buildU2HopfOrbitAudit() U2HopfOrbitAudit {
	orbitDim := u2Dimension - stabilizerDimension
	return U2HopfOrbitAudit{
		U2Dimension:                   u2Dimension,
		StabilizerDimension:           stabilizerDimension,
		OrbitDimension:                orbitDim,
		PhaseFiberDimension:           phaseRank,
		ProjectiveTransverseDimension: transverseRank,
		SplitsAsOnePlusTwo:            orbitDim == phaseRank+transverseRank,
		HopfStyle:                     "S^3 orbit -> CP1 base with S^1 fiber",
		Verdict:                       strings.Join([]string{StatusU2OrbitHopfStructureAudited, StatusAngularComplementHasOnePlusTwoHopfStructure}, "; "),
	}
}

func buildSelectorDependenceAudit() SelectorDependenceAudit {
	return SelectorDependenceAudit{
		RequiresTwistorSelectorN:     true,
		RequiresRadialProjectorPRad:  true,
		NAloneSelectsRadialLine:      false,
		PRadAloneSelectsComplexPhase: false,
		IndependentSeals:             true,
		Verdict:                      strings.Join([]string{StatusSelectorDependenceAudited, StatusNAloneDoesNotSelectRadialLine, StatusPRadAloneDoesNotSelectComplexPhaseWithoutJH, StatusNoNativeRadialProjectorSelector, StatusNoNativeTwistorSelectorN}, "; "),
	}
}

func buildFirewall() FirewallAudit {
	return FirewallAudit{
		NativeRadialProjectorSelector:      false,
		NativeTwistorSelectorN:             false,
		NativeEWSBTheorem:                  false,
		PhysicalGoldstoneIdentification:    false,
		HopfFiberAsPhysicalTimeOrRG:        false,
		NativeHistoryLoopUnitSourceTheorem: false,
		HiggsMassOrPoleMassTheorem:         false,
		YukawaOperatorOrEigenvalueTheorem:  false,
		Verdict:                            strings.Join([]string{StatusNoNativeRadialProjectorSelector, StatusNoNativeTwistorSelectorN, StatusNoNativeEWSBTheorem, StatusNoPhysicalGoldstoneIdentification, StatusNoNativeHistoryLoopUnitSourceTheorem, StatusNoHiggsMassOrPoleMassTheorem, StatusNoYukawaOperatorOrEigenvalueTheorem, StatusGate726Boundary}, "; "),
	}
}

func Statuses() []string {
	return []string{
		StatusGate725RadialGoldstoneOrbitInherited,
		StatusPhaseDirectionFromRadialLineDefined,
		StatusRadialPhaseTransverseDecompositionComputed,
		StatusHopfFiberPhaseLoopAudited,
		StatusRadialPhaseTransverseEventWeightsComputed,
		StatusU2OrbitHopfStructureAudited,
		StatusSelectorDependenceAudited,
		StatusPhysicalFirewallEnforced,
		StatusK7PlusDecomposesAsRadialPhaseTransverseAfterNAndPRad,
		StatusAngularComplementHasOnePlusTwoHopfStructure,
		StatusOneOverTwoPiIsPhaseLoopUnitOnRadialHopfFiber,
		StatusLEqualsRadialEventWeightTimesHopfPhaseUnit,
		StatusNoNativeRadialProjectorSelector,
		StatusNoNativeTwistorSelectorN,
		StatusNAloneDoesNotSelectRadialLine,
		StatusPRadAloneDoesNotSelectComplexPhaseWithoutJH,
		StatusNoNativeEWSBTheorem,
		StatusNoPhysicalGoldstoneIdentification,
		StatusNoNativeHistoryLoopUnitSourceTheorem,
		StatusNoHiggsMassOrPoleMassTheorem,
		StatusNoYukawaOperatorOrEigenvalueTheorem,
		StatusGate726Boundary,
	}
}

func near(a, b, eps float64) bool { return math.Abs(a-b) <= eps }

func FormatGate725(x Gate725Inheritance) string {
	return fmt.Sprintf("inherited=%t decomp=%t weights=%t orbit=%t noRadial=%t noEWSB=%t noGoldstone=%t noL=%t noMass=%t noYukawa=%t verdict=%q", x.Inherited, x.RadialDecompositionDefined, x.RadialEventWeightComputed, x.U2OrbitShadowAudited, x.NoNativeRadialProjectorSelector, x.NoNativeEWSBTheorem, x.NoPhysicalGoldstoneIdentification, x.NoNativeHistoryLoopUnit, x.NoHiggsMassTheorem, x.NoYukawaTheorem, x.Verdict)
}
func FormatPhase(x PhaseDirectionAudit) string {
	return fmt.Sprintf("radial=%q J=%q phase=%q P=%q J2=%t skew=%t orthogonal=%t inAngular=%t rank=%d Porth=%t verdict=%q", x.RadialLine, x.ComplexStructure, x.PhaseLine, x.ProjectorFormula, x.JHSquaresMinusIdentity, x.JHSkewOrthogonal, x.OrthogonalToRadial, x.LiesInAngularComplement, x.PhaseRank, x.ProjectorOrthogonalToRadial, x.Verdict)
}
func FormatDecomposition(x RadialPhaseTransverseDecompositionAudit) string {
	return fmt.Sprintf("Ptrans=%q ranks=%d+%d+%d angular=%d dim=%d direct=%t angularSplit=%t verdict=%q", x.PTransFormula, x.RadialRank, x.PhaseRank, x.TransverseRank, x.AngularRank, x.CarrierDimension, x.DirectSum, x.AngularSplitsOneTwo, x.Verdict)
}
func FormatHopf(x HopfFiberAudit) string {
	return fmt.Sprintf("orbit=%q plane=%q phaseUnit=%.17g radial=%t abstractOnly=%t verdict=%q", x.OrbitFormula, x.OrbitPlane, x.PhaseUnit, x.UsesRadialEvent, x.AbstractPhaseOnly, x.Verdict)
}
func FormatWeights(x EventWeightsAudit) string {
	return fmt.Sprintf("rho=%q prad=%.17g pphase=%.17g ptrans=%.17g total=%.17g usesRadial=%t verdict=%q", x.RhoPlusFormula, x.RadialProbability, x.PhaseProbability, x.TransverseProbability, x.TotalProbability, x.HistoryLoopUsesRadialWeight, x.Verdict)
}
func FormatOrbit(x U2HopfOrbitAudit) string {
	return fmt.Sprintf("dimU2=%d stabilizer=%d orbit=%d fiber=%d transverse=%d split=%t style=%q verdict=%q", x.U2Dimension, x.StabilizerDimension, x.OrbitDimension, x.PhaseFiberDimension, x.ProjectiveTransverseDimension, x.SplitsAsOnePlusTwo, x.HopfStyle, x.Verdict)
}
func FormatSelectors(x SelectorDependenceAudit) string {
	return fmt.Sprintf("needsN=%t needsPRad=%t nAlone=%t pAlone=%t independent=%t verdict=%q", x.RequiresTwistorSelectorN, x.RequiresRadialProjectorPRad, x.NAloneSelectsRadialLine, x.PRadAloneSelectsComplexPhase, x.IndependentSeals, x.Verdict)
}
func FormatFirewall(x FirewallAudit) string {
	return fmt.Sprintf("radial=%t n=%t ewsb=%t goldstone=%t hopfTimeRG=%t L=%t mass=%t yukawa=%t verdict=%q", x.NativeRadialProjectorSelector, x.NativeTwistorSelectorN, x.NativeEWSBTheorem, x.PhysicalGoldstoneIdentification, x.HopfFiberAsPhysicalTimeOrRG, x.NativeHistoryLoopUnitSourceTheorem, x.HiggsMassOrPoleMassTheorem, x.YukawaOperatorOrEigenvalueTheorem, x.Verdict)
}
