package generation2scalarflavordeficitclosuretriangleaudit

import (
	"math"
	"strings"

	"github.com/bagherbal/asha-engine/pkg/theorem"
)

func Generation2ScalarFlavorDeficitClosureTriangleAuditTheorem() theorem.Theorem {
	const id = AuditID
	const name = "Gate 659 — Scalar-Flavor Deficit Closure Triangle Audit"
	return theorem.Theorem{ID: id, Name: name, Layer: theorem.LayerBridge, Status: theorem.BridgeRequired, Verify: func() theorem.Result {
		a, err := BuildDefault()
		if err != nil {
			return theorem.Result{ID: id, Name: name, Layer: theorem.LayerBridge, Status: theorem.FailedRoute, Checks: []theorem.Check{{Name: "construct Gate659 scalar-flavor deficit closure audit", Passed: false, Detail: err.Error()}}}
		}
		checks := []theorem.Check{
			{Name: "inherit Gate658 scalar transport spine", Passed: a.Inherited.ScalarTransportSpineInherited && a.Inherited.ScalarBoundarySpineActive && a.Inherited.LowScaleMatchingActive && a.Inherited.BoundaryStressTransportActive && a.Inherited.KappaLambdaDefined && a.Inherited.NoNativeProxyRuntimeTheorem && a.Inherited.NoNativeRGThresholdTheorem && a.Inherited.NoNativeBoundaryStressTheorem && a.Inherited.NoHiggsMassOrStabilityClaim && a.Inherited.FirewallPreserved, Detail: FormatInherited(a.Inherited)},
			{Name: "inherit flavor kappa_e environmental seal", Passed: a.Flavor.EnvironmentalSeal && !a.Flavor.NativeFlavorTheorem && math.Abs(a.Flavor.KappaE-kappaE) < 1e-15 && strings.Contains(a.Flavor.OrientationBalanceExpression, "J_CKM"), Detail: FormatFlavor(a.Flavor)},
			{Name: "compute kappa closure triangle", Passed: math.Abs(a.Closure.KSum-0.04982659728765166) < 5e-15 && math.Abs(a.Closure.DeltaClosure-0.0001256552099683575) < 5e-15 && a.Closure.ClosesOnScalarWound && a.Closure.RawClosureResidualSmall && a.Closure.RelativeToAbsLambda < 0.003, Detail: FormatClosure(a.Closure)},
			{Name: "compute boundary split ratio against typed candidates", Passed: math.Abs(a.Boundary.BoundarySplit-0.0012924448188162962) < 5e-15 && math.Abs(a.Boundary.ObservedWeight-0.09722288188941064) < 5e-13 && a.Boundary.SevenOver72Closest && a.Boundary.UsedTypedSetOnly && len(a.Boundary.Candidates) == 4, Detail: FormatBoundaryWeight(a.Boundary)},
			{Name: "audit 7/72 boundary interpolation", Passed: math.Abs(a.Interpolation.WeightedTarget-0.04982659643506822) < 5e-15 && math.Abs(a.Interpolation.WeightedResidual-8.525834413464217e-10) < 5e-18 && a.Interpolation.ImprovementFactor > 100000 && a.Interpolation.ResidualRelativeToKSum < 2e-8 && a.Interpolation.BridgeLayerOnly, Detail: FormatInterpolation(a.Interpolation)},
			{Name: "classify source types and active lane location", Passed: len(a.Sources.Objects) == 5 && !a.Sources.SevenOver72InFanoLane && a.Sources.SevenOver72InTransportLane && !a.Sources.FanoBoundaryMapConstructed && !a.Sources.RandomConstantsSearched && a.Sources.TypedCandidatesOnly, Detail: FormatSources(a.Sources)},
			{Name: "preserve scalar-flavor-boundary firewalls", Passed: !a.Firewalls.ClaimsNativeFlavorTheorem && !a.Firewalls.ClaimsNativeScalarTheorem && !a.Firewalls.ClaimsNativeSevenOver72Theorem && !a.Firewalls.ClaimsBoundaryStressDerivation && !a.Firewalls.ClaimsHiggsPrediction && !a.Firewalls.ClaimsGaugeUnification && !a.Firewalls.ClaimsCKMPMNSDerivation && !a.Firewalls.ClaimsPhysicalSpacetime && !a.Firewalls.ClaimsNativeClosureTheorem && a.Firewalls.Verdict == StatusGate659Boundary, Detail: FormatFirewalls(a.Firewalls)},
		}
		notes := append(Statuses(), a.Truth)
		return theorem.Result{ID: id, Name: name, Layer: theorem.LayerBridge, Status: theorem.BridgeRequired, Checks: checks, Notes: notes}
	}}
}
