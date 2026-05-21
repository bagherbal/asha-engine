package generation2boundaryactivationmeasurenaturalityanduniquenessaudit

import (
	"strings"
	"testing"
)

func TestGate921DomainBasepointAndDegreeNaturality(t *testing.T) {
	a, err := BuildDefault()
	if err != nil {
		t.Fatal(err)
	}
	if a.Domain.InheritedStatus != Gate920ShortStatus || a.Domain.ArbitraryPolynomial || !a.Domain.NaturalOnReducedB2 || a.Domain.NativeTheorem {
		t.Fatalf("bad domain: %s", FormatDomain(a.Domain))
	}
	if !a.Basepoint.UnreducedHasLambda0 || !a.Basepoint.ReducedResponseUsed || !a.Basepoint.NoConstantTermForced || !a.Basepoint.UniqueIfNoConstantAlpha || a.Basepoint.NativeTheorem {
		t.Fatalf("bad basepoint: %s", FormatBasepoint(a.Basepoint))
	}
	if a.Degree.DegreePowers[1] != 1 || a.Degree.DegreePowers[2] != 2 || !a.Degree.PowerAssignmentUnique || a.Degree.NativeTheorem {
		t.Fatalf("bad degree naturality: %s", FormatDegree(a.Degree))
	}
	if !near(a.Degree.Coefficients[1], SBoundary) || !near(a.Degree.Coefficients[2], SBoundary*SBoundary) {
		t.Fatalf("bad degree coefficients: %s", FormatDegree(a.Degree))
	}
}

func TestGate921SelectorNormalizationZ2AndAlternatives(t *testing.T) {
	a, err := BuildDefault()
	if err != nil {
		t.Fatal(err)
	}
	if !a.Selector.FunctionhoodAssumed || !a.Selector.UniquePerDegree || !a.Selector.CrossLanesExcluded || a.Selector.NativeTheorem {
		t.Fatalf("bad selector: %s", FormatSelector(a.Selector))
	}
	if !a.Normalization.LaneLocalityAccepted || !a.Normalization.UniqueGivenLocalGlobal || a.Normalization.Ranks[1] != 10 || a.Normalization.Ranks[2] != 72 || a.Normalization.NativeTheorem {
		t.Fatalf("bad normalization: %s", FormatNormalization(a.Normalization))
	}
	if !a.Z2.RepresentativesExchanged || !a.Z2.RanksInvariant || !a.Z2.MeasureInvariant || a.Z2.NativeTheorem {
		t.Fatalf("bad z2 independence: %s", FormatZ2(a.Z2))
	}
	if !a.Alternatives.UnreducedRejected || !a.Alternatives.CrossLaneRejected || !a.Alternatives.BareChamberRejected || !a.Alternatives.CommonDenominatorRejected || !a.Alternatives.UniqueAmongTested || a.Alternatives.FullNativeUniqueness {
		t.Fatalf("bad alternatives: %s", FormatAlternatives(a.Alternatives))
	}
	if !nearLoose(a.Alternatives.PollutedLinearWeight, float64(143)/360) || !nearLoose(a.Alternatives.BareLinearWeight, float64(3)/8) || !nearLoose(a.Alternatives.BareQuadraticWeight, float64(7)/70) {
		t.Fatalf("bad alternative weights: %s", FormatAlternatives(a.Alternatives))
	}
}

func TestGate921AlphaAndNativeFirewalls(t *testing.T) {
	a, err := BuildDefault()
	if err != nil {
		t.Fatal(err)
	}
	if a.Alpha.Formula != MeasureFormula || a.Alpha.BoundaryFormula != BoundaryAlphaFormula || a.Alpha.NativeAlpha {
		t.Fatalf("bad alpha metadata: %s", FormatAlpha(a.Alpha))
	}
	if !near(a.Alpha.LinearContribution, AlphaLinear) || !near(a.Alpha.QuadraticContribution, AlphaQuad) || !near(a.Alpha.Alpha, AlphaB) {
		t.Fatalf("bad alpha reconstruction: %s", FormatAlpha(a.Alpha))
	}
	if !a.NativeStatus.NaturalMeasureCandidate || !a.NativeStatus.UniqueUnderConstraints || a.NativeStatus.NativeMeasure || a.NativeStatus.NativeAlpha || a.NativeStatus.NativeR3 {
		t.Fatalf("bad native status: %s", FormatNativeStatus(a.NativeStatus))
	}
	if !firewallsOK(a.Firewalls) {
		t.Fatalf("firewall leak: %s", FormatFirewalls(a.Firewalls))
	}
	for _, want := range []string{FailureNoNativeBoundaryActivationMeasure, FailureNoNativeMeasureUniqueness, FailureDomainNaturalityNotNative, FailureNoNativeBasepointReduction, FailureNoNativeDegreeRespectingMeasure, FailureNoNativeUniqueSelector, FailureNoNativeSelectorFunctionhood, FailureNoNativeLaneLocalityToChamber, FailureNoNativeChamberNormalization, FailureZ2InvarianceNotNative, FailureAlternativeRejectionNotFullNative, FailureAlphaBridgeCandidateNotNative, FailureNotNativeR3} {
		if !containsAll(a.Firewalls.List(), []string{want}) {
			t.Fatalf("missing firewall %s", want)
		}
	}
}

func TestGate921Theorem(t *testing.T) {
	res := Generation2BoundaryActivationMeasureNaturalityAndUniquenessAuditTheorem().Verify()
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
	for _, want := range []string{FinalTruth, Classification, ShortStatus, StrategicConclusion, NextGate, BoundaryMeasureObject, MeasureFormula, BranchMeasureFormula, BoundaryAlphaFormula} {
		if !strings.Contains(joined, want) {
			t.Fatalf("missing marker %s", want)
		}
	}
}
