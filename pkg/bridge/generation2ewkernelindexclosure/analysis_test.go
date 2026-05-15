package generation2ewkernelindexclosure

import (
	"strings"
	"testing"
)

func TestGate503ElectroweakKernelIndexNativeClosure(t *testing.T) {
	a, err := BuildDefault()
	if err != nil {
		t.Fatalf("BuildDefault: %v", err)
	}
	if !a.Inheritance.QuotientBridgeAccepted || !a.Inheritance.QuotientPhotonKernel || !a.Inheritance.QuotientBrokenRankThree || !a.Inheritance.StructuralHiggsDoubletProvenance || !a.Inheritance.StructuralDphiSocket {
		t.Fatalf("bad inheritance: %+v", a.Inheritance)
	}
	if !a.Sieve.Executed || a.Sieve.ScalarRealDimension != 4 || a.Sieve.ComplexDoublets != 1 || !a.Sieve.AssumesNonzeroHiggsRay || a.Sieve.UsesVacuumScale || a.Sieve.UsesGaugeCouplings || a.Sieve.UsesObservedElectroweakData {
		t.Fatalf("bad representation sieve: %+v", a.Sieve)
	}
	if a.Kernel.StabilizerDimension != 1 || a.Kernel.BrokenOrbitDimension != 3 || a.Kernel.RadialQuotientDimension != 1 || !a.Kernel.PhotonKernelIndexProven || !a.Kernel.BrokenOrbitIndexProven || !a.Kernel.RadialIndexProven || !a.Kernel.ConditionalOnNonzeroRay || a.Kernel.UnconditionalNativeVacuumProvenance {
		t.Fatalf("kernel index over-promoted or missing: %+v", a.Kernel)
	}
	if !a.Hessian.KernelRankMatchesGate502 || !a.Hessian.Diag114ShapeInherited || a.Hessian.Diag114NativeHessian || a.Hessian.KappaNative || a.Hessian.WeakAngleDerived || a.Hessian.GaugeCouplingsDerived || a.Hessian.PhysicalWZMassMatrix || a.Hessian.ObservedMassRatioClaimed {
		t.Fatalf("Hessian compatibility over-promoted or missing: %+v", a.Hessian)
	}
	if !a.Boundary.ConditionalRepresentationIndexAccepted || a.Boundary.UnconditionalNativeElectroweakAction || a.Boundary.NativeNonzeroVacuumRaySelected || a.Boundary.NativeVacuumOrientationSelected || a.Boundary.NativeKappaSelected || a.Boundary.NativeWeakAngleDerived || a.Boundary.NativeWZMassMatrixDerived {
		t.Fatalf("boundary over-promoted: %+v", a.Boundary)
	}
	if a.Firewall.ObservedWMassImported || a.Firewall.ObservedZMassImported || a.Firewall.ObservedWZRatioImported || a.Firewall.ObservedWeakAngleImported || a.Firewall.ObservedGaugeCouplingImported || a.Firewall.ObservedHiggsVEVImported || a.Firewall.ObservedYukawaImported || a.Firewall.NativeKappaWritten || a.Firewall.NativeWZMassWritten {
		t.Fatalf("firewall violation: %+v", a.Firewall)
	}
	if a.Next.Gate != 504 {
		t.Fatalf("expected Gate504 redirect, got %+v", a.Next)
	}
	md := Markdown(a)
	for _, want := range []string{"# Gate 503 Registry Audit", "U(1)em", StatusKernelRankPromotedConditionally, StatusFailedWZMassMatrixStillBlocked, "Gate 504"} {
		if !strings.Contains(md, want) {
			t.Fatalf("markdown missing %q", want)
		}
	}
}

func TestGate503TheoremPasses(t *testing.T) {
	res := Generation2ElectroweakKernelIndexNativeClosureAuditTheorem().Run()
	if !res.Passed() {
		t.Fatalf("theorem failed:\n%s", res.Details())
	}
}
