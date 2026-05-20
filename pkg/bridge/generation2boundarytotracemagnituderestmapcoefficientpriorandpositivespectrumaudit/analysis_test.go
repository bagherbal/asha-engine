package generation2boundarytotracemagnituderestmapcoefficientpriorandpositivespectrumaudit

import (
	"math"
	"strings"
	"testing"
)

func TestGate816NumericalLedgerAndCoefficientPriors(t *testing.T) {
	a, err := BuildDefault()
	if err != nil {
		t.Fatal(err)
	}
	if math.Abs(a.Ledger.DeltaNBFN-0.002327375081808316) > 1e-18 {
		t.Fatalf("bad BFN delta: %s", FormatLedger(a.Ledger))
	}
	if math.Abs(a.Ledger.C2Obs-5.8299915725) > 1e-9 {
		t.Fatalf("bad c2 obs: %s", FormatLedger(a.Ledger))
	}
	if !strings.Contains(a.Coeff95.Factorization, "3/5") || !strings.Contains(a.Coeff6.Factorization, "2") {
		t.Fatalf("bad coefficient priors: %s %s", FormatCoefficient(a.Coeff95), FormatCoefficient(a.Coeff6))
	}
	if !containsAll(a.Coeff95.Failures, []string{StatusCoeff95NotTheorem}) || !containsAll(a.Coeff6.Failures, []string{StatusCoeff6NotTheorem}) {
		t.Fatalf("missing coefficient firewalls")
	}
}

func TestGate816AlphaAndDeltaTables(t *testing.T) {
	a, err := BuildDefault()
	if err != nil {
		t.Fatal(err)
	}
	if !containsAlpha(a.AlphaRows, "1/2", false, false) {
		t.Fatalf("half M2 should fail positivity: %s", FormatAlphaRows(a.AlphaRows))
	}
	if !containsAlpha(a.AlphaRows, "3/5", true, true) || !containsAlpha(a.AlphaRows, "6/11", true, true) || !containsAlpha(a.AlphaRows, "1", true, true) {
		t.Fatalf("expected positive-compatible alpha rows: %s", FormatAlphaRows(a.AlphaRows))
	}
	if len(a.DeltaRows) != 24 {
		t.Fatalf("expected 24 delta reconstructions, got %d", len(a.DeltaRows))
	}
	if math.Abs(a.DeltaRows[0].ResidualToBFN) > 1e-8 {
		t.Fatalf("best reconstruction not close enough: %s", FormatDeltaRows(a.DeltaRows, 3))
	}
}

func TestGate816NoGoStatusImpactAndTheorem(t *testing.T) {
	a, err := BuildDefault()
	if err != nil {
		t.Fatal(err)
	}
	if a.NoGo.ConstructsAlphaBetaQ {
		t.Fatalf("coefficient package must not construct map")
	}
	if a.Positive.Level == "R2" || a.Positive.Level == "R3" || a.Positive.Level == "R4" {
		t.Fatalf("status over-promoted: %s", a.Positive.Level)
	}
	if math.Abs(a.Impact.CYukawaBFN-0.9992248096922658) > 1e-15 || math.Abs(a.Impact.OfficialCYukawa-CYukawa) > 1e-15 {
		t.Fatalf("bad impact: %s", FormatImpact(a.Impact))
	}
	res := Generation2BoundaryToTraceMagnitudeRestMapCoefficientPriorAndPositiveSpectrumAuditTheorem().Verify()
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
