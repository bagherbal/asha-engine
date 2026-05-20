package generation2cp1gaugeorbitdemotionandradialrankinvarianthistoryloopaudit

import (
	"math"
	"strings"

	"github.com/bagherbal/asha-engine/pkg/theorem"
)

func Generation2CP1GaugeOrbitDemotionAndRadialRankInvariantHistoryLoopAuditTheorem() theorem.Theorem {
	const id = AuditID
	const name = "Gate 764 — CP1 Gauge-Orbit Demotion and Radial Rank-Invariant HistoryLoop Audit"
	return theorem.Theorem{ID: id, Name: name, Layer: theorem.LayerBridge, Status: theorem.BridgeRequired, Verify: func() theorem.Result {
		a, err := BuildDefault()
		if err != nil {
			return theorem.Result{ID: id, Name: name, Layer: theorem.LayerBridge, Status: theorem.FailedRoute, Checks: []theorem.Check{{Name: "build default audit", Passed: false, Detail: err.Error()}}}
		}
		near := func(x, y, tol float64) bool { return math.Abs(x-y) <= tol }
		checks := []theorem.Check{
			{Name: "inherit Gate763 CP1 selector firewall", Passed: a.Gate763.Inherited && a.Gate763.Socket == "K7+_J(n) ~= C^2" && a.Gate763.CP1 == "CP1 = P(C^2)" && strings.Contains(a.Gate763.PiVacC, "CP1 point") && strings.Contains(a.Gate763.PRad, "gauge-fixed") && strings.Contains(a.Gate763.LHopfFormula, "P_rad") && a.Gate763.NoNativeCP1Selector && a.Gate763.RequiresHermitianSU2Axis && a.Gate763.ComplexVacuumLineSeal && a.Gate763.RadialGaugeFixingSecondary, Detail: FormatGate763(a.Gate763)},
			{Name: "audit CP1 as SU2 orbit", Passed: strings.Contains(a.Orbit.Carrier, "C^2") && a.Orbit.OrbitFormula == "CP1 ~= SU(2)/U(1)" && a.Orbit.SU2ActsTransitivelyOnCP1 && a.Orbit.Stabilizer == "U(1)" && a.Orbit.CP1RealDimension == cp1RealDim && a.Orbit.SU2Dimension == su2Dim && a.Orbit.StabilizerDimension == u1StabilizerDim && a.Orbit.InvariantFunctionalConstant && a.Orbit.AnisotropicAxisRequiredToSelect && a.Orbit.ExplainsGate763Failure && a.Orbit.ConditionalOnGaugeAirlock && strings.Contains(a.Orbit.Verdict, StatusInternalCNotCertifiedAsPhysicalSU2L), Detail: FormatOrbit(a.Orbit)},
			{Name: "define gauge-orbit demotion", Passed: strings.Contains(a.Demotion.Object, "Pi_vac_C") && !a.Demotion.ScalarRuntimeNumericalSource && a.Demotion.GaugeRepresentative && a.Demotion.RequiresPhysicalAnisotropy && !a.Demotion.AbsenceOfSelectorScalarFailure && a.Demotion.AbsenceOfSelectorVacuumTheoremGap && a.Demotion.DemotionConditional && strings.Contains(a.Demotion.Verdict, StatusPiVacCIsGaugeRepresentativeNotScalarNumericalSource), Detail: FormatDemotion(a.Demotion)},
			{Name: "compute radial rank invariance", Passed: strings.Contains(a.RankInvariance.RhoPlusFormula, "rank(P_rad)/4") && a.RankInvariance.ProjectorRank == realRadialRank && a.RankInvariance.K7PlusDimension == k7PlusRealDim && near(a.RankInvariance.TraceWeight, 0.25, 1e-15) && near(a.RankInvariance.PhaseLoopPayoff, 1/(2*math.Pi), 1e-15) && near(a.RankInvariance.LHopf, a.RankInvariance.ExpectedLHopf, 1e-15) && near(a.RankInvariance.LHopf, 1/(8*math.Pi), 1e-15) && a.RankInvariance.DependsOnRankOnly && a.RankInvariance.IndependentOfCP1Position && a.RankInvariance.RankInvariant, Detail: FormatRankInvariance(a.RankInvariance)},
			{Name: "audit complex-line versus radial-event distinction", Passed: a.LineVsRadial.ComplexLineRealRank == complexLineRealRank && a.LineVsRadial.RadialEventRealRank == realRadialRank && near(a.LineVsRadial.ComplexLineWeight, 0.5, 1e-15) && near(a.LineVsRadial.RadialEventWeight, 0.25, 1e-15) && near(a.LineVsRadial.ComplexLineLoopUnit, 1/(4*math.Pi), 1e-15) && near(a.LineVsRadial.RadialEventLoopUnit, 1/(8*math.Pi), 1e-15) && a.LineVsRadial.ActiveUsesRadialEvent && a.LineVsRadial.FullComplexLineRejected && strings.Contains(a.LineVsRadial.ScalarRuntimeQuestion, "real rank-one radial"), Detail: FormatLineVsRadial(a.LineVsRadial)},
			{Name: "preserve conditional demotion firewall", Passed: strings.Contains(a.Conditional.InternalSocketName, "SU(2)") && a.Conditional.InternalSocketCertified && !a.Conditional.PhysicalSU2LCertified && !a.Conditional.CP1GaugeOrbitPhysicalTheorem && a.Conditional.DemotionAllowedAsInternalSocket && !a.Conditional.DemotionPromotedToPhysicalEWSB && strings.Contains(a.Conditional.Verdict, StatusInternalCNotCertifiedAsPhysicalSU2L), Detail: FormatConditional(a.Conditional)},
			{Name: "record updated missing-object hierarchy", Passed: a.Hierarchy.Updated && strings.Contains(a.Hierarchy.BeforeGate764PrimaryMissing, "ComplexVacuumLineSeal") && strings.Contains(a.Hierarchy.AfterGate764ScalarTarget, "radial event type") && strings.Contains(a.Hierarchy.AfterGate764PhysicalTarget, "electroweak") && !a.Hierarchy.CP1LocationScalarDatum && a.Hierarchy.RadialEventTypeScalarDatum && a.Hierarchy.ElectroweakVacuumStillOpen && strings.Contains(a.Hierarchy.Verdict, StatusNextScalarSourceTargetRadialEventTypeSelection), Detail: FormatHierarchy(a.Hierarchy)},
			{Name: "enforce physical firewalls", Passed: a.Firewalls.Audited && !a.Firewalls.CP1PointPhysicalTheorem && !a.Firewalls.AbsenceOfCP1SelectorScalarFailure && !a.Firewalls.CP1GaugeOrbitPhysicalSU2LTheorem && !a.Firewalls.RankOneRadialEventNativeTheorem && !a.Firewalls.LHopfNativeTransportTheorem && !a.Firewalls.ScalarRuntimeIndependentTheorem && !a.Firewalls.TreeProxyPoleMassTheorem && !a.Firewalls.ElectroweakSymmetryBreakingTheorem && !a.Firewalls.HiggsMassOrPoleMassTheorem && !a.Firewalls.YukawaOperatorOrEigenvalueTheorem && strings.Contains(a.Firewalls.Verdict, StatusGate764CP1GaugeOrbitRadialRankBoundary), Detail: FormatFirewalls(a.Firewalls)},
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
