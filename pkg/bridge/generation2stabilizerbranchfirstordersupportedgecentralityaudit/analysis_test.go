package generation2stabilizerbranchfirstordersupportedgecentralityaudit

import (
	"strings"
	"testing"
)

func TestGate859FirstOrderSupportNotOperatorTheorem(t *testing.T) {
	a, err := BuildDefault()
	if err != nil {
		t.Fatal(err)
	}
	if !a.FirstOrder.OrderZeroInherited || !a.FirstOrder.SupportAuditable || !a.FirstOrder.DRhoCommutatorAllowedOneFormSource || !a.FirstOrder.FirstOrderSupportConditionAudited || a.FirstOrder.OperatorTheoremCertified {
		t.Fatalf("bad first-order support state: %s", FormatFirstOrder(a.FirstOrder))
	}
	if !containsAll(a.FirstOrder.Failures, []string{FailureNoFullOperatorFirstOrder, FailureNoCompleteJOppositeProof}) {
		t.Fatalf("missing first-order firewalls: %+v", a.FirstOrder.Failures)
	}
}

func TestGate859ColorEdgesAreCentralOnP3(t *testing.T) {
	a, err := BuildDefault()
	if err != nil {
		t.Fatal(err)
	}
	if !colorEdgesCentral(a.Edges) {
		t.Fatalf("color edge centrality not certified: %s", FormatEdges(a.Edges))
	}
	for _, e := range a.Edges {
		if e.ColorEdge && e.RequiredForm != "y_+3 I_{P_3}" && e.RequiredForm != "y_-3 I_{P_3}" {
			t.Fatalf("bad color form: %+v", e)
		}
	}
}

func TestGate859LeptonEdgeAndPuncture(t *testing.T) {
	a, err := BuildDefault()
	if err != nil {
		t.Fatal(err)
	}
	if !leptonAndPunctureOK(a.Edges) {
		t.Fatalf("bad lepton/puncture state: %s", FormatEdges(a.Edges))
	}
	if !a.PunctureKernel.PunctureCoefficientZero || a.PunctureKernel.PunctureReintroduced || !a.PunctureKernel.LeftKernelPresent {
		t.Fatalf("bad puncture kernel: %s", FormatPunctureKernel(a.PunctureKernel))
	}
}

func TestGate859NoMagnitudeOrLedgerPromotion(t *testing.T) {
	a, err := BuildDefault()
	if err != nil {
		t.Fatal(err)
	}
	for _, e := range a.Edges {
		if e.YukawaMagnitude || e.OperatorIntertwinerCertified {
			t.Fatalf("edge overpromoted: %+v", e)
		}
	}
	if !a.Ledger.OfficialFrozen || a.Ledger.AlphaNative || a.Ledger.R3 || a.Ledger.R4 || a.Impact.CanUpdateNEff || a.Impact.CanUpdateCYukawa || a.Impact.CanUpdateCHiggs {
		t.Fatalf("ledger overpromoted: %s | %s", FormatLedger(a.Ledger), FormatImpact(a.Impact))
	}
}

func TestGate859Theorem(t *testing.T) {
	res := Generation2StabilizerBranchFirstOrderSupportEdgeCentralityAuditTheorem().Verify()
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
