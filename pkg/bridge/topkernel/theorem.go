package topkernel

import (
	"fmt"

	"github.com/bagherbal/asha-engine/pkg/theorem"
)

func TopLikeOverlapKernelSearchTheorem() theorem.Theorem {
	const id = "BRIDGE-TOP-LIKE-OVERLAP-KERNEL-SEARCH"
	const name = "top-like overlap / condensate kernel search"
	return theorem.Theorem{ID: id, Name: name, Layer: theorem.LayerBridge, Status: theorem.BridgeRequired, Verify: func() theorem.Result {
		a, err := BuildDefault()
		if err != nil {
			return theorem.Result{ID: id, Name: name, Layer: theorem.LayerBridge, Status: theorem.FailedRoute, Checks: []theorem.Check{{Name: "construct top-like kernel search", Passed: false, Detail: err.Error()}}}
		}
		return theorem.Result{ID: id, Name: name, Layer: theorem.LayerBridge, Status: theorem.BridgeRequired, Checks: []theorem.Check{
			{Name: "unit incidence kernel is equal-weight", Passed: a.UnitRightRowNormsEqual, Detail: fmt.Sprintf("right-channel row norm²=%.10f for every one-generation right state", a.UnitRowNorm)},
			{Name: "three-color amplification skeleton", Passed: a.QuarkLeptonAmplification == 3, Detail: fmt.Sprintf("quark fiber entries/gen=%d, lepton fiber entries/gen=%d, amplification=%.1f", a.QuarkFiberEntriesPerGen, a.LeptonFiberEntriesPerGen, a.QuarkLeptonAmplification)},
			{Name: "up/down degeneracy exposed", Passed: a.UpDownDegeneracyResidual == 0, Detail: "up-type and down-type quark incidence pressures are equal; topology alone has not selected top over bottom"},
			{Name: "diagonal generation kernel available", Passed: a.DiagonalGenerationKernelFound, Detail: fmt.Sprintf("weights=%s, normalized=%s, spread=%.10f", FormatFloatSlice(a.GenerationWeights), FormatFloatSlice(a.NormalizedGenerationWeights), a.GenerationWeightSpread)},
			{Name: "generation kernel canonical", Passed: a.GenerationKernelCanonical, Detail: "open; the diagonal spurion is finite and natural, but its map to physical generation labels is not yet canonical"},
			{Name: "non-uniform overlap ingredients", Passed: a.NonUniformOverlapKernelFound, Detail: fmt.Sprintf("candidate inventory includes %s", FormatCandidate(a.BestCandidate))},
			{Name: "top-like channel selected", Passed: a.TopLikeChannelSelected, Detail: "open; current ingredients select quark amplification and a heaviest generation direction, but not a unique top-like up channel"},
			{Name: "top-dominance kernel derived", Passed: a.TopDominanceKernelDerived, Detail: "open; no finite attractive kernel yet produces y_top-like dominance"},
			{Name: "condensate strength", Passed: a.CondensateStrengthDerived && a.NativeGapKernelDerived, Detail: "open; no NJL/gap-kernel criticality or condensate scale is derived"},
			{Name: "hidden observed couplings", Passed: !a.HiddenObservedCouplingsUsed, Detail: "no observed y_t, v, Higgs mass, or fermion masses were inserted"},
		}, Notes: []string{
			a.TruthStatement,
			fmt.Sprintf("kind kernels: %s", FormatKindKernels(a.KindKernels)),
			fmt.Sprintf("recommended next gate: %s", a.RecommendedNextGate),
			fmt.Sprintf("remaining unknowns: %s", FormatUnknowns(a.RemainingUnknowns)),
		}}
	}}
}
