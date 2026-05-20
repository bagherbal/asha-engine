package generation2chargedleptonrootextensionbranchchambermonodromyaudit

import (
	"strings"
	"testing"
)

func TestGate600SplittingFieldAndMonodromy(t *testing.T) {
	a, err := Build()
	if err != nil {
		t.Fatal(err)
	}
	if !a.Inherited.TraceRingDefined || !a.Inherited.CharacteristicPolynomial || !a.Inherited.AlgebraicExtension {
		t.Fatalf("bad inherited state: %+v", a.Inherited)
	}
	if !a.Splitting.Typed || a.Splitting.TraceRingOrdersRoots || !strings.Contains(a.Splitting.SplittingField, "lambda_1") {
		t.Fatalf("bad splitting field: %+v", a.Splitting)
	}
	if a.Monodromy.NativeBranchSelector || a.Monodromy.NativeOrdering || !strings.Contains(a.Monodromy.GenericMonodromy, "S3") {
		t.Fatalf("bad monodromy audit: %+v", a.Monodromy)
	}
}

func TestGate600FourthRootAndChamberSeal(t *testing.T) {
	a, err := Build()
	if err != nil {
		t.Fatal(err)
	}
	if a.FourthRoot.ComplexSheetsPerEigenvalue != 4 || !a.FourthRoot.PositiveRealBranchUnique || a.FourthRoot.FourthRootNative || a.FourthRoot.PositivityNative {
		t.Fatalf("bad fourth-root branch: %+v", a.FourthRoot)
	}
	if a.Chamber.NativeChamberSelector || a.Chamber.TraceRingSelectsWall || a.Chamber.DiscriminantSelectsWall || a.Chamber.MonodromySelectsOrder {
		t.Fatalf("bad chamber audit: %+v", a.Chamber)
	}
	if !a.BranchSeal.AlgebraicOverTrace || a.BranchSeal.Native || !a.BranchSeal.Environmental {
		t.Fatalf("bad branch seal: %+v", a.BranchSeal)
	}
}

func TestGate600TheoremAndFirewalls(t *testing.T) {
	res := Generation2ChargedLeptonRootExtensionBranchChamberMonodromyAuditTheorem().Verify()
	if !res.Passed() {
		t.Fatalf("theorem checks failed: %+v", res)
	}
	joined := strings.Join(res.Notes, "\n")
	for _, want := range []string{StatusGate599Inherited, StatusCubicSplittingFieldTyped, StatusDiscriminantDefined, StatusTraceRingNoOrdering, StatusNoNativePositiveFourthRoot, StatusNoNativeElectronWall, StatusRootBranchChamberSealDefined, StatusBFlavEnvironmental, StatusGate352Preserved, StatusGate596Preserved, StatusGate600Boundary} {
		if !strings.Contains(joined, want) {
			t.Fatalf("missing status %s", want)
		}
	}
}
