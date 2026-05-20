package generation2flavorspectralorientationbalancefunctionalaudit

import (
	"math"
	"strings"
	"testing"
)

func TestGate594ConstructsBFlav(t *testing.T) {
	a, err := Build()
	if err != nil {
		t.Fatal(err)
	}
	if !a.Final.AllTermsInOneSpectralAlgebra || a.Algebra.Name != "A_flav" {
		t.Fatalf("expected common flavor spectral algebra: algebra=%+v final=%+v", a.Algebra, a.Final)
	}
	if math.Abs(a.Balance.BFlav+a.Balance.Delta590) > 1e-18 {
		t.Fatalf("expected B_flav=-Delta590: balance=%+v", a.Balance)
	}
	if !strings.Contains(a.Balance.Definition, "epsilon(H_e)") || !strings.Contains(a.Balance.Definition, "J(H_u,H_d)") {
		t.Fatalf("B_flav definition missing spectral terms: %s", a.Balance.Definition)
	}
}

func TestGate594RejectsNativePromotion(t *testing.T) {
	a, err := Build()
	if err != nil {
		t.Fatal(err)
	}
	if a.Availability.NativeBFlavZeroTheorem || a.Target.NativePresent || a.Final.NativeBFlavOperatorPresent {
		t.Fatalf("unexpected native B_flav theorem/operator: availability=%+v target=%+v final=%+v", a.Availability, a.Target, a.Final)
	}
	if !a.ChargedLepton.Gate352ObstructionPreserved || a.ChargedLepton.NativeFunctionalPresent {
		t.Fatalf("root-spectrum/Gate352 firewall failed: %+v", a.ChargedLepton)
	}
}

func TestGate594TheoremAndStatuses(t *testing.T) {
	th := Generation2FlavorSpectralOrientationBalanceFunctionalAuditTheorem()
	res := th.Verify()
	if !res.Passed() {
		t.Fatalf("theorem checks failed: %+v", res)
	}
	joined := strings.Join(res.Notes, "\n")
	for _, want := range []string{StatusFlavorSpectralAlgebraDefined, StatusBalanceFunctionalConstructed, StatusNumericReproducesGate590593, StatusNoBFlavZeroTheorem, StatusBFlavEnvironmental, StatusGate594Boundary} {
		if !strings.Contains(joined, want) {
			t.Fatalf("missing status %s", want)
		}
	}
}
