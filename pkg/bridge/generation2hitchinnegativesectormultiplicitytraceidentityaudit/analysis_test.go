package generation2hitchinnegativesectormultiplicitytraceidentityaudit

import (
	"math"
	"strings"
	"testing"
)

func TestGate646Build(t *testing.T) {
	a, err := Build()
	if err != nil {
		t.Fatal(err)
	}
	if !a.Inherited.NegativeWeightCertified || !a.Inherited.ProjectiveAngleDerived || !a.Inherited.ComponentAuditComputed || a.Inherited.RouteCount != 3 || !a.Inherited.MinusThreeSourceCandidate || a.Inherited.FullSymbolicTheoremCertified || a.Inherited.SplitG2Certified || a.Inherited.BoundaryStressAssignment || a.Inherited.SevenOver72Theorem || a.Inherited.ScalarFlavorTransport || a.Inherited.PhysicalMetric || !a.Inherited.Gate645FirewallPreserved {
		t.Fatalf("bad inheritance: %+v", a.Inherited)
	}
	if !a.Components.AllFamiliesAudited || len(a.Components.Families) != 4 || a.Components.AnyFamilyContributionCertified || a.Components.SymbolicComponentTheoremCertified {
		t.Fatalf("bad components: %+v", a.Components)
	}
	seen := map[string]bool{}
	for _, f := range a.Components.Families {
		seen[f.Family] = true
		if f.BlockContributionCertified || f.SymbolicContributionCertified {
			t.Fatalf("family unexpectedly promoted: %+v", f)
		}
	}
	for _, want := range []string{"Ω+++", "Ω++-", "Ω+--", "Ω---"} {
		if !seen[want] {
			t.Fatalf("missing family %s", want)
		}
	}
	if !a.OffBlockCancellation.NumericalCancellation || a.OffBlockCancellation.MaxOffBlockFrobeniusNorm > blockTolerance || a.OffBlockCancellation.SymbolicCancellationCertified {
		t.Fatalf("bad off-block audit: %+v", a.OffBlockCancellation)
	}
	if !a.PositiveUnit.UnitWeightCertified || a.PositiveUnit.SymbolicUnitWeightCertified || a.PositiveUnit.PositiveDim != 4 || math.Abs(a.PositiveUnit.ObservedPositiveWeight-1) > strictTolerance {
		t.Fatalf("bad positive audit: %+v", a.PositiveUnit)
	}
	if !a.NegativeMultiplicity.MultiplicityWeightCertified || a.NegativeMultiplicity.SymbolicMultiplicityCertified || a.NegativeMultiplicity.NegativeDim != 3 || a.NegativeMultiplicity.ObservedNegativeWeight != -3 || math.Abs(a.NegativeMultiplicity.MaxRatioDrift) > blockTolerance {
		t.Fatalf("bad negative audit: %+v", a.NegativeMultiplicity)
	}
	if !a.ProjectorIdentity.IdentityMatchesRouteData || a.ProjectorIdentity.FullSymbolicTheoremCertified || math.Abs(a.ProjectorIdentity.GHatNormalizerSq-31) > strictTolerance || math.Abs(a.ProjectorIdentity.BHatNormalizerSq-7) > strictTolerance || math.Abs(a.ProjectorIdentity.Cosine-13/math.Sqrt(217)) > strictTolerance || math.Abs(a.ProjectorIdentity.ResidualSquared-48.0/217.0) > strictTolerance {
		t.Fatalf("bad projector identity: %+v", a.ProjectorIdentity)
	}
	if !a.RouteUniversality.AllRoutesPass || !a.RouteUniversality.RouteUniversalCandidate || a.RouteUniversality.RouteDependentFailure || len(a.RouteUniversality.Routes) != 3 {
		t.Fatalf("bad route universality: %+v", a.RouteUniversality)
	}
	for _, r := range a.RouteUniversality.Routes {
		if !r.MatchesPQIdentity || !r.BlockFormCertified || math.Abs(r.MinusToPlusRatio+3) > blockTolerance || r.OffBlockFrobeniusNorm > blockTolerance {
			t.Fatalf("bad route row: %+v", r)
		}
	}
	if a.Interpretation.SymbolicTheoremCertified || !a.Interpretation.PQIdentityMatches || !a.Interpretation.RouteUniversal {
		t.Fatalf("bad interpretation: %+v", a.Interpretation)
	}
	if a.Firewalls.ClaimsFullSymbolicHitchinTheorem || a.Firewalls.ClaimsSplitG2 || a.Firewalls.ClaimsBoundaryStress || a.Firewalls.ClaimsSevenOver72Theorem || a.Firewalls.ClaimsScalarFlavor || a.Firewalls.ClaimsPhysicalMetric || a.Firewalls.ClaimsFlavor || a.Firewalls.ClaimsHiggsMass || a.Firewalls.ClaimsCKMPMNS || a.Firewalls.ClaimsGaugeUnification {
		t.Fatalf("firewall breach: %+v", a.Firewalls)
	}
}

func TestTheorem(t *testing.T) {
	res := Generation2HitchinNegativeSectorMultiplicityTraceIdentityAuditTheorem().Verify()
	if res.Status == "FAILED_ROUTE" {
		t.Fatalf("unexpected construction failure: %+v", res)
	}
	for _, c := range res.Checks {
		if !c.Passed {
			t.Fatalf("check failed: %s %s", c.Name, c.Detail)
		}
	}
	joined := strings.Join(res.Notes, "\n")
	for _, want := range []string{StatusGate645NegativeWeightInherited, StatusHitchinBlockComponentAuditComputed, StatusOffBlockCancellationAudited, StatusPositiveSectorUnitWeightAudited, StatusNegativeSectorMultiplicityAudited, StatusProjectorPlaneIdentityDerived, StatusRouteUniversalityAudited, StatusMinusThreeMultiplicity, StatusAngleFromPQTraceIdentity, StatusRouteUniversalHitchinIdentity, StatusNoFullSymbolicHitchinTheorem, StatusNoCertifiedSplitG2, StatusNoBoundaryStress, StatusNoSevenOver72Theorem, StatusNoScalarFlavorTransport, StatusNoPhysicalMetric, StatusNoHiggsFlavorGauge, StatusGate646Boundary} {
		if !strings.Contains(joined, want) {
			t.Fatalf("missing note %s", want)
		}
	}
}
