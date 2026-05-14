// Package spectralgraphf0index implements Gate 381:
// Spectral Graph Projection / f0 Index Theorem Sieve.
//
// Gate 380 found a striking near-closure: f0=10 predicts the Higgs mass
// close to 125 GeV and the finite Dirac graph has exactly ten J-doubled edge
// slots. Gate 381 asks the only lawful next question: is the analytic CCM
// spectral-action moment f0 mathematically identical to the discrete edge
// projection trace, or is the equality only a numerical capacity witness?
package spectralgraphf0index

import (
	"fmt"
	"math"
	"strings"
	"sync"

	"github.com/bagherbal/asha-engine/pkg/bridge/ccmpfaffianf0closure"
)

const (
	AuditID = "GATE381-SPECTRAL-GRAPH-F0-INDEX-THEOREM-SIEVE"

	StatusProjectionFormalized          = "CONDITIONAL_SUPPORT_SPECTRAL_EDGE_PROJECTION_FORMALIZED"
	StatusDiscreteTraceComputed         = "CONDITIONAL_SUPPORT_DISCRETE_EDGE_PROJECTION_TRACE_COMPUTED"
	StatusJDoubledTraceEqualsTen        = "CONDITIONAL_SUPPORT_J_DOUBLED_EDGE_PROJECTION_TRACE_EQUALS_TEN"
	StatusHiggsNearClosureInherited     = "CONDITIONAL_SUPPORT_F0_TEN_HIGGS_NEAR_CLOSURE_INHERITED"
	StatusCCMF0DefinitionAudited        = "CONDITIONAL_SUPPORT_CCM_F0_MOMENT_DEFINITION_AUDITED"
	StatusIndexTheoremSieveExecuted     = "CONDITIONAL_SUPPORT_INDEX_THEOREM_SIEVE_EXECUTED"
	StatusAnalyticDiscreteBridgeAudited = "CONDITIONAL_SUPPORT_ANALYTIC_TO_DISCRETE_BRIDGE_AUDITED"

	StatusTensionEdgeSlotsNotHFStates          = "CONDITIONAL_TENSION_EDGE_MODES_ARE_OPERATOR_SLOTS_NOT_HF_STATE_VECTORS"
	StatusTensionCCMF0MomentNotTrace           = "CONDITIONAL_TENSION_CCM_F0_IS_TEST_FUNCTION_MOMENT_NOT_EDGE_TRACE"
	StatusTensionSharpCutoffF0WouldBeOne       = "CONDITIONAL_TENSION_SHARP_CUTOFF_F0_VALUE_WOULD_BE_ONE_NOT_TEN"
	StatusTensionIndexCountsKernelNotAllEdges  = "CONDITIONAL_TENSION_INDEX_THEOREM_COUNTS_KERNEL_OR_CHIRAL_INDEX_NOT_GENERIC_EDGES"
	StatusTensionProjectionTraceIsMultiplicity = "CONDITIONAL_TENSION_EDGE_PROJECTION_TRACE_IS_MULTIPLICITY_FACTOR_NOT_MOMENT_VALUE"

	StatusFailedF0MomentIndexNotDerived     = "FAILED_ROUTE_F0_MOMENT_INDEX_NOT_DERIVED"
	StatusFailedEdgeProjectionNotCCMF0      = "FAILED_ROUTE_EDGE_PROJECTION_TRACE_NOT_CCM_F0"
	StatusFailedHiggsMassNotSealed          = "FAILED_ROUTE_HIGGS_MASS_NOT_GEOMETRICALLY_SEALED"
	StatusFailedSpectralAmplitudeNotNative  = "FAILED_ROUTE_CUTOFF_FUNCTION_AMPLITUDE_TEN_NOT_NATIVE"
	StatusFailedFullNumericalTOEClosureOpen = "FAILED_ROUTE_FULL_NUMERICAL_TOE_CLOSURE_STILL_NOT_REACHED"
)

const (
	TraceRatioEOverA2 = 1197.0 / 4624.0
	HiggsBoundaryGeV  = 125.10
	StandardVEVGeV    = 246.22
)

type EdgeProjection struct {
	FundamentalEdges       []string
	FundamentalEdgeCount   int
	JDoubledEdgeSlotCount  int
	ProjectionTraceOnEdges float64
	ProjectionTraceOnHF    string
	ProjectionFormula      string
	IsProjectionNative     bool
	IsTraceOverHFWellTyped bool
	Verdict                string
}

type MomentDefinition struct {
	CCMDefinition          string
	SharpCutoffValue       float64
	EdgeProjectionTrace    float64
	SameMathematicalObject bool
	WouldRequireTheorem    string
	Verdict                string
}

type IndexSieve struct {
	IndexCandidate         string
	KernelIndexDerived     bool
	AllEdgeCountIndex      bool
	ChiralSignedIndex      string
	CanIdentifyF0WithIndex bool
	Verdict                string
}

type HiggsNearClosure struct {
	F0Candidate            float64
	LambdaH                float64
	MassStandardVEVGeV     float64
	MassPfaffianVEVGeV     float64
	EffectiveF0StandardVEV float64
	EffectiveF0PfaffianVEV float64
	MassErrorPfaffianGeV   float64
	GeometricallySealed    bool
	Verdict                string
}

type Calculation struct {
	Executed                bool
	EdgeProjection          EdgeProjection
	Moment                  MomentDefinition
	Index                   IndexSieve
	Higgs                   HiggsNearClosure
	Statuses                []string
	F0MomentIndexDerived    bool
	HiggsMassSealed         bool
	FullNumericalTOEClosure bool
	Truth                   string
}

type Analysis struct{ Calculation Calculation }

var defaultOnce sync.Once
var defaultA Analysis
var defaultErr error

func BuildDefault() (Analysis, error) {
	defaultOnce.Do(func() { defaultA, defaultErr = Build() })
	return defaultA, defaultErr
}

func Build() (Analysis, error) {
	native := ccmpfaffianf0closure.NativeConstants()
	if len(native) == 0 {
		return Analysis{}, fmt.Errorf("could not read Gate 380 f0 closure constants")
	}

	edges := []string{
		"Q_L ↔ u_R",
		"Q_L ↔ d_R",
		"L_L ↔ e_R",
		"L_L ↔ ν_R",
		"ν_R ↔ ν_R^c",
	}
	projection := EdgeProjection{
		FundamentalEdges:       edges,
		FundamentalEdgeCount:   len(edges),
		JDoubledEdgeSlotCount:  2 * len(edges),
		ProjectionTraceOnEdges: float64(2 * len(edges)),
		ProjectionTraceOnHF:    "not well-typed without an additional theorem embedding edge-slots as an orthogonal H_F subspace; edges are D_F operator entries, not ordinary H_F state vectors",
		ProjectionFormula:      "Tr_E(P_edge)=dim(E_edges)=2×5=10 on the finite Dirac edge-slot space E_edges",
		IsProjectionNative:     true,
		IsTraceOverHFWellTyped: false,
		Verdict:                "The finite Dirac graph natively supplies a ten-dimensional J-doubled edge-slot projection. This is a real topological multiplicity witness, but it is a trace over edge/operator slots, not automatically the Hilbert-space trace Tr_{H_F} nor the CCM moment f₀.",
	}

	moment := MomentDefinition{
		CCMDefinition:          "In the CCM spectral action, f₀ is the zeroth test-function moment/value f(0) entering the heat-kernel coefficient ledger; for a unit sharp cutoff, f(0)=1.",
		SharpCutoffValue:       1,
		EdgeProjectionTrace:    projection.ProjectionTraceOnEdges,
		SameMathematicalObject: false,
		WouldRequireTheorem:    "A new SpectralGraphMomentTheorem proving f(0) is replaced by, or canonically paired with, Tr_E(P_edge)=10 for the finite Dirac edge graph. This theorem is not present in the current ledger.",
		Verdict:                "The equality f₀=Tr_E(P_edge)=10 is not a consequence of the CCM definition. A projection trace can multiply a moment as a degeneracy, but it does not redefine the scalar moment value unless the cutoff functional itself is changed.",
	}

	index := IndexSieve{
		IndexCandidate:         "Atiyah-Singer / finite Fredholm index analogy for D_F edge modes",
		KernelIndexDerived:     false,
		AllEdgeCountIndex:      false,
		ChiralSignedIndex:      "not computed as 10; a true index is a signed kernel or Fredholm index, not the unsigned count of all allowed Yukawa/Majorana edges",
		CanIdentifyF0WithIndex: false,
		Verdict:                "The usual index theorem does not count every allowed finite Dirac edge. It counts a kernel/chiral/Fredholm index. The current ASHA edge count 10 is an unsigned interaction-slot dimension, so it cannot by itself prove f₀=10.",
	}

	f0 := 10.0
	pfMass := native["mh_pred_f0_10_pfaffian_vev"]
	stdMass := native["mh_pred_f0_10_standard_ew_vev"]
	higgs := HiggsNearClosure{
		F0Candidate:            f0,
		LambdaH:                native["lambda_f0_10"],
		MassStandardVEVGeV:     stdMass,
		MassPfaffianVEVGeV:     pfMass,
		EffectiveF0StandardVEV: native["f0_eff_standard_ew_vev"],
		EffectiveF0PfaffianVEV: native["f0_eff_pfaffian_unreduced_planck"],
		MassErrorPfaffianGeV:   pfMass - HiggsBoundaryGeV,
		GeometricallySealed:    false,
		Verdict:                fmt.Sprintf("Inherited from Gate 380: f₀=10 gives λ_H=%.12g, m_H=%.12g GeV with v=246.22 GeV and m_H=%.12g GeV with the Pfaffian VEV. This is near-closure, not a sealed theorem, until f₀=10 is derived as the spectral-action moment.", native["lambda_f0_10"], stdMass, pfMass),
	}

	statuses := []string{
		StatusProjectionFormalized,
		StatusDiscreteTraceComputed,
		StatusJDoubledTraceEqualsTen,
		StatusHiggsNearClosureInherited,
		StatusCCMF0DefinitionAudited,
		StatusIndexTheoremSieveExecuted,
		StatusAnalyticDiscreteBridgeAudited,
		StatusTensionEdgeSlotsNotHFStates,
		StatusTensionCCMF0MomentNotTrace,
		StatusTensionSharpCutoffF0WouldBeOne,
		StatusTensionIndexCountsKernelNotAllEdges,
		StatusTensionProjectionTraceIsMultiplicity,
		StatusFailedF0MomentIndexNotDerived,
		StatusFailedEdgeProjectionNotCCMF0,
		StatusFailedHiggsMassNotSealed,
		StatusFailedSpectralAmplitudeNotNative,
		StatusFailedFullNumericalTOEClosureOpen,
	}

	truth := "Gate 381 proves the precise state of the f₀=10 idea. The finite ASHA Dirac graph really has five structural edge classes and ten J-doubled edge slots, so Tr_E(P_edge)=10 on the edge-slot space. This exactly matches the f₀ value that gives near-Higgs closure in Gate 380. However, CCM f₀ is the zeroth test-function moment/value f(0), while the edge projection trace is a finite multiplicity over D_F operator slots. Edge slots are not automatically H_F eigenvectors, and an Atiyah-Singer index counts a signed kernel/Fredholm index rather than the unsigned set of all Yukawa/Majorana edges. Therefore Gate 381 preserves f₀=10 as a powerful capacity witness, but it does not derive f₀=10 as a native spectral-action moment. Higgs mass geometric sealing remains open."

	return Analysis{Calculation: Calculation{
		Executed:                true,
		EdgeProjection:          projection,
		Moment:                  moment,
		Index:                   index,
		Higgs:                   higgs,
		Statuses:                statuses,
		F0MomentIndexDerived:    false,
		HiggsMassSealed:         false,
		FullNumericalTOEClosure: false,
		Truth:                   truth,
	}}, nil
}

func NativeConstants() map[string]float64 {
	a, err := BuildDefault()
	if err != nil {
		return map[string]float64{}
	}
	c := a.Calculation
	return map[string]float64{
		"edge_projection_trace":        c.EdgeProjection.ProjectionTraceOnEdges,
		"j_doubled_edge_slot_count":    float64(c.EdgeProjection.JDoubledEdgeSlotCount),
		"ccm_sharp_cutoff_f0_value":    c.Moment.SharpCutoffValue,
		"lambda_f0_10":                 c.Higgs.LambdaH,
		"mh_f0_10_standard_ew_vev":     c.Higgs.MassStandardVEVGeV,
		"mh_f0_10_pfaffian_vev":        c.Higgs.MassPfaffianVEVGeV,
		"f0_eff_standard_ew_vev":       c.Higgs.EffectiveF0StandardVEV,
		"f0_eff_pfaffian_unreduced_mp": c.Higgs.EffectiveF0PfaffianVEV,
	}
}

func FormatFloat(x float64) string { return fmt.Sprintf("%.15g", x) }

func StatusLine(c Calculation) string { return strings.Join(c.Statuses, ";") }

func NearlyEqual(a, b, tol float64) bool { return math.Abs(a-b) <= tol }
