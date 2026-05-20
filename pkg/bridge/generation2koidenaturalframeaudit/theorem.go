package generation2koidenaturalframeaudit

import "github.com/bagherbal/asha-engine/pkg/theorem"

func Generation2KoideNaturalFrameAuditTheorem() theorem.Theorem {
	const id = AuditID
	const name = "Generation 2 Koide natural frame audit"
	return theorem.Theorem{ID: id, Name: name, Layer: theorem.LayerBridge, Status: theorem.BridgeRequired, Verify: func() theorem.Result {
		a, err := BuildDefault()
		if err != nil {
			return theorem.Result{ID: id, Name: name, Layer: theorem.LayerBridge, Status: theorem.FailedRoute, Checks: []theorem.Check{{Name: "construct Gate579 Koide natural frame audit", Passed: false, Detail: err.Error()}}}
		}
		checks := []theorem.Check{
			{Name: "inherit Gate578 and history transport runtime", Passed: a.Runtime.Mu0GeV > 0 && a.Runtime.Lambda12GeV > 0 && a.Runtime.Gate578PhiMZDeg > 257, Detail: FormatRuntime(a.Runtime)},
			{Name: "certify Koide democratic frame", Passed: abs(a.Frame.DotNE1) < 1e-12 && abs(a.Frame.DotNE2) < 1e-12 && abs(a.Frame.DotE1E2) < 1e-12 && abs(a.Frame.NormN-1) < 1e-12 && abs(a.Frame.NormE1-1) < 1e-12 && abs(a.Frame.NormE2-1) < 1e-12 && a.Frame.RightHanded, Detail: FormatFrame(a.Frame)},
			{Name: "compute pole-mass Koide frame", Passed: a.Compare.Pole.Rho > 1 && a.Compare.Pole.PhiDeg > 257.26 && a.Compare.Pole.PhiDeg < 257.28 && abs(a.Compare.Pole.Q-KoideTarget) < 1e-5, Detail: FormatPoint(a.Compare.Pole)},
			{Name: "compute M_Z Yukawa Koide frame", Passed: a.Compare.MZ.Rho > 0.1 && a.Compare.MZ.PhiDeg > 257.26 && a.Compare.MZ.PhiDeg < 257.28 && abs(a.Compare.MZ.Q-KoideTarget) < 1e-5, Detail: FormatPoint(a.Compare.MZ)},
			{Name: "compute Lambda_12 Yukawa Koide frame", Passed: a.Compare.Lambda12.PhiDeg > 257.26 && a.Compare.Lambda12.PhiDeg < 257.28 && abs(a.Compare.Lambda12.Q-KoideTarget) < abs(a.Compare.MZ.Q-KoideTarget), Detail: FormatPoint(a.Compare.Lambda12)},
			{Name: "prove pole/M_Z angle degeneracy in v1", Passed: a.Compare.MZEqualPole && abs(a.Compare.DeltaPhiPoleToMZDeg) < 1e-12 && abs(a.Compare.DeltaQPoleToMZ) < 1e-14, Detail: FormatComparison(a.Compare)},
			{Name: "record Lambda_12 cleaner but not certified", Passed: a.Compare.LambdaCloserThanMZ && a.Compare.AzimuthStable && !a.Natural.BoundaryFrameCertified, Detail: FormatNatural(a.Natural)},
			{Name: "preserve natural-frame firewalls", Passed: !a.Firewalls.DerivesLeptonMasses && !a.Firewalls.DerivesYukawaEigenvalues && !a.Firewalls.DerivesPMNS && !a.Firewalls.DerivesCKM && !a.Firewalls.DerivesGenerationHierarchy && !a.Firewalls.PromotesObservedAsNative && !a.Firewalls.AddsNewCarrier && a.Firewalls.PreservesGate352, Detail: FormatFirewalls(a.Firewalls)},
			{Name: "return Gate579 final verdict", Passed: a.Final.SealName == "ChargedLeptonKoideNaturalFrameSeal" && !a.Final.NaturalFrameCertified && a.Final.MinimalNextRequirement != "", Detail: FormatFinal(a.Final)},
		}
		notes := append(Statuses(), a.Truth, a.Final.MinimalNextRequirement)
		return theorem.Result{ID: id, Name: name, Layer: theorem.LayerBridge, Status: theorem.BridgeRequired, Checks: checks, Notes: notes}
	}}
}

func abs(x float64) float64 {
	if x < 0 {
		return -x
	}
	return x
}
