package generation2electroweakrootclosurecoordinatenaturalityaudit

import (
	"math"
	"strings"
	"testing"
)

func TestGate665Build(t *testing.T) {
	a, err := Build()
	if err != nil {
		t.Fatal(err)
	}
	if !a.Inherited.DualRootInherited || !a.Inherited.TransverseCrossing || !a.Inherited.NoNativeDualRoot || !a.Inherited.NoNativeSevenOver72 || !a.Inherited.NoFullUncertainty || !a.Inherited.NoBoundaryStress {
		t.Fatalf("bad inheritance: %+v", a.Inherited)
	}
	if math.Abs(a.Inherited.ClosureRootRatio-0.9999991071689) > 1e-9 {
		t.Fatalf("bad inherited root ratio: %+v", a.Inherited)
	}
	if a.Seed.Mu0GeV <= 0 || a.Seed.Lambda12GeV < 9e13 || len(a.Seed.InitialVector) != 13 {
		t.Fatalf("bad seed: %+v", a.Seed)
	}
	if !a.CommonRoot.ConditionalRootPass || math.Abs(a.CommonRoot.F12AtRoot) > 1e-12 || math.Abs(a.CommonRoot.E72AmplitudeAtRoot) > 1e-8 || math.Abs(a.CommonRoot.WBestMinus7Over72) > 1e-6 {
		t.Fatalf("bad common root: %+v", a.CommonRoot)
	}
	if a.Factorization.Samples != 5 || a.Factorization.RelativeResidualF12 > 0.01 || a.Factorization.RelativeResidualU12 > 0.01 {
		t.Fatalf("bad factorization: %+v", a.Factorization)
	}
	if len(a.Coordinates.Rows) != 5 || a.Coordinates.AmplitudeRowsNearWeight != 1 || a.Coordinates.InverseRowsNearWeight != 0 || !a.Coordinates.AmplitudeNatural || a.Coordinates.RGNativeInverseNatural || a.Coordinates.CoordinateRobust {
		t.Fatalf("bad coordinate family: %+v", a.Coordinates)
	}
	var amp, inv CoordinateRow
	for _, r := range a.Coordinates.Rows {
		if r.Name == "amplitude ratio" {
			amp = r
		}
		if r.Name == "inverse-coupling ratio" {
			inv = r
		}
	}
	if math.Abs(amp.WBestMinus7Over72) > 1e-6 || !amp.RootFoundNearLambda12 || !amp.NearSevenOver72 {
		t.Fatalf("amplitude coordinate should pass: %+v", amp)
	}
	if math.Abs(inv.WBestMinus7Over72) < 0.01 || inv.RootFoundNearLambda12 || inv.NearSevenOver72 {
		t.Fatalf("inverse coordinate should fail: %+v", inv)
	}
	if !strings.Contains(a.CoordinateSeal.Classification, "amplitude-coordinate") || len(a.CoordinateSeal.Outcomes) != 4 {
		t.Fatalf("bad coordinate seal: %+v", a.CoordinateSeal)
	}
	if len(a.Source.Interpretations) != 4 || a.Discipline.ClaimsNativeDualRootTheorem || a.Discipline.ClaimsNativeSevenOver72Theorem || a.Discipline.ClaimsFullUncertaintyPropagation || a.Discipline.ClaimsBoundaryStressDerivation || a.Discipline.ClaimsNativeTransportTheorem || a.Discipline.ClaimsHiggsPrediction || a.Discipline.ClaimsScalarStability || a.Discipline.ClaimsFlavorDerivation || a.Discipline.ClaimsGaugeUnification || a.Discipline.ClaimsCKMPMNSDerivation || a.Discipline.Verdict != StatusGate665Boundary {
		t.Fatalf("firewall breach: discipline=%+v source=%+v", a.Discipline, a.Source)
	}
}

func TestTheorem(t *testing.T) {
	res := Generation2ElectroweakRootClosureCoordinateNaturalityAuditTheorem().Verify()
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
