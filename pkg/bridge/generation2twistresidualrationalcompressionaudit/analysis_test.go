package generation2twistresidualrationalcompressionaudit

import (
	"math"
	"strings"
	"testing"
)

func TestGate640Build(t *testing.T) {
	a, err := Build()
	if err != nil {
		t.Fatal(err)
	}
	if !a.Inherited.RepeatedAcrossRoutes || !a.Inherited.ResidualInvariant || !a.Inherited.CompactSplitObstruction || a.Inherited.Gate639ClassifiedAsArtifact || a.Inherited.Gate639SplitG2Certified || a.Inherited.Gate639BoundaryStressAssignment || a.Inherited.Gate639SevenOver72Theorem || a.Inherited.Gate639ScalarFlavorTransport || !a.Inherited.Gate639FirewallPreserved {
		t.Fatalf("bad inheritance: %+v", a.Inherited)
	}
	if !a.Compression.Compressed || math.Abs(a.Compression.ResidualSquared) > rationalTolerance || math.Abs(a.Compression.RhoResidual) > rationalTolerance {
		t.Fatalf("bad compression: %+v", a.Compression)
	}
	if a.Compression.CandidateNumerator != 48 || a.Compression.CandidateDenominator != 217 {
		t.Fatalf("bad rational candidate: %+v", a.Compression)
	}
	if !a.Routes.AllClusterRoutesCompress || len(a.Routes.Routes) < 3 || a.Routes.MaxSquaredDelta > routeTolerance {
		t.Fatalf("bad route compression: %+v", a.Routes)
	}
	seen := map[string]bool{}
	for _, name := range a.Routes.CompressedRouteNames {
		seen[name] = true
	}
	for _, want := range []string{"omega_1_alt", "omega_2_alt", "omega_B_alt"} {
		if !seen[want] {
			t.Fatalf("missing compressed route %s: %+v", want, a.Routes.CompressedRouteNames)
		}
	}
	if !a.Skeleton.NumeratorMatches || !a.Skeleton.DenominatorMatches || a.Skeleton.NumeratorFromHodgePolarity != 48 || a.Skeleton.DenominatorFromSelfDualGap != 217 || a.Skeleton.SelfDualComplementToK7PlusDim != 31 {
		t.Fatalf("bad skeleton: %+v", a.Skeleton)
	}
	if a.Projectors.TraceDerivationCertified || a.Projectors.BestCandidateResidual > rationalTolerance {
		t.Fatalf("bad projector audit: %+v", a.Projectors)
	}
	if !a.Classification.CompressionCandidate || a.Classification.ExactFromFiniteMatrixClaim || a.Classification.ConsequenceOfHodgeSplitClaim || a.Classification.ArtifactClaim || !a.Classification.ObstructionOnly {
		t.Fatalf("bad classification: %+v", a.Classification)
	}
	if a.Firewalls.ClaimsExactTraceTheorem || a.Firewalls.ClaimsSplitG2 || a.Firewalls.ClaimsBoundaryStress || a.Firewalls.ClaimsSevenOver72Theorem || a.Firewalls.ClaimsScalarFlavor || a.Firewalls.ClaimsPhysicalMetric || a.Firewalls.ClaimsFlavor || a.Firewalls.ClaimsHiggsMass || a.Firewalls.ClaimsCKMPMNS || a.Firewalls.ClaimsGaugeUnification {
		t.Fatalf("firewall breach: %+v", a.Firewalls)
	}
}

func TestTheorem(t *testing.T) {
	res := Generation2TwistResidualRationalCompressionAuditTheorem().Verify()
	if res.Status == "FAILED_ROUTE" {
		t.Fatalf("unexpected construction failure: %+v", res)
	}
	for _, c := range res.Checks {
		if !c.Passed {
			t.Fatalf("check failed: %s %s", c.Name, c.Detail)
		}
	}
	joined := strings.Join(res.Notes, "\n")
	for _, want := range []string{StatusGate639RhoInherited, StatusRhoSquaredCompressionTested, StatusRhoSquaredEquals48Over217, StatusDenominator217TypedCandidate, StatusNumerator48TypedCandidate, StatusRouteCompressionRepeated, StatusProjectorContractionsAudited, StatusNoTraceDerivation, StatusNoCertifiedSplitG2, StatusNoBoundaryStress, StatusNoSevenOver72Theorem, StatusNoScalarFlavorTransport, StatusNoPhysicalMetric, StatusGate640Boundary} {
		if !strings.Contains(joined, want) {
			t.Fatalf("missing note %s", want)
		}
	}
}
