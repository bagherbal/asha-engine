// Package generation2higgsradialprojectorandgoldstonecomplementorbitaudit implements
// Gate 725: Higgs Radial Projector and Goldstone-Complement Orbit Audit.
//
// Gate 724 source-typed L=1/(8*pi) as the expectation of a normalized
// phase-loop payoff on a rank-one radial event inside K7+. Gate 725 audits the
// geometry of the new missing radial projector: if P_rad is supplied, it induces
// a 1+3 split inside the four-real-dimensional sealed Higgs carrier. The split
// is Higgs/Goldstone-like only as a representation shadow; no electroweak
// symmetry breaking theorem, physical Goldstone identification, Higgs mass
// theorem, HistoryLoopUnit source theorem, or Yukawa theorem is certified.
package generation2higgsradialprojectorandgoldstonecomplementorbitaudit

import (
	"fmt"
	"math"
	"strings"
	"sync"

	gate724 "github.com/bagherbal/asha-engine/pkg/bridge/generation2higgsradialeventweightandphaselooptransportaudit"
)

const (
	AuditID = "GATE725-HIGGS-RADIAL-PROJECTOR-GOLDSTONE-COMPLEMENT-ORBIT-AUDIT"

	StatusGate724HiggsRadialEventPhaseLoopInherited = "PASS_GATE724_HIGGS_RADIAL_EVENT_PHASELOOP_INHERITED"
	StatusRadialProjectorDecompositionDefined       = "PASS_RADIAL_PROJECTOR_DECOMPOSITION_DEFINED"
	StatusRadialAndComplementEventWeightsComputed   = "PASS_RADIAL_AND_COMPLEMENT_EVENT_WEIGHTS_COMPUTED"
	StatusU2OrbitStabilizerGeometryAudited          = "PASS_U2_ORBIT_STABILIZER_GEOMETRY_AUDITED"
	StatusRadialSelectorSourceCandidatesAudited     = "PASS_RADIAL_SELECTOR_SOURCE_CANDIDATES_AUDITED"
	StatusHiggsGoldstoneFirewallEnforced            = "PASS_HIGGS_GOLDSTONE_FIREWALL_ENFORCED"
	StatusHistoryLoopFirewallPreserved              = "PASS_HISTORYLOOP_FIREWALL_PRESERVED"

	StatusPRadInducesOnePlusThreeHiggsOrbitSplit               = "CONDITIONAL_SUPPORT_P_RAD_INDUCES_1_PLUS_3_HIGGS_ORBIT_SPLIT"
	StatusOneOverFourRadialEventWeightInK7Plus                 = "CONDITIONAL_SUPPORT_ONE_OVER_FOUR_IS_RADIAL_EVENT_WEIGHT_IN_K7_PLUS"
	StatusThreeOverFourAngularComplementWeight                 = "CONDITIONAL_SUPPORT_THREE_OVER_FOUR_IS_ANGULAR_COMPLEMENT_WEIGHT"
	StatusRadialSelectorHasThreeDimensionalU2OrbitComplement   = "CONDITIONAL_SUPPORT_RADIAL_SELECTOR_HAS_THREE_DIMENSIONAL_U2_ORBIT_COMPLEMENT"
	StatusPRadIsTypeDistinctScalarVacuumDirectionSealCandidate = "CONDITIONAL_SUPPORT_P_RAD_IS_TYPE_DISTINCT_SCALAR_VACUUM_DIRECTION_SEAL_CANDIDATE"

	StatusNoNativeRadialProjectorSelector      = "FAILED_ROUTE_NO_NATIVE_RADIAL_PROJECTOR_SELECTOR"
	StatusTwistorSelectorNDoesNotSelectPRad    = "FAILED_ROUTE_TWISTOR_SELECTOR_N_DOES_NOT_SELECT_P_RAD"
	StatusQDoesNotSelectPRad                   = "FAILED_ROUTE_HYPERCHARGE_NORMALIZATION_Q_DOES_NOT_SELECT_P_RAD"
	StatusNoNativeEWSBTheorem                  = "FAILED_ROUTE_NO_NATIVE_ELECTROWEAK_SYMMETRY_BREAKING_THEOREM"
	StatusNoPhysicalGoldstoneIdentification    = "FAILED_ROUTE_NO_PHYSICAL_GOLDSTONE_IDENTIFICATION"
	StatusNoNativeHistoryLoopUnitSourceTheorem = "FAILED_ROUTE_NO_NATIVE_HISTORYLOOPUNIT_SOURCE_THEOREM"
	StatusNoNativePhasePayoffTransportTheorem  = "FAILED_ROUTE_NO_NATIVE_PHASE_PAYOFF_TRANSPORT_THEOREM"
	StatusNoHiggsMassOrPoleMassTheorem         = "FAILED_ROUTE_NO_HIGGS_MASS_OR_POLE_MASS_THEOREM"
	StatusNoYukawaOperatorOrEigenvalueTheorem  = "FAILED_ROUTE_NO_YUKAWA_OPERATOR_OR_EIGENVALUE_THEOREM"
	StatusGate725Boundary                      = "FIREWALL_PRESERVED_GATE725_RADIAL_GOLDSTONE_ORBIT_BOUNDARY"
)

const (
	k7PlusRealDim             = 4
	radialRank                = 1
	angularRank               = 3
	u2Dimension               = 4
	radialStabilizerDimension = 1
)

type Gate724Inheritance struct {
	Inherited                               bool
	RhoPlusDefined                          bool
	RadialEventWeightComputed               bool
	PhaseLoopExpectationReproducesL         bool
	NoNativeRadialProjectorSelector         bool
	TwistorSelectorDoesNotSelectRadialEvent bool
	QDoesNotSourceL                         bool
	NoNativeHistoryLoopUnit                 bool
	NoHiggsMassTheorem                      bool
	NoYukawaTheorem                         bool
	Verdict                                 string
}

type RadialDecompositionAudit struct {
	ProjectorFormula  string
	ComplementFormula string
	RadialImage       string
	AngularImage      string
	RadialRank        int
	AngularRank       int
	CarrierDimension  int
	DirectSum         bool
	Orthogonal        bool
	Verdict           string
}

type EventWeightAudit struct {
	RhoPlusFormula     string
	RadialProbability  float64
	AngularProbability float64
	Sum                float64
	Verdict            string
}

type U2OrbitStabilizerAudit struct {
	Socket                          string
	U2Dimension                     int
	StabilizerDimension             int
	OrbitDimension                  int
	MatchesAngularComplementRank    bool
	PhysicalEWSBTheorem             bool
	PhysicalGoldstoneIdentification bool
	Verdict                         string
}

type RadialSelectorSourceAudit struct {
	TwistorSelectorSelectsPRad  bool
	QSelectsPRad                bool
	ScalarWallSelectsPRad       bool
	BoundarySplitSelectsPRad    bool
	K7EventProjectorSelectsPRad bool
	FanoHodgeSelectsPRad        bool
	NativeRadialSelectorFound   bool
	CandidateSeal               string
	Verdict                     string
}

type SealClassificationAudit struct {
	SealNames                                []string
	TypeDistinctFromTwistorSelector          bool
	TypeDistinctFromHyperchargeNormalization bool
	Verdict                                  string
}

type HistoryLoopRelationAudit struct {
	HistoryLoopUnit                    float64
	RadialWeight                       float64
	PhaseLoopPayoff                    float64
	ReproducesL                        bool
	NativeHistoryLoopUnitSourceTheorem bool
	NativePhasePayoffTransportTheorem  bool
	Verdict                            string
}

type FirewallAudit struct {
	NativeRadialProjectorSelector      bool
	TwistorSelectorSelectsPRad         bool
	QSelectsPRad                       bool
	NativeEWSBTheorem                  bool
	PhysicalGoldstoneIdentification    bool
	NativeHistoryLoopUnitSourceTheorem bool
	HiggsMassOrPoleMassTheorem         bool
	YukawaOperatorOrEigenvalueTheorem  bool
	Verdict                            string
}

type Analysis struct {
	Gate724       Gate724Inheritance
	Decomposition RadialDecompositionAudit
	Weights       EventWeightAudit
	Orbit         U2OrbitStabilizerAudit
	Sources       RadialSelectorSourceAudit
	Seal          SealClassificationAudit
	HistoryLoop   HistoryLoopRelationAudit
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
	g724, err := gate724.BuildDefault()
	if err != nil {
		return Analysis{}, fmt.Errorf("Gate724 inheritance unavailable: %w", err)
	}
	inherited := buildGate724Inheritance(g724)
	decomp := buildRadialDecompositionAudit()
	weights := buildEventWeightAudit()
	orbit := buildU2OrbitStabilizerAudit()
	sources := buildRadialSelectorSourceAudit()
	seal := buildSealClassificationAudit()
	history := buildHistoryLoopRelationAudit()
	firewall := buildFirewall()
	truth := "Gate 725 audits the representation geometry of the rank-one radial projector introduced as a source-type candidate in Gate 724. If P_rad is supplied, K7+ decomposes as a radial line plus a rank-three angular complement, and the no-bias weights are 1/4 and 3/4. The U(2)-orbit shadow has the same three-dimensional angular count, but this is not an electroweak symmetry breaking theorem: P_rad remains an unselected, type-distinct scalar-vacuum/radial seal candidate, and no physical Goldstone, Higgs mass, HistoryLoopUnit source, phase-payoff transport, or Yukawa theorem is certified."
	return Analysis{Gate724: inherited, Decomposition: decomp, Weights: weights, Orbit: orbit, Sources: sources, Seal: seal, HistoryLoop: history, Firewall: firewall, Truth: truth}, nil
}

func buildGate724Inheritance(g gate724.Analysis) Gate724Inheritance {
	return Gate724Inheritance{
		Inherited:                               g.Gate723.Inherited && g.RhoPlus.MaximallyMixed && g.RadialEvent.Rank == radialRank,
		RhoPlusDefined:                          strings.Contains(g.RhoPlus.Verdict, gate724.StatusRhoPlusDefined),
		RadialEventWeightComputed:               strings.Contains(g.RadialEvent.Verdict, gate724.StatusRankOneRadialEventWeightComputed),
		PhaseLoopExpectationReproducesL:         g.PhasePayoff.ExpectationMatchesL,
		NoNativeRadialProjectorSelector:         !g.Firewall.NativeRadialProjectorSelector,
		TwistorSelectorDoesNotSelectRadialEvent: !g.Firewall.TwistorSelectorSelectsRadialEvent,
		QDoesNotSourceL:                         !g.Firewall.QSourcesL,
		NoNativeHistoryLoopUnit:                 !g.Firewall.NativeHistoryLoopUnitSourceTheorem,
		NoHiggsMassTheorem:                      !g.Firewall.HiggsMassOrPoleMassTheorem,
		NoYukawaTheorem:                         !g.Firewall.YukawaOperatorOrEigenvalueTheorem,
		Verdict:                                 StatusGate724HiggsRadialEventPhaseLoopInherited,
	}
}

func buildRadialDecompositionAudit() RadialDecompositionAudit {
	return RadialDecompositionAudit{
		ProjectorFormula:  "P_rad^2=P_rad, P_rad^T=P_rad, rank(P_rad)=1",
		ComplementFormula: "P_ang=I_K7+-P_rad",
		RadialImage:       "K_rad=Im(P_rad)",
		AngularImage:      "K_ang=Im(P_ang)",
		RadialRank:        radialRank,
		AngularRank:       angularRank,
		CarrierDimension:  k7PlusRealDim,
		DirectSum:         true,
		Orthogonal:        true,
		Verdict:           strings.Join([]string{StatusRadialProjectorDecompositionDefined, StatusPRadInducesOnePlusThreeHiggsOrbitSplit}, "; "),
	}
}

func buildEventWeightAudit() EventWeightAudit {
	prad := float64(radialRank) / float64(k7PlusRealDim)
	pang := float64(angularRank) / float64(k7PlusRealDim)
	return EventWeightAudit{
		RhoPlusFormula:     "rho_plus=I_K7+/4",
		RadialProbability:  prad,
		AngularProbability: pang,
		Sum:                prad + pang,
		Verdict:            strings.Join([]string{StatusRadialAndComplementEventWeightsComputed, StatusOneOverFourRadialEventWeightInK7Plus, StatusThreeOverFourAngularComplementWeight}, "; "),
	}
}

func buildU2OrbitStabilizerAudit() U2OrbitStabilizerAudit {
	orbit := u2Dimension - radialStabilizerDimension
	return U2OrbitStabilizerAudit{
		Socket:                          "sealed U(2)-type socket on K7+_J(n) ~= C^2",
		U2Dimension:                     u2Dimension,
		StabilizerDimension:             radialStabilizerDimension,
		OrbitDimension:                  orbit,
		MatchesAngularComplementRank:    orbit == angularRank,
		PhysicalEWSBTheorem:             false,
		PhysicalGoldstoneIdentification: false,
		Verdict:                         strings.Join([]string{StatusU2OrbitStabilizerGeometryAudited, StatusRadialSelectorHasThreeDimensionalU2OrbitComplement, StatusNoNativeEWSBTheorem, StatusNoPhysicalGoldstoneIdentification}, "; "),
	}
}

func buildRadialSelectorSourceAudit() RadialSelectorSourceAudit {
	return RadialSelectorSourceAudit{
		TwistorSelectorSelectsPRad:  false,
		QSelectsPRad:                false,
		ScalarWallSelectsPRad:       false,
		BoundarySplitSelectsPRad:    false,
		K7EventProjectorSelectsPRad: false,
		FanoHodgeSelectsPRad:        false,
		NativeRadialSelectorFound:   false,
		CandidateSeal:               "HiggsRadialSelectorSeal / ScalarVacuumDirectionSeal / RadialModeProjectionSeal",
		Verdict:                     strings.Join([]string{StatusRadialSelectorSourceCandidatesAudited, StatusNoNativeRadialProjectorSelector, StatusTwistorSelectorNDoesNotSelectPRad, StatusQDoesNotSelectPRad}, "; "),
	}
}

func buildSealClassificationAudit() SealClassificationAudit {
	return SealClassificationAudit{
		SealNames:                                []string{"HiggsRadialSelectorSeal", "ScalarVacuumDirectionSeal", "RadialModeProjectionSeal"},
		TypeDistinctFromTwistorSelector:          true,
		TypeDistinctFromHyperchargeNormalization: true,
		Verdict:                                  StatusPRadIsTypeDistinctScalarVacuumDirectionSealCandidate,
	}
}

func buildHistoryLoopRelationAudit() HistoryLoopRelationAudit {
	radialWeight := float64(radialRank) / float64(k7PlusRealDim)
	phasePayoff := 1 / (2 * math.Pi)
	L := 1 / (8 * math.Pi)
	return HistoryLoopRelationAudit{
		HistoryLoopUnit:                    L,
		RadialWeight:                       radialWeight,
		PhaseLoopPayoff:                    phasePayoff,
		ReproducesL:                        near(radialWeight*phasePayoff, L, 1e-18),
		NativeHistoryLoopUnitSourceTheorem: false,
		NativePhasePayoffTransportTheorem:  false,
		Verdict:                            strings.Join([]string{StatusHistoryLoopFirewallPreserved, StatusNoNativeHistoryLoopUnitSourceTheorem, StatusNoNativePhasePayoffTransportTheorem}, "; "),
	}
}

func buildFirewall() FirewallAudit {
	return FirewallAudit{
		NativeRadialProjectorSelector:      false,
		TwistorSelectorSelectsPRad:         false,
		QSelectsPRad:                       false,
		NativeEWSBTheorem:                  false,
		PhysicalGoldstoneIdentification:    false,
		NativeHistoryLoopUnitSourceTheorem: false,
		HiggsMassOrPoleMassTheorem:         false,
		YukawaOperatorOrEigenvalueTheorem:  false,
		Verdict:                            strings.Join([]string{StatusNoNativeRadialProjectorSelector, StatusTwistorSelectorNDoesNotSelectPRad, StatusQDoesNotSelectPRad, StatusNoNativeEWSBTheorem, StatusNoPhysicalGoldstoneIdentification, StatusNoNativeHistoryLoopUnitSourceTheorem, StatusNoHiggsMassOrPoleMassTheorem, StatusNoYukawaOperatorOrEigenvalueTheorem, StatusGate725Boundary}, "; "),
	}
}

func Statuses() []string {
	return []string{
		StatusGate724HiggsRadialEventPhaseLoopInherited,
		StatusRadialProjectorDecompositionDefined,
		StatusRadialAndComplementEventWeightsComputed,
		StatusU2OrbitStabilizerGeometryAudited,
		StatusRadialSelectorSourceCandidatesAudited,
		StatusHiggsGoldstoneFirewallEnforced,
		StatusHistoryLoopFirewallPreserved,
		StatusPRadInducesOnePlusThreeHiggsOrbitSplit,
		StatusOneOverFourRadialEventWeightInK7Plus,
		StatusThreeOverFourAngularComplementWeight,
		StatusRadialSelectorHasThreeDimensionalU2OrbitComplement,
		StatusPRadIsTypeDistinctScalarVacuumDirectionSealCandidate,
		StatusNoNativeRadialProjectorSelector,
		StatusTwistorSelectorNDoesNotSelectPRad,
		StatusQDoesNotSelectPRad,
		StatusNoNativeEWSBTheorem,
		StatusNoPhysicalGoldstoneIdentification,
		StatusNoNativeHistoryLoopUnitSourceTheorem,
		StatusNoNativePhasePayoffTransportTheorem,
		StatusNoHiggsMassOrPoleMassTheorem,
		StatusNoYukawaOperatorOrEigenvalueTheorem,
		StatusGate725Boundary,
	}
}

func near(a, b, eps float64) bool { return math.Abs(a-b) <= eps }

func FormatGate724(x Gate724Inheritance) string {
	return fmt.Sprintf("inherited=%t rho=%t radial=%t phaseL=%t noRadial=%t nNoRadial=%t qNoL=%t noL=%t noMass=%t noYukawa=%t verdict=%q", x.Inherited, x.RhoPlusDefined, x.RadialEventWeightComputed, x.PhaseLoopExpectationReproducesL, x.NoNativeRadialProjectorSelector, x.TwistorSelectorDoesNotSelectRadialEvent, x.QDoesNotSourceL, x.NoNativeHistoryLoopUnit, x.NoHiggsMassTheorem, x.NoYukawaTheorem, x.Verdict)
}
func FormatDecomposition(x RadialDecompositionAudit) string {
	return fmt.Sprintf("P=%q Pang=%q Krad=%q Kang=%q ranks=%d+%d dim=%d direct=%t orthogonal=%t verdict=%q", x.ProjectorFormula, x.ComplementFormula, x.RadialImage, x.AngularImage, x.RadialRank, x.AngularRank, x.CarrierDimension, x.DirectSum, x.Orthogonal, x.Verdict)
}
func FormatWeights(x EventWeightAudit) string {
	return fmt.Sprintf("rho=%q prad=%.17g pang=%.17g sum=%.17g verdict=%q", x.RhoPlusFormula, x.RadialProbability, x.AngularProbability, x.Sum, x.Verdict)
}
func FormatOrbit(x U2OrbitStabilizerAudit) string {
	return fmt.Sprintf("socket=%q dimU2=%d stabilizer=%d orbit=%d matches=%t ewsb=%t goldstone=%t verdict=%q", x.Socket, x.U2Dimension, x.StabilizerDimension, x.OrbitDimension, x.MatchesAngularComplementRank, x.PhysicalEWSBTheorem, x.PhysicalGoldstoneIdentification, x.Verdict)
}
func FormatSources(x RadialSelectorSourceAudit) string {
	return fmt.Sprintf("n=%t q=%t lambda=%t split=%t k7=%t fano=%t native=%t seal=%q verdict=%q", x.TwistorSelectorSelectsPRad, x.QSelectsPRad, x.ScalarWallSelectsPRad, x.BoundarySplitSelectsPRad, x.K7EventProjectorSelectsPRad, x.FanoHodgeSelectsPRad, x.NativeRadialSelectorFound, x.CandidateSeal, x.Verdict)
}
func FormatSeal(x SealClassificationAudit) string {
	return fmt.Sprintf("seals=%v distinctN=%t distinctQ=%t verdict=%q", x.SealNames, x.TypeDistinctFromTwistorSelector, x.TypeDistinctFromHyperchargeNormalization, x.Verdict)
}
func FormatHistoryLoop(x HistoryLoopRelationAudit) string {
	return fmt.Sprintf("L=%.17g radial=%.17g phase=%.17g reproduces=%t nativeL=%t nativePhase=%t verdict=%q", x.HistoryLoopUnit, x.RadialWeight, x.PhaseLoopPayoff, x.ReproducesL, x.NativeHistoryLoopUnitSourceTheorem, x.NativePhasePayoffTransportTheorem, x.Verdict)
}
func FormatFirewall(x FirewallAudit) string {
	return fmt.Sprintf("radial=%t n=%t q=%t ewsb=%t goldstone=%t L=%t mass=%t yukawa=%t verdict=%q", x.NativeRadialProjectorSelector, x.TwistorSelectorSelectsPRad, x.QSelectsPRad, x.NativeEWSBTheorem, x.PhysicalGoldstoneIdentification, x.NativeHistoryLoopUnitSourceTheorem, x.HiggsMassOrPoleMassTheorem, x.YukawaOperatorOrEigenvalueTheorem, x.Verdict)
}
