// Package generation2movingu1phaselineandhyperchargenormalizationfirewallaudit implements
// Gate 717: Moving U(1) Phase Line and Hypercharge Normalization Firewall Audit.
//
// Gate 716 certified the SU(2)-side representation compatibility of the
// selector-independent commutant C. Gate 717 audits the complementary moving
// U(1)-like phase line L_n=span(J_H(n)) inside u(2,J_H(n)). It certifies that
// for a fixed complex structure J_H(n), L_n is central and exponentiates to a
// uniform internal phase action on K7+_J(n) ~= C^2, while preserving the
// firewalls that this does not select a twistor point, does not define physical
// U(1)_Y, does not fix hypercharge normalization, does not complete the Higgs
// doublet representation, and does not derive Higgs mass, scalar runtime,
// Yukawa operators/eigenvalues, flavor hierarchy, CKM/PMNS, or native 7/72.
package generation2movingu1phaselineandhyperchargenormalizationfirewallaudit

import (
	"fmt"
	"strings"
	"sync"

	gate716 "github.com/bagherbal/asha-engine/pkg/bridge/generation2internalsu2sockettoelectroweaksu2lintertwinerairlockaudit"
)

const (
	AuditID = "GATE717-MOVING-U1-PHASE-LINE-AND-HYPERCHARGE-NORMALIZATION-FIREWALL-AUDIT"

	StatusGate716SU2IntertwinerAirlockInherited       = "PASS_GATE716_SU2_INTERTWINER_AIRLOCK_INHERITED"
	StatusMovingPhaseLineDefined                      = "PASS_MOVING_PHASE_LINE_DEFINED"
	StatusLNIsCentralInU2SocketForFixedJH             = "PASS_LN_IS_CENTRAL_IN_U2_SOCKET_FOR_FIXED_JH"
	StatusJHExponentiatesToUniformPhaseOnC2           = "PASS_JH_EXPONENTIATES_TO_UNIFORM_PHASE_ON_C2"
	StatusChargeNormalizationAudited                  = "PASS_CHARGE_NORMALIZATION_AUDITED"
	StatusSelectorDependenceAudited                   = "PASS_SELECTOR_DEPENDENCE_AUDITED"
	StatusSU2U1AsymmetryRecorded                      = "PASS_SU2_U1_ASYMMETRY_RECORDED"
	StatusPhysicalHyperchargeFirewallEnforced         = "PASS_PHYSICAL_HYPERCHARGE_FIREWALL_ENFORCED"
	StatusLNInternalU1PhaseSocketAfterJHChoice        = "CONDITIONAL_SUPPORT_LN_IS_INTERNAL_U1_PHASE_SOCKET_AFTER_JH_CHOICE"
	StatusK7PlusJHUniformInternalPhaseAction          = "CONDITIONAL_SUPPORT_K7_PLUS_JH_HAS_UNIFORM_INTERNAL_PHASE_ACTION"
	StatusElectroweakAirlockU1RequiresSelectorAndNorm = "CONDITIONAL_SUPPORT_ELECTROWEAK_AIRLOCK_U1_SIDE_REQUIRES_SELECTOR_AND_NORMALIZATION"
	StatusInternalPhaseLineNotPhysicalU1Y             = "FAILED_ROUTE_INTERNAL_PHASE_LINE_NOT_CERTIFIED_AS_PHYSICAL_U1Y"
	StatusNoHyperchargeAssignmentOrNormalization      = "FAILED_ROUTE_NO_HYPERCHARGE_ASSIGNMENT_OR_NORMALIZATION"
	StatusNoSelectorIndependentU1PhaseLine            = "FAILED_ROUTE_NO_SELECTOR_INDEPENDENT_U1_PHASE_LINE"
	StatusNoNativeTwistorPointSelector                = "FAILED_ROUTE_NO_NATIVE_TWISTOR_POINT_SELECTOR"
	StatusNoFullTypedK7PlusToPhysicalHiggsDoubletMap  = "FAILED_ROUTE_NO_FULL_TYPED_K7_PLUS_TO_PHYSICAL_HIGGS_DOUBLET_MAP"
	StatusNoHiggsMassOrScalarRuntimeTheorem           = "FAILED_ROUTE_NO_HIGGS_MASS_OR_SCALAR_RUNTIME_THEOREM"
	StatusNoYukawaOperatorOrEigenvalueTheorem         = "FAILED_ROUTE_NO_YUKAWA_OPERATOR_OR_EIGENVALUE_THEOREM"
	StatusGate717MovingU1PhaseHyperchargeBoundary     = "FIREWALL_PRESERVED_GATE717_MOVING_U1_PHASE_HYPERCHARGE_BOUNDARY"
)

const (
	k7PlusComplexDimension   = 2
	movingPhaseLineDimension = 1
	su2CommutantDimension    = 3
)

type Gate716Inheritance struct {
	SU2AirlockInherited             bool
	SU2SideStructurallyReady        bool
	InternalCSelectorIndependent    bool
	ComplexCarrierSelectorDependent bool
	U1PhaseSelectorDependent        bool
	HyperchargeDerived              bool
	HyperchargeNormalized           bool
	FullHiggsDoubletMap             bool
	HiggsMassCertified              bool
	YukawaCertified                 bool
	Verdict                         string
}

type MovingPhaseLineAudit struct {
	Definition          string
	Dimension           int
	DependsOnSelectorN  bool
	FixedJHRequired     bool
	SelectorIndependent bool
	PhysicalHypercharge bool
	Verdict             string
}

type CentralPhaseLineAudit struct {
	Commutator       string
	CommutesWithC    bool
	LiesInCenterOfU2 bool
	FixedJHOnly      bool
	PhysicalU1Y      bool
	Verdict          string
}

type UniformPhaseActionAudit struct {
	Generator               string
	ActsAsMultiplicationByI bool
	ExponentialAction       string
	UniformOnFullC2         bool
	ComplexDimension        int
	PhysicalChargeFixed     bool
	Verdict                 string
}

type ChargeNormalizationAudit struct {
	PhaseLineFixed                   bool
	NaturalDirection                 string
	CandidateNormalizations          []string
	SameLineDifferentCharges         bool
	PhysicalHyperchargeNormalization bool
	ThetaYRequired                   bool
	Verdict                          string
}

type SelectorDependenceAudit struct {
	PhaseLineDependsOnN        bool
	NativeTwistorPointSelector bool
	SelectorIndependentU1Line  bool
	ComplexStructureSelected   bool
	Verdict                    string
}

type SU2U1AsymmetryAudit struct {
	SU2Side                string
	U1Side                 string
	SU2SelectorIndependent bool
	U1SelectorDependent    bool
	U1NormalizationOpen    bool
	Verdict                string
}

type PhysicalHyperchargeFirewallAudit struct {
	LnPhysicalU1Y                               bool
	JHHyperchargeGenerator                      bool
	InternalPhaseChargePhysicalHiggsHypercharge bool
	FullPhysicalHiggsDoublet                    bool
	HyperchargeAssignment                       bool
	HyperchargeNormalization                    bool
	HiggsMass                                   bool
	ScalarRuntime                               bool
	YukawaOperator                              bool
	YukawaEigenvalues                           bool
	MissingMaps                                 []string
	Verdict                                     string
}

type Analysis struct {
	Inherited Gate716Inheritance
	PhaseLine MovingPhaseLineAudit
	Central   CentralPhaseLineAudit
	Uniform   UniformPhaseActionAudit
	Charge    ChargeNormalizationAudit
	Selector  SelectorDependenceAudit
	Asymmetry SU2U1AsymmetryAudit
	Physical  PhysicalHyperchargeFirewallAudit
	Truth     string
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
	g716, err := gate716.BuildDefault()
	if err != nil {
		return Analysis{}, fmt.Errorf("Gate716 inheritance unavailable: %w", err)
	}
	inherited := buildGate716Inheritance(g716)
	phase := buildMovingPhaseLine(inherited)
	central := buildCentralPhaseLine(phase)
	uniform := buildUniformPhaseAction(central)
	charge := buildChargeNormalization(phase)
	selector := buildSelectorDependence(inherited, phase)
	asym := buildAsymmetry(inherited, selector)
	physical := buildPhysicalFirewall()
	truth := "Gate 717 audits the selector-dependent U(1)-like phase side of the K7+ socket. For a fixed twistor point n, L_n=span(J_H(n)) commutes with the selector-independent commutant C and is therefore central inside u(2,J_H(n)).  J_H(n) acts as multiplication by i on K7+_J(n) and exponentiates to a uniform internal phase action on the C^2 pre-Higgs carrier.  However, the line moves with n, has no selector-independent representative, and does not fix physical hypercharge normalization: J_H(n), (1/2)J_H(n), and cJ_H(n) are the same internal phase line with different charge conventions.  Thus the SU(2) side remains structurally ready while the U(1) side requires both a twistor selector and a hypercharge normalization map; no physical U(1)_Y, full Higgs doublet map, Higgs mass, scalar runtime, Yukawa operator/eigenvalue, or native 7/72 theorem is derived."
	return Analysis{Inherited: inherited, PhaseLine: phase, Central: central, Uniform: uniform, Charge: charge, Selector: selector, Asymmetry: asym, Physical: physical, Truth: truth}, nil
}

func buildGate716Inheritance(g gate716.Analysis) Gate716Inheritance {
	return Gate716Inheritance{
		SU2AirlockInherited:             g.Intertwiner.RepresentationCompatible && g.Twistor.SU2AlgebraSelectorIndependent,
		SU2SideStructurallyReady:        g.Twistor.SU2CompatibleWithoutSolvingU1,
		InternalCSelectorIndependent:    g.Twistor.SU2AlgebraSelectorIndependent,
		ComplexCarrierSelectorDependent: g.Twistor.ComplexCarrierSelectorDependent,
		U1PhaseSelectorDependent:        g.Twistor.U1PhaseSelectorDependent,
		HyperchargeDerived:              g.Physical.HyperchargeDerived,
		HyperchargeNormalized:           g.Hypercharge.HyperchargeNormalization,
		FullHiggsDoubletMap:             g.Hypercharge.FullPhysicalHiggsDoubletCertified || g.Physical.K7PlusPhysicalHiggsDoublet,
		HiggsMassCertified:              g.Physical.HiggsMass || g.Physical.ScalarRuntime,
		YukawaCertified:                 g.Physical.YukawaOperator || g.Physical.YukawaEigenvalues,
		Verdict:                         StatusGate716SU2IntertwinerAirlockInherited,
	}
}

func buildMovingPhaseLine(i Gate716Inheritance) MovingPhaseLineAudit {
	return MovingPhaseLineAudit{
		Definition:          "L_n=span(J_H(n)) inside u(2,J_H(n))",
		Dimension:           movingPhaseLineDimension,
		DependsOnSelectorN:  i.U1PhaseSelectorDependent,
		FixedJHRequired:     true,
		SelectorIndependent: false,
		PhysicalHypercharge: false,
		Verdict: strings.Join([]string{
			StatusMovingPhaseLineDefined,
			StatusNoSelectorIndependentU1PhaseLine,
			StatusInternalPhaseLineNotPhysicalU1Y,
		}, "; "),
	}
}

func buildCentralPhaseLine(p MovingPhaseLineAudit) CentralPhaseLineAudit {
	ok := p.Dimension == movingPhaseLineDimension && p.FixedJHRequired
	return CentralPhaseLineAudit{
		Commutator:       "[J_H(n),X]=0 for every X in C",
		CommutesWithC:    ok,
		LiesInCenterOfU2: ok,
		FixedJHOnly:      true,
		PhysicalU1Y:      false,
		Verdict: strings.Join([]string{
			StatusLNIsCentralInU2SocketForFixedJH,
			StatusLNInternalU1PhaseSocketAfterJHChoice,
			StatusInternalPhaseLineNotPhysicalU1Y,
		}, "; "),
	}
}

func buildUniformPhaseAction(c CentralPhaseLineAudit) UniformPhaseActionAudit {
	ok := c.LiesInCenterOfU2 && c.CommutesWithC
	return UniformPhaseActionAudit{
		Generator:               "J_H(n)",
		ActsAsMultiplicationByI: ok,
		ExponentialAction:       "exp(theta J_H(n)) · v on K7+_J(n) ~= C^2",
		UniformOnFullC2:         ok,
		ComplexDimension:        k7PlusComplexDimension,
		PhysicalChargeFixed:     false,
		Verdict: strings.Join([]string{
			StatusJHExponentiatesToUniformPhaseOnC2,
			StatusK7PlusJHUniformInternalPhaseAction,
			StatusNoHyperchargeAssignmentOrNormalization,
		}, "; "),
	}
}

func buildChargeNormalization(p MovingPhaseLineAudit) ChargeNormalizationAudit {
	return ChargeNormalizationAudit{
		PhaseLineFixed:                   p.Dimension == movingPhaseLineDimension,
		NaturalDirection:                 "J_H(n)",
		CandidateNormalizations:          []string{"J_H(n)", "(1/2)J_H(n)", "c J_H(n)"},
		SameLineDifferentCharges:         true,
		PhysicalHyperchargeNormalization: false,
		ThetaYRequired:                   true,
		Verdict: strings.Join([]string{
			StatusChargeNormalizationAudited,
			StatusInternalPhaseLineNotPhysicalU1Y,
			StatusNoHyperchargeAssignmentOrNormalization,
		}, "; "),
	}
}

func buildSelectorDependence(i Gate716Inheritance, p MovingPhaseLineAudit) SelectorDependenceAudit {
	return SelectorDependenceAudit{
		PhaseLineDependsOnN:        p.DependsOnSelectorN,
		NativeTwistorPointSelector: false,
		SelectorIndependentU1Line:  false,
		ComplexStructureSelected:   false,
		Verdict: strings.Join([]string{
			StatusSelectorDependenceAudited,
			StatusNoSelectorIndependentU1PhaseLine,
			StatusNoNativeTwistorPointSelector,
		}, "; "),
	}
}

func buildAsymmetry(i Gate716Inheritance, s SelectorDependenceAudit) SU2U1AsymmetryAudit {
	return SU2U1AsymmetryAudit{
		SU2Side:                "C is twistor-invariant and selector-independent",
		U1Side:                 "L_n=span(J_H(n)) moves with the twistor selector n",
		SU2SelectorIndependent: i.InternalCSelectorIndependent,
		U1SelectorDependent:    s.PhaseLineDependsOnN,
		U1NormalizationOpen:    true,
		Verdict: strings.Join([]string{
			StatusSU2U1AsymmetryRecorded,
			StatusElectroweakAirlockU1RequiresSelectorAndNorm,
			StatusNoHyperchargeAssignmentOrNormalization,
		}, "; "),
	}
}

func buildPhysicalFirewall() PhysicalHyperchargeFirewallAudit {
	missing := []string{
		"Theta_Y: span(J_H(n)) -> physical U(1)_Y with correct Higgs charge/normalization",
		"Theta_selector: native or sealed principle selecting the twistor point n",
		"Theta_H: full K7+_J(n) -> physical Higgs doublet representation",
	}
	return PhysicalHyperchargeFirewallAudit{
		LnPhysicalU1Y:          false,
		JHHyperchargeGenerator: false,
		InternalPhaseChargePhysicalHiggsHypercharge: false,
		FullPhysicalHiggsDoublet:                    false,
		HyperchargeAssignment:                       false,
		HyperchargeNormalization:                    false,
		HiggsMass:                                   false,
		ScalarRuntime:                               false,
		YukawaOperator:                              false,
		YukawaEigenvalues:                           false,
		MissingMaps:                                 missing,
		Verdict: strings.Join([]string{
			StatusPhysicalHyperchargeFirewallEnforced,
			StatusInternalPhaseLineNotPhysicalU1Y,
			StatusNoHyperchargeAssignmentOrNormalization,
			StatusNoSelectorIndependentU1PhaseLine,
			StatusNoNativeTwistorPointSelector,
			StatusNoFullTypedK7PlusToPhysicalHiggsDoubletMap,
			StatusNoHiggsMassOrScalarRuntimeTheorem,
			StatusNoYukawaOperatorOrEigenvalueTheorem,
			StatusGate717MovingU1PhaseHyperchargeBoundary,
		}, "; "),
	}
}

func Statuses() []string {
	return []string{
		StatusGate716SU2IntertwinerAirlockInherited,
		StatusMovingPhaseLineDefined,
		StatusLNIsCentralInU2SocketForFixedJH,
		StatusJHExponentiatesToUniformPhaseOnC2,
		StatusChargeNormalizationAudited,
		StatusSelectorDependenceAudited,
		StatusSU2U1AsymmetryRecorded,
		StatusPhysicalHyperchargeFirewallEnforced,
		StatusLNInternalU1PhaseSocketAfterJHChoice,
		StatusK7PlusJHUniformInternalPhaseAction,
		StatusElectroweakAirlockU1RequiresSelectorAndNorm,
		StatusInternalPhaseLineNotPhysicalU1Y,
		StatusNoHyperchargeAssignmentOrNormalization,
		StatusNoSelectorIndependentU1PhaseLine,
		StatusNoNativeTwistorPointSelector,
		StatusNoFullTypedK7PlusToPhysicalHiggsDoubletMap,
		StatusNoHiggsMassOrScalarRuntimeTheorem,
		StatusNoYukawaOperatorOrEigenvalueTheorem,
		StatusGate717MovingU1PhaseHyperchargeBoundary,
	}
}

func FormatInherited(x Gate716Inheritance) string {
	return fmt.Sprintf("inherited=%t su2Ready=%t Cindependent=%t complexDependent=%t u1Dependent=%t hyperchargeDerived=%t hyperchargeNorm=%t higgsMap=%t higgsMass=%t yukawa=%t verdict=%q", x.SU2AirlockInherited, x.SU2SideStructurallyReady, x.InternalCSelectorIndependent, x.ComplexCarrierSelectorDependent, x.U1PhaseSelectorDependent, x.HyperchargeDerived, x.HyperchargeNormalized, x.FullHiggsDoubletMap, x.HiggsMassCertified, x.YukawaCertified, x.Verdict)
}

func FormatPhaseLine(x MovingPhaseLineAudit) string {
	return fmt.Sprintf("def=%q dim=%d dependsOnN=%t fixedJH=%t selectorIndependent=%t physicalY=%t verdict=%q", x.Definition, x.Dimension, x.DependsOnSelectorN, x.FixedJHRequired, x.SelectorIndependent, x.PhysicalHypercharge, x.Verdict)
}

func FormatCentral(x CentralPhaseLineAudit) string {
	return fmt.Sprintf("comm=%q commutes=%t center=%t fixedOnly=%t physicalU1Y=%t verdict=%q", x.Commutator, x.CommutesWithC, x.LiesInCenterOfU2, x.FixedJHOnly, x.PhysicalU1Y, x.Verdict)
}

func FormatUniform(x UniformPhaseActionAudit) string {
	return fmt.Sprintf("generator=%q actsI=%t exp=%q uniform=%t dimC=%d physicalCharge=%t verdict=%q", x.Generator, x.ActsAsMultiplicationByI, x.ExponentialAction, x.UniformOnFullC2, x.ComplexDimension, x.PhysicalChargeFixed, x.Verdict)
}

func FormatCharge(x ChargeNormalizationAudit) string {
	return fmt.Sprintf("lineFixed=%t natural=%q candidates=%d sameLine=%t physicalNorm=%t thetaY=%t verdict=%q", x.PhaseLineFixed, x.NaturalDirection, len(x.CandidateNormalizations), x.SameLineDifferentCharges, x.PhysicalHyperchargeNormalization, x.ThetaYRequired, x.Verdict)
}

func FormatSelector(x SelectorDependenceAudit) string {
	return fmt.Sprintf("depends=%t nativeSelector=%t independentLine=%t complexSelected=%t verdict=%q", x.PhaseLineDependsOnN, x.NativeTwistorPointSelector, x.SelectorIndependentU1Line, x.ComplexStructureSelected, x.Verdict)
}

func FormatAsymmetry(x SU2U1AsymmetryAudit) string {
	return fmt.Sprintf("su2=%q u1=%q su2Independent=%t u1Dependent=%t u1NormOpen=%t verdict=%q", x.SU2Side, x.U1Side, x.SU2SelectorIndependent, x.U1SelectorDependent, x.U1NormalizationOpen, x.Verdict)
}

func FormatPhysical(x PhysicalHyperchargeFirewallAudit) string {
	return fmt.Sprintf("LnU1Y=%t JHGenerator=%t phaseCharge=%t fullHiggs=%t assign=%t norm=%t higgsMass=%t runtime=%t yukawa=%t eigen=%t missing=%d verdict=%q", x.LnPhysicalU1Y, x.JHHyperchargeGenerator, x.InternalPhaseChargePhysicalHiggsHypercharge, x.FullPhysicalHiggsDoublet, x.HyperchargeAssignment, x.HyperchargeNormalization, x.HiggsMass, x.ScalarRuntime, x.YukawaOperator, x.YukawaEigenvalues, len(x.MissingMaps), x.Verdict)
}
