package generation2electroweakmeetingdeficitclosuredualrootaudit

import (
	"math"
	"strings"
	"testing"
)

func TestGate664Build(t *testing.T) {
	a, err := Build()
	if err != nil {
		t.Fatal(err)
	}
	if !a.Inherited.ZeroCrossingInherited || !a.Inherited.NoStationaryClaim || !a.Inherited.NoNativeScale || !a.Inherited.NoNativeSevenOver72 || !a.Inherited.NoUncertainty || !a.Inherited.NoBoundaryStress {
		t.Fatalf("bad inheritance: %+v", a.Inherited)
	}
	if a.Seed.Mu0GeV <= 0 || a.Seed.Lambda12GeV < 9e13 || len(a.Seed.InitialVector) != 13 {
		t.Fatalf("bad seed: %+v", a.Seed)
	}
	if math.Abs(a.Meeting.T12Analytic-a.Seed.T12) > 1e-12 || math.Abs(a.Meeting.F12AtRoot) > 1e-12 || math.Abs(a.Meeting.U12AtRoot) > 1e-10 {
		t.Fatalf("bad meeting root: %+v", a.Meeting)
	}
	if !a.ClosureRoot.ClosureIsTransverse || math.Abs(a.ClosureRoot.E72AtClosureZero) > 1e-12 {
		t.Fatalf("bad closure root: %+v", a.ClosureRoot)
	}
	if !a.DualRoot.AlignedInV1 || math.Abs(a.DualRoot.DeltaLogMuEOverMu12) > 1e-5 || math.Abs(a.DualRoot.MuEOverMu12-0.9999991071689) > 1e-9 {
		t.Fatalf("bad dual root offset: %+v", a.DualRoot)
	}
	if !a.Transversality.F12Transverse || !a.Transversality.U12Transverse || !a.Transversality.E72Transverse || a.Transversality.SlopeTied || a.Transversality.DE72DtAtLambda12 < 9e-4 {
		t.Fatalf("bad transversality: %+v", a.Transversality)
	}
	if a.Proportionality.Samples != 5 || a.Proportionality.RelativeResidualF12 > 0.2 || a.Proportionality.RelativeResidualU12 > 0.2 {
		t.Fatalf("bad proportionality: %+v", a.Proportionality)
	}
	if len(a.Conventions.Rows) != 5 || a.Conventions.DirectCouplingConventionsPass < 4 || !a.Conventions.ConventionStable {
		t.Fatalf("bad convention audit: %+v", a.Conventions)
	}
	if len(a.WeightRoot.Rows) != 5 || math.Abs(a.WeightRoot.WBestMinus7Over72AtLambda12) > 1e-6 || !a.WeightRoot.CrossesSevenOver72NearLambda || a.WeightRoot.WeightIndependentlySelected {
		t.Fatalf("bad weight root: %+v", a.WeightRoot)
	}
	if len(a.Source.Outcomes) != 5 || a.Discipline.ClaimsNativeDualRootTheorem || a.Discipline.ClaimsNativeSevenOver72Theorem || a.Discipline.ClaimsFullUncertaintyPropagation || a.Discipline.ClaimsBoundaryStressDerivation || a.Discipline.ClaimsNativeTransportTheorem || a.Discipline.ClaimsHiggsPrediction || a.Discipline.ClaimsScalarStability || a.Discipline.ClaimsFlavorDerivation || a.Discipline.ClaimsGaugeUnification || a.Discipline.ClaimsCKMPMNSDerivation || a.Discipline.Verdict != StatusGate664Boundary {
		t.Fatalf("firewall breach: %+v source=%+v", a.Discipline, a.Source)
	}
}

func TestTheorem(t *testing.T) {
	res := Generation2ElectroweakMeetingDeficitClosureDualRootAlignmentAuditTheorem().Verify()
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
