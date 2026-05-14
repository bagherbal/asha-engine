package fullphysicalfirstorder

import "github.com/bagherbal/asha-engine/pkg/theorem"

func FullPhysicalFirstOrderVerificationFiniteSpectralTripleCompletionAuditTheorem() theorem.Theorem {
	const id = "BRIDGE-FULL-PHYSICAL-FIRST-ORDER-VERIFICATION-FINITE-SPECTRAL-TRIPLE-COMPLETION-AUDIT"
	const name = "Full Physical First-Order Verification / Finite Spectral Triple Completion Audit"
	return theorem.Theorem{ID: id, Name: name, Layer: theorem.LayerBridge, Status: theorem.BridgeRequired, Verify: func() theorem.Result {
		a, err := BuildDefault()
		if err != nil {
			return theorem.Result{ID: id, Name: name, Layer: theorem.LayerBridge, Status: theorem.FailedRoute, Checks: []theorem.Check{{Name: "build Gate 297 audit", Passed: false, Detail: err.Error()}}}
		}
		checks := []theorem.Check{
			{Name: "Gate 295/296 inputs inherited", Passed: a.Input.Gate295TrueBimodule && a.Input.Gate296HyperchargeRay && a.Input.Gate296DiracGraph, Detail: FormatInput(a.Input)},
			{Name: "full left and opposite representations assembled structurally", Passed: a.Representation.LeftRepresentationAssembled && a.Representation.OppositeRepresentationAssembled && a.Representation.ParticleDimension == 16 && a.Representation.DoubledDimension == 32, Detail: FormatRepresentation(a.Representation)},
			{Name: "zero-order condition verified on true bimodule", Passed: a.ZeroOrder.ZeroOrderVerified && a.ZeroOrder.WeakColorCommutatorNorm < 1e-12, Detail: FormatZeroOrder(a.ZeroOrder)},
			{Name: "full structural first-order sweep verified on canonical edge graph", Passed: a.FirstOrder.FullSweepVerified && a.FirstOrder.MaxLegalResidual < 1e-12 && a.FirstOrder.MinRejectedResidual > 0.1, Detail: FormatFirstOrder(a.FirstOrder)},
			{Name: "finite spectral-triple skeleton completed but dynamics remain open", Passed: a.Triple.StructuralSkeletonComplete && !a.Triple.DynamicalTripleComplete && !a.Triple.NumericalYukawas && !a.Triple.BGapMajorana, Detail: FormatTriple(a.Triple)},
			{Name: "firewalls preserve Higgs/B-gap dynamics", Passed: !a.Firewalls.FiniteCorePolluted && a.Firewalls.DoesNotClaimDynamics && a.Firewalls.DoesNotUnlockHiggs && a.Firewalls.DoesNotUnlockBGap, Detail: FormatFirewalls(a.Firewalls)},
		}
		return theorem.Result{ID: id, Name: name, Layer: theorem.LayerBridge, Status: theorem.BridgeRequired, Checks: checks, Notes: []string{a.Truth, FormatSummary(a.Summary)}}
	}}
}
