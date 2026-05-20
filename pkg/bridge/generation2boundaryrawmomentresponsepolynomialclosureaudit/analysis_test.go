package generation2boundaryrawmomentresponsepolynomialclosureaudit

import (
	"math"
	"strings"
	"testing"
)

func TestGate733CubicPolynomialClosure(t *testing.T) {
	a, err := Build()
	if err != nil {
		t.Fatal(err)
	}
	if !a.Gate732.Inherited || !a.Gate732.RawMomentCoordinateActive || !a.Gate732.RawM3BestCompression || !a.Gate732.NoNativeRawMomentTheorem {
		t.Fatalf("bad Gate732 inheritance: %+v", a.Gate732)
	}
	if math.Abs(a.Polynomial.LeadingTerm-a.Gate732.M1Wall) > 1e-18 || math.Abs(a.Polynomial.QuadraticTerm-a.Gate732.KappaE*a.Gate732.M2Wall) > 1e-18 || math.Abs(a.Polynomial.CubicTerm+2*a.Gate732.P_K7*a.Gate732.M3Wall) > 1e-18 {
		t.Fatalf("bad polynomial: %+v", a.Polynomial)
	}
	if math.Abs(a.Closure.Residual-(a.Gate732.DBase-a.Polynomial.Value)) > 1e-18 || math.Abs(a.Closure.Residual-a.Gate732.RawCubicResidual) > 1e-18 || !a.Closure.StrongCompression || a.Closure.CompressionFactor < 1000 {
		t.Fatalf("bad closure: %+v", a.Closure)
	}
}

func TestGate733FourthOrderStopAndRuntimeFirewalls(t *testing.T) {
	a, err := Build()
	if err != nil {
		t.Fatal(err)
	}
	expectedM4 := a.Gate732.P_K7 * math.Pow(a.Gate732.SSplit, 4)
	if math.Abs(a.FourthOrder.M4Wall-expectedM4) > 1e-18 || math.Abs(a.FourthOrder.RequiredCoeff-a.Closure.Residual/expectedM4) > 1e-18 || a.FourthOrder.TypedSourceFound || a.FourthOrder.PromoteFourthOrder {
		t.Fatalf("bad fourth order audit: %+v", a.FourthOrder)
	}
	if a.Stop.ProjectorPowersSupplyNewDirections || !a.Stop.HigherMomentsOnlyScalarPowers || a.Stop.TypedFourthOrderSourceFound || !a.Stop.StoppingAtCubicMoreLawful {
		t.Fatalf("bad stop condition: %+v", a.Stop)
	}
	if !a.Generating.CandidateTruncationSupported || a.Generating.NativeGeneratingFunction {
		t.Fatalf("bad generating-function audit: %+v", a.Generating)
	}
	if math.Abs(a.Runtime.RuntimeResidual-a.Runtime.LambdaProxy*a.Runtime.L*a.Closure.Residual) > 1e-18 || !a.Runtime.NearEliminated {
		t.Fatalf("bad runtime propagation: %+v", a.Runtime)
	}
	if !a.NonCircular.KappaEPartiallyDependent || a.NonCircular.DoubleK7CoefficientNative || a.NonCircular.BoundaryMomentExpansionNative {
		t.Fatalf("noncircularity firewall failed: %+v", a.NonCircular)
	}
	if a.Firewall.ScalarRuntimeTheoremNative || a.Firewall.HiggsMassTheoremNative || a.Firewall.YukawaTheoremNative {
		t.Fatalf("physical firewall failed: %+v", a.Firewall)
	}

	res := Generation2BoundaryRawMomentResponsePolynomialClosureAuditTheorem().Verify()
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
