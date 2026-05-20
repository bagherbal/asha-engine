package generation2kappaehyperchargeboundaryresidualandboundarystressmomentaudit

import (
	"strings"
	"testing"
)

func TestGate748BoundaryStressMomentCompression(t *testing.T) {
	a, err := Build()
	if err != nil {
		t.Fatal(err)
	}
	if !a.Gate747.Inherited || !a.Gate747.HyperBoundaryCorrectionClose || !a.Gate747.ResidualNotZero || !a.Gate747.FlavorFirewallKept {
		t.Fatalf("bad Gate747 inheritance: %+v", a.Gate747)
	}
	if !Near(a.Residual.Residual, 8.149544918e-9, 1e-17) || !Near(a.Residual.M2Wall, 1.624013231638281e-7, 1e-19) || !Near(a.Residual.Ratio, 0.05018151797950743, 1e-12) || !a.Residual.SecondMomentScale {
		t.Fatalf("bad residual over M2 audit: %+v", a.Residual)
	}
	if !Near(a.Stress.XiBoundary, 0.0503471644870914, 1e-15) || a.Stress.BestCandidate != "xi_boundary midpoint" {
		t.Fatalf("bad boundary stress candidates: %+v", a.Stress)
	}
	if !Near(a.Correction.StressMoment, 8.176446130250547e-9, 1e-18) || !Near(a.Correction.KappaEHyperStress, 0.005503554218475772, 1e-18) || !Near(a.Correction.ResidualAfterCorrection, -2.6901212160646004e-11, 1e-18) || a.Correction.CompressionFactor < 250 || !a.Correction.CorrectionNotExact {
		t.Fatalf("bad stress moment correction: %+v", a.Correction)
	}
}

func TestGate748RuntimeReplacementAndFirewalls(t *testing.T) {
	a, err := Build()
	if err != nil {
		t.Fatal(err)
	}
	if !Near(a.Replacement.RuntimeOrientShift, 1.3795355913170937e-8, 5e-14) || !Near(a.Replacement.RuntimeHyperBoundaryShift, -4.050107471620379e-11, 5e-14) || !Near(a.Replacement.RuntimeHyperStressShift, 1.3369860774048448e-13, 5e-15) || a.Replacement.StressImprovementOverOrient < 1e5 || a.Replacement.StressImprovementOverHyper < 100 || !a.Replacement.ReplacementNotNative {
		t.Fatalf("bad replacement audit: %+v", a.Replacement)
	}
	if a.Firewall.DerivesFlavorTheorem || a.Firewall.DerivesScalarRuntime || a.Firewall.DerivesHiggsMass || a.Firewall.DerivesYukawa || !a.Firewall.XiBoundaryBridgeStressQuantity || !a.Firewall.M2WallBoundaryMomentNotFlavor {
		t.Fatalf("bad firewall: %+v", a.Firewall)
	}
	res := Generation2KappaEHyperchargeBoundaryResidualAndBoundaryStressMomentAuditTheorem().Verify()
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
