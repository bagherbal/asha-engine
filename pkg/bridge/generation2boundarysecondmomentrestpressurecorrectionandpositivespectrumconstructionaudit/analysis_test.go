package generation2boundarysecondmomentrestpressurecorrectionandpositivespectrumconstructionaudit

import (
	"math"
	"strings"
	"testing"
)

func TestGate813NumericalInheritanceAndDirectClosure(t *testing.T) {
	a, err := BuildDefault()
	if err != nil {
		t.Fatal(err)
	}
	if !a.Inheritance.Gate811Inherited || !a.Inheritance.Gate812Inherited || !a.Inheritance.BoundarySecondMomentSelected {
		t.Fatalf("bad inheritance: %+v", a.Inheritance)
	}
	if math.Abs(a.Inheritance.M2-1.624013231638281e-7) > 1e-20 {
		t.Fatalf("bad M2: %s", FormatInheritance(a.Inheritance))
	}
	if math.Abs(a.Inheritance.C2Obs-5.8299915725) > 1e-9 {
		t.Fatalf("bad c2 obs: %s", FormatInheritance(a.Inheritance))
	}
	if math.Abs(a.DirectDelta.Residual+2.76095936e-8) > 1e-15 || a.DirectDelta.ResidualImprovement < 30 {
		t.Fatalf("bad direct second-moment closure: %s", FormatDirectDelta(a.DirectDelta))
	}
	if !a.DirectDelta.PositiveBandExists || !(a.DirectDelta.AlphaMin < a.DirectDelta.AlphaMaxTopBranch) {
		t.Fatalf("direct closure positive band missing: %s", FormatDirectDelta(a.DirectDelta))
	}
}

func TestGate813AlphaPositivityFamily(t *testing.T) {
	a, err := BuildDefault()
	if err != nil {
		t.Fatal(err)
	}
	if a.NaiveAlpha.BetaRequired >= 0 || a.NaiveAlpha.Positive {
		t.Fatalf("naive alpha should require negative beta: %+v", a.NaiveAlpha)
	}
	if math.Abs(a.LowerBound.CorrectionOverM2-0.5086108926) > 1e-9 {
		t.Fatalf("bad lower-bound correction: %+v", a.LowerBound)
	}
	if a.AlphaFamily.HalfM2.Beta >= 0 || a.AlphaFamily.HalfM2.ValidBeta {
		t.Fatalf("half M2 should still be slightly negative: %s", FormatAlphaCandidate(a.AlphaFamily.HalfM2))
	}
	for _, c := range []AlphaCandidate{a.AlphaFamily.ThreeFifths, a.AlphaFamily.SixElevenths, a.AlphaFamily.ObservedMin} {
		if !c.ValidBeta || !c.ValidQRest || c.QRest < -1e-15 || c.QRest > 1+1e-15 {
			t.Fatalf("candidate should be positive-rest compatible: %s", FormatAlphaCandidate(c))
		}
	}
	if !containsAll(a.AlphaFamily.Failures, []string{StatusHalfM2StillNegative, StatusNoNativeCAlpha}) {
		t.Fatalf("missing alpha-family failures: %s", FormatAlphaFamily(a.AlphaFamily))
	}
}

func TestGate813SpurionImpactAndFirewalls(t *testing.T) {
	a, err := BuildDefault()
	if err != nil {
		t.Fatal(err)
	}
	if !(math.Abs(a.Spurion.B2Diff) < math.Abs(a.Spurion.B1Diff)) {
		t.Fatalf("B2 spurion should improve B1: %s", FormatSpurion(a.Spurion))
	}
	if math.Abs(a.Impact.CYukawaOfficial-CYukawa) > 1e-15 || math.Abs(a.Impact.CHiggsOfficial-CHiggs) > 1e-15 {
		t.Fatalf("official ledger should be preserved: %s", FormatImpact(a.Impact))
	}
	if !a.Firewalls.Enforced || a.Firewalls.Verdict != StatusFirewallGate813 {
		t.Fatalf("bad firewalls: %+v", a.Firewalls)
	}
	res := Generation2BoundarySecondMomentRestPressureCorrectionAndPositiveSpectrumConstructionAuditTheorem().Verify()
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
