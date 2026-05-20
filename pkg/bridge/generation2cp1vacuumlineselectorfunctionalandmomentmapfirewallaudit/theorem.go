package generation2cp1vacuumlineselectorfunctionalandmomentmapfirewallaudit

import (
	"strings"

	"github.com/bagherbal/asha-engine/pkg/theorem"
)

func Generation2CP1VacuumLineSelectorFunctionalAndMomentMapFirewallAuditTheorem() theorem.Theorem {
	const id = AuditID
	const name = "Gate 763 — CP1 Vacuum-Line Selector Functional and Moment-Map Firewall Audit"
	return theorem.Theorem{ID: id, Name: name, Layer: theorem.LayerBridge, Status: theorem.BridgeRequired, Verify: func() theorem.Result {
		a, err := BuildDefault()
		if err != nil {
			return theorem.Result{ID: id, Name: name, Layer: theorem.LayerBridge, Status: theorem.FailedRoute, Checks: []theorem.Check{{Name: "construct Gate763 CP1 selector-functional audit", Passed: false, Detail: err.Error()}}}
		}
		checks := []theorem.Check{
			{Name: "inherit Gate762 complex vacuum-line seal", Passed: a.Gate762.Inherited && a.Gate762.Socket == "K7+_J(n) ~= C^2" && strings.Contains(a.Gate762.MissingObject, "Pi_vac_C") && a.Gate762.MissingObjectRankR == complexLineRealRank && a.Gate762.MissingObjectRankC == complexLineRank && a.Gate762.PRadSecondary && a.Gate762.NoCurrentCP1Selector && a.Gate762.ComplexVacuumLineSealRemains, Detail: FormatGate762(a.Gate762)},
			{Name: "define functional selector requirement", Passed: strings.Contains(a.Requirement.Question, "functional") && a.Requirement.Domain == "CP1 = P(C^2)" && strings.Contains(a.Requirement.Target, "Pi_vac_C") && strings.Contains(a.Requirement.EquivalentHermitianAxis, "su(2)") && a.Requirement.RequiresNonconstantFunctional && a.Requirement.RequiresIsolatedCriticalLine && !a.Requirement.U2InvariantFunctionalSelectsLine && a.Requirement.ConstantFunctionalLeavesCP1Flat && !a.Requirement.RadialGaugeRequiredForLineSelector, Detail: FormatRequirement(a.Requirement)},
			{Name: "audit CP1 moment-map candidate", Passed: a.MomentMap.MomentMapExistsOnCP1 && strings.Contains(a.MomentMap.MomentMapFormula, "zz^dagger") && strings.Contains(a.MomentMap.HamiltonianFormula, "<h") && a.MomentMap.RequiresSU2Axis && !a.MomentMap.SuppliedAxisCertified && a.MomentMap.WouldSelectEigenlineIfAxis && !a.MomentMap.NativeSelectorCertified && strings.Contains(a.MomentMap.Verdict, StatusNoNativeSU2MomentMapAxis), Detail: FormatMomentMap(a.MomentMap)},
			{Name: "audit scalar-potential functional candidate", Passed: strings.Contains(a.ScalarPotential.InvariantPotentialFormula, "|z|^2") && strings.Contains(a.ScalarPotential.AnisotropicPotentialFormula, "H") && a.ScalarPotential.U2Invariant && a.ScalarPotential.FlatOnCP1 && a.ScalarPotential.AnisotropicTermWouldSelectLine && !a.ScalarPotential.AnisotropicTermCertified && !a.ScalarPotential.NativeOrientationSelector && strings.Contains(a.ScalarPotential.Verdict, StatusScalarPotentialNotTypedOrientationSelector), Detail: FormatScalarPotential(a.ScalarPotential)},
			{Name: "audit boundary-history stress candidate", Passed: len(a.BoundaryStress.BoundaryObjects) == 5 && a.BoundaryStress.CollapsedScalarLayer && !a.BoundaryStress.ProvidesVectorInK7Plus && !a.BoundaryStress.ProvidesHermitianAxisOnC2 && !a.BoundaryStress.CanDefineCP1Functional && !a.BoundaryStress.CanSelectCP1Point && a.BoundaryStress.NeedsTypedVectorCoupling && strings.Contains(a.BoundaryStress.Verdict, StatusBoundaryStressScalarNotCP1Functional), Detail: FormatBoundaryStress(a.BoundaryStress)},
			{Name: "audit Fano/quaternionic invariant candidate", Passed: a.Fano.SuppliesTwistorFamily && a.Fano.SuppliesJHSocket && a.Fano.SuppliesU2Socket && !a.Fano.SelectsComplexStructureN && !a.Fano.SelectsCP1Point && !a.Fano.InvariantVectorInK7Plus && a.Fano.SymmetryWouldMakeLineArbitrary && strings.Contains(a.Fano.Verdict, StatusFanoQuaternionicDoesNotSelectCP1Point), Detail: FormatFano(a.Fano)},
			{Name: "audit orientation seal option", Passed: strings.Contains(a.Orientation.SealName, "ComplexVacuumLineSeal") && a.Orientation.CanSelectCP1IfSupplied && !a.Orientation.NativeTheoremCertified && a.Orientation.WouldPrecedeRadialGauge && a.Orientation.WouldNotDeriveEWSB && a.Orientation.WouldNotDeriveHiggsMass && strings.Contains(a.Orientation.Verdict, StatusPiVacCRemainsComplexVacuumLineSeal), Detail: FormatOrientation(a.Orientation)},
			{Name: "record candidate functional ranking", Passed: a.Ranking.RankingRecorded && len(a.Ranking.Candidates) == 5 && !a.Ranking.AnyNativeSelectorCertified && strings.Contains(a.Ranking.BestTypedFutureCandidate, "SU(2)") && strings.Contains(a.Ranking.HighestNativeResult, "no CP1 point") && strings.Contains(a.Ranking.Verdict, StatusNoNativeCP1SelectorFunctional), Detail: FormatRanking(a.Ranking)},
			{Name: "preserve line-before-gauge order", Passed: a.SealOrder.Preserved && a.SealOrder.LineBeforeGauge && a.SealOrder.LHopfAfterGauge && strings.Contains(a.SealOrder.Step1, "J_H") && strings.Contains(a.SealOrder.Step2, "CP1") && strings.Contains(a.SealOrder.Step3, "P_rad") && strings.Contains(a.SealOrder.Step4, "L_Hopf"), Detail: FormatSealOrder(a.SealOrder)},
			{Name: "enforce physical firewalls", Passed: a.Firewalls.Audited && !a.Firewalls.CP1FunctionalNativeSelector && !a.Firewalls.MomentMapAxisNative && !a.Firewalls.HermitianAxisNative && !a.Firewalls.ScalarPotentialNativeOrientation && !a.Firewalls.BoundaryStressCP1Functional && !a.Firewalls.PiVacCNativeEWSBTheorem && !a.Firewalls.RadialGaugeFixingNativeTheorem && !a.Firewalls.LHopfNativeHistoryLoopTheorem && !a.Firewalls.ScalarRuntimeIndependentTheorem && !a.Firewalls.HiggsMassOrPoleMassTheorem && !a.Firewalls.YukawaOperatorOrEigenvalueTheorem && strings.Contains(a.Firewalls.Verdict, StatusGate763CP1FunctionalSelectorBoundary), Detail: FormatFirewalls(a.Firewalls)},
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
