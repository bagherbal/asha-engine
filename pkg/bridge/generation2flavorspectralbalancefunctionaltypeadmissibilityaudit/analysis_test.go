package generation2flavorspectralbalancefunctionaltypeadmissibilityaudit

import (
	"strings"
	"testing"
)

func TestGate595TypeAdmissibilityLocatesEpsilonObstruction(t *testing.T) {
	a, err := Build()
	if err != nil {
		t.Fatal(err)
	}
	if !a.Final.BFlavEnvironmentalWellDefined {
		t.Fatalf("expected B_flav environmental well-defined: %+v", a.Final)
	}
	if !a.TermTyping.Epsilon.RequiresFractional || a.TermTyping.Epsilon.NativePresent {
		t.Fatalf("epsilon(H_e) obstruction not captured: %+v", a.TermTyping.Epsilon)
	}
	if !a.Obstruction.EpsilonHEBlocked || !strings.Contains(a.Final.PrimaryNativeObstruction, "epsilon") {
		t.Fatalf("primary obstruction not epsilon(H_e): obstruction=%+v final=%+v", a.Obstruction, a.Final)
	}
}

func TestGate595ProjectorAndCommutatorAreOnlyConditionallyAdmissible(t *testing.T) {
	a, err := Build()
	if err != nil {
		t.Fatal(err)
	}
	if !a.Obstruction.PMNSMoreAdmissible || !a.Obstruction.CKMMoreAdmissible {
		t.Fatalf("expected PMNS/CKM terms more admissible as observed ledgers: %+v", a.Obstruction)
	}
	if a.TermTyping.PMNS.NativePresent || a.TermTyping.CKM.NativePresent {
		t.Fatalf("PMNS/CKM terms unexpectedly native: pmns=%+v ckm=%+v", a.TermTyping.PMNS, a.TermTyping.CKM)
	}
	if a.Final.NativeBFlavZeroTheoremPresent {
		t.Fatalf("unexpected native B_flav zero theorem: %+v", a.Final)
	}
}

func TestGate595TheoremAndStatuses(t *testing.T) {
	th := Generation2FlavorSpectralBalanceFunctionalTypeAdmissibilityAuditTheorem()
	res := th.Verify()
	if !res.Passed() {
		t.Fatalf("theorem checks failed: %+v", res)
	}
	joined := strings.Join(res.Notes, "\n")
	for _, want := range []string{StatusEnvironmentalWellDefined, StatusPrimaryObstructionEpsilon, StatusNoNativeFourthRootHE, StatusNoNativeBFlavZero, StatusGate595Boundary} {
		if !strings.Contains(joined, want) {
			t.Fatalf("missing status %s", want)
		}
	}
}
