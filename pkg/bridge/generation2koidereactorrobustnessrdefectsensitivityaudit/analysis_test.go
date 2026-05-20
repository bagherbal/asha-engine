package generation2koidereactorrobustnessrdefectsensitivityaudit

import (
	"math"
	"testing"
)

func TestGate589KoideReactorRobustnessRDefectSensitivityAudit(t *testing.T) {
	a, err := BuildDefault()
	if err != nil {
		t.Fatal(err)
	}
	if math.Abs(a.Runtime.KappaObs-0.005503554191574556) > 2e-15 {
		t.Fatalf("unexpected observed kappa: %s", FormatRuntime(a.Runtime))
	}
	if math.Abs(a.Runtime.KappaR1-0.00531293763388241) > 2e-15 {
		t.Fatalf("unexpected exact-R1 kappa: %s", FormatRuntime(a.Runtime))
	}
	if math.Abs(a.Robustness.Observed.RelativeResidual-0.00616797931733138) > 2e-12 {
		t.Fatalf("unexpected observed residual: %s", FormatKappaComparison(a.Robustness.Observed))
	}
	if math.Abs(a.Robustness.ExactR1.RelativeResidual-0.0422670811502625) > 2e-12 {
		t.Fatalf("unexpected exact-R1 residual: %s", FormatKappaComparison(a.Robustness.ExactR1))
	}
	if !a.Robustness.Observed.WithinThetaOneSigma || a.Robustness.ExactR1.WithinThetaOneSigma {
		t.Fatalf("unexpected one-sigma coverage: %s", FormatRobustness(a.Robustness))
	}
	if math.Abs(a.Robustness.ExactR1.Theta13PredDeg-8.38243836864531) > 1e-11 {
		t.Fatalf("unexpected R1 theta prediction: %s", FormatKappaComparison(a.Robustness.ExactR1))
	}
	if math.Abs(a.RDefect.RequiredC-20.6455256996) > 1e-9 {
		t.Fatalf("unexpected required R-defect coefficient: %s", FormatRDefect(a.RDefect))
	}
	if a.RDefect.BestCandidate.Name != "8*pi" || a.RDefect.BestCandidateCertified {
		t.Fatalf("unexpected R-defect candidate: %s", FormatRDefect(a.RDefect))
	}
	if !a.Shift.ControlledByEpsilonShift || a.Shift.ControlledByRDefectTyped || a.Shift.ControlledByQResidualTyped {
		t.Fatalf("unexpected shift control: %s", FormatShift(a.Shift))
	}
	if a.Operator.NativeKoideReactorOperatorPresent || a.Operator.DerivesKappa || a.Firewalls.PromotesObservedAsNative || !a.Firewalls.PreservesGate352 {
		t.Fatalf("firewall/operator broken: %s / %s", FormatOperator(a.Operator), FormatFirewalls(a.Firewalls))
	}
}

func TestGate589Theorem(t *testing.T) {
	res := Generation2KoideReactorRobustnessRDefectSensitivityAuditTheorem().Run()
	if res.Status == "FAILED_ROUTE" {
		t.Fatalf("unexpected failed route: %+v", res)
	}
	for _, c := range res.Checks {
		if !c.Passed {
			t.Fatalf("check failed: %+v", c)
		}
	}
}
