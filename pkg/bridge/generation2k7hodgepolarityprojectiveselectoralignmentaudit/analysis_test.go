package generation2k7hodgepolarityprojectiveselectoralignmentaudit

import (
	"math"
	"strings"
	"testing"
)

func TestGate635Build(t *testing.T) {
	a, err := Build()
	if err != nil {
		t.Fatal(err)
	}
	if !a.Inherited.HodgeStable || !a.Inherited.MixedHodgePolarity || !a.Inherited.NoBoundaryAssignment || !a.Inherited.NoSevenOver72Theorem || !a.Inherited.Gate634FirewallPreserved {
		t.Fatalf("bad Gate634 inheritance: %+v", a.Inherited)
	}
	if a.Inherited.K7Dimension != 7 || a.Inherited.PlusDimension != 4 || a.Inherited.MinusDimension != 3 || math.Abs(a.Inherited.Trace-1) > 1e-10 || math.Abs(a.Inherited.Determinant+1) > 1e-8 {
		t.Fatalf("bad inherited signature values: %+v", a.Inherited)
	}
	if !a.K7Subspaces.ProjectorsCertified || !a.K7Subspaces.PlusMinusOrthogonal || !a.K7Subspaces.NativeHodgePolarity || a.K7Subspaces.PlusDimension != 4 || a.K7Subspaces.MinusDimension != 3 || a.K7Subspaces.SumDimension != 7 {
		t.Fatalf("bad K7 subspaces: %+v", a.K7Subspaces)
	}
	if a.ProjectiveSelector.CarrierComplexDimension != 4 || a.ProjectiveSelector.LineComplexDimension != 1 || a.ProjectiveSelector.SpatialBlockComplexDimension != 3 || !a.ProjectiveSelector.CP0CP2CriticalStrata || !a.ProjectiveSelector.ProjectiveOnePlusThree || !a.ProjectiveSelector.MatchesGate555Commutant || a.ProjectiveSelector.CP3ToK7FunctorFound {
		t.Fatalf("bad projective selector inheritance: %+v", a.ProjectiveSelector)
	}
	if !a.Alignment.FourDimensionalMatch || !a.Alignment.ThreeDimensionalMatch || a.Alignment.SameCarrier || a.Alignment.TypedThetaMapFound || a.Alignment.K7ToCP3FunctorFound || !a.Alignment.AlignmentCandidateOnly {
		t.Fatalf("bad alignment posture: %+v", a.Alignment)
	}
	if !strings.Contains(a.Alignment.Verdict, StatusResemblesSelectorSplit) || !strings.Contains(a.Alignment.Verdict, StatusNoK7ToFockMap) {
		t.Fatalf("bad alignment verdict: %s", a.Alignment.Verdict)
	}
	if !a.K7PlusRefinement.HodgeProjectorActsAsIdentity || a.K7PlusRefinement.InternalRankOneLineDerived || a.K7PlusRefinement.InternalThreePlaneDerived || a.K7PlusRefinement.NativeOnePlusThreeRefinement || a.K7PlusRefinement.Verdict != StatusNoK7PlusOnePlusThree {
		t.Fatalf("bad K7+ refinement audit: %+v", a.K7PlusRefinement)
	}
	if !a.K7MinusTriplet.DimensionMatch || a.K7MinusTriplet.TypedTripletIdentification || a.K7MinusTriplet.UsesBMinusLCarrier {
		t.Fatalf("bad K7- triplet audit: %+v", a.K7MinusTriplet)
	}
	if math.Abs(a.TraceImbalance.Trace-1) > 1e-10 || a.TraceImbalance.PlusMinusDifference != 1 || a.TraceImbalance.DistinguishedLineDerived || a.TraceImbalance.TraceAsRankOneProjector || !a.TraceImbalance.NeedsAdditionalSelector {
		t.Fatalf("bad trace imbalance audit: %+v", a.TraceImbalance)
	}
	if !a.CarrierMap.DimensionResemblance || a.CarrierMap.TypedIntertwinerFound || a.CarrierMap.FunctorFromProjectiveFockToK7Found || a.CarrierMap.FunctorFromK7ToProjectiveFockFound || !strings.Contains(a.CarrierMap.MissingObject, "Theta") {
		t.Fatalf("bad carrier map audit: %+v", a.CarrierMap)
	}
	if a.BoundaryReadiness.BoundaryStressAssignment || a.BoundaryReadiness.SevenOver72Promoted || a.BoundaryReadiness.K7ToW7PairingReopened || a.BoundaryReadiness.K7ToFockMapWouldSuffice {
		t.Fatalf("boundary readiness breached: %+v", a.BoundaryReadiness)
	}
	if a.Firewalls.ClaimsK7ToFockSelectorMap || a.Firewalls.ClaimsK7PlusOnePlusThree || a.Firewalls.ClaimsBoundaryStressAssignment || a.Firewalls.ClaimsSevenOver72Theorem || a.Firewalls.ClaimsScalarRGMatching || a.Firewalls.ClaimsHiggsMassDerivation || a.Firewalls.ClaimsFlavorDerivation || a.Firewalls.ClaimsCKMPMNSDerivation || a.Firewalls.ClaimsGaugeUnification || a.Firewalls.ClaimsPhysicalOrientation {
		t.Fatalf("firewall breach: %+v", a.Firewalls)
	}
}

func TestTheorem(t *testing.T) {
	res := Generation2K7HodgePolarityProjectiveSelectorAlignmentAuditTheorem().Verify()
	if res.Status == "FAILED_ROUTE" {
		t.Fatalf("unexpected construction failure: %+v", res)
	}
	for _, c := range res.Checks {
		if !c.Passed {
			t.Fatalf("check failed: %s %s", c.Name, c.Detail)
		}
	}
	joined := strings.Join(res.Notes, "\n")
	for _, want := range []string{StatusGate634Inherited, StatusK7PlusMinusSubspacesDefined, StatusProjectiveSelectorInherited, StatusFourPlusThreeAudited, StatusResemblesSelectorSplit, StatusNoK7ToFockMap, StatusNoK7PlusOnePlusThree, StatusTraceNotDistinguishedLine, StatusNoBoundaryStressAssignment, StatusGate635Boundary} {
		if !strings.Contains(joined, want) {
			t.Fatalf("missing note %s", want)
		}
	}
}
