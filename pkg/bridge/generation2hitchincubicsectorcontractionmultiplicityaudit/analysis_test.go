package generation2hitchincubicsectorcontractionmultiplicityaudit

import (
	"math"
	"strings"
	"testing"
)

func TestGate647Build(t *testing.T) {
	a, err := Build()
	if err != nil {
		t.Fatal(err)
	}
	if !a.Inherited.ProjectorPlaneIdentityInherited || !a.Inherited.RouteUniversal || a.Inherited.PositiveDim != 4 || a.Inherited.NegativeDim != 3 || a.Inherited.FullSymbolicTheoremCertified || a.Inherited.SplitG2Certified || a.Inherited.BoundaryStressAssignment || a.Inherited.SevenOver72Theorem || a.Inherited.ScalarFlavorTransport || a.Inherited.PhysicalMetric || !a.Inherited.Gate646FirewallPreserved {
		t.Fatalf("bad inheritance: %+v", a.Inherited)
	}
	if !a.Families.AllFamiliesAudited || len(a.Families.Families) != 4 || a.Families.AnySymbolicFamilyTheorem || !a.Families.AntisymmetrizedTwistsAudited {
		t.Fatalf("bad family ledger: %+v", a.Families)
	}
	seen := map[string]bool{}
	for _, f := range a.Families.Families {
		seen[f.Family] = true
		if f.SymbolicSourceFound {
			t.Fatalf("family unexpectedly promoted: %+v", f)
		}
	}
	for _, want := range []string{"Ω+++", "Ω++-", "Ω+--", "Ω---"} {
		if !seen[want] {
			t.Fatalf("missing family %s", want)
		}
	}
	if !a.Contributions.AllRoutesReconstruct || !a.Contributions.AllRoutesBlockRayCertified || !a.Contributions.SameProjectorPlaneShadow || a.Contributions.AnyRouteSymbolicCertified || len(a.Contributions.Routes) != 3 {
		t.Fatalf("bad contributions: %+v", a.Contributions)
	}
	for _, r := range a.Contributions.Routes {
		if r.AdditiveReconstructionError > reconstructionTolerance || !r.BlockRayCertified || !r.MinusQCertified || !r.OffBlockCancelledAtTotal || r.SymbolicContractionCertified {
			t.Fatalf("bad route ledger: %+v", r)
		}
		if math.Abs(r.NormalizedPositiveWeight-1) > strictTolerance || math.Abs(r.NormalizedNegativeWeight+3) > blockTolerance || math.Abs(r.RawMinusToPlusRatio+3) > blockTolerance {
			t.Fatalf("bad route weights: %+v", r)
		}
		if r.SignificantTripleCount == 0 || r.TotalTripleCount != 64 || len(r.TopContributions) == 0 {
			t.Fatalf("missing triple contributions: %+v", r)
		}
	}
	if !a.PositiveUnit.AllRoutesUnitPositive || a.PositiveUnit.SymbolicUnitTheoremCertified || a.PositiveUnit.MaxPositiveWeightDrift > strictTolerance {
		t.Fatalf("bad positive unit: %+v", a.PositiveUnit)
	}
	if !a.NegativeMultiplicity.AllRoutesMinusQ || !a.NegativeMultiplicity.CubicSectorMultiplicitySupported || a.NegativeMultiplicity.SymbolicMultiplicityTheoremCertified || math.Abs(a.NegativeMultiplicity.TargetRatio+3) > strictTolerance || a.NegativeMultiplicity.MaxRatioDrift > blockTolerance {
		t.Fatalf("bad negative multiplicity: %+v", a.NegativeMultiplicity)
	}
	if a.OffBlockCancellation.MaxTotalOffBlockNorm > blockTolerance || a.OffBlockCancellation.StructuralCancellationCertified {
		t.Fatalf("bad off-block: %+v", a.OffBlockCancellation)
	}
	if a.RouteUniversality.RouteCount != 3 || !a.RouteUniversality.AllRoutesSameFinalRay || a.RouteUniversality.ComponentContributionLedgersEqual {
		t.Fatalf("bad route universality: %+v", a.RouteUniversality)
	}
	if !a.TheoremReadiness.FiniteLedgerSupportsTheorem || !a.TheoremReadiness.ComponentLedgerComputed || !a.TheoremReadiness.BlockContributionComputed || a.TheoremReadiness.FullSymbolicTheoremCertified {
		t.Fatalf("bad theorem readiness: %+v", a.TheoremReadiness)
	}
	if a.Firewalls.ClaimsFullSymbolicHitchinTheorem || a.Firewalls.ClaimsSplitG2 || a.Firewalls.ClaimsBoundaryStress || a.Firewalls.ClaimsSevenOver72Theorem || a.Firewalls.ClaimsScalarFlavor || a.Firewalls.ClaimsPhysicalMetric || a.Firewalls.ClaimsFlavor || a.Firewalls.ClaimsHiggsMass || a.Firewalls.ClaimsCKMPMNS || a.Firewalls.ClaimsGaugeUnification {
		t.Fatalf("firewall breach: %+v", a.Firewalls)
	}
}

func TestTheorem(t *testing.T) {
	res := Generation2HitchinCubicSectorContractionMultiplicityAuditTheorem().Verify()
	if res.Status == "FAILED_ROUTE" {
		t.Fatalf("unexpected construction failure: %+v", res)
	}
	for _, c := range res.Checks {
		if !c.Passed {
			t.Fatalf("check failed: %s %s", c.Name, c.Detail)
		}
	}
	joined := strings.Join(res.Notes, "\n")
	for _, want := range []string{StatusGate646ProjectorPlaneInherited, StatusComponentFamilyLedgerComputed, StatusHitchinBlockContributionComputed, StatusPositiveSectorUnitCoefficientAudited, StatusNegativeSectorMultiplicityAudited, StatusOffBlockCancellationSourceAudited, StatusRouteUniversalityComparisonComputed, StatusMinusQFromCubicSectorMultiplicity, StatusHitchinMultiplicityTheoremSharpened, StatusSameProjectorPlaneRouteUniversal, StatusNoFullSymbolicHitchinTheorem, StatusNoCertifiedSplitG2, StatusNoBoundaryStress, StatusNoSevenOver72Theorem, StatusNoScalarFlavorTransport, StatusNoPhysicalMetric, StatusNoHiggsFlavorGauge, StatusGate647Boundary} {
		if !strings.Contains(joined, want) {
			t.Fatalf("missing note %s", want)
		}
	}
}
