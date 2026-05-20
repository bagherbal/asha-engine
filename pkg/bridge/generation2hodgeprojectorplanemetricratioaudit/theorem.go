package generation2hodgeprojectorplanemetricratioaudit

import (
	"strings"

	"github.com/bagherbal/asha-engine/pkg/theorem"
)

func Generation2HodgeProjectorPlaneMetricRatioAuditTheorem() theorem.Theorem {
	const id = AuditID
	const name = "Gate 644 — HodgeProjector Plane MetricRatio Audit"
	return theorem.Theorem{ID: id, Name: name, Layer: theorem.LayerBridge, Status: theorem.BridgeRequired, Verify: func() theorem.Result {
		a, err := BuildDefault()
		if err != nil {
			return theorem.Result{ID: id, Name: name, Layer: theorem.LayerBridge, Status: theorem.FailedRoute, Checks: []theorem.Check{{Name: "construct Gate644 projector-plane metric-ratio audit", Passed: false, Detail: err.Error()}}}
		}
		checks := []theorem.Check{
			{Name: "inherit Gate643 same-sector residual tensor and firewall", Passed: a.Inherited.Verdict == StatusGate643ResidualInherited && a.Inherited.ResidualTensorCertified && a.Inherited.SameSectorHodgeDiagonal && a.Inherited.OffSectorCarrierRejected && !a.Inherited.NativeTraceIdentityFound && !a.Inherited.SplitG2Certified && !a.Inherited.BoundaryStressAssignment && !a.Inherited.SevenOver72Theorem && !a.Inherited.ScalarFlavorTransport && !a.Inherited.PhysicalAngle && !a.Inherited.PhysicalMetric && a.Inherited.Gate643FirewallPreserved, Detail: FormatInherited(a.Inherited)},
			{Name: "define B_hat and G_hat projector-plane targets", Passed: a.Definitions.ProjectorPlaneMetricsCertified && a.Definitions.BHatTargetResidual < strictTolerance && a.Definitions.ProjectorPlaneTargetResidual < strictTolerance, Detail: FormatDefinitions(a.Definitions)},
			{Name: "reconstruct G_hat as the 1:-3 Hodge-sector metric ratio across routes", Passed: a.MetricRatio.AllRoutesRatioCertified && len(a.MetricRatio.Routes) >= 3 && a.MetricRatio.MaxProjectorPlaneResidual < ratioTolerance && a.MetricRatio.MaxReconstructedResidual < ratioTolerance && a.MetricRatio.MaxOffDiagonalNorm < ratioTolerance && a.MetricRatio.MaxRatioDrift < ratioTolerance, Detail: FormatMetricRatioDetails(a.MetricRatio)},
			{Name: "derive the projective angle from the two projector-plane rays", Passed: a.AngleFromPlane.AngleDerivedFromPlane && !a.AngleFromPlane.NativeTraceIdentityFound, Detail: FormatAngleFromPlane(a.AngleFromPlane)},
			{Name: "classify minus-three weight as candidate, not theorem", Passed: !a.MinusThree.CertifiedNativeSource && strings.Contains(a.MinusThree.Verdict, StatusNoMinusThreeSource), Detail: FormatMinusThree(a.MinusThree)},
			{Name: "preserve split-G2, boundary, scalar/flavor, physical-geometry, and 7/72 firewalls", Passed: !a.Firewalls.ClaimsNativeTraceIdentity && !a.Firewalls.ClaimsSplitG2 && !a.Firewalls.ClaimsBoundaryStress && !a.Firewalls.ClaimsSevenOver72Theorem && !a.Firewalls.ClaimsScalarFlavor && !a.Firewalls.ClaimsPhysicalAngle && !a.Firewalls.ClaimsPhysicalMetric && !a.Firewalls.ClaimsFlavor && !a.Firewalls.ClaimsHiggsMass && !a.Firewalls.ClaimsCKMPMNS && !a.Firewalls.ClaimsGaugeUnification && a.Firewalls.Verdict == StatusGate644Boundary, Detail: FormatFirewalls(a.Firewalls)},
		}
		notes := []string{
			StatusGate643ResidualInherited,
			StatusGHatReconstructed,
			StatusProjectorPlaneMetricsDefined,
			StatusRouteMetricRatiosComputed,
			StatusHodgeDiagonalRatio,
			StatusProjectorPlaneAngle,
			StatusMinusThreeSourceCandidate,
			StatusNoMinusThreeSource,
			StatusNoNativeTraceIdentity,
			StatusNoCertifiedSplitG2,
			StatusNoBoundaryStress,
			StatusNoSevenOver72Theorem,
			StatusNoScalarFlavorTransport,
			StatusNoPhysicalAngle,
			StatusNoPhysicalMetric,
			StatusNoHiggsFlavorGauge,
			StatusGate644Boundary,
			"Gate644 reconstructs the normalized twist metric behind Gate643: B_hat=(P_{K7+}-P_{K7-})/sqrt(7), while every audited split-twist route gives G_hat=(P_{K7+}-3P_{K7-})/sqrt(31) after projective sign alignment. Therefore the Gate642 angle follows from <G_hat,B_hat>=(4+9)/sqrt(31*7)=13/sqrt(217), and the residual square 48/217 follows automatically. The -3 weight is a live source-pressure point, not a theorem.",
		}
		return theorem.Result{ID: id, Name: name, Layer: theorem.LayerBridge, Status: theorem.BridgeRequired, Checks: checks, Notes: notes}
	}}
}
