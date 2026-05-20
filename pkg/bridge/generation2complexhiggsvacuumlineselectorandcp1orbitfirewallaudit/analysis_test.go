package generation2complexhiggsvacuumlineselectorandcp1orbitfirewallaudit

import (
	"strings"
	"testing"
)

func TestGate762InheritancePrimaryTargetAndCP1Orbit(t *testing.T) {
	a, err := BuildDefault()
	if err != nil {
		t.Fatal(err)
	}
	if !a.Gate761.Inherited || !a.Gate761.PRadGaugeFixedInsideComplexLine || !a.Gate761.ComplexLineNotNativelySelected || !a.Gate761.RadialGaugeNotNativelySelected || !a.Gate761.ActiveHistoryUsesRealRadialHalf {
		t.Fatalf("bad Gate761 inheritance: %+v", a.Gate761)
	}
	if a.Primary.NewPrimaryTarget != "Pi_vac_C" || !a.Primary.ComplexLineBeforePRad || !a.Primary.RadialGaugeAfterLine || !a.Primary.LineLivesInK7PlusJ || !a.Primary.LineIsCP1Point {
		t.Fatalf("bad primary target: %+v", a.Primary)
	}
	if !a.Orbit.Recorded || a.Orbit.Socket != "K7+_J(n) ~= C^2" || !strings.Contains(a.Orbit.ComplexLines, "CP1") || a.Orbit.CP1RealDimension != cp1RealDim || a.Orbit.UnitRepresentatives != "S3" || a.Orbit.S3RealDimension != s3RealDim || !strings.Contains(a.Orbit.HopfFibration, "S1 -> S3 -> CP1") || !a.Orbit.BasePointSelectsLine || !a.Orbit.FiberGaugeSelectsRadialRep {
		t.Fatalf("bad CP1 orbit geometry: %+v", a.Orbit)
	}
}

func TestGate762SelectorQuestionAndSourceAudit(t *testing.T) {
	a, err := BuildDefault()
	if err != nil {
		t.Fatal(err)
	}
	if !a.Question.Formulated || !strings.Contains(a.Question.Question, "CP1") || !strings.Contains(a.Question.RequiredObject, "Pi_vac_C") || a.Question.RequiredRankR != complexLineRealRank || a.Question.RequiredRankC != complexLineRank || !a.Question.RequiresJH || !a.Question.RequiresCP1BasePoint || !a.Question.DoesNotRequireRadialGauge {
		t.Fatalf("bad selector question: %+v", a.Question)
	}
	if !a.SourceAudit.Completed || len(a.SourceAudit.Candidates) != 7 || a.SourceAudit.NativeComplexVacuumLineSelector || a.SourceAudit.NativeCP1BasePointSelector || a.SourceAudit.NativeRadialGaugeSelector || !a.SourceAudit.NSelectsJHOnly || !a.SourceAudit.RhoPlusNoBias || !a.SourceAudit.PRadAssumesLineAndGauge || !a.SourceAudit.AllCurrentCandidatesFailCP1 {
		t.Fatalf("bad source audit: %+v", a.SourceAudit)
	}
	for _, c := range a.SourceAudit.Candidates {
		if c.SelectsCP1Point {
			t.Fatalf("candidate unexpectedly selects CP1: %+v", c)
		}
	}
}

func TestGate762ConstructibilityGaugeHierarchyAndHistoryLoop(t *testing.T) {
	a, err := BuildDefault()
	if err != nil {
		t.Fatal(err)
	}
	if !a.Constructible.CanConstructPiFromSuppliedPRad || !strings.Contains(a.Constructible.ConstructionFormula, "J_H") || !a.Constructible.ConstructionDependsOnSeal || a.Constructible.ConstructionIsSelectionTheorem || !a.Constructible.PRadMayNotBeUsedAsNativeCause || a.Constructible.NativeLineSelectorCertified {
		t.Fatalf("bad constructibility audit: %+v", a.Constructible)
	}
	if a.GaugeHierarchy.ComplexLineSeal != "ComplexVacuumLineSeal" || a.GaugeHierarchy.RadialGaugeFixingSeal != "RadialGaugeFixingSeal" || !a.GaugeHierarchy.LineSelectionPrecedesGauge || !a.GaugeHierarchy.GaugeWithoutLineIllTyped || !a.GaugeHierarchy.PRadRequiresBothChoices || !a.GaugeHierarchy.SecondaryMarked {
		t.Fatalf("bad gauge hierarchy: %+v", a.GaugeHierarchy)
	}
	if !a.HistoryLoop.Refined || !a.HistoryLoop.DependsOnComplexLine || !a.HistoryLoop.DependsOnRadialGauge || a.HistoryLoop.RealRadialWeight != 0.25 || a.HistoryLoop.ComplexLineWeight != 0.50 || !a.HistoryLoop.FullComplexLineRejected || len(a.HistoryLoop.UnsolvedObjects) != 2 {
		t.Fatalf("bad HistoryLoop dependency: %+v", a.HistoryLoop)
	}
}

func TestGate762LayersFirewallsAndTheoremStatuses(t *testing.T) {
	a, err := BuildDefault()
	if err != nil {
		t.Fatal(err)
	}
	if !a.Layers.Audited || !a.Layers.ObjectsNotSameOperator || !strings.Contains(a.Layers.NLayer, "J_H") || !strings.Contains(a.Layers.CP1Layer, "Pi_vac_C") || !strings.Contains(a.Layers.RadialGaugeLayer, "P_rad") || !strings.Contains(a.Layers.HistoryLayer, "L_Hopf") || !strings.Contains(a.Layers.YukawaLayer, "N_eff") {
		t.Fatalf("bad layer separation: %+v", a.Layers)
	}
	if !a.Firewalls.Audited || a.Firewalls.PiVacCNativeVacuumTheorem || a.Firewalls.CP1PointNativeEWSBTheorem || a.Firewalls.RadialGaugeFixingNativeTheorem || a.Firewalls.LHopfNativeHistoryLoopTheorem || a.Firewalls.ComplexLineHiggsMassTheorem || a.Firewalls.ComplexLineYukawaTheorem || a.Firewalls.ScalarRuntimeIndependentTheorem || a.Firewalls.HiggsMassOrPoleMassTheorem || a.Firewalls.YukawaOperatorOrEigenvalueTheorem {
		t.Fatalf("bad firewalls: %+v", a.Firewalls)
	}
	res := Generation2ComplexHiggsVacuumLineSelectorAndCP1OrbitFirewallAuditTheorem().Verify()
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
