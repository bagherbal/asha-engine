package generation2higgsradialprojectorandgoldstonecomplementorbitaudit

import (
	"math"
	"strings"

	"github.com/bagherbal/asha-engine/pkg/theorem"
)

func Generation2HiggsRadialProjectorAndGoldstoneComplementOrbitAuditTheorem() theorem.Theorem {
	const id = AuditID
	const name = "Gate 725 — Higgs Radial Projector and Goldstone-Complement Orbit Audit"
	return theorem.Theorem{ID: id, Name: name, Layer: theorem.LayerBridge, Status: theorem.BridgeRequired, Verify: func() theorem.Result {
		a, err := BuildDefault()
		if err != nil {
			return theorem.Result{ID: id, Name: name, Layer: theorem.LayerBridge, Status: theorem.FailedRoute, Checks: []theorem.Check{{Name: "construct Gate725 radial projector orbit audit", Passed: false, Detail: err.Error()}}}
		}
		checks := []theorem.Check{
			{Name: "inherit Gate724 radial event phase-loop audit", Passed: a.Gate724.Inherited && a.Gate724.RhoPlusDefined && a.Gate724.RadialEventWeightComputed && a.Gate724.PhaseLoopExpectationReproducesL && a.Gate724.NoNativeRadialProjectorSelector && a.Gate724.TwistorSelectorDoesNotSelectRadialEvent && a.Gate724.QDoesNotSourceL && a.Gate724.NoNativeHistoryLoopUnit && a.Gate724.NoHiggsMassTheorem && a.Gate724.NoYukawaTheorem && a.Gate724.Verdict == StatusGate724HiggsRadialEventPhaseLoopInherited, Detail: FormatGate724(a.Gate724)},
			{Name: "define radial projector decomposition", Passed: a.Decomposition.RadialRank == radialRank && a.Decomposition.AngularRank == angularRank && a.Decomposition.CarrierDimension == k7PlusRealDim && a.Decomposition.DirectSum && a.Decomposition.Orthogonal && strings.Contains(a.Decomposition.Verdict, StatusRadialProjectorDecompositionDefined), Detail: FormatDecomposition(a.Decomposition)},
			{Name: "compute radial and complement event weights", Passed: near(a.Weights.RadialProbability, 0.25, 1e-18) && near(a.Weights.AngularProbability, 0.75, 1e-18) && near(a.Weights.Sum, 1, 1e-18) && strings.Contains(a.Weights.Verdict, StatusRadialAndComplementEventWeightsComputed), Detail: FormatWeights(a.Weights)},
			{Name: "audit U2 orbit-stabilizer geometry", Passed: a.Orbit.U2Dimension == u2Dimension && a.Orbit.StabilizerDimension == radialStabilizerDimension && a.Orbit.OrbitDimension == angularRank && a.Orbit.MatchesAngularComplementRank && !a.Orbit.PhysicalEWSBTheorem && !a.Orbit.PhysicalGoldstoneIdentification && strings.Contains(a.Orbit.Verdict, StatusU2OrbitStabilizerGeometryAudited), Detail: FormatOrbit(a.Orbit)},
			{Name: "audit radial selector source candidates", Passed: !a.Sources.TwistorSelectorSelectsPRad && !a.Sources.QSelectsPRad && !a.Sources.ScalarWallSelectsPRad && !a.Sources.BoundarySplitSelectsPRad && !a.Sources.K7EventProjectorSelectsPRad && !a.Sources.FanoHodgeSelectsPRad && !a.Sources.NativeRadialSelectorFound && strings.Contains(a.Sources.Verdict, StatusNoNativeRadialProjectorSelector), Detail: FormatSources(a.Sources)},
			{Name: "classify type-distinct scalar radial seal", Passed: len(a.Seal.SealNames) == 3 && a.Seal.TypeDistinctFromTwistorSelector && a.Seal.TypeDistinctFromHyperchargeNormalization && a.Seal.Verdict == StatusPRadIsTypeDistinctScalarVacuumDirectionSealCandidate, Detail: FormatSeal(a.Seal)},
			{Name: "preserve HistoryLoop relation and firewalls", Passed: near(a.HistoryLoop.HistoryLoopUnit, 1/(8*math.Pi), 1e-18) && near(a.HistoryLoop.RadialWeight, 0.25, 1e-18) && near(a.HistoryLoop.PhaseLoopPayoff, 1/(2*math.Pi), 1e-18) && a.HistoryLoop.ReproducesL && !a.HistoryLoop.NativeHistoryLoopUnitSourceTheorem && !a.HistoryLoop.NativePhasePayoffTransportTheorem && strings.Contains(a.HistoryLoop.Verdict, StatusNoNativeHistoryLoopUnitSourceTheorem), Detail: FormatHistoryLoop(a.HistoryLoop)},
			{Name: "enforce Gate725 physical firewalls", Passed: !a.Firewall.NativeRadialProjectorSelector && !a.Firewall.TwistorSelectorSelectsPRad && !a.Firewall.QSelectsPRad && !a.Firewall.NativeEWSBTheorem && !a.Firewall.PhysicalGoldstoneIdentification && !a.Firewall.NativeHistoryLoopUnitSourceTheorem && !a.Firewall.HiggsMassOrPoleMassTheorem && !a.Firewall.YukawaOperatorOrEigenvalueTheorem && strings.Contains(a.Firewall.Verdict, StatusGate725Boundary), Detail: FormatFirewall(a.Firewall)},
		}
		notes := append(Statuses(), a.Truth)
		return theorem.Result{ID: id, Name: name, Layer: theorem.LayerBridge, Status: theorem.BridgeRequired, Checks: checks, Notes: notes}
	}}
}
