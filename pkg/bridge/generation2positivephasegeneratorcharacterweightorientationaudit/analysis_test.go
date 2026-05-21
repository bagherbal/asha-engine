package generation2positivephasegeneratorcharacterweightorientationaudit

import (
	"strings"
	"testing"
)

func TestGate906QPhiWeightOperator(t *testing.T) {
	a, err := BuildDefault()
	if err != nil {
		t.Fatal(err)
	}
	q := a.WeightOperator
	if !q.Exists || !q.OrderEquivalentToSign || q.PositiveSignNative || q.LambdaWeight != 1 || q.BarLambdaWeight != -1 {
		t.Fatalf("bad Q_phi audit: %s", FormatWeightOperator(q))
	}
}

func TestGate906HopfAndCL17AreCandidatesNotTheorems(t *testing.T) {
	a, err := BuildDefault()
	if err != nil {
		t.Fatal(err)
	}
	if !a.HopfReeb.SuppliesPositiveGeneratorCandidate || !a.HopfReeb.SelectsLambdaWeightIfSealed || a.HopfReeb.TypedActionOnCR2 || a.HopfReeb.NativeSelector {
		t.Fatalf("bad Hopf audit: %s", FormatHopfReeb(a.HopfReeb))
	}
	if !a.CL17Sign.SignMatchesPhaseWeight || !a.CL17Sign.CanSourceQPhiIfTyped || a.CL17Sign.TypedMapToQPhi || a.CL17Sign.SelectsSocketOrder {
		t.Fatalf("bad Cl17 audit: %s", FormatCL17Sign(a.CL17Sign))
	}
}

func TestGate906JAndBoundaryDoNotOrient(t *testing.T) {
	a, err := BuildDefault()
	if err != nil {
		t.Fatal(err)
	}
	if !a.JConjugation.ExchangesWeights || !a.JConjugation.ExplainsPair || a.JConjugation.OrientsSign {
		t.Fatalf("bad J audit: %s", FormatJConjugation(a.JConjugation))
	}
	if !a.BoundaryOrientation.OrientsExteriorDegree || a.BoundaryOrientation.SelectsPhaseWeightSign {
		t.Fatalf("bad boundary audit: %s", FormatBoundaryOrientation(a.BoundaryOrientation))
	}
}

func TestGate906WoundAndFreeze(t *testing.T) {
	a, err := BuildDefault()
	if err != nil {
		t.Fatal(err)
	}
	if a.Wound.NativeSolved || !strings.Contains(a.Wound.ReducedWound, "positive phase generator") {
		t.Fatalf("bad wound: %s", FormatWound(a.Wound))
	}
	if !a.Freeze.Frozen || !a.Freeze.DiagnosticOnly || a.Freeze.CanUpdate || !near(a.Freeze.OperatorNEff, OperatorNEffDiagnostic) || near(a.Freeze.OperatorNEff, a.Freeze.OfficialNEff) {
		t.Fatalf("bad freeze: %s", FormatFreeze(a.Freeze))
	}
}

func TestGate906Firewalls(t *testing.T) {
	a, err := BuildDefault()
	if err != nil {
		t.Fatal(err)
	}
	if !firewallsOK(a.Firewalls) {
		t.Fatalf("firewall leak: %s", FormatFirewalls(a.Firewalls))
	}
}

func TestGate906Theorem(t *testing.T) {
	res := Generation2PositivePhaseGeneratorCharacterWeightOrientationAuditTheorem().Verify()
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
