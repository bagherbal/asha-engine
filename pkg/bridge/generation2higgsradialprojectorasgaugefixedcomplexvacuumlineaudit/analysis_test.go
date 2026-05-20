package generation2higgsradialprojectorasgaugefixedcomplexvacuumlineaudit

import (
	"math"
	"strings"
	"testing"
)

func TestGate761InheritanceComplexStructureAndLine(t *testing.T) {
	a, err := Build()
	if err != nil {
		t.Fatal(err)
	}
	if !a.Gate760.Inherited || !a.Gate760.PRadPriorityInherited || a.Gate760.HighestPrioritySeal != "P_rad" || !a.Gate760.LHopfDependsOnPRad || a.Gate760.NativePRadSelector || a.Gate760.NativeHistoryLoopTheorem || a.Gate760.NativeScalarRuntimeTheorem {
		t.Fatalf("bad Gate760 inheritance: %+v", a.Gate760)
	}
	if !a.Complex.JHSquaresMinusIdentity || !a.Complex.JHSkewOrthogonal || a.Complex.K7PlusRealDimension != k7PlusRealDim || a.Complex.K7PlusComplexDimension != k7PlusComplexDim || !a.Complex.K7PlusAsC2 || !a.Complex.Gate726SplitInherited {
		t.Fatalf("bad complex structure inheritance: %+v", a.Complex)
	}
	if !a.Directions.Defined || !a.Directions.RadialPhaseOrthogonal || a.Directions.RadialRank != realRadialRank || a.Directions.PhaseRank != realPhaseRank || !strings.Contains(a.Directions.PhaseVector, "J_H") {
		t.Fatalf("bad radial/phase directions: %+v", a.Directions)
	}
	if !a.VacuumLine.ConstructedFromPRadJH || a.VacuumLine.RealRank != complexLineRealRank || a.VacuumLine.ComplexRank != complexLineRank || !a.VacuumLine.JInvariant || !a.VacuumLine.CommutesWithJH || !a.VacuumLine.InsideK7PlusC2 || !a.VacuumLine.ContainsRadial || !a.VacuumLine.ContainsPhase {
		t.Fatalf("bad complex vacuum line: %+v", a.VacuumLine)
	}
}

func TestGate761WeightsGaugeFixingAndHistoryLoop(t *testing.T) {
	a, err := Build()
	if err != nil {
		t.Fatal(err)
	}
	if math.Abs(a.Weights.RadialWeight-0.25) > 1e-18 || math.Abs(a.Weights.PhaseWeight-0.25) > 1e-18 || math.Abs(a.Weights.ComplexLineWeight-0.5) > 1e-18 || math.Abs(a.Weights.TransverseWeight-0.5) > 1e-18 || math.Abs(a.Weights.TotalWeight-1.0) > 1e-18 {
		t.Fatalf("bad event weights: %+v", a.Weights)
	}
	if math.Abs(a.Weights.LHopfFromRadialEvent-1/(8*math.Pi)) > 1e-18 || math.Abs(a.Weights.LoopFromComplexLine-1/(4*math.Pi)) > 1e-18 || !a.Weights.ActiveHistoryUsesRadial || !a.Weights.ComplexLineTooLargeForL {
		t.Fatalf("bad loop weights: %+v", a.Weights)
	}
	if !a.GaugeFixing.PRadInsideComplexLine || a.GaugeFixing.PRadArbitraryPrimitiveLine || !a.GaugeFixing.ComplexLineContainsHopfPhase || !a.GaugeFixing.ScalarVacuumDirectionSealSplit || !strings.Contains(a.GaugeFixing.RefinedTyping, "GaugeFixedRadialRepresentativeSeal") {
		t.Fatalf("bad gauge-fixing typing: %+v", a.GaugeFixing)
	}
	if !a.HistoryLoop.ActiveUsesRealRadialAmplitude || !a.HistoryLoop.FullComplexLineRejected || math.Abs(a.HistoryLoop.ActiveLoopUnit-1/(8*math.Pi)) > 1e-18 || math.Abs(a.HistoryLoop.FullComplexLineLoopUnit-1/(4*math.Pi)) > 1e-18 {
		t.Fatalf("bad HistoryLoop implication: %+v", a.HistoryLoop)
	}
}

func TestGate761SelectorsOrbitSourceAuditAndFirewalls(t *testing.T) {
	a, err := Build()
	if err != nil {
		t.Fatal(err)
	}
	if !a.Selectors.NSelectsComplexStructure || a.Selectors.NSelectsVacuumLine || a.Selectors.ComplexLineSelectsRadialGauge || !a.Selectors.PRadSelectsGaugeRepresentative || !a.Selectors.ThreeChoicesDistinct {
		t.Fatalf("bad selector distinction: %+v", a.Selectors)
	}
	if !a.Orbit.CP1BasePoint || !a.Orbit.S1FiberGaugeRepresentative || !a.Orbit.RealRadialAmplitudeAxis || !a.Orbit.RefinesGate726S3Orbit || a.Orbit.ComplexVacuumLineOrbit != "CP1" || a.Orbit.UnitRepresentativeOrbit != "S3" {
		t.Fatalf("bad U2 Hopf orbit interpretation: %+v", a.Orbit)
	}
	if !a.SourceAudit.Completed || len(a.SourceAudit.Candidates) != 6 || a.SourceAudit.NativeComplexVacuumLineSelector || a.SourceAudit.NativeRadialGaugeFixingSelector || a.SourceAudit.RhoPlusSelectsLine || a.SourceAudit.NSelectsLine || a.SourceAudit.QSelectsLine || a.SourceAudit.BoundaryScalarsSelectLine || a.SourceAudit.FanoQuaternionicSelectsLine {
		t.Fatalf("bad source-candidate audit: %+v", a.SourceAudit)
	}
	if !a.Firewalls.Audited || a.Firewalls.PRadNativeVacuumTheorem || a.Firewalls.ComplexLineNativeEWSBTheorem || a.Firewalls.RadialGaugeFixingPhysicalEWSB || a.Firewalls.ComplexLineWeightActiveL || a.Firewalls.PRadHiggsMassTheorem || a.Firewalls.PRadYukawaTheorem || a.Firewalls.LHopfNativeHistoryLoopTheorem || a.Firewalls.HiggsMassOrPoleMassTheorem || a.Firewalls.YukawaOperatorOrEigenvalueTheorem {
		t.Fatalf("bad physical firewalls: %+v", a.Firewalls)
	}
}

func TestGate761TheoremVerdictStatuses(t *testing.T) {
	res := Generation2HiggsRadialProjectorAsGaugeFixedComplexVacuumLineAuditTheorem().Verify()
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
