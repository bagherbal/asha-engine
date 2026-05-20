package generation2negativesectormultiplicityhitchinmetricsourceaudit

import (
	"math"
	"strings"
	"testing"
)

func TestGate645Build(t *testing.T) {
	a, err := Build()
	if err != nil {
		t.Fatal(err)
	}
	if !a.Inherited.ProjectorPlaneRatioCertified || !a.Inherited.MinusThreeCandidate || a.Inherited.MinusThreeSourceFound || a.Inherited.NativeTraceIdentityFound || a.Inherited.SplitG2Certified || a.Inherited.BoundaryStressAssignment || a.Inherited.SevenOver72Theorem || a.Inherited.ScalarFlavorTransport || a.Inherited.PhysicalMetric || !a.Inherited.Gate644FirewallPreserved {
		t.Fatalf("bad inheritance: %+v", a.Inherited)
	}
	if !a.Components.AllFamiliesAudited || !a.Components.AntisymmetrizedTwistUsed || len(a.Components.Families) != 4 || a.Components.Omega0TotalNormSq <= 0 || a.Components.Omega1AltTotalNormSq <= 0 {
		t.Fatalf("bad component decomposition: %+v", a.Components)
	}
	seen := map[string]bool{}
	for _, f := range a.Components.Families {
		seen[f.Family] = true
	}
	for _, want := range []string{"Ω+++", "Ω++-", "Ω+--", "Ω---"} {
		if !seen[want] {
			t.Fatalf("missing family %s", want)
		}
	}
	if !a.HitchinBlocks.AllRoutesBlockCertified || len(a.HitchinBlocks.Routes) != 3 || a.HitchinBlocks.MaxPlusSpread > blockTolerance || a.HitchinBlocks.MaxMinusSpread > blockTolerance || a.HitchinBlocks.MaxOffDiagonalNorm > blockTolerance || a.HitchinBlocks.MaxRatioDrift > blockTolerance {
		t.Fatalf("bad Hitchin blocks: %+v", a.HitchinBlocks)
	}
	for _, r := range a.HitchinBlocks.Routes {
		if !r.BlockFormCertified || !r.MinusThreeCertified || math.Abs(r.GHatPlusMean-1/math.Sqrt(31)) > blockTolerance || math.Abs(r.GHatMinusMean+3/math.Sqrt(31)) > blockTolerance || math.Abs(r.GHatMinusToPlusRatio+3) > blockTolerance || r.PlusMinusFrobNorm > blockTolerance {
			t.Fatalf("bad route: %+v", r)
		}
	}
	if !a.Multiplicity.PerDirectionWeightCertified || a.Multiplicity.DerivedBySymbolicTheorem || a.Multiplicity.NegativeSectorDim != 3 || a.Multiplicity.ObservedNegativeWeight != -3 {
		t.Fatalf("bad multiplicity: %+v", a.Multiplicity)
	}
	if !a.Angle.AngleFromBlockTrace || math.Abs(a.Angle.ComputedCosine-13/math.Sqrt(217)) > strictTolerance || math.Abs(a.Angle.ComputedResidualSq-48.0/217.0) > strictTolerance {
		t.Fatalf("bad angle: %+v", a.Angle)
	}
	if a.Firewalls.ClaimsSymbolicMultiplicityTheorem || a.Firewalls.ClaimsSplitG2 || a.Firewalls.ClaimsBoundaryStress || a.Firewalls.ClaimsSevenOver72Theorem || a.Firewalls.ClaimsScalarFlavor || a.Firewalls.ClaimsPhysicalMetric || a.Firewalls.ClaimsFlavor || a.Firewalls.ClaimsHiggsMass || a.Firewalls.ClaimsCKMPMNS || a.Firewalls.ClaimsGaugeUnification {
		t.Fatalf("firewall breach: %+v", a.Firewalls)
	}
}

func TestTheorem(t *testing.T) {
	res := Generation2NegativeSectorMultiplicityHitchinMetricSourceAuditTheorem().Verify()
	if res.Status == "FAILED_ROUTE" {
		t.Fatalf("unexpected construction failure: %+v", res)
	}
	for _, c := range res.Checks {
		if !c.Passed {
			t.Fatalf("check failed: %s %s", c.Name, c.Detail)
		}
	}
	joined := strings.Join(res.Notes, "\n")
	for _, want := range []string{StatusGate644ProjectorPlaneInherited, StatusOmegaSectorDecompositionComputed, StatusTwistedOmegaConstructed, StatusHitchinMetricBlockFormComputed, StatusNegativeSectorWeightCertified, StatusMinusThreeMultiplicityCandidate, StatusProjectiveAngleFromHitchinBlockTrace, StatusNoSymbolicMultiplicityTheorem, StatusNoCertifiedSplitG2, StatusNoBoundaryStress, StatusNoSevenOver72Theorem, StatusNoScalarFlavorTransport, StatusNoPhysicalMetric, StatusNoHiggsFlavorGauge, StatusGate645Boundary} {
		if !strings.Contains(joined, want) {
			t.Fatalf("missing note %s", want)
		}
	}
}
