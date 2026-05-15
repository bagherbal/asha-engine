package generation2ckmnullmirror

import (
	"strings"
	"testing"
)

func TestBuildDefaultGate486CKMNullMirrorAudit(t *testing.T) {
	a, err := BuildDefault()
	if err != nil {
		t.Fatal(err)
	}
	if !a.Geometry.CoordinateChartBridgeOnly || a.Geometry.RelativeCoordinateChartDim != 2 {
		t.Fatalf("expected bridge two-coordinate chart: %+v", a.Geometry)
	}
	if a.Geometry.CKMFourToTwoForcedByCone || a.Geometry.CKMEigenbasisMismatchDerived {
		t.Fatalf("illegal native CKM promotion: %+v", a.Geometry)
	}
	if a.Rephasing.DerivedIndependentConstraints != 0 || a.Rephasing.RephasingInvariantConstraintsOK {
		t.Fatalf("unexpected invariant constraints: %+v", a.Rephasing)
	}
	if a.Operators.NativeUpOperatorDerived || a.Operators.NativeDownOperatorDerived || a.Operators.CKMAsUuDaggerUdConstructed || a.Operators.InvariantPolynomialProduced {
		t.Fatalf("unexpected native CKM operators: %+v", a.Operators)
	}
	if a.Firewall.ObservedCKMImported || a.Firewall.CKMFourToTwoNativeWritten || a.Firewall.NativeRegistryWritten || !a.Firewall.NullMirrorSocketBridgeWritten {
		t.Fatalf("firewall leak: %+v", a.Firewall)
	}
}

func TestRenderAuditGate486(t *testing.T) {
	a, err := BuildDefault()
	if err != nil {
		t.Fatal(err)
	}
	md := RenderAudit(a)
	for _, want := range []string{
		"# Gate 486 Registry Audit",
		StatusNullMirrorCoordinateChartFound,
		StatusCKMNativeTheoremNotProven,
		"V_CKM = U_u^† U_d",
		"rephasing-invariant",
		"Gate 487",
	} {
		if !strings.Contains(md, want) {
			t.Fatalf("audit missing %q", want)
		}
	}
}
