package spectralgraphf0index

import (
	"math"
	"strings"

	"github.com/bagherbal/asha-engine/pkg/theorem"
)

func SpectralGraphProjectionF0IndexTheoremSieveTheorem() theorem.Theorem {
	const id = "BRIDGE-SPECTRAL-GRAPH-F0-INDEX-SIEVE"
	const name = "Spectral Graph Projection / f0 Index Theorem Sieve"
	return theorem.Theorem{ID: id, Name: name, Layer: theorem.LayerBridge, Status: theorem.BridgeRequired, Verify: func() theorem.Result {
		a, err := BuildDefault()
		if err != nil {
			return theorem.Result{ID: id, Name: name, Layer: theorem.LayerBridge, Status: theorem.FailedRoute, Checks: []theorem.Check{{Name: "build spectral graph f0 index ledger", Passed: false, Detail: err.Error()}}}
		}
		c := a.Calculation
		checks := []theorem.Check{
			{Name: "finite Dirac edge projection is formalized", Passed: c.Executed && c.EdgeProjection.IsProjectionNative && c.EdgeProjection.FundamentalEdgeCount == 5, Detail: c.EdgeProjection.ProjectionFormula},
			{Name: "discrete edge projection trace equals ten", Passed: math.Abs(c.EdgeProjection.ProjectionTraceOnEdges-10) < 1e-12 && c.EdgeProjection.JDoubledEdgeSlotCount == 10, Detail: c.EdgeProjection.Verdict},
			{Name: "edge-slot trace is not mis-typed as Tr_HF without an embedding theorem", Passed: !c.EdgeProjection.IsTraceOverHFWellTyped && strings.Contains(c.EdgeProjection.ProjectionTraceOnHF, "not well-typed"), Detail: c.EdgeProjection.ProjectionTraceOnHF},
			{Name: "CCM f0 moment definition is audited", Passed: strings.Contains(c.Moment.CCMDefinition, "f(0)") && c.Moment.SharpCutoffValue == 1 && !c.Moment.SameMathematicalObject, Detail: c.Moment.Verdict},
			{Name: "index theorem analogy does not prove all-edge count", Passed: !c.Index.KernelIndexDerived && !c.Index.AllEdgeCountIndex && !c.Index.CanIdentifyF0WithIndex, Detail: c.Index.Verdict},
			{Name: "f0=10 Higgs near-closure is inherited but not sealed", Passed: math.Abs(c.Higgs.F0Candidate-10) < 1e-12 && math.Abs(c.Higgs.MassPfaffianVEVGeV-HiggsBoundaryGeV) < 0.3 && !c.Higgs.GeometricallySealed && !c.HiggsMassSealed, Detail: c.Higgs.Verdict},
			{Name: "f0 moment index theorem remains open", Passed: !c.F0MomentIndexDerived && strings.Contains(StatusLine(c), StatusFailedF0MomentIndexNotDerived) && strings.Contains(StatusLine(c), StatusFailedEdgeProjectionNotCCMF0), Detail: c.Truth},
			{Name: "full numerical ToE closure is not claimed", Passed: !c.FullNumericalTOEClosure && strings.Contains(StatusLine(c), StatusFailedFullNumericalTOEClosureOpen), Detail: c.Truth},
		}
		return theorem.Result{ID: id, Name: name, Layer: theorem.LayerBridge, Status: theorem.BridgeRequired, Checks: checks, Notes: []string{c.Truth}}
	}}
}
