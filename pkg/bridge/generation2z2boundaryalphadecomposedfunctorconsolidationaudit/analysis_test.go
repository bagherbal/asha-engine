package generation2z2boundaryalphadecomposedfunctorconsolidationaudit

import (
	"strings"
	"testing"
)

func TestGate918FiveSubobjectsInherited(t *testing.T) {
	a, err := BuildDefault()
	if err != nil {
		t.Fatal(err)
	}
	if !a.Subobjects.AllAuditedAtShape || a.Subobjects.NativeTheorem {
		t.Fatalf("bad subobject stack: %s", FormatSubobjects(a.Subobjects))
	}
	for _, want := range []string{Gate913ShortStatus, Gate914ShortStatus, Gate915ShortStatus, Gate916ShortStatus, Gate917ShortStatus} {
		if !strings.Contains(FormatSubobjects(a.Subobjects), want) {
			t.Fatalf("missing inherited status %s in %s", want, FormatSubobjects(a.Subobjects))
		}
	}
}

func TestGate918BoundaryAlphaReassemblesBridgeCandidate(t *testing.T) {
	a, err := BuildDefault()
	if err != nil {
		t.Fatal(err)
	}
	if !a.BoundaryAlpha.InternalCoherent || a.BoundaryAlpha.OpaqueSeal || !a.BoundaryAlpha.BridgeCandidate || a.BoundaryAlpha.NativeTheorem {
		t.Fatalf("bad BoundaryAlpha status: %s", FormatBoundaryAlpha(a.BoundaryAlpha))
	}
	if !near(a.BoundaryAlpha.LinearContribution, AlphaLinear) || !near(a.BoundaryAlpha.QuadraticContribution, AlphaQuad) || !near(a.BoundaryAlpha.TotalAlpha, AlphaB) {
		t.Fatalf("bad BoundaryAlpha reconstruction: %s", FormatBoundaryAlpha(a.BoundaryAlpha))
	}
}

func TestGate918RepresentativeIndependence(t *testing.T) {
	a, err := BuildDefault()
	if err != nil {
		t.Fatal(err)
	}
	if a.RepresentativeIndependence.RankPair != [2]int{3, 7} || !a.RepresentativeIndependence.TauPhiPreservesRankPair || a.RepresentativeIndependence.PhaseSignEntersAlpha || !a.RepresentativeIndependence.CorrectAlphaDomain || a.RepresentativeIndependence.NativeAirlockFunctor {
		t.Fatalf("bad representative independence: %s", FormatRepresentativeIndependence(a.RepresentativeIndependence))
	}
}

func TestGate918BridgeCandidateObligations(t *testing.T) {
	a, err := BuildDefault()
	if err != nil {
		t.Fatal(err)
	}
	if !a.BridgeCandidate.AllVisibleComponents || a.BridgeCandidate.NativeTheorem || len(a.BridgeCandidate.TheoremObligations) != 5 {
		t.Fatalf("bad bridge candidate: %s", FormatBridgeCandidate(a.BridgeCandidate))
	}
	for _, want := range []string{"native reduced B2 response functional", "native degree-indexed Z2 flag functor", "native Z2 cross-lane exclusion theorem", "native S_split transport map", "native response-chamber normalization theorem"} {
		if !containsAll(a.BridgeCandidate.TheoremObligations, []string{want}) {
			t.Fatalf("missing obligation %s: %s", want, FormatBridgeCandidate(a.BridgeCandidate))
		}
	}
}

func TestGate918R3TraceLedgerDiagnosticOnly(t *testing.T) {
	a, err := BuildDefault()
	if err != nil {
		t.Fatal(err)
	}
	if len(a.R3TraceLedger.TraceRows) != 3 || !a.R3TraceLedger.DiagnosticOnly || a.R3TraceLedger.OfficialUpdateAllowed {
		t.Fatalf("bad R3 trace ledger flags: %s", FormatR3TraceLedger(a.R3TraceLedger))
	}
	if !nearLoose(a.R3TraceLedger.NEffOperator, NEffOperator) || !nearLoose(a.R3TraceLedger.CYukawaOperator, CYukawaOperator) || !nearLoose(a.R3TraceLedger.CHiggsOperator, CHiggsOperator) {
		t.Fatalf("bad operator values: %s", FormatR3TraceLedger(a.R3TraceLedger))
	}
}

func TestGate918NativeGapsAndFirewalls(t *testing.T) {
	a, err := BuildDefault()
	if err != nil {
		t.Fatal(err)
	}
	if len(a.NativeGaps.AlphaGaps) != 5 || len(a.NativeGaps.FiniteLayerGaps) != 2 || !a.NativeGaps.GenerationFlavorR4OrLater || a.NativeGaps.NativeR3 {
		t.Fatalf("bad native gaps: %s", FormatNativeGaps(a.NativeGaps))
	}
	if !firewallsOK(a.Firewalls) {
		t.Fatalf("firewall leak: %s", FormatFirewalls(a.Firewalls))
	}
	for _, want := range []string{FailureNotNativeR3, FailureNoNativeReducedB2ResponseFunctional, FailureNoNativeDegreeToZ2FlagClassFunctor, FailureNoNativeZ2CrossLaneExclusionTheorem, FailureNoNativeSSplitTransportMap, FailureNoNativeResponseChamberNormalization, FailureAlphaBridgeCandidateNotNative, FailureFullAFDescentStillBlocked, FailureNoOfficialNEffUpdateAllowed} {
		if !containsAll(a.Firewalls.List(), []string{want}) {
			t.Fatalf("missing firewall %s", want)
		}
	}
}

func TestGate918Theorem(t *testing.T) {
	res := Generation2Z2BoundaryAlphaDecomposedFunctorConsolidationAuditTheorem().Verify()
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
