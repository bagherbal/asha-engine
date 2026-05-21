package generation2puncturedsocketresponsefunctionalaudit

import (
	"strings"
	"testing"
)

func TestGate846ResponseTable(t *testing.T) {
	a, err := BuildDefault()
	if err != nil {
		t.Fatal(err)
	}
	if !a.Table.ReconstructsGate845 || a.Table.NativeSource || a.Table.ActiveRank != 7 || a.Table.PunctureRank != 1 || a.Table.RightFullRank != 8 {
		t.Fatalf("bad table support: %s", FormatTable(a.Table))
	}
	if !a.Table.TopColor.Included || a.Table.TopColor.Rank != 3 || !nearly(a.Table.TopColor.Eigenvalue, 1) {
		t.Fatalf("bad top cell: %s", FormatCell(a.Table.TopColor))
	}
	if a.Table.Puncture.Included || a.Table.Puncture.Rank != 1 || a.Table.Puncture.HasEdgeTarget {
		t.Fatalf("bad puncture cell: %s", FormatCell(a.Table.Puncture))
	}
	if !a.Table.RestColor.Included || a.Table.RestColor.Rank != 3 || !nearly(a.Table.RestColor.Eigenvalue, AlphaB*(1-AlphaB)) {
		t.Fatalf("bad rest color cell: %s", FormatCell(a.Table.RestColor))
	}
	if !a.Table.RestLepton.Included || a.Table.RestLepton.Rank != 1 || !nearly(a.Table.RestLepton.Eigenvalue, 3*AlphaB*AlphaB) {
		t.Fatalf("bad rest lepton cell: %s", FormatCell(a.Table.RestLepton))
	}
}

func TestGate846FunctionalAndTraceDiagnostics(t *testing.T) {
	a, err := BuildDefault()
	if err != nil {
		t.Fatal(err)
	}
	if !a.Functional.ReconstructsTable || a.Functional.NativeFunctional || a.Functional.AlphaDerived || a.Functional.PunctureDerived || !a.Functional.BLTransferTraceZero {
		t.Fatalf("bad functional: %s", FormatFunctional(a.Functional))
	}
	if !nearly(a.Table.Trace, 3+3*AlphaB) {
		t.Fatalf("bad trace: %s", FormatTable(a.Table))
	}
	wantSquare := 3 + 3*AlphaB*AlphaB - 6*AlphaB*AlphaB*AlphaB + 12*AlphaB*AlphaB*AlphaB*AlphaB
	if !nearly(a.Table.SquareTrace, wantSquare) || !nearly(a.Table.OperatorNEff, operatorNEff(AlphaB)) {
		t.Fatalf("bad square/N_eff: %s", FormatTable(a.Table))
	}
	if a.Ledger.OperatorNEff == a.Ledger.OfficialNEff {
		t.Fatalf("operator and official ledgers aliased: %s", FormatLedger(a.Ledger))
	}
}

func TestGate846EdgeCompatibilityAndFirewalls(t *testing.T) {
	a, err := BuildDefault()
	if err != nil {
		t.Fatal(err)
	}
	if !a.Edges.CompatibleWithGate844 || !a.Edges.SupportOnly || !a.Edges.ActiveCellsHaveTargets || a.Edges.PunctureHasTarget || a.Edges.ExplicitDFMatrix || a.Edges.FirstOrderCertified || a.Edges.BimoduleCommutantCertified || a.Edges.Magnitudes {
		t.Fatalf("bad edge compatibility: %s", FormatEdges(a.Edges))
	}
	if !a.Firewalls.Enforced || !a.Firewalls.FunctionalSealNotNative || !a.Firewalls.NoNativeCompressionFunctional || !a.Firewalls.AlphaStillSealed || !a.Firewalls.PunctureStillSealed || !a.Firewalls.NoTraceMagnitudeReadout || !a.Firewalls.NoNEffUpdate || !a.Firewalls.NotR3 || !a.Firewalls.NotR4 || a.Firewalls.Verdict != StatusFirewallGate846 {
		t.Fatalf("firewalls invalid: %+v", a.Firewalls)
	}
}

func TestGate846Theorem(t *testing.T) {
	res := Generation2PuncturedSocketResponseFunctionalAuditTheorem().Verify()
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
