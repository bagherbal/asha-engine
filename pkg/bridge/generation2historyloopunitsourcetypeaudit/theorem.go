package generation2historyloopunitsourcetypeaudit

import (
	"math"
	"strings"

	"github.com/bagherbal/asha-engine/pkg/theorem"
)

func Generation2HistoryLoopUnitSourceTypeAuditTheorem() theorem.Theorem {
	const id = AuditID
	const name = "Generation 2 history loop unit source-type audit"
	return theorem.Theorem{ID: id, Name: name, Layer: theorem.LayerBridge, Status: theorem.BridgeRequired, Verify: func() theorem.Result {
		a, err := BuildDefault()
		if err != nil {
			return theorem.Result{ID: id, Name: name, Layer: theorem.LayerBridge, Status: theorem.FailedRoute, Checks: []theorem.Check{{Name: "construct Gate624 source-type audit", Passed: false, Detail: err.Error()}}}
		}
		L := 1 / (8 * math.Pi)
		checks := []theorem.Check{
			{Name: "inherit Gate623 HistoryLoopUnitSeal", Passed: a.Inherited.Verdict == StatusGate623Inherited && math.Abs(a.Inherited.LoopUnit-L) < 1e-18 && !a.Inherited.NativeLTheorem && !a.Inherited.NativeCrossSeal, Detail: FormatInherited(a.Inherited)},
			{Name: "type algebraic decompositions of L", Passed: a.Decompositions.Verdict == StatusDecompositionsTyped && len(a.Decompositions.Rows) == 5 && a.Decompositions.AllValuesMatch && a.Decompositions.AllRowsTyped && a.Decompositions.NoArbitrarySearch, Detail: FormatDecompositions(a.Decompositions)},
			{Name: "audit Hopf/circle phase source candidate", Passed: a.HopfPhase.Verdict == StatusHopfPhaseAudited && a.HopfPhase.Gate570HopfS7Certified && a.HopfPhase.Gate570ReebPhaseCertified && a.HopfPhase.Gate572CP3Certified && a.HopfPhase.CirclePhaseNormalization && a.HopfPhase.QuarterProjectionCandidate && !a.HopfPhase.QuarterProjectionCertified && !a.HopfPhase.MapToFlavorWallCertified && !a.HopfPhase.MapToScalarMatchingCertified && !a.HopfPhase.PhysicalTimeClaimed, Detail: FormatHopfPhase(a.HopfPhase)},
			{Name: "audit weak-quarter source candidate", Passed: a.WeakQuarter.Verdict == StatusWeakQuarterAudited && math.Abs(a.WeakQuarter.Factor-0.25) < 1e-18 && a.WeakQuarter.WeakNormalizationTyped && a.WeakQuarter.PMNSOverlapTyped && !a.WeakQuarter.NativeConnectionToL && !a.WeakQuarter.NativeWeakQuarterLoopLaw, Detail: FormatWeakQuarter(a.WeakQuarter)},
			{Name: "audit heat-kernel and loop-factor descendants", Passed: a.HeatKernel.Verdict == StatusHeatKernelAudited && math.Abs(a.HeatKernel.LoopUnit-L) < 1e-18 && a.HeatKernel.FourDLoopUnit > 0 && a.HeatKernel.BoundarySurfaceUnit > 0 && !a.HeatKernel.AnyCertifiedReduction, Detail: FormatHeatKernel(a.HeatKernel)},
			{Name: "audit scalar role and kappa_lambda source", Passed: a.ScalarRole.Verdict == StatusScalarFlavorAudited && math.Abs(a.ScalarRole.KappaLambda-0.0443230430960771) < 1e-14 && a.ScalarRole.ClosestName != "" && !a.ScalarRole.KappaSourceCertified, Detail: FormatScalarRole(a.ScalarRole)},
			{Name: "audit flavor orientation-corrected phase-wall role", Passed: a.FlavorRole.Verdict == StatusScalarFlavorAudited && a.FlavorRole.Classification == "orientation-corrected phase-wall loop unit" && math.Abs(a.FlavorRole.Residual) < 2e-7 && !a.FlavorRole.NativeDerived, Detail: FormatFlavorRole(a.FlavorRole)},
			{Name: "build cross-seal comparison table", Passed: a.CrossSeal.Verdict == StatusSharedHistoryLoopUnitSeal && len(a.CrossSeal.Rows) == 3 && a.CrossSeal.SharedLBridgeSeal && !a.CrossSeal.NativeCrossSeal, Detail: FormatCrossSeal(a.CrossSeal)},
			{Name: "record missing native source theorems", Passed: !a.NativeStatus.NativeLTheorem && !a.NativeStatus.NativeHopfToFlavorWallMap && !a.NativeStatus.NativeHopfToScalarMatchingMap && !a.NativeStatus.NativeHeatKernelToLReduction && !a.NativeStatus.NativeWeakQuarterLoopTheorem && !a.NativeStatus.NativeCrossSealOrientationLaw, Detail: FormatNativeStatus(a.NativeStatus)},
			{Name: "preserve Gate624 firewalls", Passed: !a.Firewalls.ClaimsKoideDerived && !a.Firewalls.ClaimsHiggsMassDerived && !a.Firewalls.ClaimsScalarStability && !a.Firewalls.ClaimsPMNSCKMDerived && !a.Firewalls.ClaimsGaugeUnification && !a.Firewalls.ClaimsNativeLoopTheorem && !a.Firewalls.ClaimsPhysicalTime, Detail: FormatFirewalls(a.Firewalls)},
		}
		notes := append(Statuses(), a.Truth)
		notes = append(notes, "Quarter phase candidate: "+strings.TrimSpace(FormatDecompositions(a.Decompositions)))
		return theorem.Result{ID: id, Name: name, Layer: theorem.LayerBridge, Status: theorem.BridgeRequired, Checks: checks, Notes: notes}
	}}
}
