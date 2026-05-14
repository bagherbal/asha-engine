package condensate

import (
	"fmt"

	"github.com/bagherbal/asha-engine/pkg/theorem"
)

func CompositeHiggsCondensateDirectionAuditTheorem() theorem.Theorem {
	const id = "BRIDGE-COMPOSITE-HIGGS-CONDENSATE-DIRECTION"
	const name = "composite Higgs / fermion-condensate direction audit"
	return theorem.Theorem{ID: id, Name: name, Layer: theorem.LayerBridge, Status: theorem.BridgeRequired, Verify: func() theorem.Result {
		a, err := BuildDefault()
		if err != nil {
			return theorem.Result{ID: id, Name: name, Layer: theorem.LayerBridge, Status: theorem.FailedRoute, Checks: []theorem.Check{{Name: "construct condensate audit", Passed: false, Detail: err.Error()}}}
		}
		return theorem.Result{ID: id, Name: name, Layer: theorem.LayerBridge, Status: theorem.BridgeRequired, Checks: []theorem.Check{
			{Name: "Fock matter substrate", Passed: a.FockModes == 4 && a.FockStates == 16 && a.VacuumNeutral, Detail: fmt.Sprintf("modes=%d, states=%d, temporal=%d, spatial=%d, |Ω⟩ B−L neutral=%v", a.FockModes, a.FockStates, a.TemporalModes, a.SpatialModes, a.VacuumNeutral)},
			{Name: "finite scalar/contact doublet substrate", Passed: a.ScalarActiveRealDimension == 4 && a.ScalarComplexDimension == 2, Detail: fmt.Sprintf("active real directions=%d, complex components=%d, radial+angular=%d+%d", a.ScalarActiveRealDimension, a.ScalarComplexDimension, a.ScalarRadialDirections, a.ScalarAngularDirections)},
			{Name: "gauge-compatible bilinear channels", Passed: a.BilinearScalarCandidateAvailable, Detail: fmt.Sprintf("minimal left×scalar→right channels=%d; leptonic=%d, up-color=%d, down-color=%d", a.GaugeCompatibleYukawaChannels, a.LeptonicChannels, a.UpTypeColorChannels, a.DownTypeColorChannels)},
			{Name: "three-color amplification available", Passed: a.ColorAmplificationAvailable, Detail: fmt.Sprintf("spatial Fock modes=%d, up channels=%d, down channels=%d", a.SpatialModes, a.UpTypeColorChannels, a.DownTypeColorChannels)},
			{Name: "composite-Higgs direction preferred", Passed: a.CompositeHiggsDirectionPreferred, Detail: "the next physics calculation should be a native scalar bilinear/condensate computation, not another scalar-frame orientation convention"},
			{Name: "native one-loop potential", Passed: a.NativeOneLoopPotentialComputed, Detail: "open; finite Fock/Yukawa loop computation not yet implemented"},
			{Name: "NJL / gap equation scale", Passed: a.NJLGapEquationSolved && a.CondensationScaleDerived, Detail: "open; no four-fermion kernel or non-fitted condensate scale derived"},
			{Name: "CPT boundary vacuum selection", Passed: a.CPTBoundarySelectionComputed, Detail: "open; boundary condition has not yet selected potential parameters or condensate orientation"},
		}, Notes: []string{
			a.TruthStatement,
			fmt.Sprintf("critical computations: %s", FormatComputations(a.CriticalComputations)),
			fmt.Sprintf("recommended next gate: %s", a.RecommendedNextGate),
			fmt.Sprintf("remaining unknowns: %s", FormatUnknowns(a.RemainingUnknowns)),
		}}
	}}
}
