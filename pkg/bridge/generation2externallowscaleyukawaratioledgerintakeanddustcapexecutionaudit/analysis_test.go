package generation2externallowscaleyukawaratioledgerintakeanddustcapexecutionaudit

import (
	"math"
	"testing"
)

func TestGate823NoLedgerDataRequired(t *testing.T) {
	a, err := BuildDefault()
	if err != nil {
		t.Fatal(err)
	}
	if a.Search.Found || !a.Status.DataRequired || !a.Status.LiteralSectorFrozen {
		t.Fatalf("expected absent ledger DATA_REQUIRED freeze: search=%+v status=%+v", a.Search, a.Status)
	}
	if math.Abs(a.Ledger.SqrtBOverT-0.019691251452864992) > 1e-16 {
		t.Fatalf("bad inherited target: %s", FormatLedger(a.Ledger))
	}
	if math.Abs(a.Ledger.UncoloredCap-0.0006718553149936293) > 1e-16 {
		t.Fatalf("bad uncolored cap: %s", FormatLedger(a.Ledger))
	}
}

func TestGate823ConventionLockValidation(t *testing.T) {
	partial := ExternalLowScaleYukawaRatioLedger{SourceLabel: "test", ScaleMu: "M_Z"}
	if partial.ConventionLocked() {
		t.Fatalf("partial ledger should not be convention locked")
	}
	missing := partial.MissingConventionObjects()
	if len(missing) == 0 {
		t.Fatalf("expected missing objects")
	}
	full := samplePassingLedger()
	if !full.ConventionLocked() {
		t.Fatalf("full sample should be convention locked, missing=%v", full.MissingConventionObjects())
	}
}

func TestGate823ExecutionHelpers(t *testing.T) {
	pass := samplePassingLedger()
	res := ExecuteBranch(pass, "b", 0.05)
	if !res.LargeTripletMatched || !res.ColoredDustOK || !res.UncoloredDustOK || !res.LiteralSectorSurvives {
		t.Fatalf("expected literal sector pass for synthetic pass ledger: %s", FormatExecution(res))
	}

	coloredFail := samplePassingLedger()
	coloredFail.ColoredRatios["c"] = AlphaB(SBoundary) * 2
	res = ExecuteBranch(coloredFail, "b", 0.05)
	if res.ColoredDustOK || res.LiteralSectorSurvives {
		t.Fatalf("expected colored dust failure: %s", FormatExecution(res))
	}

	uncoloredFail := samplePassingLedger()
	uncoloredFail.UncoloredRatios["tau"] = UncoloredCap(AlphaB(SBoundary)) * 2
	res = ExecuteBranch(uncoloredFail, "b", 0.05)
	if res.UncoloredDustOK || res.LiteralSectorSurvives {
		t.Fatalf("expected uncolored dust failure: %s", FormatExecution(res))
	}

	abstract := ExecuteAbstractColoredChamber(pass, 0.05)
	if !abstract.LiteralSectorSurvives || abstract.Selected != "b" {
		t.Fatalf("expected abstract chamber to select b and pass: %s", FormatExecution(abstract))
	}
}

func TestGate823TheoremAndFirewalls(t *testing.T) {
	res := Generation2ExternalLowScaleYukawaRatioLedgerIntakeAndDustCapExecutionAuditTheorem().Verify()
	if string(res.Status) == "FAILED_ROUTE" {
		t.Fatalf("theorem failed: %+v", res)
	}
	a, err := BuildDefault()
	if err != nil {
		t.Fatal(err)
	}
	if !a.Firewalls.Enforced || !a.Firewalls.NoInferenceWithoutLedger || !a.Firewalls.NoCYukawaUpdate || a.Firewalls.Verdict != StatusFirewallGate823 {
		t.Fatalf("firewalls not enforced: %+v", a.Firewalls)
	}
}

func samplePassingLedger() ExternalLowScaleYukawaRatioLedger {
	alpha := AlphaB(SBoundary)
	return ExternalLowScaleYukawaRatioLedger{
		SourceLabel: "synthetic unit-test ledger", ScaleMu: "M_Z", Scheme: "declared", Normalization: "dimensionless", TopSelector: "t", ColorConvention: "coefficient-color", NeutrinoConvention: "absent", Uncertainties: map[string]float64{},
		ColoredRatios: map[string]float64{
			"b": math.Sqrt(BOverT(alpha)),
			"c": alpha * 0.2,
			"s": alpha * 0.1,
			"u": alpha * 0.01,
			"d": alpha * 0.02,
		},
		UncoloredRatios: map[string]float64{
			"tau": UncoloredCap(alpha) * 0.1,
			"mu":  UncoloredCap(alpha) * 0.01,
			"e":   UncoloredCap(alpha) * 0.001,
		},
	}
}
