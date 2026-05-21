package generation2finitebodyaggregatetracecompressionshadowmapaudit

import (
	"strings"
	"testing"
)

func TestGate845DomainSplitAndPuncture(t *testing.T) {
	a, err := BuildDefault()
	if err != nil {
		t.Fatal(err)
	}
	if !a.Domain.InheritedGate844 || !a.Domain.Orthogonal || !a.Domain.CompleteOnHRMin || a.Domain.HRMinRank != 7 || a.Domain.Top.Rank != 3 || a.Domain.Rest.Rank != 4 {
		t.Fatalf("bad domain split: %s", FormatDomain(a.Domain))
	}
	if !a.Domain.PunctureExcluded || a.Domain.Puncture.Included || a.Domain.Puncture.Rank != 1 {
		t.Fatalf("bad puncture status: %s", FormatDomain(a.Domain))
	}
	if a.Domain.Top.BMinusLTrace != 1 || a.Domain.Rest.BMinusLTrace != 0 || a.Domain.Puncture.BMinusLTrace != -1 {
		t.Fatalf("bad B-L traces: %s", FormatDomain(a.Domain))
	}
}

func TestGate845AggregateTraceReconstruction(t *testing.T) {
	a, err := BuildDefault()
	if err != nil {
		t.Fatal(err)
	}
	if !a.Operator.FiniteBodyLocationAtSealLevel || a.Operator.NativeCompressionTheorem || a.Operator.NativeTraceCompressionFunctional {
		t.Fatalf("bad operator placement: %s", FormatOperator(a.Operator))
	}
	if !nearly(a.Operator.TopTrace, 3) || !nearly(a.Operator.RestTrace, 3*AlphaB) || !nearly(a.Operator.TotalTrace, 3+3*AlphaB) {
		t.Fatalf("bad trace reconstruction: %s", FormatOperator(a.Operator))
	}
	wantSquare := 3 + 3*AlphaB*AlphaB - 6*AlphaB*AlphaB*AlphaB + 12*AlphaB*AlphaB*AlphaB*AlphaB
	if !nearly(a.Operator.TotalSquareTrace, wantSquare) || !nearly(a.Operator.OperatorNEff, operatorNEff(AlphaB)) {
		t.Fatalf("bad square/N_eff reconstruction: %s", FormatOperator(a.Operator))
	}
	if a.Operator.OperatorNEff == a.Operator.OfficialNEff {
		t.Fatalf("operator and official ledgers were aliased: %s", FormatOperator(a.Operator))
	}
}

func TestGate845EdgeCompatibilityAndFirewalls(t *testing.T) {
	a, err := BuildDefault()
	if err != nil {
		t.Fatal(err)
	}
	if !a.Edges.CompatibleWithGate844 || !a.Edges.SupportOnly || a.Edges.ExplicitDFMatrix || a.Edges.FirstOrderCertified || a.Edges.BimoduleCommutantCertified || a.Edges.Magnitudes {
		t.Fatalf("bad edge compatibility: %s", FormatEdges(a.Edges))
	}
	if a.Operator.AlphaDerived || a.Operator.TraceMagnitudeReadout || a.Operator.R3 || a.Operator.R4 {
		t.Fatalf("operator over-promoted: %s", FormatOperator(a.Operator))
	}
	if !a.Firewalls.Enforced || !a.Firewalls.NoNativeCompressionMap || !a.Firewalls.AlphaStillSealed || !a.Firewalls.NoTraceMagnitudeReadout || !a.Firewalls.NoNEffUpdate || !a.Firewalls.NotR3 || !a.Firewalls.NotR4 || a.Firewalls.Verdict != StatusFirewallGate845 {
		t.Fatalf("firewalls invalid: %+v", a.Firewalls)
	}
}

func TestGate845Theorem(t *testing.T) {
	res := Generation2FiniteBodyAggregateTraceCompressionShadowMapAuditTheorem().Verify()
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
