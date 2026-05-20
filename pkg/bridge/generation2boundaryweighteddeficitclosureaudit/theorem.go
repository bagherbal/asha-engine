package generation2boundaryweighteddeficitclosureaudit

import (
	"math"
	"strings"

	"github.com/bagherbal/asha-engine/pkg/theorem"
)

func Generation2BoundaryWeightedDeficitClosureAuditTheorem() theorem.Theorem {
	const id = AuditID
	const name = "Generation 2 boundary-weighted deficit closure audit"
	return theorem.Theorem{ID: id, Name: name, Layer: theorem.LayerBridge, Status: theorem.BridgeRequired, Verify: func() theorem.Result {
		a, err := BuildDefault()
		if err != nil {
			return theorem.Result{ID: id, Name: name, Layer: theorem.LayerBridge, Status: theorem.FailedRoute, Checks: []theorem.Check{{Name: "construct Gate626 boundary-weighted closure audit", Passed: false, Detail: err.Error()}}}
		}
		checks := []theorem.Check{
			{Name: "inherit Gate625 deficit closure without native promotion", Passed: a.Inherited.Verdict == StatusGate625Inherited && a.Inherited.Gate625ClosureSealDefined && !a.Inherited.Gate625NativeClosureTheorem && !a.Inherited.Gate625NativeScalarRGMatching && !a.Inherited.Gate625NativeFlavorOrientation, Detail: FormatInherited(a.Inherited)},
			{Name: "compute boundary split residual", Passed: a.BoundarySplit.Verdict == StatusBoundarySplitComputed && a.BoundarySplit.SplitPositive && a.BoundarySplit.ResidualInsideBoundarySplit && a.BoundarySplit.BoundaryStressLaneInherited, Detail: FormatBoundarySplit(a.BoundarySplit)},
			{Name: "audit 7/72 boundary weight candidate", Passed: a.WeightCandidate.Verdict == StatusSevenOverSeventyTwoCandidate && a.WeightCandidate.Expression == "7/72" && a.WeightCandidate.TypedOperands && !a.WeightCandidate.NativeSourceCertified && a.WeightCandidate.AbsoluteRatioResidual < 7e-7, Detail: FormatWeightCandidate(a.WeightCandidate)},
			{Name: "compute boundary-weighted closure", Passed: a.WeightedClosure.Verdict == StatusBoundaryWeightedClosure && a.WeightedClosure.BridgeOnly && a.WeightedClosure.ImprovesGate625 && a.WeightedClosure.AbsoluteResidual < 1e-9 && math.Abs(a.WeightedClosure.BoundaryWeight-7.0/72.0) < 1e-15, Detail: FormatWeightedClosure(a.WeightedClosure)},
			{Name: "rewrite scalar deficit with weighted scalar/gauge wound mixture", Passed: a.ScalarFormula.Verdict == StatusScalarFormulaComputed && !a.ScalarFormula.NativeScalarFormulaClaimed && math.Abs(a.ScalarFormula.KappaLambdaResidualExact-a.WeightedClosure.Residual) < 1e-15 && math.Abs(a.ScalarFormula.KappaLambdaResidualOrient) < 3e-6, Detail: FormatScalarFormula(a.ScalarFormula)},
			{Name: "compute boundary-weighted scalar prediction", Passed: a.ScalarPrediction.Verdict == StatusScalarFormulaComputed && a.ScalarPrediction.DiagnosticOnly && a.ScalarPrediction.ImprovesGate625Prediction && a.ScalarPrediction.BestResidual < 5e-12, Detail: FormatScalarPrediction(a.ScalarPrediction)},
			{Name: "compare residual scales against Gate625", Passed: a.ResidualScales.Verdict == StatusBoundaryWeightedClosure && a.ResidualScales.ClosureImprovementFactor > 100000 && a.ResidualScales.ScalarImprovementFactor > 100000, Detail: FormatResidualScales(a.ResidualScales)},
			{Name: "audit sign and role of scalar/gauge boundary mixture", Passed: a.SignRole.Verdict == StatusBoundaryWeightedClosure && !a.SignRole.NativeTheoremClaimed && a.SignRole.BoundaryWeight > 0 && a.SignRole.ScalarWeight > 0.9, Detail: FormatSignAndRole(a.SignRole)},
			{Name: "record missing native 7/72 and transport theorems", Passed: !a.NativeStatus.NativeSevenOverSeventyTwoSource && !a.NativeStatus.NativeGaugeScalarFlavorDeficitTransport && !a.NativeStatus.NativeBoundaryWeightedClosureTheorem && !a.NativeStatus.NativeScalarRGMatchingTheorem && !a.NativeStatus.NativeFlavorOrientationTheorem, Detail: FormatNativeStatus(a.NativeStatus)},
			{Name: "preserve Gate626 firewalls", Passed: !a.Firewalls.ClaimsHiggsMassDerived && !a.Firewalls.ClaimsScalarStability && !a.Firewalls.ClaimsKoideDerived && !a.Firewalls.ClaimsPMNSCKMDerived && !a.Firewalls.ClaimsGaugeUnification && !a.Firewalls.ClaimsNativeWeightTheorem && !a.Firewalls.ClaimsNativeTransportTheorem && !a.Firewalls.ClaimsEndpointDerivation, Detail: FormatFirewalls(a.Firewalls)},
		}
		notes := append(Statuses(), a.Truth)
		notes = append(notes, "Boundary-weighted closure: "+strings.TrimSpace(FormatWeightedClosure(a.WeightedClosure)))
		return theorem.Result{ID: id, Name: name, Layer: theorem.LayerBridge, Status: theorem.BridgeRequired, Checks: checks, Notes: notes}
	}}
}
