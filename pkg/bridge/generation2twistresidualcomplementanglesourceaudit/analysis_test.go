package generation2twistresidualcomplementanglesourceaudit

import (
	"math"
	"strings"
	"testing"
)

func TestGate641Build(t *testing.T) {
	a, err := Build()
	if err != nil {
		t.Fatal(err)
	}
	if !a.Inherited.RhoSquaredCompresses || !a.Inherited.RouteCompressionRepeated || !a.Inherited.DimensionalSkeletonTyped || a.Inherited.TraceDerivationCertifiedByGate640 || a.Inherited.SplitG2CertifiedByGate640 || a.Inherited.BoundaryStressByGate640 || a.Inherited.SevenOver72TheoremByGate640 || a.Inherited.ScalarFlavorByGate640 || a.Inherited.PhysicalMetricByGate640 || !a.Inherited.Gate640FirewallPreserved {
		t.Fatalf("bad inheritance: %+v", a.Inherited)
	}
	if !a.Complement.ComplementIdentified || a.Complement.AlignmentRoot != 13 || a.Complement.AlignmentNumerator != 169 || a.Complement.Denominator != 217 || math.Abs(a.Complement.ComplementResidual) > complementTolerance || math.Abs(a.Complement.PythagoreanResidual) > complementTolerance {
		t.Fatalf("bad complement: %+v", a.Complement)
	}
	if math.Abs(a.Complement.CosTheta-(13/math.Sqrt(217))) > complementTolerance {
		t.Fatalf("bad cosine: %+v", a.Complement)
	}
	if math.Abs(a.Complement.TanTheta-(4*math.Sqrt(3)/13)) > complementTolerance {
		t.Fatalf("bad tangent: %+v", a.Complement)
	}
	if !a.Projective.AllRoutesAlign || len(a.Projective.Contractions) < 3 || a.Projective.MaxCosSquaredDelta > routeComplementTolerance {
		t.Fatalf("bad projective audit: %+v", a.Projective)
	}
	seen := map[string]bool{}
	for _, c := range a.Projective.Contractions {
		seen[c.RouteName] = true
		if !c.Matches13Squared || math.Abs(c.SinSquared+c.CosSquared-1) > routeComplementTolerance || c.NormGTwist != 1 || c.NormBK != 1 {
			t.Fatalf("bad contraction: %+v", c)
		}
	}
	for _, want := range []string{"omega_1_alt", "omega_2_alt", "omega_B_alt"} {
		if !seen[want] {
			t.Fatalf("missing contraction route %s", want)
		}
	}
	if !a.Thirteen.StrongestCandidateTyped || a.Thirteen.StrongestCandidateValue != 13 || a.Thirteen.TraceIdentityCertified {
		t.Fatalf("bad thirteen audit: %+v", a.Thirteen)
	}
	if a.TraceIdentity.NativeTraceIdentityFound || a.TraceIdentity.BestCandidateResidual > complementTolerance {
		t.Fatalf("bad trace identity audit: %+v", a.TraceIdentity)
	}
	if !a.Classification.SinSquared48Over217 || !a.Classification.CosSquared169Over217 || !a.Classification.FiniteAngleCandidate || a.Classification.TraceAngleDecomposition || a.Classification.NormalizationArtifact || !a.Classification.ObstructionOnly {
		t.Fatalf("bad classification: %+v", a.Classification)
	}
	if a.Firewalls.ClaimsNativeTraceIdentity || a.Firewalls.ClaimsSplitG2 || a.Firewalls.ClaimsBoundaryStress || a.Firewalls.ClaimsSevenOver72Theorem || a.Firewalls.ClaimsScalarFlavor || a.Firewalls.ClaimsPhysicalAngle || a.Firewalls.ClaimsPhysicalMetric || a.Firewalls.ClaimsFlavor || a.Firewalls.ClaimsHiggsMass || a.Firewalls.ClaimsCKMPMNS || a.Firewalls.ClaimsGaugeUnification {
		t.Fatalf("firewall breach: %+v", a.Firewalls)
	}
}

func TestTheorem(t *testing.T) {
	res := Generation2TwistResidualComplementAngleSourceAuditTheorem().Verify()
	if res.Status == "FAILED_ROUTE" {
		t.Fatalf("unexpected construction failure: %+v", res)
	}
	for _, c := range res.Checks {
		if !c.Passed {
			t.Fatalf("check failed: %s %s", c.Name, c.Detail)
		}
	}
	joined := strings.Join(res.Notes, "\n")
	for _, want := range []string{StatusGate640RhoSquaredInherited, StatusComplement169Identified, StatusProjectiveAlignmentAngleAudited, StatusAlignment13SquaredCandidate, StatusThirteenSourceCandidatesAudited, StatusRouteComplementRepeated, StatusRawFrobeniusContractionsAudited, StatusTraceIdentitySearched, StatusNoNativeTraceIdentityFor13, StatusNoCertifiedSplitG2, StatusNoBoundaryStress, StatusNoSevenOver72Theorem, StatusNoScalarFlavorTransport, StatusNoPhysicalAngle, StatusNoPhysicalMetric, StatusNoHiggsFlavorGauge, StatusGate641Boundary} {
		if !strings.Contains(joined, want) {
			t.Fatalf("missing note %s", want)
		}
	}
}
