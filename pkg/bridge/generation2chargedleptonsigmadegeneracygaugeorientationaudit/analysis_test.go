package generation2chargedleptonsigmadegeneracygaugeorientationaudit

import (
	"math"
	"strings"
	"testing"
)

func TestGate603S3Degeneracy(t *testing.T) {
	a, err := Build()
	if err != nil {
		t.Fatal(err)
	}
	if len(a.S3Action) != 6 {
		t.Fatalf("expected 6 S3 rows, got %d", len(a.S3Action))
	}
	eps := a.S3Action[0].ElectronWallEpsilonRad
	signs := map[int]bool{}
	for _, row := range a.S3Action {
		if math.Abs(row.ElectronWallEpsilonRad-eps) > 5e-14 {
			t.Fatalf("electron wall epsilon should be sigma-degenerate: %+v vs %.18g", row, eps)
		}
		signs[row.SignVandermondeX] = true
		if math.Abs(row.BFlavInvariantValue-a.S3Action[0].BFlavInvariantValue) > 5e-14 {
			t.Fatalf("B_flav should not see cyclic sigma: %+v", row)
		}
	}
	if !signs[+1] || !signs[-1] {
		t.Fatalf("expected both Vandermonde signs: %v", signs)
	}
}

func TestGate603SealAndFirewalls(t *testing.T) {
	a, err := Build()
	if err != nil {
		t.Fatal(err)
	}
	if !a.MinimalRemaining.SigmaGaugeForBFlav || !a.MinimalRemaining.PhysicalFullOrderingRequiresSeal {
		t.Fatalf("bad minimal remaining statement: %+v", a.MinimalRemaining)
	}
	if a.SignedDiscriminant.NativeSignedVTheorem {
		t.Fatalf("signed discriminant theorem should not be native")
	}
	if a.FourierCyclic.BFlavDependsOnCyclicOrientation {
		t.Fatalf("B_flav should not depend on cyclic sigma")
	}
	if a.Firewalls.DerivesKoide || a.Firewalls.DerivesBFlavZero || a.Firewalls.AddsSelector {
		t.Fatalf("firewall breach: %+v", a.Firewalls)
	}
}

func TestGate603TheoremStatuses(t *testing.T) {
	res := Generation2ChargedLeptonSigmaDegeneracyGaugeOrientationAuditTheorem().Verify()
	if !res.Passed() {
		t.Fatalf("theorem checks failed: %+v", res)
	}
	joined := strings.Join(res.Notes, "\n")
	for _, want := range []string{StatusSigmaDegeneracySourceIdentified, StatusSigmaGaugeRedundancy, StatusDiscriminantOrientationSealRequired, StatusBFlavDoesNotSeeCyclicSigma, StatusNoNativeSignedDiscriminantTheorem, StatusGate603Boundary} {
		if !strings.Contains(joined, want) {
			t.Fatalf("missing status %s", want)
		}
	}
}
