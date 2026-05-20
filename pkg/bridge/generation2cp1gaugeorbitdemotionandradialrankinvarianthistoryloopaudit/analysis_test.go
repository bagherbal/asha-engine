package generation2cp1gaugeorbitdemotionandradialrankinvarianthistoryloopaudit

import (
	"math"
	"strings"
	"testing"
)

func near(x, y float64) bool { return math.Abs(x-y) <= 1e-15 }

func TestGate764InheritanceOrbitAndDemotion(t *testing.T) {
	a, err := BuildDefault()
	if err != nil {
		t.Fatal(err)
	}
	if !a.Gate763.Inherited || a.Gate763.Socket != "K7+_J(n) ~= C^2" || a.Gate763.CP1 != "CP1 = P(C^2)" || !strings.Contains(a.Gate763.PiVacC, "CP1 point") || !strings.Contains(a.Gate763.PRad, "gauge-fixed") || !a.Gate763.NoNativeCP1Selector || !a.Gate763.RequiresHermitianSU2Axis || !a.Gate763.ComplexVacuumLineSeal || !a.Gate763.RadialGaugeFixingSecondary {
		t.Fatalf("bad Gate763 inheritance: %+v", a.Gate763)
	}
	if a.Orbit.OrbitFormula != "CP1 ~= SU(2)/U(1)" || !a.Orbit.SU2ActsTransitivelyOnCP1 || a.Orbit.Stabilizer != "U(1)" || a.Orbit.CP1RealDimension != cp1RealDim || a.Orbit.SU2Dimension != su2Dim || a.Orbit.StabilizerDimension != u1StabilizerDim || !a.Orbit.InvariantFunctionalConstant || !a.Orbit.AnisotropicAxisRequiredToSelect || !a.Orbit.ExplainsGate763Failure || !a.Orbit.ConditionalOnGaugeAirlock {
		t.Fatalf("bad SU2 orbit audit: %+v", a.Orbit)
	}
	if !strings.Contains(a.Demotion.Object, "Pi_vac_C") || a.Demotion.ScalarRuntimeNumericalSource || !a.Demotion.GaugeRepresentative || !a.Demotion.RequiresPhysicalAnisotropy || a.Demotion.AbsenceOfSelectorScalarFailure || !a.Demotion.AbsenceOfSelectorVacuumTheoremGap || !a.Demotion.DemotionConditional {
		t.Fatalf("bad demotion audit: %+v", a.Demotion)
	}
}

func TestGate764RankInvarianceAndLineRadialDistinction(t *testing.T) {
	a, err := BuildDefault()
	if err != nil {
		t.Fatal(err)
	}
	if a.RankInvariance.ProjectorRank != realRadialRank || a.RankInvariance.K7PlusDimension != k7PlusRealDim || !near(a.RankInvariance.TraceWeight, 0.25) || !near(a.RankInvariance.PhaseLoopPayoff, 1/(2*math.Pi)) || !near(a.RankInvariance.LHopf, 1/(8*math.Pi)) || !near(a.RankInvariance.LHopf, a.RankInvariance.ExpectedLHopf) || !a.RankInvariance.DependsOnRankOnly || !a.RankInvariance.IndependentOfCP1Position || !a.RankInvariance.RankInvariant {
		t.Fatalf("bad radial rank-invariance audit: %+v", a.RankInvariance)
	}
	if a.LineVsRadial.ComplexLineRealRank != complexLineRealRank || a.LineVsRadial.RadialEventRealRank != realRadialRank || !near(a.LineVsRadial.ComplexLineWeight, 0.5) || !near(a.LineVsRadial.RadialEventWeight, 0.25) || !near(a.LineVsRadial.ComplexLineLoopUnit, 1/(4*math.Pi)) || !near(a.LineVsRadial.RadialEventLoopUnit, 1/(8*math.Pi)) || !a.LineVsRadial.ActiveUsesRadialEvent || !a.LineVsRadial.FullComplexLineRejected || !strings.Contains(a.LineVsRadial.ScalarRuntimeQuestion, "real rank-one radial") {
		t.Fatalf("bad line-vs-radial audit: %+v", a.LineVsRadial)
	}
}

func TestGate764ConditionalFirewallAndHierarchy(t *testing.T) {
	a, err := BuildDefault()
	if err != nil {
		t.Fatal(err)
	}
	if !strings.Contains(a.Conditional.InternalSocketName, "SU(2)") || !a.Conditional.InternalSocketCertified || a.Conditional.PhysicalSU2LCertified || a.Conditional.CP1GaugeOrbitPhysicalTheorem || !a.Conditional.DemotionAllowedAsInternalSocket || a.Conditional.DemotionPromotedToPhysicalEWSB {
		t.Fatalf("bad conditional firewall: %+v", a.Conditional)
	}
	if !a.Hierarchy.Updated || !strings.Contains(a.Hierarchy.BeforeGate764PrimaryMissing, "ComplexVacuumLineSeal") || !strings.Contains(a.Hierarchy.AfterGate764ScalarTarget, "radial event type") || !strings.Contains(a.Hierarchy.AfterGate764PhysicalTarget, "electroweak") || a.Hierarchy.CP1LocationScalarDatum || !a.Hierarchy.RadialEventTypeScalarDatum || !a.Hierarchy.ElectroweakVacuumStillOpen {
		t.Fatalf("bad hierarchy update: %+v", a.Hierarchy)
	}
}

func TestGate764FirewallsAndTheoremStatuses(t *testing.T) {
	a, err := BuildDefault()
	if err != nil {
		t.Fatal(err)
	}
	if !a.Firewalls.Audited || a.Firewalls.CP1PointPhysicalTheorem || a.Firewalls.AbsenceOfCP1SelectorScalarFailure || a.Firewalls.CP1GaugeOrbitPhysicalSU2LTheorem || a.Firewalls.RankOneRadialEventNativeTheorem || a.Firewalls.LHopfNativeTransportTheorem || a.Firewalls.ScalarRuntimeIndependentTheorem || a.Firewalls.TreeProxyPoleMassTheorem || a.Firewalls.ElectroweakSymmetryBreakingTheorem || a.Firewalls.HiggsMassOrPoleMassTheorem || a.Firewalls.YukawaOperatorOrEigenvalueTheorem {
		t.Fatalf("bad firewalls: %+v", a.Firewalls)
	}
	res := Generation2CP1GaugeOrbitDemotionAndRadialRankInvariantHistoryLoopAuditTheorem().Verify()
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
