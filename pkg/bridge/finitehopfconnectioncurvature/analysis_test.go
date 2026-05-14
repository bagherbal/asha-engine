package finitehopfconnectioncurvature

import (
	"math"
	"testing"
)

func TestGate285FiniteHopfConnectionAudit(t *testing.T) {
	a, err := Build()
	if err != nil {
		t.Fatalf("Build() error = %v", err)
	}
	if !a.Gate284.Gate284Inherited || !a.Gate284.InstantonFunctionalFormalized || a.Gate284.ContactVacuumMapDerived {
		t.Fatalf("expected inherited Gate 284 formal action without map: %s", FormatGate284(a.Gate284))
	}
	if !a.Connection.HopfFibrationAvailable || !a.Connection.S3FiberAvailable || !a.Connection.LocalQuaternionicAlgebraHint {
		t.Fatalf("expected Hopf/Quaternionic target availability: %s", FormatConnection(a.Connection))
	}
	if a.Connection.NativeFiniteConnectionDerived || a.Connection.FiniteConnectionOneFormDerived {
		t.Fatalf("finite connection must not be invented: %s", FormatConnection(a.Connection))
	}
	if a.Curvature.CurvatureTwoFormDerived || a.Curvature.FiniteExteriorDDerived || a.Curvature.TracePairingDerived {
		t.Fatalf("curvature route should remain missing: %s", FormatCurvature(a.Curvature))
	}
	if a.ChernSimons.BoundaryWindingEvaluated || a.ChernSimons.IntegerWindingNumberDerived || a.ChernSimons.ChernSimonsThreeFormDerived {
		t.Fatalf("CS winding should not be evaluated: %s", FormatChernSimons(a.ChernSimons))
	}
	if math.Abs(a.Action.Coefficient-4/math.Pi) > 1e-12 || !a.Action.TopologicalRatioAvailable {
		t.Fatalf("expected inherited 4/pi ratio: %s", FormatAction(a.Action))
	}
	if a.Action.BGapAsInverseCouplingDerived || a.Action.ActionEvaluationDerived || a.Action.IntermediateScaleTheorem {
		t.Fatalf("action theorem should not be promoted: %s", FormatAction(a.Action))
	}
	if a.Coupling.InverseCouplingMapDerived || !a.Coupling.GaugeKineticNormalizationOpen {
		t.Fatalf("B_gap coupling map should stay open: %s", FormatCoupling(a.Coupling))
	}
	if a.Firewall.FiniteCorePolluted || !a.Firewall.DoesNotInventConnection || !a.Firewall.DoesNotPromoteBGapToCoupling || !a.Firewall.DoesNotGrantIntermediateSeal {
		t.Fatalf("firewall failure: %s", FormatFirewall(a.Firewall))
	}
	if a.Summary.IntermediateTheorem || a.Summary.ActionEvaluated || a.Summary.FiniteConnectionDerived {
		t.Fatalf("summary should preserve failed route: %s", FormatSummary(a.Summary))
	}
}

func TestGate285TheoremPassesChecks(t *testing.T) {
	res := FiniteHopfConnectionCurvatureChernSimonsBoundaryWindingAuditTheorem().Run()
	if !res.Passed() {
		t.Fatalf("theorem checks failed:\n%s", res.Details())
	}
	if res.Status != "BRIDGE_REQUIRED" {
		t.Fatalf("Gate 285 should remain BridgeRequired, got %s", res.Status)
	}
}
