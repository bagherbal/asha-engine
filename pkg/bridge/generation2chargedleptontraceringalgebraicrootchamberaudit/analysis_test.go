package generation2chargedleptontraceringalgebraicrootchamberaudit

import (
	"strings"
	"testing"
)

func TestGate599TraceRingAndCharacteristicPolynomial(t *testing.T) {
	a, err := Build()
	if err != nil {
		t.Fatal(err)
	}
	if !a.Inherited.NativeTraceCableVisible || !a.Inherited.RootOrientationMissing {
		t.Fatalf("bad inherited state: %+v", a.Inherited)
	}
	if !a.TraceRing.Admissible || len(a.TraceRing.Generators) != 3 {
		t.Fatalf("trace ring not defined: %+v", a.TraceRing)
	}
	if !a.Characteristic.BuiltFromTraceRing || !strings.Contains(a.Characteristic.Polynomial, "lambda^3") {
		t.Fatalf("bad characteristic polynomial: %+v", a.Characteristic)
	}
}

func TestGate599RootExtensionAndChamber(t *testing.T) {
	a, err := Build()
	if err != nil {
		t.Fatal(err)
	}
	if !a.RootExtension.AlgebraicOverTraceRing || !a.RootExtension.RequiresFourthRoot || a.RootExtension.Native || a.RootExtension.AvoidsGate596Obstruction {
		t.Fatalf("bad root extension: %+v", a.RootExtension)
	}
	if !a.Chamber.AlgebraicOverRootExt || !a.Chamber.RequiresChamberSeal || a.Chamber.NativePolynomial {
		t.Fatalf("bad chamber functional: %+v", a.Chamber)
	}
	if !a.Epsilon.AlgebraicOverTraceRing || a.Epsilon.NativePolynomial || a.Epsilon.PurelyRawInsertion {
		t.Fatalf("bad epsilon status: %+v", a.Epsilon)
	}
}

func TestGate599TheoremAndFirewalls(t *testing.T) {
	res := Generation2ChargedLeptonTraceRingAlgebraicRootChamberAuditTheorem().Verify()
	if !res.Passed() {
		t.Fatalf("theorem checks failed: %+v", res)
	}
	joined := strings.Join(res.Notes, "\n")
	for _, want := range []string{StatusNativeTraceRingDefined, StatusCharacteristicPolynomialDefined, StatusEpsilonAlgebraicOverTraceRing, StatusEpsilonNotPolynomialInvariant, StatusNoNativeHEOneFourthTheorem, StatusDoesNotAvoidGate596, StatusNoNativeBFlavZeroTheorem, StatusGate352Preserved, StatusGate596Preserved, StatusGate599Boundary} {
		if !strings.Contains(joined, want) {
			t.Fatalf("missing status %s", want)
		}
	}
}
