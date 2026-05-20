package generation2hitchinchannelsignequalunitcalibrationaudit

import (
	"math"
	"strings"
	"testing"
)

func TestGate651Build(t *testing.T) {
	a, err := Build()
	if err != nil {
		t.Fatal(err)
	}
	if !a.Inherited.DegreeSelectionInherited || !a.Inherited.SectorLedgerDefined || !a.Inherited.PositiveAAAOnly || !a.Inherited.NegativeAABOnly || !a.Inherited.MixedZeroByDegree || !a.Inherited.SignCalibrationGapInherited || !a.Inherited.SlotFormulaRecovered || a.Inherited.SplitG2Certified || a.Inherited.BoundaryStressAssignment || a.Inherited.SevenOver72Theorem || a.Inherited.ScalarFlavorTransport || a.Inherited.PhysicalMetric || !a.Inherited.Gate650FirewallPreserved {
		t.Fatalf("bad inheritance: %+v", a.Inherited)
	}
	if a.Orientation.PositiveDim != 4 || a.Orientation.NegativeDim != 3 || !a.Orientation.OrientationCompatible || !a.Orientation.ConventionDependentSign {
		t.Fatalf("bad orientation audit: %+v", a.Orientation)
	}
	if !a.Maps.MapsComputed || len(a.Maps.Rows) != 4 || !contains(a.Maps.OnlySurvivors, "AAA") || !contains(a.Maps.OnlySurvivors, "AAB") || !contains(a.Maps.OnlySurvivors, "ABA") || !contains(a.Maps.OnlySurvivors, "BAA") {
		t.Fatalf("bad maps: %+v", a.Maps)
	}
	for _, r := range a.Maps.Rows {
		if !r.MatchesCalibratedRay || r.AnisotropyResidual > tol || r.OffBlockResidual > tol {
			t.Fatalf("bad channel row: %+v", r)
		}
	}
	if a.Positive.CPlus != 1 || !a.Positive.ScalarMultipleOfP || a.Positive.AnisotropyResidual > tol || !a.Positive.ContributesOnlyPlus {
		t.Fatalf("bad positive audit: %+v", a.Positive)
	}
	if !a.Negative.EqualToMinusCPlus || !a.Negative.EachScalarMultipleOfP || a.Negative.CAAB != -1 || a.Negative.CABA != -1 || a.Negative.CBAA != -1 || a.Negative.CombinedCoefficient != -3 || a.Negative.CombinedAnisotropy > tol {
		t.Fatalf("bad negative audit: %+v", a.Negative)
	}
	if !a.Sign.FiniteNegativeSignObserved || !a.Sign.FiniteEqualUnitObserved || a.Sign.BasisFreeSourceCertified || !a.Sign.RequiresCalibrationIdentityProof {
		t.Fatalf("bad sign audit: %+v", a.Sign)
	}
	if !a.Routes.AllRoutesPass || !a.Routes.SamePatternAfterNorm || a.Routes.RouteDependentMagnitude || len(a.Routes.Rows) != 3 {
		t.Fatalf("bad route audit: %+v", a.Routes)
	}
	for _, r := range a.Routes.Rows {
		if !r.EqualPattern || !r.ReconstructsProjectorRay || r.AAAUnit != 1 || r.AABUnit != -1 || r.ABAUnit != -1 || r.BAAUnit != -1 {
			t.Fatalf("bad route row: %+v", r)
		}
	}
	if !a.Reconstruction.ReconstructsPPlusMinus3P || !a.Reconstruction.RecoversGate642Angle || math.Abs(a.Reconstruction.NormSquared-31) > tol || math.Abs(a.Reconstruction.Cosine-13/math.Sqrt(217)) > tol || math.Abs(a.Reconstruction.ResidualSquared-48.0/217.0) > tol {
		t.Fatalf("bad reconstruction: %+v", a.Reconstruction)
	}
	if !a.Theorem.FiniteCalibrationIdentityPasses || a.Theorem.FullSymbolicCalibrationTheorem {
		t.Fatalf("bad theorem target: %+v", a.Theorem)
	}
	if a.Firewalls.ClaimsFullSymbolicCalibration || a.Firewalls.ClaimsSplitG2 || a.Firewalls.ClaimsBoundaryStress || a.Firewalls.ClaimsSevenOver72 || a.Firewalls.ClaimsScalarFlavor || a.Firewalls.ClaimsPhysicalMetric || a.Firewalls.ClaimsHiggsMass || a.Firewalls.ClaimsCKMPMNS || a.Firewalls.ClaimsGaugeUnification || a.Firewalls.Verdict != StatusGate651Boundary {
		t.Fatalf("firewall breach: %+v", a.Firewalls)
	}
}

func TestTheorem(t *testing.T) {
	res := Generation2HitchinChannelSignEqualUnitCalibrationAuditTheorem().Verify()
	if res.Status == "FAILED_ROUTE" {
		t.Fatalf("unexpected construction failure: %+v", res)
	}
	for _, c := range res.Checks {
		if !c.Passed {
			t.Fatalf("check failed: %s %s", c.Name, c.Detail)
		}
	}
	joined := strings.Join(res.Notes, "\n")
	for _, want := range []string{StatusGate650DegreeSelectionInherited, StatusOrientationVolumeAudited, StatusSurvivingChannelMapsComputed, StatusAAAPositiveUnitAudited, StatusAABNegativeEqualUnitAudited, StatusEqualUnitMagnitudeSupported, StatusNegativeSignSourceClassified, StatusReconstructionComputed, StatusCalibrationIdentitySharpened, StatusNoFullSymbolicCalibrationTheorem, StatusNoSplitG2, StatusNoBoundaryStress, StatusNoSevenOver72, StatusNoScalarFlavor, StatusNoPhysicalMetric, StatusNoHiggsFlavorGauge, StatusGate651Boundary} {
		if !strings.Contains(joined, want) {
			t.Fatalf("missing note %s", want)
		}
	}
}
