// Package generation2higgsradialselectorsourcecandidateandvacuumdirectionfirewallaudit implements
// Gate 737: Higgs Radial Selector Source-Candidate and Vacuum-Direction Firewall Audit.
//
// Gate 736 source-typed rho_plus=I_K7+/4 as the maximum-entropy observer state
// on K7+ and showed that any supplied rank-one radial projector has no-bias
// weight 1/4. Gate 737 audits whether any currently typed ASHA object selects the
// radial projector itself. The verdict preserves that P_rad remains a type-distinct
// scalar/vacuum-direction seal: rho_plus weighs rank-one events but does not choose
// one, n selects complex structure but not a radial line, q normalizes phase charge
// but not direction, and scalar bridge data contain no vector in K7+.
package generation2higgsradialselectorsourcecandidateandvacuumdirectionfirewallaudit

import (
	"fmt"
	"strings"
	"sync"

	gate736 "github.com/bagherbal/asha-engine/pkg/bridge/generation2k7plusmaximumentropyobserverstateandradialeventweightaudit"
)

const (
	AuditID = "GATE737-HIGGS-RADIAL-SELECTOR-SOURCE-CANDIDATE-VACUUM-DIRECTION-FIREWALL-AUDIT"

	StatusGate736K7PlusMaximumEntropyObserverInherited = "PASS_GATE736_K7_PLUS_MAXIMUM_ENTROPY_OBSERVER_INHERITED"
	StatusRadialSelectorProblemDefined                 = "PASS_RADIAL_SELECTOR_PROBLEM_DEFINED"
	StatusCandidateSourceAuditCompleted                = "PASS_CANDIDATE_SOURCE_AUDIT_COMPLETED"
	StatusSymmetryObstructionAudited                   = "PASS_SYMMETRY_OBSTRUCTION_AUDITED"
	StatusSealClassificationDefined                    = "PASS_SEAL_CLASSIFICATION_DEFINED"
	StatusHistoryLoopDependenceOnPRadRecorded          = "PASS_HISTORYLOOP_DEPENDENCE_ON_P_RAD_RECORDED"
	StatusPhysicalFirewallsEnforced                    = "PASS_PHYSICAL_FIREWALLS_ENFORCED"

	StatusPRadTypeDistinctScalarVacuumDirectionSeal    = "CONDITIONAL_SUPPORT_P_RAD_IS_TYPE_DISTINCT_SCALAR_VACUUM_DIRECTION_SEAL"
	StatusPRadRequiresSymmetryBreakingOrVacuumSelector = "CONDITIONAL_SUPPORT_P_RAD_REQUIRES_SYMMETRY_BREAKING_OR_VACUUM_SELECTOR"

	StatusRhoPlusDoesNotSelectPRad                    = "FAILED_ROUTE_RHO_PLUS_DOES_NOT_SELECT_P_RAD"
	StatusTwistorSelectorNDoesNotSelectPRad           = "FAILED_ROUTE_TWISTOR_SELECTOR_N_DOES_NOT_SELECT_P_RAD"
	StatusQDoesNotSelectPRad                          = "FAILED_ROUTE_Q_DOES_NOT_SELECT_P_RAD"
	StatusHodgePolarityDoesNotSelectLineInsideK7Plus  = "FAILED_ROUTE_HODGE_POLARITY_DOES_NOT_SELECT_LINE_INSIDE_K7_PLUS"
	StatusQuaternionicFanoStructureDoesNotSelectPRad  = "FAILED_ROUTE_QUATERNIONIC_FANO_STRUCTURE_DOES_NOT_SELECT_P_RAD"
	StatusBoundaryScalarDataDoNotSelectPRad           = "FAILED_ROUTE_BOUNDARY_SCALAR_DATA_DO_NOT_SELECT_P_RAD"
	StatusNoNativeRadialProjectorSelectorFound        = "FAILED_ROUTE_NO_NATIVE_RADIAL_PROJECTOR_SELECTOR_FOUND"
	StatusHistoryLoopUnitSourceConditionalWithoutPRad = "FAILED_ROUTE_HISTORYLOOPUNIT_SOURCE_REMAINS_CONDITIONAL_WITHOUT_P_RAD"
	StatusNoNativeElectroweakSymmetryBreakingTheorem  = "FAILED_ROUTE_NO_NATIVE_ELECTROWEAK_SYMMETRY_BREAKING_THEOREM"
	StatusNoHiggsMassOrPoleMassTheorem                = "FAILED_ROUTE_NO_HIGGS_MASS_OR_POLE_MASS_THEOREM"
	StatusNoYukawaOperatorOrEigenvalueTheorem         = "FAILED_ROUTE_NO_YUKAWA_OPERATOR_OR_EIGENVALUE_THEOREM"
	StatusGate737Boundary                             = "FIREWALL_PRESERVED_GATE737_RADIAL_SELECTOR_BOUNDARY"
)

type Gate736Inheritance struct {
	Inherited                bool
	RhoPlusMaxEntropy        bool
	RadialWeightCertified    bool
	RhoPlusDoesNotSelectPRad bool
	RhoPlusDoesNotSelectN    bool
	HistoryLoopConditional   bool
	NoHiggsMassTheorem       bool
	NoYukawaTheorem          bool
	Verdict                  string
}

type RadialSelectorProblem struct {
	DesiredProjector      string
	Carrier               string
	Rank                  int
	NeedsLineInsideK7Plus bool
	CurrentlyNative       bool
	Verdict               string
}

type SourceCandidate struct {
	Name           string
	Type           string
	SuppliesPRad   bool
	ReasonRejected string
	Verdict        string
}

type CandidateSourceAudit struct {
	Candidates                   []SourceCandidate
	AnyNativeSelectorFound       bool
	BoundaryScalarsContainVector bool
	Verdict                      string
}

type SymmetryObstructionAudit struct {
	SymmetricData            []string
	CandidateSymmetry        string
	PRadWouldBreakSymmetryTo string
	RequiresVacuumSelector   bool
	CurrentDataSelectsLine   bool
	Verdict                  string
}

type SealClassificationAudit struct {
	SealNames           []string
	DistinctFromN       bool
	DistinctFromQ       bool
	DistinctFromRhoPlus bool
	TypeRole            string
	Verdict             string
}

type HistoryLoopDependenceAudit struct {
	RhoPlusSuppliesWeight  bool
	PRadSuppliesEvent      bool
	NSuppliesPhaseLoop     bool
	QSuppliesPRad          bool
	LFormula               string
	ConditionalWithoutPRad bool
	Verdict                string
}

type PhysicalFirewallAudit struct {
	PRadIsHiggsVacuumTheorem          bool
	PRadIsElectroweakBreakingTheorem  bool
	PhaseTransverseGoldstoneTheorem   bool
	PRadIsHiggsMassTheorem            bool
	YukawaOperatorOrEigenvalueTheorem bool
	Verdict                           string
}

type Analysis struct {
	Gate736     Gate736Inheritance
	Problem     RadialSelectorProblem
	Candidates  CandidateSourceAudit
	Symmetry    SymmetryObstructionAudit
	Seal        SealClassificationAudit
	HistoryLoop HistoryLoopDependenceAudit
	Firewall    PhysicalFirewallAudit
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
	g736, err := gate736.BuildDefault()
	if err != nil {
		return Analysis{}, fmt.Errorf("Gate736 inheritance unavailable: %w", err)
	}
	inherited := buildGate736Inheritance(g736)
	problem := buildProblem()
	candidates := buildCandidateSourceAudit()
	symmetry := buildSymmetryObstruction()
	seal := buildSealClassification()
	historyLoop := buildHistoryLoopDependence()
	firewall := buildPhysicalFirewall()
	truth := "Gate 737 audits the source problem for the rank-one Higgs radial projector P_rad. It finds no native selector among rho_plus, the twistor selector n, hypercharge normalization q, Hodge polarity, quaternionic/Fano structure, boundary scalar data, P_K7, or lambda_proxy. P_rad is therefore classified as a type-distinct scalar/vacuum-direction seal candidate: it is required for the Radial-Hopf source law, but it is not selected by the current symmetric K7+ data and no Higgs, electroweak-breaking, Goldstone, mass, or Yukawa theorem follows."
	return Analysis{Gate736: inherited, Problem: problem, Candidates: candidates, Symmetry: symmetry, Seal: seal, HistoryLoop: historyLoop, Firewall: firewall, Truth: truth}, nil
}

func buildGate736Inheritance(g gate736.Analysis) Gate736Inheritance {
	return Gate736Inheritance{
		Inherited:                g.Entropy.UniqueMaximumEntropy && g.Radial.Rank == 1 && !g.Radial.RhoPlusSelectsEvent,
		RhoPlusMaxEntropy:        g.Entropy.UniqueMaximumEntropy,
		RadialWeightCertified:    g.Radial.Weight == 0.25,
		RhoPlusDoesNotSelectPRad: !g.Radial.RhoPlusSelectsEvent && !g.Selectors.RhoPlusSelectsPRad,
		RhoPlusDoesNotSelectN:    !g.Selectors.RhoPlusSelectsN,
		HistoryLoopConditional:   !g.HistoryLoop.NativeTransportTheorem,
		NoHiggsMassTheorem:       !g.Firewall.HiggsMassOrPoleMassTheorem,
		NoYukawaTheorem:          !g.Firewall.YukawaOperatorOrEigenvalueTheorem,
		Verdict:                  StatusGate736K7PlusMaximumEntropyObserverInherited,
	}
}

func buildProblem() RadialSelectorProblem {
	return RadialSelectorProblem{
		DesiredProjector:      "P_rad^2=P_rad, P_rad^T=P_rad, rank(P_rad)=1",
		Carrier:               "K7+",
		Rank:                  1,
		NeedsLineInsideK7Plus: true,
		CurrentlyNative:       false,
		Verdict:               StatusRadialSelectorProblemDefined,
	}
}

func buildCandidateSourceAudit() CandidateSourceAudit {
	cs := []SourceCandidate{
		{Name: "rho_plus", Type: "maximum-entropy observer state on K7+", SuppliesPRad: false, ReasonRejected: "isotropic state assigns equal weight to every rank-one event; selects no preferred line", Verdict: StatusRhoPlusDoesNotSelectPRad},
		{Name: "TwistorSelectorSeal n", Type: "unit direction in K7- selecting J_H(n)", SuppliesPRad: false, ReasonRejected: "selects complex structure / phase rule, not a real radial vector in K7+", Verdict: StatusTwistorSelectorNDoesNotSelectPRad},
		{Name: "HyperchargeNormalizationSeal q", Type: "scalar normalization of phase generator", SuppliesPRad: false, ReasonRejected: "rescales charge convention; carries no K7+ direction", Verdict: StatusQDoesNotSelectPRad},
		{Name: "K7 Hodge polarity", Type: "K7=K7+⊕K7- split", SuppliesPRad: false, ReasonRejected: "separates sectors but gives no line inside K7+", Verdict: StatusHodgePolarityDoesNotSelectLineInsideK7Plus},
		{Name: "Quaternionic/Fano structure", Type: "J_1,J_2,J_3 and twistor family", SuppliesPRad: false, ReasonRejected: "Sp(1)/SO(3)-covariant structure selects no vector", Verdict: StatusQuaternionicFanoStructureDoesNotSelectPRad},
		{Name: "boundary scalar data lambda, S_split, W_3", Type: "scalar bridge coordinates", SuppliesPRad: false, ReasonRejected: "scalars contain no vector in K7+", Verdict: StatusBoundaryScalarDataDoNotSelectPRad},
		{Name: "P_K7", Type: "rank-seven event projector in H72", SuppliesPRad: false, ReasonRejected: "selects full K7 event, not a rank-one line inside K7+", Verdict: StatusNoNativeRadialProjectorSelectorFound},
		{Name: "lambda_proxy", Type: "scalar proxy lane quantity", SuppliesPRad: false, ReasonRejected: "scalar coefficient lane, no radial direction", Verdict: StatusNoNativeRadialProjectorSelectorFound},
	}
	return CandidateSourceAudit{
		Candidates:                   cs,
		AnyNativeSelectorFound:       false,
		BoundaryScalarsContainVector: false,
		Verdict: strings.Join([]string{
			StatusCandidateSourceAuditCompleted,
			StatusRhoPlusDoesNotSelectPRad,
			StatusTwistorSelectorNDoesNotSelectPRad,
			StatusQDoesNotSelectPRad,
			StatusHodgePolarityDoesNotSelectLineInsideK7Plus,
			StatusQuaternionicFanoStructureDoesNotSelectPRad,
			StatusBoundaryScalarDataDoNotSelectPRad,
			StatusNoNativeRadialProjectorSelectorFound,
		}, "; "),
	}
}

func buildSymmetryObstruction() SymmetryObstructionAudit {
	return SymmetryObstructionAudit{
		SymmetricData:            []string{"rho_plus", "K7+ metric", "quaternionic/Fano twistor family", "internal U(2)-socket before radial choice"},
		CandidateSymmetry:        "quaternionic / U(2)-socket symmetry on K7+",
		PRadWouldBreakSymmetryTo: "stabilizer of a unit radial vector",
		RequiresVacuumSelector:   true,
		CurrentDataSelectsLine:   false,
		Verdict: strings.Join([]string{
			StatusSymmetryObstructionAudited,
			StatusPRadRequiresSymmetryBreakingOrVacuumSelector,
			StatusNoNativeRadialProjectorSelectorFound,
		}, "; "),
	}
}

func buildSealClassification() SealClassificationAudit {
	return SealClassificationAudit{
		SealNames:           []string{"HiggsRadialSelectorSeal", "ScalarVacuumDirectionSeal", "RadialModeProjectionSeal"},
		DistinctFromN:       true,
		DistinctFromQ:       true,
		DistinctFromRhoPlus: true,
		TypeRole:            "rank-one line / radial-vacuum direction inside K7+",
		Verdict: strings.Join([]string{
			StatusSealClassificationDefined,
			StatusPRadTypeDistinctScalarVacuumDirectionSeal,
		}, "; "),
	}
}

func buildHistoryLoopDependence() HistoryLoopDependenceAudit {
	return HistoryLoopDependenceAudit{
		RhoPlusSuppliesWeight:  true,
		PRadSuppliesEvent:      true,
		NSuppliesPhaseLoop:     true,
		QSuppliesPRad:          false,
		LFormula:               "L=Tr(rho_plus [(1/(2*pi))P_rad])=1/(8*pi)",
		ConditionalWithoutPRad: true,
		Verdict: strings.Join([]string{
			StatusHistoryLoopDependenceOnPRadRecorded,
			StatusHistoryLoopUnitSourceConditionalWithoutPRad,
		}, "; "),
	}
}

func buildPhysicalFirewall() PhysicalFirewallAudit {
	return PhysicalFirewallAudit{
		PRadIsHiggsVacuumTheorem:          false,
		PRadIsElectroweakBreakingTheorem:  false,
		PhaseTransverseGoldstoneTheorem:   false,
		PRadIsHiggsMassTheorem:            false,
		YukawaOperatorOrEigenvalueTheorem: false,
		Verdict: strings.Join([]string{
			StatusPhysicalFirewallsEnforced,
			StatusNoNativeElectroweakSymmetryBreakingTheorem,
			StatusNoHiggsMassOrPoleMassTheorem,
			StatusNoYukawaOperatorOrEigenvalueTheorem,
			StatusGate737Boundary,
		}, "; "),
	}
}

func Statuses() []string {
	return []string{
		StatusGate736K7PlusMaximumEntropyObserverInherited,
		StatusRadialSelectorProblemDefined,
		StatusCandidateSourceAuditCompleted,
		StatusSymmetryObstructionAudited,
		StatusSealClassificationDefined,
		StatusHistoryLoopDependenceOnPRadRecorded,
		StatusPhysicalFirewallsEnforced,
		StatusPRadTypeDistinctScalarVacuumDirectionSeal,
		StatusPRadRequiresSymmetryBreakingOrVacuumSelector,
		StatusRhoPlusDoesNotSelectPRad,
		StatusTwistorSelectorNDoesNotSelectPRad,
		StatusQDoesNotSelectPRad,
		StatusHodgePolarityDoesNotSelectLineInsideK7Plus,
		StatusQuaternionicFanoStructureDoesNotSelectPRad,
		StatusBoundaryScalarDataDoNotSelectPRad,
		StatusNoNativeRadialProjectorSelectorFound,
		StatusHistoryLoopUnitSourceConditionalWithoutPRad,
		StatusNoNativeElectroweakSymmetryBreakingTheorem,
		StatusNoHiggsMassOrPoleMassTheorem,
		StatusNoYukawaOperatorOrEigenvalueTheorem,
		StatusGate737Boundary,
	}
}

func FormatGate736(x Gate736Inheritance) string {
	return fmt.Sprintf("inherited=%t maxEnt=%t radialWeight=%t rhoNoPRad=%t rhoNoN=%t conditionalL=%t noMass=%t noYukawa=%t verdict=%q", x.Inherited, x.RhoPlusMaxEntropy, x.RadialWeightCertified, x.RhoPlusDoesNotSelectPRad, x.RhoPlusDoesNotSelectN, x.HistoryLoopConditional, x.NoHiggsMassTheorem, x.NoYukawaTheorem, x.Verdict)
}
func FormatProblem(x RadialSelectorProblem) string {
	return fmt.Sprintf("projector=%q carrier=%q rank=%d needsLine=%t native=%t verdict=%q", x.DesiredProjector, x.Carrier, x.Rank, x.NeedsLineInsideK7Plus, x.CurrentlyNative, x.Verdict)
}
func FormatCandidates(x CandidateSourceAudit) string {
	parts := make([]string, 0, len(x.Candidates))
	for _, c := range x.Candidates {
		parts = append(parts, fmt.Sprintf("%s supplies=%t reason=%s", c.Name, c.SuppliesPRad, c.ReasonRejected))
	}
	return fmt.Sprintf("anyNative=%t boundaryScalarsVector=%t candidates=[%s] verdict=%q", x.AnyNativeSelectorFound, x.BoundaryScalarsContainVector, strings.Join(parts, " | "), x.Verdict)
}
func FormatSymmetry(x SymmetryObstructionAudit) string {
	return fmt.Sprintf("symmetry=%q data=%q breaksTo=%q requiresVacuum=%t currentSelects=%t verdict=%q", x.CandidateSymmetry, strings.Join(x.SymmetricData, ","), x.PRadWouldBreakSymmetryTo, x.RequiresVacuumSelector, x.CurrentDataSelectsLine, x.Verdict)
}
func FormatSeal(x SealClassificationAudit) string {
	return fmt.Sprintf("seals=%q role=%q distinctN=%t distinctQ=%t distinctRho=%t verdict=%q", strings.Join(x.SealNames, ","), x.TypeRole, x.DistinctFromN, x.DistinctFromQ, x.DistinctFromRhoPlus, x.Verdict)
}
func FormatHistoryLoop(x HistoryLoopDependenceAudit) string {
	return fmt.Sprintf("rhoWeight=%t pRadEvent=%t nPhase=%t qSupplies=%t conditional=%t formula=%q verdict=%q", x.RhoPlusSuppliesWeight, x.PRadSuppliesEvent, x.NSuppliesPhaseLoop, x.QSuppliesPRad, x.ConditionalWithoutPRad, x.LFormula, x.Verdict)
}
func FormatFirewall(x PhysicalFirewallAudit) string {
	return fmt.Sprintf("vacuum=%t ewsb=%t goldstone=%t mass=%t yukawa=%t verdict=%q", x.PRadIsHiggsVacuumTheorem, x.PRadIsElectroweakBreakingTheorem, x.PhaseTransverseGoldstoneTheorem, x.PRadIsHiggsMassTheorem, x.YukawaOperatorOrEigenvalueTheorem, x.Verdict)
}
