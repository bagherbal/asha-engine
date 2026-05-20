package generation2gaugescalarboundarypairingrobustnessaudit

import (
	"math"
	"testing"
)

func TestBuildGate612Audit(t *testing.T) {
	a, err := Build()
	if err != nil {
		t.Fatal(err)
	}
	if len(a.CandidateScales) != 4 || len(a.PairingRatios) != 4 {
		t.Fatalf("expected four candidate scales and pairings")
	}
	l12 := pairingFor(a.PairingRatios, "Lambda_12")
	if math.Abs(l12.GaugeRelativeResidual-0.0509933868964996) > 1e-12 {
		t.Fatalf("Lambda12 gauge residual drifted: %.15g", l12.GaugeRelativeResidual)
	}
	if math.Abs(l12.AbsLambda-0.049700942077683274) > 2e-6 {
		t.Fatalf("Lambda12 scalar residual drifted: %.15g", l12.AbsLambda)
	}
	if !a.UniquenessAudit.Lambda12UniqueBest {
		t.Fatalf("expected Lambda12 to be the sharpest audited v1 pairing: %+v", a.UniquenessAudit)
	}
	if a.NativeStatus.ProvidesNativeJointCorrectionTheorem || a.Firewalls.ClaimsGaugeUnification {
		t.Fatalf("native theorem/firewall failure")
	}
}

func TestTheoremChecks(t *testing.T) {
	res := Generation2GaugeScalarBoundaryPairingRobustnessAndScaleDependenceAuditTheorem().Verify()
	for _, c := range res.Checks {
		if !c.Passed {
			t.Fatalf("check %q failed: %s", c.Name, c.Detail)
		}
	}
}
