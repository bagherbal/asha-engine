package generation2finiteoneformpuncturekernelweaksocketselectoraudit

import (
	"strings"
	"testing"
)

func TestGate893NullEdgePointsToHPlusButDoesNotDerive(t *testing.T) {
	a, err := BuildDefault()
	if err != nil {
		t.Fatal(err)
	}
	n := a.NullEdge
	if !n.YPlus1Zero || !n.SelectsHPlusCandidate || n.NativeOrientationTheorem {
		t.Fatalf("bad null-edge audit: %s", FormatNullEdge(n))
	}
	if !containsAll(n.Failures, []string{FailureNullEdgeNotNativeOrientation, FailureNoNativeHiggsOrientationSource}) {
		t.Fatalf("missing null-edge firewalls: %s", FormatNullEdge(n))
	}
}

func TestGate893FiniteOneFormIsCompatibleButCircular(t *testing.T) {
	a, err := BuildDefault()
	if err != nil {
		t.Fatal(err)
	}
	o := a.OneForm
	if !o.UsesWeakFrame || !o.CompatibleWithFrame || o.ForcesFrame || !o.CircularIfUsedAsSource {
		t.Fatalf("bad finite one-form route: %s", FormatFiniteOneForm(o))
	}
	if !containsAll(o.Supports, []string{SupportFiniteOneFormCompatibleWithOrientation}) || !containsAll(o.Failures, []string{FailureDFPatternRestatesOrientation}) {
		t.Fatalf("missing one-form statuses: %s", FormatFiniteOneForm(o))
	}
}

func TestGate893NonCircularityRequiresWeakSocketSelector(t *testing.T) {
	a, err := BuildDefault()
	if err != nil {
		t.Fatal(err)
	}
	n := a.NonCircularity
	if !n.RequiresIndependentSelector || n.DFToOrientationCertified || n.NativeSelectorFunctional {
		t.Fatalf("noncircularity leak: %s", FormatNonCircularity(n))
	}
	if n.MissingObject != "WeakSocketSelectorFunctional" || n.CandidatePrinciple != "MinimalNullEdgeOrientationPrinciple" {
		t.Fatalf("wrong missing object: %s", FormatNonCircularity(n))
	}
}

func TestGate893WeakFrameAndFreeze(t *testing.T) {
	a, err := BuildDefault()
	if err != nil {
		t.Fatal(err)
	}
	if !a.WeakFrame.FullHMixesFrame || !a.WeakFrame.StabilizerPreservesFrame || a.WeakFrame.NativeOrientationSource {
		t.Fatalf("weak frame leak: %s", FormatWeakFrame(a.WeakFrame))
	}
	if !a.Freeze.Frozen || !a.Freeze.DiagnosticOnly || a.Freeze.CanUpdate || near(a.Freeze.OperatorNEff, a.Freeze.OfficialNEff) {
		t.Fatalf("freeze leak: %s", FormatFreeze(a.Freeze))
	}
}

func TestGate893Firewalls(t *testing.T) {
	a, err := BuildDefault()
	if err != nil {
		t.Fatal(err)
	}
	if !firewallsOK(a.Firewalls) {
		t.Fatalf("firewall leak: %s", FormatFirewalls(a.Firewalls))
	}
}

func TestGate893Theorem(t *testing.T) {
	res := Generation2FiniteOneFormPunctureKernelWeakSocketSelectorAuditTheorem().Verify()
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
