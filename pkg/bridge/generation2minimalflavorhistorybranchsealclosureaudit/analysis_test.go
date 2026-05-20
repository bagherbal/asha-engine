package generation2minimalflavorhistorybranchsealclosureaudit

import (
	"strings"
	"testing"
)

func TestGate604MinimalSeal(t *testing.T) {
	a, err := Build()
	if err != nil {
		t.Fatal(err)
	}
	if !a.Inherited.SelectsElectronWall || !a.Inherited.SelectsP3Nu || !a.Inherited.SelectsPositiveJ {
		t.Fatalf("bad inherited branch selection: %+v", a.Inherited)
	}
	if a.MinimalSeal.IsNative || !a.MinimalSeal.IsEnvironmental {
		t.Fatalf("minimal seal should be environmental: %+v", a.MinimalSeal)
	}
	if !contains(a.MinimalSeal.SelectedByBFlav, "third neutrino projector P_3^nu") {
		t.Fatalf("missing P3 in minimal seal: %+v", a.MinimalSeal)
	}
	if !contains(a.MinimalSeal.NotIncluded, "signed Vandermonde orientation") {
		t.Fatalf("signed Vandermonde should not be required for B_flav: %+v", a.MinimalSeal)
	}
}

func TestGate604OptionalFullOrder(t *testing.T) {
	a, err := Build()
	if err != nil {
		t.Fatal(err)
	}
	if a.OptionalFullOrder.RequiredForBFlav {
		t.Fatalf("full-order seal should not be required for B_flav: %+v", a.OptionalFullOrder)
	}
	if !a.OptionalFullOrder.RequiredForFullOrder {
		t.Fatalf("full-order seal should be required for full ordered history: %+v", a.OptionalFullOrder)
	}
	if a.OptionalFullOrder.NativeTheoremPresent {
		t.Fatalf("no native signed discriminant theorem should be present")
	}
	if a.Firewalls.DerivesBFlavZero || a.Firewalls.AddsSelector {
		t.Fatalf("firewall breach: %+v", a.Firewalls)
	}
}

func TestGate604TheoremStatuses(t *testing.T) {
	res := Generation2MinimalFlavorHistoryBranchSealClosureAuditTheorem().Verify()
	if !res.Passed() {
		t.Fatalf("theorem checks failed: %+v", res)
	}
	joined := strings.Join(res.Notes, "\n")
	for _, want := range []string{StatusFlavorHistoryBranchStackConstructed, StatusMinimalFlavorHistoryBranchSealDefined, StatusSigmaGaugeLikeForBFlav, StatusOptionalDiscriminantSeal, StatusBFlavBranchCompatibilityFilter, StatusNoNativeBranchSelectionTheorem, StatusGate604Boundary} {
		if !strings.Contains(joined, want) {
			t.Fatalf("missing status %s", want)
		}
	}
}
