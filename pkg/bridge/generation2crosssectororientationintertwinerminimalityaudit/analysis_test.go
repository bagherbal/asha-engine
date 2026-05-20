package generation2crosssectororientationintertwinerminimalityaudit

import (
	"strings"
	"testing"
)

func TestGate592DefinesMinimalOrientationBalanceSeal(t *testing.T) {
	a, err := Build()
	if err != nil {
		t.Fatal(err)
	}
	if a.Seal.Name != "OrientationBalanceSeal" || a.Seal.Native {
		t.Fatalf("expected non-native OrientationBalanceSeal: %+v", a.Seal)
	}
	if a.Seal.KappaCandidate != a.Inherited.OrientationCandidate || a.Seal.KappaResidual != a.Inherited.Delta590 {
		t.Fatalf("seal must inherit Gate590 relation exactly: seal=%+v inherited=%+v", a.Seal, a.Inherited)
	}
	if a.Precision.AdditionalCorrectionJustified {
		t.Fatalf("Gate592 must not justify an additional residual fit at v1 precision: %+v", a.Precision)
	}
}

func TestGate592RejectsNativeIntertwiner(t *testing.T) {
	a, err := Build()
	if err != nil {
		t.Fatal(err)
	}
	if a.Repository.AnyNativeCrossSectorIntertwiner || a.Lawfulness.CrossSectorOrientationIntertwinerPresent || a.Lawfulness.NativeRootTraceOperatorPresent || a.Lawfulness.AbsoluteDiracObservablePresent {
		t.Fatalf("unexpected native intertwiner/root operator: repo=%+v law=%+v", a.Repository, a.Lawfulness)
	}
	if !a.Final.KappaRemainsEnvironmental || a.Final.NativeIntertwinerPresent {
		t.Fatalf("kappa must remain environmental: %+v", a.Final)
	}
}

func TestGate592TheoremAndStatuses(t *testing.T) {
	th := Generation2CrossSectorOrientationIntertwinerMinimalityAuditTheorem()
	res := th.Verify()
	if !res.Passed() {
		t.Fatalf("theorem checks failed: %+v", res)
	}
	joined := strings.Join(res.Notes, "\n")
	for _, want := range []string{StatusOrientationBalanceSealDefined, StatusNoCrossSectorIntertwiner, StatusNoNativeRootTraceAbsoluteDirac, StatusNoDeltaFitJustified, StatusGate592Boundary} {
		if !strings.Contains(joined, want) {
			t.Fatalf("missing status %s", want)
		}
	}
}
