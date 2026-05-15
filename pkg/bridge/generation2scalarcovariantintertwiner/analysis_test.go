package generation2scalarcovariantintertwiner

import (
	"strings"
	"testing"
)

func TestBuildDefaultGate492ScalarCovariantIntertwiner(t *testing.T) {
	a, err := BuildDefault()
	if err != nil {
		t.Fatal(err)
	}
	if !a.Inheritance.Gate491ScalarEdgeStability || !a.Inheritance.NativeGaugeEatingPreviouslyOpen {
		t.Fatalf("bad inheritance: %+v", a.Inheritance)
	}
	if !a.Representation.AbstractDoubletRepresentation || a.Representation.FullSU2SelectedByScalarData || a.Representation.CanonicalComplexStructure {
		t.Fatalf("scalar SU(2) should remain bridge-level: %+v", a.Representation)
	}
	if !a.Dphi.AbstractTemplateAvailable || !a.Dphi.DimensionlessWZPhotonSignature || a.Dphi.MassMatrixRank != 3 || a.Dphi.NativeDphiDerived || a.Dphi.PhysicalMassesDerived {
		t.Fatalf("bad Dphi audit: %+v", a.Dphi)
	}
	if !a.Intertwiner.BrokenImagesIndependent || a.Intertwiner.BrokenImageRank != 3 || !a.Intertwiner.GoldstoneImageDiagnostic || a.Intertwiner.CanonicalProtectedToBrokenDerived {
		t.Fatalf("bad intertwiner audit: %+v", a.Intertwiner)
	}
	if !a.Photon.QEMAnnihilatesVacuum || a.Photon.WeakMixingAngleDerived || a.Photon.PhotonPhysicallyNormalized {
		t.Fatalf("bad photon audit: %+v", a.Photon)
	}
	if !a.Boundary.DiagnosticSocketPromotable || a.Boundary.NativeIntertwinerDerived || a.Boundary.PhysicalMassMatrixDerived {
		t.Fatalf("native promotion should be blocked: %+v", a.Boundary)
	}
	if a.Firewall.ObservedWMassImported || a.Firewall.NativeWZMassWritten || a.Firewall.NativeWeakAngleWritten {
		t.Fatalf("firewall leak: %+v", a.Firewall)
	}
}

func TestGate492RenderAudit(t *testing.T) {
	a, err := BuildDefault()
	if err != nil {
		t.Fatal(err)
	}
	md := RenderAudit(a)
	for _, want := range []string{
		"# Gate 492 Registry Audit",
		StatusAbstractDphiTemplateFound,
		StatusGoldstoneImageDiagnosticFound,
		StatusPhotonExemptionDiagnostic,
		StatusFailedNativeDphiNotDerived,
		"Gate 493",
		"Qem",
	} {
		if !strings.Contains(md, want) {
			t.Fatalf("audit missing %q", want)
		}
	}
}
