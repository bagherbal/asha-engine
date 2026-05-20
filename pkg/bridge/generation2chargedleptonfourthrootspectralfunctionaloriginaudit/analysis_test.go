package generation2chargedleptonfourthrootspectralfunctionaloriginaudit

import (
	"strings"
	"testing"
)

func TestGate596RootFunctionalRequiresFourthRootAndChamber(t *testing.T) {
	a, err := Build()
	if err != nil {
		t.Fatal(err)
	}
	if !a.RootFunctional.EnvironmentalWellDefined || !a.RootFunctional.RequiresFourthRoot || !a.RootFunctional.RequiresOrderedChamber {
		t.Fatalf("root functional not typed correctly: %+v", a.RootFunctional)
	}
	if a.RootFunctional.NativePresent || a.Final.NativeFourthRootPresent {
		t.Fatalf("unexpected native fourth-root promotion: root=%+v final=%+v", a.RootFunctional, a.Final)
	}
}

func TestGate596RoutesAndSeal(t *testing.T) {
	a, err := Build()
	if err != nil {
		t.Fatal(err)
	}
	if a.Routes.AnyNativeRoute {
		t.Fatalf("unexpected native route: %+v", a.Routes)
	}
	if a.Seal.Name != "ChargedLeptonRootChamberSeal" || !a.Seal.MayEnterBFlav || a.Seal.NativeLaw {
		t.Fatalf("minimal seal not defined correctly: %+v", a.Seal)
	}
	if !strings.Contains(a.Routes.Verdict, StatusNoNativeSpectralZeta) || !strings.Contains(a.Routes.Verdict, StatusNoAbsoluteDirac) {
		t.Fatalf("route obstruction statuses missing: %s", a.Routes.Verdict)
	}
}

func TestGate596TheoremAndStatuses(t *testing.T) {
	th := Generation2ChargedLeptonFourthRootSpectralFunctionalOriginAuditTheorem()
	res := th.Verify()
	if !res.Passed() {
		t.Fatalf("theorem checks failed: %+v", res)
	}
	joined := strings.Join(res.Notes, "\n")
	for _, want := range []string{StatusEpsilonEnvironmentalWellDefined, StatusNoNativeFourthRoot, StatusNoNativeRootTrace, StatusMinimalSealDefined, StatusGate596Boundary} {
		if !strings.Contains(joined, want) {
			t.Fatalf("missing status %s", want)
		}
	}
}
