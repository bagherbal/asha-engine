package generation2fullafdescenthiggsorientationsourceobstructionaudit

import (
	"strings"
	"testing"
)

func TestGate891ProjectorFullAFInstability(t *testing.T) {
	a, err := BuildDefault()
	if err != nil {
		t.Fatal(err)
	}
	if !projectorsOK(a.Projectors) {
		t.Fatalf("bad projectors: %s", FormatProjectors(a.Projectors))
	}
	for _, p := range a.Projectors {
		if !p.StableUnderAFOrient || p.StableUnderFullAF || p.StableUnderFullH || p.PhysicalSector || p.GenerationResolved || p.FlavorResolved || p.IndividualYukawaValue {
			t.Fatalf("projector promoted incorrectly: %s", FormatProjector(p))
		}
	}
}

func TestGate891StabilizerNotFullH(t *testing.T) {
	a, err := BuildDefault()
	if err != nil {
		t.Fatal(err)
	}
	s := a.Stabilizer
	if !s.PreservesHPlusHMinus || !s.PreservesProjectors || s.IsFullH || s.IsFullAF || s.NativeDescentCertified {
		t.Fatalf("bad stabilizer audit: %s", FormatStabilizer(s))
	}
	if !containsAll(s.Failures, []string{FailureStabilizerNotFullNativeAF, FailureAFOrientNotFullAF, FailureNoNativeDescentFullToOrient}) {
		t.Fatalf("missing stabilizer failures: %s", FormatStabilizer(s))
	}
}

func TestGate891OrientationSourcesAuditedNoNativeSource(t *testing.T) {
	a, err := BuildDefault()
	if err != nil {
		t.Fatal(err)
	}
	o := a.OrientationSources
	if len(o.Candidates) < 7 || o.AnyNativeSourceCertified || !o.RequiresOrientationSeal {
		t.Fatalf("bad orientation source audit: %s", FormatOrientationSources(o))
	}
	for _, c := range o.Candidates {
		if !c.Audited || c.Certified {
			t.Fatalf("candidate promoted incorrectly: %s", FormatCandidate(c))
		}
	}
}

func TestGate891FullDescentBlocked(t *testing.T) {
	a, err := BuildDefault()
	if err != nil {
		t.Fatal(err)
	}
	d := a.Descent
	if !d.AFOrientStable || d.FullAFStable || !d.GenericHMixesWeakSockets || d.WeakFrameFullHInvariant || d.NativeDescentCertified {
		t.Fatalf("descent leak: %s", FormatDescent(d))
	}
	if !containsAll(d.Failures, []string{FailureFullHMixesWeakSockets, FailureSocketProjectorsNotStableFullH, FailureNoNativeDescentFullToOrient}) {
		t.Fatalf("missing descent failures: %s", FormatDescent(d))
	}
}

func TestGate891FreezeAndFirewalls(t *testing.T) {
	a, err := BuildDefault()
	if err != nil {
		t.Fatal(err)
	}
	if !a.Freeze.Frozen || !a.Freeze.DiagnosticOnly || a.Freeze.CanUpdate || near(a.Freeze.OperatorNEff, a.Freeze.OfficialNEff) {
		t.Fatalf("freeze leak: %s", FormatFreeze(a.Freeze))
	}
	if hasPhysicalLeak(a) {
		t.Fatalf("physical leak: %s", FormatProjectors(a.Projectors))
	}
	if !firewallsOK(a.Firewalls) {
		t.Fatalf("firewall leak: %s", FormatFirewalls(a.Firewalls))
	}
}

func TestGate891Theorem(t *testing.T) {
	res := Generation2FullAFDescentHiggsOrientationSourceObstructionAuditTheorem().Verify()
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
