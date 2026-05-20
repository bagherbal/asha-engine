package generation2hyperchargecolorboundarycoefficientandpositiverestcorrectionaudit

import (
	"math"
	"strings"
	"testing"
)

func TestGate811FactorizationsAndCorrectionScale(t *testing.T) {
	a, err := BuildDefault()
	if err != nil {
		t.Fatal(err)
	}
	if !a.Inheritance.Gate810Inherited || !a.Inheritance.CoeffTargetSelected || !a.Inheritance.CorrectionTargetSelected {
		t.Fatalf("bad inheritance: %+v", a.Inheritance)
	}
	if !containsAll(a.NineFifths.Supports, []string{StatusNineFiveTypedCandidate, StatusRestCoeffColorHypercharge}) || !strings.Contains(a.NineFifths.Expression, "9/5") {
		t.Fatalf("bad 9/5 audit: %s", FormatFactor(a.NineFifths))
	}
	if !containsAll(a.ThreeTenths.Supports, []string{StatusThreeTenthsTypedCandidate, StatusHalfBoundaryPairCandidate}) || !strings.Contains(a.ThreeTenths.Expression, "3/10") {
		t.Fatalf("bad 3/10 audit: %s", FormatFactor(a.ThreeTenths))
	}
	if math.Abs(a.Correction.DeltaAlpha-8.2599081954e-8) > 5e-17 || math.Abs(a.Correction.M2-1.624013231638281e-7) > 1e-20 || math.Abs(a.Correction.HalfM2-8.120066158191404e-8) > 1e-20 {
		t.Fatalf("bad correction: %s", FormatCorrection(a.Correction))
	}
	if math.Abs(a.Correction.DeltaOverHalfM2-1.01722178515078) > 5e-13 {
		t.Fatalf("expected delta_alpha slightly above half M2: %s", FormatCorrection(a.Correction))
	}
}

func TestGate811CorrectedAlphaAndDeltaResidual(t *testing.T) {
	a, err := BuildDefault()
	if err != nil {
		t.Fatal(err)
	}
	if a.AlphaCorr.BetaNonnegative || a.AlphaCorr.QRestValid {
		t.Fatalf("half-M2 alpha correction should remain slightly positivity-blocked: %s", FormatAlphaCorr(a.AlphaCorr))
	}
	if math.Abs(a.AlphaCorr.BetaCorr-(-2.795756404161409e-9)) > 5e-17 || math.Abs(a.AlphaCorr.AlphaCorrMinusAlphaMin-(-1.398420347890287e-9)) > 5e-18 {
		t.Fatalf("bad corrected alpha: %s", FormatAlphaCorr(a.AlphaCorr))
	}
	if !containsAll(a.AlphaCorr.Failures, []string{StatusAlphaCorrNeedsBeta, StatusAlphaCorrNotNative}) {
		t.Fatalf("missing alpha correction failures: %s", FormatAlphaCorr(a.AlphaCorr))
	}
	if math.Abs(a.DeltaCorr.C2Obs-5.8299915722461693) > 5e-13 || math.Abs(a.DeltaCorr.CandidateResidual-(-2.7609593616136768e-8)) > 5e-18 {
		t.Fatalf("bad direct delta correction: %s", FormatDeltaCorr(a.DeltaCorr))
	}
	if math.Abs(a.DeltaCorr.CandidateRelErr-(-1.1863116249617649e-5)) > 5e-16 {
		t.Fatalf("bad corrected relative residual: %s", FormatDeltaCorr(a.DeltaCorr))
	}
}

func TestGate811ControlsImpactAndTheorem(t *testing.T) {
	a, err := BuildDefault()
	if err != nil {
		t.Fatal(err)
	}
	if len(a.Controls.Controls) != 5 || !strings.Contains(a.Controls.Controls[4].Name, "9/5 + 6ps") || math.Abs(a.Controls.Controls[4].AbsResidual-2.7609593616136768e-8) > 5e-18 {
		t.Fatalf("bad controls: %s", FormatControls(a.Controls))
	}
	if math.Abs(a.Impact.CYukawaBoundary-0.9992248096922658) > 1e-15 || math.Abs(a.Impact.CHiggsBoundary-1.0372205108665145) > 1e-15 {
		t.Fatalf("bad impact: %s", FormatImpact(a.Impact))
	}
	if !a.Firewalls.Enforced || a.Firewalls.Verdict != StatusFirewallGate811 {
		t.Fatalf("bad firewalls: %+v", a.Firewalls)
	}
	res := Generation2HyperchargeColorBoundaryCoefficientAndPositiveRestCorrectionAuditTheorem().Verify()
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
