package generation2historyloopdeficitclosuretriangleaudit

import (
	"math"
	"strings"

	"github.com/bagherbal/asha-engine/pkg/theorem"
)

func Generation2HistoryLoopDeficitClosureTriangleAuditTheorem() theorem.Theorem {
	const id = AuditID
	const name = "Generation 2 history loop deficit closure triangle audit"
	return theorem.Theorem{ID: id, Name: name, Layer: theorem.LayerBridge, Status: theorem.BridgeRequired, Verify: func() theorem.Result {
		a, err := BuildDefault()
		if err != nil {
			return theorem.Result{ID: id, Name: name, Layer: theorem.LayerBridge, Status: theorem.FailedRoute, Checks: []theorem.Check{{Name: "construct Gate625 deficit closure audit", Passed: false, Detail: err.Error()}}}
		}
		checks := []theorem.Check{
			{Name: "inherit Gate624 history loop unit source-type audit", Passed: a.Inherited.Verdict == StatusGate624Inherited && a.Inherited.Gate624QuarterPhase && !a.Inherited.NativeHistoryLoopUnit && !a.Inherited.NativeHopfToScalar && !a.Inherited.NativeHopfToFlavor, Detail: FormatInherited(a.Inherited)},
			{Name: "define kappa_e and kappa_lambda as positive L-seal deficits", Passed: a.Kappas.Verdict == StatusKappasDefined && a.Kappas.BothPositive && a.Kappas.ScalarDeficitLarger, Detail: FormatKappas(a.Kappas)},
			{Name: "compute deficit closure against typed boundary candidates", Passed: a.ClosureTable.Verdict == StatusDeficitClosureComputed && len(a.ClosureTable.Rows) == 3 && a.ClosureTable.ClosesOnAbsLambda && a.ClosureTable.ClosestTarget == "|lambda(Lambda_12)|" && a.ClosureTable.ClosestRelative < 0.003, Detail: FormatClosureTable(a.ClosureTable)},
			{Name: "rewrite scalar deficit using boundary wound minus flavor deficit", Passed: a.ScalarFormula.Verdict == StatusClosureOnAbsLambda12 && a.ScalarFormula.BridgeOnly && math.Abs(a.ScalarFormula.ResidualExact-a.ClosureTable.ClosestResidual) < 1e-15 && math.Abs(a.ScalarFormula.ResidualOrient) < 1.5e-4, Detail: FormatScalarFormula(a.ScalarFormula)},
			{Name: "compute full scalar prediction from closure", Passed: a.ScalarPrediction.Verdict == StatusScalarPredictionComputed && len(a.ScalarPrediction.Rows) == 2 && a.ScalarPrediction.DiagnosticOnly && a.ScalarPrediction.ImprovesGate623RawLAnsatz && a.ScalarPrediction.BestResidual < 7e-7, Detail: FormatScalarPrediction(a.ScalarPrediction)},
			{Name: "compare residual scale against prior gates", Passed: a.ResidualScales.Verdict == StatusClosureSealDefined && a.ResidualScales.ClosureSharperThanRawScalarAnsatz && a.ResidualScales.ScalarImprovementFactor > 100, Detail: FormatResidualScales(a.ResidualScales)},
			{Name: "audit signs and roles of the closure triangle", Passed: a.SignRole.Verdict == StatusClosureSealDefined && a.SignRole.OpposedRGWoundSign && !a.SignRole.NativeTheoremClaimed, Detail: FormatSignAndRole(a.SignRole)},
			{Name: "record missing native closure theorems", Passed: !a.NativeStatus.NativeKappaClosureTheorem && !a.NativeStatus.NativeScalarRGMatchingTheorem && !a.NativeStatus.NativeFlavorOrientationTheorem && !a.NativeStatus.NativeLowScaleMatchingToHighScaleWoundLaw && !a.NativeStatus.NativeHistoryLoopDeficitClosureTheorem, Detail: FormatNativeStatus(a.NativeStatus)},
			{Name: "preserve Gate625 firewalls", Passed: !a.Firewalls.ClaimsHiggsMassDerived && !a.Firewalls.ClaimsScalarStability && !a.Firewalls.ClaimsKoideDerived && !a.Firewalls.ClaimsPMNSCKMDerived && !a.Firewalls.ClaimsGaugeUnification && !a.Firewalls.ClaimsNativeASHAClosure && !a.Firewalls.ClaimsNativeHistoryLoopUnit, Detail: FormatFirewalls(a.Firewalls)},
		}
		notes := append(Statuses(), a.Truth)
		notes = append(notes, "Deficit closure: "+strings.TrimSpace(FormatClosureTable(a.ClosureTable)))
		return theorem.Result{ID: id, Name: name, Layer: theorem.LayerBridge, Status: theorem.BridgeRequired, Checks: checks, Notes: notes}
	}}
}
