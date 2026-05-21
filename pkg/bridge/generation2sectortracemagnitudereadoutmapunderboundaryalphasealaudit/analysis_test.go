package generation2sectortracemagnitudereadoutmapunderboundaryalphasealaudit

import (
	"strings"
	"testing"
)

func TestGate884ReadoutRows(t *testing.T) {
	a, err := BuildDefault()
	if err != nil {
		t.Fatal(err)
	}
	if len(a.Readout.Rows) != 3 {
		t.Fatalf("expected three readout rows: %s", FormatReadout(a.Readout))
	}
	if a.Readout.Rows[0].Atom != AtomPiPlus3 || a.Readout.Rows[1].Atom != AtomPiMinus3 || a.Readout.Rows[2].Atom != AtomPiMinus1 {
		t.Fatalf("bad row order: %s", FormatRows(a.Readout.Rows))
	}
	if a.Readout.ActiveRank != RankHRMin || !a.Readout.Orthogonal || !a.Readout.CompleteOnHRMin {
		t.Fatalf("bad completeness: %s", FormatReadout(a.Readout))
	}
}

func TestGate884PositiveTraceMagnitudeLedger(t *testing.T) {
	a, err := BuildDefault()
	if err != nil {
		t.Fatal(err)
	}
	for _, row := range a.Readout.Rows {
		if !row.Positive || row.Weight <= 0 || row.TraceContribution <= 0 || row.SquareContribution <= 0 {
			t.Fatalf("bad positive readout row: %s", FormatRow(row))
		}
	}
	if !near(a.Readout.TraceTotal, 3+3*AlphaB) {
		t.Fatalf("trace mismatch: %s", FormatReadout(a.Readout))
	}
	if !near(a.Readout.SquareTrace, 3+3*AlphaB*AlphaB-6*AlphaB*AlphaB*AlphaB+12*AlphaB*AlphaB*AlphaB*AlphaB) {
		t.Fatalf("square trace mismatch: %s", FormatReadout(a.Readout))
	}
	if !near(a.Readout.OperatorNEff, OperatorNEffDiagnostic) || !near(a.Readout.OperatorCYukawa, OperatorCYukawaDiagnostic) {
		t.Fatalf("operator readout mismatch: %s", FormatReadout(a.Readout))
	}
}

func TestGate884ClassificationAndFirewalls(t *testing.T) {
	a, err := BuildDefault()
	if err != nil {
		t.Fatal(err)
	}
	if a.Readout.NativeR3 || a.Readout.OfficialUpdatesAllowed || a.Readout.Classification != Classification {
		t.Fatalf("bad classification: %s", FormatReadout(a.Readout))
	}
	if !containsAll(a.Readout.Failures, []string{FailureAlphaStillSealed, FailureReadoutUnderSealNotNative, FailureSocketAtomsNotPhysical, FailureNoGenerationCarrierMap, FailureNoFlavorOrientationMap, FailureNoIndividualYukawaValues}) {
		t.Fatalf("missing readout failures: %s", FormatReadout(a.Readout))
	}
	if !firewallsOK(a.Firewalls) {
		t.Fatalf("bad firewalls: %s", FormatFirewalls(a.Firewalls))
	}
}

func TestGate884OfficialFreeze(t *testing.T) {
	a, err := BuildDefault()
	if err != nil {
		t.Fatal(err)
	}
	if !a.Freeze.Frozen || !a.Freeze.DiagnosticOnly || a.Freeze.CanUpdate {
		t.Fatalf("official freeze leaked: %s", FormatFreeze(a.Freeze))
	}
	if near(a.Freeze.OperatorNEff, a.Freeze.OfficialNEff) || near(a.Freeze.OperatorCYukawa, a.Freeze.OfficialCYukawa) {
		t.Fatalf("operator and official ledgers collapsed: %s", FormatFreeze(a.Freeze))
	}
}

func TestGate884Theorem(t *testing.T) {
	res := Generation2SectorTraceMagnitudeReadoutMapUnderBoundaryAlphaSealAuditTheorem().Verify()
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
