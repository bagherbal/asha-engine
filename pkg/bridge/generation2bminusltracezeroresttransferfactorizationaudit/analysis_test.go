package generation2bminusltracezeroresttransferfactorizationaudit

import (
	"math"
	"strings"
	"testing"
)

func TestGate826BMinusLTransferAlgebra(t *testing.T) {
	a, err := BuildDefault()
	if err != nil {
		t.Fatal(err)
	}
	if a.Projectors.P1Rank != 1 || a.Projectors.P3Rank != 3 || !a.Projectors.Orthogonal || !a.Projectors.Complete {
		t.Fatalf("bad projector data: %s", FormatProjectors(a.Projectors))
	}
	if math.Abs(a.Projectors.TraceQ) > 1e-15 || math.Abs(a.Projectors.TraceP3Q+3) > 1e-15 || math.Abs(a.Projectors.TraceQ2-12) > 1e-15 {
		t.Fatalf("bad Q_BL traces: %s", FormatProjectors(a.Projectors))
	}
}

func TestGate826RestOperatorFactorization(t *testing.T) {
	a, err := BuildDefault()
	if err != nil {
		t.Fatal(err)
	}
	alpha := a.Ledger.AlphaB
	if a.Factorization.MaxAbsResidual > 1e-18 {
		t.Fatalf("factorization residual too large: %s", FormatFactorization(a.Factorization))
	}
	if math.Abs(a.Factorization.TraceQuadratic) > 1e-21 || math.Abs(a.Factorization.TraceRest-3*alpha) > 1e-18 {
		t.Fatalf("trace preservation failed: %s", FormatFactorization(a.Factorization))
	}
	wantSquare := 3*alpha*alpha - 6*math.Pow(alpha, 3) + 12*math.Pow(alpha, 4)
	if math.Abs(a.Factorization.SquareTrace-wantSquare) > 1e-21 {
		t.Fatalf("square trace failed: %s", FormatFactorization(a.Factorization))
	}
}

func TestGate826FirewallsAndTheorem(t *testing.T) {
	a, err := BuildDefault()
	if err != nil {
		t.Fatal(err)
	}
	if !a.Boundary.CertifiedTransferFactorization || a.Boundary.CertifiedAlphaSource || a.Boundary.CertifiedSectorLedger {
		t.Fatalf("boundary separation failed: %s", FormatBoundary(a.Boundary))
	}
	if a.Impact.CanUpdateCYukawa || a.Impact.CanUpdateCHiggs {
		t.Fatalf("impact firewall failed: %s", FormatImpact(a.Impact))
	}
	res := Generation2BMinusLTraceZeroRestTransferFactorizationAuditTheorem().Verify()
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
