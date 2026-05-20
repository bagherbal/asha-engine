package generation2higgsradialeventweightandphaselooptransportaudit

import (
	"math"
	"strings"

	"github.com/bagherbal/asha-engine/pkg/theorem"
)

func Generation2HiggsRadialEventWeightAndPhaseLoopTransportAuditTheorem() theorem.Theorem {
	const id = AuditID
	const name = "Gate 724 — Higgs Radial Event Weight and PhaseLoop Transport Audit"
	return theorem.Theorem{ID: id, Name: name, Layer: theorem.LayerBridge, Status: theorem.BridgeRequired, Verify: func() theorem.Result {
		a, err := BuildDefault()
		if err != nil {
			return theorem.Result{ID: id, Name: name, Layer: theorem.LayerBridge, Status: theorem.FailedRoute, Checks: []theorem.Check{{Name: "construct Gate724 radial event phaseloop audit", Passed: false, Detail: err.Error()}}}
		}
		checks := []theorem.Check{
			{Name: "inherit Gate723 quarter phase transport", Passed: a.Gate723.Inherited && a.Gate723.QuarterPhaseCandidate && a.Gate723.PhaseLoopMeasureCandidate && a.Gate723.FourComponentAverageCandidate && a.Gate723.ScalarTransportNotRepresentationLayer && a.Gate723.NoNativeHistoryLoopUnit && a.Gate723.NoNativePhaseLoopMeasure && a.Gate723.NoScalarProxyRuntimeTheorem && a.Gate723.QDoesNotSourceL && a.Gate723.LDoesNotSelectN && a.Gate723.SevenOver72DoesNotSourceL && a.Gate723.Verdict == StatusGate723QuarterPhaseTransportInherited, Detail: FormatGate723(a.Gate723)},
			{Name: "define rho_plus", Passed: a.RhoPlus.CarrierDimension == k7PlusRealDim && near(a.RhoPlus.Trace, 1, 1e-18) && a.RhoPlus.MaximallyMixed && strings.Contains(a.RhoPlus.StateFormula, "I_K7+/4") && a.RhoPlus.Verdict == StatusRhoPlusDefined, Detail: FormatRhoPlus(a.RhoPlus)},
			{Name: "compute rank-one radial event weight", Passed: a.RadialEvent.Rank == radialRank && a.RadialEvent.ProjectorIdempotent && a.RadialEvent.ProjectorSymmetric && a.RadialEvent.ActsInsideK7Plus && near(a.RadialEvent.Weight, 0.25, 1e-18) && !a.RadialEvent.NativeSelector && strings.Contains(a.RadialEvent.Verdict, StatusRankOneRadialEventWeightComputed), Detail: FormatRadial(a.RadialEvent)},
			{Name: "define phase-loop payoff observable", Passed: near(a.PhasePayoff.PhaseLoopPayoff, 1/(2*math.Pi), 1e-18) && strings.Contains(a.PhasePayoff.ObservableFormula, "P_rad") && strings.Contains(a.PhasePayoff.PhaseAction, "exp") && strings.Contains(a.PhasePayoff.Verdict, StatusPhaseLoopPayoffObservableDefined), Detail: FormatPhasePayoff(a.PhasePayoff)},
			{Name: "expectation reproduces one over 8 pi", Passed: near(a.PhasePayoff.Expectation, 1/(8*math.Pi), 1e-18) && near(a.PhasePayoff.HistoryLoopUnit, 1/(8*math.Pi), 1e-18) && a.PhasePayoff.ExpectationMatchesL && !a.PhasePayoff.NativePayoffTheorem && strings.Contains(a.PhasePayoff.Verdict, StatusExpectationReproducesOneOver8Pi), Detail: FormatPhasePayoff(a.PhasePayoff)},
			{Name: "audit alternative ranks", Passed: len(a.Ranks.Alternatives) == 4 && a.Ranks.ActiveName == "rank-one radial event" && oneActiveRank(a.Ranks.Alternatives, radialRank) && strings.Contains(a.Ranks.Verdict, StatusAlternativeRanksAudited) && strings.Contains(a.Ranks.Verdict, StatusRankOneRadialEventBestTypedQuarterSource), Detail: FormatRanks(a.Ranks)},
			{Name: "audit radial, twistor, and q firewalls", Passed: !a.Selectors.P_radNativelySelected && !a.Selectors.TwistorSelectorSelectsRadial && !a.Selectors.QSourcesRadialOrL && strings.Contains(a.Selectors.Verdict, StatusNoNativeRadialProjectorSelector) && strings.Contains(a.Selectors.Verdict, StatusTwistorSelectorDoesNotSelectRadialEvent) && strings.Contains(a.Selectors.Verdict, StatusQDoesNotSourceL), Detail: FormatSelectors(a.Selectors)},
			{Name: "audit event-weight analogy to 7 over 72", Passed: near(a.Analogy.K7EventProbability, float64(k7Dim)/float64(h72Dim), 1e-18) && near(a.Analogy.RadialEventProbability, 0.25, 1e-18) && a.Analogy.AnalogousEventWeights && !a.Analogy.SevenOver72SourcesQuarter && !a.Analogy.SevenOver72SourcesL && strings.Contains(a.Analogy.Verdict, StatusEventWeightAnalogyHistoryLoopK7Response), Detail: FormatAnalogy(a.Analogy)},
			{Name: "preserve scalar transport placement", Passed: a.Placement.BelongsAfterScalarProxy && !a.Placement.DerivedFromRepresentationAlone && a.Placement.NoScalarProxyRuntimeTheorem && strings.Contains(a.Placement.Verdict, StatusNoNativeScalarProxyToRuntimeTheorem), Detail: FormatPlacement(a.Placement)},
			{Name: "preserve Gate724 firewalls", Passed: !a.Firewall.NativeRadialProjectorSelector && !a.Firewall.NativeHistoryLoopUnitSourceTheorem && !a.Firewall.NativePhaseLoopPayoffTheorem && !a.Firewall.TwistorSelectorSelectsRadialEvent && !a.Firewall.QSourcesL && !a.Firewall.SevenOver72SourcesOneOver8Pi && !a.Firewall.NativeScalarProxyToRuntimeTheorem && !a.Firewall.HiggsMassOrPoleMassTheorem && !a.Firewall.YukawaOperatorOrEigenvalueTheorem && strings.Contains(a.Firewall.Verdict, StatusGate724Boundary), Detail: FormatFirewall(a.Firewall)},
		}
		notes := append(Statuses(), a.Truth)
		return theorem.Result{ID: id, Name: name, Layer: theorem.LayerBridge, Status: theorem.BridgeRequired, Checks: checks, Notes: notes}
	}}
}

func oneActiveRank(xs []AlternativeRank, rank int) bool {
	count := 0
	for _, x := range xs {
		if x.Active {
			count++
			if x.Rank != rank || !near(x.Value, 1/(8*math.Pi), 1e-18) {
				return false
			}
		}
	}
	return count == 1
}
