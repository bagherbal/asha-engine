package gravityspectralactionf2

import (
	"math"
	"testing"
)

func close(a, b, tol float64) bool {
	if a > b {
		return a-b <= tol
	}
	return b-a <= tol
}

func TestBuildDefault(t *testing.T) {
	a, err := BuildDefault()
	if err != nil {
		t.Fatalf("BuildDefault returned error: %v", err)
	}
	if a.Inputs.HighestInheritedGate != inheritedHighestGate || a.Inputs.NGen != 3 {
		t.Fatalf("bad inputs: %s", FormatInputs(a.Inputs))
	}
}

func TestEinsteinHilbertCoefficient(t *testing.T) {
	a, err := BuildDefault()
	if err != nil {
		t.Fatalf("BuildDefault returned error: %v", err)
	}
	want := 8 / (math.Pi * math.Pi)
	if !close(a.EH.SpectralCoefficientCG, want, 1e-15) {
		t.Fatalf("bad C_G: %s", FormatEH(a.EH))
	}
}

func TestF2Targets(t *testing.T) {
	a, err := BuildDefault()
	if err != nil {
		t.Fatalf("BuildDefault returned error: %v", err)
	}
	if !close(a.Target.InvariantF2LambdaOverUnreduced, math.Pi/64, 1e-15) {
		t.Fatalf("bad unreduced invariant: %s", FormatTarget(a.Target))
	}
	if !close(a.Target.ReducedPlanckCutoffF2Target, math.Pi*math.Pi/8, 1e-15) {
		t.Fatalf("bad reduced invariant: %s", FormatTarget(a.Target))
	}
	if a.Target.VEVCutoffF2Target < 1e31 {
		t.Fatalf("VEV cutoff target should be huge: %s", FormatTarget(a.Target))
	}
}

func TestScaleChoiceSieve(t *testing.T) {
	a, err := BuildDefault()
	if err != nil {
		t.Fatalf("BuildDefault returned error: %v", err)
	}
	if len(a.ScaleSieve) != 3 {
		t.Fatalf("expected 3 scale candidates: %v", a.ScaleSieve)
	}
	if a.ScaleSieve[0].Name != "Λ = unreduced Planck mass" || !close(a.ScaleSieve[0].F2Required, math.Pi/64, 1e-15) {
		t.Fatalf("bad Planck candidate: %s", FormatScale(a.ScaleSieve[0]))
	}
}

func TestResonanceAuditPreservesFirewall(t *testing.T) {
	a, err := BuildDefault()
	if err != nil {
		t.Fatalf("BuildDefault returned error: %v", err)
	}
	if a.Resonance.NativeMatchFound {
		t.Fatalf("unexpected native f2 resonance: %s", FormatResonance(a.Resonance))
	}
	if a.Resonance.BestCandidate.Name == "π/64 target itself" {
		t.Fatalf("identity target should not count as best independent candidate: %s", FormatResonance(a.Resonance))
	}
}

func TestStatusesAndTheorem(t *testing.T) {
	a, err := BuildDefault()
	if err != nil {
		t.Fatalf("BuildDefault returned error: %v", err)
	}
	statuses := Statuses(a)
	required := []string{StatusF2MomentTargetExtracted, StatusF2LambdaInvariantDerived, StatusFailedF2MomentStillUnlocked, StatusFailedCosmologicalF4Firewall}
	for _, req := range required {
		found := false
		for _, got := range statuses {
			if got == req {
				found = true
				break
			}
		}
		if !found {
			t.Fatalf("missing status %s in %v", req, statuses)
		}
	}
	res := GravitationalSpectralActionF2CutoffMomentSieveTheorem().Run()
	if !res.Passed() {
		t.Fatalf("theorem did not pass:\n%s", res.Details())
	}
}
