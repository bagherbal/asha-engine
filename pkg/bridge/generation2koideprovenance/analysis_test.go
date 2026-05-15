package generation2koideprovenance

import (
	"math"
	"strings"
	"testing"
)

func TestBuildDefaultGate485KoideProvenance(t *testing.T) {
	a, err := BuildDefault()
	if err != nil {
		t.Fatal(err)
	}
	if !a.Basis.DemocraticPhaseOrthogonality || !a.Basis.PhaseNormEqualsThreeHalves || !a.Basis.NoEmpiricalMassesUsed {
		t.Fatalf("bad C3 basis proof: %+v", a.Basis)
	}
	if !a.Derivation.NullForcesRatio || math.Abs(a.Derivation.RatioDerived-math.Sqrt2) > 1e-12 {
		t.Fatalf("bad null ratio derivation: %+v", a.Derivation)
	}
	if !a.Derivation.KoideEquivalent || math.Abs(a.Derivation.KoideQDerived-2.0/3.0) > 1e-12 {
		t.Fatalf("bad Koide derivation: %+v", a.Derivation)
	}
	if a.Collapse.AbsoluteMassesDerived || a.Collapse.FullFlavorModuliCollapsed || a.Collapse.CollapsedShapeDOF != 1 {
		t.Fatalf("bad collapse/firewall boundary: %+v", a.Collapse)
	}
	if a.Firewall.ObservedMassImportedForProof || a.Firewall.KoideAsPhysicalMassPrediction || a.Firewall.CKMMatrixConstructed || a.Firewall.PMNSMatrixConstructed || a.Firewall.NativeRegistryWritten {
		t.Fatalf("firewall leak: %+v", a.Firewall)
	}
}

func TestRenderAuditGate485(t *testing.T) {
	a, err := BuildDefault()
	if err != nil {
		t.Fatal(err)
	}
	md := RenderAudit(a)
	for _, want := range []string{StatusNullKoideRatioDerived, "R/S = sqrt(2)", "Q = 2/3", "absolute masses", "CKM/PMNS", StatusFailedFullFlavorCollapseRejected} {
		if !strings.Contains(md, want) {
			t.Fatalf("audit missing %q", want)
		}
	}
}
