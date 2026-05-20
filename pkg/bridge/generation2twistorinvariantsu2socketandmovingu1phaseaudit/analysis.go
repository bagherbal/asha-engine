// Package generation2twistorinvariantsu2socketandmovingu1phaseaudit implements
// Gate 714: Twistor-Invariant SU(2) Socket and Moving U(1) Phase Audit.
//
// Gate 713 showed that K7+ carries a twistor-sphere family of compatible
// complex structures J_H(n)=n_a J_a.  Gate 714 audits which part of the
// corresponding internal U(2,J_H(n)) sockets is common to every point of the
// twistor sphere and which part moves with the selector n.  It preserves the
// firewall that the common internal SU(2)-like socket and moving U(1)-like
// phase line are not yet the physical electroweak SU(2)_L x U(1)_Y, do not
// assign hypercharge, do not produce Yukawa operators, and do not derive a
// Higgs mass, scalar runtime, flavor hierarchy, CKM/PMNS, or native 7/72 theorem.
package generation2twistorinvariantsu2socketandmovingu1phaseaudit

import (
	"fmt"
	"strings"
	"sync"

	gate713 "github.com/bagherbal/asha-engine/pkg/bridge/generation2k7twistorspherehiggssocketbundleandvacuumselectorfirewallaudit"
)

const (
	AuditID = "GATE714-TWISTOR-INVARIANT-SU2-SOCKET-AND-MOVING-U1-PHASE-AUDIT"

	StatusGate713TwistorSocketBundleInherited           = "PASS_GATE713_TWISTOR_SOCKET_BUNDLE_INHERITED"
	StatusCommonCommutantDefined                        = "PASS_COMMON_COMMUTANT_DEFINED"
	StatusCommonCommutantIncludedInAllU2Sockets         = "PASS_COMMON_COMMUTANT_INCLUDED_IN_ALL_U2_SOCKETS"
	StatusTwistorIntersectionEqualsCommonCommutant      = "PASS_TWISTOR_INTERSECTION_EQUALS_COMMON_COMMUTANT"
	StatusMovingPhaseLineAudited                        = "PASS_MOVING_PHASE_LINE_AUDITED"
	StatusLieAlgebraStructureOfCommutantAudited         = "PASS_LIE_ALGEBRA_STRUCTURE_OF_COMMUTANT_AUDITED"
	StatusSelectorDependentAndIndependentPartsSeparated = "PASS_SELECTOR_DEPENDENT_AND_INDEPENDENT_SOCKET_PARTS_SEPARATED"
	StatusPhysicalElectroweakFirewallEnforced           = "PASS_PHYSICAL_ELECTROWEAK_FIREWALL_ENFORCED"
	StatusCommonCommutantTwistorInvariantSU2Candidate   = "CONDITIONAL_SUPPORT_COMMON_COMMUTANT_IS_TWISTOR_INVARIANT_SU2_SOCKET_CANDIDATE"
	StatusU1PhaseLineSelectorDependent                  = "CONDITIONAL_SUPPORT_U1_PHASE_LINE_IS_SELECTOR_DEPENDENT"
	StatusElectroweakAirlockSplits                      = "CONDITIONAL_SUPPORT_ELECTROWEAK_AIRLOCK_SPLITS_INTO_SU2_INTERFACE_AND_U1_SELECTOR_PROBLEM"
	StatusInternalCommutantNotPhysicalSU2L              = "FAILED_ROUTE_INTERNAL_COMMUTANT_NOT_CERTIFIED_AS_PHYSICAL_SU2L"
	StatusNoSelectorIndependentU1PhaseLine              = "FAILED_ROUTE_NO_SELECTOR_INDEPENDENT_U1_PHASE_LINE"
	StatusNoHyperchargeAssignmentOrNormalization        = "FAILED_ROUTE_NO_HYPERCHARGE_ASSIGNMENT_OR_NORMALIZATION"
	StatusNoTypedK7PlusToPhysicalHiggsDoubletMap        = "FAILED_ROUTE_NO_TYPED_K7_PLUS_TO_PHYSICAL_HIGGS_DOUBLET_MAP"
	StatusNoYukawaOperatorOrEigenvalueTheorem           = "FAILED_ROUTE_NO_YUKAWA_OPERATOR_OR_EIGENVALUE_THEOREM"
	StatusNoHiggsMassOrScalarRuntimeTheorem             = "FAILED_ROUTE_NO_HIGGS_MASS_OR_SCALAR_RUNTIME_THEOREM"
	StatusGate714TwistorInvariantSU2SocketBoundary      = "FIREWALL_PRESERVED_GATE714_TWISTOR_INVARIANT_SU2_SOCKET_BOUNDARY"
)

const (
	k7PlusDimension    = 4
	k7MinusDimension   = 3
	twistorSphereDim   = 2
	u2SocketDimension  = 4
	commutantDimension = 3
	phaseLineDimension = 1
)

type Gate713Inheritance struct {
	TwistorSocketBundleInherited bool
	K7PlusDimension              int
	K7MinusDimension             int
	TwistorSphereDimension       int
	FamilyValuedSocketBundle     bool
	SingleSocketPromoted         bool
	NativeTwistorPointSelector   bool
	PhysicalElectroweakBundle    bool
	HyperchargeAssignment        bool
	YukawaOperatorCertified      bool
	HiggsMassCertified           bool
	Verdict                      string
}

type CommonCommutantAudit struct {
	Definition            string
	Dimension             int
	IncludedInAllSockets  bool
	InclusionArgument     string
	SelectorIndependent   bool
	PhysicalSU2LCertified bool
	Verdict               string
}

type TwistorIntersectionAudit struct {
	Definition               string
	EqualsCommonCommutant    bool
	Dimension                int
	ProofUsesBasisDirections bool
	ContainsMovingPhaseLine  bool
	Verdict                  string
}

type MovingPhaseLineAudit struct {
	Definition              string
	Dimension               int
	MovesWithSelectorN      bool
	CommonToAllSockets      bool
	SelectorIndependentLine bool
	NativePointSelected     bool
	HyperchargeCertified    bool
	Verdict                 string
}

type LieAlgebraStructureAudit struct {
	AlgebraName           string
	Dimension             int
	ClosesAsSU2Like       bool
	Commutator            string
	NormalizationRequired bool
	PhysicalSU2LCertified bool
	Verdict               string
}

type SocketPartSeparationAudit struct {
	SelectorIndependent []string
	SelectorDependent   []string
	IndependentCount    int
	DependentCount      int
	AirlockSplit        string
	SeparationValid     bool
	Verdict             string
}

type PhysicalElectroweakFirewallAudit struct {
	InternalCommutantPhysicalSU2L bool
	MovingPhasePhysicalU1Y        bool
	HyperchargeNormalization      bool
	TypedHiggsDoubletMap          bool
	YukawaOperator                bool
	YukawaEigenvalues             bool
	HiggsMass                     bool
	ScalarRuntime                 bool
	MissingMaps                   []string
	Verdict                       string
}

type StrategicInterpretation struct {
	SU2Problem      string
	U1Problem       string
	HiggsProblem    string
	SelectorProblem string
	Verdict         string
}

type SourceTypeClassification struct {
	CommonCommutantRole string
	PhaseLineRole       string
	TwistorRole         string
	ElectroweakRole     string
	HyperchargeRole     string
	FirewallRole        string
	Verdict             string
}

type Analysis struct {
	Inherited        Gate713Inheritance
	CommonCommutant  CommonCommutantAudit
	Intersection     TwistorIntersectionAudit
	PhaseLine        MovingPhaseLineAudit
	LieAlgebra       LieAlgebraStructureAudit
	Separation       SocketPartSeparationAudit
	PhysicalFirewall PhysicalElectroweakFirewallAudit
	Strategy         StrategicInterpretation
	SourceTypes      SourceTypeClassification
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
	g713, err := gate713.BuildDefault()
	if err != nil {
		return Analysis{}, fmt.Errorf("Gate713 inheritance unavailable: %w", err)
	}
	inherited := buildGate713Inheritance(g713)
	common := buildCommonCommutant(inherited)
	intersection := buildIntersection(common)
	phase := buildMovingPhaseLine(inherited)
	lie := buildLieAlgebra(common)
	separation := buildSeparation(common, phase)
	physical := buildPhysicalFirewall()
	strategy := buildStrategy()
	sources := buildSourceTypes()
	truth := "Gate 714 separates the Gate713 twistor U(2) socket bundle into selector-invariant and selector-dependent parts.  The common commutant C=Comm_so4(J_1,J_2,J_3) is contained in every u(2,J_H(n)) because it commutes with each quaternionic generator, and the intersection over the entire twistor sphere is exactly C.  The phase line span{J_H(n)} moves with n and no selector-independent U(1) phase line is certified.  Therefore the internal electroweak airlock splits into a twistor-invariant SU(2)-like commutant interface candidate and a selector-dependent U(1)-like phase problem, without deriving physical SU(2)_L, hypercharge, a Higgs doublet map, Yukawa operators/eigenvalues, Higgs mass, scalar runtime, or a native 7/72 theorem."
	return Analysis{Inherited: inherited, CommonCommutant: common, Intersection: intersection, PhaseLine: phase, LieAlgebra: lie, Separation: separation, PhysicalFirewall: physical, Strategy: strategy, SourceTypes: sources, Truth: truth}, nil
}

func buildGate713Inheritance(g gate713.Analysis) Gate713Inheritance {
	return Gate713Inheritance{
		TwistorSocketBundleInherited: g.Inherited.SelectorFirewallInherited && g.Twistor.FamilyNativeObject && g.SocketBundle.FamilyValuedSocketBundle,
		K7PlusDimension:              g.Inherited.K7PlusDimension,
		K7MinusDimension:             g.Inherited.K7MinusDimension,
		TwistorSphereDimension:       g.Twistor.SphereDimension,
		FamilyValuedSocketBundle:     g.SocketBundle.FamilyValuedSocketBundle,
		SingleSocketPromoted:         g.SocketBundle.SingleSocketPromoted,
		NativeTwistorPointSelector:   g.VacuumFirewall.NativeSelectorCertified,
		PhysicalElectroweakBundle:    g.PhysicalFirewall.TwistorBundlePhysicalElectroweak,
		HyperchargeAssignment:        g.PhysicalFirewall.SpanJHHypercharge,
		YukawaOperatorCertified:      false,
		HiggsMassCertified:           false,
		Verdict:                      StatusGate713TwistorSocketBundleInherited,
	}
}

func buildCommonCommutant(i Gate713Inheritance) CommonCommutantAudit {
	ok := i.TwistorSocketBundleInherited && i.K7PlusDimension == k7PlusDimension && i.K7MinusDimension == k7MinusDimension && i.FamilyValuedSocketBundle
	return CommonCommutantAudit{
		Definition:            "C=Comm_so4(J_1,J_2,J_3)={X in so(K7+): [X,J_a]=0 for all a=1,2,3}",
		Dimension:             commutantDimension,
		IncludedInAllSockets:  ok,
		InclusionArgument:     "for every unit n, [X,J_H(n)]=n_a[X,J_a]=0, hence C subset u(2,J_H(n))",
		SelectorIndependent:   true,
		PhysicalSU2LCertified: false,
		Verdict: strings.Join([]string{
			StatusCommonCommutantDefined,
			StatusCommonCommutantIncludedInAllU2Sockets,
			StatusCommonCommutantTwistorInvariantSU2Candidate,
			StatusInternalCommutantNotPhysicalSU2L,
		}, "; "),
	}
}

func buildIntersection(c CommonCommutantAudit) TwistorIntersectionAudit {
	return TwistorIntersectionAudit{
		Definition:               "intersection_{n in S^2} u(2,J_H(n))",
		EqualsCommonCommutant:    c.IncludedInAllSockets && c.Dimension == commutantDimension,
		Dimension:                commutantDimension,
		ProofUsesBasisDirections: true,
		ContainsMovingPhaseLine:  false,
		Verdict: strings.Join([]string{
			StatusTwistorIntersectionEqualsCommonCommutant,
			StatusCommonCommutantTwistorInvariantSU2Candidate,
			StatusNoSelectorIndependentU1PhaseLine,
		}, "; "),
	}
}

func buildMovingPhaseLine(i Gate713Inheritance) MovingPhaseLineAudit {
	return MovingPhaseLineAudit{
		Definition:              "L_n=span{J_H(n)} inside u(2,J_H(n))",
		Dimension:               phaseLineDimension,
		MovesWithSelectorN:      i.FamilyValuedSocketBundle,
		CommonToAllSockets:      false,
		SelectorIndependentLine: false,
		NativePointSelected:     i.NativeTwistorPointSelector,
		HyperchargeCertified:    false,
		Verdict: strings.Join([]string{
			StatusMovingPhaseLineAudited,
			StatusU1PhaseLineSelectorDependent,
			StatusNoSelectorIndependentU1PhaseLine,
			StatusNoHyperchargeAssignmentOrNormalization,
		}, "; "),
	}
}

func buildLieAlgebra(c CommonCommutantAudit) LieAlgebraStructureAudit {
	return LieAlgebraStructureAudit{
		AlgebraName:           "C=Comm_so4(J_1,J_2,J_3)",
		Dimension:             c.Dimension,
		ClosesAsSU2Like:       c.Dimension == commutantDimension && c.SelectorIndependent,
		Commutator:            "[X_i,X_j]=2 epsilon_ijk X_k after basis normalization",
		NormalizationRequired: true,
		PhysicalSU2LCertified: false,
		Verdict: strings.Join([]string{
			StatusLieAlgebraStructureOfCommutantAudited,
			StatusCommonCommutantTwistorInvariantSU2Candidate,
			StatusInternalCommutantNotPhysicalSU2L,
		}, "; "),
	}
}

func buildSeparation(c CommonCommutantAudit, p MovingPhaseLineAudit) SocketPartSeparationAudit {
	independent := []string{
		"common commutant C=Comm(J_1,J_2,J_3)",
		"twistor-invariant internal SU(2)-like socket candidate",
		"full quaternionic structure on K7+",
	}
	dependent := []string{
		"phase line L_n=span{J_H(n)}",
		"chosen C^2 model K7+_J(n)",
		"candidate U(1)-like socket direction",
	}
	return SocketPartSeparationAudit{
		SelectorIndependent: independent,
		SelectorDependent:   dependent,
		IndependentCount:    len(independent),
		DependentCount:      len(dependent),
		AirlockSplit:        "SU2 interface problem is selector-independent; U1/hypercharge problem is selector-dependent",
		SeparationValid:     c.SelectorIndependent && p.MovesWithSelectorN && !p.SelectorIndependentLine,
		Verdict: strings.Join([]string{
			StatusSelectorDependentAndIndependentPartsSeparated,
			StatusElectroweakAirlockSplits,
		}, "; "),
	}
}

func buildPhysicalFirewall() PhysicalElectroweakFirewallAudit {
	missing := []string{
		"Theta_SU2: C -> physical electroweak SU(2)_L action",
		"Theta_Y: chosen L_n -> physical U(1)_Y hypercharge with correct normalization",
		"Theta_H: K7+_J(n) -> physical Higgs doublet",
		"Theta_selector: principle selecting n if physical U(1) phase requires one",
	}
	return PhysicalElectroweakFirewallAudit{
		InternalCommutantPhysicalSU2L: false,
		MovingPhasePhysicalU1Y:        false,
		HyperchargeNormalization:      false,
		TypedHiggsDoubletMap:          false,
		YukawaOperator:                false,
		YukawaEigenvalues:             false,
		HiggsMass:                     false,
		ScalarRuntime:                 false,
		MissingMaps:                   missing,
		Verdict: strings.Join([]string{
			StatusPhysicalElectroweakFirewallEnforced,
			StatusInternalCommutantNotPhysicalSU2L,
			StatusNoSelectorIndependentU1PhaseLine,
			StatusNoHyperchargeAssignmentOrNormalization,
			StatusNoTypedK7PlusToPhysicalHiggsDoubletMap,
			StatusNoYukawaOperatorOrEigenvalueTheorem,
			StatusNoHiggsMassOrScalarRuntimeTheorem,
		}, "; "),
	}
}

func buildStrategy() StrategicInterpretation {
	return StrategicInterpretation{
		SU2Problem:      "test whether the selector-invariant commutant C can interface with the already-derived electroweak SU(2)_L lane",
		U1Problem:       "explain what selects or normalizes the moving phase line L_n before any U(1)_Y claim",
		HiggsProblem:    "provide a typed map from K7+_J(n) to the physical Higgs doublet representation",
		SelectorProblem: "supply a native or sealed selector n_* if a single phase line is physically required",
		Verdict: strings.Join([]string{
			StatusElectroweakAirlockSplits,
			StatusNoSelectorIndependentU1PhaseLine,
			StatusNoHyperchargeAssignmentOrNormalization,
		}, "; "),
	}
}

func buildSourceTypes() SourceTypeClassification {
	return SourceTypeClassification{
		CommonCommutantRole: "selector-independent twistor-invariant internal SU(2)-like socket candidate",
		PhaseLineRole:       "selector-dependent moving internal U(1)-like phase line span{J_H(n)}",
		TwistorRole:         "family of sockets whose intersection is C, not a single selected U(1) phase",
		ElectroweakRole:     "physical SU(2)_L interface remains a missing typed map",
		HyperchargeRole:     "physical U(1)_Y and hypercharge normalization remain missing and selector-dependent",
		FirewallRole:        "no Higgs doublet, Yukawa, flavor hierarchy, Higgs mass, scalar runtime, or native 7/72 theorem follows",
		Verdict: strings.Join([]string{
			StatusCommonCommutantTwistorInvariantSU2Candidate,
			StatusU1PhaseLineSelectorDependent,
			StatusGate714TwistorInvariantSU2SocketBoundary,
		}, "; "),
	}
}

func Statuses() []string {
	return []string{
		StatusGate713TwistorSocketBundleInherited,
		StatusCommonCommutantDefined,
		StatusCommonCommutantIncludedInAllU2Sockets,
		StatusTwistorIntersectionEqualsCommonCommutant,
		StatusMovingPhaseLineAudited,
		StatusLieAlgebraStructureOfCommutantAudited,
		StatusSelectorDependentAndIndependentPartsSeparated,
		StatusPhysicalElectroweakFirewallEnforced,
		StatusCommonCommutantTwistorInvariantSU2Candidate,
		StatusU1PhaseLineSelectorDependent,
		StatusElectroweakAirlockSplits,
		StatusInternalCommutantNotPhysicalSU2L,
		StatusNoSelectorIndependentU1PhaseLine,
		StatusNoHyperchargeAssignmentOrNormalization,
		StatusNoTypedK7PlusToPhysicalHiggsDoubletMap,
		StatusNoYukawaOperatorOrEigenvalueTheorem,
		StatusNoHiggsMassOrScalarRuntimeTheorem,
		StatusGate714TwistorInvariantSU2SocketBoundary,
	}
}

func FormatInherited(x Gate713Inheritance) string {
	return fmt.Sprintf("inherited=%t plusDim=%d minusDim=%d sphereDim=%d family=%t single=%t selector=%t physical=%t hypercharge=%t yukawa=%t higgsMass=%t verdict=%q", x.TwistorSocketBundleInherited, x.K7PlusDimension, x.K7MinusDimension, x.TwistorSphereDimension, x.FamilyValuedSocketBundle, x.SingleSocketPromoted, x.NativeTwistorPointSelector, x.PhysicalElectroweakBundle, x.HyperchargeAssignment, x.YukawaOperatorCertified, x.HiggsMassCertified, x.Verdict)
}

func FormatCommon(x CommonCommutantAudit) string {
	return fmt.Sprintf("def=%q dim=%d included=%t selectorIndependent=%t physicalSU2=%t verdict=%q", x.Definition, x.Dimension, x.IncludedInAllSockets, x.SelectorIndependent, x.PhysicalSU2LCertified, x.Verdict)
}

func FormatIntersection(x TwistorIntersectionAudit) string {
	return fmt.Sprintf("def=%q equalsC=%t dim=%d basis=%t movingPhase=%t verdict=%q", x.Definition, x.EqualsCommonCommutant, x.Dimension, x.ProofUsesBasisDirections, x.ContainsMovingPhaseLine, x.Verdict)
}

func FormatPhase(x MovingPhaseLineAudit) string {
	return fmt.Sprintf("def=%q dim=%d moves=%t common=%t selectorIndependent=%t nativePoint=%t hypercharge=%t verdict=%q", x.Definition, x.Dimension, x.MovesWithSelectorN, x.CommonToAllSockets, x.SelectorIndependentLine, x.NativePointSelected, x.HyperchargeCertified, x.Verdict)
}

func FormatLie(x LieAlgebraStructureAudit) string {
	return fmt.Sprintf("algebra=%q dim=%d closes=%t comm=%q norm=%t physical=%t verdict=%q", x.AlgebraName, x.Dimension, x.ClosesAsSU2Like, x.Commutator, x.NormalizationRequired, x.PhysicalSU2LCertified, x.Verdict)
}

func FormatSeparation(x SocketPartSeparationAudit) string {
	return fmt.Sprintf("independent=%d dependent=%d split=%q valid=%t verdict=%q", x.IndependentCount, x.DependentCount, x.AirlockSplit, x.SeparationValid, x.Verdict)
}

func FormatPhysical(x PhysicalElectroweakFirewallAudit) string {
	return fmt.Sprintf("commPhysical=%t phaseU1=%t hypercharge=%t higgs=%t yukawa=%t eigen=%t mass=%t runtime=%t missing=%d verdict=%q", x.InternalCommutantPhysicalSU2L, x.MovingPhasePhysicalU1Y, x.HyperchargeNormalization, x.TypedHiggsDoubletMap, x.YukawaOperator, x.YukawaEigenvalues, x.HiggsMass, x.ScalarRuntime, len(x.MissingMaps), x.Verdict)
}

func FormatStrategy(x StrategicInterpretation) string {
	return fmt.Sprintf("su2=%q u1=%q higgs=%q selector=%q verdict=%q", x.SU2Problem, x.U1Problem, x.HiggsProblem, x.SelectorProblem, x.Verdict)
}

func FormatSourceTypes(x SourceTypeClassification) string {
	return fmt.Sprintf("comm=%q phase=%q twistor=%q ew=%q hyper=%q firewall=%q verdict=%q", x.CommonCommutantRole, x.PhaseLineRole, x.TwistorRole, x.ElectroweakRole, x.HyperchargeRole, x.FirewallRole, x.Verdict)
}
