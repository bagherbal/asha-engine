package generation2socketcharacteridentificationedgeintertwinerpromotionaudit

import (
	"strings"
	"testing"
)

func TestGate862CharacterLedgers(t *testing.T) {
	a, err := BuildDefault()
	if err != nil {
		t.Fatal(err)
	}
	if !charactersOK(a.Characters) {
		t.Fatalf("bad characters: %s", FormatCharacters(a.Characters))
	}
}

func TestGate862CharacterIdentificationSeal(t *testing.T) {
	a, err := BuildDefault()
	if err != nil {
		t.Fatal(err)
	}
	if !a.IDMap.Defined || !a.IDMap.OrientationSeal || a.IDMap.Native || a.IDMap.OperatorCertified || a.IDMap.ForcesPunctureEdge {
		t.Fatalf("bad character identification map: %s", FormatIdentification(a.IDMap))
	}
	if !containsAll(a.IDMap.Failures, []string{FailureCRToCHNotNative, FailureSocketMatchSeal}) {
		t.Fatalf("missing character firewall: %s", FormatIdentification(a.IDMap))
	}
}

func TestGate862ActiveEdgesIntertwineGivenSeal(t *testing.T) {
	a, err := BuildDefault()
	if err != nil {
		t.Fatal(err)
	}
	if !activeIntertwinersOK(a.Edges) {
		t.Fatalf("bad edge intertwiners: %s", FormatEdges(a.Edges))
	}
}

func TestGate862PunctureNotForced(t *testing.T) {
	a, err := BuildDefault()
	if err != nil {
		t.Fatal(err)
	}
	if !punctureOK(a.Edges) || !a.Kernel.PunctureZero || !a.Kernel.PunctureNotForcedByCharacterID {
		t.Fatalf("puncture not preserved: %s | %s", FormatEdges(a.Edges), FormatKernel(a.Kernel))
	}
}

func TestGate862FirstOrderSharpenedNotPromoted(t *testing.T) {
	a, err := BuildDefault()
	if err != nil {
		t.Fatal(err)
	}
	if !a.FirstOrder.StabilizerOperatorCompatibilitySharpened || a.FirstOrder.StabilizerOperatorCompatibilityNative || a.FirstOrder.FullUnbrokenCompatibilityCertified {
		t.Fatalf("first-order overpromoted or not sharpened: %s", FormatFirstOrder(a.FirstOrder))
	}
}

func TestGate862NoMagnitudeOrLedgerPromotion(t *testing.T) {
	a, err := BuildDefault()
	if err != nil {
		t.Fatal(err)
	}
	if !a.Ledger.OfficialFrozen || a.Ledger.AlphaNative || a.Ledger.R3 || a.Ledger.R4 || a.Impact.CanUpdateNEff || a.Impact.CanUpdateCYukawa || a.Impact.CanUpdateCHiggs || a.Impact.CanPromoteToR3 || a.Impact.CanPromoteToR4 {
		t.Fatalf("ledger overpromoted: %s | %s", FormatLedger(a.Ledger), FormatImpact(a.Impact))
	}
	for _, e := range a.Edges {
		if e.NumericalValue || e.YukawaMagnitude || e.OperatorCertified {
			t.Fatalf("edge overpromoted: %+v", e)
		}
	}
}

func TestGate862Theorem(t *testing.T) {
	res := Generation2SocketCharacterIdentificationEdgeIntertwinerPromotionAuditTheorem().Verify()
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
