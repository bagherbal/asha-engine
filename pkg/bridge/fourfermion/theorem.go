package fourfermion

import (
	"fmt"

	"github.com/bagherbal/asha-engine/pkg/theorem"
)

func NativeFourFermionKernelTheorem() theorem.Theorem {
	const id = "BRIDGE-NATIVE-FOUR-FERMION-KERNEL"
	const name = "native four-fermion kernel from x∧p/u(4) sector"
	return theorem.Theorem{ID: id, Name: name, Layer: theorem.LayerBridge, Status: theorem.BridgeRequired, Verify: func() theorem.Result {
		a, err := BuildDefault()
		if err != nil {
			return theorem.Result{ID: id, Name: name, Layer: theorem.LayerBridge, Status: theorem.FailedRoute, Checks: []theorem.Check{{Name: "construct native four-fermion kernel audit", Passed: false, Detail: err.Error()}}}
		}
		return theorem.Result{ID: id, Name: name, Layer: theorem.LayerBridge, Status: theorem.BridgeRequired, Checks: []theorem.Check{
			{Name: "x∧p/u(4) current inventory", Passed: a.CurrentAlgebraAvailable && a.U4Dimension == 16, Detail: fmt.Sprintf("dim u(4)=%d from %d Fock modes", a.U4Dimension, a.FockModes)},
			{Name: "Pati-Salam-shaped decomposition", Passed: a.DecompositionComplete, Detail: fmt.Sprintf("u(4)=%d central + %d su(3)c + %d B-L + %d leptoquark = %d", a.CentralU1Dimension, a.ColorSU3Dimension, a.BLDimension, a.LeptoquarkDimension, a.U4Dimension)},
			{Name: "formal current-current template", Passed: a.CurrentCurrentTemplateAvailable, Detail: "L_eff ∼ -g_A²(J_AJ_A)/M_A² is a formal template; normalization is open"},
			{Name: "scalar left-right channel available", Passed: a.ScalarLRChannelAvailable, Detail: fmt.Sprintf("finite pressure skeleton inherited: %s", a.KnownFiniteSkeleton)},
			{Name: "finite Fierz projection", Passed: a.FierzProjectionDerived, Detail: "open; current-current → scalar bilinear coefficients c_A are not computed"},
			{Name: "attractive scalar-channel sign", Passed: a.AttractiveChannelSignDerived, Detail: "open; attraction cannot be assumed from generator count alone"},
			{Name: "four-fermion strength G_hat", Passed: a.FourFermionStrengthDerived && a.NativeNJLKernelDerived, Detail: fmt.Sprintf("open; %s", a.FormalKernelExpression)},
			{Name: "up/down splitting", Passed: a.UpDownSplittingDerived, Detail: "open; x∧p/u(4) inventory still does not select top over bottom"},
			{Name: "regulator and criticality", Passed: a.RegulatorDerived && a.CriticalityClosed, Detail: "open; no C_reg or gap solution is derived"},
			{Name: "hidden observed input", Passed: !a.HiddenObservedCouplingsUsed && !a.HiddenMassScaleUsed, Detail: "no observed y_t, v, Higgs mass, or mass scale was inserted"},
		}, Notes: []string{
			a.TruthStatement,
			fmt.Sprintf("u(4) sectors: %s", FormatSectors(a.Sectors)),
			fmt.Sprintf("kernel conditions: %s", FormatConditions(a.KernelConditions)),
			fmt.Sprintf("recommended next gate: %s", a.RecommendedNextGate),
			fmt.Sprintf("remaining unknowns: %s", FormatUnknowns(a.RemainingUnknowns)),
		}}
	}}
}
