package generation2jfoppositeactionorderzerobimodulerealizationaudit

import (
	"strings"
	"testing"
)

func TestGate858OppositeActionSeal(t *testing.T) {
	a, err := BuildDefault()
	if err != nil {
		t.Fatal(err)
	}
	if !a.Opposite.FormalJExchangeDefined || !a.Opposite.OppositeSupportDefined || a.Opposite.OppositeOperatorCertified || !a.Opposite.OrderZeroTargetTyped {
		t.Fatalf("bad opposite action: %s", FormatOpposite(a.Opposite))
	}
	if !containsAll(a.Opposite.Failures, []string{FailureJOppositeSealOnly, FailureNoOperatorJOppositeProof}) {
		t.Fatalf("missing opposite firewalls: %+v", a.Opposite.Failures)
	}
}

func TestGate858OrderZeroSupportNotOperatorTheorem(t *testing.T) {
	a, err := BuildDefault()
	if err != nil {
		t.Fatal(err)
	}
	if !a.OrderZero.SupportAuditable || !a.OrderZero.BlockSupportCompatible || a.OrderZero.OperatorTheoremCertified || !a.OrderZero.RequiresOperatorJOpposite {
		t.Fatalf("bad order-zero state: %s", FormatOrderZero(a.OrderZero))
	}
	if !containsAll(a.OrderZero.Supports, []string{SupportOrderZeroBlockSupport, SupportAForientSupportBimodule}) {
		t.Fatalf("missing order-zero supports: %+v", a.OrderZero.Supports)
	}
}

func TestGate858MinimalCarrierJClosureDoesNotRestoreAmbientPuncture(t *testing.T) {
	a, err := BuildDefault()
	if err != nil {
		t.Fatal(err)
	}
	if a.Carrier.HPartMinRank != 15 || a.Carrier.HFMinRank != 30 || a.Carrier.AmbientPartRank != 16 || a.Carrier.AmbientFRank != 32 {
		t.Fatalf("bad carrier ranks: %s", FormatCarrier(a.Carrier))
	}
	if !a.Carrier.RightPunctureOutsideMinimal || !a.Carrier.LeftKernelPresent || a.Carrier.JCopyRestoresAmbientPuncture {
		t.Fatalf("bad minimal/J closure: %s", FormatCarrier(a.Carrier))
	}
}

func TestGate858EdgesRemainSupportOnlyAndNoMagnitudes(t *testing.T) {
	a, err := BuildDefault()
	if err != nil {
		t.Fatal(err)
	}
	if len(a.Edges) != 3 {
		t.Fatalf("expected three edges: %s", FormatEdges(a.Edges))
	}
	for _, e := range a.Edges {
		if !e.LeftSupportCompatible || !e.RightSupportCompatible || e.OperatorIntertwinerCertified || e.YukawaMagnitude {
			t.Fatalf("edge overpromoted: %+v", e)
		}
	}
}

func TestGate858Theorem(t *testing.T) {
	res := Generation2JFOppositeActionOrderZeroBimoduleRealizationAuditTheorem().Verify()
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
