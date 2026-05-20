package generation2boundaryexteriordegreeresponsemapandcubicstopsourceaudit

import (
	"strings"
	"testing"
)

func TestGate784BoundaryPairAndExteriorCandidate(t *testing.T) {
	a, err := BuildDefault()
	if err != nil {
		t.Fatal(err)
	}
	if !a.Gate783.Inherited || !strings.Contains(a.Gate783.BoundarySubBottleneck, "F_wall_3_red") {
		t.Fatalf("bad Gate783 inheritance: %+v", a.Gate783)
	}
	if !closeRel(a.Ledger.M1, 0.0001256543573849177, 1e-14) || !closeRel(a.Ledger.M2, 1.624013231638281e-07, 1e-14) || !closeRel(a.Ledger.M3, 2.0989474869200057e-10, 1e-14) || !a.Ledger.Matches {
		t.Fatalf("bad ledger: %+v", a.Ledger)
	}
	if !a.BoundaryPair.Typed || a.BoundaryPair.Dim != 2 || !a.BoundaryPair.TwoDimensionalBridge || a.BoundaryPair.IsSpacetimeCarrier || a.BoundaryPair.IsFlavorCarrier || a.BoundaryPair.IsK7Carrier {
		t.Fatalf("bad boundary pair audit: %+v", a.BoundaryPair)
	}
	if !a.Exterior.Audited || !a.Exterior.Lambda3Zero || !a.Exterior.CubicStopCandidate || a.Exterior.DimensionAloneProvesStop || a.Exterior.NativeThetaExtMapCertified || !strings.Contains(a.Exterior.DegreeAssignments["M4"], "Lambda^3") {
		t.Fatalf("bad exterior candidate: %+v", a.Exterior)
	}
}

func TestGate784ThetaExtCubicM4AndGeneratingCandidate(t *testing.T) {
	a, err := BuildDefault()
	if err != nil {
		t.Fatal(err)
	}
	if !a.ThetaExt.Identified || a.ThetaExt.Certified || !a.ThetaExt.DimensionShortcutRejected || !strings.Contains(a.ThetaExt.Codomain, "Lambda^(n-1)") {
		t.Fatalf("bad Theta_ext audit: %+v", a.ThetaExt)
	}
	if !a.Cubic.Audited || !closeRel(a.Cubic.TwoP, 7.0/36.0, 1e-15) || !a.Cubic.CompatibleWithDegree2 || a.Cubic.SignSourceCertified || a.Cubic.StressPullTheorem {
		t.Fatalf("bad cubic audit: %+v", a.Cubic)
	}
	if !a.M4.Audited || !a.M4.BlockedIfThetaExt || a.M4.NativeThetaExtCertified || a.M4.TypedCoefficientSource || a.M4.NativeCubicStopTheorem || !a.M4.UntypedM4FitRejected || a.M4.M4 <= 0 {
		t.Fatalf("bad M4 audit: %+v", a.M4)
	}
	if !a.Generating.Recorded || !strings.Contains(a.Generating.Truncation, "x+kappa_e_red") || a.Generating.NativeFunctionTheorem || a.Generating.HigherTermsProvedVanish {
		t.Fatalf("bad generating function audit: %+v", a.Generating)
	}
}

func TestGate784ResponseTableRelationPredictionAndFirewalls(t *testing.T) {
	a, err := BuildDefault()
	if err != nil {
		t.Fatal(err)
	}
	if !a.Table.Recorded || !a.Table.KappaERedNaturallyDegree1 || a.Table.FlavorBoundaryModulationNative || !strings.Contains(a.Table.Degree2, "-2p M3") {
		t.Fatalf("bad degree table: %+v", a.Table)
	}
	if !a.Projector.Preserved || a.Projector.ProjectorStopsExpansion || !a.Projector.NeedsExteriorMap || !strings.Contains(a.Projector.PowerLaw, "s^n P_7") {
		t.Fatalf("bad projector firewall: %+v", a.Projector)
	}
	if !a.Relation.Recorded || !a.Relation.WouldUpgradeFWall || a.Relation.KappaLambdaNative || a.Relation.CHistoryIndependent || !strings.Contains(a.Relation.CHiggsStatus, "not Level C") {
		t.Fatalf("bad relation audit: %+v", a.Relation)
	}
	if !a.Prediction.Recorded || !strings.Contains(a.Prediction.FWall3Level, "Level B+") || a.Prediction.FWallLevelC || !strings.Contains(a.Prediction.CHiggsLevel, "not Level C") {
		t.Fatalf("bad prediction classification: %+v", a.Prediction)
	}
	if !a.Firewalls.Enforced || a.Firewalls.BoundaryExteriorDegreeNative || a.Firewalls.DimensionTwoProof || a.Firewalls.TwoPCoefficientNative || a.Firewalls.NegativeSignNative || a.Firewalls.FWallNative || a.Firewalls.KappaLambdaNative || a.Firewalls.CHistoryIndependent || a.Firewalls.TreeProxyPoleMass || a.Firewalls.YukawaNative {
		t.Fatalf("bad firewalls: %+v", a.Firewalls)
	}
	if !strings.Contains(a.FinalStatement, "does not prove the cubic stop") || !strings.Contains(a.FinalStatement, "Theta_ext") || !strings.Contains(a.FinalStatement, "next bottleneck") {
		t.Fatalf("bad final statement: %s", a.FinalStatement)
	}
}

func TestGate784TheoremStatuses(t *testing.T) {
	res := Generation2BoundaryExteriorDegreeResponseMapAndCubicStopSourceAuditTheorem().Verify()
	if res.Status == "FAILED_ROUTE" {
		t.Fatalf("unexpected theorem failure: %+v", res)
	}
	for _, c := range res.Checks {
		if !c.Passed {
			t.Fatalf("check failed: %s %s", c.Name, c.Detail)
		}
	}
	joined := strings.Join(res.Notes, "\n")
	for _, want := range Statuses() {
		if !strings.Contains(joined, want) {
			t.Fatalf("missing status note %s", want)
		}
	}
}
