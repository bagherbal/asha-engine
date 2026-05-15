package generation2ewresidualgeometryairlock

import (
	"strings"
	"testing"
)

func TestGate508ResidualGeometryAirlock(t *testing.T) {
	a, err := BuildDefault()
	if err != nil {
		t.Fatalf("BuildDefault: %v", err)
	}
	if !a.Inheritance.Gate507FileLoaded || !a.Inheritance.Gate507SyntheticOnly || a.Inheritance.Gate507ObservedNumbersImported || !a.Inheritance.Gate507AdapterExecuted || !a.Inheritance.Gate507ResidualsComputed || !a.Inheritance.Gate507ResidualsAllZero || !a.Inheritance.Gate507NativeWriteBlocked || !a.Inheritance.Gate502QuotientAccepted || !a.Inheritance.Gate503KernelIndexAccepted {
		t.Fatalf("bad inheritance: %+v", a.Inheritance)
	}
	if !a.Quotient.Executed || a.Quotient.PhotonKernelDimension != 1 || a.Quotient.BrokenOrbitRank != 3 || a.Quotient.RadialQuotientDimension != 1 || !nearly(a.Quotient.Diag114NeutralChargedRatio, 4, 1e-12) || a.Quotient.Diag114NativeMassRatio || a.Quotient.KappaNative || a.Quotient.WeakAngleDerived || a.Quotient.GaugeCouplingsDerived || a.Quotient.VEVDerived || a.Quotient.PhysicalWZMassMatrixDerived {
		t.Fatalf("bad quotient: %+v", a.Quotient)
	}
	if !a.Residuals.Executed || !a.Residuals.BridgeOnly || !a.Residuals.SyntheticOnly || a.Residuals.ObservedValuesImported || !a.Residuals.FileResidualsComputed || !a.Residuals.FileResidualsAllZero || !nearly(a.Residuals.FileNeutralChargedRatio, 25.0/9.0, 1e-12) || !nearly(a.Residuals.QuotientNeutralChargedRatio, 4, 1e-12) || !nearly(a.Residuals.Diag114ToFileRatioResidual, 11.0/9.0, 1e-12) || a.Residuals.Diag114RatioMatchedByFile || !a.Residuals.PhotonZeroAlignment || !a.Residuals.RhoIdentityConfirmed || a.Residuals.RhoIdentityNativeMassPrediction {
		t.Fatalf("bad residual geometry: %+v", a.Residuals)
	}
	if !a.Classification.PhotonZeroIsStructuralAlignment || !a.Classification.RhoIdentityIsBridgeFormula || !a.Classification.FileResidualsAreAdapterResiduals || !a.Classification.Diag114MismatchIsExpected || a.Classification.Diag114UsedAsMassRatio || a.Classification.WeakAngleNativePrediction || a.Classification.GaugeCouplingNativePrediction || a.Classification.VEVNativePrediction || a.Classification.WZMassNativePrediction || a.Classification.KappaNativePromotion {
		t.Fatalf("bad classification: %+v", a.Classification)
	}
	if a.Firewall.ObservedNumbersImported || a.Firewall.FileAdapterOutputsNative || a.Firewall.FileResidualsNative || a.Firewall.Diag114RatioNativeMassRatio || a.Firewall.WeakAngleNativeWritten || a.Firewall.GaugeCouplingsNativeWritten || a.Firewall.VEVNativeWritten || a.Firewall.WZMassNativeWritten || a.Firewall.KappaNativeWritten || a.Firewall.NativeRegistryWritten || a.Firewall.PhysicalElectroweakPredictionMade {
		t.Fatalf("firewall violation: %+v", a.Firewall)
	}
	if a.Next.Gate != 509 {
		t.Fatalf("expected Gate509 redirect: %+v", a.Next)
	}
	md := Markdown(a)
	for _, want := range []string{"# Gate 508 Registry Audit", StatusDiag114ResidualComputedBridgeOnly, StatusFailedDiag114NotMassRatio, StatusFirewallNativeWriteBlocked, "25/9", "11/9", "Gate 509"} {
		if !strings.Contains(md, want) {
			t.Fatalf("markdown missing %q", want)
		}
	}
}

func TestGate508TheoremPasses(t *testing.T) {
	res := Generation2ElectroweakComparatorResidualGeometryAirlockTheorem().Run()
	if !res.Passed() {
		t.Fatalf("theorem failed:\n%s", res.Details())
	}
}
