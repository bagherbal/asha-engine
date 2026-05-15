package generation2ewnquotient

import (
	"math"
	"strings"
	"testing"
)

func TestGate502ScalarNormalizationIndependentElectroweakQuotient(t *testing.T) {
	a, err := BuildDefault()
	if err != nil {
		t.Fatalf("BuildDefault: %v", err)
	}
	if !a.Inheritance.ScalarNormalizationSealed || !a.Inheritance.ScalarNormalizationIndependentRoute || !a.Inheritance.PhotonKernelAvailable || !a.Inheritance.BrokenOrbitRankThreeAvailable || !a.Inheritance.DimensionlessDiag114Candidate {
		t.Fatalf("bad inheritance: %+v", a.Inheritance)
	}
	if !a.Quotient.ScalarNormalizationRemoved || !a.Quotient.YukawaTraceRemoved || !a.Quotient.HiggsVEVRemoved || !a.Quotient.OnlyDimensionlessStatements || a.Quotient.PhysicalMassStatementsAllowed || a.Quotient.PhysicalCouplingStatementsAllowed {
		t.Fatalf("bad quotient definition: %+v", a.Quotient)
	}
	if a.KernelRank.PhotonKernelDimension != 1 || a.KernelRank.BrokenOrbitRank != 3 || !a.KernelRank.PhotonKernelSurvivesScaleQuotient || !a.KernelRank.BrokenRankSurvivesScaleQuotient || a.KernelRank.RadialModeAfterQuotient != 1 || a.KernelRank.NativeKernelIndexClosed {
		t.Fatalf("kernel/rank quotient over-promoted or missing: %+v", a.KernelRank)
	}
	if !a.Hessian.ChargedPairDegenerate || !a.Hessian.Diag114Shape || !a.Hessian.DimensionlessShapeSurvives || math.Abs(a.Hessian.NeutralToChargedRatio-4) > eps || a.Hessian.KappaNative || a.Hessian.WeakAngleDerived || a.Hessian.GaugeCouplingsDerived || a.Hessian.PhysicalWZMassMatrixDerived || a.Hessian.ObservedWZMassRatioClaimed {
		t.Fatalf("Hessian quotient over-promoted or missing: %+v", a.Hessian)
	}
	if !a.Boundary.BridgeQuotientAccepted || a.Boundary.NativeElectroweakActionClosed || a.Boundary.NativeScalarNormalizationClosed || a.Boundary.NativeKappaClosed || a.Boundary.NativeWeakAngleClosed || a.Boundary.NativeGaugeCouplingsClosed || a.Boundary.NativeHiggsVEVClosed || a.Boundary.NativeWZMassMatrixClosed || a.Boundary.NativeMassRatioClosed {
		t.Fatalf("boundary over-promoted: %+v", a.Boundary)
	}
	if a.Firewall.YukawaTraceValueImported || a.Firewall.ObservedWMassImported || a.Firewall.ObservedZMassImported || a.Firewall.ObservedHiggsVEVImported || a.Firewall.ObservedWeakAngleImported || a.Firewall.ObservedGaugeCouplingImported || a.Firewall.ObservedWZMassRatioImported || a.Firewall.NativeWZMassWritten || a.Firewall.NativeMassRatioWritten {
		t.Fatalf("firewall violation: %+v", a.Firewall)
	}
	if a.Next.Gate != 503 {
		t.Fatalf("expected Gate503 redirect, got %+v", a.Next)
	}
	md := Markdown(a)
	for _, want := range []string{"# Gate 502 Registry Audit", "diag(1,1,4)", StatusBridgeQuotientAccepted, StatusFailedWZMassMatrixStillBlocked, "Gate 503"} {
		if !strings.Contains(md, want) {
			t.Fatalf("markdown missing %q", want)
		}
	}
}

func TestGate502TheoremPasses(t *testing.T) {
	res := Generation2ScalarNormalizationIndependentElectroweakQuotientAuditTheorem().Run()
	if !res.Passed() {
		t.Fatalf("theorem failed:\n%s", res.Details())
	}
}
