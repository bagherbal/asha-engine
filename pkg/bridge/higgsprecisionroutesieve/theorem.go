package higgsprecisionroutesieve

import "github.com/bagherbal/asha-engine/pkg/theorem"

func HiggsPrecisionRepairRouteSievePoleCorrectionVsContactShapeAuditTheorem() theorem.Theorem {
	const id = "BRIDGE-HIGGS-PRECISION-REPAIR-ROUTE-SIEVE-POLE-CORRECTION-VS-CONTACT-SHAPE-AUDIT"
	const name = "Higgs Precision Repair Route Sieve / Pole Correction vs Contact Shape Audit"
	return theorem.Theorem{ID: id, Name: name, Layer: theorem.LayerBridge, Status: theorem.BridgeRequired, Verify: func() theorem.Result {
		a, err := BuildDefault()
		if err != nil {
			return theorem.Result{ID: id, Name: name, Layer: theorem.LayerBridge, Status: theorem.FailedRoute, Checks: []theorem.Check{{Name: "build Gate 337 precision route sieve", Passed: false, Detail: err.Error()}}}
		}
		checks := []theorem.Check{
			{Name: "Gate 336 exact inverse gap inherited", Passed: a.Inputs.HighestInheritedGate == inheritedHighestGate && a.Inputs.ContactShape.Cmp(rat(1197, 4624)) == 0 && a.Inputs.RequiredRePiGeV2.RatString() == "504067437/11560000", Detail: FormatInputs(a.Inputs)},
			{Name: "repair routes audited and contact-shape fit rejected", Passed: len(a.Routes.Routes) == 3 && !a.Routes.Routes[0].AllowedByNativeCore && a.Routes.Routes[2].AllowedByNativeCore, Detail: FormatRoutes(a.Routes)},
			{Name: "high-precision one-loop component kernel recomputed", Passed: len(a.Kernel.Components) == 4 && nearlyFloat(a.Kernel.RawKernelGeV2, -991.5670298916105, 1e-9), Detail: FormatKernel(a.Kernel)},
			{Name: "finite counterterm target solved", Passed: nearlyFloat(a.Counterterm.FiniteRemainderGeV2, 1035.1714794590845, 1e-9) && nearlyFloat(a.Counterterm.RemainderOverRequired, 23.7400423518074, 1e-12), Detail: FormatCounterterm(a.Counterterm)},
			{Name: "pole-correction branch preferred without collider claim", Passed: a.Recommendation.BestRoute != "" && a.Firewalls.NoFullRenormalizedSM && a.Firewalls.NoColliderClaim, Detail: FormatRecommendation(a.Recommendation) + " | " + FormatFirewalls(a.Firewalls)},
		}
		return theorem.Result{ID: id, Name: name, Layer: theorem.LayerBridge, Status: theorem.BridgeRequired, Checks: checks, Notes: []string{a.Truth, FormatSummary(a.Summary), FormatStatuses(Statuses(a))}}
	}}
}
