package generation2internalobstructionsealclosurepivot

import (
	"strings"

	"github.com/bagherbal/asha-engine/pkg/theorem"
)

func Generation2InternalObstructionSealClosureAndActiveBoundaryTransportPivotAuditTheorem() theorem.Theorem {
	const id = AuditID
	const name = "Gate 657 — Internal Obstruction Seal Closure and Active Boundary-Transport Pivot Audit"
	return theorem.Theorem{ID: id, Name: name, Layer: theorem.LayerBridge, Status: theorem.BridgeRequired, Verify: func() theorem.Result {
		a, err := BuildDefault()
		if err != nil {
			return theorem.Result{ID: id, Name: name, Layer: theorem.LayerBridge, Status: theorem.FailedRoute, Checks: []theorem.Check{{Name: "construct Gate657 closure/pivot audit", Passed: false, Detail: err.Error()}}}
		}
		checks := []theorem.Check{
			{Name: "inherit Gate656 half-trace non-certification", Passed: a.Inherited.HalfTraceAudited && a.Inherited.HalfTraceTypedClue && a.Inherited.FanoNumeratorStrengthened && a.Inherited.NoNativeHalfTraceMap && a.Inherited.NoSevenOver144Theorem && a.Inherited.NoSevenOver72Theorem && a.Inherited.NoBoundaryStressFromK7 && a.Inherited.NoBoundaryStressDerived && a.Inherited.NoHistoryLoopSource && a.Inherited.NoScalarFlavorMap && !a.Inherited.ClaimsBoundaryStress && !a.Inherited.ClaimsSevenOver144 && !a.Inherited.ClaimsSevenOver72 && !a.Inherited.ClaimsHistoryLoopUnit && !a.Inherited.ClaimsScalarFlavor && a.Inherited.Gate656Firewall, Detail: FormatInherited(a.Inherited)},
			{Name: "classify Fano-Hitchin lane as internal and boundary-closed", Passed: a.Closure.InternalTheoremPathMature && a.Closure.BoundaryInterfaceFailed && a.Closure.PhysicsPromotionBlocked && a.Closure.FutureUseRequiresExplicitPsi && a.Closure.SealName == "FanoHitchinObstructionSeal", Detail: FormatClosure(a.Closure)},
			{Name: "rebuild active bridge seal vector", Passed: a.Active.ActiveCount == 5 && a.Active.Primary == "GaugeScalarBoundaryStressSeal" && a.Active.Seals[0].Name == "GaugeScalarBoundaryStressSeal" && a.Active.Seals[1].Name == "HistoryLoopUnitSeal" && a.Active.Seals[2].Name == "OrientationBalanceSeal" && a.Active.Seals[3].Name == "ScalarProxyMatchingSeal" && a.Active.Seals[4].Name == "StrongBoundaryCorrectionSlot", Detail: FormatActive(a.Active)},
			{Name: "classify inactive/sealed lanes", Passed: len(a.Inactive.Lanes) == 5 && a.Inactive.FanoHitchinInactive && a.Inactive.HalfTraceInactive && a.Inactive.K7TraceInactive && a.Inactive.HodgeK7W7Inactive && a.Inactive.SplitG2Inactive, Detail: FormatInactive(a.Inactive)},
			{Name: "construct next-action ranking", Passed: len(a.Ranking.Actions) == 5 && a.Ranking.PrimaryPath == "RG/threshold transport refinement" && a.Ranking.SecondaryPath == "Scalar proxy-to-runtime matching theorem" && a.Ranking.Actions[4].Path == "K_7 boundary trace theorem" && !a.Ranking.Actions[4].Actionable && a.Ranking.K7BoundaryLow, Detail: FormatRanking(a.Ranking)},
			{Name: "pivot strategy back to transport", Passed: a.Strategic.StopFanoBoundaryLane && a.Strategic.ReturnToTransport && a.Strategic.BoundaryStressLive && a.Strategic.ScalarMatchingLive && a.Strategic.HistoryLoopLive && a.Strategic.FlavorOrientationLive && a.Strategic.K7BoundaryBlocked, Detail: FormatStrategic(a.Strategic)},
			{Name: "preserve closure/pivot firewalls", Passed: !a.Firewalls.ClaimsBoundaryStressDerived && !a.Firewalls.ClaimsScalarRGDerived && !a.Firewalls.ClaimsHiggsMass && !a.Firewalls.ClaimsGaugeUnification && !a.Firewalls.ClaimsFlavorDerived && !a.Firewalls.ClaimsPhysicalSpacetime && !a.Firewalls.ClaimsSplitG2 && !a.Firewalls.ClaimsSevenOver72Theorem && !a.Firewalls.ClaimsFanoBoundaryInterface && a.Firewalls.Verdict == StatusGate657Boundary, Detail: FormatFirewalls(a.Firewalls)},
		}
		notes := append(Statuses(), a.Truth)
		if !strings.Contains(a.Closure.Verdict, StatusFanoHitchinBoundaryClosed) {
			notes = append(notes, "WARNING_FANO_HITCHIN_ROUTE_NOT_CLOSED")
		}
		return theorem.Result{ID: id, Name: name, Layer: theorem.LayerBridge, Status: theorem.BridgeRequired, Checks: checks, Notes: notes}
	}}
}
