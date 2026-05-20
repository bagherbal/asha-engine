package generation2k7nativeomegasourcesplitg2audit

import (
	"math"
	"strings"
	"testing"
)

func TestGate637Build(t *testing.T) {
	a, err := Build()
	if err != nil {
		t.Fatal(err)
	}
	if a.Inherited.K7Dimension != 7 || a.Inherited.BKInertiaPlus != 4 || a.Inherited.BKInertiaMinus != 3 || a.Inherited.BKInertiaZero != 0 || math.Abs(a.Inherited.BKTrace-1) > 1e-10 || math.Abs(a.Inherited.BKDeterminant+1) > 1e-8 {
		t.Fatalf("bad inheritance: %+v", a.Inherited)
	}
	if !a.Inherited.NativeSplitSignature || !a.Inherited.BilinearNotSelector || !a.Inherited.NoFockSelectorMap || !a.Inherited.NoSplitG2Yet || !a.Inherited.NoBoundaryAssignment || !a.Inherited.NoSevenOver72Theorem || !a.Inherited.Gate636FirewallPreserved {
		t.Fatalf("bad inherited firewall posture: %+v", a.Inherited)
	}
	if a.Source.PGSectorDimension != 14 || a.Source.RawCalibrationRows != 70 || a.Source.RawCalibrationColumns != 14 || a.Source.AssociativeFanoTerms != 7 || a.Source.CoassociativeTerms != 7 || !a.Source.K7ToPGCoordinatesComputed || a.Source.UsesArbitrarySplitG2Normal || a.Source.UsesExternalThreeForm || a.Source.HodgePolarityAloneSufficient {
		t.Fatalf("bad source audit: %+v", a.Source)
	}
	if len(a.Candidates.Candidates) != 4 || !a.Candidates.PullbackCandidatesComputed || a.Candidates.NonZeroStableCandidates < 3 || !a.Candidates.CandidateStabilityCertified || a.Candidates.CompatibleNativeOmegaCertified || a.Candidates.CompatibleWithBKCount != 0 {
		t.Fatalf("bad candidate summary: %+v", a.Candidates)
	}
	for _, c := range a.Candidates.Candidates {
		if !c.FullyAntisymmetric || c.AntisymmetryResidual > 1e-9 || !c.HitchinMetricComputed {
			t.Fatalf("bad candidate tensor certificate: %+v", c)
		}
		if c.NonZero && (!c.Stable || c.HitchinMetricInertiaPlus != 7 || c.HitchinMetricInertiaMinus != 0 || c.HitchinMetricInertiaZero != 0) {
			t.Fatalf("nonzero candidate must be stable compact-positive, got %+v", c)
		}
		if c.CompatibleWithBK {
			t.Fatalf("no candidate should be B_K compatible: %+v", c)
		}
	}
	if a.Compatibility.GomegaProportionalToBK || a.Compatibility.GomegaSignatureMatchesBK || a.Compatibility.CertifiedScaleNotFitted || a.Compatibility.BestOmegaInertia != "(7,0,0)" || a.Compatibility.BestRelativeResidualToBK < 0.9 {
		t.Fatalf("bad compatibility obstruction: %+v", a.Compatibility)
	}
	if a.CrossProduct.OmegaCompatibleWithBK || a.CrossProduct.CrossProductDefined || a.CrossProduct.BKPairingIdentityCertified || a.CrossProduct.SplitCrossProductIdentity {
		t.Fatalf("cross-product firewall breach: %+v", a.CrossProduct)
	}
	if a.Stabilizer.SplitG2Certified || a.Stabilizer.StabilizerDimensionComputed || a.Stabilizer.ExpectedSplitG2Dimension != 14 {
		t.Fatalf("split-G2 firewall breach: %+v", a.Stabilizer)
	}
	if !a.NativeStatus.NativePullbackTensorExists || a.NativeStatus.CompatibleOmegaKCertified || a.NativeStatus.SplitG2CandidateCertified || a.NativeStatus.BoundaryStressAssignment || a.NativeStatus.SevenOver72TraceTheorem || a.NativeStatus.PhysicalSpacetimeMetricClaimed || a.NativeStatus.FockSelectorClaimed || a.NativeStatus.ScalarRGMatchingClaimed || a.NativeStatus.FlavorClaimed || a.NativeStatus.GaugeUnificationClaimed {
		t.Fatalf("bad native status: %+v", a.NativeStatus)
	}
	if a.Firewalls.ClaimsPhysicalSpacetimeMetric || a.Firewalls.ClaimsFockSelector || a.Firewalls.ClaimsBoundaryStress || a.Firewalls.ClaimsSevenOver72Theorem || a.Firewalls.ClaimsScalarRGMatching || a.Firewalls.ClaimsHiggsMass || a.Firewalls.ClaimsFlavor || a.Firewalls.ClaimsCKMPMNS || a.Firewalls.ClaimsGaugeUnification || a.Firewalls.ClaimsSplitG2WithoutOmega {
		t.Fatalf("firewall breach: %+v", a.Firewalls)
	}
}

func TestTheorem(t *testing.T) {
	res := Generation2K7NativeOmegaSourceSplitG2CompatibilityAuditTheorem().Verify()
	if res.Status == "FAILED_ROUTE" {
		t.Fatalf("unexpected construction failure: %+v", res)
	}
	for _, c := range res.Checks {
		if !c.Passed {
			t.Fatalf("check failed: %s %s", c.Name, c.Detail)
		}
	}
	joined := strings.Join(res.Notes, "\n")
	for _, want := range []string{StatusGate636Inherited, StatusOctonionicCalibrationSource, StatusPGPullbackCandidatesComputed, StatusOmegaCandidateStable, StatusOmegaCompactNotSplitBK, StatusNoCompatibleOmegaK, StatusSplitSignatureAloneNoSplitG2, StatusNoCertifiedSplitG2, StatusNoBoundaryStressAssignment, StatusNoSevenOver72Theorem, StatusGate637Boundary} {
		if !strings.Contains(joined, want) {
			t.Fatalf("missing note %s", want)
		}
	}
}
