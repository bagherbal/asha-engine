package generation2koidereactorckmorientationcombinationaudit

import (
	"strings"
	"testing"
)

func TestGate590CombinationImprovesKappaResidual(t *testing.T) {
	a, err := Build()
	if err != nil {
		t.Fatal(err)
	}
	if !a.Combination.BOutperformsA {
		t.Fatal("expected reactor-minus-CKM candidate to outperform reactor quarter")
	}
	if a.Combination.BImprovementFactor <= 10 {
		t.Fatalf("expected >10x improvement, got %.12g", a.Combination.BImprovementFactor)
	}
	if a.Combination.BReactorMinusCKM.AbsResidual >= 3e-6 {
		t.Fatalf("combined residual too large: %.12g", a.Combination.BReactorMinusCKM.AbsResidual)
	}
}

func TestGate590InverseTheta13InsideOneSigma(t *testing.T) {
	a, err := Build()
	if err != nil {
		t.Fatal(err)
	}
	if !a.Inverse.WithinSin2OneSigma || !a.Inverse.WithinThetaOneSigma {
		t.Fatalf("expected inverse prediction inside one sigma: %+v", a.Inverse)
	}
	if a.Inverse.Theta13PredDeg < 8.55 || a.Inverse.Theta13PredDeg > 8.57 {
		t.Fatalf("unexpected theta13 prediction %.15g", a.Inverse.Theta13PredDeg)
	}
}

func TestGate590Firewalls(t *testing.T) {
	a, err := Build()
	if err != nil {
		t.Fatal(err)
	}
	if a.Lawfulness.CrossSectorOrientationIntertwinerPresent || a.Lawfulness.NativeRootTraceOperatorPresent || !a.Final.KappaRemainsEnvironmental {
		t.Fatalf("firewall failure: law=%+v final=%+v", a.Lawfulness, a.Final)
	}
	th := Generation2KoideReactorCKMOrientationCombinationAuditTheorem()
	res := th.Verify()
	if !res.Passed() {
		t.Fatalf("theorem checks failed: %+v", res)
	}
	joined := strings.Join(res.Notes, "\n")
	for _, want := range []string{StatusBOutperformsA, StatusNoCrossSectorIntertwiner, StatusGate590Boundary} {
		if !strings.Contains(joined, want) {
			t.Fatalf("missing status %s", want)
		}
	}
}
