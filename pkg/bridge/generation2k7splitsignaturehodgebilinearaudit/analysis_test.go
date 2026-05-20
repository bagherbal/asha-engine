package generation2k7splitsignaturehodgebilinearaudit

import (
	"math"
	"strings"
	"testing"
)

func TestGate636Build(t *testing.T) {
	a, err := Build()
	if err != nil {
		t.Fatal(err)
	}
	if a.Inherited.K7Dimension != 7 || a.Inherited.PlusDimension != 4 || a.Inherited.MinusDimension != 3 || math.Abs(a.Inherited.Trace-1) > 1e-10 || math.Abs(a.Inherited.Determinant+1) > 1e-8 {
		t.Fatalf("bad inheritance values: %+v", a.Inherited)
	}
	if !a.Inherited.CarrierFirewallPreserved || !a.Inherited.NoK7ToFockSelectorMap || !a.Inherited.NoOnePlusThreeRefinement || !a.Inherited.TraceNotDistinguishedLine || !a.Inherited.NoBoundaryAssignment || !a.Inherited.NoSevenOver72Theorem {
		t.Fatalf("bad inherited firewalls: %+v", a.Inherited)
	}
	if a.Definition.Dimension != 7 || a.Definition.Rows != 7 || a.Definition.Cols != 7 || !a.Definition.Symmetric || !a.Definition.Nondegenerate || !a.Definition.InheritedFromSK {
		t.Fatalf("bad bilinear definition: %+v", a.Definition)
	}
	if a.Signature.InertiaPlus != 4 || a.Signature.InertiaMinus != 3 || a.Signature.InertiaZero != 0 || a.Signature.DeterminantSign != -1 || !a.Signature.NullConeExists || a.Signature.PositiveDefinite || a.Signature.NegativeDefinite || !a.Signature.SplitIndefinite {
		t.Fatalf("bad signature: %+v", a.Signature)
	}
	if !a.MetricConversion.SKOrthogonal || !a.MetricConversion.SKSymmetric || !a.MetricConversion.SKInvolutive || !a.MetricConversion.BEqualsGComposedWithSK {
		t.Fatalf("bad metric conversion: %+v", a.MetricConversion)
	}
	if !a.Orthogonality.GOrthogonal || !a.Orthogonality.BOrthogonal || !a.Orthogonality.CrossTermZero || a.Orthogonality.PlusDimension != 4 || a.Orthogonality.MinusDimension != 3 {
		t.Fatalf("bad plus/minus orthogonality: %+v", a.Orthogonality)
	}
	if !a.Octonionic.SplitSignatureMatchesDimension || a.Octonionic.OmegaKThreeFormCertified || a.Octonionic.CrossProductCertified || a.Octonionic.CalibrationCertified || a.Octonionic.G2SplitStructureCertified {
		t.Fatalf("bad octonionic compatibility posture: %+v", a.Octonionic)
	}
	if a.Stabilizer.BilinearStabilizerCandidate != "O(4,3)" || a.Stabilizer.OrientationPreservingCandidate != "SO(4,3)" || !a.Stabilizer.StabilizerCertified || a.Stabilizer.SplitG2Certified || !a.Stabilizer.NeedsOmegaK || a.Stabilizer.PhysicalMetricClaimed {
		t.Fatalf("bad stabilizer posture: %+v", a.Stabilizer)
	}
	if a.Firewalls.K7ToFockMapCertified || a.Firewalls.OnePlusThreeSelectorDerived || a.Firewalls.BoundaryStressAssigned || a.Firewalls.SevenOver72Promoted || a.Firewalls.PhysicalSpacetimeMetric || a.Firewalls.ScalarRGMatchingClaimed || a.Firewalls.HiggsMassClaimed || a.Firewalls.FlavorClaimed || a.Firewalls.CKMPMNSClaimed || a.Firewalls.GaugeUnificationClaimed {
		t.Fatalf("firewall breach: %+v", a.Firewalls)
	}
	if !strings.Contains(a.MissingObject.CurrentMissingObject, "Omega_K") || a.MissingObject.CanSupportSplitG2 || a.MissingObject.CanSupportBoundary {
		t.Fatalf("bad missing object: %+v", a.MissingObject)
	}
}

func TestTheorem(t *testing.T) {
	res := Generation2K7SplitSignatureHodgeBilinearAuditTheorem().Verify()
	if res.Status == "FAILED_ROUTE" {
		t.Fatalf("unexpected construction failure: %+v", res)
	}
	for _, c := range res.Checks {
		if !c.Passed {
			t.Fatalf("check failed: %s %s", c.Name, c.Detail)
		}
	}
	joined := strings.Join(res.Notes, "\n")
	for _, want := range []string{StatusGate635Inherited, StatusBKHodgeBilinearDefined, StatusBKSignatureCertified, StatusNativeSplitSignature, StatusBilinearNotSelector, StatusNoK7ToFockMap, StatusNoSplitG2Yet, StatusNoOmegaK, StatusNoBoundaryStressAssignment, StatusNoPhysicalMetric, StatusGate636Boundary} {
		if !strings.Contains(joined, want) {
			t.Fatalf("missing note %s", want)
		}
	}
}
