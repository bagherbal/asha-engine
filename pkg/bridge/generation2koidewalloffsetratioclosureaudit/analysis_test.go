package generation2koidewalloffsetratioclosureaudit

import (
	"math"
	"testing"
)

func TestGate584KoideWallOffsetRatioClosureAudit(t *testing.T) {
	a, err := BuildDefault()
	if err != nil {
		t.Fatal(err)
	}
	if math.Abs(a.MZ.FromElectronMuon.SolvedEpsilonDeg-2.26761458653505) > 1e-10 {
		t.Fatalf("unexpected epsilon solved from e/mu: %s", FormatPrediction(a.MZ.FromElectronMuon))
	}
	if math.Abs(a.MZ.FromElectronMuon.PredictedRootRatio-0.243843978487765) > 1e-12 {
		t.Fatalf("unexpected predicted mu/tau ratio: %s", FormatPrediction(a.MZ.FromElectronMuon))
	}
	if math.Abs(a.MZ.FromElectronMuon.RootResidual+7.48094669156e-06) > 1e-12 {
		t.Fatalf("unexpected mu/tau residual: %s", FormatPrediction(a.MZ.FromElectronMuon))
	}
	if math.Abs(a.MZ.FromMuonTau.PredictedRootRatio-0.069519379373936) > 1e-12 {
		t.Fatalf("unexpected predicted e/mu ratio: %s", FormatPrediction(a.MZ.FromMuonTau))
	}
	if !a.MZ.ClosureCertified || !a.Lambda12.ClosureCertified {
		t.Fatalf("charged-lepton one-parameter closure should certify: %s / %s", FormatClosure(a.MZ), FormatClosure(a.Lambda12))
	}
	if !(a.Transport.ClosureStable && a.Transport.ResidualImprovesAtBoundary) {
		t.Fatalf("closure should be stable and improve at boundary: %s", FormatTransport(a.Transport))
	}
	if a.Quarks.OneParameterClosure {
		t.Fatalf("quarks should not certify one-parameter Koide wall closure: %s", FormatQuarks(a.Quarks))
	}
	if a.Firewalls.DerivesEpsilon || a.Firewalls.PromotesObservedAsNative || !a.Firewalls.PreservesGate352 {
		t.Fatalf("firewall broken: %s", FormatFirewalls(a.Firewalls))
	}
}

func TestGate584Theorem(t *testing.T) {
	res := Generation2KoideWallOffsetRatioClosureAuditTheorem().Run()
	if res.Status == "FAILED_ROUTE" {
		t.Fatalf("unexpected failed route: %+v", res)
	}
	for _, c := range res.Checks {
		if !c.Passed {
			t.Fatalf("check failed: %+v", c)
		}
	}
}
