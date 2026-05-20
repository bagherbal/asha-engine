package generation2koidecoordinatebetafunctionaudit

import "github.com/bagherbal/asha-engine/pkg/theorem"

func Generation2KoideCoordinateBetaFunctionAuditTheorem() theorem.Theorem {
	const id = AuditID
	const name = "Generation 2 Koide coordinate beta-function audit"
	return theorem.Theorem{ID: id, Name: name, Layer: theorem.LayerBridge, Status: theorem.BridgeRequired, Verify: func() theorem.Result {
		a, err := BuildDefault()
		if err != nil {
			return theorem.Result{ID: id, Name: name, Layer: theorem.LayerBridge, Status: theorem.FailedRoute, Checks: []theorem.Check{{Name: "construct Gate581 Koide coordinate beta audit", Passed: false, Detail: err.Error()}}}
		}
		checks := []theorem.Check{
			{Name: "inherit Gate580 and history runtime", Passed: a.Runtime.Mu0GeV > 0 && a.Runtime.Lambda12GeV > a.Runtime.Mu0GeV && a.Runtime.LogLambdaOverMZ > 20, Detail: FormatRuntime(a.Runtime)},
			{Name: "derive coordinate beta formulas", Passed: a.Formula.CommonRateCancels && a.Formula.DLnRhoFormula != "" && a.Formula.DThetaFormula != "" && a.Formula.DPhiFormula != "", Detail: FormatFormula(a.Formula)},
			{Name: "compute M_Z local Koide coordinate beta", Passed: a.MZ.DThetaDTDeg > 0 && a.MZ.DPhiDTDeg > 0 && a.MZ.ProjectiveSpeedRad > 0 && a.MZ.PointsTowardCone && a.MZ.PhiSlow, Detail: FormatEndpointBeta(a.MZ)},
			{Name: "compute Lambda_12 local Koide coordinate beta", Passed: a.Lambda12.DThetaDTDeg > 0 && a.Lambda12.DPhiDTDeg > 0 && a.Lambda12.ProjectiveSpeedRad > 0 && a.Lambda12.PointsTowardCone && a.Lambda12.PhiSlow, Detail: FormatEndpointBeta(a.Lambda12)},
			{Name: "show common multiplicative running gives no projective motion", Passed: a.Source.ProjectiveMotionRequiresRateSplitting && a.MZ.CommonOnlyProjectiveSpeedRad < 1e-18 && a.Lambda12.CommonOnlyProjectiveSpeedRad < 1e-18, Detail: FormatProjectiveSource(a.Source)},
			{Name: "block Koide cone invariant/attractor theorem in v1", Passed: !a.Cone.ConeInvariantInV1 && !a.Cone.AttractorCertified && a.Cone.MZExactConeDThetaDTDeg > 0 && a.Cone.LambdaExactConeDThetaDTDeg > 0, Detail: FormatCone(a.Cone)},
			{Name: "preserve flavor/root-trace firewalls", Passed: !a.Firewalls.DerivesKoide && !a.Firewalls.DerivesLeptonMasses && !a.Firewalls.DerivesYukawaEigenvalues && !a.Firewalls.DerivesCKM && !a.Firewalls.DerivesPMNS && !a.Firewalls.DerivesGenerationHierarchy && !a.Firewalls.IntroducesNewCarrier && !a.Firewalls.PromotesObservedAsNative && a.Firewalls.PreservesGate352, Detail: FormatFirewalls(a.Firewalls)},
			{Name: "return Gate581 final verdict", Passed: a.Final.SealName == "ChargedLeptonKoideCoordinateBetaSeal" && !a.Final.ConeInvariantInV1 && !a.Final.AttractorCertified && a.Final.MinimalNextRequirement != "", Detail: FormatFinal(a.Final)},
		}
		notes := append(Statuses(), a.Truth, a.Final.MinimalNextRequirement)
		return theorem.Result{ID: id, Name: name, Layer: theorem.LayerBridge, Status: theorem.BridgeRequired, Checks: checks, Notes: notes}
	}}
}
