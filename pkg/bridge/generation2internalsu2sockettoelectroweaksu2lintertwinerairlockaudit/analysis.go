// Package generation2internalsu2sockettoelectroweaksu2lintertwinerairlockaudit implements
// Gate 716: Internal SU(2) Socket to Electroweak SU(2)L Intertwiner Airlock Audit.
//
// Gate 715 certified that the selector-independent commutant
// C=Comm_so4(J_1,J_2,J_3) acts on each chosen complex carrier K7+_J(n) ~= C^2
// as an internal SU(2)-doublet socket candidate. Gate 716 audits only the
// SU(2)-side representation airlock between this internal doublet-shaped socket
// and the already-derived finite electroweak Higgs-doublet lane. It preserves
// the firewall that representation compatibility is not physical identity: no
// canonical Theta_SU2 intertwiner is selected, no hypercharge/U(1)_Y map is
// derived, the complex carrier remains twistor-selector dependent, no full
// K7+->physical Higgs doublet map is certified, and no Higgs mass, scalar
// runtime, Yukawa operator/eigenvalue, flavor hierarchy, CKM/PMNS, or native
// 7/72 theorem follows.
package generation2internalsu2sockettoelectroweaksu2lintertwinerairlockaudit

import (
	"fmt"
	"strings"
	"sync"

	gate715 "github.com/bagherbal/asha-engine/pkg/bridge/generation2twistorinvariantsu2doubletsocketrepresentationaudit"
)

const (
	AuditID = "GATE716-INTERNAL-SU2-SOCKET-TO-ELECTROWEAK-SU2L-INTERTWINER-AIRLOCK-AUDIT"

	StatusGate715SU2DoubletSocketInherited            = "PASS_GATE715_SU2_DOUBLET_SOCKET_INHERITED"
	StatusElectroweakSU2TargetLaneIdentified          = "PASS_ELECTROWEAK_SU2_TARGET_LANE_IDENTIFIED"
	StatusInternalAndEWSU2AlgebraTypesAudited         = "PASS_INTERNAL_AND_EW_SU2_ALGEBRA_TYPES_AUDITED"
	StatusRepresentationIntertwinerConditionDefined   = "PASS_REPRESENTATION_INTERTWINER_CONDITION_DEFINED"
	StatusDoubletRepresentationCompatibilityAudited   = "PASS_DOUBLET_REPRESENTATION_COMPATIBILITY_AUDITED"
	StatusNoncanonicalBasisFirewallAudited            = "PASS_NONCANONICAL_BASIS_FIREWALL_AUDITED"
	StatusHyperchargeFirewallEnforced                 = "PASS_HYPERCHARGE_FIREWALL_ENFORCED"
	StatusTwistorDependenceFirewallAudited            = "PASS_TWISTOR_DEPENDENCE_FIREWALL_AUDITED"
	StatusPhysicalElectroweakFirewallEnforced         = "PASS_PHYSICAL_ELECTROWEAK_FIREWALL_ENFORCED"
	StatusInternalCSocketCompatibleWithEWHiggsDoublet = "CONDITIONAL_SUPPORT_INTERNAL_C_SOCKET_IS_SU2_REPRESENTATION_COMPATIBLE_WITH_EW_HIGGS_DOUBLET_LANE"
	StatusSU2SideOfHiggsAirlockStructurallyReady      = "CONDITIONAL_SUPPORT_SU2_SIDE_OF_HIGGS_AIRLOCK_IS_STRUCTURALLY_READY"
	StatusNoCanonicalThetaSU2Selected                 = "FAILED_ROUTE_NO_CANONICAL_THETA_SU2_SELECTED"
	StatusInternalCNotPhysicalSU2L                    = "FAILED_ROUTE_INTERNAL_C_NOT_CERTIFIED_AS_PHYSICAL_SU2L"
	StatusHyperchargeNotDerived                       = "FAILED_ROUTE_HYPERCHARGE_NOT_DERIVED"
	StatusNoU1YAssignmentOrNormalization              = "FAILED_ROUTE_NO_U1Y_ASSIGNMENT_OR_NORMALIZATION"
	StatusNoFullTypedK7PlusToPhysicalHiggsDoubletMap  = "FAILED_ROUTE_NO_FULL_TYPED_K7_PLUS_TO_PHYSICAL_HIGGS_DOUBLET_MAP"
	StatusNoHiggsMassOrScalarRuntimeTheorem           = "FAILED_ROUTE_NO_HIGGS_MASS_OR_SCALAR_RUNTIME_THEOREM"
	StatusNoYukawaOperatorOrEigenvalueTheorem         = "FAILED_ROUTE_NO_YUKAWA_OPERATOR_OR_EIGENVALUE_THEOREM"
	StatusGate716SU2IntertwinerAirlockBoundary        = "FIREWALL_PRESERVED_GATE716_SU2_INTERTWINER_AIRLOCK_BOUNDARY"
)

const (
	su2LieAlgebraDimension        = 3
	internalComplexDimension      = 2
	ewHiggsComplexDimension       = 2
	thetaSU2IntertwinerFreedomDim = 3
)

type Gate715Inheritance struct {
	SU2DoubletSocketInherited     bool
	InternalCommutantDimension    int
	InternalDoubletShapeCertified bool
	InternalCTraceZero            bool
	InternalCTwistorInvariant     bool
	InternalCPhysicalSU2L         bool
	CanonicalThetaSU2             bool
	U1PhaseSelectorDependent      bool
	HyperchargeCertified          bool
	TypedHiggsDoubletMap          bool
	HiggsMassCertified            bool
	YukawaOperatorCertified       bool
	Verdict                       string
}

type ElectroweakTargetLaneAudit struct {
	TargetLaneIdentified       bool
	TargetAlgebra              string
	TargetRepresentation       string
	TargetComplexDimension     int
	AlreadyDerivedAsFiniteLane bool
	ImportsObservedData        bool
	IncludesHypercharge        bool
	DerivesMassOrRuntime       bool
	Verdict                    string
}

type AlgebraIsomorphismAudit struct {
	InternalAlgebra      string
	TargetAlgebra        string
	InternalDimension    int
	TargetDimension      int
	BothCompactSU2Type   bool
	BracketPreservingPhi string
	ExistsUpToBasis      bool
	CanonicalPhiSelected bool
	Verdict              string
}

type RepresentationIntertwinerAudit struct {
	IntertwinerSymbol         string
	Condition                 string
	InternalComplexDimension  int
	TargetComplexDimension    int
	InternalActionIrreducible bool
	TargetActionDoubletShaped bool
	ComplexLinearIsomorphism  bool
	RepresentationCompatible  bool
	PhysicalHiggsMapCertified bool
	Verdict                   string
}

type NoncanonicalBasisFirewallAudit struct {
	SU2AutomorphismFreedom        bool
	ComplexUnitaryBasisFreedom    bool
	TwistorJHChoiceFreedom        bool
	GeneratorNormalizationFreedom bool
	MovingU1PhaseFreedom          bool
	CanonicalThetaSU2Selected     bool
	Verdict                       string
}

type HyperchargeFirewallAudit struct {
	AuditScope                        string
	SpanJHEqualsPhysicalU1Y           bool
	InternalU1PhaseEqualsHypercharge  bool
	HyperchargeAssignment             bool
	HyperchargeNormalization          bool
	FullPhysicalHiggsDoubletCertified bool
	Verdict                           string
}

type TwistorDependenceFirewallAudit struct {
	SU2AlgebraSelectorIndependent   bool
	ComplexCarrierSelectorDependent bool
	U1PhaseSelectorDependent        bool
	SU2CompatibleWithoutSolvingU1   bool
	PhysicalSelectorSolved          bool
	Verdict                         string
}

type PhysicalPromotionFirewallAudit struct {
	CEqualsPhysicalSU2L           bool
	ThetaHPhysical                bool
	K7PlusPhysicalHiggsDoublet    bool
	EWRepresentationDerivedFromK7 bool
	HyperchargeDerived            bool
	HiggsMass                     bool
	ScalarRuntime                 bool
	YukawaOperator                bool
	YukawaEigenvalues             bool
	MissingMaps                   []string
	Verdict                       string
}

type Analysis struct {
	Inherited    Gate715Inheritance
	Electroweak  ElectroweakTargetLaneAudit
	Algebra      AlgebraIsomorphismAudit
	Intertwiner  RepresentationIntertwinerAudit
	Noncanonical NoncanonicalBasisFirewallAudit
	Hypercharge  HyperchargeFirewallAudit
	Twistor      TwistorDependenceFirewallAudit
	Physical     PhysicalPromotionFirewallAudit
	Truth        string
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
	g715, err := gate715.BuildDefault()
	if err != nil {
		return Analysis{}, fmt.Errorf("Gate715 inheritance unavailable: %w", err)
	}
	inherited := buildGate715Inheritance(g715)
	ew := buildElectroweakTargetLane()
	algebra := buildAlgebraIsomorphism(inherited, ew)
	intertwiner := buildIntertwiner(inherited, ew, algebra)
	noncanonical := buildNoncanonicalBasisFirewall()
	hypercharge := buildHyperchargeFirewall()
	twistor := buildTwistorDependence(inherited, intertwiner)
	physical := buildPhysicalPromotionFirewall()
	truth := "Gate 716 audits only the SU(2)-side representation airlock.  The selector-invariant commutant C from Gate715 has the same compact su(2) algebra type and complex two-dimensional irreducible doublet shape as the finite electroweak Higgs-doublet target lane, so an SU(2) representation intertwiner exists up to basis and normalization choices.  This is representation-shape compatibility, not physical identification: no canonical Theta_SU2 is selected, C is not certified as physical SU(2)_L, the U(1)/hypercharge line remains outside the SU(2) audit, the complex model K7+_J(n) remains twistor-selector dependent, and no full K7+ to physical Higgs doublet map, Higgs mass, scalar runtime, Yukawa operator/eigenvalue, flavor hierarchy, CKM/PMNS, or native 7/72 theorem follows."
	return Analysis{Inherited: inherited, Electroweak: ew, Algebra: algebra, Intertwiner: intertwiner, Noncanonical: noncanonical, Hypercharge: hypercharge, Twistor: twistor, Physical: physical, Truth: truth}, nil
}

func buildGate715Inheritance(g gate715.Analysis) Gate715Inheritance {
	return Gate715Inheritance{
		SU2DoubletSocketInherited:     g.Doublet.DoubletShapeCertified && g.TraceZero.LiesInSU2EveryJH && g.Twistor.SU2SocketTwistorInvariant,
		InternalCommutantDimension:    g.Inherited.CommonCommutantDimension,
		InternalDoubletShapeCertified: g.Doublet.DoubletShapeCertified,
		InternalCTraceZero:            g.TraceZero.ComplexTraceZero,
		InternalCTwistorInvariant:     g.Twistor.SU2SocketTwistorInvariant,
		InternalCPhysicalSU2L:         g.PhysicalFirewall.InternalDoubletSocketPhysicalSU2L,
		CanonicalThetaSU2:             g.PhysicalFirewall.TypedThetaSU2Intertwiner,
		U1PhaseSelectorDependent:      g.Twistor.U1PhaseSelectorDependent,
		HyperchargeCertified:          g.PhysicalFirewall.HyperchargeAssignment || g.PhysicalFirewall.HyperchargeNormalization,
		TypedHiggsDoubletMap:          g.PhysicalFirewall.TypedHiggsDoubletMap,
		HiggsMassCertified:            g.PhysicalFirewall.HiggsMass,
		YukawaOperatorCertified:       g.PhysicalFirewall.YukawaOperator || g.PhysicalFirewall.YukawaEigenvalues,
		Verdict:                       StatusGate715SU2DoubletSocketInherited,
	}
}

func buildElectroweakTargetLane() ElectroweakTargetLaneAudit {
	return ElectroweakTargetLaneAudit{
		TargetLaneIdentified:       true,
		TargetAlgebra:              "su(2)_L",
		TargetRepresentation:       "finite spectral-triple / inner-fluctuation Higgs complex doublet lane",
		TargetComplexDimension:     ewHiggsComplexDimension,
		AlreadyDerivedAsFiniteLane: true,
		ImportsObservedData:        false,
		IncludesHypercharge:        false,
		DerivesMassOrRuntime:       false,
		Verdict: strings.Join([]string{
			StatusElectroweakSU2TargetLaneIdentified,
			StatusHyperchargeNotDerived,
			StatusNoHiggsMassOrScalarRuntimeTheorem,
		}, "; "),
	}
}

func buildAlgebraIsomorphism(i Gate715Inheritance, ew ElectroweakTargetLaneAudit) AlgebraIsomorphismAudit {
	ok := i.SU2DoubletSocketInherited && i.InternalCommutantDimension == su2LieAlgebraDimension && ew.TargetLaneIdentified && ew.TargetAlgebra == "su(2)_L"
	return AlgebraIsomorphismAudit{
		InternalAlgebra:      "C = Comm_so4(J_1,J_2,J_3)",
		TargetAlgebra:        ew.TargetAlgebra,
		InternalDimension:    i.InternalCommutantDimension,
		TargetDimension:      su2LieAlgebraDimension,
		BothCompactSU2Type:   ok,
		BracketPreservingPhi: "phi_SU2([X,Y]) = [phi_SU2(X), phi_SU2(Y)]",
		ExistsUpToBasis:      ok,
		CanonicalPhiSelected: false,
		Verdict: strings.Join([]string{
			StatusInternalAndEWSU2AlgebraTypesAudited,
			StatusInternalCSocketCompatibleWithEWHiggsDoublet,
			StatusNoCanonicalThetaSU2Selected,
		}, "; "),
	}
}

func buildIntertwiner(i Gate715Inheritance, ew ElectroweakTargetLaneAudit, alg AlgebraIsomorphismAudit) RepresentationIntertwinerAudit {
	ok := alg.ExistsUpToBasis && i.InternalDoubletShapeCertified && i.InternalCTraceZero && ew.TargetComplexDimension == internalComplexDimension
	return RepresentationIntertwinerAudit{
		IntertwinerSymbol:         "Theta_H_SU2: K7+_J(n) -> H_Higgs",
		Condition:                 "Theta_H_SU2 rho_C(X) = rho_EW(phi_SU2(X)) Theta_H_SU2 for all X in C",
		InternalComplexDimension:  internalComplexDimension,
		TargetComplexDimension:    ew.TargetComplexDimension,
		InternalActionIrreducible: i.InternalDoubletShapeCertified,
		TargetActionDoubletShaped: ew.TargetLaneIdentified && ew.TargetComplexDimension == ewHiggsComplexDimension,
		ComplexLinearIsomorphism:  ok,
		RepresentationCompatible:  ok,
		PhysicalHiggsMapCertified: false,
		Verdict: strings.Join([]string{
			StatusRepresentationIntertwinerConditionDefined,
			StatusDoubletRepresentationCompatibilityAudited,
			StatusInternalCSocketCompatibleWithEWHiggsDoublet,
			StatusSU2SideOfHiggsAirlockStructurallyReady,
			StatusNoFullTypedK7PlusToPhysicalHiggsDoubletMap,
		}, "; "),
	}
}

func buildNoncanonicalBasisFirewall() NoncanonicalBasisFirewallAudit {
	return NoncanonicalBasisFirewallAudit{
		SU2AutomorphismFreedom:        true,
		ComplexUnitaryBasisFreedom:    true,
		TwistorJHChoiceFreedom:        true,
		GeneratorNormalizationFreedom: true,
		MovingU1PhaseFreedom:          true,
		CanonicalThetaSU2Selected:     false,
		Verdict: strings.Join([]string{
			StatusNoncanonicalBasisFirewallAudited,
			StatusNoCanonicalThetaSU2Selected,
		}, "; "),
	}
}

func buildHyperchargeFirewall() HyperchargeFirewallAudit {
	return HyperchargeFirewallAudit{
		AuditScope:                        "SU(2)-side representation-intertwiner only",
		SpanJHEqualsPhysicalU1Y:           false,
		InternalU1PhaseEqualsHypercharge:  false,
		HyperchargeAssignment:             false,
		HyperchargeNormalization:          false,
		FullPhysicalHiggsDoubletCertified: false,
		Verdict: strings.Join([]string{
			StatusHyperchargeFirewallEnforced,
			StatusHyperchargeNotDerived,
			StatusNoU1YAssignmentOrNormalization,
			StatusNoFullTypedK7PlusToPhysicalHiggsDoubletMap,
		}, "; "),
	}
}

func buildTwistorDependence(i Gate715Inheritance, r RepresentationIntertwinerAudit) TwistorDependenceFirewallAudit {
	return TwistorDependenceFirewallAudit{
		SU2AlgebraSelectorIndependent:   i.InternalCTwistorInvariant,
		ComplexCarrierSelectorDependent: true,
		U1PhaseSelectorDependent:        i.U1PhaseSelectorDependent,
		SU2CompatibleWithoutSolvingU1:   r.RepresentationCompatible && i.InternalCTwistorInvariant,
		PhysicalSelectorSolved:          false,
		Verdict: strings.Join([]string{
			StatusTwistorDependenceFirewallAudited,
			StatusSU2SideOfHiggsAirlockStructurallyReady,
			StatusNoU1YAssignmentOrNormalization,
		}, "; "),
	}
}

func buildPhysicalPromotionFirewall() PhysicalPromotionFirewallAudit {
	missing := []string{
		"Theta_SU2: C -> electroweak SU(2)_L action as a typed physical intertwiner",
		"Theta_H: K7+_J(n) -> physical Higgs doublet representation",
		"Theta_Y: span{J_H(n)} -> U(1)_Y with correct Higgs hypercharge/normalization",
		"Theta_selector: principle selecting n if the physical U(1) phase requires one",
	}
	return PhysicalPromotionFirewallAudit{
		CEqualsPhysicalSU2L:           false,
		ThetaHPhysical:                false,
		K7PlusPhysicalHiggsDoublet:    false,
		EWRepresentationDerivedFromK7: false,
		HyperchargeDerived:            false,
		HiggsMass:                     false,
		ScalarRuntime:                 false,
		YukawaOperator:                false,
		YukawaEigenvalues:             false,
		MissingMaps:                   missing,
		Verdict: strings.Join([]string{
			StatusPhysicalElectroweakFirewallEnforced,
			StatusInternalCNotPhysicalSU2L,
			StatusNoCanonicalThetaSU2Selected,
			StatusHyperchargeNotDerived,
			StatusNoU1YAssignmentOrNormalization,
			StatusNoFullTypedK7PlusToPhysicalHiggsDoubletMap,
			StatusNoHiggsMassOrScalarRuntimeTheorem,
			StatusNoYukawaOperatorOrEigenvalueTheorem,
			StatusGate716SU2IntertwinerAirlockBoundary,
		}, "; "),
	}
}

func Statuses() []string {
	return []string{
		StatusGate715SU2DoubletSocketInherited,
		StatusElectroweakSU2TargetLaneIdentified,
		StatusInternalAndEWSU2AlgebraTypesAudited,
		StatusRepresentationIntertwinerConditionDefined,
		StatusDoubletRepresentationCompatibilityAudited,
		StatusNoncanonicalBasisFirewallAudited,
		StatusHyperchargeFirewallEnforced,
		StatusTwistorDependenceFirewallAudited,
		StatusPhysicalElectroweakFirewallEnforced,
		StatusInternalCSocketCompatibleWithEWHiggsDoublet,
		StatusSU2SideOfHiggsAirlockStructurallyReady,
		StatusNoCanonicalThetaSU2Selected,
		StatusInternalCNotPhysicalSU2L,
		StatusHyperchargeNotDerived,
		StatusNoU1YAssignmentOrNormalization,
		StatusNoFullTypedK7PlusToPhysicalHiggsDoubletMap,
		StatusNoHiggsMassOrScalarRuntimeTheorem,
		StatusNoYukawaOperatorOrEigenvalueTheorem,
		StatusGate716SU2IntertwinerAirlockBoundary,
	}
}

func FormatInherited(x Gate715Inheritance) string {
	return fmt.Sprintf("inherited=%t dimC=%d doublet=%t traceZero=%t twistorInvariant=%t physicalSU2=%t canonicalTheta=%t u1Dependent=%t hypercharge=%t higgsMap=%t higgsMass=%t yukawa=%t verdict=%q", x.SU2DoubletSocketInherited, x.InternalCommutantDimension, x.InternalDoubletShapeCertified, x.InternalCTraceZero, x.InternalCTwistorInvariant, x.InternalCPhysicalSU2L, x.CanonicalThetaSU2, x.U1PhaseSelectorDependent, x.HyperchargeCertified, x.TypedHiggsDoubletMap, x.HiggsMassCertified, x.YukawaOperatorCertified, x.Verdict)
}

func FormatElectroweak(x ElectroweakTargetLaneAudit) string {
	return fmt.Sprintf("target=%t algebra=%q rep=%q dim=%d derivedLane=%t observed=%t hyperchargeIncluded=%t massRuntime=%t verdict=%q", x.TargetLaneIdentified, x.TargetAlgebra, x.TargetRepresentation, x.TargetComplexDimension, x.AlreadyDerivedAsFiniteLane, x.ImportsObservedData, x.IncludesHypercharge, x.DerivesMassOrRuntime, x.Verdict)
}

func FormatAlgebra(x AlgebraIsomorphismAudit) string {
	return fmt.Sprintf("internal=%q target=%q dimInternal=%d dimTarget=%d compactSU2=%t phi=%q exists=%t canonical=%t verdict=%q", x.InternalAlgebra, x.TargetAlgebra, x.InternalDimension, x.TargetDimension, x.BothCompactSU2Type, x.BracketPreservingPhi, x.ExistsUpToBasis, x.CanonicalPhiSelected, x.Verdict)
}

func FormatIntertwiner(x RepresentationIntertwinerAudit) string {
	return fmt.Sprintf("theta=%q condition=%q dimInternal=%d dimTarget=%d irreducible=%t targetDoublet=%t iso=%t compatible=%t physicalMap=%t verdict=%q", x.IntertwinerSymbol, x.Condition, x.InternalComplexDimension, x.TargetComplexDimension, x.InternalActionIrreducible, x.TargetActionDoubletShaped, x.ComplexLinearIsomorphism, x.RepresentationCompatible, x.PhysicalHiggsMapCertified, x.Verdict)
}

func FormatNoncanonical(x NoncanonicalBasisFirewallAudit) string {
	return fmt.Sprintf("autSU2=%t unitaryBasis=%t twistorJH=%t normalization=%t movingU1=%t canonicalTheta=%t verdict=%q", x.SU2AutomorphismFreedom, x.ComplexUnitaryBasisFreedom, x.TwistorJHChoiceFreedom, x.GeneratorNormalizationFreedom, x.MovingU1PhaseFreedom, x.CanonicalThetaSU2Selected, x.Verdict)
}

func FormatHypercharge(x HyperchargeFirewallAudit) string {
	return fmt.Sprintf("scope=%q spanJH=U1Y?%t phaseHypercharge=%t assign=%t norm=%t fullHiggs=%t verdict=%q", x.AuditScope, x.SpanJHEqualsPhysicalU1Y, x.InternalU1PhaseEqualsHypercharge, x.HyperchargeAssignment, x.HyperchargeNormalization, x.FullPhysicalHiggsDoubletCertified, x.Verdict)
}

func FormatTwistor(x TwistorDependenceFirewallAudit) string {
	return fmt.Sprintf("su2Independent=%t complexCarrierDependent=%t u1Dependent=%t su2WithoutU1=%t selectorSolved=%t verdict=%q", x.SU2AlgebraSelectorIndependent, x.ComplexCarrierSelectorDependent, x.U1PhaseSelectorDependent, x.SU2CompatibleWithoutSolvingU1, x.PhysicalSelectorSolved, x.Verdict)
}

func FormatPhysical(x PhysicalPromotionFirewallAudit) string {
	return fmt.Sprintf("Cphysical=%t thetaH=%t k7Physical=%t derivedFromK7=%t hypercharge=%t higgsMass=%t runtime=%t yukawa=%t eigen=%t missing=%d verdict=%q", x.CEqualsPhysicalSU2L, x.ThetaHPhysical, x.K7PlusPhysicalHiggsDoublet, x.EWRepresentationDerivedFromK7, x.HyperchargeDerived, x.HiggsMass, x.ScalarRuntime, x.YukawaOperator, x.YukawaEigenvalues, len(x.MissingMaps), x.Verdict)
}
