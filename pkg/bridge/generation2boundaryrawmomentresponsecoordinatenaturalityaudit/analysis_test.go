package generation2boundaryrawmomentresponsecoordinatenaturalityaudit

import (
	"math"
	"strings"
	"testing"
)

func TestGate732RawMomentCoordinate(t *testing.T) {
	a, err := Build()
	if err != nil {
		t.Fatal(err)
	}
	if !a.Gate731.Inherited || !a.Gate731.MomentPolynomialAvailable || !a.Gate731.NoNativeMomentExpansion {
		t.Fatalf("bad Gate731 inheritance: %+v", a.Gate731)
	}
	if !a.RawMoment.UsesRawMoments || math.Abs(a.RawMoment.M1Wall-a.Gate731.P_K7*a.Gate731.SSplit) > 1e-18 || math.Abs(a.RawMoment.M2Wall-a.Gate731.P_K7*a.Gate731.SSplit*a.Gate731.SSplit) > 1e-18 || math.Abs(a.RawMoment.M3Wall-a.Gate731.P_K7*math.Pow(a.Gate731.SSplit, 3)) > 1e-18 {
		t.Fatalf("bad raw moment response: %+v", a.RawMoment)
	}
	if !strings.Contains(a.RawMoment.FactoredFunction, "p_K7 S_split") {
		t.Fatalf("missing factored response function: %+v", a.RawMoment)
	}
	if !a.Degeneracy.AllPowersSupportedOnK7 || a.Degeneracy.IndependentOperatorDirections || !a.Degeneracy.ScalarResponseFunctionOnly {
		t.Fatalf("bad projector degeneracy: %+v", a.Degeneracy)
	}
}

func TestGate732VarianceCentralAndFirewalls(t *testing.T) {
	a, err := Build()
	if err != nil {
		t.Fatal(err)
	}
	expectedVar := a.Gate731.P_K7 * (1 - a.Gate731.P_K7) * a.Gate731.SSplit * a.Gate731.SSplit
	if math.Abs(a.Variance.VarianceWall-expectedVar) > 1e-18 || math.Abs(a.Variance.CoefficientInVariance-a.Gate731.EWall/expectedVar) > 1e-18 || !a.Variance.RawM2CloserToKappaE || !a.Variance.TypedButInactive {
		t.Fatalf("bad variance audit: %+v", a.Variance)
	}
	expectedMu3 := a.Gate731.P_K7 * (1 - a.Gate731.P_K7) * (1 - 2*a.Gate731.P_K7) * math.Pow(a.Gate731.SSplit, 3)
	if math.Abs(a.CentralM3.Mu3Wall-expectedMu3) > 1e-18 || !a.CentralM3.RawCompressesBetter {
		t.Fatalf("bad central m3 audit: %+v", a.CentralM3)
	}
	if !a.Comparison.RawSelectedByCurrentCompression || a.Comparison.CentralResidualAbs <= a.Comparison.RawResidualAbs {
		t.Fatalf("bad raw vs central comparison: %+v", a.Comparison)
	}
	if a.Firewall.RawMomentsNativelySelected || a.Firewall.BoundaryMomentTheoremNative || a.Firewall.ScalarRuntimeTheoremNative || a.Firewall.HiggsMassTheoremNative || a.Firewall.YukawaTheoremNative {
		t.Fatalf("firewall failed: %+v", a.Firewall)
	}

	res := Generation2BoundaryRawMomentResponseCoordinateNaturalityAuditTheorem().Verify()
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
