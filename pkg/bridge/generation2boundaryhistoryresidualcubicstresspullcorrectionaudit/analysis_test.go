package generation2boundaryhistoryresidualcubicstresspullcorrectionaudit

import (
	"math"
	"strings"
	"testing"
)

func TestGate730CubicStressPullCorrection(t *testing.T) {
	a, err := Build()
	if err != nil {
		t.Fatal(err)
	}
	if !a.Gate729.Inherited || math.Abs(a.Gate729.P_K7-7.0/72.0) > 1e-18 || !a.Gate729.KappaEPartiallyDependent || !a.Gate729.NoNativeSecondOrderBoundaryTheorem {
		t.Fatalf("bad Gate729 inheritance: %+v", a.Gate729)
	}
	expectedM3 := a.Gate729.P_K7 * math.Pow(a.Gate729.SSplit, 3)
	if math.Abs(a.CubicMoment.M3Wall-expectedM3) > 1e-18 || math.Abs(a.CubicMoment.NegativeE2OverM3-(-a.Gate729.E2Residual/expectedM3)) > 1e-18 || !a.CubicMoment.SecondResidualCubicScale {
		t.Fatalf("bad cubic moment: %+v", a.CubicMoment)
	}
	if !a.Coefficients.SevenOver36Closest || a.Coefficients.ClosestName != "7/36" || !a.Coefficients.NoArbitrarySearch {
		t.Fatalf("bad cubic coefficient audit: %+v", a.Coefficients)
	}
	if math.Abs(a.CubicCorr.QuadraticTerm-a.Gate729.KappaE*a.Gate729.M2Wall) > 1e-18 || math.Abs(a.CubicCorr.CubicStressPullTerm-(7.0/36.0)*a.CubicMoment.M3Wall) > 1e-18 {
		t.Fatalf("bad correction terms: %+v", a.CubicCorr)
	}
	if math.Abs(a.CubicCorr.ResidualAfterCubicCorrection-(a.Gate729.EWall-a.CubicCorr.CombinedCorrection)) > 1e-18 || !a.CubicCorr.ImprovesSecondOrderResidual || !a.CubicCorr.NotExact || a.CubicCorr.RawCompressionFactor < 1000 {
		t.Fatalf("bad cubic correction: %+v", a.CubicCorr)
	}
}

func TestGate730RuntimeNoncircularityAndFirewalls(t *testing.T) {
	a, err := Build()
	if err != nil {
		t.Fatal(err)
	}
	if math.Abs(a.Runtime.CubicCorrectedRuntimeResidual-a.Gate729.LambdaProxy*a.Gate729.L*a.CubicCorr.ResidualAfterCubicCorrection) > 1e-18 || !a.Runtime.CompressedToNearFloatScale || !a.Runtime.ImprovesSecondOrderRuntime {
		t.Fatalf("bad runtime propagation: %+v", a.Runtime)
	}
	if a.SourceType.MomentExpansionTheoremNative || !strings.Contains(a.SourceType.ExpansionFormula, "Tr(rho_72 R_wall^3)") {
		t.Fatalf("bad source type: %+v", a.SourceType)
	}
	if !a.NonCircular.DBaseContainsKappaE || !a.NonCircular.KappaEUsedAsQuadraticCoeff || !a.NonCircular.CubicCoeffTypedButUnexplained || a.NonCircular.NativeExpansionTheorem {
		t.Fatalf("bad noncircularity: %+v", a.NonCircular)
	}
	if a.Firewall.ClaimsNativeBoundaryHistory || a.Firewall.ClaimsNativeMomentExpansion || a.Firewall.ClaimsNativeScalarRuntime || a.Firewall.ClaimsHiggsMassTheorem || a.Firewall.ClaimsYukawaTheorem || a.Firewall.ClaimsCKMPMNSTheorem {
		t.Fatalf("firewall failed: %+v", a.Firewall)
	}

	res := Generation2BoundaryHistoryResidualCubicStressPullCorrectionAuditTheorem().Verify()
	if res.Status == "FAILED_ROUTE" {
		t.Fatalf("unexpected construction failure: %+v", res)
	}
	for _, c := range res.Checks {
		if !c.Passed {
			t.Fatalf("check failed: %s %s", c.Name, c.Detail)
		}
	}
	joined := strings.Join(res.Notes, "\n")
	for _, want := range Statuses() {
		if !strings.Contains(joined, want) {
			t.Fatalf("missing note %s", want)
		}
	}
}
