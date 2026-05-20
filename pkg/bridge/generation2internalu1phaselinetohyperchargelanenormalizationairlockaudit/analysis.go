// Package generation2internalu1phaselinetohyperchargelanenormalizationairlockaudit implements
// Gate 718: Internal U(1) Phase Line to Hypercharge Lane Normalization Airlock Audit.
//
// Gate 717 certified that, after a twistor point n is chosen, the moving phase
// line L_n=span(J_H(n)) is central inside u(2,J_H(n)) and exponentiates to a
// uniform internal phase action on K7+_J(n) ~= C^2. Gate 718 audits only the
// representation-normalization airlock from that internal one-dimensional phase
// socket to the already-derived finite spectral-triple U(1)_Y Higgs lane. It
// conditionally supports one-dimensional abelian representation compatibility
// after choosing both a twistor selector n and a normalization constant q, while
// preserving the firewalls that the phase line does not itself fix physical
// hypercharge, does not select n, does not complete the physical Higgs map, and
// does not derive Higgs mass, scalar runtime, Yukawa operators/eigenvalues,
// flavor hierarchy, CKM/PMNS, or native 7/72.
package generation2internalu1phaselinetohyperchargelanenormalizationairlockaudit

import (
	"fmt"
	"strings"
	"sync"

	gate717 "github.com/bagherbal/asha-engine/pkg/bridge/generation2movingu1phaselineandhyperchargenormalizationfirewallaudit"
)

const (
	AuditID = "GATE718-INTERNAL-U1-PHASE-LINE-TO-HYPERCHARGE-LANE-NORMALIZATION-AIRLOCK-AUDIT"

	StatusGate717MovingU1PhaseInherited           = "PASS_GATE717_MOVING_U1_PHASE_INHERITED"
	StatusInternalPhaseLineShapeAudited           = "PASS_INTERNAL_PHASE_LINE_SHAPE_AUDITED"
	StatusUniformPhaseActionInherited             = "PASS_UNIFORM_PHASE_ACTION_INHERITED"
	StatusNormalizationFreedomAudited             = "PASS_NORMALIZATION_FREEDOM_AUDITED"
	StatusU1YTargetLaneIdentified                 = "PASS_U1Y_TARGET_LANE_IDENTIFIED"
	StatusU1RepresentationCompatibilityAudited    = "PASS_U1_REPRESENTATION_COMPATIBILITY_AUDITED"
	StatusSelectorDependenceFirewallAudited       = "PASS_SELECTOR_DEPENDENCE_FIREWALL_AUDITED"
	StatusCombinedElectroweakAirlockStatusUpdated = "PASS_COMBINED_ELECTROWEAK_AIRLOCK_STATUS_UPDATED"
	StatusLnU1YCompatibleAfterSelectorAndNorm     = "CONDITIONAL_SUPPORT_LN_IS_U1Y_COMPATIBLE_PHASE_LINE_AFTER_SELECTOR_AND_NORMALIZATION"
	StatusFullU2SocketCompatibleOnlyAfterNAndQ    = "CONDITIONAL_SUPPORT_FULL_U2_SOCKET_IS_REPRESENTATION_COMPATIBLE_ONLY_AFTER_N_AND_Q_CHOICES"
	StatusPhaseLineDoesNotFixHyperchargeNorm      = "FAILED_ROUTE_PHASE_LINE_DOES_NOT_FIX_HYPERCHARGE_NORMALIZATION"
	StatusNoNativeTwistorPointSelector            = "FAILED_ROUTE_NO_NATIVE_TWISTOR_POINT_SELECTOR"
	StatusNoNativeThetaYNormalizationTheorem      = "FAILED_ROUTE_NO_NATIVE_THETA_Y_NORMALIZATION_THEOREM"
	StatusNoFullTypedK7PlusToPhysicalHiggsMap     = "FAILED_ROUTE_NO_FULL_TYPED_K7_PLUS_TO_PHYSICAL_HIGGS_DOUBLET_MAP"
	StatusNoHiggsMassOrScalarRuntimeTheorem       = "FAILED_ROUTE_NO_HIGGS_MASS_OR_SCALAR_RUNTIME_THEOREM"
	StatusNoYukawaOperatorOrEigenvalueTheorem     = "FAILED_ROUTE_NO_YUKAWA_OPERATOR_OR_EIGENVALUE_THEOREM"
	StatusGate718U1HyperchargeAirlockBoundary     = "FIREWALL_PRESERVED_GATE718_U1_HYPERCHARGE_AIRLOCK_BOUNDARY"
)

const (
	internalPhaseLineDimension = 1
	higgsComplexDimension      = 2
)

type Gate717Inheritance struct {
	MovingPhaseInherited        bool
	CentralPhaseSocket          bool
	UniformPhaseAction          bool
	U1RequiresSelectorAndNorm   bool
	SelectorIndependentU1Line   bool
	NativeTwistorSelector       bool
	HyperchargeAssigned         bool
	HyperchargeNormalized       bool
	FullPhysicalHiggsDoubletMap bool
	HiggsMassCertified          bool
	YukawaCertified             bool
	Verdict                     string
}

type InternalPhaseLineShapeAudit struct {
	LineDefinition      string
	Generator           string
	Dimension           int
	ComplexCarrier      string
	UniformChargeSymbol string
	HasCorrectLineShape bool
	PhysicalU1Y         bool
	Verdict             string
}

type UniformPhaseActionInheritance struct {
	GeneratorActsAsI bool
	ExponentialForm  string
	UniformOnC2      bool
	ComplexDimension int
	Verdict          string
}

type NormalizationFreedomAudit struct {
	CandidateNormalizations []string
	SameLineDifferentQ      bool
	ChargeUnitFixed         bool
	PhysicalHyperchargeNorm bool
	Verdict                 string
}

type HyperchargeTargetLaneAudit struct {
	TargetLaneIdentified     bool
	TargetAction             string
	TargetComplexDimension   int
	FiniteSpectralTripleLane bool
	PhysicalIdentityClaimed  bool
	Verdict                  string
}

type U1RepresentationCompatibilityAudit struct {
	ThetaYMap                   string
	DomainDimension             int
	TargetDimension             int
	AbelianLieAlgebraTypesMatch bool
	NonzeroNormalizationNeeded  bool
	RepresentationCompatible    bool
	NormalizationNative         bool
	Verdict                     string
}

type SelectorDependenceFirewallAudit struct {
	PhaseLineDependsOnN   bool
	NativeTwistorSelector bool
	CanonicalPhysicalMap  bool
	RequiresSelector      bool
	Verdict               string
}

type CombinedElectroweakAirlockStatus struct {
	SU2Side                    string
	U1Side                     string
	SU2SelectorIndependent     bool
	U1SelectorDependent        bool
	U1NormalizationDependent   bool
	RequiredChoices            []string
	FullU2CompatibleAfterNAndQ bool
	Verdict                    string
}

type PhysicalFirewallAudit struct {
	LnPhysicalU1Y            bool
	JHHyperchargeGenerator   bool
	QDerivedHiggsHypercharge bool
	FullPhysicalHiggsDoublet bool
	HiggsMass                bool
	ScalarRuntime            bool
	YukawaOperator           bool
	YukawaEigenvalues        bool
	MissingMaps              []string
	Verdict                  string
}

type Analysis struct {
	Inherited     Gate717Inheritance
	Shape         InternalPhaseLineShapeAudit
	Uniform       UniformPhaseActionInheritance
	Normalization NormalizationFreedomAudit
	Target        HyperchargeTargetLaneAudit
	Compatibility U1RepresentationCompatibilityAudit
	Selector      SelectorDependenceFirewallAudit
	Combined      CombinedElectroweakAirlockStatus
	Physical      PhysicalFirewallAudit
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
	g717, err := gate717.BuildDefault()
	if err != nil {
		return Analysis{}, fmt.Errorf("Gate717 inheritance unavailable: %w", err)
	}
	inherited := buildGate717Inheritance(g717)
	shape := buildPhaseLineShape(inherited)
	uniform := buildUniformInheritance(g717, shape)
	normalization := buildNormalizationFreedom(shape)
	target := buildHyperchargeTargetLane()
	compatibility := buildU1Compatibility(shape, target, normalization)
	selector := buildSelectorFirewall(inherited, shape)
	combined := buildCombinedAirlock(inherited, compatibility, selector)
	physical := buildPhysicalFirewall()
	truth := "Gate 718 audits the U(1)-side representation-normalization airlock.  After choosing a twistor point n, the internal line L_n=span(J_H(n)) has the correct one-dimensional abelian phase-line shape: Y_int=q J_H(n) acts uniformly on K7+_J(n) ~= C^2.  Because both L_n and the finite electroweak U(1)_Y target are one-dimensional abelian Lie algebras, a representation-compatible map Theta_Y exists after choosing a nonzero normalization constant q.  This is only compatibility: the phase line does not fix physical hypercharge normalization, no native twistor point selector is known, and no full physical Higgs-doublet map, Higgs mass/scalar-runtime theorem, Yukawa operator/eigenvalue theorem, or native 7/72 theorem is derived."
	return Analysis{Inherited: inherited, Shape: shape, Uniform: uniform, Normalization: normalization, Target: target, Compatibility: compatibility, Selector: selector, Combined: combined, Physical: physical, Truth: truth}, nil
}

func buildGate717Inheritance(g gate717.Analysis) Gate717Inheritance {
	return Gate717Inheritance{
		MovingPhaseInherited:        g.PhaseLine.Dimension == 1 && g.PhaseLine.DependsOnSelectorN,
		CentralPhaseSocket:          g.Central.LiesInCenterOfU2 && g.Central.CommutesWithC,
		UniformPhaseAction:          g.Uniform.ActsAsMultiplicationByI && g.Uniform.UniformOnFullC2,
		U1RequiresSelectorAndNorm:   g.Asymmetry.U1SelectorDependent && g.Asymmetry.U1NormalizationOpen,
		SelectorIndependentU1Line:   g.Selector.SelectorIndependentU1Line,
		NativeTwistorSelector:       g.Selector.NativeTwistorPointSelector,
		HyperchargeAssigned:         g.Physical.HyperchargeAssignment,
		HyperchargeNormalized:       g.Physical.HyperchargeNormalization,
		FullPhysicalHiggsDoubletMap: g.Physical.FullPhysicalHiggsDoublet,
		HiggsMassCertified:          g.Physical.HiggsMass || g.Physical.ScalarRuntime,
		YukawaCertified:             g.Physical.YukawaOperator || g.Physical.YukawaEigenvalues,
		Verdict:                     StatusGate717MovingU1PhaseInherited,
	}
}

func buildPhaseLineShape(i Gate717Inheritance) InternalPhaseLineShapeAudit {
	ok := i.MovingPhaseInherited && i.CentralPhaseSocket && i.UniformPhaseAction
	return InternalPhaseLineShapeAudit{
		LineDefinition:      "L_n=span(J_H(n))",
		Generator:           "Y_int=q J_H(n)",
		Dimension:           internalPhaseLineDimension,
		ComplexCarrier:      "K7+_J(n) ~= C^2",
		UniformChargeSymbol: "q",
		HasCorrectLineShape: ok,
		PhysicalU1Y:         false,
		Verdict: strings.Join([]string{
			StatusInternalPhaseLineShapeAudited,
			StatusLnU1YCompatibleAfterSelectorAndNorm,
			StatusPhaseLineDoesNotFixHyperchargeNorm,
		}, "; "),
	}
}

func buildUniformInheritance(g gate717.Analysis, s InternalPhaseLineShapeAudit) UniformPhaseActionInheritance {
	ok := s.HasCorrectLineShape && g.Uniform.ActsAsMultiplicationByI && g.Uniform.UniformOnFullC2
	return UniformPhaseActionInheritance{
		GeneratorActsAsI: ok,
		ExponentialForm:  "exp(theta q J_H(n)) · v on K7+_J(n)",
		UniformOnC2:      ok,
		ComplexDimension: higgsComplexDimension,
		Verdict:          StatusUniformPhaseActionInherited,
	}
}

func buildNormalizationFreedom(s InternalPhaseLineShapeAudit) NormalizationFreedomAudit {
	return NormalizationFreedomAudit{
		CandidateNormalizations: []string{"J_H(n)", "(1/2)J_H(n)", "c J_H(n)", "q J_H(n)"},
		SameLineDifferentQ:      s.Dimension == internalPhaseLineDimension,
		ChargeUnitFixed:         false,
		PhysicalHyperchargeNorm: false,
		Verdict: strings.Join([]string{
			StatusNormalizationFreedomAudited,
			StatusPhaseLineDoesNotFixHyperchargeNorm,
			StatusNoNativeThetaYNormalizationTheorem,
		}, "; "),
	}
}

func buildHyperchargeTargetLane() HyperchargeTargetLaneAudit {
	return HyperchargeTargetLaneAudit{
		TargetLaneIdentified:     true,
		TargetAction:             "rho_Y : u(1)_Y -> End_C(H_Higgs)",
		TargetComplexDimension:   higgsComplexDimension,
		FiniteSpectralTripleLane: true,
		PhysicalIdentityClaimed:  false,
		Verdict:                  StatusU1YTargetLaneIdentified,
	}
}

func buildU1Compatibility(s InternalPhaseLineShapeAudit, t HyperchargeTargetLaneAudit, n NormalizationFreedomAudit) U1RepresentationCompatibilityAudit {
	ok := s.HasCorrectLineShape && t.TargetLaneIdentified && t.TargetComplexDimension == higgsComplexDimension && n.SameLineDifferentQ && !n.PhysicalHyperchargeNorm
	return U1RepresentationCompatibilityAudit{
		ThetaYMap:                   "Theta_Y : L_n -> u(1)_Y,  J_H(n) |-> q_Y Y_H",
		DomainDimension:             s.Dimension,
		TargetDimension:             internalPhaseLineDimension,
		AbelianLieAlgebraTypesMatch: true,
		NonzeroNormalizationNeeded:  true,
		RepresentationCompatible:    ok,
		NormalizationNative:         false,
		Verdict: strings.Join([]string{
			StatusU1RepresentationCompatibilityAudited,
			StatusLnU1YCompatibleAfterSelectorAndNorm,
			StatusNoNativeThetaYNormalizationTheorem,
		}, "; "),
	}
}

func buildSelectorFirewall(i Gate717Inheritance, s InternalPhaseLineShapeAudit) SelectorDependenceFirewallAudit {
	return SelectorDependenceFirewallAudit{
		PhaseLineDependsOnN:   i.MovingPhaseInherited,
		NativeTwistorSelector: false,
		CanonicalPhysicalMap:  false,
		RequiresSelector:      s.HasCorrectLineShape,
		Verdict: strings.Join([]string{
			StatusSelectorDependenceFirewallAudited,
			StatusNoNativeTwistorPointSelector,
		}, "; "),
	}
}

func buildCombinedAirlock(i Gate717Inheritance, c U1RepresentationCompatibilityAudit, s SelectorDependenceFirewallAudit) CombinedElectroweakAirlockStatus {
	req := []string{"twistor selector n", "hypercharge normalization q"}
	return CombinedElectroweakAirlockStatus{
		SU2Side:                    "C is selector-independent and doublet-compatible from Gate716",
		U1Side:                     "L_n is phase-compatible only after choosing n and q",
		SU2SelectorIndependent:     true,
		U1SelectorDependent:        s.PhaseLineDependsOnN,
		U1NormalizationDependent:   c.NonzeroNormalizationNeeded && !c.NormalizationNative,
		RequiredChoices:            req,
		FullU2CompatibleAfterNAndQ: c.RepresentationCompatible && s.RequiresSelector,
		Verdict: strings.Join([]string{
			StatusCombinedElectroweakAirlockStatusUpdated,
			StatusFullU2SocketCompatibleOnlyAfterNAndQ,
		}, "; "),
	}
}

func buildPhysicalFirewall() PhysicalFirewallAudit {
	missing := []string{
		"Theta_selector: native or sealed principle selecting the twistor point n",
		"Theta_Y: normalized map L_n -> U(1)_Y with correct Higgs charge convention",
		"Theta_H: full K7+_J(n) -> physical Higgs doublet representation",
	}
	return PhysicalFirewallAudit{
		LnPhysicalU1Y:            false,
		JHHyperchargeGenerator:   false,
		QDerivedHiggsHypercharge: false,
		FullPhysicalHiggsDoublet: false,
		HiggsMass:                false,
		ScalarRuntime:            false,
		YukawaOperator:           false,
		YukawaEigenvalues:        false,
		MissingMaps:              missing,
		Verdict: strings.Join([]string{
			StatusPhaseLineDoesNotFixHyperchargeNorm,
			StatusNoNativeTwistorPointSelector,
			StatusNoNativeThetaYNormalizationTheorem,
			StatusNoFullTypedK7PlusToPhysicalHiggsMap,
			StatusNoHiggsMassOrScalarRuntimeTheorem,
			StatusNoYukawaOperatorOrEigenvalueTheorem,
			StatusGate718U1HyperchargeAirlockBoundary,
		}, "; "),
	}
}

func Statuses() []string {
	return []string{
		StatusGate717MovingU1PhaseInherited,
		StatusInternalPhaseLineShapeAudited,
		StatusUniformPhaseActionInherited,
		StatusNormalizationFreedomAudited,
		StatusU1YTargetLaneIdentified,
		StatusU1RepresentationCompatibilityAudited,
		StatusSelectorDependenceFirewallAudited,
		StatusCombinedElectroweakAirlockStatusUpdated,
		StatusLnU1YCompatibleAfterSelectorAndNorm,
		StatusFullU2SocketCompatibleOnlyAfterNAndQ,
		StatusPhaseLineDoesNotFixHyperchargeNorm,
		StatusNoNativeTwistorPointSelector,
		StatusNoNativeThetaYNormalizationTheorem,
		StatusNoFullTypedK7PlusToPhysicalHiggsMap,
		StatusNoHiggsMassOrScalarRuntimeTheorem,
		StatusNoYukawaOperatorOrEigenvalueTheorem,
		StatusGate718U1HyperchargeAirlockBoundary,
	}
}

func FormatInherited(x Gate717Inheritance) string {
	return fmt.Sprintf("phase=%t central=%t uniform=%t requires=%t independentU1=%t selector=%t assign=%t norm=%t higgsMap=%t higgsMass=%t yukawa=%t verdict=%q", x.MovingPhaseInherited, x.CentralPhaseSocket, x.UniformPhaseAction, x.U1RequiresSelectorAndNorm, x.SelectorIndependentU1Line, x.NativeTwistorSelector, x.HyperchargeAssigned, x.HyperchargeNormalized, x.FullPhysicalHiggsDoubletMap, x.HiggsMassCertified, x.YukawaCertified, x.Verdict)
}

func FormatShape(x InternalPhaseLineShapeAudit) string {
	return fmt.Sprintf("line=%q generator=%q dim=%d carrier=%q q=%q shape=%t physical=%t verdict=%q", x.LineDefinition, x.Generator, x.Dimension, x.ComplexCarrier, x.UniformChargeSymbol, x.HasCorrectLineShape, x.PhysicalU1Y, x.Verdict)
}

func FormatUniform(x UniformPhaseActionInheritance) string {
	return fmt.Sprintf("actsI=%t exp=%q uniform=%t dim=%d verdict=%q", x.GeneratorActsAsI, x.ExponentialForm, x.UniformOnC2, x.ComplexDimension, x.Verdict)
}

func FormatNormalization(x NormalizationFreedomAudit) string {
	return fmt.Sprintf("candidates=%d sameLine=%t fixed=%t physicalNorm=%t verdict=%q", len(x.CandidateNormalizations), x.SameLineDifferentQ, x.ChargeUnitFixed, x.PhysicalHyperchargeNorm, x.Verdict)
}

func FormatTarget(x HyperchargeTargetLaneAudit) string {
	return fmt.Sprintf("identified=%t action=%q dim=%d finite=%t identityClaimed=%t verdict=%q", x.TargetLaneIdentified, x.TargetAction, x.TargetComplexDimension, x.FiniteSpectralTripleLane, x.PhysicalIdentityClaimed, x.Verdict)
}

func FormatCompatibility(x U1RepresentationCompatibilityAudit) string {
	return fmt.Sprintf("theta=%q domain=%d target=%d abelian=%t qNeeded=%t compatible=%t nativeNorm=%t verdict=%q", x.ThetaYMap, x.DomainDimension, x.TargetDimension, x.AbelianLieAlgebraTypesMatch, x.NonzeroNormalizationNeeded, x.RepresentationCompatible, x.NormalizationNative, x.Verdict)
}

func FormatSelector(x SelectorDependenceFirewallAudit) string {
	return fmt.Sprintf("depends=%t native=%t canonical=%t requires=%t verdict=%q", x.PhaseLineDependsOnN, x.NativeTwistorSelector, x.CanonicalPhysicalMap, x.RequiresSelector, x.Verdict)
}

func FormatCombined(x CombinedElectroweakAirlockStatus) string {
	return fmt.Sprintf("su2=%q u1=%q su2Independent=%t u1Dependent=%t u1Norm=%t required=%d fullAfter=%t verdict=%q", x.SU2Side, x.U1Side, x.SU2SelectorIndependent, x.U1SelectorDependent, x.U1NormalizationDependent, len(x.RequiredChoices), x.FullU2CompatibleAfterNAndQ, x.Verdict)
}

func FormatPhysical(x PhysicalFirewallAudit) string {
	return fmt.Sprintf("lnU1Y=%t jhY=%t qDerived=%t higgs=%t mass=%t runtime=%t yukawa=%t eigen=%t missing=%d verdict=%q", x.LnPhysicalU1Y, x.JHHyperchargeGenerator, x.QDerivedHiggsHypercharge, x.FullPhysicalHiggsDoublet, x.HiggsMass, x.ScalarRuntime, x.YukawaOperator, x.YukawaEigenvalues, len(x.MissingMaps), x.Verdict)
}
