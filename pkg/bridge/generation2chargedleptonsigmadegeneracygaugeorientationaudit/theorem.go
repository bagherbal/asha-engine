package generation2chargedleptonsigmadegeneracygaugeorientationaudit

import "github.com/bagherbal/asha-engine/pkg/theorem"

func Generation2ChargedLeptonSigmaDegeneracyGaugeOrientationAuditTheorem() theorem.Theorem {
	const id = AuditID
	const name = "Generation 2 charged-lepton sigma degeneracy gauge-or-orientation audit"
	return theorem.Theorem{ID: id, Name: name, Layer: theorem.LayerBridge, Status: theorem.BridgeRequired, Verify: func() theorem.Result {
		a, err := BuildDefault()
		if err != nil {
			return theorem.Result{ID: id, Name: name, Layer: theorem.LayerBridge, Status: theorem.FailedRoute, Checks: []theorem.Check{{Name: "construct Gate603 sigma degeneracy gauge-orientation audit", Passed: false, Detail: err.Error()}}}
		}
		checks := []theorem.Check{
			{Name: "inherit Gate602 selector result", Passed: a.Inherited.SelectsElectronRow && a.Inherited.SelectsP3 && a.Inherited.SelectsPositiveJ && !a.Inherited.SelectsFullSigma && a.Inherited.MinimalClassSize == 6, Detail: FormatInherited(a.Inherited)},
			{Name: "audit S3 action on Fourier coordinates", Passed: len(a.S3Action) == 6 && allSameElectronEpsilon(a.S3Action) && hasBothVSigns(a.S3Action), Detail: FormatS3Table(a.S3Action, 6)},
			{Name: "classify invariant versus orientation-sensitive quantities", Passed: len(a.Invariants) >= 6 && containsQuantity(a.Invariants, "signed Vandermonde sign(V_x)"), Detail: FormatInvariants(a.Invariants)},
			{Name: "audit physical charged-lepton labels", Passed: a.PhysicalLabels.EnvironmentalOrdering && !a.PhysicalLabels.NativeOrdering, Detail: FormatPhysicalLabels(a.PhysicalLabels)},
			{Name: "audit signed discriminant obstruction", Passed: a.SignedDiscriminant.TraceRingSuppliesDeltaOnly && a.SignedDiscriminant.SignedVRequiresOrdering && a.SignedDiscriminant.SignDistinguishesOrientation && !a.SignedDiscriminant.NativeSignedVTheorem, Detail: FormatSignedDiscriminant(a.SignedDiscriminant)},
			{Name: "audit Fourier cyclic orientation", Passed: a.FourierCyclic.RequiresCyclicConvention && a.FourierCyclic.BFlavUsesUnsignedWallDistance && !a.FourierCyclic.BFlavDependsOnCyclicOrientation, Detail: FormatFourierCyclic(a.FourierCyclic)},
			{Name: "audit PMNS/CKM orientation coupling candidates", Passed: !a.OrientationCoupling.TypedASHAOperatorPresent && a.OrientationCoupling.JCKMSign == +1 && a.OrientationCoupling.JPMNSSign == -1, Detail: FormatOrientationCoupling(a.OrientationCoupling)},
			{Name: "define minimal remaining seal or gauge statement", Passed: a.MinimalRemaining.SigmaGaugeForBFlav && a.MinimalRemaining.PhysicalFullOrderingRequiresSeal && a.MinimalRemaining.SealName == "ChargedLeptonDiscriminantOrientationSeal", Detail: FormatMinimalSeal(a.MinimalRemaining)},
			{Name: "preserve sigma/orientation firewalls", Passed: !a.Firewalls.DerivesKoide && !a.Firewalls.DerivesChargedLeptonMasses && !a.Firewalls.DerivesPMNS && !a.Firewalls.DerivesCKM && !a.Firewalls.DerivesBFlavZero && !a.Firewalls.AddsCarrier && !a.Firewalls.AddsSelector && a.Firewalls.PreservesGate352 && a.Firewalls.PreservesGate596 && a.Firewalls.PreservesGate600 && a.Firewalls.PreservesGate602, Detail: FormatFirewalls(a.Firewalls)},
		}
		notes := append(Statuses(), a.Truth, a.MinimalRemaining.Statement)
		return theorem.Result{ID: id, Name: name, Layer: theorem.LayerBridge, Status: theorem.BridgeRequired, Checks: checks, Notes: notes}
	}}
}

func allSameElectronEpsilon(rows []S3ActionRow) bool {
	if len(rows) == 0 {
		return false
	}
	v := rows[0].ElectronWallEpsilonRad
	for _, r := range rows {
		if abs(r.ElectronWallEpsilonRad-v) > 1e-15 {
			return false
		}
	}
	return true
}
func hasBothVSigns(rows []S3ActionRow) bool {
	pos, neg := false, false
	for _, r := range rows {
		if r.SignVandermondeX > 0 {
			pos = true
		}
		if r.SignVandermondeX < 0 {
			neg = true
		}
	}
	return pos && neg
}
func containsQuantity(rows []InvariantVsOrientationSensitive, q string) bool {
	for _, r := range rows {
		if r.Quantity == q {
			return true
		}
	}
	return false
}
func abs(x float64) float64 {
	if x < 0 {
		return -x
	}
	return x
}
