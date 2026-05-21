package generation2totalrelativetracemagnitudeoperatorandledgerconsistencyaudit

import (
	"math"
	"strings"
	"testing"
)

func TestGate829TotalOperatorTraceAndSquareTrace(t *testing.T) {
	a, err := BuildDefault()
	if err != nil {
		t.Fatal(err)
	}
	alpha := a.Ledger.AlphaB
	if len(a.Operator.TotalSpectrum) != 7 || len(a.Operator.TopBlock) != 3 || len(a.Operator.RestBlock) != 4 {
		t.Fatalf("bad operator dimensions: %s", FormatOperator(a.Operator))
	}
	if math.Abs(a.Operator.TraceTotal-(3+3*alpha)) > 1e-15 {
		t.Fatalf("bad total trace: %s", FormatOperator(a.Operator))
	}
	wantSquare := 3 + 3*alpha*alpha - 6*math.Pow(alpha, 3) + 12*math.Pow(alpha, 4)
	if math.Abs(a.Operator.SquareTraceTotal-wantSquare) > 1e-15 {
		t.Fatalf("bad total square trace: %s", FormatOperator(a.Operator))
	}
	if math.Abs(a.Operator.RestTrace-3*alpha) > 1e-18 {
		t.Fatalf("bad rest trace: %s", FormatOperator(a.Operator))
	}
}

func TestGate829OperatorNEffAndLedgerSeparation(t *testing.T) {
	a, err := BuildDefault()
	if err != nil {
		t.Fatal(err)
	}
	if math.Abs(a.Ledger.OperatorNEff-3.0023273750818085) > 1e-15 {
		t.Fatalf("bad operator N_eff: %s", FormatLedger(a.Ledger))
	}
	if math.Abs(a.Ledger.BFNTruncatedNEff-3.0023273750818085) > 1e-15 {
		t.Fatalf("bad BFN truncated N_eff: %s", FormatLedger(a.Ledger))
	}
	if math.Abs(a.Ledger.OfficialNEff-3.0023273474722147) > 1e-15 {
		t.Fatalf("bad official N_eff: %s", FormatLedger(a.Ledger))
	}
	if math.Abs(a.Ledger.OperatorMinusBFN-a.Consistency.FifthOrderResidual) > 1e-15 {
		t.Fatalf("bad fifth-order residual: %s", FormatConsistency(a.Consistency))
	}
	if math.Abs(a.Ledger.OperatorMinusOfficial) < 1e-10 {
		t.Fatalf("operator and official ledger silently collapsed: %s", FormatLedger(a.Ledger))
	}
	if !a.Ledger.OfficialFrozen || !a.Ledger.DiagnosticSeparated || !a.Ledger.AliasBlocked {
		t.Fatalf("ledger separation flags failed: %s", FormatLedger(a.Ledger))
	}
}

func TestGate829FreezeAndTheorem(t *testing.T) {
	a, err := BuildDefault()
	if err != nil {
		t.Fatal(err)
	}
	if a.Consistency.CanPromoteOperatorToOfficial || a.Consistency.CanUpdateNEff || a.Consistency.CanUpdateCYukawa || a.Consistency.CanUpdateCHiggs {
		t.Fatalf("freeze failed: %s", FormatConsistency(a.Consistency))
	}
	if a.Source.AlphaSourceCertified || a.Source.BoundaryAlphaTransportMapCertified || a.Source.SectorLedgerCertified || a.Source.NativeYukawaTheoremCertified {
		t.Fatalf("source firewall failed: %s", FormatSource(a.Source))
	}
	res := Generation2TotalRelativeTraceMagnitudeOperatorAndLedgerConsistencyAuditTheorem().Verify()
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
