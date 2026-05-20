package generation2radialphasehopffiberandangularcomplementdecompositionaudit

import (
	"math"
	"strings"

	"github.com/bagherbal/asha-engine/pkg/theorem"
)

func Generation2RadialPhaseHopfFiberAndAngularComplementDecompositionAuditTheorem() theorem.Theorem {
	const id = AuditID
	const name = "Gate 726 — Radial-Phase Hopf Fiber and Angular Complement Decomposition Audit"
	return theorem.Theorem{ID: id, Name: name, Layer: theorem.LayerBridge, Status: theorem.BridgeRequired, Verify: func() theorem.Result {
		a, err := BuildDefault()
		if err != nil {
			return theorem.Result{ID: id, Name: name, Layer: theorem.LayerBridge, Status: theorem.FailedRoute, Checks: []theorem.Check{{Name: "construct Gate726 radial phase Hopf audit", Passed: false, Detail: err.Error()}}}
		}
		checks := []theorem.Check{
			{Name: "inherit Gate725 radial orbit audit", Passed: a.Gate725.Inherited && a.Gate725.RadialDecompositionDefined && a.Gate725.RadialEventWeightComputed && a.Gate725.U2OrbitShadowAudited && a.Gate725.NoNativeRadialProjectorSelector && a.Gate725.NoNativeEWSBTheorem && a.Gate725.NoPhysicalGoldstoneIdentification && a.Gate725.NoNativeHistoryLoopUnit && a.Gate725.NoHiggsMassTheorem && a.Gate725.NoYukawaTheorem && a.Gate725.Verdict == StatusGate725RadialGoldstoneOrbitInherited, Detail: FormatGate725(a.Gate725)},
			{Name: "define phase direction from radial line", Passed: a.Phase.JHSquaresMinusIdentity && a.Phase.JHSkewOrthogonal && a.Phase.OrthogonalToRadial && a.Phase.LiesInAngularComplement && a.Phase.PhaseRank == phaseRank && a.Phase.ProjectorOrthogonalToRadial && strings.Contains(a.Phase.Verdict, StatusPhaseDirectionFromRadialLineDefined), Detail: FormatPhase(a.Phase)},
			{Name: "compute radial phase transverse decomposition", Passed: a.Decomposition.RadialRank == radialRank && a.Decomposition.PhaseRank == phaseRank && a.Decomposition.TransverseRank == transverseRank && a.Decomposition.AngularRank == angularRank && a.Decomposition.CarrierDimension == k7PlusRealDim && a.Decomposition.DirectSum && a.Decomposition.AngularSplitsOneTwo && strings.Contains(a.Decomposition.Verdict, StatusRadialPhaseTransverseDecompositionComputed), Detail: FormatDecomposition(a.Decomposition)},
			{Name: "audit Hopf fiber phase loop", Passed: near(a.Hopf.PhaseUnit, 1/(2*math.Pi), 1e-18) && a.Hopf.UsesRadialEvent && !a.Hopf.AbstractPhaseOnly && strings.Contains(a.Hopf.Verdict, StatusOneOverTwoPiIsPhaseLoopUnitOnRadialHopfFiber), Detail: FormatHopf(a.Hopf)},
			{Name: "compute radial phase transverse event weights", Passed: near(a.Weights.RadialProbability, 0.25, 1e-18) && near(a.Weights.PhaseProbability, 0.25, 1e-18) && near(a.Weights.TransverseProbability, 0.5, 1e-18) && near(a.Weights.TotalProbability, 1, 1e-18) && a.Weights.HistoryLoopUsesRadialWeight && strings.Contains(a.Weights.Verdict, StatusRadialPhaseTransverseEventWeightsComputed), Detail: FormatWeights(a.Weights)},
			{Name: "audit U2 orbit Hopf 1+2 structure", Passed: a.Orbit.U2Dimension == u2Dimension && a.Orbit.StabilizerDimension == stabilizerDimension && a.Orbit.OrbitDimension == angularRank && a.Orbit.PhaseFiberDimension == phaseRank && a.Orbit.ProjectiveTransverseDimension == transverseRank && a.Orbit.SplitsAsOnePlusTwo && strings.Contains(a.Orbit.Verdict, StatusU2OrbitHopfStructureAudited), Detail: FormatOrbit(a.Orbit)},
			{Name: "audit selector dependence", Passed: a.Selectors.RequiresTwistorSelectorN && a.Selectors.RequiresRadialProjectorPRad && !a.Selectors.NAloneSelectsRadialLine && !a.Selectors.PRadAloneSelectsComplexPhase && a.Selectors.IndependentSeals && strings.Contains(a.Selectors.Verdict, StatusNAloneDoesNotSelectRadialLine), Detail: FormatSelectors(a.Selectors)},
			{Name: "enforce Gate726 physical firewalls", Passed: !a.Firewall.NativeRadialProjectorSelector && !a.Firewall.NativeTwistorSelectorN && !a.Firewall.NativeEWSBTheorem && !a.Firewall.PhysicalGoldstoneIdentification && !a.Firewall.HopfFiberAsPhysicalTimeOrRG && !a.Firewall.NativeHistoryLoopUnitSourceTheorem && !a.Firewall.HiggsMassOrPoleMassTheorem && !a.Firewall.YukawaOperatorOrEigenvalueTheorem && strings.Contains(a.Firewall.Verdict, StatusGate726Boundary), Detail: FormatFirewall(a.Firewall)},
		}
		notes := append(Statuses(), a.Truth)
		return theorem.Result{ID: id, Name: name, Layer: theorem.LayerBridge, Status: theorem.BridgeRequired, Checks: checks, Notes: notes}
	}}
}
