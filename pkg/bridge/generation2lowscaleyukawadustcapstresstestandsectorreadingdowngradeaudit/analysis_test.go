package generation2lowscaleyukawadustcapstresstestandsectorreadingdowngradeaudit

import (
	"math"
	"testing"
)

func TestGate822DustCapNumbers(t *testing.T) {
	a, err := BuildDefault()
	if err != nil {
		t.Fatal(err)
	}
	if math.Abs(a.Ledger.SqrtBOverT-0.019691251452864992) > 1e-16 {
		t.Fatalf("bad large triplet target: %s", FormatLedger(a.Ledger))
	}
	if math.Abs(a.Ledger.ExtraColoredTraceCap-1.5046318809506294e-7) > 1e-20 {
		t.Fatalf("bad extra colored trace cap: %s", FormatLedger(a.Ledger))
	}
	if math.Abs(a.Ledger.ExtraColoredYukawaCap-a.Ledger.AlphaB) > 1e-18 {
		t.Fatalf("bad extra colored Yukawa cap: %s", FormatLedger(a.Ledger))
	}
	if math.Abs(a.Ledger.UncoloredYukawaCap-0.0006718553149936293) > 1e-16 {
		t.Fatalf("bad uncolored Yukawa cap: %s", FormatLedger(a.Ledger))
	}
}

func TestGate822StressTestHelpers(t *testing.T) {
	alpha := AlphaB(SBoundary)
	pass := RatioLedger{
		Scale: "M_Z", Scheme: "test", Normalization: "dimensionless", TopSelector: "t",
		ColoredYukawaRatios: map[string]float64{
			"b": math.Sqrt(BOverT(alpha)),
			"c": alpha * 0.25,
			"s": alpha * 0.1,
		},
		UncoloredYukawaRatios: map[string]float64{
			"tau": math.Sqrt(DustOverT(alpha)) * 0.1,
		},
	}
	res := StressTestLiteralSector(pass, "b", 0.1)
	if !res.LargeTripletMatch || !res.ColoredDustOK || !res.UncoloredDustOK {
		t.Fatalf("expected dust-cap pass for selected branch: %+v", res)
	}

	coloredFail := pass
	coloredFail.ColoredYukawaRatios = map[string]float64{"b": math.Sqrt(BOverT(alpha)), "c": alpha * 2}
	coloredFail.UncoloredYukawaRatios = map[string]float64{}
	res = StressTestLiteralSector(coloredFail, "b", 0.1)
	if res.ColoredDustOK || res.LiteralSectorReadingSurvives {
		t.Fatalf("expected colored dust failure: %+v", res)
	}

	uncoloredFail := pass
	uncoloredFail.ColoredYukawaRatios = map[string]float64{"b": math.Sqrt(BOverT(alpha))}
	uncoloredFail.UncoloredYukawaRatios = map[string]float64{"tau": math.Sqrt(DustOverT(alpha)) * 2}
	res = StressTestLiteralSector(uncoloredFail, "b", 0.1)
	if res.UncoloredDustOK || res.LiteralSectorReadingSurvives {
		t.Fatalf("expected uncolored dust failure: %+v", res)
	}
}

func TestGate822TheoremAndFirewalls(t *testing.T) {
	a, err := BuildDefault()
	if err != nil {
		t.Fatal(err)
	}
	if !a.Protocol.CanFalsify {
		t.Fatalf("expected falsification protocol")
	}
	if a.Status.ExternalLedgerSupplied || a.Status.CanUpdateCYukawa {
		t.Fatalf("gate should not decide sector status or update C_Yukawa: %+v", a.Status)
	}
	if !a.Firewalls.Enforced || !a.Firewalls.ExtraColoredCap || !a.Firewalls.UncoloredCap || a.Firewalls.Verdict != StatusFirewallGate822 {
		t.Fatalf("firewalls not enforced: %+v", a.Firewalls)
	}
	th := Generation2LowScaleYukawaDustCapStressTestAndSectorReadingDowngradeAuditTheorem()
	result := th.Verify()
	if string(result.Status) == "FAILED_ROUTE" {
		t.Fatalf("theorem failed: %+v", result)
	}
}
