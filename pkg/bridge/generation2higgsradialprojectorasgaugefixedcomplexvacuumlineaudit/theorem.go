package generation2higgsradialprojectorasgaugefixedcomplexvacuumlineaudit

import (
	"math"
	"strings"

	"github.com/bagherbal/asha-engine/pkg/theorem"
)

func Generation2HiggsRadialProjectorAsGaugeFixedComplexVacuumLineAuditTheorem() theorem.Theorem {
	const id = AuditID
	const name = "Gate 761 — Higgs Radial Projector as Gauge-Fixed Complex Vacuum Line Audit"
	return theorem.Theorem{ID: id, Name: name, Layer: theorem.LayerBridge, Status: theorem.BridgeRequired, Verify: func() theorem.Result {
		a, err := BuildDefault()
		if err != nil {
			return theorem.Result{ID: id, Name: name, Layer: theorem.LayerBridge, Status: theorem.FailedRoute, Checks: []theorem.Check{{Name: "construct Gate761 radial projector gauge-fixed complex-line audit", Passed: false, Detail: err.Error()}}}
		}
		checks := []theorem.Check{
			{Name: "inherit Gate760 master form and P_rad priority", Passed: a.Gate760.Inherited && a.Gate760.PRadPriorityInherited && a.Gate760.HighestPrioritySeal == "P_rad" && a.Gate760.LHopfDependsOnPRad && !a.Gate760.NativePRadSelector && !a.Gate760.NativeHistoryLoopTheorem && !a.Gate760.NativeScalarRuntimeTheorem && strings.Contains(a.Gate760.MasterFormula, "3/N_eff") && strings.Contains(a.Gate760.HighestPriorityReason, "P_rad"), Detail: FormatGate760(a.Gate760)},
			{Name: "inherit complex structure J_H(n)", Passed: a.Complex.JHSquaresMinusIdentity && a.Complex.JHSkewOrthogonal && a.Complex.K7PlusRealDimension == k7PlusRealDim && a.Complex.K7PlusComplexDimension == k7PlusComplexDim && a.Complex.K7PlusAsC2 && a.Complex.Gate726SplitInherited && strings.Contains(a.Complex.TwistorSelector, "S^2") && strings.Contains(a.Complex.ComplexStructure, "J_H"), Detail: FormatComplex(a.Complex)},
			{Name: "define real radial and phase directions", Passed: a.Directions.Defined && a.Directions.JHSkewOrthogonal && a.Directions.RadialPhaseOrthogonal && a.Directions.RadialRank == realRadialRank && a.Directions.PhaseRank == realPhaseRank && a.Directions.PhaseInAngularComplement && strings.Contains(a.Directions.PRadFormula, "v_rad") && strings.Contains(a.Directions.PhaseVector, "J_H"), Detail: FormatDirections(a.Directions)},
			{Name: "construct complex vacuum line from P_rad and J_H", Passed: a.VacuumLine.ConstructedFromPRadJH && a.VacuumLine.RealRank == complexLineRealRank && a.VacuumLine.ComplexRank == complexLineRank && a.VacuumLine.JInvariant && a.VacuumLine.CommutesWithJH && a.VacuumLine.InsideK7PlusC2 && a.VacuumLine.ContainsRadial && a.VacuumLine.ContainsPhase && strings.Contains(a.VacuumLine.Formula, "P_rad+P_phase"), Detail: FormatVacuumLine(a.VacuumLine)},
			{Name: "compute event weights", Passed: near(a.Weights.RadialWeight, 0.25, 1e-18) && near(a.Weights.PhaseWeight, 0.25, 1e-18) && near(a.Weights.ComplexLineWeight, 0.5, 1e-18) && near(a.Weights.TransverseWeight, 0.5, 1e-18) && near(a.Weights.TotalWeight, 1.0, 1e-18) && near(a.Weights.LHopfFromRadialEvent, 1/(8*math.Pi), 1e-18) && near(a.Weights.LoopFromComplexLine, 1/(4*math.Pi), 1e-18) && a.Weights.ActiveHistoryUsesRadial && a.Weights.ComplexLineTooLargeForL, Detail: FormatWeights(a.Weights)},
			{Name: "type P_rad as gauge-fixed radial representative", Passed: a.GaugeFixing.PRadInsideComplexLine && !a.GaugeFixing.PRadArbitraryPrimitiveLine && a.GaugeFixing.ComplexLineContainsHopfPhase && a.GaugeFixing.ScalarVacuumDirectionSealSplit && strings.Contains(a.GaugeFixing.RefinedTyping, "GaugeFixedRadialRepresentativeSeal") && strings.Contains(a.GaugeFixing.SealDecomposition, "ComplexVacuumLineSeal") && strings.Contains(a.GaugeFixing.Verdict, StatusPRadGaugeFixedRepresentativeInsideComplexLine), Detail: FormatGaugeFixing(a.GaugeFixing)},
			{Name: "separate n, complex line, and radial gauge representative", Passed: a.Selectors.NSelectsComplexStructure && !a.Selectors.NSelectsVacuumLine && !a.Selectors.ComplexLineSelectsRadialGauge && a.Selectors.PRadSelectsGaugeRepresentative && a.Selectors.ThreeChoicesDistinct && strings.Contains(a.Selectors.NRole, "J_H") && strings.Contains(a.Selectors.ComplexLineRole, "complex rank-one") && strings.Contains(a.Selectors.PRadRole, "radial representative"), Detail: FormatSelectors(a.Selectors)},
			{Name: "record U(2) Hopf orbit interpretation", Passed: a.Orbit.CP1BasePoint && a.Orbit.S1FiberGaugeRepresentative && a.Orbit.RealRadialAmplitudeAxis && a.Orbit.RefinesGate726S3Orbit && a.Orbit.Socket == "K7+_J(n) ~= C^2" && a.Orbit.ComplexVacuumLineOrbit == "CP1" && a.Orbit.UnitRepresentativeOrbit == "S3" && strings.Contains(a.Orbit.HopfFibration, "S1 -> S3 -> CP1"), Detail: FormatOrbit(a.Orbit)},
			{Name: "complete source-candidate audit", Passed: a.SourceAudit.Completed && len(a.SourceAudit.Candidates) == 6 && !a.SourceAudit.NativeComplexVacuumLineSelector && !a.SourceAudit.NativeRadialGaugeFixingSelector && !a.SourceAudit.RhoPlusSelectsLine && !a.SourceAudit.NSelectsLine && !a.SourceAudit.QSelectsLine && !a.SourceAudit.BoundaryScalarsSelectLine && !a.SourceAudit.FanoQuaternionicSelectsLine && strings.Contains(a.SourceAudit.Verdict, StatusRhoPlusDoesNotSelectVacuumLine), Detail: FormatSourceAudit(a.SourceAudit)},
			{Name: "refine HistoryLoop quarter-factor interpretation", Passed: a.HistoryLoop.ActiveUsesRealRadialAmplitude && a.HistoryLoop.FullComplexLineRejected && near(a.HistoryLoop.RadialQuarterWeight, 0.25, 1e-18) && near(a.HistoryLoop.FullComplexLineWeight, 0.5, 1e-18) && near(a.HistoryLoop.ActiveLoopUnit, 1/(8*math.Pi), 1e-18) && near(a.HistoryLoop.FullComplexLineLoopUnit, 1/(4*math.Pi), 1e-18) && strings.Contains(a.HistoryLoop.QuarterFactorInterpretation, "real radial amplitude"), Detail: FormatHistoryLoop(a.HistoryLoop)},
			{Name: "enforce physical firewalls", Passed: a.Firewalls.Audited && !a.Firewalls.PRadNativeVacuumTheorem && !a.Firewalls.ComplexLineNativeEWSBTheorem && !a.Firewalls.RadialGaugeFixingPhysicalEWSB && !a.Firewalls.ComplexLineWeightActiveL && !a.Firewalls.PRadHiggsMassTheorem && !a.Firewalls.PRadYukawaTheorem && !a.Firewalls.LHopfNativeHistoryLoopTheorem && !a.Firewalls.HiggsMassOrPoleMassTheorem && !a.Firewalls.YukawaOperatorOrEigenvalueTheorem && strings.Contains(a.Firewalls.Verdict, StatusGate761RadialProjectorGaugeFixedComplexLineBoundary), Detail: FormatFirewalls(a.Firewalls)},
		}
		ok := true
		for _, c := range checks {
			if !c.Passed {
				ok = false
				break
			}
		}
		status := theorem.BridgeRequired
		if !ok {
			status = theorem.FailedRoute
		}
		notes := append([]string{a.Truth}, Statuses()...)
		return theorem.Result{ID: id, Name: name, Layer: theorem.LayerBridge, Status: status, Checks: checks, Notes: notes}
	}}
}
