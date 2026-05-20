package generation2boundaryfnrestpressurespurionandneffminusthreeclosureaudit

import (
	"math"
	"strings"
	"testing"
)

func TestGate810DirectClosureAndSpurion(t *testing.T) {
	a, err := BuildDefault()
	if err != nil {
		t.Fatal(err)
	}
	if math.Abs(a.Direct.CObs-1.8007325638446063) > 5e-14 || math.Abs(a.Direct.Residual-9.4679834536684e-7) > 2e-18 || math.Abs(a.Direct.RelativeResidual-0.0004068143483817) > 5e-16 {
		t.Fatalf("bad direct closure: %s", FormatDirect(a.Direct))
	}
	if !containsAll(a.Direct.Failures, []string{StatusNineFifthsNotExact, StatusNumericalNotTheorem}) {
		t.Fatalf("missing direct failures: %s", FormatDirect(a.Direct))
	}
	if math.Abs(a.Spurion.EpsilonN-0.21964195823344188) > 3e-15 || math.Abs(a.Spurion.EpsilonB-0.21961961644976352) > 1e-15 || math.Abs(a.Spurion.Difference-0.00002234178367725) > 3e-14 {
		t.Fatalf("bad spurion: %s", FormatSpurion(a.Spurion))
	}
	if !containsAll(a.Spurion.Failures, []string{StatusEpsilonBNotNative, StatusNoBoundarySpurionMap, StatusCabibboNotTheorem}) {
		t.Fatalf("missing spurion failures: %s", FormatSpurion(a.Spurion))
	}
}

func TestGate810CoefficientAlphaAndPositivity(t *testing.T) {
	a, err := BuildDefault()
	if err != nil {
		t.Fatal(err)
	}
	if !containsAll(a.Coefficient.Supports, []string{StatusNineFifthsTypedCandidate, StatusFiveThirdsNonarbitrary}) || !containsAll(a.Coefficient.Failures, []string{StatusNoColorHyperchargeTheorem, StatusInverseHyperchargeNotAuto, StatusNoRationalFit}) {
		t.Fatalf("bad coefficient audit: %s", FormatCoeff(a.Coefficient))
	}
	if math.Abs(a.Alpha.AlphaApprox-0.0003878912453691245) > 1e-15 || math.Abs(a.Alpha.AlphaBoundary-0.00038773344564488885) > 1e-15 || math.Abs(a.Alpha.Residual-1.5779972422780667e-7) > 1e-18 {
		t.Fatalf("bad alpha closure: %s", FormatAlpha(a.Alpha))
	}
	if a.Positivity.BetaRequired >= 0 || math.Abs(a.Positivity.BetaRequired-(-1.651341154285823e-7)) > 3e-16 {
		t.Fatalf("bad beta positivity: %s", FormatPositivity(a.Positivity))
	}
	if math.Abs(a.Positivity.AlphaMinOverS-0.3000639091748843) > 5e-14 || math.Abs(a.Positivity.CorrectionNeeded-8.2599081954e-8) > 5e-17 {
		t.Fatalf("bad alpha min: %s", FormatPositivity(a.Positivity))
	}
	if !containsAll(a.Positivity.Failures, []string{StatusAlphaThreeTenthsNotExact, StatusFirstOrderNotTheorem, StatusPositiveBlocksExact}) {
		t.Fatalf("missing positivity failures: %s", FormatPositivity(a.Positivity))
	}
}

func TestGate810RestConcentrationControlsAndCHiggs(t *testing.T) {
	a, err := BuildDefault()
	if err != nil {
		t.Fatal(err)
	}
	if len(a.Concentration.Regimes) != 3 {
		t.Fatalf("expected 3 regimes: %s", FormatConcentration(a.Concentration))
	}
	if math.Abs(a.Concentration.Regimes[1].QRest-0.33307493962706697) > 5e-10 || math.Abs(a.Concentration.Regimes[2].AlphaOverS-0.3002387347866694) > 5e-13 {
		t.Fatalf("bad concentration: %s", FormatConcentration(a.Concentration))
	}
	if !containsAll(a.Map.Failures, []string{StatusNoBoundaryFNMap, StatusNoSectorTraceRule, StatusNoPositiveConcentration, StatusNoScaleStability}) {
		t.Fatalf("bad map failures: %s", FormatMap(a.Map))
	}
	if len(a.Controls.Controls) != 4 || math.Abs(a.Controls.Controls[3].AbsResidual-9.4679834536684e-7) > 2e-18 || !strings.Contains(a.Controls.BestTyped, "9/5") {
		t.Fatalf("bad controls: %s", FormatControls(a.Controls))
	}
	if math.Abs(a.CHiggs.CYukawaCandidate-0.9992251339916449) > 1e-15 || math.Abs(a.CHiggs.CYukawaResidual-3.151104440712871e-7) > 2e-18 || math.Abs(a.CHiggs.CHiggsCandidate-1.0372208474974351) > 1e-15 {
		t.Fatalf("bad C_Higgs candidate: %s", FormatCHiggs(a.CHiggs))
	}
	if !containsAll(a.CHiggs.Failures, []string{StatusNoCYukawaUpdate, StatusCHiggsLevelB, StatusCandidateNotLevelC}) {
		t.Fatalf("missing C_Higgs failures: %s", FormatCHiggs(a.CHiggs))
	}
}

func TestGate810BranchFirewallsAndTheorem(t *testing.T) {
	a, err := BuildDefault()
	if err != nil {
		t.Fatal(err)
	}
	if !a.Inheritance.Gate809Inherited || !a.Inheritance.CandidateSelected {
		t.Fatalf("bad inheritance")
	}
	if !strings.Contains(a.Branch.Next, "Gate 811") || !strings.Contains(a.Branch.Next, "Hypercharge-Color") {
		t.Fatalf("bad branch: %+v", a.Branch)
	}
	if !a.Firewalls.Enforced || a.Firewalls.Verdict != StatusFirewallGate810 {
		t.Fatalf("bad firewall: %+v", a.Firewalls)
	}
	res := Generation2BoundaryFNRestPressureSpurionAndNEffMinusThreeClosureAuditTheorem().Verify()
	if res.Status == "FAILED_ROUTE" {
		t.Fatalf("unexpected theorem failure: %+v", res)
	}
	for _, c := range res.Checks {
		if !c.Passed {
			t.Fatalf("check failed: %s %s", c.Name, c.Detail)
		}
	}
	joined := strings.Join(res.Notes, "\n")
	for _, want := range Statuses() {
		if !strings.Contains(joined, want) {
			t.Fatalf("missing status %s", want)
		}
	}
}
