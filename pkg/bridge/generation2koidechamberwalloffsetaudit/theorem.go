package generation2koidechamberwalloffsetaudit

import "github.com/bagherbal/asha-engine/pkg/theorem"

func Generation2KoideChamberWallOffsetAuditTheorem() theorem.Theorem {
	const id = AuditID
	const name = "Generation 2 Koide chamber-wall offset audit"
	return theorem.Theorem{ID: id, Name: name, Layer: theorem.LayerBridge, Status: theorem.BridgeRequired, Verify: func() theorem.Result {
		a, err := BuildDefault()
		if err != nil {
			return theorem.Result{ID: id, Name: name, Layer: theorem.LayerBridge, Status: theorem.FailedRoute, Checks: []theorem.Check{{Name: "construct Gate583 Koide chamber-wall audit", Passed: false, Detail: err.Error()}}}
		}
		checks := []theorem.Check{
			{Name: "inherit Gate582 Fourier/circulant runtime", Passed: a.Runtime.Mu0GeV > 0 && a.Runtime.Lambda12GeV > a.Runtime.Mu0GeV && a.Runtime.Gate582DeltaMZDeg > 132 && a.Runtime.Gate582DeltaL12Deg > 132, Detail: FormatRuntime(a.Runtime)},
			{Name: "define positive S3 Koide chamber and electron-zero wall", Passed: a.Chamber.PositiveChamberDeg[0] == 105 && a.Chamber.PositiveChamberDeg[1] == 135 && abs(a.Chamber.UpperWallZeroCheck) < 1e-14 && abs(a.Chamber.LowerWallZeroCheck) < 1e-14, Detail: FormatChamber(a.Chamber)},
			{Name: "compute M_Z electron-wall offset", Passed: a.MZ.InsideCanonicalChamber && a.MZ.EpsilonDeg > 2.26 && a.MZ.EpsilonDeg < 2.28 && a.MZ.ElectronRootOverA < 0.041 && abs(a.MZ.QuadraticResidual) < 2e-6, Detail: FormatWallPoint(a.MZ)},
			{Name: "compute Lambda_12 electron-wall offset", Passed: a.Lambda12.InsideCanonicalChamber && a.Lambda12.EpsilonDeg > 2.26 && a.Lambda12.EpsilonDeg < 2.28 && a.Lambda12.ElectronRootOverA < 0.041, Detail: FormatWallPoint(a.Lambda12)},
			{Name: "certify v1 wall-offset stability and chamber preservation", Passed: a.Transport.EpsilonStable && a.Transport.AbsDriftDeg < 3e-4 && a.Transport.ChamberPreserved && a.Transport.AmplitudeMovesToward1, Detail: FormatTransport(a.Transport)},
			{Name: "audit quark formal coordinates without Koide wall promotion", Passed: !a.Quarks.Up.WallSealValid && !a.Quarks.Down.WallSealValid && a.Quarks.Up.R > 1.2 && a.Quarks.Down.R > 1.1, Detail: FormatQuarks(a.Quarks)},
			{Name: "preserve root-trace and flavor firewalls", Passed: !a.Firewalls.DerivesLeptonMasses && !a.Firewalls.DerivesYukawaEigenvalues && !a.Firewalls.DerivesKoide && !a.Firewalls.DerivesEpsilon && !a.Firewalls.DerivesCKM && !a.Firewalls.DerivesPMNS && !a.Firewalls.DerivesGenerationHierarchy && !a.Firewalls.AddsNewCarrier && !a.Firewalls.PromotesObservedAsNative && a.Firewalls.PreservesGate352, Detail: FormatFirewalls(a.Firewalls)},
			{Name: "return Gate583 final wall-offset seal verdict", Passed: a.Final.SealName == "ChargedLeptonKoideChamberWallOffsetSeal" && a.Final.WallOffsetStableInV1 && a.Final.HierarchyNearWall && !a.Final.QuarkWallSealCertified && !a.Final.NativeSelectorCertified && a.Final.MinimalNextRequirement != "", Detail: FormatFinal(a.Final)},
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
