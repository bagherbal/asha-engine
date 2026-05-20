package generation2koidereactorckmresidualclosureuncertaintyaudit

import (
	"strings"
	"testing"
)

func TestGate591ResidualInsideUncertainty(t *testing.T) {
	a, err := Build()
	if err != nil {
		t.Fatal(err)
	}
	if !a.Uncertainty.CoversKappa {
		t.Fatalf("expected one-sigma band to cover kappa: %+v", a.Uncertainty)
	}
	if a.Uncertainty.SigmaFractionPlus > 0.03 || a.Uncertainty.SigmaFractionMinus > 0.03 {
		t.Fatalf("residual should be tiny compared with uncertainty band: %+v", a.Uncertainty)
	}
	if a.Final.DeltaStatisticallyMeaningful {
		t.Fatalf("residual should not be statistically meaningful at current precision")
	}
}

func TestGate591DefectScaleAndCorrectionAudit(t *testing.T) {
	a, err := Build()
	if err != nil {
		t.Fatal(err)
	}
	if !a.Defects.DeltaSmallerThanR || !a.Defects.DeltaSmallerThanQ {
		t.Fatalf("expected delta590 smaller than Koide defects: %+v", a.Defects)
	}
	if a.Corrections.AnyCertified {
		t.Fatalf("no R/Q correction should be certified: %+v", a.Corrections)
	}
	if a.Corrections.BestCandidate.AbsResidual >= a.Residual.AbsDelta590 {
		t.Fatalf("best trial should reduce residual numerically without certification")
	}
}

func TestGate591Firewalls(t *testing.T) {
	a, err := Build()
	if err != nil {
		t.Fatal(err)
	}
	if a.Lawfulness.CrossSectorOrientationIntertwinerPresent || a.Lawfulness.NativeRootTraceOperatorPresent || !a.Final.KappaRemainsEnvironmental {
		t.Fatalf("firewall failure: law=%+v final=%+v", a.Lawfulness, a.Final)
	}
	th := Generation2KoideReactorCKMResidualClosureAndUncertaintyAuditTheorem()
	res := th.Verify()
	if !res.Passed() {
		t.Fatalf("theorem checks failed: %+v", res)
	}
	joined := strings.Join(res.Notes, "\n")
	for _, want := range []string{StatusResidualInsideOneSigma, StatusNoRQCorrectionCertified, StatusNoCrossSectorIntertwiner, StatusGate591Boundary} {
		if !strings.Contains(joined, want) {
			t.Fatalf("missing status %s", want)
		}
	}
}
