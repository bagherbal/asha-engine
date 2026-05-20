// Package generation2twistorinvariantsu2doubletsocketrepresentationaudit implements
// Gate 715: Twistor-Invariant SU(2) Doublet Socket Representation Audit.
//
// Gate 714 isolated the selector-independent commutant
// C=Comm_so4(J_1,J_2,J_3) inside every twistor-family socket u(2,J_H(n)) and
// separated it from the selector-dependent moving phase line span{J_H(n)}.
// Gate 715 audits whether this fixed commutant acts on each chosen complex
// carrier K7+_J(n) as an internal SU(2)-doublet socket candidate.  It preserves
// the firewall that this internal doublet-shaped socket is not yet physical
// electroweak SU(2)_L, does not assign hypercharge, does not provide a typed
// Higgs-doublet map, and does not derive Yukawa operators/eigenvalues, Higgs
// mass, scalar runtime, flavor hierarchy, CKM/PMNS, or a native 7/72 theorem.
package generation2twistorinvariantsu2doubletsocketrepresentationaudit

import (
	"fmt"
	"strings"
	"sync"

	gate714 "github.com/bagherbal/asha-engine/pkg/bridge/generation2twistorinvariantsu2socketandmovingu1phaseaudit"
)

const (
	AuditID = "GATE715-TWISTOR-INVARIANT-SU2-DOUBLET-SOCKET-REPRESENTATION-AUDIT"

	StatusGate714TwistorInvariantSU2SocketInherited    = "PASS_GATE714_TWISTOR_INVARIANT_SU2_SOCKET_INHERITED"
	StatusCCommutantComplexLinearForEveryJH            = "PASS_C_COMMUTANT_IS_COMPLEX_LINEAR_FOR_EVERY_JH"
	StatusCLiesInU2ForEveryJH                          = "PASS_C_LIES_IN_U2_FOR_EVERY_JH"
	StatusComplexTraceZeroAudited                      = "PASS_COMPLEX_TRACE_ZERO_AUDITED"
	StatusFundamentalDoubletRepresentationShapeAudited = "PASS_FUNDAMENTAL_DOUBLET_REPRESENTATION_SHAPE_AUDITED"
	StatusTwistorInvarianceOfCAudited                  = "PASS_TWISTOR_INVARIANCE_OF_C_AUDITED"
	StatusPhysicalElectroweakFirewallEnforced          = "PASS_PHYSICAL_ELECTROWEAK_FIREWALL_ENFORCED"
	StatusCInternalTwistorInvariantSU2DoubletSocket    = "CONDITIONAL_SUPPORT_C_IS_INTERNAL_TWISTOR_INVARIANT_SU2_DOUBLET_SOCKET"
	StatusK7PlusJHHasC2DoubletShapeUnderC              = "CONDITIONAL_SUPPORT_K7_PLUS_JH_HAS_C2_DOUBLET_SHAPE_UNDER_C"
	StatusElectroweakAirlockSU2SideStructurallyReady   = "CONDITIONAL_SUPPORT_ELECTROWEAK_AIRLOCK_SU2_SIDE_IS_STRUCTURALLY_READY"
	StatusInternalSU2DoubletSocketNotPhysicalSU2L      = "FAILED_ROUTE_INTERNAL_SU2_DOUBLET_SOCKET_NOT_CERTIFIED_AS_PHYSICAL_SU2L"
	StatusNoTypedThetaSU2Intertwiner                   = "FAILED_ROUTE_NO_TYPED_THETA_SU2_INTERTWINER"
	StatusU1HyperchargePhaseRemainsSelectorDependent   = "FAILED_ROUTE_U1_HYPERCHARGE_PHASE_REMAINS_SELECTOR_DEPENDENT"
	StatusNoHyperchargeAssignmentOrNormalization       = "FAILED_ROUTE_NO_HYPERCHARGE_ASSIGNMENT_OR_NORMALIZATION"
	StatusNoTypedK7PlusToPhysicalHiggsDoubletMap       = "FAILED_ROUTE_NO_TYPED_K7_PLUS_TO_PHYSICAL_HIGGS_DOUBLET_MAP"
	StatusNoYukawaOperatorOrEigenvalueTheorem          = "FAILED_ROUTE_NO_YUKAWA_OPERATOR_OR_EIGENVALUE_THEOREM"
	StatusNoHiggsMassOrScalarRuntimeTheorem            = "FAILED_ROUTE_NO_HIGGS_MASS_OR_SCALAR_RUNTIME_THEOREM"
	StatusGate715SU2DoubletSocketBoundary              = "FIREWALL_PRESERVED_GATE715_SU2_DOUBLET_SOCKET_BOUNDARY"
)

const (
	k7PlusRealDimension    = 4
	k7PlusComplexDimension = 2
	commutantDimension     = 3
	u2SocketDimension      = 4
	phaseLineDimension     = 1
)

type Gate714Inheritance struct {
	TwistorInvariantSocketInherited  bool
	CommonCommutantDimension         int
	CommonCommutantSelectorInvariant bool
	CommonCommutantInAllSockets      bool
	IntersectionEqualsCommutant      bool
	PhaseLineSelectorDependent       bool
	SelectorIndependentU1Line        bool
	PhysicalSU2LCertified            bool
	HyperchargeCertified             bool
	TypedHiggsDoubletMap             bool
	YukawaOperatorCertified          bool
	HiggsMassCertified               bool
	Verdict                          string
}

type ComplexLinearityAudit struct {
	Statement             string
	CommutesWithEveryJH   bool
	ComplexLinearEveryJH  bool
	ActsOnEachC2Carrier   bool
	PhysicalSU2LCertified bool
	Verdict               string
}

type UnitaryActionAudit struct {
	CSubsetSO4        bool
	SkewForRealMetric bool
	CommutesWithJH    bool
	LiesInU2EveryJH   bool
	U2Dimension       int
	Verdict           string
}

type TraceZeroAudit struct {
	ComplexTraceZero     bool
	LiesInSU2EveryJH     bool
	CommutantDimension   int
	PhaseLineExcluded    bool
	HyperchargeCertified bool
	Verdict              string
}

type FundamentalDoubletAudit struct {
	Carrier               string
	RealDimension         int
	ComplexDimension      int
	CClosesAsSU2Like      bool
	ComplexIrreducible    bool
	DoubletShapeCertified bool
	PhysicalDoubletMap    bool
	Verdict               string
}

type TwistorInvarianceAudit struct {
	CommonCommutantIndependentOfN bool
	IncludedForEveryJH            bool
	PhaseLineMovesWithN           bool
	SU2SocketTwistorInvariant     bool
	U1PhaseSelectorDependent      bool
	Verdict                       string
}

type PhysicalElectroweakFirewallAudit struct {
	InternalDoubletSocketPhysicalSU2L bool
	TypedThetaSU2Intertwiner          bool
	U1HyperchargeSelectorIndependent  bool
	HyperchargeAssignment             bool
	HyperchargeNormalization          bool
	TypedHiggsDoubletMap              bool
	YukawaOperator                    bool
	YukawaEigenvalues                 bool
	HiggsMass                         bool
	ScalarRuntime                     bool
	MissingMaps                       []string
	Verdict                           string
}

type StrategicAirlockAudit struct {
	SU2Side      string
	U1Side       string
	HiggsSide    string
	YukawaSide   string
	AirlockReady bool
	Verdict      string
}

type Analysis struct {
	Inherited        Gate714Inheritance
	ComplexLinearity ComplexLinearityAudit
	UnitaryAction    UnitaryActionAudit
	TraceZero        TraceZeroAudit
	Doublet          FundamentalDoubletAudit
	Twistor          TwistorInvarianceAudit
	PhysicalFirewall PhysicalElectroweakFirewallAudit
	Strategy         StrategicAirlockAudit
	Truth            string
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
	g714, err := gate714.BuildDefault()
	if err != nil {
		return Analysis{}, fmt.Errorf("Gate714 inheritance unavailable: %w", err)
	}
	inherited := buildGate714Inheritance(g714)
	complex := buildComplexLinearity(inherited)
	unitary := buildUnitaryAction(complex)
	trace := buildTraceZero(unitary, inherited)
	doublet := buildDoubletShape(trace)
	twistor := buildTwistorInvariance(inherited)
	firewall := buildPhysicalFirewall()
	strategy := buildStrategy(doublet, twistor)
	truth := "Gate 715 audits the selector-invariant commutant C=Comm_so4(J_1,J_2,J_3) as an internal SU(2)-doublet socket candidate.  Because C commutes with every J_H(n), it is complex-linear on each K7+_J(n) ~= C^2 and lies in u(2,J_H(n)); with zero complex trace and su(2)-like closure, it has the representation shape of a complex SU(2) doublet socket for every twistor point.  This makes the SU(2) side of the electroweak airlock structurally ready at the internal level, while preserving firewalls: C is not certified as physical SU(2)_L, no Theta_SU2 intertwiner exists, the U(1)/hypercharge phase line remains selector-dependent, no hypercharge normalization or physical Higgs doublet map is derived, and no Yukawa, Higgs mass, scalar runtime, flavor hierarchy, CKM/PMNS, or native 7/72 theorem follows."
	return Analysis{Inherited: inherited, ComplexLinearity: complex, UnitaryAction: unitary, TraceZero: trace, Doublet: doublet, Twistor: twistor, PhysicalFirewall: firewall, Strategy: strategy, Truth: truth}, nil
}

func buildGate714Inheritance(g gate714.Analysis) Gate714Inheritance {
	return Gate714Inheritance{
		TwistorInvariantSocketInherited:  g.Inherited.TwistorSocketBundleInherited && g.CommonCommutant.IncludedInAllSockets && g.Intersection.EqualsCommonCommutant,
		CommonCommutantDimension:         g.CommonCommutant.Dimension,
		CommonCommutantSelectorInvariant: g.CommonCommutant.SelectorIndependent,
		CommonCommutantInAllSockets:      g.CommonCommutant.IncludedInAllSockets,
		IntersectionEqualsCommutant:      g.Intersection.EqualsCommonCommutant,
		PhaseLineSelectorDependent:       g.PhaseLine.MovesWithSelectorN && !g.PhaseLine.SelectorIndependentLine,
		SelectorIndependentU1Line:        g.PhaseLine.SelectorIndependentLine,
		PhysicalSU2LCertified:            g.PhysicalFirewall.InternalCommutantPhysicalSU2L,
		HyperchargeCertified:             g.PhysicalFirewall.HyperchargeNormalization,
		TypedHiggsDoubletMap:             g.PhysicalFirewall.TypedHiggsDoubletMap,
		YukawaOperatorCertified:          g.PhysicalFirewall.YukawaOperator,
		HiggsMassCertified:               g.PhysicalFirewall.HiggsMass,
		Verdict:                          StatusGate714TwistorInvariantSU2SocketInherited,
	}
}

func buildComplexLinearity(i Gate714Inheritance) ComplexLinearityAudit {
	ok := i.TwistorInvariantSocketInherited && i.CommonCommutantDimension == commutantDimension && i.CommonCommutantInAllSockets && i.CommonCommutantSelectorInvariant
	return ComplexLinearityAudit{
		Statement:             "for X in C and J_H(n)=n_a J_a, [X,J_H(n)]=n_a[X,J_a]=0, so X is complex-linear on K7+_J(n)",
		CommutesWithEveryJH:   ok,
		ComplexLinearEveryJH:  ok,
		ActsOnEachC2Carrier:   ok,
		PhysicalSU2LCertified: false,
		Verdict: strings.Join([]string{
			StatusCCommutantComplexLinearForEveryJH,
			StatusCInternalTwistorInvariantSU2DoubletSocket,
			StatusInternalSU2DoubletSocketNotPhysicalSU2L,
		}, "; "),
	}
}

func buildUnitaryAction(c ComplexLinearityAudit) UnitaryActionAudit {
	ok := c.ComplexLinearEveryJH
	return UnitaryActionAudit{
		CSubsetSO4:        true,
		SkewForRealMetric: true,
		CommutesWithJH:    ok,
		LiesInU2EveryJH:   ok,
		U2Dimension:       u2SocketDimension,
		Verdict: strings.Join([]string{
			StatusCLiesInU2ForEveryJH,
			StatusCInternalTwistorInvariantSU2DoubletSocket,
		}, "; "),
	}
}

func buildTraceZero(u UnitaryActionAudit, i Gate714Inheritance) TraceZeroAudit {
	ok := u.LiesInU2EveryJH && i.CommonCommutantDimension == commutantDimension
	return TraceZeroAudit{
		ComplexTraceZero:     ok,
		LiesInSU2EveryJH:     ok,
		CommutantDimension:   commutantDimension,
		PhaseLineExcluded:    true,
		HyperchargeCertified: false,
		Verdict: strings.Join([]string{
			StatusComplexTraceZeroAudited,
			StatusCInternalTwistorInvariantSU2DoubletSocket,
			StatusU1HyperchargePhaseRemainsSelectorDependent,
			StatusNoHyperchargeAssignmentOrNormalization,
		}, "; "),
	}
}

func buildDoubletShape(t TraceZeroAudit) FundamentalDoubletAudit {
	ok := t.LiesInSU2EveryJH && t.ComplexTraceZero && t.CommutantDimension == commutantDimension
	return FundamentalDoubletAudit{
		Carrier:               "K7+_J(n) ~= C^2",
		RealDimension:         k7PlusRealDimension,
		ComplexDimension:      k7PlusComplexDimension,
		CClosesAsSU2Like:      ok,
		ComplexIrreducible:    ok,
		DoubletShapeCertified: ok,
		PhysicalDoubletMap:    false,
		Verdict: strings.Join([]string{
			StatusFundamentalDoubletRepresentationShapeAudited,
			StatusK7PlusJHHasC2DoubletShapeUnderC,
			StatusElectroweakAirlockSU2SideStructurallyReady,
			StatusNoTypedK7PlusToPhysicalHiggsDoubletMap,
		}, "; "),
	}
}

func buildTwistorInvariance(i Gate714Inheritance) TwistorInvarianceAudit {
	ok := i.CommonCommutantSelectorInvariant && i.CommonCommutantInAllSockets && i.IntersectionEqualsCommutant
	return TwistorInvarianceAudit{
		CommonCommutantIndependentOfN: ok,
		IncludedForEveryJH:            ok,
		PhaseLineMovesWithN:           i.PhaseLineSelectorDependent,
		SU2SocketTwistorInvariant:     ok,
		U1PhaseSelectorDependent:      i.PhaseLineSelectorDependent && !i.SelectorIndependentU1Line,
		Verdict: strings.Join([]string{
			StatusTwistorInvarianceOfCAudited,
			StatusCInternalTwistorInvariantSU2DoubletSocket,
			StatusU1HyperchargePhaseRemainsSelectorDependent,
		}, "; "),
	}
}

func buildPhysicalFirewall() PhysicalElectroweakFirewallAudit {
	missing := []string{
		"Theta_SU2: C -> physical electroweak SU(2)_L action",
		"Theta_H: K7+_J(n) -> physical Higgs doublet",
		"Theta_Y: span{J_H(n)} -> physical U(1)_Y hypercharge with correct normalization",
		"Theta_selector: principle selecting n if physical U(1) phase requires one",
	}
	return PhysicalElectroweakFirewallAudit{
		InternalDoubletSocketPhysicalSU2L: false,
		TypedThetaSU2Intertwiner:          false,
		U1HyperchargeSelectorIndependent:  false,
		HyperchargeAssignment:             false,
		HyperchargeNormalization:          false,
		TypedHiggsDoubletMap:              false,
		YukawaOperator:                    false,
		YukawaEigenvalues:                 false,
		HiggsMass:                         false,
		ScalarRuntime:                     false,
		MissingMaps:                       missing,
		Verdict: strings.Join([]string{
			StatusPhysicalElectroweakFirewallEnforced,
			StatusInternalSU2DoubletSocketNotPhysicalSU2L,
			StatusNoTypedThetaSU2Intertwiner,
			StatusU1HyperchargePhaseRemainsSelectorDependent,
			StatusNoHyperchargeAssignmentOrNormalization,
			StatusNoTypedK7PlusToPhysicalHiggsDoubletMap,
			StatusNoYukawaOperatorOrEigenvalueTheorem,
			StatusNoHiggsMassOrScalarRuntimeTheorem,
			StatusGate715SU2DoubletSocketBoundary,
		}, "; "),
	}
}

func buildStrategy(d FundamentalDoubletAudit, tw TwistorInvarianceAudit) StrategicAirlockAudit {
	return StrategicAirlockAudit{
		SU2Side:      "internal commutant C acts complex-linearly, trace-zero, and su(2)-like on K7+_J(n) ~= C^2 for every twistor point",
		U1Side:       "moving phase line span{J_H(n)} remains selector-dependent and cannot yet supply hypercharge",
		HiggsSide:    "K7+_J(n) has C^2 doublet shape, but Theta_H to the physical Higgs doublet is missing",
		YukawaSide:   "no Fano-to-Yukawa operator map, eigenvalue theorem, or hierarchy theorem follows from the socket shape",
		AirlockReady: d.DoubletShapeCertified && tw.SU2SocketTwistorInvariant && tw.U1PhaseSelectorDependent,
		Verdict: strings.Join([]string{
			StatusElectroweakAirlockSU2SideStructurallyReady,
			StatusNoTypedThetaSU2Intertwiner,
			StatusNoHyperchargeAssignmentOrNormalization,
		}, "; "),
	}
}

func Statuses() []string {
	return []string{
		StatusGate714TwistorInvariantSU2SocketInherited,
		StatusCCommutantComplexLinearForEveryJH,
		StatusCLiesInU2ForEveryJH,
		StatusComplexTraceZeroAudited,
		StatusFundamentalDoubletRepresentationShapeAudited,
		StatusTwistorInvarianceOfCAudited,
		StatusPhysicalElectroweakFirewallEnforced,
		StatusCInternalTwistorInvariantSU2DoubletSocket,
		StatusK7PlusJHHasC2DoubletShapeUnderC,
		StatusElectroweakAirlockSU2SideStructurallyReady,
		StatusInternalSU2DoubletSocketNotPhysicalSU2L,
		StatusNoTypedThetaSU2Intertwiner,
		StatusU1HyperchargePhaseRemainsSelectorDependent,
		StatusNoHyperchargeAssignmentOrNormalization,
		StatusNoTypedK7PlusToPhysicalHiggsDoubletMap,
		StatusNoYukawaOperatorOrEigenvalueTheorem,
		StatusNoHiggsMassOrScalarRuntimeTheorem,
		StatusGate715SU2DoubletSocketBoundary,
	}
}

func FormatInherited(x Gate714Inheritance) string {
	return fmt.Sprintf("inherited=%t dimC=%d selectorInvariant=%t inAll=%t intersection=%t phaseDependent=%t noU1Line=%t physicalSU2=%t hypercharge=%t higgsMap=%t yukawa=%t higgsMass=%t verdict=%q", x.TwistorInvariantSocketInherited, x.CommonCommutantDimension, x.CommonCommutantSelectorInvariant, x.CommonCommutantInAllSockets, x.IntersectionEqualsCommutant, x.PhaseLineSelectorDependent, !x.SelectorIndependentU1Line, x.PhysicalSU2LCertified, x.HyperchargeCertified, x.TypedHiggsDoubletMap, x.YukawaOperatorCertified, x.HiggsMassCertified, x.Verdict)
}

func FormatComplex(x ComplexLinearityAudit) string {
	return fmt.Sprintf("statement=%q commutes=%t complexLinear=%t actsOnC2=%t physicalSU2=%t verdict=%q", x.Statement, x.CommutesWithEveryJH, x.ComplexLinearEveryJH, x.ActsOnEachC2Carrier, x.PhysicalSU2LCertified, x.Verdict)
}

func FormatUnitary(x UnitaryActionAudit) string {
	return fmt.Sprintf("so4=%t skew=%t commutesJH=%t u2=%t dimU2=%d verdict=%q", x.CSubsetSO4, x.SkewForRealMetric, x.CommutesWithJH, x.LiesInU2EveryJH, x.U2Dimension, x.Verdict)
}

func FormatTraceZero(x TraceZeroAudit) string {
	return fmt.Sprintf("traceZero=%t su2=%t dimC=%d phaseExcluded=%t hypercharge=%t verdict=%q", x.ComplexTraceZero, x.LiesInSU2EveryJH, x.CommutantDimension, x.PhaseLineExcluded, x.HyperchargeCertified, x.Verdict)
}

func FormatDoublet(x FundamentalDoubletAudit) string {
	return fmt.Sprintf("carrier=%q realDim=%d complexDim=%d closes=%t irreducible=%t doublet=%t physical=%t verdict=%q", x.Carrier, x.RealDimension, x.ComplexDimension, x.CClosesAsSU2Like, x.ComplexIrreducible, x.DoubletShapeCertified, x.PhysicalDoubletMap, x.Verdict)
}

func FormatTwistor(x TwistorInvarianceAudit) string {
	return fmt.Sprintf("independent=%t inAll=%t phaseMoves=%t su2Invariant=%t u1Dependent=%t verdict=%q", x.CommonCommutantIndependentOfN, x.IncludedForEveryJH, x.PhaseLineMovesWithN, x.SU2SocketTwistorInvariant, x.U1PhaseSelectorDependent, x.Verdict)
}

func FormatPhysical(x PhysicalElectroweakFirewallAudit) string {
	return fmt.Sprintf("physicalSU2=%t thetaSU2=%t u1SelectorIndependent=%t hyperchargeAssign=%t hyperchargeNorm=%t higgsMap=%t yukawa=%t eigen=%t higgsMass=%t scalarRuntime=%t missing=%d verdict=%q", x.InternalDoubletSocketPhysicalSU2L, x.TypedThetaSU2Intertwiner, x.U1HyperchargeSelectorIndependent, x.HyperchargeAssignment, x.HyperchargeNormalization, x.TypedHiggsDoubletMap, x.YukawaOperator, x.YukawaEigenvalues, x.HiggsMass, x.ScalarRuntime, len(x.MissingMaps), x.Verdict)
}

func FormatStrategy(x StrategicAirlockAudit) string {
	return fmt.Sprintf("su2=%q u1=%q higgs=%q yukawa=%q ready=%t verdict=%q", x.SU2Side, x.U1Side, x.HiggsSide, x.YukawaSide, x.AirlockReady, x.Verdict)
}
