package generation2dualsealr3candidatejoppositefulldescentincidencefunctorrecheckaudit

import (
	"strings"
	"testing"
)

func TestGate890InheritedLedger(t *testing.T) {
	a, err := BuildDefault()
	if err != nil {
		t.Fatal(err)
	}
	if !a.Ledger.R3DualSealCandidate || a.Ledger.NativeR3 || a.Ledger.Rank != RankHRMin || len(a.Ledger.Atoms) != 3 {
		t.Fatalf("bad inherited ledger: %s", FormatLedger(a.Ledger))
	}
	if !near(a.Ledger.OperatorNEff, OperatorNEffDiagnostic) || !near(a.Ledger.OperatorCYukawa, OperatorCYukawaDiagnostic) {
		t.Fatalf("diagnostic drift: %s", FormatLedger(a.Ledger))
	}
}

func TestGate890JMirrorExtension(t *testing.T) {
	a, err := BuildDefault()
	if err != nil {
		t.Fatal(err)
	}
	j := a.JExtension
	if !j.MirrorExists || !j.PreservesRanks || !j.PreservesOrthogonality || j.SourceRank != RankHRMin || j.MirrorRank != RankHRMin || j.CombinedActiveRank != RankActiveMirrorLedger {
		t.Fatalf("bad J mirror: %s", FormatJExtension(j))
	}
	if j.CompletesFullHFMin || j.CombinedActiveRank == j.HFMinRank || j.OperatorLevelJFKO || j.FullJOppositeAction || j.NativeFiniteSector || j.YukawaMagnitudeSource {
		t.Fatalf("J extension promoted incorrectly: %s", FormatJExtension(j))
	}
}

func TestGate890DescentAudit(t *testing.T) {
	a, err := BuildDefault()
	if err != nil {
		t.Fatal(err)
	}
	d := a.Descent
	if !d.AFOrientLedgerStable || d.FullAFLedgerStable || !d.FullToOrientHiggsRestriction || d.NativeDescentCertified || d.WeakSocketFrameFullHInvariant || d.NativeFiniteSectorProjectors {
		t.Fatalf("bad descent audit: %s", FormatDescent(d))
	}
	if !containsAll(d.Failures, []string{FailureAFOrientNotFullAF, FailureSocketProjectorsNotStableFullH, FailureNoNativeDescentFullToOrient}) {
		t.Fatalf("missing descent failures: %s", FormatDescent(d))
	}
}

func TestGate890IncidenceRecheck(t *testing.T) {
	a, err := BuildDefault()
	if err != nil {
		t.Fatal(err)
	}
	i := a.Incidence
	if !i.BoundaryAlphaSealCoherent || !i.AlphaReconstructedUnderSeal || i.NewNativeBoundarySourceFound || i.NativeFunctorCertified || i.CrossLaneExclusionCertified || !i.AlphaStillSealed {
		t.Fatalf("bad incidence recheck: %s", FormatIncidence(i))
	}
	if !containsAll(i.Failures, []string{FailureGate890AddsNoBoundarySource, FailureNoNativeIncidenceFunctor, FailureNoNativeCrossLaneExclusion, FailureAlphaStillSealed}) {
		t.Fatalf("missing incidence failures: %s", FormatIncidence(i))
	}
}

func TestGate890FirewallsAndFreeze(t *testing.T) {
	a, err := BuildDefault()
	if err != nil {
		t.Fatal(err)
	}
	if !a.Freeze.Frozen || !a.Freeze.DiagnosticOnly || a.Freeze.CanUpdate || near(a.Freeze.OperatorNEff, a.Freeze.OfficialNEff) {
		t.Fatalf("freeze leak: %s", FormatFreeze(a.Freeze))
	}
	if hasPhysicalLeak(a) {
		t.Fatalf("physical leak: %s / %s", FormatLedger(a.Ledger), FormatJExtension(a.JExtension))
	}
	if !firewallsOK(a.Firewalls) {
		t.Fatalf("firewall leak: %s", FormatFirewalls(a.Firewalls))
	}
}

func TestGate890Theorem(t *testing.T) {
	res := Generation2DualSealR3CandidateJOppositeFullDescentIncidenceFunctorRecheckAuditTheorem().Verify()
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
