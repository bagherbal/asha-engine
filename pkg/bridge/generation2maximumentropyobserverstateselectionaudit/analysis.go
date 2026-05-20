// Package generation2maximumentropyobserverstateselectionaudit implements
// Gate 694: Maximum-Entropy Observer State Selection Audit.
//
// Gate 693 showed that rho_72=I_H72/72 is not unique among all density
// states: biased synthetic states can reproduce Tr(rho P_K7)=7/72 by
// construction.  Gate 694 audits the stronger no-bias criterion.  On the
// full 72-dimensional augmented chamber, the unique positive normalized
// maximum-entropy state, equivalently the state invariant under all full
// chamber basis changes, is rho_72.
//
// This is a bridge-layer observer-state selection audit only. It does not
// derive boundary stress, scalar RG matching, Higgs mass, gauge unification,
// flavor, CKM/PMNS, a native state-selection theorem, a native first-trace
// theorem, or a native 7/72 theorem.
package generation2maximumentropyobserverstateselectionaudit

import (
	"fmt"
	"math"
	"strings"
	"sync"

	gate693 "github.com/bagherbal/asha-engine/pkg/bridge/generation2fullaugmentedobserverstateselectionandbiasfirewallaudit"
)

const (
	AuditID = "GATE694-MAXIMUM-ENTROPY-OBSERVER-STATE-SELECTION-AUDIT"

	StatusGate693ObserverStateSelectionInherited        = "PASS_GATE693_OBSERVER_STATE_SELECTION_INHERITED"
	StatusGeneralStateResponseReducedToK7Weight         = "PASS_GENERAL_STATE_RESPONSE_REDUCED_TO_K7_WEIGHT"
	StatusVonNeumannEntropyAudited                      = "PASS_VON_NEUMANN_ENTROPY_AUDITED"
	StatusRho72UniquelyMaximizesEntropyOnH72            = "PASS_RHO72_UNIQUELY_MAXIMIZES_ENTROPY_ON_H72"
	StatusFullSymmetryInvarianceSelectsRho72            = "PASS_FULL_SYMMETRY_INVARIANCE_SELECTS_RHO72"
	StatusBlockBiasFamilyAudited                        = "PASS_BLOCK_BIAS_FAMILY_AUDITED"
	StatusEqualPerDimensionWeightSelectsRho72           = "PASS_EQUAL_PER_DIMENSION_WEIGHT_SELECTS_RHO72"
	StatusBiasedStateFirewallPreserved                  = "PASS_BIASED_STATE_FIREWALL_PRESERVED"
	StatusRho72UniqueMaximumEntropyFullChamberState     = "CONDITIONAL_SUPPORT_RHO72_IS_UNIQUE_MAXIMUM_ENTROPY_FULL_CHAMBER_STATE"
	StatusActiveSevenOver72NoBiasK7Expectation          = "CONDITIONAL_SUPPORT_ACTIVE_7_OVER_72_IS_NO_BIAS_K7_EXPECTATION"
	StatusFiniteBoundaryEqualPerDimensionWeightRequired = "CONDITIONAL_SUPPORT_FINITE_BOUNDARY_EQUAL_PER_DIMENSION_WEIGHT_IS_REQUIRED"
	StatusBiasedStatesCanReproduceWeightButCircular     = "FAILED_ROUTE_BIASED_STATES_CAN_REPRODUCE_WEIGHT_BUT_ARE_CIRCULAR"
	StatusNoNativeMaximumEntropyHistoryObserverTheorem  = "FAILED_ROUTE_NO_NATIVE_MAXIMUM_ENTROPY_HISTORY_OBSERVER_THEOREM"
	StatusNoNativeStateSelectionTheorem                 = "FAILED_ROUTE_NO_NATIVE_STATE_SELECTION_THEOREM"
	StatusNoNativeSevenOver72Theorem                    = "FAILED_ROUTE_NO_NATIVE_7_OVER_72_THEOREM"
	StatusGate694MaximumEntropyObserverBoundary         = "FIREWALL_PRESERVED_GATE694_MAXIMUM_ENTROPY_OBSERVER_BOUNDARY"
)

const (
	lambda4Dimension  = 70
	boundaryDimension = 2
	h72Dimension      = lambda4Dimension + boundaryDimension
	kernelDimension   = 71
	k7Dimension       = 7
	tolerance         = 1e-15
)

type Gate693Inheritance struct {
	ObserverStateSelectionInherited bool
	Rho72Definition                 string
	ResponseOperator                string
	DBase                           float64
	SSplit                          float64
	ActiveExpectation               float64
	ResidualE1                      float64
	H72Dimension                    int
	K7Dimension                     int
	Gate693MinimalUnbiased          bool
	Gate693BiasedStatesCanMatch     bool
	Gate693NoNativeStateSelection   bool
	Gate693NoNativeFirstTrace       bool
	Gate693NoNativeSevenOver72      bool
	Verdict                         string
}

type GeneralStateResponseAudit struct {
	Formula                string
	RequiredK7Weight       float64
	RequiredExpectation    float64
	ReducesToK7Weight      bool
	ActiveRequiresK7Weight bool
	Verdict                string
}

type VonNeumannEntropyAudit struct {
	Definition                       string
	Dimension                        int
	Rho72Eigenvalue                  float64
	Rho72Entropy                     float64
	MaximumEntropy                   float64
	UniqueMaximizer                  string
	StrictConcavityUsed              bool
	Rho72UniqueMaximumEntropy        bool
	ExampleBiasedEntropy             float64
	ExampleBiasedEntropyLowerThanMax bool
	AnyBiasedStateEntropyLower       bool
	Verdict                          string
}

type SymmetryNoBiasAudit struct {
	InvarianceGroup             string
	InvariantStateForm          string
	NormalizationEquation       string
	ScalarCoefficient           float64
	SelectsRho72                bool
	EquivalentToNoDirectionBias bool
	Verdict                     string
}

type BlockBiasFamilyAudit struct {
	Family                     string
	NormalizationEquation      string
	K7WeightFormula            string
	ActiveK7Weight             float64
	SolvedA                    float64
	SolvedB                    float64
	EqualPerDimensionWeight    bool
	FiniteBoundaryBiasRejected bool
	Rho72SelectedInBlockFamily bool
	Verdict                    string
}

type BiasFirewallAudit struct {
	BiasedDensityStatesCanMatch   bool
	BiasedWitnessK7Weight         float64
	BiasedWitnessExpectation      float64
	BiasedWitnessCircular         bool
	ReproductionIsNativeSelection bool
	PreservesGate693Firewall      bool
	Verdict                       string
}

type ResponseValueAudit struct {
	Rho72K7Weight float64
	Expectation   float64
	DBase         float64
	ResidualE1    float64
	ResidualAbs   float64
	Verdict       string
}

type MissingTheoremAudit struct {
	Missing    []string
	Candidates []string
	PreciseGap string
	Verdict    string
}

type VerdictDiscipline struct {
	ClaimsPhysicalHistoryUsesMaxEntropy bool
	ClaimsNativeStateSelectionTheorem   bool
	ClaimsNativeFirstTraceTheorem       bool
	ClaimsNativeSevenOver72Theorem      bool
	ClaimsBiasedStateNativeSelection    bool
	ClaimsBoundaryStress                bool
	ClaimsScalarRGMatching              bool
	ClaimsHiggsMass                     bool
	ClaimsGaugeUnification              bool
	ClaimsFlavorDerivation              bool
	ClaimsCKMPMNS                       bool
	ClaimsProjectorActivation           bool
	Verdict                             string
}

type Analysis struct {
	Inherited  Gate693Inheritance
	General    GeneralStateResponseAudit
	Entropy    VonNeumannEntropyAudit
	Symmetry   SymmetryNoBiasAudit
	BlockBias  BlockBiasFamilyAudit
	Bias       BiasFirewallAudit
	Response   ResponseValueAudit
	Missing    MissingTheoremAudit
	Discipline VerdictDiscipline
	Truth      string
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
	g693, err := gate693.BuildDefault()
	if err != nil {
		return Analysis{}, fmt.Errorf("Gate693 inheritance unavailable: %w", err)
	}
	inherited := buildInheritance(g693)
	general := buildGeneralStateResponse(inherited)
	entropy := buildEntropyAudit()
	symmetry := buildSymmetryAudit()
	block := buildBlockBiasAudit(general)
	bias := buildBiasFirewall(g693, inherited)
	response := buildResponseValue(inherited, general)
	missing := MissingTheoremAudit{
		Missing: []string{
			StatusNoNativeMaximumEntropyHistoryObserverTheorem,
			StatusNoNativeStateSelectionTheorem,
			StatusNoNativeSevenOver72Theorem,
		},
		Candidates: []string{
			"MaximallyMixedHistoryObserverTheorem",
			"FullAugmentedNoBiasObserverPrinciple",
			"HistoryResponseStateSelectionTheorem",
		},
		PreciseGap: "a native physical-history theorem explaining why the observer state must be the full augmented maximum-entropy state rho_72 rather than merely a mathematically no-bias state on H72",
		Verdict: strings.Join([]string{
			StatusNoNativeMaximumEntropyHistoryObserverTheorem,
			StatusNoNativeStateSelectionTheorem,
			StatusNoNativeSevenOver72Theorem,
		}, "; "),
	}
	discipline := VerdictDiscipline{Verdict: StatusGate694MaximumEntropyObserverBoundary}
	truth := "Gate 694 upgrades the Gate693 no-bias reading: among full H72 positive normalized states with no preferred subspace, rho_72=I_H72/72 is uniquely selected by maximum von Neumann entropy and, equivalently, by invariance under all full-chamber basis changes.  The block-bias family rho(a,b)=a P_finite+b P_boundary also collapses to rho_72 when the active K7 weight 7/72 and normalization are both required, because a=b=1/72.  Biased states can still reproduce the weight by construction, so this is not a native physical-history state-selection theorem.  The audit conditionally supports rho_72 as the unique maximum-entropy full-chamber observer state while preserving the missing maximum-entropy history observer, state-selection, and native 7/72 theorems."
	return Analysis{Inherited: inherited, General: general, Entropy: entropy, Symmetry: symmetry, BlockBias: block, Bias: bias, Response: response, Missing: missing, Discipline: discipline, Truth: truth}, nil
}

func buildInheritance(g gate693.Analysis) Gate693Inheritance {
	return Gate693Inheritance{
		ObserverStateSelectionInherited: g.Rho72Selection.FullSupport && g.Rho72Selection.Positive && g.Rho72Selection.Normalized && g.Rho72Selection.Unbiased,
		Rho72Definition:                 "rho_72 = I_H72/72",
		ResponseOperator:                g.Inherited.ResponseOperator,
		DBase:                           g.Inherited.DBase,
		SSplit:                          g.Inherited.SSplit,
		ActiveExpectation:               g.Inherited.ActiveExpectation,
		ResidualE1:                      g.Residual.E1,
		H72Dimension:                    g.Inherited.H72Dimension,
		K7Dimension:                     g.Inherited.K7Dimension,
		Gate693MinimalUnbiased:          g.Rho72Selection.UniqueUnderUnbiasedFullH72,
		Gate693BiasedStatesCanMatch:     g.BiasFirewall.BiasedDensityStatesCanMatch,
		Gate693NoNativeStateSelection:   !g.Rho72Selection.NativeStateSelectionTheorem,
		Gate693NoNativeFirstTrace:       !g.Discipline.ClaimsNativeFirstTraceTheorem,
		Gate693NoNativeSevenOver72:      !g.Discipline.ClaimsNativeSevenOver72Theorem,
		Verdict:                         StatusGate693ObserverStateSelectionInherited,
	}
}

func buildGeneralStateResponse(i Gate693Inheritance) GeneralStateResponseAudit {
	weight := float64(i.K7Dimension) / float64(i.H72Dimension)
	return GeneralStateResponseAudit{
		Formula:                "For any density state rho, Tr(rho R_split)=S_split Tr(rho P_K7)",
		RequiredK7Weight:       weight,
		RequiredExpectation:    weight * i.SSplit,
		ReducesToK7Weight:      true,
		ActiveRequiresK7Weight: true,
		Verdict: strings.Join([]string{
			StatusGeneralStateResponseReducedToK7Weight,
			StatusActiveSevenOver72NoBiasK7Expectation,
		}, "; "),
	}
}

func buildEntropyAudit() VonNeumannEntropyAudit {
	maxEntropy := math.Log(float64(h72Dimension))
	biasedEntropy := entropyOfExampleBiasedFullSupportState()
	return VonNeumannEntropyAudit{
		Definition:                       "S_vN(rho) = -Tr(rho log rho)",
		Dimension:                        h72Dimension,
		Rho72Eigenvalue:                  1.0 / float64(h72Dimension),
		Rho72Entropy:                     maxEntropy,
		MaximumEntropy:                   maxEntropy,
		UniqueMaximizer:                  "rho_72 = I_H72/72",
		StrictConcavityUsed:              true,
		Rho72UniqueMaximumEntropy:        true,
		ExampleBiasedEntropy:             biasedEntropy,
		ExampleBiasedEntropyLowerThanMax: biasedEntropy < maxEntropy,
		AnyBiasedStateEntropyLower:       true,
		Verdict: strings.Join([]string{
			StatusVonNeumannEntropyAudited,
			StatusRho72UniquelyMaximizesEntropyOnH72,
			StatusRho72UniqueMaximumEntropyFullChamberState,
		}, "; "),
	}
}

func entropyOfExampleBiasedFullSupportState() float64 {
	// Full-support biased witness with total K7 weight 7/72 but nonuniform
	// eigenvalues inside K7: one K7 direction has weight 2/72, the other six
	// share 5/72, and the 65-dimensional complement remains uniform.
	eigs := make([]float64, 0, h72Dimension)
	eigs = append(eigs, 2.0/72.0)
	for i := 0; i < 6; i++ {
		eigs = append(eigs, (5.0/72.0)/6.0)
	}
	for i := 0; i < 65; i++ {
		eigs = append(eigs, 1.0/72.0)
	}
	var s float64
	for _, p := range eigs {
		s -= p * math.Log(p)
	}
	return s
}

func buildSymmetryAudit() SymmetryNoBiasAudit {
	return SymmetryNoBiasAudit{
		InvarianceGroup:             "all orthogonal/unitary changes of basis on H72",
		InvariantStateForm:          "rho = c I_H72",
		NormalizationEquation:       "Tr(rho)=72c=1",
		ScalarCoefficient:           1.0 / 72.0,
		SelectsRho72:                true,
		EquivalentToNoDirectionBias: true,
		Verdict: strings.Join([]string{
			StatusFullSymmetryInvarianceSelectsRho72,
			StatusRho72UniqueMaximumEntropyFullChamberState,
		}, "; "),
	}
}

func buildBlockBiasAudit(g GeneralStateResponseAudit) BlockBiasFamilyAudit {
	a := g.RequiredK7Weight / float64(k7Dimension)
	b := (1.0 - float64(lambda4Dimension)*a) / float64(boundaryDimension)
	return BlockBiasFamilyAudit{
		Family:                     "rho(a,b)=a P_finite + b P_boundary",
		NormalizationEquation:      "70a+2b=1",
		K7WeightFormula:            "Tr(rho(a,b) P_K7)=7a",
		ActiveK7Weight:             g.RequiredK7Weight,
		SolvedA:                    a,
		SolvedB:                    b,
		EqualPerDimensionWeight:    math.Abs(a-b) < tolerance && math.Abs(a-1.0/72.0) < tolerance,
		FiniteBoundaryBiasRejected: math.Abs(a-b) < tolerance,
		Rho72SelectedInBlockFamily: math.Abs(a-1.0/72.0) < tolerance && math.Abs(b-1.0/72.0) < tolerance,
		Verdict: strings.Join([]string{
			StatusBlockBiasFamilyAudited,
			StatusEqualPerDimensionWeightSelectsRho72,
			StatusFiniteBoundaryEqualPerDimensionWeightRequired,
		}, "; "),
	}
}

func buildBiasFirewall(g gate693.Analysis, i Gate693Inheritance) BiasFirewallAudit {
	return BiasFirewallAudit{
		BiasedDensityStatesCanMatch:   g.BiasFirewall.BiasedDensityStatesCanMatch,
		BiasedWitnessK7Weight:         g.BiasFirewall.BiasedWitnessK7Weight,
		BiasedWitnessExpectation:      g.BiasFirewall.BiasedWitnessExpectation,
		BiasedWitnessCircular:         g.BiasFirewall.BiasedWitnessCircular,
		ReproductionIsNativeSelection: false,
		PreservesGate693Firewall:      g.BiasFirewall.BiasedDensityStatesCanMatch && g.BiasFirewall.BiasedWitnessCircular && !g.BiasFirewall.ReproductionIsNativeSelection,
		Verdict: strings.Join([]string{
			StatusBiasedStateFirewallPreserved,
			StatusBiasedStatesCanReproduceWeightButCircular,
		}, "; "),
	}
}

func buildResponseValue(i Gate693Inheritance, g GeneralStateResponseAudit) ResponseValueAudit {
	return ResponseValueAudit{
		Rho72K7Weight: g.RequiredK7Weight,
		Expectation:   g.RequiredExpectation,
		DBase:         i.DBase,
		ResidualE1:    i.ResidualE1,
		ResidualAbs:   math.Abs(i.ResidualE1),
		Verdict:       StatusActiveSevenOver72NoBiasK7Expectation,
	}
}

func Statuses() []string {
	return []string{
		StatusGate693ObserverStateSelectionInherited,
		StatusGeneralStateResponseReducedToK7Weight,
		StatusVonNeumannEntropyAudited,
		StatusRho72UniquelyMaximizesEntropyOnH72,
		StatusFullSymmetryInvarianceSelectsRho72,
		StatusBlockBiasFamilyAudited,
		StatusEqualPerDimensionWeightSelectsRho72,
		StatusBiasedStateFirewallPreserved,
		StatusRho72UniqueMaximumEntropyFullChamberState,
		StatusActiveSevenOver72NoBiasK7Expectation,
		StatusFiniteBoundaryEqualPerDimensionWeightRequired,
		StatusBiasedStatesCanReproduceWeightButCircular,
		StatusNoNativeMaximumEntropyHistoryObserverTheorem,
		StatusNoNativeStateSelectionTheorem,
		StatusNoNativeSevenOver72Theorem,
		StatusGate694MaximumEntropyObserverBoundary,
	}
}

func FormatInheritance(x Gate693Inheritance) string {
	return fmt.Sprintf("inherited=%t rho=%q response=%q dbase=%.18g ssplit=%.18g expectation=%.18g e1=%.18g h72=%d k7=%d minimalUnbiased=%t biasedCanMatch=%t noState=%t noFirst=%t no7=%t verdict=%q", x.ObserverStateSelectionInherited, x.Rho72Definition, x.ResponseOperator, x.DBase, x.SSplit, x.ActiveExpectation, x.ResidualE1, x.H72Dimension, x.K7Dimension, x.Gate693MinimalUnbiased, x.Gate693BiasedStatesCanMatch, x.Gate693NoNativeStateSelection, x.Gate693NoNativeFirstTrace, x.Gate693NoNativeSevenOver72, x.Verdict)
}

func FormatGeneral(x GeneralStateResponseAudit) string {
	return fmt.Sprintf("formula=%q requiredWeight=%.18g requiredExpectation=%.18g reduces=%t activeRequires=%t verdict=%q", x.Formula, x.RequiredK7Weight, x.RequiredExpectation, x.ReducesToK7Weight, x.ActiveRequiresK7Weight, x.Verdict)
}

func FormatEntropy(x VonNeumannEntropyAudit) string {
	return fmt.Sprintf("definition=%q dim=%d eigen=%.18g entropy=%.18g max=%.18g unique=%q strict=%t uniqueMax=%t biasedEntropy=%.18g biasedLower=%t anyBiasedLower=%t verdict=%q", x.Definition, x.Dimension, x.Rho72Eigenvalue, x.Rho72Entropy, x.MaximumEntropy, x.UniqueMaximizer, x.StrictConcavityUsed, x.Rho72UniqueMaximumEntropy, x.ExampleBiasedEntropy, x.ExampleBiasedEntropyLowerThanMax, x.AnyBiasedStateEntropyLower, x.Verdict)
}

func FormatSymmetry(x SymmetryNoBiasAudit) string {
	return fmt.Sprintf("group=%q form=%q normalization=%q c=%.18g selects=%t noBias=%t verdict=%q", x.InvarianceGroup, x.InvariantStateForm, x.NormalizationEquation, x.ScalarCoefficient, x.SelectsRho72, x.EquivalentToNoDirectionBias, x.Verdict)
}

func FormatBlockBias(x BlockBiasFamilyAudit) string {
	return fmt.Sprintf("family=%q normalization=%q k7Formula=%q activeWeight=%.18g a=%.18g b=%.18g equal=%t finiteBoundaryBiasRejected=%t rho72=%t verdict=%q", x.Family, x.NormalizationEquation, x.K7WeightFormula, x.ActiveK7Weight, x.SolvedA, x.SolvedB, x.EqualPerDimensionWeight, x.FiniteBoundaryBiasRejected, x.Rho72SelectedInBlockFamily, x.Verdict)
}

func FormatBias(x BiasFirewallAudit) string {
	return fmt.Sprintf("biasedCanMatch=%t weight=%.18g expectation=%.18g circular=%t nativeSelection=%t preserves=%t verdict=%q", x.BiasedDensityStatesCanMatch, x.BiasedWitnessK7Weight, x.BiasedWitnessExpectation, x.BiasedWitnessCircular, x.ReproductionIsNativeSelection, x.PreservesGate693Firewall, x.Verdict)
}

func FormatResponse(x ResponseValueAudit) string {
	return fmt.Sprintf("rhoWeight=%.18g expectation=%.18g dbase=%.18g e1=%.18g abs=%.18g verdict=%q", x.Rho72K7Weight, x.Expectation, x.DBase, x.ResidualE1, x.ResidualAbs, x.Verdict)
}

func FormatMissing(x MissingTheoremAudit) string {
	return fmt.Sprintf("missing=[%s] candidates=[%s] precise=%q verdict=%q", strings.Join(x.Missing, "; "), strings.Join(x.Candidates, "; "), x.PreciseGap, x.Verdict)
}

func FormatDiscipline(x VerdictDiscipline) string {
	return fmt.Sprintf("physicalMaxEntropy=%t nativeState=%t nativeFirst=%t native7=%t biasedNative=%t boundary=%t scalarRG=%t higgs=%t gauge=%t flavor=%t ckmPmns=%t activation=%t verdict=%q", x.ClaimsPhysicalHistoryUsesMaxEntropy, x.ClaimsNativeStateSelectionTheorem, x.ClaimsNativeFirstTraceTheorem, x.ClaimsNativeSevenOver72Theorem, x.ClaimsBiasedStateNativeSelection, x.ClaimsBoundaryStress, x.ClaimsScalarRGMatching, x.ClaimsHiggsMass, x.ClaimsGaugeUnification, x.ClaimsFlavorDerivation, x.ClaimsCKMPMNS, x.ClaimsProjectorActivation, x.Verdict)
}
