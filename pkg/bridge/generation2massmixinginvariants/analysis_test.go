package generation2massmixinginvariants

import "testing"

func TestBuildDefault(t *testing.T) {
	a, err := BuildDefault()
	if err != nil {
		t.Fatalf("BuildDefault failed: %v", err)
	}
	if a.Sieve.UniqueMassAngleInvariant || !a.Firewall.RatioPredictionSealed {
		t.Fatalf("Gate 450 must seal ratio prediction: %s", FormatSieve(a.Sieve))
	}
	if !a.Identity.SumRuleExact || a.Identity.SpecificMassAngleRatio {
		t.Fatalf("texture-zero identity must be exact but not pairwise ratio: %s", FormatIdentity(a.Identity))
	}
}

func TestSymbolicCounterexamples(t *testing.T) {
	s := buildRatioSieve()
	if !s.SameAngleDifferentMassShape {
		t.Fatalf("expected same-angle/different-mass counterexample: %s", FormatSieve(s))
	}
	if !s.SameMassShapeDifferentAngle {
		t.Fatalf("expected same-mass/different-angle counterexample: %s", FormatSieve(s))
	}
	for _, x := range s.Counterexamples {
		if !x.BoundaryCompatible || x.ImportsEmpiricalData {
			t.Fatalf("bad witness: %s", FormatCounterexample(x))
		}
	}
}

func TestCharacteristicPolynomialSamples(t *testing.T) {
	x := newCounterexample("X branch", 0, 1, 0, "pure X")
	if !nearlyEqual(x.D, 2, 1e-12) || !nearlyEqual(x.P, 3, 1e-12) {
		t.Fatalf("unexpected X-branch invariants: %s", FormatCounterexample(x))
	}
	want := []float64{-0.5773502691896258, -0.5773502691896258, 1.1547005383792517}
	if !sameNormalizedEigenvalues(x.NormalizedEigenvalues, want, 1e-10) {
		t.Fatalf("unexpected normalized X eigenvalues: got %v want %v", x.NormalizedEigenvalues, want)
	}
}

func TestTheoremPassesAsFailedRouteAudit(t *testing.T) {
	res := Generation2StructuralZeroMassMixingInvariantRatioSieveTheorem().Run()
	if !res.Passed() {
		t.Fatalf("theorem audit checks did not pass:\n%s", res.Details())
	}
	if string(res.Status) != "FAILED_ROUTE" {
		t.Fatalf("Gate 450 should be a failed-route audit, got %s", res.Status)
	}
}

func TestRenderAuditContainsKeyStatuses(t *testing.T) {
	a, err := BuildDefault()
	if err != nil {
		t.Fatal(err)
	}
	md := RenderAudit(a)
	for _, want := range []string{StatusFailedRatiosRequireExactAmplitudes, StatusTextureZeroSumRuleDerived, "chi(lambda)=lambda^3", "same normalized mass spectrum while changing the local mixing angle"} {
		if !stringsContains(md, want) {
			t.Fatalf("audit missing %q", want)
		}
	}
}

func stringsContains(s, sub string) bool {
	for i := 0; i+len(sub) <= len(s); i++ {
		if s[i:i+len(sub)] == sub {
			return true
		}
	}
	return false
}
