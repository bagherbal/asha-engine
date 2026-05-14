package contactvacuumhopfaction

import (
	"math"
	"testing"
)

func TestGate284ContactVacuumHopfActionAudit(t *testing.T) {
	a, err := Build()
	if err != nil {
		t.Fatalf("Build() error = %v", err)
	}
	if !a.Gate283.Gate283Inherited || !a.Gate283.FourOverPiIdentity || !a.Gate283.BGapResonanceReproduced {
		t.Fatalf("expected Gate 283 resonance inheritance: %s", FormatGate283(a.Gate283))
	}
	if math.Abs(a.Instanton.Coefficient-4/math.Pi) > 1e-12 || a.Instanton.FiniteInstantonDerived {
		t.Fatalf("expected exact 4/pi candidate but no finite instanton theorem: %s", FormatInstanton(a.Instanton))
	}
	if a.BoundaryMap.ContactVacuumHopfActionMapDerived || a.BoundaryMap.BGapAsInverseCouplingDerived || a.BoundaryMap.ActionDensityOnFiberDerived {
		t.Fatalf("boundary action map must remain missing: %s", FormatBoundaryMap(a.BoundaryMap))
	}
	if a.OrderParameter.HiddenSectorOrderParameterDefined || a.OrderParameter.NonzeroVEVDerived || a.OrderParameter.CouplesToHopfAction {
		t.Fatalf("hidden order parameter must not be invented: %s", FormatOrderParameter(a.OrderParameter))
	}
	if a.Residual.RelativeDeltaCoefficient <= 0 || a.Residual.RelativeDeltaCoefficient > 0.004 || a.Residual.ResidualExacted {
		t.Fatalf("expected small residual but no correction theorem: %s", FormatResidual(a.Residual))
	}
	if a.Seal.IntermediateBreakingSealGranted || !a.Seal.RequiresInstantonActionMap || !a.Seal.RequiresHiddenOrderParameter {
		t.Fatalf("seal ledger violated: %s", FormatSeal(a.Seal))
	}
	if a.Firewall.FiniteCorePolluted || !a.Firewall.DoesNotFitCoefficient || !a.Firewall.DoesNotGrantIntermediateSeal {
		t.Fatalf("firewall failure: %s", FormatFirewall(a.Firewall))
	}
	if a.Summary.IntermediateTheoremUpgraded || a.Summary.ContactVacuumMapDerived || a.Summary.HiddenOrderParameterDerived {
		t.Fatalf("summary should not upgrade theorem: %s", FormatSummary(a.Summary))
	}
}

func TestGate284TheoremPassesChecks(t *testing.T) {
	res := NativeContactVacuumHopfActionMapHiddenSectorOrderParameterAuditTheorem().Run()
	if !res.Passed() {
		t.Fatalf("theorem checks failed:\n%s", res.Details())
	}
	if res.Status == "EXACT_FINITE" || res.Status == "PHENOMENOLOGY" {
		t.Fatalf("Gate 284 should remain BridgeRequired, got %s", res.Status)
	}
}
