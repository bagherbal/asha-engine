package generation2nativereducedboundarypairresponsefunctionalaudit

import (
	"strings"
	"testing"
)

func TestGate913InheritedSubobjectSelection(t *testing.T) {
	a, err := BuildDefault()
	if err != nil {
		t.Fatal(err)
	}
	if a.Inherited.SubobjectIndex != 1 || a.Inherited.TotalSubobjects != 5 || a.Inherited.SelectedSubobject != NativeReducedB2Theorem {
		t.Fatalf("bad inherited selection: %s", FormatInherited(a.Inherited))
	}
	if a.Inherited.ReopensPhaseSign || a.Inherited.ReopensSocketOrder || a.Inherited.ReopensRepresentative || a.Inherited.DerivesAlpha || a.Inherited.UpdatesOfficialLedger {
		t.Fatalf("closed wound reopened: %s", FormatInherited(a.Inherited))
	}
}

func TestGate913ExteriorLedgerAndExpansion(t *testing.T) {
	a, err := BuildDefault()
	if err != nil {
		t.Fatal(err)
	}
	if a.Ledger.Rank != BoundaryPairRank || a.Ledger.Lambda0Dim != 1 || a.Ledger.Lambda1Dim != 2 || a.Ledger.Lambda2Dim != 1 || a.Ledger.Lambda3Dim != 0 || !a.Ledger.Lambda3Zero {
		t.Fatalf("bad exterior ledger: %s", FormatLedger(a.Ledger))
	}
	if !a.Expansion.ExactShape || !a.Expansion.ConstantRemoved || a.Expansion.NativeFunctional || hasDegree(a.Expansion.ReducedTerms, 0) {
		t.Fatalf("bad expansion status: %s", FormatExpansion(a.Expansion))
	}
	if len(a.Expansion.DegreeOneTerms) != 2 || len(a.Expansion.DegreeTwoTerms) != 1 {
		t.Fatalf("bad response term count: %s", FormatExpansion(a.Expansion))
	}
	if !near(sumDegree(a.Expansion.ReducedTerms, 1), 2*SBoundary) || !near(sumDegree(a.Expansion.ReducedTerms, 2), SBoundary*SBoundary) {
		t.Fatalf("bad response coefficients: %s", FormatExpansion(a.Expansion))
	}
}

func TestGate913ZeroOrderTruncationAndNaturalityFirewalls(t *testing.T) {
	a, err := BuildDefault()
	if err != nil {
		t.Fatal(err)
	}
	if !a.ZeroOrder.RemovedByReduction || !a.ZeroOrder.ReducedStartsAtOrderOne || a.ZeroOrder.NativeReasonForReduction {
		t.Fatalf("bad zero-order status: %s", FormatZeroOrder(a.ZeroOrder))
	}
	if !a.Truncation.Lambda3Zero || !a.Truncation.NoCubicOrHigher || a.Truncation.HighestNonzeroDegree != 2 {
		t.Fatalf("bad truncation: %s", FormatTruncation(a.Truncation))
	}
	if !a.Naturality.MultiplicativeCandidate || a.Naturality.NativeASHAFunctional || a.Naturality.VariationalPrinciple || a.Naturality.FunctorialSelectionPrinciple {
		t.Fatalf("bad naturality classification: %s", FormatNaturality(a.Naturality))
	}
}

func TestGate913AlphaShapeAndSsplitTransportRemainBlocked(t *testing.T) {
	a, err := BuildDefault()
	if err != nil {
		t.Fatal(err)
	}
	if !a.AlphaShape.SuppliesPowerShape || a.AlphaShape.SelectsZ2FlagTargets || a.AlphaShape.DerivesCoefficients || a.AlphaShape.ProvesCrossLaneExclusion || a.AlphaShape.DerivesAlpha {
		t.Fatalf("alpha shape firewall leaked: %s", FormatAlphaShape(a.AlphaShape))
	}
	if !a.Transport.UsesSAsParameter || a.Transport.NativeTransport || a.Transport.TypedParameterMap {
		t.Fatalf("S_split transport firewall leaked: %s", FormatTransport(a.Transport))
	}
	for _, want := range []string{FailureNoZ2FlagTargets, FailureNoAlphaCoefficients, FailureNoCrossLaneExclusion, FailureReducedResponseNotAlphaAlone, FailureAlphaStillSealed} {
		if !containsAll(a.AlphaShape.Failures, []string{want}) {
			t.Fatalf("missing alpha firewall %s from %s", want, FormatAlphaShape(a.AlphaShape))
		}
	}
}

func TestGate913GlobalFirewalls(t *testing.T) {
	a, err := BuildDefault()
	if err != nil {
		t.Fatal(err)
	}
	if !firewallsOK(a.Firewalls) {
		t.Fatalf("firewall leak: %s", FormatFirewalls(a.Firewalls))
	}
	for _, want := range []string{FailureReducedB2NotNativeFunctional, FailureNoNativeReasonEBMinusOne, FailureMultiplicativeNotNative, FailureNoVariationalProductForm, FailureNoNativeTransportSIntoB2, FailureNoZ2FlagTargets, FailureNoCrossLaneExclusion, FailureAlphaStillSealed, FailureNotNativeR3, FailureFullAFDescentStillBlocked, FailureNoGenerationCarrierMap, FailureNoNativeYukawaOperator} {
		if !containsAll(a.FirewallsList(), []string{want}) {
			t.Fatalf("missing firewall %s from %s", want, FormatFirewalls(a.Firewalls))
		}
	}
}

func TestGate913Theorem(t *testing.T) {
	res := Generation2NativeReducedBoundaryPairResponseFunctionalAuditTheorem().Verify()
	if res.Status == "FAILED_ROUTE" {
		t.Fatalf("unexpected theorem failure: %+v", res)
	}
	for _, c := range res.Checks {
		if !c.Passed {
			t.Fatalf("failed check %s: %s", c.Name, c.Detail)
		}
	}
	joined := strings.Join(res.Notes, "\n")
	for _, want := range Statuses() {
		if !strings.Contains(joined, want) {
			t.Fatalf("missing status %s", want)
		}
	}
	for _, want := range []string{FinalTruth, Classification, ShortStatus, StrategicConclusion, NextGate} {
		if !strings.Contains(joined, want) {
			t.Fatalf("missing marker %s in notes: %s", want, joined)
		}
	}
}
