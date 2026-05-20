package generation2hodgestarinternaldestinationaudit

import (
	"math"
	"strings"
	"testing"
)

func TestGate633Build(t *testing.T) {
	a, err := Build()
	if err != nil {
		t.Fatal(err)
	}
	if !a.Inherited.TransverseFailureCertified || !a.Inherited.StarK7InsideUPlusV || !a.Inherited.Gate632FirewallPreserved || a.Inherited.LeakageRank != 0 {
		t.Fatalf("bad Gate632 inheritance: %+v", a.Inherited)
	}
	if a.L7.Rank != 7 || a.L7.QLOthonormalResidual > 1e-12 || a.L7.PWQLFrobeniusNorm > 1e-10 || math.Abs(a.L7.PUVProjectionFraction-1) > 1e-8 || a.L7.PUVContainmentResidual > 1e-10 || a.L7.StarTwoCycleResidual > 1e-12 || !a.L7.InternalContainmentCertified {
		t.Fatalf("bad L7 certificate: %+v", a.L7)
	}
	if a.K7Preservation.Rank != 7 || !a.K7Preservation.StarPreservesK7 || !a.K7Preservation.HodgeStable || math.Abs(a.K7Preservation.ProjectionFraction-1) > 1e-8 || a.K7Preservation.ContainmentResidual > 1e-10 || math.Abs(math.Abs(a.K7Preservation.Determinant)-1) > 1e-8 {
		t.Fatalf("K7 should be Hodge-stable: %+v", a.K7Preservation)
	}
	if a.T56Complement.Rank != 0 || a.T56Complement.ProjectionFraction > 1e-20 || a.T56Complement.ContainmentResidual < 2.6 || a.T56Complement.L7InsideT56 {
		t.Fatalf("L7 should not enter T56 complement: %+v", a.T56Complement)
	}
	if a.OctonionicResidual.Rank != 0 || a.OctonionicResidual.ProjectionFraction > 1e-20 || a.OctonionicResidual.ContainmentResidual < 2.6 || a.OctonionicResidual.L7EqualsV0 || a.OctonionicResidual.VDecomposesAsK7StarK7 {
		t.Fatalf("L7 should not equal V0: %+v", a.OctonionicResidual)
	}
	if a.BooleanResidual.Rank != 0 || a.BooleanResidual.ProjectionFraction > 1e-20 || a.BooleanResidual.ContainmentResidual < 2.6 || a.BooleanResidual.L7InsideU0 {
		t.Fatalf("L7 should not enter U0: %+v", a.BooleanResidual)
	}
	if a.ObliqueDecomposition.ObliquePlaneDetected || a.ObliqueDecomposition.ProjectionFraction > 1e-20 || a.ObliqueDecomposition.DirectSumCoordinateNorm > 1e-10 {
		t.Fatalf("bad oblique decomposition audit: %+v", a.ObliqueDecomposition)
	}
	if !a.StarTwoCycle.L7EqualsK7 || a.StarTwoCycle.L7EqualsV0 || a.StarTwoCycle.L7EqualsU0 || a.StarTwoCycle.StarSquaredResidual > 1e-12 {
		t.Fatalf("bad star two-cycle audit: %+v", a.StarTwoCycle)
	}
	if !a.ConsequenceFor7Over72.K7Stable || a.ConsequenceFor7Over72.Octonionic14HodgeSplit || a.ConsequenceFor7Over72.NewSevenPlaneDiscovered || a.ConsequenceFor7Over72.TraceWeightPromoted || !a.ConsequenceFor7Over72.BoundaryAssignmentMissing {
		t.Fatalf("bad 7/72 consequence: %+v", a.ConsequenceFor7Over72)
	}
	if a.Firewalls.ClaimsBoundaryStressAssignment || a.Firewalls.ClaimsSevenOver72Theorem || a.Firewalls.ClaimsScalarRGMatching || a.Firewalls.ClaimsHiggsMassDerivation || a.Firewalls.ClaimsFlavorDerivation || a.Firewalls.ClaimsCKMPMNSDerivation || a.Firewalls.ClaimsGaugeUnification {
		t.Fatalf("firewall breach: %+v", a.Firewalls)
	}
}

func TestTheorem(t *testing.T) {
	res := Generation2HodgeStarInternalDestinationAuditTheorem().Verify()
	if res.Status == "FAILED_ROUTE" {
		t.Fatalf("unexpected construction failure: %+v", res)
	}
	for _, c := range res.Checks {
		if !c.Passed {
			t.Fatalf("check failed: %s %s", c.Name, c.Detail)
		}
	}
	joined := strings.Join(res.Notes, "\n")
	for _, want := range []string{StatusGate632Inherited, StatusHodgeCompanionL7Defined, StatusHodgeStarPreservesK7, StatusNoNewInternalHodgeCompanion, StatusDoesNotPairOctonionicResidualV0, StatusNoK7ToW7Pairing, StatusNoBoundaryStressAssignment, StatusGate633Boundary} {
		if !strings.Contains(joined, want) {
			t.Fatalf("missing note %s", want)
		}
	}
}
