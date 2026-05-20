package generation2kappaeorientationresidualandhyperchargenormalizedboundarysquareaudit

import (
	"strings"
	"testing"
)

func TestGate747HyperchargeBoundarySquareCompression(t *testing.T) {
	a, err := Build()
	if err != nil {
		t.Fatal(err)
	}
	if !a.Gate746.Inherited || !a.Gate746.KappaEActiveInput || !a.Gate746.OrientationClose || !a.Gate746.FlavorFirewallKept {
		t.Fatalf("bad Gate746 inheritance: %+v", a.Gate746)
	}
	if !Near(a.Ratio.DeltaKappaE, -2.77587313789e-6, 1e-15) || !Near(a.Ratio.Ratio, -1.6617879079741393, 1e-12) || !a.Ratio.CloseToMinusFiveThirds || a.Ratio.BestCandidate == "" {
		t.Fatalf("bad residual ratio audit: %+v", a.Ratio)
	}
	if !Near(a.Correction.Correction, -2.7840226828084814e-6, 1e-18) || !Near(a.Correction.KappaEHyperBoundary, 0.005503546042029642, 1e-18) || !Near(a.Correction.ResidualAfterCorrection, 8.149544918367644e-9, 1e-18) || a.Correction.CompressionFactor < 330 || !a.Correction.CorrectionNotExact {
		t.Fatalf("bad hypercharge-boundary correction: %+v", a.Correction)
	}
}

func TestGate747RuntimeReplacementAndFirewalls(t *testing.T) {
	a, err := Build()
	if err != nil {
		t.Fatal(err)
	}
	if !Near(a.Replacement.RuntimeOrientShift, 1.3795355913170937e-8, 5e-14) || !Near(a.Replacement.RuntimeHyperBoundaryShift, -4.050107471620379e-11, 5e-14) || a.Replacement.ImprovementFactor < 330 || a.Replacement.ImprovementFactor > 350 || !a.Replacement.ReplacementNotNative {
		t.Fatalf("bad replacement audit: %+v", a.Replacement)
	}
	if a.Firewall.DerivesFlavorTheorem || a.Firewall.DerivesScalarRuntime || a.Firewall.DerivesHiggsMass || a.Firewall.DerivesYukawa || !a.Firewall.FiveThirdsMatureButUncoupled || !a.Firewall.SSplitBoundaryNotFlavorOperator {
		t.Fatalf("bad firewall: %+v", a.Firewall)
	}
	res := Generation2KappaEOrientationResidualAndHyperchargeNormalizedBoundarySquareAuditTheorem().Verify()
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
			t.Fatalf("missing status note %s", want)
		}
	}
}
