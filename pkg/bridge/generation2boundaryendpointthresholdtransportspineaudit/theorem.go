package generation2boundaryendpointthresholdtransportspineaudit

import "github.com/bagherbal/asha-engine/pkg/theorem"

func Generation2BoundaryEndpointThresholdTransportSpineAuditTheorem() theorem.Theorem {
	const id = AuditID
	const name = "Generation 2 boundary-to-endpoint RG threshold transport spine audit"
	return theorem.Theorem{ID: id, Name: name, Layer: theorem.LayerBridge, Status: theorem.BridgeRequired, Verify: func() theorem.Result {
		a, err := BuildDefault()
		if err != nil {
			return theorem.Result{ID: id, Name: name, Layer: theorem.LayerBridge, Status: theorem.FailedRoute, Checks: []theorem.Check{{Name: "construct Gate606 RG threshold transport spine audit", Passed: false, Detail: err.Error()}}}
		}
		checks := []theorem.Check{
			{Name: "inherit Gate605 master history vector", Passed: a.Inherited.MasterVectorBuilt && a.Inherited.RGTopActionable, Detail: FormatInherited(a.Inherited)},
			{Name: "classify native boundary conditions", Passed: len(a.NativeBoundaryTable) >= 8 && containsBoundary(a.NativeBoundaryTable, "k_Y=5/3") && containsBoundary(a.NativeBoundaryTable, "sin²(theta_*)=3/8") && containsBoundary(a.NativeBoundaryTable, "symbolic EW Hessian"), Detail: FormatNativeBoundaryTable(a.NativeBoundaryTable)},
			{Name: "build endpoint observed ledger", Passed: len(a.EndpointLedger) >= 10 && containsEndpoint(a.EndpointLedger, "g1(M_Z)") && containsEndpoint(a.EndpointLedger, "lambda(M_Z)") && containsEndpoint(a.EndpointLedger, "v"), Detail: FormatEndpointLedger(a.EndpointLedger)},
			{Name: "define gauge RG transport slots", Passed: len(a.GaugeTransport) >= 8 && containsGauge(a.GaugeTransport, "Lambda_12") && containsGauge(a.GaugeTransport, "Delta_3") && containsGauge(a.GaugeTransport, "Delta_sin²"), Detail: FormatGaugeTransport(a.GaugeTransport)},
			{Name: "define scalar RG transport slots", Passed: len(a.ScalarTransport) >= 6 && containsScalar(a.ScalarTransport, "lambda(Lambda_12)") && containsScalar(a.ScalarTransport, "zero_crossing_scale"), Detail: FormatScalarTransport(a.ScalarTransport)},
			{Name: "define threshold correction ledger", Passed: len(a.ThresholdSlots) >= 7 && containsThreshold(a.ThresholdSlots, "delta_i^gauge") && containsThreshold(a.ThresholdSlots, "delta_K_phi") && containsThreshold(a.ThresholdSlots, "delta_v"), Detail: FormatThresholdSlots(a.ThresholdSlots)},
			{Name: "classify kinetic normalization blockers", Passed: len(a.KineticBlockers) >= 6 && containsBlocker(a.KineticBlockers, "K_phi") && containsBlocker(a.KineticBlockers, "v") && containsBlocker(a.KineticBlockers, "continuum matching"), Detail: FormatKineticBlockers(a.KineticBlockers)},
			{Name: "record flavor seals as environmental inputs only", Passed: len(a.FlavorRelation) >= 5 && a.FlavorRelation[0].Item == "MinimalFlavorHistoryBranchSeal", Detail: FormatFlavorRelation(a.FlavorRelation)},
			{Name: "preserve RG/product-time firewall", Passed: !a.ProductTimeFirewall.RGScaleIsProductTime && !a.ProductTimeFirewall.RGScaleIsOSHilbert && !a.ProductTimeFirewall.RGScaleIsCosmoTime, Detail: FormatProductTimeFirewall(a.ProductTimeFirewall)},
			{Name: "write updated history transport formula", Passed: a.Formula.Formula != "" && containsString(a.Formula.ThresholdSlots, "delta_i^gauge") && containsString(a.Formula.BlockedPromotions, "absolute kinetic scale"), Detail: FormatFormula(a.Formula)},
			{Name: "preserve endpoint firewalls", Passed: !a.Firewalls.ClaimsFullUnification && !a.Firewalls.DerivesEndpoint && !a.Firewalls.DerivesKineticScale && !a.Firewalls.DerivesVEV && !a.Firewalls.DerivesFlavor && !a.Firewalls.DerivesProductTime && a.Firewalls.ThresholdsExplicit, Detail: FormatFirewalls(a.Firewalls)},
		}
		notes := append(Statuses(), a.Truth, a.Formula.Formula)
		return theorem.Result{ID: id, Name: name, Layer: theorem.LayerBridge, Status: theorem.BridgeRequired, Checks: checks, Notes: notes}
	}}
}
