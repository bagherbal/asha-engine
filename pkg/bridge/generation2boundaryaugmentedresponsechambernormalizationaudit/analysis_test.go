package generation2boundaryaugmentedresponsechambernormalizationaudit

import (
	"strings"
	"testing"
)

func TestGate917InheritedFirstFourSubobjects(t *testing.T) {
	a, err := BuildDefault()
	if err != nil {
		t.Fatal(err)
	}
	if !a.Inherited.ResponseShapeAudited || !a.Inherited.SelectorShapeAudited || !a.Inherited.CrossLaneShapeAudited || !a.Inherited.SSplitTransportAudited || a.Inherited.NativeAlpha || a.Inherited.NativeR3 {
		t.Fatalf("bad inherited state: %s", FormatInherited(a.Inherited))
	}
}

func TestGate917H10AndH72Chambers(t *testing.T) {
	a, err := BuildDefault()
	if err != nil {
		t.Fatal(err)
	}
	if a.DegreeOne.ChamberRank != 10 || a.DegreeOne.AmbientRank+a.DegreeOne.BoundaryRank != 10 || a.DegreeOne.NativeActivation {
		t.Fatalf("bad H10: %s", FormatDegreeOne(a.DegreeOne))
	}
	if a.DegreeTwo.ChamberRank != 72 || a.DegreeTwo.Lambda4Rank+a.DegreeTwo.BoundaryRank != 72 || a.DegreeTwo.NativeActivation {
		t.Fatalf("bad H72: %s", FormatDegreeTwo(a.DegreeTwo))
	}
}

func TestGate917LaneCompatibilityAndBoundaryAugmentation(t *testing.T) {
	a, err := BuildDefault()
	if err != nil {
		t.Fatal(err)
	}
	if !a.LaneCompatibility.LinearLaneLocal || !a.LaneCompatibility.QuadraticLaneGlobal || !a.LaneCompatibility.MatchesLaneLocality || a.LaneCompatibility.NativeFunctorTheorem {
		t.Fatalf("bad lane compatibility: %s", FormatLaneCompatibility(a.LaneCompatibility))
	}
	if !a.BoundaryAugmentation.BothAugmentedByB2 || !a.BoundaryAugmentation.UniformAugmentation || a.BoundaryAugmentation.NativeNormalization {
		t.Fatalf("bad boundary augmentation: %s", FormatBoundaryAugmentation(a.BoundaryAugmentation))
	}
}

func TestGate917WrongDenominatorsDetectedButNotNativeProof(t *testing.T) {
	a, err := BuildDefault()
	if err != nil {
		t.Fatal(err)
	}
	if !a.Contamination.BareLinearMismatches || !a.Contamination.BareQuadraticMismatches || !a.Contamination.CommonDenominatorMismatches || a.Contamination.NativeProof {
		t.Fatalf("bad contamination audit: %s", FormatContamination(a.Contamination))
	}
	if near(a.Contamination.BareLinearAlpha, AlphaLinear) || near(a.Contamination.BareQuadraticAlpha, AlphaQuad) || near(a.Contamination.SameH72LinearAlpha, AlphaLinear) {
		t.Fatalf("wrong denominators unexpectedly matched: %s", FormatContamination(a.Contamination))
	}
}

func TestGate917FiveSubobjectReconstruction(t *testing.T) {
	a, err := BuildDefault()
	if err != nil {
		t.Fatal(err)
	}
	if !a.Reconstruction.ResponseShape || !a.Reconstruction.DegreeSelector || !a.Reconstruction.CrossLaneExclusion || !a.Reconstruction.SSplitTransport || !a.Reconstruction.ChamberNormalization || a.Reconstruction.NativeAlphaTheorem {
		t.Fatalf("bad reconstruction flags: %s", FormatReconstruction(a.Reconstruction))
	}
	if !near(a.Reconstruction.LinearContribution, AlphaLinear) || !near(a.Reconstruction.QuadraticContribution, AlphaQuad) || !near(a.Reconstruction.TotalAlpha, AlphaB) {
		t.Fatalf("bad alpha reconstruction: %s", FormatReconstruction(a.Reconstruction))
	}
}

func TestGate917Firewalls(t *testing.T) {
	a, err := BuildDefault()
	if err != nil {
		t.Fatal(err)
	}
	if !firewallsOK(a.Firewalls) {
		t.Fatalf("firewall leak: %s", FormatFirewalls(a.Firewalls))
	}
	for _, want := range []string{FailureH10NotNativeActivation, FailureH72NotNativeActivation, FailureLocalVsGlobalNotNative, FailureBoundaryAugmentationNotNative, FailureNumericalMismatchNotNative, FailureReconstructionNotNativeAlpha, FailureAlphaStillSealed, FailureNotNativeR3, FailureFullAFDescentStillBlocked, FailureNoGenerationCarrierMap, FailureNoNativeYukawaOperator} {
		if !containsAll(a.Firewalls.List(), []string{want}) {
			t.Fatalf("missing firewall %s", want)
		}
	}
}

func TestGate917Theorem(t *testing.T) {
	res := Generation2BoundaryAugmentedResponseChamberNormalizationAuditTheorem().Verify()
	if res.Status == "FAILED_ROUTE" {
		t.Fatalf("unexpected theorem failure: %+v", res)
	}
	for _, c := range res.Checks {
		if !c.Passed {
			t.Fatalf("failed check %s: %s", c.Name, c.Detail)
		}
	}
	joined := strings.Join(res.Notes, "\n")
	for _, want := range append(append(Statuses(), Supports()...), Failures()...) {
		if !strings.Contains(joined, want) {
			t.Fatalf("missing marker %s", want)
		}
	}
	for _, want := range []string{FinalTruth, Classification, ShortStatus, StrategicConclusion, NextGate} {
		if !strings.Contains(joined, want) {
			t.Fatalf("missing marker %s", want)
		}
	}
}
