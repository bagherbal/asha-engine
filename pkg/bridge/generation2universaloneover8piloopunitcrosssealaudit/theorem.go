package generation2universaloneover8piloopunitcrosssealaudit

import (
	"math"
	"strings"

	"github.com/bagherbal/asha-engine/pkg/theorem"
)

func Generation2UniversalOneOver8PiLoopUnitCrossSealAuditTheorem() theorem.Theorem {
	const id = AuditID
	const name = "Generation 2 universal one-over-8pi loop unit cross-seal audit"
	return theorem.Theorem{ID: id, Name: name, Layer: theorem.LayerBridge, Status: theorem.BridgeRequired, Verify: func() theorem.Result {
		a, err := BuildDefault()
		if err != nil {
			return theorem.Result{ID: id, Name: name, Layer: theorem.LayerBridge, Status: theorem.FailedRoute, Checks: []theorem.Check{{Name: "construct Gate623 shared loop unit audit", Passed: false, Detail: err.Error()}}}
		}
		L := 1 / (8 * math.Pi)
		checks := []theorem.Check{
			{Name: "inherit Gate622 scalar loop match", Passed: a.ScalarInherited.Verdict == StatusGate622Inherited && a.ScalarInherited.RhoLambdaMatch > 0 && math.Abs(a.ScalarInherited.RhoLambdaMatch-L)/L < 0.05, Detail: FormatScalarInherited(a.ScalarInherited)},
			{Name: "inherit flavor loop unit structure", Passed: a.FlavorInherited.Verdict == StatusFlavorLoopInherited && a.FlavorInherited.EpsilonE > 0 && math.Abs(a.FlavorInherited.OrientationResidual) < 3e-6, Detail: FormatFlavorInherited(a.FlavorInherited)},
			{Name: "write shared loop-unit normal form", Passed: a.NormalForm.Verdict == StatusNormalFormWritten && a.NormalForm.ScalarKappaLambda > 0 && a.NormalForm.FlavorKappaE > 0, Detail: FormatNormalForm(a.NormalForm)},
			{Name: "compute scalar and flavor L units", Passed: a.Kappas.Verdict == StatusUnitsComputed && a.Kappas.KappaLambda > a.Kappas.KappaE && a.Kappas.ClosestName != "", Detail: FormatKappaComparison(a.Kappas)},
			{Name: "scalar L ansatz close to runtime lambda", Passed: a.ScalarQuality.Verdict == StatusScalarAnsatzClose && math.Abs(a.ScalarQuality.RelativeRuntimeResidual) < 0.002 && a.ScalarQuality.DiagnosticOnly, Detail: FormatScalarQuality(a.ScalarQuality)},
			{Name: "flavor orientation balance improves raw L", Passed: a.FlavorQuality.Verdict == StatusFlavorBalanceClose && a.FlavorQuality.ResidualImprovementFactor > 70, Detail: FormatFlavorQuality(a.FlavorQuality)},
			{Name: "opposite bridge roles audited", Passed: a.SignRole.Verdict == StatusAppearsInBoth && a.SignRole.FlavorUsesBelowL && a.SignRole.ScalarUsesAboveProxy && !a.SignRole.NativeTheoremClaimed, Detail: FormatSignAndRole(a.SignRole)},
			{Name: "define bridge-only history loop unit seal", Passed: a.CrossSealType.Verdict == StatusCrossSealBridgeOnly && a.CrossSealType.BridgeOnly, Detail: FormatCrossSealType(a.CrossSealType)},
			{Name: "record missing native cross-seal theorems", Passed: !a.NativeStatus.NativeOneOver8PiTheorem && !a.NativeStatus.NativeScalarMatchingTheorem && !a.NativeStatus.NativeKoideWallTheorem && !a.NativeStatus.NativeCrossSealTheorem && !a.NativeStatus.NativeOrientationBalance && !a.NativeStatus.NativeHiggsPoleTheorem, Detail: FormatNativeStatus(a.NativeStatus)},
			{Name: "preserve Gate623 firewalls", Passed: !a.Firewalls.ClaimsKoideDerived && !a.Firewalls.ClaimsHiggsMassDerived && !a.Firewalls.ClaimsScalarStability && !a.Firewalls.ClaimsPMNSCKMDerived && !a.Firewalls.ClaimsGaugeUnification && !a.Firewalls.ClaimsNativeLoopTheorem, Detail: FormatFirewalls(a.Firewalls)},
		}
		notes := append(Statuses(), a.Truth)
		notes = append(notes, "Shared L normal form: "+strings.TrimSpace(FormatNormalForm(a.NormalForm)))
		return theorem.Result{ID: id, Name: name, Layer: theorem.LayerBridge, Status: theorem.BridgeRequired, Checks: checks, Notes: notes}
	}}
}
