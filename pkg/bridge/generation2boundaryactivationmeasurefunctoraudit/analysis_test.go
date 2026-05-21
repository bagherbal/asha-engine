package generation2boundaryactivationmeasurefunctoraudit

import (
	"strings"
	"testing"
)

func TestGate920DomainAndDegreeExtraction(t *testing.T) {
	a, err := BuildDefault()
	if err != nil {
		t.Fatal(err)
	}
	if a.Domain.InheritedStatus != Gate919ShortStatus || a.Domain.IncludesLambda0 || !a.Domain.ReducedBasepoint || a.Domain.NativeTheorem {
		t.Fatalf("bad domain: %s", FormatDomain(a.Domain))
	}
	if len(a.Domain.NonzeroDegrees) != 2 || a.Domain.NonzeroDegrees[0] != 1 || a.Domain.NonzeroDegrees[1] != 2 {
		t.Fatalf("bad active degrees: %s", FormatDomain(a.Domain))
	}
	if a.DegreeExtraction.SeparateS2Transport || !a.DegreeExtraction.ExteriorGeneratedS2 || a.DegreeExtraction.NativeTheorem {
		t.Fatalf("bad degree extraction: %s", FormatDegreeExtraction(a.DegreeExtraction))
	}
	if !near(a.DegreeExtraction.DegreeCoefficients[1], SBoundary) || !near(a.DegreeExtraction.DegreeCoefficients[2], SBoundary*SBoundary) {
		t.Fatalf("bad coefficients: %s", FormatDegreeExtraction(a.DegreeExtraction))
	}
}

func TestGate920SelectorChambersAndCrossLanes(t *testing.T) {
	a, err := BuildDefault()
	if err != nil {
		t.Fatal(err)
	}
	if a.Selector.Targets[1] != "[F_1/F_0]_{Z2}" || a.Selector.Targets[2] != "[F_2/F_0]_{Z2}" || a.Selector.Ranks[1] != 3 || a.Selector.Ranks[2] != 7 || !a.Selector.RepresentativeIndependent || a.Selector.UniqueNativeSelector {
		t.Fatalf("bad selector: %s", FormatSelector(a.Selector))
	}
	if a.Chambers.Ranks[1] != 10 || a.Chambers.Ranks[2] != 72 || !a.Chambers.ExplicitLaneWeights || a.Chambers.NativeTheorem {
		t.Fatalf("bad chambers: %s", FormatChambers(a.Chambers))
	}
	if !a.CrossLanes.ExcludedIfFunctional || !a.CrossLanes.AbsorbedInIndexing || a.CrossLanes.FunctionhoodNative {
		t.Fatalf("bad cross-lane audit: %s", FormatCrossLanes(a.CrossLanes))
	}
}

func TestGate920AlphaReconstruction(t *testing.T) {
	a, err := BuildDefault()
	if err != nil {
		t.Fatal(err)
	}
	if a.Alpha.Formula != MeasureFormula || a.Alpha.BoundaryFormula != BoundaryAlphaMeasureFormula || !a.Alpha.ReassemblesFive || a.Alpha.NativeAlpha {
		t.Fatalf("bad alpha flags: %s", FormatAlpha(a.Alpha))
	}
	if !near(a.Alpha.LinearContribution, AlphaLinear) || !near(a.Alpha.QuadraticContribution, AlphaQuad) || !near(a.Alpha.Alpha, AlphaB) {
		t.Fatalf("bad alpha reconstruction: %s", FormatAlpha(a.Alpha))
	}
}

func TestGate920NativeFirewalls(t *testing.T) {
	a, err := BuildDefault()
	if err != nil {
		t.Fatal(err)
	}
	if !a.NativeStatus.BridgeMeasureCandidate || a.NativeStatus.NativeMeasure || a.NativeStatus.NativeAlpha || a.NativeStatus.NativeR3 || len(a.NativeStatus.MissingNativeTheorems) != 6 {
		t.Fatalf("bad native status: %s", FormatNativeStatus(a.NativeStatus))
	}
	if !firewallsOK(a.Firewalls) {
		t.Fatalf("firewall leak: %s", FormatFirewalls(a.Firewalls))
	}
	for _, want := range []string{FailureNoNativeBoundaryActivationMeasure, FailureMuBFormalNotNative, FailureNoNativeMeasureUniqueness, FailureNoNativeMuBDomain, FailureNoNativeDegreeExtraction, FailureNoNativeSSplitTransportMap, FailureNoNativeDegreeToZ2FlagFunctor, FailureNoNativeChamberNormalization, FailureNoNativeZ2CrossLaneExclusion, FailureAlphaBridgeCandidateNotNative, FailureNotNativeR3, FailureFullAFDescentStillBlocked, FailureNoGenerationCarrierMap, FailureNoNativeYukawaOperator} {
		if !containsAll(a.Firewalls.List(), []string{want}) {
			t.Fatalf("missing firewall %s", want)
		}
	}
}

func TestGate920Theorem(t *testing.T) {
	res := Generation2BoundaryActivationMeasureFunctorAuditTheorem().Verify()
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
	for _, want := range []string{FinalTruth, Classification, ShortStatus, StrategicConclusion, NextGate, BoundaryActivationMeasure, BoundaryResponseFunctor, MeasureFormula, BranchMeasureFormula, BoundaryAlphaMeasureFormula} {
		if !strings.Contains(joined, want) {
			t.Fatalf("missing marker %s", want)
		}
	}
}
