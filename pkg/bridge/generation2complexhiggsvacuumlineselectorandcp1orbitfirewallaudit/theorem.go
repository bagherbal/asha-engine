package generation2complexhiggsvacuumlineselectorandcp1orbitfirewallaudit

import (
	"strings"

	"github.com/bagherbal/asha-engine/pkg/theorem"
)

func Generation2ComplexHiggsVacuumLineSelectorAndCP1OrbitFirewallAuditTheorem() theorem.Theorem {
	const id = AuditID
	const name = "Gate 762 — Complex Higgs Vacuum Line Selector and CP1 Orbit Firewall Audit"
	return theorem.Theorem{ID: id, Name: name, Layer: theorem.LayerBridge, Status: theorem.BridgeRequired, Verify: func() theorem.Result {
		a, err := BuildDefault()
		if err != nil {
			return theorem.Result{ID: id, Name: name, Layer: theorem.LayerBridge, Status: theorem.FailedRoute, Checks: []theorem.Check{{Name: "construct Gate762 CP1 selector audit", Passed: false, Detail: err.Error()}}}
		}
		checks := []theorem.Check{
			{Name: "inherit Gate761 radial-projector refinement", Passed: a.Gate761.Inherited && a.Gate761.PRadGaugeFixedInsideComplexLine && a.Gate761.ComplexLineNotNativelySelected && a.Gate761.RadialGaugeNotNativelySelected && a.Gate761.ActiveHistoryUsesRealRadialHalf && a.Gate761.RealRadialWeight == 0.25 && a.Gate761.FullComplexLineWeight == 0.50 && strings.Contains(a.Gate761.RefinedPRadTyping, "GaugeFixedRadialRepresentativeSeal") && strings.Contains(a.Gate761.ScalarVacuumSealSplit, "ComplexVacuumLineSeal"), Detail: FormatGate761(a.Gate761)},
			{Name: "promote complex vacuum line to primary target", Passed: a.Primary.NewPrimaryTarget == "Pi_vac_C" && a.Primary.ComplexLineBeforePRad && a.Primary.RadialGaugeAfterLine && a.Primary.LineLivesInK7PlusJ && a.Primary.LineIsCP1Point && strings.Contains(a.Primary.PreviousTarget, "P_rad") && strings.Contains(a.Primary.Reason, "pre-gauge"), Detail: FormatPrimary(a.Primary)},
			{Name: "record CP1 orbit geometry", Passed: a.Orbit.Recorded && a.Orbit.Socket == "K7+_J(n) ~= C^2" && strings.Contains(a.Orbit.ComplexLines, "CP1") && a.Orbit.CP1RealDimension == cp1RealDim && a.Orbit.UnitRepresentatives == "S3" && a.Orbit.S3RealDimension == s3RealDim && strings.Contains(a.Orbit.HopfFibration, "S1 -> S3 -> CP1") && a.Orbit.FiberDimension == s1RealDim && a.Orbit.BasePointSelectsLine && a.Orbit.FiberGaugeSelectsRadialRep, Detail: FormatOrbit(a.Orbit)},
			{Name: "formulate CP1 selector question", Passed: a.Question.Formulated && strings.Contains(a.Question.Question, "CP1") && strings.Contains(a.Question.RequiredObject, "Pi_vac_C") && a.Question.RequiredRankR == complexLineRealRank && a.Question.RequiredRankC == complexLineRank && a.Question.RequiresJH && a.Question.RequiresCP1BasePoint && a.Question.DoesNotRequireRadialGauge, Detail: FormatQuestion(a.Question)},
			{Name: "audit current source candidates", Passed: a.SourceAudit.Completed && len(a.SourceAudit.Candidates) == 7 && !a.SourceAudit.NativeComplexVacuumLineSelector && !a.SourceAudit.NativeCP1BasePointSelector && !a.SourceAudit.NativeRadialGaugeSelector && a.SourceAudit.NSelectsJHOnly && a.SourceAudit.RhoPlusNoBias && a.SourceAudit.PRadAssumesLineAndGauge && a.SourceAudit.AllCurrentCandidatesFailCP1 && strings.Contains(a.SourceAudit.Verdict, StatusNoNativeComplexVacuumLineSelector), Detail: FormatSourceAudit(a.SourceAudit)},
			{Name: "separate construction from supplied P_rad from native line selection", Passed: a.Constructible.CanConstructPiFromSuppliedPRad && strings.Contains(a.Constructible.ConstructionFormula, "J_H") && a.Constructible.ConstructionDependsOnSeal && !a.Constructible.ConstructionIsSelectionTheorem && a.Constructible.PRadMayNotBeUsedAsNativeCause && !a.Constructible.NativeLineSelectorCertified && strings.Contains(a.Constructible.Verdict, StatusPRadCannotBeUsedAsNativeLineSelector), Detail: FormatConstructible(a.Constructible)},
			{Name: "mark radial gauge fixing secondary", Passed: a.GaugeHierarchy.ComplexLineSeal == "ComplexVacuumLineSeal" && a.GaugeHierarchy.RadialGaugeFixingSeal == "RadialGaugeFixingSeal" && a.GaugeHierarchy.LineSelectionPrecedesGauge && a.GaugeHierarchy.GaugeWithoutLineIllTyped && a.GaugeHierarchy.PRadRequiresBothChoices && a.GaugeHierarchy.SecondaryMarked, Detail: FormatGaugeHierarchy(a.GaugeHierarchy)},
			{Name: "refine HistoryLoop dependency", Passed: a.HistoryLoop.Refined && strings.Contains(a.HistoryLoop.ActiveLHopfFormula, "P_rad") && a.HistoryLoop.DependsOnComplexLine && a.HistoryLoop.DependsOnRadialGauge && a.HistoryLoop.RealRadialWeight == 0.25 && a.HistoryLoop.ComplexLineWeight == 0.50 && a.HistoryLoop.FullComplexLineRejected && len(a.HistoryLoop.UnsolvedObjects) == 2, Detail: FormatHistoryLoop(a.HistoryLoop)},
			{Name: "audit layer separation", Passed: a.Layers.Audited && a.Layers.ObjectsNotSameOperator && strings.Contains(a.Layers.NLayer, "J_H") && strings.Contains(a.Layers.CP1Layer, "Pi_vac_C") && strings.Contains(a.Layers.RadialGaugeLayer, "P_rad") && strings.Contains(a.Layers.HistoryLayer, "L_Hopf") && strings.Contains(a.Layers.YukawaLayer, "N_eff"), Detail: FormatLayers(a.Layers)},
			{Name: "enforce physical firewalls", Passed: a.Firewalls.Audited && !a.Firewalls.PiVacCNativeVacuumTheorem && !a.Firewalls.CP1PointNativeEWSBTheorem && !a.Firewalls.RadialGaugeFixingNativeTheorem && !a.Firewalls.LHopfNativeHistoryLoopTheorem && !a.Firewalls.ComplexLineHiggsMassTheorem && !a.Firewalls.ComplexLineYukawaTheorem && !a.Firewalls.ScalarRuntimeIndependentTheorem && !a.Firewalls.HiggsMassOrPoleMassTheorem && !a.Firewalls.YukawaOperatorOrEigenvalueTheorem && strings.Contains(a.Firewalls.Verdict, StatusGate762ComplexVacuumLineCP1Boundary), Detail: FormatFirewalls(a.Firewalls)},
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
