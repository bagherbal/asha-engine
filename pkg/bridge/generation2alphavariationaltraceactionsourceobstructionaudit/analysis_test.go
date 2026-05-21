package generation2alphavariationaltraceactionsourceobstructionaudit

import (
	"math"
	"strings"
	"testing"
)

func TestGate830TraceExpansionReconstructsAlphaButIsRestatement(t *testing.T) {
	a, err := BuildDefault()
	if err != nil {
		t.Fatal(err)
	}
	if math.Abs(a.Ledger.AlphaB-0.0003878958469680527) > 1e-18 {
		t.Fatalf("bad alpha: %s", FormatLedger(a.Ledger))
	}
	if !a.Trace.ReconstructsAlpha || math.Abs(a.Trace.AlphaFromTrace-a.Ledger.AlphaB) > 1e-18 {
		t.Fatalf("trace expansion failed: %s", FormatTrace(a.Trace))
	}
	if !a.Trace.ClassifiedAsRestatement || a.Trace.X1NaturallyProduced || a.Trace.X2NaturallyProduced || a.Trace.TraceActionCertified {
		t.Fatalf("trace expansion over-promoted: %s", FormatTrace(a.Trace))
	}
	if !containsAll(a.Trace.Failures, []string{FailureTraceExpansionRestatesRule, FailureX1NotNative, FailureX2NotNative}) {
		t.Fatalf("missing trace failures: %s", strings.Join(a.Trace.Failures, ","))
	}
}

func TestGate830ResponseOrderAndVariationalObstruction(t *testing.T) {
	a, err := BuildDefault()
	if err != nil {
		t.Fatal(err)
	}
	if a.ResponseOrder.LinearPower != 1 || a.ResponseOrder.QuadraticPower != 2 {
		t.Fatalf("bad powers: %s", FormatResponseOrder(a.ResponseOrder))
	}
	if a.ResponseOrder.LinearOrderDerived || a.ResponseOrder.QuadraticOrderDerived || a.ResponseOrder.ResponseOrderTheoremCertified {
		t.Fatalf("response order over-promoted: %s", FormatResponseOrder(a.ResponseOrder))
	}
	if !a.Variational.StationarityWorksFormally || math.Abs(a.Variational.StationaryAlpha-a.Ledger.AlphaB) > 1e-18 {
		t.Fatalf("formal stationarity failed: %s", FormatVariational(a.Variational))
	}
	if a.Variational.ActionNative || !a.Variational.UsesInsertedAlphaRule || !a.Variational.IsFormalRepackaging || a.Variational.CertifiesAlphaTheorem {
		t.Fatalf("variational route over-promoted: %s", FormatVariational(a.Variational))
	}
}

func TestGate830FreezeAndTheorem(t *testing.T) {
	a, err := BuildDefault()
	if err != nil {
		t.Fatal(err)
	}
	if a.Impact.CanPromoteAlpha || a.Impact.CanPromoteOperatorNEff || a.Impact.CanUpdateNEff || a.Impact.CanUpdateCYukawa || a.Impact.CanUpdateCHiggs {
		t.Fatalf("impact freeze failed: %s", FormatImpact(a.Impact))
	}
	if !a.Firewalls.Enforced || !a.Firewalls.AlphaSealed || !a.Firewalls.NoBoundaryAlphaMap || !a.Firewalls.NotR3 || !a.Firewalls.NotR4 {
		t.Fatalf("firewall failed: %s", a.Firewalls.Verdict)
	}
	res := Generation2AlphaVariationalTraceActionSourceObstructionAuditTheorem().Verify()
	if res.Status == "FAILED_ROUTE" {
		t.Fatalf("unexpected theorem construction failure: %+v", res)
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
