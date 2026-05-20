package generation2hodgepolarityprojectiveangletraceidentityaudit

import (
	"math"
	"strings"
	"testing"
)

func TestGate642Build(t *testing.T) {
	a, err := Build()
	if err != nil {
		t.Fatal(err)
	}
	if !a.Inherited.ComplementIdentified || !a.Inherited.ProjectiveAngleAudited || !a.Inherited.ThirteenSourcesAudited || a.Inherited.TraceIdentityCertifiedByGate641 || a.Inherited.SplitG2CertifiedByGate641 || a.Inherited.BoundaryStressByGate641 || a.Inherited.SevenOver72TheoremByGate641 || a.Inherited.ScalarFlavorByGate641 || a.Inherited.PhysicalAngleByGate641 || !a.Inherited.Gate641FirewallPreserved {
		t.Fatalf("bad inheritance: %+v", a.Inherited)
	}
	if math.Abs(a.Inherited.CosTheta-(float64(alignmentRoot)/math.Sqrt(float64(angleDenominator)))) > traceIdentityTolerance {
		t.Fatalf("bad cosine: %+v", a.Inherited)
	}
	if math.Abs(a.Inherited.SinTheta-(4*math.Sqrt(3)/math.Sqrt(float64(angleDenominator)))) > traceIdentityTolerance {
		t.Fatalf("bad sine: %+v", a.Inherited)
	}
	if !a.RawContractions.IntegerRatioVerified || a.RawContractions.NativeTraceIdentityFound || len(a.RawContractions.Contractions) < 3 || a.RawContractions.MaxCosSquaredDelta > traceIdentityTolerance || a.RawContractions.MaxSinSquaredDelta > traceIdentityTolerance {
		t.Fatalf("bad raw contractions: %+v", a.RawContractions)
	}
	seen := map[string]bool{}
	for _, c := range a.RawContractions.Contractions {
		seen[c.RouteName] = true
		if !c.ProjectivePairMatches || c.IntegerInnerProductSquare != alignmentNumerator || c.IntegerFailureSquare != failureNumerator || c.IntegerProductNormSquare != angleDenominator {
			t.Fatalf("bad contraction: %+v", c)
		}
	}
	for _, want := range []string{"omega_1_alt", "omega_2_alt", "omega_B_alt"} {
		if !seen[want] {
			t.Fatalf("missing contraction route %s", want)
		}
	}
	if !a.SectorBlocks.BlockSkeletonMatches || a.SectorBlocks.NativeTraceIdentity || a.SectorBlocks.PDim != 4 || a.SectorBlocks.QDim != 3 || a.SectorBlocks.AlignmentAmplitude != 13 || a.SectorBlocks.FailureAmplitudeSquared != 48 || a.SectorBlocks.Denominator != 217 {
		t.Fatalf("bad sector block audit: %+v", a.SectorBlocks)
	}
	if !a.ProjectivePair.PairMatches || a.ProjectivePair.DerivedFromNativeTraceIdentity || a.ProjectivePair.TanSquaredNumerator != 48 || a.ProjectivePair.TanSquaredDenominator != 169 || a.ProjectivePair.PythagoreanIntegerResidual != 0 {
		t.Fatalf("bad projective pair: %+v", a.ProjectivePair)
	}
	if a.TraceIdentity.NativeTraceIdentityFound || a.TraceIdentity.BestCandidateResidual > traceIdentityTolerance {
		t.Fatalf("bad trace identity audit: %+v", a.TraceIdentity)
	}
	if !a.Classification.ProjectiveAngleInherited || !a.Classification.RawContractionsComputed || !a.Classification.HodgeSectorBlocksComputed || !a.Classification.BlockSkeletonSupported || a.Classification.NativeTraceIdentityCertified || !a.Classification.ObstructionOnly {
		t.Fatalf("bad classification: %+v", a.Classification)
	}
	if a.Firewalls.ClaimsNativeTraceIdentity || a.Firewalls.ClaimsSplitG2 || a.Firewalls.ClaimsBoundaryStress || a.Firewalls.ClaimsSevenOver72Theorem || a.Firewalls.ClaimsScalarFlavor || a.Firewalls.ClaimsPhysicalAngle || a.Firewalls.ClaimsPhysicalMetric || a.Firewalls.ClaimsFlavor || a.Firewalls.ClaimsHiggsMass || a.Firewalls.ClaimsCKMPMNS || a.Firewalls.ClaimsGaugeUnification {
		t.Fatalf("firewall breach: %+v", a.Firewalls)
	}
}

func TestTheorem(t *testing.T) {
	res := Generation2HodgePolarityProjectiveAngleTraceIdentityAuditTheorem().Verify()
	if res.Status == "FAILED_ROUTE" {
		t.Fatalf("unexpected construction failure: %+v", res)
	}
	for _, c := range res.Checks {
		if !c.Passed {
			t.Fatalf("check failed: %s %s", c.Name, c.Detail)
		}
	}
	joined := strings.Join(res.Notes, "\n")
	for _, want := range []string{StatusGate641AngleInherited, StatusRawContractionsComputed, StatusHodgeSectorBlocksComputed, StatusProjectivePairAudited, StatusHodgePolarityBlockSkeleton, StatusOffSectorObstructionBlock, StatusTraceIdentityCandidates, StatusNoNativeTraceIdentity, StatusNoCertifiedSplitG2, StatusNoBoundaryStress, StatusNoSevenOver72Theorem, StatusNoScalarFlavorTransport, StatusNoPhysicalAngle, StatusNoPhysicalMetric, StatusNoHiggsFlavorGauge, StatusGate642Boundary} {
		if !strings.Contains(joined, want) {
			t.Fatalf("missing note %s", want)
		}
	}
}
