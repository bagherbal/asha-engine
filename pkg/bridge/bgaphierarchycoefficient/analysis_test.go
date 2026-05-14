package bgaphierarchycoefficient

import (
	"math"
	"testing"
)

func TestGate283BGapHierarchyCoefficient(t *testing.T) {
	a, err := Build()
	if err != nil {
		t.Fatalf("Build() error = %v", err)
	}
	if !a.PreviousGate282.PathBClosed || !a.PreviousGate282.HiggsFirewallActive {
		t.Fatalf("expected Gate 282 capstone inherited: %s", FormatGate282(a.PreviousGate282))
	}
	if math.Abs(a.Volumes.UnitS3Volume-2*math.Pi*math.Pi) > 1e-12 {
		t.Fatalf("unexpected S3 volume: %s", FormatVolumes(a.Volumes))
	}
	if math.Abs(a.ContactAction.Coefficient-4/math.Pi) > 1e-12 {
		t.Fatalf("coefficient should be exactly 4/pi: %s", FormatContactAction(a.ContactAction))
	}
	if a.ContactAction.ContactBoundaryActionMapDerived || a.ContactAction.HopfFiberNormalizationDerived {
		t.Fatalf("native action map must remain missing: %s", FormatContactAction(a.ContactAction))
	}
	if !a.Hierarchy.TightNearResonance || a.Hierarchy.ExactIntermediateMatch || a.Hierarchy.TheoremUpgradeGranted {
		t.Fatalf("expected tight resonance but no theorem upgrade: %s", FormatHierarchy(a.Hierarchy))
	}
	if a.Hierarchy.Log10Gap > 0.02 {
		t.Fatalf("expected Gate 229 tight residual, got: %s", FormatHierarchy(a.Hierarchy))
	}
	if !a.Sensitivity.BindingWarning || a.Sensitivity.OnePercentShiftDecades < 0.05 || a.Sensitivity.TenPercentShiftDecades < 0.5 {
		t.Fatalf("expected binding exponential sensitivity: %s", FormatSensitivity(a.Sensitivity))
	}
	if a.Seal.IntermediateBreakingSealGranted || !a.Seal.RequiresContactActionMap || !a.Seal.RequiresFiniteOrderParameter {
		t.Fatalf("seal ledger violated: %s", FormatSeal(a.Seal))
	}
	if a.Firewall.FiniteCorePolluted || !a.Firewall.DoesNotFitCoefficient || !a.Firewall.DoesNotClaimExactMIntTheorem {
		t.Fatalf("firewall failure: %s", FormatFirewall(a.Firewall))
	}
}

func TestGate283TheoremPassesChecks(t *testing.T) {
	res := BGapHierarchyCoefficientTopologicalVolumeRatioAuditTheorem().Run()
	if !res.Passed() {
		t.Fatalf("theorem checks failed:\n%s", res.Details())
	}
	if res.Status == "EXACT_FINITE" || res.Status == "PHENOMENOLOGY" {
		t.Fatalf("Gate 283 should remain a bridge-required audit, got %s", res.Status)
	}
}
