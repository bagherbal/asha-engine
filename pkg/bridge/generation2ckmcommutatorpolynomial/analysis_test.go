package generation2ckmcommutatorpolynomial

import (
	"math"
	"strings"
	"testing"
)

func TestBuildDefaultGate487CKMCommutatorPolynomialAudit(t *testing.T) {
	a, err := BuildDefault()
	if err != nil {
		t.Fatal(err)
	}
	if math.Abs(a.Spectrum.RatioRoverS-math.Sqrt2) > rankTolerance || math.Abs(a.Spectrum.MinkowskiResidual) > rankTolerance {
		t.Fatalf("expected null-C3 spectrum: %+v", a.Spectrum)
	}
	if a.Operators.NativeUpOperatorDerived || a.Operators.NativeDownOperatorDerived || a.Operators.NativeDiagonalizersDerived || a.Operators.NullBoundaryConstrainsEigenbasis {
		t.Fatalf("unexpected native operators/eigenbasis: %+v", a.Operators)
	}
	if !a.Sieve.RankVariabilityObserved || !a.Sieve.ZeroCommutatorPossible || !a.Sieve.RankTwoCommutatorPossible || !a.Sieve.RankThreeCommutatorPossible {
		t.Fatalf("expected variable commutator ranks under same spectrum: %+v", a.Sieve)
	}
	if a.Sieve.CommutatorRankSuppressedByNull || a.Sieve.JarlskogDeterminantLocked || a.Sieve.InvariantPolynomialProduced {
		t.Fatalf("illegal invariant theorem promotion: %+v", a.Sieve)
	}
	if a.Constraints.DerivedIndependentConstraints != 0 || a.Constraints.TwoConstraintTheoremPassed {
		t.Fatalf("unexpected CKM constraints: %+v", a.Constraints)
	}
	if a.Firewall.ObservedCKMImported || a.Firewall.ObservedWolfensteinImported || a.Firewall.ObservedQuarkMassesImported || a.Firewall.PolynomialConstraintsNativeWrite || a.Firewall.NativeRegistryWritten {
		t.Fatalf("firewall leak: %+v", a.Firewall)
	}
}

func TestRenderAuditGate487(t *testing.T) {
	a, err := BuildDefault()
	if err != nil {
		t.Fatal(err)
	}
	md := RenderAudit(a)
	for _, want := range []string{
		"# Gate 487 Registry Audit",
		StatusCommutatorRankNotSuppressed,
		StatusNoJarlskogPolynomialDerived,
		"[O_u,O_d]",
		"rephasing-invariant",
		"Gate 488",
	} {
		if !strings.Contains(md, want) {
			t.Fatalf("audit missing %q", want)
		}
	}
}
