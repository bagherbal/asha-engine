package generation2bordismcobordismclassifierairlock

import (
	"strings"
	"testing"
)

func TestGate521BordismClassifierAirlock(t *testing.T) {
	a, err := BuildDefault()
	if err != nil {
		t.Fatalf("BuildDefault failed: %v", err)
	}
	if !a.Inheritance.Gate520FileAdapterDefined || !a.Inheritance.Gate520SyntheticOnly || a.Inheritance.Gate520NativePrediction || !a.Inheritance.Gate520NativeWriteBlocked {
		t.Fatalf("bad Gate520 inheritance: %s", FormatInheritance(a.Inheritance))
	}
	if !a.Socket.OrientedSocket || !a.Socket.SpinSocket || !a.Socket.SpinCSocket || !a.Socket.BoundaryBordismSocket {
		t.Fatalf("missing bordism sockets: %s", FormatSocket(a.Socket))
	}
	if !a.Socket.RequiresW1ZeroForOriented || !a.Socket.RequiresW2ZeroForSpin || !a.Socket.RequiresW3ZeroForSpinC || !a.Socket.RequiresC1Mod2EqualsW2ForSpinC {
		t.Fatalf("missing obstruction constraints: %s", FormatSocket(a.Socket))
	}
	if a.Socket.SelectsSpecificClass || a.Socket.SelectsManifoldRepresentative {
		t.Fatalf("classifier selected topology: %s", FormatSocket(a.Socket))
	}
	if !nearly(a.Constraints.SignatureP1Residual, 0, 1e-12) || !a.Constraints.SpinDivisibilityPassed || !nearly(a.Constraints.SyntheticAHat, 2, 1e-12) {
		t.Fatalf("bad characteristic checks: %s", FormatConstraints(a.Constraints))
	}
	if a.Constraints.GlobalNumbersDerived || a.Constraints.PhysicalThetaSelected {
		t.Fatalf("illegal global derivation: %s", FormatConstraints(a.Constraints))
	}
	if !a.Scale.ClassifierScaleFree || a.Scale.UsesLambda || a.Scale.UsesNewton || a.Scale.UsesObservedTopology || a.Scale.UsesBoundarySpectrum {
		t.Fatalf("scale firewall failed: %s", FormatScale(a.Scale))
	}
	if !a.Rejection.NativeRegistryWriteBlocked || !a.Rejection.SpecificBordismClassNativeBlocked || !a.Rejection.EtaInvariantNativeBlocked {
		t.Fatalf("rejection failed: %s", FormatRejection(a.Rejection))
	}
}

func TestGate521Markdown(t *testing.T) {
	a, err := BuildDefault()
	if err != nil {
		t.Fatalf("BuildDefault failed: %v", err)
	}
	md := Markdown(a)
	for _, want := range []string{
		"# Gate 521 Registry Audit — Bordism and Cobordism Classifier Airlock",
		StatusBordismLedgerDefined,
		StatusFailedSpecificClassNotSelected,
		"Gate 522 — Bordism Comparator File Adapter and Stiefel-Whitney Metadata Firewall",
	} {
		if !strings.Contains(md, want) {
			t.Fatalf("markdown missing %q\n%s", want, md)
		}
	}
}

func TestGate521Theorem(t *testing.T) {
	th := Generation2BordismCobordismClassifierAirlockTheorem()
	res := th.Verify()
	if len(res.Checks) != 5 {
		t.Fatalf("expected 5 checks, got %d", len(res.Checks))
	}
	for _, c := range res.Checks {
		if !c.Passed {
			t.Fatalf("check failed: %s: %s", c.Name, c.Detail)
		}
	}
}
