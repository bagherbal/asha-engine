package generation2phaseanchoredneutralpunctureairlockfunctoraudit

import (
	"strings"
	"testing"
)

func TestGate901PhaseAnchorOrdersButDoesNotNativeSelect(t *testing.T) {
	a, err := BuildDefault()
	if err != nil {
		t.Fatal(err)
	}
	p := a.PhaseAnchor
	if !p.OrdersPair || p.SelectsNatively || p.ELambda != ELambda || p.EBarLambda != EBarLambda {
		t.Fatalf("bad phase anchor: %s", FormatPhaseAnchor(p))
	}
	if !containsAll(p.Failures, []string{FailureNoNativeRightPhaseOrientation, FailureNoNativeSelectionLambda}) {
		t.Fatalf("missing phase failures: %s", FormatPhaseAnchor(p))
	}
}

func TestGate901PhaseAnchorSelectsSocketEdgeWeakAndAlphaUnderSeal(t *testing.T) {
	a, err := BuildDefault()
	if err != nil {
		t.Fatal(err)
	}
	if !a.SocketOrder.OrderedByPhaseAnchor || a.SocketOrder.OrderedNatively || a.SocketOrder.Puncture != PPhase {
		t.Fatalf("socket order leak: %s", FormatSocketOrder(a.SocketOrder))
	}
	if !a.EdgeOrdering.GeneratedByPhaseAnchor || a.EdgeOrdering.GeneratedNatively {
		t.Fatalf("edge ordering leak: %s", FormatEdgeOrdering(a.EdgeOrdering))
	}
	if !a.WeakKernel.PhaseIndexedFrame || a.WeakKernel.SelectorNative || a.WeakKernel.Kernel != "h_lambda tensor P_1" {
		t.Fatalf("weak kernel leak: %s", FormatWeakKernel(a.WeakKernel))
	}
	if !a.BoundaryAlpha.SelectedByPhasePuncture || a.BoundaryAlpha.NativeAlphaFunctor || !near(a.BoundaryAlpha.Alpha, AlphaB) {
		t.Fatalf("alpha leak: %s", FormatBoundaryAlpha(a.BoundaryAlpha))
	}
}

func TestGate901MasterFunctorReduction(t *testing.T) {
	a, err := BuildDefault()
	if err != nil {
		t.Fatal(err)
	}
	m := a.MasterFunctor
	if !m.UnifiesSocketOrder || !m.UnifiesEdgeOrdering || !m.UnifiesWeakKernel || !m.UnifiesBoundaryAlpha || !m.SingleMasterBlocker || m.NativeFunctor {
		t.Fatalf("master functor mismatch: %s", FormatMasterFunctor(m))
	}
	if !containsAll(m.Supports, []string{SupportAirlockUnifiesAlphaAndHiggs, SupportR3SealedReducesToSingleFunctor}) {
		t.Fatalf("missing supports: %s", FormatMasterFunctor(m))
	}
}

func TestGate901StabilizerAndDiagnosticsFreeze(t *testing.T) {
	a, err := BuildDefault()
	if err != nil {
		t.Fatal(err)
	}
	if !a.Stabilizer.PhaseAnchored || a.Stabilizer.FullDescent {
		t.Fatalf("stabilizer leak: %s", FormatStabilizer(a.Stabilizer))
	}
	if !a.Diagnostics.Coherent || !a.Diagnostics.OfficialFrozen || a.Diagnostics.CanUpdate || !near(a.Diagnostics.NEff, OperatorNEffDiagnostic) || near(a.Diagnostics.NEff, a.Diagnostics.OfficialNEff) {
		t.Fatalf("diagnostic leak: %s", FormatDiagnostics(a.Diagnostics))
	}
}

func TestGate901Firewalls(t *testing.T) {
	a, err := BuildDefault()
	if err != nil {
		t.Fatal(err)
	}
	if !firewallsOK(a.Firewalls) {
		t.Fatalf("firewall leak: %s", FormatFirewalls(a.Firewalls))
	}
}

func TestGate901Theorem(t *testing.T) {
	res := Generation2PhaseAnchoredNeutralPunctureAirlockFunctorAuditTheorem().Verify()
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
}
