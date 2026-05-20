package generation2boundarystresssplitlinepullbacksourceaudit

import (
	"math"
	"strings"

	"github.com/bagherbal/asha-engine/pkg/theorem"
)

func Generation2BoundaryStressSplitLinePullbackSourceAuditTheorem() theorem.Theorem {
	const id = AuditID
	const name = "Gate 673 — BoundaryStressSplit Line-Pullback Source Audit"
	return theorem.Theorem{ID: id, Name: name, Layer: theorem.LayerBridge, Status: theorem.BridgeRequired, Verify: func() theorem.Result {
		a, err := BuildDefault()
		if err != nil {
			return theorem.Result{ID: id, Name: name, Layer: theorem.LayerBridge, Status: theorem.FailedRoute, Checks: []theorem.Check{{Name: "construct Gate673 stress-split line-pullback audit", Passed: false, Detail: err.Error()}}}
		}
		checks := []theorem.Check{
			{Name: "inherit Gate672 stress-split pullback firewall", Passed: a.Inherited.InheritedStressSplitPullback && a.Inherited.BaseClosureComputed && a.Inherited.StressSplitComputed && a.Inherited.NoNativeStressSplitTheorem && a.Inherited.NoNativeSevenOver72Theorem && a.Inherited.NoWallDistanceAirlockTheorem && a.Inherited.NoBoundaryStressDerivation && a.Inherited.FirewallPreserved, Detail: FormatInherited(a.Inherited)},
			{Name: "define boundary split line", Passed: math.Abs(a.BoundaryLine.SSplit-0.0012924448188163) < 1e-14 && a.BoundaryLine.Verdict == StatusBoundarySplitLineDefined, Detail: FormatBoundaryLine(a.BoundaryLine)},
			{Name: "define scalar/flavor base-defect line", Passed: math.Abs(a.BaseLine.DBase-0.00012565520996836) < 1e-14 && a.BaseLine.Verdict == StatusScalarFlavorBaseDefectLineDefined, Detail: FormatBaseLine(a.BaseLine)},
			{Name: "compute q_pull and typed candidates", Passed: math.Abs(a.Coefficient.QPull-0.0972228818894104) < 1e-12 && a.Coefficient.BestTypedCandidate == "7/72" && math.Abs(a.Coefficient.SevenOver72Residual-8.52583727234e-10) < 1e-14 && len(a.Coefficient.Candidates) == 6 && strings.Contains(a.Coefficient.Verdict, StatusSevenOver72StressSplitLinePullback), Detail: FormatCoefficient(a.Coefficient)},
			{Name: "audit line-map source types", Passed: len(a.Source.CandidateSupport) == 3 && len(a.Source.MissingTheorems) == 4 && strings.Contains(a.Source.Verdict, StatusLinePullbackSharperThanFullMap) && strings.Contains(a.Source.Verdict, StatusAugmentedChamberTraceCandidate), Detail: FormatSource(a.Source)},
			{Name: "preserve full-boundary-map firewall", Passed: a.Firewall.FullK7ToBoundaryMapFailed && a.Firewall.FanoHitchinRouteRemainsSealed && a.Firewall.LinePullbackStillPossible && a.Firewall.Verdict == StatusFullBoundaryMapFirewallAudited, Detail: FormatFirewall(a.Firewall)},
			{Name: "audit Lambda12-local line coefficient", Passed: a.ScaleLocal.Lambda12Local && a.ScaleLocal.CrossingBased && a.ScaleLocal.StationarityRejected && a.ScaleLocal.QPullNearSevenOver72OnlyAtLambda12 && a.ScaleLocal.Verdict == StatusScaleLocalityAudited, Detail: FormatScale(a.ScaleLocal)},
			{Name: "preserve firewalls", Passed: !a.Discipline.ClaimsNativeStressSplitPullback && !a.Discipline.ClaimsNativeSevenOver72 && !a.Discipline.ClaimsFullK7BoundaryMap && !a.Discipline.ClaimsWallDistanceAirlock && !a.Discipline.ClaimsBoundaryStressDerivation && !a.Discipline.ClaimsHiggsMassPrediction && !a.Discipline.ClaimsScalarStability && !a.Discipline.ClaimsGaugeUnification && !a.Discipline.ClaimsFlavorDerivation && !a.Discipline.ClaimsCKMPMNSDerivation && a.Discipline.Verdict == StatusGate673Boundary, Detail: FormatDiscipline(a.Discipline)},
		}
		notes := append(Statuses(), a.Truth)
		return theorem.Result{ID: id, Name: name, Layer: theorem.LayerBridge, Status: theorem.BridgeRequired, Checks: checks, Notes: notes}
	}}
}
