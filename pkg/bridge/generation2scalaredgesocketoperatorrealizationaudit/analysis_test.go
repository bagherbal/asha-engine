package generation2scalaredgesocketoperatorrealizationaudit

import (
	"strings"
	"testing"
)

func TestGate860OperatorValuedYMap(t *testing.T) {
	a, err := BuildDefault()
	if err != nil {
		t.Fatal(err)
	}
	if !a.Y.OperatorValued || !a.Y.SymbolicSocketMatrix || !a.Y.ColorCentrality || !a.Y.LeptonTriviality || !a.Y.PunctureZero || a.Y.RankIfActiveSocketsNonzero != YRankFull {
		t.Fatalf("bad Y operator state: %s", FormatY(a.Y))
	}
	if !containsAll(a.Y.Failures, []string{FailureOperatorYSymbolicNotYukawa, FailureNoNumericalYukawa, FailureZeroSocketBranchEnlargesKer}) {
		t.Fatalf("missing Y firewalls: %+v", a.Y.Failures)
	}
}

func TestGate860ColorCentralityOperators(t *testing.T) {
	a, err := BuildDefault()
	if err != nil {
		t.Fatal(err)
	}
	if !colorEdgesScalar(a.Edges) {
		t.Fatalf("color edges not scalar operators: %s", FormatEdges(a.Edges))
	}
}

func TestGate860LeptonAndPunctureOperators(t *testing.T) {
	a, err := BuildDefault()
	if err != nil {
		t.Fatal(err)
	}
	if !leptonAndPunctureOK(a.Edges) {
		t.Fatalf("lepton/puncture operators invalid: %s", FormatEdges(a.Edges))
	}
}

func TestGate860RankKernelLedger(t *testing.T) {
	a, err := BuildDefault()
	if err != nil {
		t.Fatal(err)
	}
	if !a.D.BuiltFromY || !a.D.SelfAdjointByBlockForm || a.D.RankIfActiveSocketsNonzero != DSymRankFull || a.D.KernelRankIfNonzero != KernelRank || a.D.KernelSingleton != "h_+ tensor P_1" {
		t.Fatalf("bad D ledger: %s", FormatD(a.D))
	}
}

func TestGate860NoMagnitudeOrLedgerPromotion(t *testing.T) {
	a, err := BuildDefault()
	if err != nil {
		t.Fatal(err)
	}
	for _, e := range a.Edges {
		if e.NumericalValue || e.YukawaMagnitude {
			t.Fatalf("edge overpromoted: %+v", e)
		}
	}
	if !a.Ledger.OfficialFrozen || a.Ledger.AlphaNative || a.Ledger.R3 || a.Ledger.R4 || a.Impact.CanUpdateNEff || a.Impact.CanUpdateCYukawa || a.Impact.CanUpdateCHiggs || a.Impact.CanPromoteToR3 || a.Impact.CanPromoteToR4 {
		t.Fatalf("ledger overpromoted: %s | %s", FormatLedger(a.Ledger), FormatImpact(a.Impact))
	}
}

func TestGate860Theorem(t *testing.T) {
	res := Generation2ScalarEdgeSocketOperatorRealizationAuditTheorem().Verify()
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
