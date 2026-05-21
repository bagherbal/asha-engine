package generation2higgsorientationsourcecandidateweaksocketselectoraudit

import (
	"strings"
	"testing"
)

func TestGate892WeakSelectorStillRequiresOrientationSeal(t *testing.T) {
	a, err := BuildDefault()
	if err != nil {
		t.Fatal(err)
	}
	w := a.WeakSelector
	if !w.FullHMixesWeakSockets || !w.StabilizerPreservesFrame || w.NativeOrientationSource {
		t.Fatalf("bad weak selector: %s", FormatWeakSelector(w))
	}
	if !containsAll(w.Failures, []string{FailureFullHMixesWeakSockets, FailureNoNativeHiggsOrientationSource}) {
		t.Fatalf("missing weak selector failures: %s", FormatWeakSelector(w))
	}
}

func TestGate892StrongestCandidatesButNoNativeSource(t *testing.T) {
	a, err := BuildDefault()
	if err != nil {
		t.Fatal(err)
	}
	c := a.Candidates
	if len(c.Candidates) < 7 || c.AnyNativeSourceCertified || !c.RequiresOrientationSeal {
		t.Fatalf("bad candidate audit: %s", FormatCandidateAudit(c))
	}
	if !containsAll(c.Supports, []string{SupportFiniteOneFormStrongestCandidate, SupportPunctureKernelPointsToHPlus}) {
		t.Fatalf("missing strongest supports: %s", FormatCandidateAudit(c))
	}
	for _, candidate := range c.Candidates {
		if !candidate.Audited || candidate.Certified {
			t.Fatalf("candidate promoted incorrectly: %s", FormatCandidate(candidate))
		}
	}
}

func TestGate892DFSupportCircularity(t *testing.T) {
	a, err := BuildDefault()
	if err != nil {
		t.Fatal(err)
	}
	e := a.EdgeOrientation
	if !e.AssumesWeakFrame || e.DerivesWeakFrame || !e.CircularIfUsedAsSource || !e.CompatibleWithFrame {
		t.Fatalf("bad edge orientation audit: %s", FormatEdgeOrientation(e))
	}
}

func TestGate892PunctureKernelAndBLDoNotPromote(t *testing.T) {
	a, err := BuildDefault()
	if err != nil {
		t.Fatal(err)
	}
	if !a.PunctureKernel.PointsToHPlus || a.PunctureKernel.NativeSourceCertified {
		t.Fatalf("puncture/kernel leak: %s", FormatPunctureKernel(a.PunctureKernel))
	}
	if !a.BMinusL.Compatible || a.BMinusL.SelectsWeakFrame || a.BMinusL.FullRectangleTrace != 0 {
		t.Fatalf("B-L leak: %s", FormatBMinusL(a.BMinusL))
	}
}

func TestGate892FreezeAndFirewalls(t *testing.T) {
	a, err := BuildDefault()
	if err != nil {
		t.Fatal(err)
	}
	if !a.Freeze.Frozen || !a.Freeze.DiagnosticOnly || a.Freeze.CanUpdate || near(a.Freeze.OperatorNEff, a.Freeze.OfficialNEff) {
		t.Fatalf("freeze leak: %s", FormatFreeze(a.Freeze))
	}
	if !firewallsOK(a.Firewalls) {
		t.Fatalf("firewall leak: %s", FormatFirewalls(a.Firewalls))
	}
}

func TestGate892Theorem(t *testing.T) {
	res := Generation2HiggsOrientationSourceCandidateWeakSocketSelectorAuditTheorem().Verify()
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
