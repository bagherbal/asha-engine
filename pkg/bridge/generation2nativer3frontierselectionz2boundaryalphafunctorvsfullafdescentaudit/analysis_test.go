package generation2nativer3frontierselectionz2boundaryalphafunctorvsfullafdescentaudit

import (
	"strings"
	"testing"
)

func TestGate911InheritedPlateauAndDiagnostics(t *testing.T) {
	a, err := BuildDefault()
	if err != nil {
		t.Fatal(err)
	}
	if !a.Inherited.R3TraceLedgerSealed || a.Inherited.NativeR3 || a.Inherited.LoopBackPhase || a.Inherited.LoopBackRepAlpha || a.Inherited.LoopBackSocketOrder {
		t.Fatalf("bad inherited plateau: %s", FormatInherited(a.Inherited))
	}
	if !a.Trace.DiagnosticsOnly || a.Trace.OfficialUpdates || !near(a.Trace.Alpha, AlphaB) || !near(a.Trace.OperatorNEff, OperatorNEffDiagnostic) || !near(a.Trace.OperatorCYukawa, OperatorCYukawaDiagnostic) || !near(a.Trace.OperatorCHiggs, OperatorCHiggsDiagnostic) {
		t.Fatalf("bad trace diagnostics: %s", FormatTrace(a.Trace))
	}
}

func TestGate911FrontierASelectedFirst(t *testing.T) {
	a, err := BuildDefault()
	if err != nil {
		t.Fatal(err)
	}
	f := a.FrontierA
	if !f.SelectedFirst || !f.RequiredForNativeR3 || !f.DirectlyControlsAlpha || !f.DirectlyControlsTraceWeights || !f.DirectlyControlsNEff || !f.DirectlyControlsCYukawa || !f.DirectlyControlsCHiggs || f.Deferred || f.R4OrLater {
		t.Fatalf("bad Frontier A selection: %s", FormatFrontier(f))
	}
	for _, want := range []string{MissingZ2BoundaryAlphaFunctor, "degree-to-Z2-flag-class functor", "native Z2 cross-lane exclusion theorem"} {
		found := false
		for _, got := range f.MissingObjects {
			if got == want {
				found = true
			}
		}
		if !found {
			t.Fatalf("missing Frontier A object %s in %v", want, f.MissingObjects)
		}
	}
}

func TestGate911FrontierBAndCDeferred(t *testing.T) {
	a, err := BuildDefault()
	if err != nil {
		t.Fatal(err)
	}
	b := a.FrontierB
	if b.SelectedFirst || !b.RequiredForNativeR3 || !b.FullAFDescentProblem || !b.Deferred || b.DirectlyControlsAlpha || b.R4OrLater {
		t.Fatalf("bad Frontier B classification: %s", FormatFrontier(b))
	}
	c := a.FrontierC
	if c.SelectedFirst || c.RequiredForNativeR3 || !c.GenerationFlavorYukawaBranch || !c.R4OrLater || !c.Deferred {
		t.Fatalf("bad Frontier C classification: %s", FormatFrontier(c))
	}
}

func TestGate911SelectionAndFirewalls(t *testing.T) {
	a, err := BuildDefault()
	if err != nil {
		t.Fatal(err)
	}
	s := a.Selection
	if !s.AttackAFirst || s.AttackBBeforeA || s.EnterGenerationNow || s.NativeR3 || s.LoopBackToPhase || s.LoopBackToRepAlpha || s.LoopBackToSocket || s.UpdateOfficialLedger {
		t.Fatalf("bad selection: %s", FormatSelection(s))
	}
	if !firewallsOK(a.Firewalls) {
		t.Fatalf("firewall leak: %s", FormatFirewalls(a.Firewalls))
	}
	for _, want := range []string{FailureNotNativeR3, FailureNoNativeZ2BoundaryAlphaFunctor, FailureFullAFDescentStillBlocked, FailureNoGenerationCarrierMap, FailureNoNativeYukawaOperator} {
		if !containsAll(a.FirewallsList(), []string{want}) {
			t.Fatalf("missing firewall %s", want)
		}
	}
}

func TestGate911Theorem(t *testing.T) {
	res := Generation2NativeR3FrontierSelectionZ2BoundaryAlphaFunctorVsFullAFDescentAuditTheorem().Verify()
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
			t.Fatalf("missing final marker %s in notes: %s", want, joined)
		}
	}
}
