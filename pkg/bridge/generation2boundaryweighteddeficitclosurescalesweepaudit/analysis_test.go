package generation2boundaryweighteddeficitclosurescalesweepaudit

import (
	"math"
	"strings"
	"testing"
)

func TestGate662Build(t *testing.T) {
	a, err := Build()
	if err != nil {
		t.Fatal(err)
	}
	if !a.Inherited.ClosureInherited || math.Abs(a.Inherited.E72-8.525834413464217e-10) > 5e-18 || !a.Inherited.Lambda12OnlyComputed || !a.Inherited.FormulaLiftCircular || !a.Inherited.NoNativeSevenOver72 || !a.Inherited.NoNativeTransport || !a.Inherited.NoIndependentEndpoint || !a.Inherited.FirewallPreserved {
		t.Fatalf("bad inheritance: %+v", a.Inherited)
	}
	if a.Seed.Mu0GeV <= 0 || a.Seed.Lambda12GeV < 9e13 || a.Seed.T13 <= a.Seed.T12 || a.Seed.T23 <= a.Seed.T13 || len(a.Seed.InitialVector) != 13 {
		t.Fatalf("bad seed: %+v", a.Seed)
	}
	if len(a.ScaleSweep.Rows) != 4 || a.ScaleSweep.BestEWMeanScale != "Lambda_12" || a.ScaleSweep.BestPairScale != "Lambda_12" || !a.ScaleSweep.Lambda12UniquelyMinimalEW || !a.ScaleSweep.Lambda12UniquelyMinimalPair || a.ScaleSweep.BestEWMeanResidual > 1e-9 || a.ScaleSweep.BestPairResidual > 1e-9 {
		t.Fatalf("bad scale sweep: %+v", a.ScaleSweep)
	}
	if math.Abs(a.ScaleSweep.Rows[0].E72EWMean-8.525834691019967e-10) > 5e-17 || math.Abs(a.ScaleSweep.Rows[1].E72EWMean-0.0024261778886282) > 5e-12 || math.Abs(a.ScaleSweep.Rows[2].E72EWMean-0.00794360239762586) > 5e-12 {
		t.Fatalf("unexpected scale sweep values: %+v", a.ScaleSweep.Rows)
	}
	if len(a.LocalSweep.Rows) != 9 || !a.LocalSweep.LocalGridSelectsLambda12 || math.Abs(a.LocalSweep.MinimumDeltaLog) > 1e-15 || a.LocalSweep.MinimumAbsResidual > 1e-9 || a.LocalSweep.Threshold1eMinus4Width < 0.1 || a.LocalSweep.FiniteDifferenceSlope < 0.0009 {
		t.Fatalf("bad local sweep: %+v", a.LocalSweep)
	}
	if math.Abs(a.Weight.WBestExact-0.09722288188941036) > 5e-15 || math.Abs(a.Weight.WBestExactMinus7Over72-6.596671881381466e-07) > 5e-15 || math.Abs(a.Weight.WBestOrientation-0.09937065106104444) > 5e-15 || !a.Weight.ExactWeightNear7Over72 || !a.Weight.OrientationWeightNear7Over72 {
		t.Fatalf("bad weight audit: %+v", a.Weight)
	}
	if a.Jacobian.DE_DKappaE != 1 || math.Abs(a.Jacobian.DE_DAbsLambda+65.0/72.0) > 1e-15 || math.Abs(a.Jacobian.DE_DR3Minus1+7.0/72.0) > 1e-15 || a.Jacobian.DKappa_DLambdaRuntime > -200 || a.Jacobian.DKappa_DLambdaProxy < 200 || a.Jacobian.DKappa_DL < 20 || len(a.Jacobian.Notes) != 3 {
		t.Fatalf("bad jacobian: %+v", a.Jacobian)
	}
	if math.Abs(a.Orientation.OrientationE72AtLambda12-2.7767257213331953e-06) > 5e-18 || a.Orientation.ClosureResidualAmplification < 3000 || a.Orientation.BestWeightShift < 0.002 {
		t.Fatalf("bad orientation scale audit: %+v", a.Orientation)
	}
	if a.Discipline.ClaimsNativeScaleSelection || a.Discipline.ClaimsNativeSevenOver72Theorem || a.Discipline.ClaimsFullUncertaintyPropagation || a.Discipline.ClaimsNativeTransportTheorem || a.Discipline.ClaimsBoundaryStressDerivation || a.Discipline.ClaimsHiggsPrediction || a.Discipline.ClaimsScalarStability || a.Discipline.ClaimsFlavorDerivation || a.Discipline.ClaimsGaugeUnification || a.Discipline.ClaimsCKMPMNSDerivation || a.Discipline.Verdict != StatusGate662Boundary {
		t.Fatalf("firewall breach: %+v", a.Discipline)
	}
}

func TestTheorem(t *testing.T) {
	res := Generation2BoundaryWeightedDeficitClosureScaleSweepAndSensitivityAuditTheorem().Verify()
	if res.Status == "FAILED_ROUTE" {
		t.Fatalf("unexpected construction failure: %+v", res)
	}
	for _, c := range res.Checks {
		if !c.Passed {
			t.Fatalf("check failed: %s %s", c.Name, c.Detail)
		}
	}
	joined := strings.Join(res.Notes, "\n")
	for _, want := range []string{StatusGate661ClosureInherited, StatusScaleSweepComputed, StatusLambda12SelectedInV1, StatusLocalPerturbationComputed, StatusLocalMinimumAtLambda12, StatusWeightSensitivityComputed, StatusSevenOver72WeightRobustInV1ExactLedger, StatusOrientationApproximationPerturbsWeight, StatusInputJacobianComputed, StatusScaleSpecificityNotNative, StatusNoNativeSevenOver72Theorem, StatusNoFullUncertaintyPropagation, StatusNoNativeScalarFlavorBoundaryTheorem, StatusNoBoundaryStressDerivation, StatusNoHiggsStabilityGaugeFlavorClaim, StatusGate662Boundary} {
		if !strings.Contains(joined, want) {
			t.Fatalf("missing note %s", want)
		}
	}
}
